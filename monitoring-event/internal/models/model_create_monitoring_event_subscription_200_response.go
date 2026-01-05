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
	"errors"
	"time"
)

type CreateMonitoringEventSubscription200Response struct {
	ImeiChange AssociationType `json:"imeiChange,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId,omitempty"`

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	PduSessInfo *PduSessionInformation `json:"pduSessInfo,omitempty"`

	IdleStatusInfo IdleStatusInfo `json:"idleStatusInfo,omitempty"`

	LocationInfo LocationInfo `json:"locationInfo,omitempty"`

	LocFailureCause LocationFailureCause `json:"locFailureCause,omitempty"`

	// If \"monitoringType\" is \"LOSS_OF_CONNECTIVITY\", this parameter shall be included if available to identify the reason why loss of connectivity is reported. Refer to 3GPP TS 29.336 clause 8.4.58.
	LossOfConnectReason int32 `json:"lossOfConnectReason,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	UnavailPerDur int32 `json:"unavailPerDur,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MaxUEAvailabilityTime time.Time `json:"maxUEAvailabilityTime,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn string `json:"msisdn,omitempty"`

	MonitoringType MonitoringType `json:"monitoringType"`

	UePerLocationReport UePerLocationReport `json:"uePerLocationReport,omitempty"`

	PlmnId PlmnId `json:"plmnId,omitempty"`

	ReachabilityType ReachabilityType `json:"reachabilityType,omitempty"`

	// If \"monitoringType\" is \"ROAMING_STATUS\", this parameter shall be set to \"true\" if the new serving PLMN is different from the HPLMN. Set to false or omitted otherwise.
	RoamingStatus bool `json:"roamingStatus,omitempty"`

	FailureCause FailureCause `json:"failureCause,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	EventTime time.Time `json:"eventTime,omitempty"`

	PdnConnInfoList []PdnConnectionInformation `json:"pdnConnInfoList,omitempty"`

	DddStatus DlDataDeliveryStatus `json:"dddStatus,omitempty"`

	DddTrafDescriptor DddTrafficDescriptor `json:"dddTrafDescriptor,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MaxWaitTime time.Time `json:"maxWaitTime,omitempty"`

	ApiCaps []ApiCapabilityInfo `json:"apiCaps,omitempty"`

	NSStatusInfo SacEventStatus `json:"nSStatusInfo,omitempty"`

	AfServiceId string `json:"afServiceId,omitempty"`

	// If \"monitoringType\" is \"AREA_OF_INTEREST\", this parameter may be included to identify the UAV.
	ServLevelDevId string `json:"servLevelDevId,omitempty"`

	// If \"monitoringType\" is \"AREA_OF_INTEREST\", this parameter shall be set to true if the specified UAV is in the monitoring area. Set to false or omitted otherwise.
	UavPresInd bool `json:"uavPresInd,omitempty"`

	GroupMembListChanges *GroupMembListChanges `json:"groupMembListChanges,omitempty"`

	MonitoringEventReports []MonitoringEventReport `json:"monitoringEventReports"`
}

// AssertCreateMonitoringEventSubscription200ResponseRequired checks if the required fields are not zero-ed
func AssertCreateMonitoringEventSubscription200ResponseRequired(obj CreateMonitoringEventSubscription200Response) error {
	elements := map[string]interface{}{
		"monitoringType":         obj.MonitoringType,
		"monitoringEventReports": obj.MonitoringEventReports,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAssociationTypeRequired(obj.ImeiChange); err != nil {
		return err
	}
	if obj.PduSessInfo != nil {
		if err := AssertPduSessionInformationRequired(*obj.PduSessInfo); err != nil {
			return err
		}
	}
	if err := AssertIdleStatusInfoRequired(obj.IdleStatusInfo); err != nil {
		return err
	}
	if err := AssertLocationInfoRequired(obj.LocationInfo); err != nil {
		return err
	}
	if err := AssertLocationFailureCauseRequired(obj.LocFailureCause); err != nil {
		return err
	}
	if err := AssertMonitoringTypeRequired(obj.MonitoringType); err != nil {
		return err
	}
	if err := AssertUePerLocationReportRequired(obj.UePerLocationReport); err != nil {
		return err
	}
	if err := AssertPlmnIdRequired(obj.PlmnId); err != nil {
		return err
	}
	if err := AssertReachabilityTypeRequired(obj.ReachabilityType); err != nil {
		return err
	}
	if err := AssertFailureCauseRequired(obj.FailureCause); err != nil {
		return err
	}
	for _, el := range obj.PdnConnInfoList {
		if err := AssertPdnConnectionInformationRequired(el); err != nil {
			return err
		}
	}
	if err := AssertDlDataDeliveryStatusRequired(obj.DddStatus); err != nil {
		return err
	}
	if err := AssertDddTrafficDescriptorRequired(obj.DddTrafDescriptor); err != nil {
		return err
	}
	for _, el := range obj.ApiCaps {
		if err := AssertApiCapabilityInfoRequired(el); err != nil {
			return err
		}
	}
	if err := AssertSacEventStatusRequired(obj.NSStatusInfo); err != nil {
		return err
	}
	if obj.GroupMembListChanges != nil {
		if err := AssertGroupMembListChangesRequired(*obj.GroupMembListChanges); err != nil {
			return err
		}
	}
	for _, el := range obj.MonitoringEventReports {
		if err := AssertMonitoringEventReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCreateMonitoringEventSubscription200ResponseConstraints checks if the values respects the defined constraints
func AssertCreateMonitoringEventSubscription200ResponseConstraints(obj CreateMonitoringEventSubscription200Response) error {
	if err := AssertAssociationTypeConstraints(obj.ImeiChange); err != nil {
		return err
	}
	if obj.PduSessInfo != nil {
		if err := AssertPduSessionInformationConstraints(*obj.PduSessInfo); err != nil {
			return err
		}
	}
	if err := AssertIdleStatusInfoConstraints(obj.IdleStatusInfo); err != nil {
		return err
	}
	if err := AssertLocationInfoConstraints(obj.LocationInfo); err != nil {
		return err
	}
	if err := AssertLocationFailureCauseConstraints(obj.LocFailureCause); err != nil {
		return err
	}
	if obj.UnavailPerDur < 0 {
		return &ParsingError{Param: "UnavailPerDur", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertMonitoringTypeConstraints(obj.MonitoringType); err != nil {
		return err
	}
	if err := AssertUePerLocationReportConstraints(obj.UePerLocationReport); err != nil {
		return err
	}
	if err := AssertPlmnIdConstraints(obj.PlmnId); err != nil {
		return err
	}
	if err := AssertReachabilityTypeConstraints(obj.ReachabilityType); err != nil {
		return err
	}
	if err := AssertFailureCauseConstraints(obj.FailureCause); err != nil {
		return err
	}
	for _, el := range obj.PdnConnInfoList {
		if err := AssertPdnConnectionInformationConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertDlDataDeliveryStatusConstraints(obj.DddStatus); err != nil {
		return err
	}
	if err := AssertDddTrafficDescriptorConstraints(obj.DddTrafDescriptor); err != nil {
		return err
	}
	for _, el := range obj.ApiCaps {
		if err := AssertApiCapabilityInfoConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertSacEventStatusConstraints(obj.NSStatusInfo); err != nil {
		return err
	}
	if obj.GroupMembListChanges != nil {
		if err := AssertGroupMembListChangesConstraints(*obj.GroupMembListChanges); err != nil {
			return err
		}
	}
	for _, el := range obj.MonitoringEventReports {
		if err := AssertMonitoringEventReportConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
