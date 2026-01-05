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

// AsSessionWithQoSSubscription - Represents an individual AS session with required QoS subscription resource.
type AsSessionWithQoSSubscription struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination"`

	// Identifies the external Application Identifier.
	ExterAppId string `json:"exterAppId,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExtGroupId string `json:"extGroupId,omitempty"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" validate:"regexp=^(msisdn-[0-9]{5,15}|extid-[^@]+@[^@]+|.+)$"`

	// Describe the data flow which requires QoS.
	FlowInfo []FlowInfo `json:"flowInfo,omitempty"`

	// Identifies Ethernet packet flows.
	EthFlowInfo []EthFlowDescription `json:"ethFlowInfo,omitempty"`

	// Identifies the Ethernet flows which require QoS. Each Ethernet flow consists of a flow idenifer and the corresponding UL and/or DL flows.
	EnEthFlowInfo []EthFlowInfo `json:"enEthFlowInfo,omitempty"`

	// Identifies the list of UE address.
	ListUeAddrs []UeAddInfo `json:"listUeAddrs,omitempty"`

	// This data type contains a multi-modal service identifier.
	MultiModalId string `json:"multiModalId,omitempty"`

	ProtoDescUl ProtocolDescription `json:"protoDescUl,omitempty"`

	ProtoDescDl ProtocolDescription `json:"protoDescDl,omitempty"`

	// Identifies a pre-defined QoS information
	QosReference string `json:"qosReference,omitempty"`

	// Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority.
	AltQoSReferences []string `json:"altQoSReferences,omitempty"`

	// Identifies an ordered list of alternative service requirements that include individual QoS parameter sets. The lower the index of the array for a given entry, the higher the priority.
	AltQosReqs []AlternativeServiceRequirementsData `json:"altQosReqs,omitempty"`

	// Indicates whether the QoS flow parameters signalling to the UE when the SMF is notified by the NG-RAN of changes in the fulfilled QoS situation is disabled (true) or not (false). Default value is false. The fulfilled situation is either the QoS profile or an Alternative QoS Profile.
	DisUeNotif bool `json:"disUeNotif,omitempty"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	UeIpv4Addr string `json:"ueIpv4Addr,omitempty"`

	IpDomain string `json:"ipDomain,omitempty"`

	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	UeIpv6Addr string `json:"ueIpv6Addr,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	MacAddr string `json:"macAddr,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`

	UsageThreshold UsageThreshold `json:"usageThreshold,omitempty"`

	SponsorInfo SponsorInformation `json:"sponsorInfo,omitempty"`

	QosMonInfo QosMonitoringInformation `json:"qosMonInfo,omitempty"`

	PdvMon QosMonitoringInformation `json:"pdvMon,omitempty"`

	// indicating a time in seconds.
	QosDuration int32 `json:"qosDuration,omitempty"`

	// indicating a time in seconds.
	QosInactInt int32 `json:"qosInactInt,omitempty"`

	// Indicates whether the direct event notification is requested (true) or not (false) for the provided and/or previously provided QoS monitoring parameter(s).
	DirectNotifInd bool `json:"directNotifInd,omitempty"`

	TscQosReq TscQosRequirement `json:"tscQosReq,omitempty"`

	L4sInd UplinkDownlinkSupport `json:"l4sInd,omitempty"`

	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	TempInValidity TemporalInValidity `json:"tempInValidity,omitempty"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`

	// Represents the list of user plane event(s) to which the SCS/AS requests to subscribe to.
	Events []UserPlaneEvent `json:"events,omitempty"`

	// Contains media component data for a single-modal data flow(s). The key of the map is the medCompN attribute.
	MultiModDatFlows map[string]AsSessionMediaComponent `json:"multiModDatFlows,omitempty"`

	PduSetQosDl PduSetQosPara `json:"pduSetQosDl,omitempty"`

	PduSetQosUl PduSetQosPara `json:"pduSetQosUl,omitempty"`

	// Indicates the service data flow needs to meet the Round-Trip (RT) latency requirement of the service, when it is included and set to \"true\". The default value is \"false\" if omitted.
	RTLatencyInd bool `json:"rTLatencyInd,omitempty"`

	// Indicates the time interval in units of milliseconds.
	PeriodUl int32 `json:"periodUl,omitempty"`

	// Indicates the time interval in units of milliseconds.
	PeriodDl int32 `json:"periodDl,omitempty"`

	RttMon QosMonitoringInformation `json:"rttMon,omitempty"`

	QosMonDatRate QosMonitoringInformation `json:"qosMonDatRate,omitempty"`

	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AvrgWndw int32 `json:"avrgWndw,omitempty"`

	ServAuthInfo ServAuthInfo `json:"servAuthInfo,omitempty"`

	QosMonConReq QosMonitoringInformation `json:"qosMonConReq,omitempty"`

	// Identifies the list of UE addresses subject for Consolidated Data Rate monitoring.
	ListUeConsDtRt []IpAddr `json:"listUeConsDtRt,omitempty"`
}

//TODO See what to do with those assertions... code does not seem correct

// AssertAsSessionWithQoSSubscriptionRequired checks if the required fields are not zero-ed
func AssertAsSessionWithQoSSubscriptionRequired(obj AsSessionWithQoSSubscription) error {
	/*	elements := map[string]interface{}{
			"notificationDestination": obj.NotificationDestination,
		}
		for name, el := range elements {
			if isZero := IsZeroValue(el); isZero {
				return &RequiredError{Field: name}
			}
		}

		if err := AssertSnssaiRequired(obj.Snssai); err != nil {
			return err
		}
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
		if err := AssertProtocolDescriptionRequired(obj.ProtoDescUl); err != nil {
			return err
		}
		if err := AssertProtocolDescriptionRequired(obj.ProtoDescDl); err != nil {
			return err
		}
		for _, el := range obj.AltQosReqs {
			if err := AssertAlternativeServiceRequirementsDataRequired(el); err != nil {
				return err
			}
		}
		if err := AssertUsageThresholdRequired(obj.UsageThreshold); err != nil {
			return err
		}
		if err := AssertSponsorInformationRequired(obj.SponsorInfo); err != nil {
			return err
		}
		if err := AssertQosMonitoringInformationRequired(obj.QosMonInfo); err != nil {
			return err
		}
		if err := AssertQosMonitoringInformationRequired(obj.PdvMon); err != nil {
			return err
		}
		if err := AssertTscQosRequirementRequired(obj.TscQosReq); err != nil {
			return err
		}
		if err := AssertUplinkDownlinkSupportRequired(obj.L4sInd); err != nil {
			return err
		}
		if err := AssertTemporalInValidityRequired(obj.TempInValidity); err != nil {
			return err
		}
		if err := AssertWebsockNotifConfigRequired(obj.WebsockNotifConfig); err != nil {
			return err
		}
		for _, el := range obj.Events {
			if err := AssertUserPlaneEventRequired(el); err != nil {
				return err
			}
		}
		if err := AssertPduSetQosParaRequired(obj.PduSetQosDl); err != nil {
			return err
		}
		if err := AssertPduSetQosParaRequired(obj.PduSetQosUl); err != nil {
			return err
		}
		if err := AssertQosMonitoringInformationRequired(obj.RttMon); err != nil {
			return err
		}
		if err := AssertQosMonitoringInformationRequired(obj.QosMonDatRate); err != nil {
			return err
		}
		if err := AssertServAuthInfoRequired(obj.ServAuthInfo); err != nil {
			return err
		}
		if err := AssertQosMonitoringInformationRequired(obj.QosMonConReq); err != nil {
			return err
		}
		for _, el := range obj.ListUeConsDtRt {
			if err := AssertIpAddrRequired(el); err != nil {
				return err
			}
		}*/
	return nil
}

// AssertAsSessionWithQoSSubscriptionConstraints checks if the values respects the defined constraints
func AssertAsSessionWithQoSSubscriptionConstraints(obj AsSessionWithQoSSubscription) error {
	/*if err := AssertSnssaiConstraints(obj.Snssai); err != nil {
		return err
	}
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
	if err := AssertProtocolDescriptionConstraints(obj.ProtoDescUl); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionConstraints(obj.ProtoDescDl); err != nil {
		return err
	}
	for _, el := range obj.AltQosReqs {
		if err := AssertAlternativeServiceRequirementsDataConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertUsageThresholdConstraints(obj.UsageThreshold); err != nil {
		return err
	}
	if err := AssertSponsorInformationConstraints(obj.SponsorInfo); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationConstraints(obj.QosMonInfo); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationConstraints(obj.PdvMon); err != nil {
		return err
	}
	if err := AssertTscQosRequirementConstraints(obj.TscQosReq); err != nil {
		return err
	}
	if err := AssertUplinkDownlinkSupportConstraints(obj.L4sInd); err != nil {
		return err
	}
	if err := AssertTemporalInValidityConstraints(obj.TempInValidity); err != nil {
		return err
	}
	if err := AssertWebsockNotifConfigConstraints(obj.WebsockNotifConfig); err != nil {
		return err
	}
	for _, el := range obj.Events {
		if err := AssertUserPlaneEventConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertPduSetQosParaConstraints(obj.PduSetQosDl); err != nil {
		return err
	}
	if err := AssertPduSetQosParaConstraints(obj.PduSetQosUl); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationConstraints(obj.RttMon); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationConstraints(obj.QosMonDatRate); err != nil {
		return err
	}
	if obj.AvrgWndw < 1 {
		return &ParsingError{Param: "AvrgWndw", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AvrgWndw > 4095 {
		return &ParsingError{Param: "AvrgWndw", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertServAuthInfoConstraints(obj.ServAuthInfo); err != nil {
		return err
	}
	if err := AssertQosMonitoringInformationConstraints(obj.QosMonConReq); err != nil {
		return err
	}
	for _, el := range obj.ListUeConsDtRt {
		if err := AssertIpAddrConstraints(el); err != nil {
			return err
		}
	}*/
	return nil
}
