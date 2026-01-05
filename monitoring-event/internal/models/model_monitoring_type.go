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

/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.5
 */

package models

// MonitoringType
//LOSS_OF_CONNECTIVITY: The SCS/AS requests to be notified when the 3GPP network detects   that the UE is no longer reachable for signalling or user plane communication
//UE_REACHABILITY: The SCS/AS requests to be notified when the UE becomes reachable for   sending either SMS or downlink data to the UE
// LOCATION_REPORTING: The SCS/AS requests to be notified of the current location or   the last known location of the UE
// CHANGE_OF_IMSI_IMEI_ASSOCIATION: The SCS/AS requests to be notified when the association   of an ME (IMEI(SV)) that uses a specific subscription (IMSI) is changed
//ROAMING_STATUS: The SCS/AS queries the UE's current roaming status and requests to get   notified when the status changes
//COMMUNICATION_FAILURE: The SCS/AS requests to be notified of communication failure events
//AVAILABILITY_AFTER_DDN_FAILURE: The SCS/AS requests to be notified when the UE has become   available after a DDN failure
//NUMBER_OF_UES_IN_AN_AREA: The SCS/AS requests to be notified the number of UEs in a given   geographic area
//PDN_CONNECTIVITY_STATUS: The SCS/AS requests to be notified when the 3GPP network detects   that the UE’s PDN connection is set up or torn down
//DOWNLINK_DATA_DELIVERY_STATUS: The AF requests to be notified when the 3GPP network detects that the downlink data delivery status is changed.
//API_SUPPORT_CAPABILITY: The SCS/AS requests to be notified of the availability of support   of service APIs.
//NUM_OF_REGD_UES: The AF requests to be notified of the current number of registered UEs   for a network slice.
//NUM_OF_ESTD_PDU_SESSIONS: The AF requests to be notified of the current number of   established PDU Sessions for a network slice.
//AREA_OF_INTEREST: The SCS/AS requests to be notified when the UAV moves in or   out of the geographic area.
//GROUP_MEMBER_LIST_CHANGE: The AF requests to be notified of the changes to a group members   list.
//APPLICATION_START: The AF requests to be notified about the start of application traffic   has been detected.
//APPLICATION_STOP: The AF requests to be notified about the stop of application traffic   has been detected.

type MonitoringType string

const (
	MonitoringTypeLossOfConnectivity          MonitoringType = "LOSS_OF_CONNECTIVITY"
	MonitoringTypeUeReachability              MonitoringType = "UE_REACHABILITY"
	MonitoringTypeLocationReporting           MonitoringType = "LOCATION_REPORTING"
	MonitoringTypeChangeOfImsiImeiAssociation MonitoringType = "CHANGE_OF_IMSI_IMEI_ASSOCIATION"
	MonitoringTypeRoamingStatus               MonitoringType = "ROAMING_STATUS"
	MonitoringTypeCommunicationFailure        MonitoringType = "COMMUNICATION_FAILURE"
	MonitoringTypeAvailabilityAfterDdnFailure MonitoringType = "AVAILABILITY_AFTER_DDN_FAILURE"
	MonitoringTypeNumberOfUesInAnArea         MonitoringType = "NUMBER_OF_UES_IN_AN_AREA"
	MonitoringTypePdnConnectivityStatus       MonitoringType = "PDN_CONNECTIVITY_STATUS"
	MonitoringTypeDownlinkDataDeliveryStatus  MonitoringType = "DOWNLINK_DATA_DELIVERY_STATUS"
	MonitoringTypeApiSupportCapability        MonitoringType = "API_SUPPORT_CAPABILITY"
	MonitoringTypeNumOfRegdUes                MonitoringType = "NUM_OF_REGD_UES"
	MonitoringTypeNumOfEstdPduSessions        MonitoringType = "NUM_OF_ESTD_PDU_SESSIONS"
	MonitoringTypeAreaOfInterest              MonitoringType = "AREA_OF_INTEREST"
	MonitoringTypeGroupMemberListChange       MonitoringType = "GROUP_MEMBER_LIST_CHANGE"
	MonitoringTypeApplicationStart            MonitoringType = "APPLICATION_START"
	MonitoringTypeApplicationStop             MonitoringType = "APPLICATION_STOP"
)

// AssertMonitoringTypeRequired checks if the required fields are not zero-ed
func AssertMonitoringTypeRequired(obj MonitoringType) error {
	return nil
}

// AssertMonitoringTypeConstraints checks if the values respects the defined constraints
func AssertMonitoringTypeConstraints(obj MonitoringType) error {
	return nil
}
