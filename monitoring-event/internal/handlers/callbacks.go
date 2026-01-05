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

package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"
)

func HandleLocationReport(subscriptionLocation string, userId string, notificationUri string, patch *models.UeInfoPatch) error {

	var jsonData []byte

	/* marshall interface into json */
	jsonData, err := json.Marshal(patch.Data)
	if err != nil {
		return fmt.Errorf("malformed core network event data: %s", err.Error())
	}

	locInfo := models.Location{}
	if err := json.Unmarshal(jsonData, &locInfo); err != nil {
		return fmt.Errorf("malformed core network event data: %s", err.Error())
	}

	nrLoc := locInfo.UserLocation.NrLocation

	report := models.MonitoringEventReport{
		ExternalId:     &patch.Imsi,
		MonitoringType: models.MonitoringTypeLocationReporting,
		PlmnId: &models.PlmnId{
			Mcc: nrLoc.Tai.PlmnId.Mcc,
			Mnc: nrLoc.Tai.PlmnId.Mnc,
		},
		LocationInfo: &models.LocationInfo{
			CellId:            nrLoc.Ncgi.NrCellId,
			AgeOfLocationInfo: int32(time.Now().Unix() - locInfo.TimeStamp),
			UserLocation:      &locInfo.UserLocation,
			GeographicArea:    locInfo.GeographicArea,
			PositionMethod:    models.PositioningMethodCellID,
		},
		EventTime: time.Unix(locInfo.TimeStamp, 0),
	}

	monitoringEvent := models.MonitoringNotification{
		Subscription: subscriptionLocation,
	}
	monitoringEvent.MonitoringEventReports = append(monitoringEvent.MonitoringEventReports, report)
	return sendNotification(notificationUri, monitoringEvent)
}

func HandleDDDSReport(subscriptionLocation string, userId string, notificationUri string, patch *models.UeInfoPatch) error {

	var jsonData []byte

	/* marshall interface into json */
	jsonData, err := json.Marshal(patch.Data)
	if err != nil {
		return fmt.Errorf("malformed core network event data: %s", err.Error())
	}

	ddds := models.Ddds{}
	if err := json.Unmarshal(jsonData, &ddds); err != nil {
		return fmt.Errorf("malformed core network event data: %s", err.Error())
	}

	var pduSessInfo *models.PduSessionInformation

	if ddds.Dnn != nil && ddds.Snssai != nil && ddds.UeIpv4Addr != nil {
		pduSessInfo = &models.PduSessionInformation{
			Snssai: *ddds.Snssai,
			Dnn:    *ddds.Dnn,
			UeIpv4: *ddds.UeIpv4Addr,
		}
	} else {
		pduSessInfo = nil
	}

	report := models.MonitoringEventReport{
		ExternalId:     &patch.Imsi,
		MonitoringType: models.MonitoringTypeDownlinkDataDeliveryStatus,
		DddStatus:      ddds.DddStatus,
		PduSessInfo:    pduSessInfo,
		EventTime:      time.Unix(ddds.TimeStamp, 0),
	}

	monitoringEvent := models.MonitoringNotification{
		Subscription: subscriptionLocation,
	}
	monitoringEvent.MonitoringEventReports = append(monitoringEvent.MonitoringEventReports, report)
	return sendNotification(notificationUri, monitoringEvent)
}

func HandleRegistrationReport(subscriptionLocation string, userId string, notificationUri string, patch *models.UeInfoPatch) error {
	log.Printf("handleRegistrationReport: notificationUri=%s, identity=%s, patch=%+v", notificationUri, patch.Imsi, patch)
	return nil
}

func HandleConnectivityReport(subscriptionLocation string, userId string, notificationUri string, patch *models.UeInfoPatch) error {
	log.Printf("handleConnectivityReport: notificationUri=%s, identity=%s, patch=%+v", notificationUri, patch.Imsi, patch)
	return nil
}

func HandlePdnStatusReport(subscriptionLocation string, userId string, notificationUri string, patch *models.UeInfoPatch) error {
	var jsonData []byte

	/* marshall interface into json */
	jsonData, err := json.Marshal(patch.Data)
	if err != nil {
		return fmt.Errorf("malformed core network event data: %s", err.Error())
	}

	var pduSessInfo *models.PduSessionInformation
	var status models.PdnConnectionStatus
	var eventTime time.Time

	if patch.Type == string(models.CORENETWORKEVENT_PDU_SES_EST) {
		pdu_est := models.PduSesEst{}
		if err := json.Unmarshal(jsonData, &pdu_est); err != nil {
			return fmt.Errorf("malformed core network event data: %s", err.Error())
		}
		if pdu_est.Dnn != nil && pdu_est.Snssai != nil && pdu_est.AdIpv4Addr != nil {
			pduSessInfo = &models.PduSessionInformation{
				Snssai: *pdu_est.Snssai,
				Dnn:    *pdu_est.Dnn,
				UeIpv4: *pdu_est.AdIpv4Addr,
			}
			status = models.PdnConnectionStatus_CREATED
			eventTime = time.Unix(pdu_est.TimeStamp, 0)

		} else {
			return fmt.Errorf("malformed core network event data: %s", err.Error())
		}

	} else {
		pdu_rel := models.PduSesRel{}
		if err := json.Unmarshal(jsonData, &pdu_rel); err != nil {
			return fmt.Errorf("malformed core network event data: %s", err.Error())
		}
		if pdu_rel.Dnn != nil && pdu_rel.Snssai != nil && pdu_rel.Ipv4Addr != nil {
			pduSessInfo = &models.PduSessionInformation{
				Snssai: *pdu_rel.Snssai,
				Dnn:    *pdu_rel.Dnn,
				UeIpv4: *pdu_rel.Ipv4Addr,
			}
			status = models.PdnConnectionStatus_RELEASED
			eventTime = time.Unix(pdu_rel.TimeStamp, 0)

		} else {
			return fmt.Errorf("malformed core network event data: %s", err.Error())
		}
	}

	pdnInfo := models.PdnConnectionInformation{
		Status:   status,
		Apn:      pduSessInfo.Dnn,
		PdnType:  models.PdnType_IPV4,
		Ipv4Addr: pduSessInfo.UeIpv4,
	}

	report := models.MonitoringEventReport{
		ExternalId:      &patch.Imsi,
		MonitoringType:  models.MonitoringTypePdnConnectivityStatus,
		PduSessInfo:     pduSessInfo,
		EventTime:       eventTime,
		PdnConnInfoList: &[]models.PdnConnectionInformation{pdnInfo},
	}

	monitoringEvent := models.MonitoringNotification{
		Subscription: subscriptionLocation,
	}
	monitoringEvent.MonitoringEventReports = append(monitoringEvent.MonitoringEventReports, report)
	return sendNotification(notificationUri, monitoringEvent)
}

func sendNotification(notificationUri string, v any) error {
	log.Printf("sendNotification: notificationUri=%s, data=%+v", notificationUri, v)
	bData, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal notification data: %w", err)
	}

	r, _ := http.NewRequest("POST", notificationUri, bytes.NewBuffer(bData))
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Printf("could not close response body correctly")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification, status code: %d", resp.StatusCode)
	}
	return err
}
