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

// MonitoringEventSubscription - Represents a subscription to event(s) monitoring.
type MonitoringEventSubscription struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	// Identifies the MTC Service Provider and/or MTC Application.
	MtcProviderId string `json:"mtcProviderId,omitempty"`

	// Identifies the Application Identifier(s)
	AppIds []string `json:"appIds,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn string `json:"msisdn,omitempty"`

	// Indicates the added external Identifier(s) within the active group.
	AddedExternalIds []string `json:"addedExternalIds,omitempty"`

	// Indicates the added MSISDN(s) within the active group.
	AddedMsisdns []string `json:"addedMsisdns,omitempty"`

	// Indicates cancellation of the external Identifier(s) within the active group.
	ExcludedExternalIds []string `json:"excludedExternalIds,omitempty"`

	// Indicates cancellation of the MSISDN(s) within the active group.
	ExcludedMsisdns []string `json:"excludedMsisdns,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId string `json:"externalGroupId,omitempty"`

	AddExtGroupId []string `json:"addExtGroupId,omitempty"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	Ipv6Addr string `json:"ipv6Addr,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination"`

	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false by the SCS/AS indicates not request SCEF to send a test notification, default false if omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`

	MonitoringType MonitoringType `json:"monitoringType"`

	// Identifies the maximum number of event reports to be generated by the HSS, MME/SGSN as specified in clause 5.6.0 of 3GPP TS 23.682.
	MaximumNumberOfReports int32 `json:"maximumNumberOfReports,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MonitorExpireTime time.Time `json:"monitorExpireTime,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	RepPeriod int32 `json:"repPeriod,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	GroupReportGuardTime int32 `json:"groupReportGuardTime,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumDetectionTime int32 `json:"maximumDetectionTime,omitempty"`

	ReachabilityType ReachabilityType `json:"reachabilityType,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumLatency int32 `json:"maximumLatency,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumResponseTime int32 `json:"maximumResponseTime,omitempty"`

	// If \"monitoringType\" is \"UE_REACHABILITY\", this parameter may be included to identify the number of packets that the serving gateway shall buffer in case that the UE is not reachable.
	SuggestedNumberOfDlPackets int32 `json:"suggestedNumberOfDlPackets,omitempty"`

	// If \"monitoringType\" is set to \"UE_REACHABILITY\" or \"AVAILABILITY_AFTER_DDN_FAILURE\", this parameter may be included to indicate the notification of when a UE, for which PSM is enabled, transitions into idle mode. \"true\"  indicates enabling of notification; \"false\"  indicate no need to notify. Default value is \"false\" if omitted.
	IdleStatusIndication bool `json:"idleStatusIndication,omitempty"`

	LocationType LocationType `json:"locationType,omitempty"`

	Accuracy Accuracy `json:"accuracy,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MinimumReportInterval int32 `json:"minimumReportInterval,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaxRptExpireIntvl int32 `json:"maxRptExpireIntvl,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	SamplingInterval int32 `json:"samplingInterval,omitempty"`

	// Indicates whether to request the location estimate for event reporting. If \"monitoringType\" is \"LOCATION_REPORTING\", this parameter may be included to indicate whether event reporting requires the location information. If set to true, the location estimation information shall be included in event reporting. If set to \"false\", indicates the location estimation information shall not be included in event reporting. Default \"false\" if omitted.
	ReportingLocEstInd bool `json:"reportingLocEstInd,omitempty"`

	// Minimum straight line distance moved by a UE to trigger a motion event report.
	LinearDistance int32 `json:"linearDistance,omitempty"`

	LocQoS LocationQoS `json:"locQoS,omitempty"`

	// Contains the service identity
	SvcId string `json:"svcId,omitempty"`

	LdrType LdrType `json:"ldrType,omitempty"`

	VelocityRequested VelocityRequested `json:"velocityRequested,omitempty"`

	// Indicates value of the age of the location estimate.
	MaxAgeOfLocEst int32 `json:"maxAgeOfLocEst,omitempty"`

	LocTimeWindow TimeWindow `json:"locTimeWindow,omitempty"`

	SupportedGADShapes []SupportedGadShapes `json:"supportedGADShapes,omitempty"`

	// Contains the codeword
	CodeWord string `json:"codeWord,omitempty"`

	// Indicates whether location reporting over user plane is requested or not. \"true\" indicates the location reporting over user plane is requested. \"false\" indicates the location reporting over user plane is not requested. Default value is \"false\" if omitted.
	UpLocRepIndAf bool `json:"upLocRepIndAf,omitempty"`

	UpLocRepAddrAf *UpLocRepAddrAfRm `json:"upLocRepAddrAf,omitempty"`

	AssociationType AssociationType `json:"associationType,omitempty"`

	// If \"monitoringType\" is \"ROAMING_STATUS\", this parameter may be included to indicate the notification of UE's Serving PLMN ID. Value \"true\" indicates enabling of notification; \"false\" indicates disabling of notification. Default value is \"false\" if omitted.
	PlmnIndication bool `json:"plmnIndication,omitempty"`

	LocationArea LocationArea `json:"locationArea,omitempty"`

	LocationArea5G LocationArea5G `json:"locationArea5G,omitempty"`

	DddTraDescriptors []DddTrafficDescriptor `json:"dddTraDescriptors,omitempty"`

	DddStati []DlDataDeliveryStatus `json:"dddStati,omitempty"`

	ApiNames []string `json:"apiNames,omitempty"`

	MonitoringEventReport MonitoringEventReport `json:"monitoringEventReport,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	TgtNsThreshold SacInfo `json:"tgtNsThreshold,omitempty"`

	NsRepFormat SacRepFormat `json:"nsRepFormat,omitempty"`

	AfServiceId string `json:"afServiceId,omitempty"`

	// Indicates whether an immediate reporting is requested or not. \"true\" indicate an immediate reporting is requested. \"false\" indicate an immediate reporting is not requested. Default value \"false\" if omitted.
	ImmediateRep bool `json:"immediateRep,omitempty"`

	UavPolicy UavPolicy `json:"uavPolicy,omitempty"`

	// Set to true by the SCS/AS so that only UAV's with \"PDU session established for DNN(s) subject to aerial service\" are to be listed in the Event report. Set to false or default false if omitted otherwise.
	SesEstInd bool `json:"sesEstInd,omitempty"`

	SubType SubType `json:"subType,omitempty"`

	AddnMonTypes []MonitoringType `json:"addnMonTypes,omitempty"`

	AddnMonEventReports []MonitoringEventReport `json:"addnMonEventReports,omitempty"`

	UeIpAddr *IpAddr `json:"ueIpAddr,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	UeMacAddr string `json:"ueMacAddr,omitempty"`

	// string providing an URI formatted according to IETF RFC 3986.
	RevocationNotifUri string `json:"revocationNotifUri,omitempty"`

	ReqRangSlRes []RangingSlResult `json:"reqRangSlRes,omitempty"`

	// Contains a list of the related UE(s) for the ranging and sidelink positioning and the corresponding information. The key of the map shall be a any unique string set to the value.
	RelatedUEs map[string]RelatedUe `json:"relatedUEs,omitempty"`

	/* TODO remove it, it is just for development porpouses*/
	Supi string
}

// AssertMonitoringEventSubscriptionRequired checks if the required fields are not zero-ed
func AssertMonitoringEventSubscriptionRequired(obj MonitoringEventSubscription) error {
	elements := map[string]interface{}{
		"notificationDestination": obj.NotificationDestination,
		"monitoringType":          obj.MonitoringType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertWebsockNotifConfigRequired(obj.WebsockNotifConfig); err != nil {
		return err
	}
	if err := AssertMonitoringTypeRequired(obj.MonitoringType); err != nil {
		return err
	}
	if err := AssertReachabilityTypeRequired(obj.ReachabilityType); err != nil {
		return err
	}
	if err := AssertLocationTypeRequired(obj.LocationType); err != nil {
		return err
	}
	if err := AssertAccuracyRequired(obj.Accuracy); err != nil {
		return err
	}
	if err := AssertLocationQoSRequired(obj.LocQoS); err != nil {
		return err
	}
	if err := AssertLdrTypeRequired(obj.LdrType); err != nil {
		return err
	}
	if err := AssertVelocityRequestedRequired(obj.VelocityRequested); err != nil {
		return err
	}
	if err := AssertTimeWindowRequired(obj.LocTimeWindow); err != nil {
		return err
	}
	for _, el := range obj.SupportedGADShapes {
		if err := AssertSupportedGadShapesRequired(el); err != nil {
			return err
		}
	}
	if obj.UpLocRepAddrAf != nil {
		if err := AssertUpLocRepAddrAfRmRequired(*obj.UpLocRepAddrAf); err != nil {
			return err
		}
	}
	if err := AssertAssociationTypeRequired(obj.AssociationType); err != nil {
		return err
	}
	if err := AssertLocationAreaRequired(obj.LocationArea); err != nil {
		return err
	}
	if err := AssertLocationArea5GRequired(obj.LocationArea5G); err != nil {
		return err
	}
	for _, el := range obj.DddTraDescriptors {
		if err := AssertDddTrafficDescriptorRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.DddStati {
		if err := AssertDlDataDeliveryStatusRequired(el); err != nil {
			return err
		}
	}
	if err := AssertMonitoringEventReportRequired(obj.MonitoringEventReport); err != nil {
		return err
	}
	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	if err := AssertSacInfoRequired(obj.TgtNsThreshold); err != nil {
		return err
	}
	if err := AssertSacRepFormatRequired(obj.NsRepFormat); err != nil {
		return err
	}
	if err := AssertUavPolicyRequired(obj.UavPolicy); err != nil {
		return err
	}
	if err := AssertSubTypeRequired(obj.SubType); err != nil {
		return err
	}
	for _, el := range obj.AddnMonTypes {
		if err := AssertMonitoringTypeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AddnMonEventReports {
		if err := AssertMonitoringEventReportRequired(el); err != nil {
			return err
		}
	}
	if obj.UeIpAddr != nil {
		if err := AssertIpAddrRequired(*obj.UeIpAddr); err != nil {
			return err
		}
	}
	for _, el := range obj.ReqRangSlRes {
		if err := AssertRangingSlResultRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMonitoringEventSubscriptionConstraints checks if the values respects the defined constraints
func AssertMonitoringEventSubscriptionConstraints(obj MonitoringEventSubscription) error {
	if err := AssertWebsockNotifConfigConstraints(obj.WebsockNotifConfig); err != nil {
		return err
	}
	if err := AssertMonitoringTypeConstraints(obj.MonitoringType); err != nil {
		return err
	}
	if obj.MaximumNumberOfReports < 1 {
		return &ParsingError{Param: "MaximumNumberOfReports", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.RepPeriod < 0 {
		return &ParsingError{Param: "RepPeriod", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.GroupReportGuardTime < 0 {
		return &ParsingError{Param: "GroupReportGuardTime", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaximumDetectionTime < 0 {
		return &ParsingError{Param: "MaximumDetectionTime", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertReachabilityTypeConstraints(obj.ReachabilityType); err != nil {
		return err
	}
	if obj.MaximumLatency < 0 {
		return &ParsingError{Param: "MaximumLatency", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaximumResponseTime < 0 {
		return &ParsingError{Param: "MaximumResponseTime", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SuggestedNumberOfDlPackets < 0 {
		return &ParsingError{Param: "SuggestedNumberOfDlPackets", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertLocationTypeConstraints(obj.LocationType); err != nil {
		return err
	}
	if err := AssertAccuracyConstraints(obj.Accuracy); err != nil {
		return err
	}
	if obj.MinimumReportInterval < 0 {
		return &ParsingError{Param: "MinimumReportInterval", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaxRptExpireIntvl < 0 {
		return &ParsingError{Param: "MaxRptExpireIntvl", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SamplingInterval < 0 {
		return &ParsingError{Param: "SamplingInterval", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.LinearDistance < 1 {
		return &ParsingError{Param: "LinearDistance", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.LinearDistance > 10000 {
		return &ParsingError{Param: "LinearDistance", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertLocationQoSConstraints(obj.LocQoS); err != nil {
		return err
	}
	if err := AssertLdrTypeConstraints(obj.LdrType); err != nil {
		return err
	}
	if err := AssertVelocityRequestedConstraints(obj.VelocityRequested); err != nil {
		return err
	}
	if obj.MaxAgeOfLocEst < 0 {
		return &ParsingError{Param: "MaxAgeOfLocEst", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaxAgeOfLocEst > 32767 {
		return &ParsingError{Param: "MaxAgeOfLocEst", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertTimeWindowConstraints(obj.LocTimeWindow); err != nil {
		return err
	}
	for _, el := range obj.SupportedGADShapes {
		if err := AssertSupportedGadShapesConstraints(el); err != nil {
			return err
		}
	}
	if obj.UpLocRepAddrAf != nil {
		if err := AssertUpLocRepAddrAfRmConstraints(*obj.UpLocRepAddrAf); err != nil {
			return err
		}
	}
	if err := AssertAssociationTypeConstraints(obj.AssociationType); err != nil {
		return err
	}
	if err := AssertLocationAreaConstraints(obj.LocationArea); err != nil {
		return err
	}
	if err := AssertLocationArea5GConstraints(obj.LocationArea5G); err != nil {
		return err
	}
	for _, el := range obj.DddTraDescriptors {
		if err := AssertDddTrafficDescriptorConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.DddStati {
		if err := AssertDlDataDeliveryStatusConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertMonitoringEventReportConstraints(obj.MonitoringEventReport); err != nil {
		return err
	}
	if err := AssertSnssaiConstraints(obj.Snssai); err != nil {
		return err
	}
	if err := AssertSacInfoConstraints(obj.TgtNsThreshold); err != nil {
		return err
	}
	if err := AssertSacRepFormatConstraints(obj.NsRepFormat); err != nil {
		return err
	}
	if err := AssertUavPolicyConstraints(obj.UavPolicy); err != nil {
		return err
	}
	if err := AssertSubTypeConstraints(obj.SubType); err != nil {
		return err
	}
	for _, el := range obj.AddnMonTypes {
		if err := AssertMonitoringTypeConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AddnMonEventReports {
		if err := AssertMonitoringEventReportConstraints(el); err != nil {
			return err
		}
	}
	if obj.UeIpAddr != nil {
		if err := AssertIpAddrConstraints(*obj.UeIpAddr); err != nil {
			return err
		}
	}
	for _, el := range obj.ReqRangSlRes {
		if err := AssertRangingSlResultConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
