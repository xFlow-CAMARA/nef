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
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.5
 */

package models

import (
	"errors"
)

// AsSessionWithQoSSubscriptionPatch - Represents parameters to modify an AS session with specific QoS subscription.
type AsSessionWithQoSSubscriptionPatch struct {

	// Identifies the external Application Identifier.
	ExterAppId string `json:"exterAppId,omitempty"`

	// Describe the IP data flow which requires QoS.
	FlowInfo []FlowInfo `json:"flowInfo,omitempty"`

	// Identifies Ethernet packet flows.
	EthFlowInfo []EthFlowDescription `json:"ethFlowInfo,omitempty"`

	// Identifies the Ethernet flows which require QoS. Each Ethernet flow consists of a flow idenifer and the corresponding UL and/or DL flows.
	EnEthFlowInfo []EthFlowInfo `json:"enEthFlowInfo,omitempty"`

	// Identifies the list of UE address.
	ListUeAddrs []UeAddInfo `json:"listUeAddrs,omitempty"`

	// Pre-defined QoS reference
	QosReference string `json:"qosReference,omitempty"`

	// Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority.
	AltQoSReferences []string `json:"altQoSReferences,omitempty"`

	// Identifies an ordered list of alternative service requirements that include individual QoS parameter sets. The lower the index of the array for a given entry, the higher the priority.
	AltQosReqs []AlternativeServiceRequirementsData `json:"altQosReqs,omitempty"`

	// Indicates whether the QoS flow parameters signalling to the UE when the SMF is notified by the NG-RAN of changes in the fulfilled QoS situation is disabled (true) or not (false). The fulfilled situation is either the QoS profile or an Alternative QoS Profile.
	DisUeNotif bool `json:"disUeNotif,omitempty"`

	UsageThreshold *UsageThresholdRm `json:"usageThreshold,omitempty"`

	QosMonInfo QosMonitoringInformationRm `json:"qosMonInfo,omitempty"`

	PdvMon QosMonitoringInformationRm `json:"pdvMon,omitempty"`

	// Indicates whether the direct event notification is requested (true) or not (false) for the provided QoS monitoring parameter(s).
	DirectNotifInd bool `json:"directNotifInd,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty"`

	TscQosReq TscQosRequirementRm `json:"tscQosReq,omitempty"`

	L4sInd UplinkDownlinkSupport `json:"l4sInd,omitempty"`

	TempInValidity TemporalInValidity `json:"tempInValidity,omitempty"`

	// Represents the updated list of user plane event(s) to which the SCS/AS requests to subscribe to.
	Events []UserPlaneEvent `json:"events,omitempty"`

	// Contains media component data for a single-modal data flow(s). The key of the map is the medCompN attribute.
	MultiModDatFlows map[string]AsSessionMediaComponentRm `json:"multiModDatFlows,omitempty"`

	PduSetQosDl *PduSetQosPara `json:"pduSetQosDl,omitempty"`

	PduSetQosUl *PduSetQosPara `json:"pduSetQosUl,omitempty"`

	// Indicates the service data flow needs to meet the Round-Trip (RT) latency requirement of the service, when it is included and set to \"true\". The default value is \"false\" if omitted.
	RTLatencyInd bool `json:"rTLatencyInd,omitempty"`

	ProtoDescUl ProtocolDescription `json:"protoDescUl,omitempty"`

	ProtoDescDl ProtocolDescription `json:"protoDescDl,omitempty"`

	// This data type is defined in the same way as the \"DurationMillisec\" data type, but with the OpenAPI nullable property set to true.
	PeriodUl int32 `json:"periodUl,omitempty"`

	// This data type is defined in the same way as the \"DurationMillisec\" data type, but with the OpenAPI nullable property set to true.
	PeriodDl int32 `json:"periodDl,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	QosDuration *int32 `json:"qosDuration,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	QosInactInt *int32 `json:"qosInactInt,omitempty"`

	RttMon QosMonitoringInformationRm `json:"rttMon,omitempty"`

	QosMonDatRate QosMonitoringInformationRm `json:"qosMonDatRate,omitempty"`

	// This data type is defined in the same way as the 'AverWindow' data type, but with the OpenAPI 'nullable: true' property.
	AvrgWndw *int32 `json:"avrgWndw,omitempty"`

	QosMonConReq QosMonitoringInformationRm `json:"qosMonConReq,omitempty"`

	// Identifies the list of UE addresses subject for Consolidated Data Rate monitoring.
	ListUeConsDtRt []IpAddr `json:"listUeConsDtRt,omitempty"`
}

// AssertAsSessionWithQoSSubscriptionPatchRequired checks if the required fields are not zero-ed
func AssertAsSessionWithQoSSubscriptionPatchRequired(obj AsSessionWithQoSSubscriptionPatch) error {
	for _, el := range obj.FlowInfo {
		if err := AssertFlowInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EthFlowInfo {
		if err := AssertEthFlowDescriptionRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EnEthFlowInfo {
		if err := AssertEthFlowInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ListUeAddrs {
		if err := AssertUeAddInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AltQosReqs {
		if err := AssertAlternativeServiceRequirementsDataRequired(el); err != nil {
			return err
		}
	}
	if obj.UsageThreshold != nil {
		if err := AssertUsageThresholdRmRequired(*obj.UsageThreshold); err != nil {
			return err
		}
	}
	if err := AssertQosMonitoringInformationRmRequired(obj.QosMonInfo); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationRmRequired(obj.PdvMon); err != nil {
		return err
	}
	if err := AssertTscQosRequirementRmRequired(obj.TscQosReq); err != nil {
		return err
	}
	if err := AssertUplinkDownlinkSupportRequired(obj.L4sInd); err != nil {
		return err
	}
	if err := AssertTemporalInValidityRequired(obj.TempInValidity); err != nil {
		return err
	}
	for _, el := range obj.Events {
		if err := AssertUserPlaneEventRequired(el); err != nil {
			return err
		}
	}
	if obj.PduSetQosDl != nil {
		if err := AssertPduSetQosParaRequired(*obj.PduSetQosDl); err != nil {
			return err
		}
	}
	if obj.PduSetQosUl != nil {
		if err := AssertPduSetQosParaRequired(*obj.PduSetQosUl); err != nil {
			return err
		}
	}
	if err := AssertProtocolDescriptionRequired(obj.ProtoDescUl); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionRequired(obj.ProtoDescDl); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationRmRequired(obj.RttMon); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationRmRequired(obj.QosMonDatRate); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationRmRequired(obj.QosMonConReq); err != nil {
		return err
	}
	for _, el := range obj.ListUeConsDtRt {
		if err := AssertIpAddrRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAsSessionWithQoSSubscriptionPatchConstraints checks if the values respects the defined constraints
func AssertAsSessionWithQoSSubscriptionPatchConstraints(obj AsSessionWithQoSSubscriptionPatch) error {
	for _, el := range obj.FlowInfo {
		if err := AssertFlowInfoConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EthFlowInfo {
		if err := AssertEthFlowDescriptionConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EnEthFlowInfo {
		if err := AssertEthFlowInfoConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ListUeAddrs {
		if err := AssertUeAddInfoConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AltQosReqs {
		if err := AssertAlternativeServiceRequirementsDataConstraints(el); err != nil {
			return err
		}
	}
	if obj.UsageThreshold != nil {
		if err := AssertUsageThresholdRmConstraints(*obj.UsageThreshold); err != nil {
			return err
		}
	}
	if err := AssertQosMonitoringInformationRmConstraints(obj.QosMonInfo); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationRmConstraints(obj.PdvMon); err != nil {
		return err
	}
	if err := AssertTscQosRequirementRmConstraints(obj.TscQosReq); err != nil {
		return err
	}
	if err := AssertUplinkDownlinkSupportConstraints(obj.L4sInd); err != nil {
		return err
	}
	if err := AssertTemporalInValidityConstraints(obj.TempInValidity); err != nil {
		return err
	}
	for _, el := range obj.Events {
		if err := AssertUserPlaneEventConstraints(el); err != nil {
			return err
		}
	}
	if obj.PduSetQosDl != nil {
		if err := AssertPduSetQosParaConstraints(*obj.PduSetQosDl); err != nil {
			return err
		}
	}
	if obj.PduSetQosUl != nil {
		if err := AssertPduSetQosParaConstraints(*obj.PduSetQosUl); err != nil {
			return err
		}
	}
	if err := AssertProtocolDescriptionConstraints(obj.ProtoDescUl); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionConstraints(obj.ProtoDescDl); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationRmConstraints(obj.RttMon); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationRmConstraints(obj.QosMonDatRate); err != nil {
		return err
	}
	if obj.AvrgWndw != nil && *obj.AvrgWndw < 1 {
		return &ParsingError{Param: "AvrgWndw", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AvrgWndw != nil && *obj.AvrgWndw > 4095 {
		return &ParsingError{Param: "AvrgWndw", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertQosMonitoringInformationRmConstraints(obj.QosMonConReq); err != nil {
		return err
	}
	for _, el := range obj.ListUeConsDtRt {
		if err := AssertIpAddrConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
