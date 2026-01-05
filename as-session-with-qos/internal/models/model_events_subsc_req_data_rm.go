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

// EventsSubscReqDataRm - This data type is defined in the same way as the EventsSubscReqData data type, but with the OpenAPI nullable property set to true.
type EventsSubscReqDataRm struct {
	Events []AfEventSubscription `json:"events"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri,omitempty"`

	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams,omitempty"`

	QosMon *QosMonitoringInformationRm1 `json:"qosMon,omitempty"`

	QosMonDatRate *QosMonitoringInformationRm1 `json:"qosMonDatRate,omitempty"`

	PdvReqMonParams []RequestedQosMonitoringParameter `json:"pdvReqMonParams,omitempty"`

	PdvMon *QosMonitoringInformationRm1 `json:"pdvMon,omitempty"`

	CongestMon *QosMonitoringInformationRm1 `json:"congestMon,omitempty"`

	ReqAnis []RequiredAccessInfo `json:"reqAnis,omitempty"`

	UsgThres *UsageThresholdRm `json:"usgThres,omitempty"`

	NotifCorreId string `json:"notifCorreId,omitempty"`

	// Indicates whether the direct event notification is requested (true) or not (false) for the provided and/or previously provided QoS monitoring parameters.
	DirectNotifInd *bool `json:"directNotifInd,omitempty"`

	// This data type is defined in the same way as the 'AverWindow' data type, but with the OpenAPI 'nullable: true' property.
	AvrgWndw *int32 `json:"avrgWndw,omitempty"`
}

// AssertEventsSubscReqDataRmRequired checks if the required fields are not zero-ed
func AssertEventsSubscReqDataRmRequired(obj EventsSubscReqDataRm) error {
	elements := map[string]interface{}{
		"events": obj.Events,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Events {
		if err := AssertAfEventSubscriptionRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ReqQosMonParams {
		if err := AssertRequestedQosMonitoringParameterRequired(el); err != nil {
			return err
		}
	}
	if obj.QosMon != nil {
		if err := AssertQosMonitoringInformationRm1Required(*obj.QosMon); err != nil {
			return err
		}
	}
	if obj.QosMonDatRate != nil {
		if err := AssertQosMonitoringInformationRm1Required(*obj.QosMonDatRate); err != nil {
			return err
		}
	}
	for _, el := range obj.PdvReqMonParams {
		if err := AssertRequestedQosMonitoringParameterRequired(el); err != nil {
			return err
		}
	}
	if obj.PdvMon != nil {
		if err := AssertQosMonitoringInformationRm1Required(*obj.PdvMon); err != nil {
			return err
		}
	}
	if obj.CongestMon != nil {
		if err := AssertQosMonitoringInformationRm1Required(*obj.CongestMon); err != nil {
			return err
		}
	}
	for _, el := range obj.ReqAnis {
		if err := AssertRequiredAccessInfoRequired(el); err != nil {
			return err
		}
	}
	if obj.UsgThres != nil {
		if err := AssertUsageThresholdRmRequired(*obj.UsgThres); err != nil {
			return err
		}
	}
	return nil
}

// AssertEventsSubscReqDataRmConstraints checks if the values respects the defined constraints
func AssertEventsSubscReqDataRmConstraints(obj EventsSubscReqDataRm) error {
	for _, el := range obj.Events {
		if err := AssertAfEventSubscriptionConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ReqQosMonParams {
		if err := AssertRequestedQosMonitoringParameterConstraints(el); err != nil {
			return err
		}
	}
	if obj.QosMon != nil {
		if err := AssertQosMonitoringInformationRm1Constraints(*obj.QosMon); err != nil {
			return err
		}
	}
	if obj.QosMonDatRate != nil {
		if err := AssertQosMonitoringInformationRm1Constraints(*obj.QosMonDatRate); err != nil {
			return err
		}
	}
	for _, el := range obj.PdvReqMonParams {
		if err := AssertRequestedQosMonitoringParameterConstraints(el); err != nil {
			return err
		}
	}
	if obj.PdvMon != nil {
		if err := AssertQosMonitoringInformationRm1Constraints(*obj.PdvMon); err != nil {
			return err
		}
	}
	if obj.CongestMon != nil {
		if err := AssertQosMonitoringInformationRm1Constraints(*obj.CongestMon); err != nil {
			return err
		}
	}
	for _, el := range obj.ReqAnis {
		if err := AssertRequiredAccessInfoConstraints(el); err != nil {
			return err
		}
	}
	if obj.UsgThres != nil {
		if err := AssertUsageThresholdRmConstraints(*obj.UsgThres); err != nil {
			return err
		}
	}
	if obj.AvrgWndw != nil && *obj.AvrgWndw < 1 {
		return &ParsingError{Param: "AvrgWndw", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AvrgWndw != nil && *obj.AvrgWndw > 4095 {
		return &ParsingError{Param: "AvrgWndw", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
