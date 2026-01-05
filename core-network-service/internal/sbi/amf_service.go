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

	amf_client "gitlab.eurecom.fr/open-exposure/nef/core-network-service/internal/amfclient"
)

// ------------------------------------------------------------------------------
func storeAmfNotificationOnDB(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		log.Printf("Received AMF Event Notification")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %s", err.Error())
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		log.Printf("Received body: %s", string(body))
		var amfNotification *amf_client.AmfEventNotification
		err = json.Unmarshal(body, &amfNotification)
		if err != nil {
			log.Printf("could not unmarshal due to: %s", err.Error())
			http.Error(w, "Error unmarshaling JSON", http.StatusBadRequest)
			return
		}
		reportList, ok := amfNotification.GetReportListOk()
		if !ok {
			log.Printf("Error: ReportList not found in notification")
			http.Error(w, "Error in getting ReportList from AMF notification", http.StatusBadRequest)
			return
		}

		// store reports one by one
		log.Printf("Processing %d reports", len(reportList))
		for _, report := range reportList {
			user := report.GetSupi()
			if user == "" {
				log.Printf("Error: SUPI not found in report")
				http.Error(w, "supi not found in report, cannot create object id", http.StatusBadRequest)
				return
			}
			log.Printf("Processing report for SUPI: %s, Type: %s", user, report.GetType())

			err := doUpdateByReport(report)
			if err != nil {
				log.Printf("Error in doUpdateByReport: %s", err.Error())
				http.Error(w, "error in getUpdateByReport", http.StatusBadRequest)
				return
			}
			log.Printf("Successfully stored report for SUPI: %s", user)

		}
		w.WriteHeader(http.StatusOK)
		log.Printf("AMF notification processing completed successfully")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ------------------------------------------------------------------------------
// doUpdateByReport - Return update bson.D by report
func doUpdateByReport(report amf_client.AmfEventReport) error {
	var update []byte
	var err error
	b, err := report.MarshalJSON()
	if err == nil {
		log.Printf("%s\n", string(b))
	}

	switch report.GetType() {
	case amf_client.AMFEVENTTYPEANYOF_REGISTRATION_STATE_REPORT:
		update, err = getUpdateRegistration(report)
	case amf_client.AMFEVENTTYPEANYOF_LOCATION_REPORT:
		update, err = getUpdateLocation(report)
	case amf_client.AMFEVENTTYPEANYOF_LOSS_OF_CONNECTIVITY:
		update, err = getUpdateLossOfConnectivity(report)
	case amf_client.AMFEVENTTYPEANYOF_CONNECTIVITY_STATE_REPORT:
		update, err = getUpdateConnectivityStateReport(report)
	default:
		log.Printf("report type %s is not supported currently", string(report.GetType()))
		return errors.New("invalid report type")
	}
	if err != nil {
		log.Printf("could not parse amf report: %s", err.Error())
		return err
	}

	userKey := "user:" + report.GetSupi()
	path := "$." + string(report.GetType())

	exists, err := redisClient.Do("EXISTS", userKey).Int()
	if err != nil {
		return err
	}
	if exists == 0 {
		err = redisClient.Do("JSON.SET", userKey, "$", `{}`).Err()
		if err != nil {
			return fmt.Errorf("failed to initialize user root: %w", err)
		}
	}

	// Append last notification to user in redis
	err = redisClient.Do("JSON.SET", userKey, path, update).Err()
	if err != nil {
		log.Printf("Error in updating redisJson: %s", err.Error())
	}

	/* publish to broadcast channel */
	ueChannel := "user:" + report.GetSupi() + ":" + string(report.GetType())
	broadcastChannel := "broadcast:" + string(report.GetType())

	// You can publish the raw update, or wrap it with metadata
	message := fmt.Sprintf(`{"imsi":"%s","type":"%s","data":%s}`,
		report.GetSupi(), string(report.GetType()), update)

	// Publish to both channels
	redisClient.Publish(ueChannel, message)
	redisClient.Publish(broadcastChannel, message)

	return nil
}

// ------------------------------------------------------------------------------
// getUpdateRegistration - Create update bson.D in case of registration
func getUpdateRegistration(report amf_client.AmfEventReport) ([]byte, error) {
	rmInfoList, ok := report.GetRmInfoListOk()
	if !ok {
		return nil, errors.New("failed to get RmInfoList")
	}
	// for _, rmInfo := range rmInfoList {
	//	b, err := rmInfo.MarshalJSON()
	//	if err == nil {
	//       	log.Printf("%s\n", string(b))
	//       	}
	//	}

	timeStamp := time.Now().Unix()
	// TODO: fix (get rid of) the "RmStateAnyOf" field
	push := rmInfo{
		RmInfo:    rmInfoList[len(rmInfoList)-1],
		TimeStamp: timeStamp,
	}
	update, err := json.Marshal(push)
	return update, err
}

// ------------------------------------------------------------------------------
// getUpdateLocation - Create update bson.D in case of Location
func getUpdateLocation(report amf_client.AmfEventReport) ([]byte, error) {
	locationObj, ok := report.GetLocationOk()
	if !ok {
		return nil, errors.New("failed to get Location")
	}
	//log.Printf("%+v\n", *locationObj)
	timeStamp := time.Now().Unix()
	push := location{UserLocation: *locationObj, TimeStamp: timeStamp}
	update, err := json.Marshal(push)
	return update, err
}

// ------------------------------------------------------------------------------
// getUpdateLossOfConnectivity - Create update bson.D in case of Loss of connectivity
func getUpdateLossOfConnectivity(report amf_client.AmfEventReport) ([]byte, error) {
	lossOfConnectReasonObj, ok := report.GetLossOfConnectReasonOk()
	if !ok {
		return nil, errors.New("failed to get lossOfConnectReason")
	}
	//log.Printf("%+v\n", *lossOfConnectReasonObj.LossOfConnectivityReasonAnyOf)
	timeStamp := time.Now().Unix()
	push := lossOfConnectReason{
		LossOfConnectReason: *lossOfConnectReasonObj.LossOfConnectivityReasonAnyOf,
		TimeStamp:           timeStamp,
	}
	update, err := json.Marshal(push)
	return update, err
}

// ------------------------------------------------------------------------------
// getUpdateConnectivityStateReport - Create update bson.D in case of Connectivity Report
func getUpdateConnectivityStateReport(report amf_client.AmfEventReport) ([]byte, error) {
	cmInfoList, ok := report.GetCmInfoListOk()
	if !ok {
		return nil, errors.New("failed to get cmInfoList")
	}

	timeStamp := time.Now().Unix()
	// TODO: fix (get rid of) the "RmStateAnyOf" field
	push := cmInfo{
		CmInfo:    cmInfoList[len(cmInfoList)-1],
		TimeStamp: timeStamp,
	}

	update, err := json.Marshal(push)
	return update, err
}
