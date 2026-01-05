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

package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/connector"
	contexts "gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/context"
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/handlers"
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"
	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/pkg/config"
)

type app interface {
	Cfg() *config.AppConfig
	Connector() *connector.Connector
	Ctx() *contexts.MonitoringEventCtx
}

type Service struct {
	app
}

func NewMonitoringEventService(app app) *Service {
	svc := &Service{app: app}
	return svc
}

// ------------------------------------------------------------------------------
func (s *Service) PostMonitoringEventSubscription(afId string, data *models.MonitoringEventSubscription, immediateReport *models.MonitoringEventReport) (string, int, error) {

	af := s.Ctx().GetAf(afId)
	if af == nil {
		af = s.Ctx().AddAf(afId)
	}

	af.Mu.Lock()
	defer af.Mu.Unlock()

	/* elaborate subscription here */
	if /* len(data.Msisdn) > 0 ||*/ len(data.ExternalId) > 0 || len(data.Supi) > 0 {

		var supi string
		var err error

		if len(data.Supi) > 0 {
			supi = data.Supi
			log.Printf("Using provided SUPI: %s", supi)
		} else {
			log.Printf("Looking up externalId=%s with afId=%s", data.ExternalId, afId)
			supi, err = s.Connector().LookupExternalId(afId, data.ExternalId)
			if err != nil {
				return "", http.StatusNotFound, fmt.Errorf("failed to lookup externalId %s: %w", data.ExternalId, err)
			}
			log.Printf("LookupExternalId returned SUPI: %s", supi)
		}

		/* get user info */
		log.Printf("Calling QueryUEInfo with SUPI: %s", supi)
		userInfo, err := s.Connector().QueryUEInfo(supi)
		if err != nil {
			return "", http.StatusInternalServerError, fmt.Errorf("failed to get current user state: %w", err)
		}

		*immediateReport, err = prepareImmediateReport(data.ExternalId, data.MonitoringType, userInfo)
		if err != nil {
			return "", http.StatusInternalServerError, fmt.Errorf("failed to prepare immediate report: %w", err)
		}

		if data.MaximumNumberOfReports == 1 {
			/* if only immediate report is requested, then return and do not create the subscription context*/
			return "", http.StatusOK, nil
		}

		sub := af.NewAfSubscription(supi, data)
		if sub == nil {
			return "", http.StatusBadRequest, fmt.Errorf("failed to validate monitoring event subscription")
		}
		loc := sub.GetLocation()
		eventType := sub.GetEventType()

		cnEvenType := mapNefTriggerToCoreNetworkEventTypes(eventType)
		if cnEvenType == "NOT_SUPPORTED" {
			return "", http.StatusNotImplemented, fmt.Errorf("monitoring event type %s is not supported", eventType)
		}

		/*subscribe to user info */
		subscription, err := s.Connector().SubscribeUserEvent(supi, cnEvenType)
		if err != nil {
			return "", http.StatusInternalServerError, fmt.Errorf("failed to subscribe to user info: %w", err)
		}

		/* create notification handler */
		notifHandler := handlers.NewNotificationHandler(loc, data.ExternalId, sub.GetNotificationUri(), subscription)

		/* assign event callback to the notification handler */
		switch eventType {
		case models.MonitoringTypeLocationReporting:
			notifHandler.SetEventCallback(handlers.HandleLocationReport)
		case models.MonitoringTypeUeReachability:
			notifHandler.SetEventCallback(handlers.HandleRegistrationReport)
		case models.MonitoringTypeLossOfConnectivity:
			notifHandler.SetEventCallback(handlers.HandleConnectivityReport)
		case models.MonitoringTypeDownlinkDataDeliveryStatus:
			notifHandler.SetEventCallback(handlers.HandleDDDSReport)
		case models.MonitoringTypePdnConnectivityStatus:
			notifHandler.SetEventCallback(handlers.HandlePdnStatusReport)
		default:
			return "", http.StatusNotImplemented, fmt.Errorf("monitoring event type %s is not supported", eventType)
		}
		notifHandler.Start()
		sub.SetNotificationHandler(notifHandler)

		return loc, http.StatusOK, nil

	} else if len(data.ExternalGroupId) > 0 {
		return "", http.StatusNotImplemented, fmt.Errorf("multiple or grouped UEs are not supported yet")
	} else {
		return "", http.StatusBadRequest, fmt.Errorf("no identifiers provided")
	}

}

// ------------------------------------------------------------------------------
func (s *Service) UpdateMonitoringEventSubscription(afId string, subId string, data *models.MonitoringEventSubscription) error {

	return nil
}

// ------------------------------------------------------------------------------
func (s *Service) DeleteMonitoringEventSubscription(afId string, subId string) (int, error) {
	af := s.Ctx().GetAf(afId)

	if af != nil {
		af.Mu.Lock()
		defer af.Mu.Unlock()

		sub := af.GetAfSubscription(subId)
		if sub != nil {
			notifHandler := sub.GetNotificationHandler()
			if notifHandler != nil {
				// stop the notification handler
				ok := notifHandler.Stop()
				if !ok {
					return http.StatusInternalServerError, fmt.Errorf("could not stop the notification handler")
				} else {
					err := af.DeleteAfscription(subId)
					if err != nil {
						return http.StatusInternalServerError, fmt.Errorf("could not stop the notification handler")
					}
					return http.StatusNoContent, nil
				}
			}
		}
	}
	return http.StatusNotFound, fmt.Errorf("could not find af/subId")
}

// ------------------------------------------------------------------------------
func (s *Service) GetMonitoringEventSubscription(afId string, subId string) (*models.MonitoringEventSubscription, int, error) {
	af := s.Ctx().GetAf(afId)

	if af != nil {
		af.Mu.Lock()
		defer af.Mu.Unlock()

		sub := af.GetAfSubscription(subId)
		if sub != nil {
			// use getter to return a copy of the subscription data
			// this is to avoid returning a pointer to the internal data structure
			// which could be modified later
			// and cause issues with the API consumers
			subData := sub.GetSubscriptionData()
			return &subData, http.StatusOK, nil
		}
	}
	return nil, http.StatusNotFound, fmt.Errorf("could not find af/subId")
}

// ------------------------------------------------------------------------------
func (s *Service) GetAllMonitoringEventSubscriptions(afId string) ([]*models.MonitoringEventSubscription, int, error) {
	af := s.Ctx().GetAf(afId)

	if af != nil {
		af.Mu.Lock()
		defer af.Mu.Unlock()

		var subList []*models.MonitoringEventSubscription
		subs := af.GetAfSubscriptions()
		for _, sub := range subs {
			// use getter to return a copy of the subscription data
			// this is to avoid returning a pointer to the internal data structure
			// which could be modified later
			// and cause issues with the API consumers
			subData := sub.GetSubscriptionData()
			subList = append(subList, &subData)
		}
		return subList, http.StatusOK, nil

	}
	return nil, http.StatusNotFound, fmt.Errorf("could not find af/subId")
}

// ------------------------------------------------------------------------------
func mapNefTriggerToCoreNetworkEventTypes(trigger models.MonitoringType) string {

	switch trigger {
	case models.MonitoringTypeUeReachability:
		return string(models.CORENETWORKEVENT_REACHABILITY_REPORT)
	case models.MonitoringTypeLocationReporting:
		return string(models.CORENETWORKEVENT_LOCATION_REPORT)
	case models.MonitoringTypeLossOfConnectivity:
		return string(models.CORENETWORKEVENT_LOSS_OF_CONNECTIVITY)
	case models.MonitoringTypeDownlinkDataDeliveryStatus:
		return string(models.CORENETWORKEVENT_DDDS)
	case models.MonitoringTypePdnConnectivityStatus:
		return string(models.CORENETWORKEVENT_PDN_CONNECTIVITY_STATUS)
	}
	return "NOT_SUPPORTED"

}

// generateMockGeographicArea creates a mock geographic area (polygon) based on cell ID
// This is needed because CoreSim only provides cell-based location (NCGI), not GPS coordinates
func generateMockGeographicArea(cellId string) *models.GeographicArea {
	// Mock base coordinates - different locations for different cell IDs
	// Using coordinates around Paris, France as default
	baseLat := 48.8566
	baseLon := 2.3522

	// Offset based on cell ID to create different locations
	// Parse last digit of cell ID for simple variation
	if len(cellId) > 0 {
		lastChar := cellId[len(cellId)-1]
		offset := float64(lastChar%10) * 0.01
		baseLat += offset
		baseLon += offset
	}

	// Create a small polygon (approximately 200m x 200m)
	// ~0.001 degrees is roughly 100 meters
	pointList := []models.GeographicalCoordinates{
		{Lat: baseLat + 0.001, Lon: baseLon + 0.001},
		{Lat: baseLat + 0.001, Lon: baseLon - 0.001},
		{Lat: baseLat - 0.001, Lon: baseLon - 0.001},
		{Lat: baseLat - 0.001, Lon: baseLon + 0.001},
	}

	return &models.GeographicArea{
		Shape:     models.SupportedGadShapesPOLYGON,
		PointList: pointList,
	}
}

func prepareImmediateReport(externalId string, eventType models.MonitoringType, ue *models.UeInfo) (models.MonitoringEventReport, error) {
	immediateReport := models.MonitoringEventReport{}

	immediateReport.ExternalId = &externalId
	immediateReport.MonitoringType = eventType
	immediateReport.EventTime = time.Now().Local()

	switch eventType {
	case models.MonitoringTypeLocationReporting:
		if ue.Location != nil {
			// Generate mock geographic area if not provided by CoreSim
			geoArea := ue.Location.GeographicArea
			if geoArea == nil {
				cellId := ue.Location.UserLocation.NrLocation.Ncgi.NrCellId
				geoArea = generateMockGeographicArea(cellId)
				log.Printf("Generated mock geographic area for cell %s", cellId)
			}

			immediateReport.LocationInfo = &models.LocationInfo{
				UserLocation:   &ue.Location.UserLocation,
				GeographicArea: geoArea,
				CellId:         ue.Location.UserLocation.NrLocation.Ncgi.NrCellId,
			}
			/*update nrLocation ageOfLocation*/
			immediateReport.LocationInfo.UserLocation.NrLocation.AgeOfLocationInformation = int32(time.Now().Unix() - ue.Location.TimeStamp)
			immediateReport.LocationInfo.UserLocation.NrLocation.UeLocationTimestamp = time.Unix(ue.Location.TimeStamp, 0)
		}

	case models.MonitoringTypeLossOfConnectivity:
		if ue.LossOfConnectivity != nil {
			reason := string(ue.LossOfConnectivity.LossOfConnectReason)
			immediateReport.LossOfConnectReason = &reason
		}

	case models.MonitoringTypeDownlinkDataDeliveryStatus:
		var oldestDdds *models.Ddds
		oldestDdds = nil
		oldestTs := int64(0)

		for _, Ddds := range ue.Ddds {
			if Ddds.TimeStamp > oldestTs {
				oldestTs = Ddds.TimeStamp
				oldestDdds = Ddds
			}
		}

		if oldestDdds != nil {
			immediateReport.PduSessInfo = &models.PduSessionInformation{
				Dnn:    *oldestDdds.Dnn,
				Snssai: *oldestDdds.Snssai,
				UeIpv4: *oldestDdds.UeIpv4Addr,
			}
			immediateReport.DddStatus = oldestDdds.DddStatus

		}

	case models.MonitoringTypePdnConnectivityStatus:
		immediateReport.PdnConnInfoList = &[]models.PdnConnectionInformation{}

		/* first add all the pdu sessions that are established */
		for pduId, pduSessEst := range ue.PduSessEst {
			if ue.PduSessRel[pduId] != nil && ue.PduSessRel[pduId].TimeStamp > pduSessEst.TimeStamp {
				/* this session is currently released */
				*immediateReport.PdnConnInfoList = append(*immediateReport.PdnConnInfoList, models.PdnConnectionInformation{
					Status:   models.PdnConnectionStatus_RELEASED,
					Apn:      *ue.PduSessRel[pduId].Dnn,
					Ipv4Addr: *ue.PduSessRel[pduId].Ipv4Addr,
					PdnType:  models.PdnType_IPV4, // TODO: remove hardcoded value
				})

			} else {
				*immediateReport.PdnConnInfoList = append(*immediateReport.PdnConnInfoList, models.PdnConnectionInformation{
					Status:   models.PdnConnectionStatus_CREATED,
					Apn:      *pduSessEst.Dnn,
					Ipv4Addr: *pduSessEst.AdIpv4Addr,
					PdnType:  models.PdnType_IPV4, // TODO: remove hardcoded value
				})
			}
		}

		/*now add all the released ones that do not have a correspective entry in established */
		for pduId, pduSessRel := range ue.PduSessRel {
			if ue.PduSessEst[pduId] == nil {
				*immediateReport.PdnConnInfoList = append(*immediateReport.PdnConnInfoList, models.PdnConnectionInformation{
					Status:   models.PdnConnectionStatus_RELEASED,
					Apn:      *pduSessRel.Dnn,
					Ipv4Addr: *pduSessRel.Ipv4Addr,
					PdnType:  models.PdnType_IPV4, // TODO: remove hardcoded value
				})

			}
		}

	default:
		return models.MonitoringEventReport{}, fmt.Errorf("monitoring event type %s is not supported", eventType)
	}
	return immediateReport, nil
}
