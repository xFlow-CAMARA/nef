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

import (
	"time"
)

// MonitoringEventReport - Represents an event monitoring report.
type MonitoringEventReport struct {
	ImeiChange *AssociationType `json:"imeiChange,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId *string `json:"externalId,omitempty"`

	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`

	PduSessInfo *PduSessionInformation `json:"pduSessInfo,omitempty"`

	IdleStatusInfo *IdleStatusInfo `json:"idleStatusInfo,omitempty"`

	LocationInfo *LocationInfo `json:"locationInfo,omitempty"`

	LocFailureCause *LocationFailureCause `json:"locFailureCause,omitempty"`

	// If \"monitoringType\" is \"LOSS_OF_CONNECTIVITY\", this parameter shall be included if available to identify the reason why loss of connectivity is reported. Refer to 3GPP TS 29.336 clause 8.4.58.
	LossOfConnectReason *string `json:"lossOfConnectReason,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	UnavailPerDur *int32 `json:"unavailPerDur,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MaxUEAvailabilityTime *time.Time `json:"maxUEAvailabilityTime,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn *string `json:"msisdn,omitempty"`

	MonitoringType MonitoringType `json:"monitoringType"`

	UePerLocationReport *UePerLocationReport `json:"uePerLocationReport,omitempty"`

	PlmnId *PlmnId `json:"plmnId,omitempty"`

	ReachabilityType *ReachabilityType `json:"reachabilityType,omitempty"`

	// If \"monitoringType\" is \"ROAMING_STATUS\", this parameter shall be set to \"true\" if the new serving PLMN is different from the HPLMN. Set to false or omitted otherwise.
	RoamingStatus *bool `json:"roamingStatus,omitempty"`

	FailureCause *FailureCause `json:"failureCause,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	EventTime time.Time `json:"eventTime,omitempty"`

	PdnConnInfoList *[]PdnConnectionInformation `json:"pdnConnInfoList,omitempty"`

	DddStatus *DlDataDeliveryStatus `json:"dddStatus,omitempty"`

	DddTrafDescriptor *DddTrafficDescriptor `json:"dddTrafDescriptor,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MaxWaitTime *time.Time `json:"maxWaitTime,omitempty"`

	ApiCaps *[]ApiCapabilityInfo `json:"apiCaps,omitempty"`

	NSStatusInfo *SacEventStatus `json:"nSStatusInfo,omitempty"`

	AfServiceId *string `json:"afServiceId,omitempty"`

	// If \"monitoringType\" is \"AREA_OF_INTEREST\", this parameter may be included to identify the UAV.
	ServLevelDevId *string `json:"servLevelDevId,omitempty"`

	// If \"monitoringType\" is \"AREA_OF_INTEREST\", this parameter shall be set to true if the specified UAV is in the monitoring area. Set to false or omitted otherwise.
	UavPresInd *bool `json:"uavPresInd,omitempty"`

	GroupMembListChanges *GroupMembListChanges `json:"groupMembListChanges,omitempty"`
}

func AssertMonitoringEventReportRequired(obj MonitoringEventReport) error { return nil }

// AssertMonitoringEventReportConstraints checks if the values respects the defined constraints
func AssertMonitoringEventReportConstraints(obj MonitoringEventReport) error { return nil }
