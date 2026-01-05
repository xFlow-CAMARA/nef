// Copyright 2025 EURECOM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Contributors:
//   Giulio CAROTA
//   Thomas DU
//   Adlen KSENTINI

package sbi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	smf_client "gitlab.eurecom.fr/open-exposure/nef/core-network-service/internal/smfclient"
)

// ------------------------------------------------------------------------------
func storeSmfNotificationOnDB(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		log.Printf("Received SMF Event Notification")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error: %s", err.Error())

			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		var smfNotification *smf_client.NsmfEventExposureNotification
		err = json.Unmarshal(body, &smfNotification)
		if err != nil {
			log.Printf("error: %s", err.Error())
			http.Error(w, "Error unmarshaling JSON", http.StatusBadRequest)
			return
		}
		notifList, ok := smfNotification.GetEventNotifsOk()
		if !ok {
			log.Printf("Error in getting EventNotifs from SMF notification")

			http.Error(w, "Error in getting EventNotifs from SMF notification", http.StatusBadRequest)
			return
		}
		// store reports one by one
		for _, notif := range notifList {

			user := notif.GetSupi()
			if user == "" {
				log.Print("supi not found in notification, cannot create object id")

				http.Error(w, "supi not found in notification, cannot create object id", http.StatusBadRequest)
				return
			}

			//parse the received notification
			err := doUpdateByNotif(notif)
			if err != nil {
				log.Print("error in getUpdateByNotif")

				http.Error(w, "error in getUpdateByNotif", http.StatusBadRequest)
				return
			}

			/*userKey := "user:" + user
			// Append last notification to user in redis
			err = redisClient.Do("JSON.ARRAPPEND", userKey, "$.smf_events", update).Err()
			if err != nil {
				// f key or $.events doesn't exist, create it with the array
				payload := map[string][]byte{
					"smf_events": update,
				}
				jsonData, _ := json.Marshal(payload)
				redisClient.Do("JSON.SET", userKey, "$", string(jsonData))
			}

			// Trigger broadcast message for event notifications
			*/
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ------------------------------------------------------------------------------
// getUpdateByNotif - Return update bson.D by notif
func doUpdateByNotif(notif smf_client.EventNotification) error {
	var update []byte
	var err error
	b, err := notif.MarshalJSON()
	if err == nil {
		log.Printf("%s\n", string(b))
	}

	var updateId int32
	updateId = -1

	// TODO: implement other report types
	switch notif.GetEvent() {
	case smf_client.SMFEVENTANYOF_PDU_SES_EST:
		update, updateId, err = getUpdatePDU_SES_EST(notif)
	case smf_client.SMFEVENTANYOF_UE_IP_CH:
		update, updateId, err = getUpdateUE_IP_CH(notif)
	case smf_client.SMFEVENTANYOF_PLMN_CH:
		update, err = getUpdatePLMN_CH(notif)
	case smf_client.SMFEVENTANYOF_DDDS:
		update, updateId, err = getUpdateDDDS(notif)
	case smf_client.SMFEVENTANYOF_PDU_SES_REL:
		update, updateId, err = getUpdatePDU_SES_REL(notif)
	case smf_client.SMFEVENTANYOF_QOS_MON:
		update, err = getUpdateQOS_MON(notif)
	default:
		log.Printf("notif event %s is not supported currently",
			string(notif.GetEvent()))
		return errors.New("invalid notif event")
	}
	if err != nil {
		return err
	}
	userKey := "user:" + notif.GetSupi()
	eventType := string(notif.GetEvent())
	eventPath := fmt.Sprintf("$.%s", eventType)

	// Step 1: Ensure root JSON object exists
	exists, err := redisClient.Exists(userKey).Result()
	if err != nil {
		return fmt.Errorf("redis EXISTS failed: %w", err)
	}
	if exists == 0 {
		if err := redisClient.Do("JSON.SET", userKey, "$", `{}`).Err(); err != nil {
			return fmt.Errorf("failed to initialize user root: %w", err)
		}
	}

	// Step 2: Ensure intermediate event type object exists
	val, err := redisClient.Do("JSON.GET", userKey, eventPath).String()
	if err != nil {
		return fmt.Errorf("failed to read event path: %w", err)
	}
	if val == "[]" || val == "{}" {
		log.Print("Setting event path object")
		if err := redisClient.Do("JSON.SET", userKey, eventPath, `{}`).Err(); err != nil {
			return fmt.Errorf("failed to initialize event type: %w", err)
		}
	}

	// Step 3: Compose final path with update ID (if applicable)
	finalPath := eventPath
	if updateId >= 0 {
		finalPath = fmt.Sprintf("%s.%d", eventPath, updateId)
	}

	// Step 4: Marshal update and set it
	if err != nil {
		return fmt.Errorf("failed to marshal update: %w", err)
	}

	err = redisClient.Do("JSON.SET", userKey, finalPath, update).Err()
	if err != nil {
		log.Printf("Error in updating Redis JSON: %s", err.Error())
		return err
	}

	/* publish to broadcast channel */
	// to semplify upper layers, we merge pdu sess reletated events into one
	if eventType == "PDU_SES_EST" || eventType == "PDU_SES_REL" {
		eventType = "PDN_CONNECTIVITY_STATUS"
	}

	broadcastChannel := "broadcast:" + eventType

	ueChannel := "user:" + notif.GetSupi() + ":" + eventType

	// You can publish the raw update, or wrap it with metadata
	message := fmt.Sprintf(`{"imsi":"%s","type":"%s","data":%s}`,
		notif.GetSupi(), string(notif.GetEvent()), update)

	// Publish to both channels
	redisClient.Publish(ueChannel, message)
	redisClient.Publish(broadcastChannel, message)

	return nil
}

// ----------------------------------------------------------------------------------------------------------------
// getUpdatePDU_SES_EST - Create update bson.D in case of PDU SESS EST
func getUpdatePDU_SES_EST(notif smf_client.EventNotification) ([]byte, int32, error) {
	adIpv4Addr, ok := notif.GetIpv4AddrOk()
	if !ok {
		return nil, -1, errors.New("failed to get AdIpv4Addr")
	}
	dnn, ok := notif.GetDnnOk()
	if !ok {
		return nil, -1, errors.New("failed to get Dnn")
	}
	pduSeId, ok := notif.GetPduSeIdOk()
	if !ok {
		return nil, -1, errors.New("failed to get PduSeId")
	}
	pduSessType, ok := notif.GetPduSessTypeOk()
	if !ok {
		return nil, -1, errors.New("failed to get PduSessType")
	}
	snssai, ok := notif.GetSnssaiOk()
	if !ok {
		return nil, -1, errors.New("failed to get Snssai")
	}
	timeStamp := time.Now().Unix()
	push := pduSesEst{
		AdIpv4Addr:  adIpv4Addr,
		Dnn:         dnn,
		PduSeId:     pduSeId,
		PduSessType: pduSessType,
		Snssai:      snssai,
		TimeStamp:   timeStamp,
	}
	update, err := json.Marshal(push)
	return update, *pduSeId, err
}

// ----------------------------------------------------------------------------------------------------------------
// getUpdateUE_IP_CH - Create update bson.D in case of UE IP CH
func getUpdateUE_IP_CH(notif smf_client.EventNotification) ([]byte, int32, error) {
	adIpv4Addr, ok := notif.GetAdIpv4AddrOk()
	if !ok {
		return nil, -1, errors.New("failed to get AdIpv4Addr")
	}
	pduSeId, ok := notif.GetPduSeIdOk()
	if !ok {
		return nil, -1, errors.New("failed to get PduSeId")
	}
	timeStamp := time.Now().Unix()
	push := ueIpCh{
		AdIpv4Addr: adIpv4Addr,
		PduSeId:    pduSeId,
		TimeStamp:  timeStamp,
	}
	update, err := json.Marshal(push)
	return update, *pduSeId, err
}

// ----------------------------------------------------------------------------------------------------------------
// getUpdatePLMN_CH - Create update bson.D in case of PLMN CH
func getUpdatePLMN_CH(notif smf_client.EventNotification) ([]byte, error) {
	return nil, errors.New("getUpdatePLMN_CH not implemented")
}

// ----------------------------------------------------------------------------------------------------------------
// getUpdateDDDS - Create update bson.D in case of DDDs
func getUpdateDDDS(notif smf_client.EventNotification) ([]byte, int32, error) {
	dddStatus, ok := notif.GetDddStatusOk()
	if !ok {
		return nil, -1, errors.New("failed to get DddStatus")
	}
	pduSeId, ok := notif.GetPduSeIdOk()
	if !ok {
		return nil, -1, errors.New("failed to get PduSeId")
	}

	ueIpv4 := notif.Ipv4Addr
	dnn := notif.Dnn
	snssai := notif.Snssai

	timeStamp := time.Now().Unix()
	push := ddds{DddStatus: dddStatus,
		PduSeId:    pduSeId,
		TimeStamp:  timeStamp,
		UeIpv4Addr: ueIpv4,
		Dnn:        dnn,
		Snssai:     snssai,
	}
	update, err := json.Marshal(push)
	return update, *pduSeId, err
}

// ----------------------------------------------------------------------------------------------------------------
// getUpdatePDU_SES_REL - Create update bson.D in case of PDU SES REL
func getUpdatePDU_SES_REL(notif smf_client.EventNotification) ([]byte, int32, error) {
	ipv4Addr, ok := notif.GetIpv4AddrOk()
	if !ok {
		return nil, -1, errors.New("failed to get AdIpv4Addr")
	}
	dnn, ok := notif.GetDnnOk()
	if !ok {
		return nil, -1, errors.New("failed to get Dnn")
	}
	pduSeId, ok := notif.GetPduSeIdOk()
	if !ok {
		return nil, -1, errors.New("failed to get PduSeId")
	}
	pduSessType, ok := notif.GetPduSessTypeOk()
	if !ok {
		return nil, -1, errors.New("failed to get PduSessType")
	}
	snssai, ok := notif.GetSnssaiOk()
	if !ok {
		return nil, -1, errors.New("failed to get Snssai")
	}
	timeStamp := time.Now().Unix()
	push := pduSesRel{
		Ipv4Addr:    ipv4Addr,
		Dnn:         dnn,
		PduSeId:     pduSeId,
		PduSessType: pduSessType,
		Snssai:      snssai,
		TimeStamp:   timeStamp,
	}
	update, err := json.Marshal(push)
	return update, *pduSeId, err
}

// ----------------------------------------------------------------------------------------------------------------
// getUpdatePDU_SES_REL - Create update bson.D in case of QoS MON
func getUpdateQOS_MON(notif smf_client.EventNotification) ([]byte, error) {
	pduSeId, ok := notif.GetPduSeIdOk()
	if !ok {
		return nil, errors.New("failed to get PduSeId")
	}
	timeStamp := time.Now().Unix()
	// include "customized_data"
	push := qosMon{
		Customized_data: notif.CustomizedData,
		PduSeId:         pduSeId,
		TimeStamp:       timeStamp,
	}
	update, err := json.Marshal(push)
	return update, err
}
