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

// QosMonitoringInformationRm - Represents the same as the QosMonitoringInformation data type but with the nullable:true property.
type QosMonitoringInformationRm struct {
	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams,omitempty"`

	RepFreqs []ReportingFrequency `json:"repFreqs,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	RepThreshDl *int32 `json:"repThreshDl,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	RepThreshUl *int32 `json:"repThreshUl,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	RepThreshRp *int32 `json:"repThreshRp,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	ConThreshDl *int32 `json:"conThreshDl,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	ConThreshUl *int32 `json:"conThreshUl,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	WaitTime *int32 `json:"waitTime,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	RepPeriod *int32 `json:"repPeriod,omitempty"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	RepThreshDatRateDl *string `json:"repThreshDatRateDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	RepThreshDatRateUl *string `json:"repThreshDatRateUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	ConsDataRateThrDl *string `json:"consDataRateThrDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	ConsDataRateThrUl *string `json:"consDataRateThrUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
}

// AssertQosMonitoringInformationRmRequired checks if the required fields are not zero-ed
func AssertQosMonitoringInformationRmRequired(obj QosMonitoringInformationRm) error {
	for _, el := range obj.ReqQosMonParams {
		if err := AssertRequestedQosMonitoringParameterRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RepFreqs {
		if err := AssertReportingFrequencyRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertQosMonitoringInformationRmConstraints checks if the values respects the defined constraints
func AssertQosMonitoringInformationRmConstraints(obj QosMonitoringInformationRm) error {
	for _, el := range obj.ReqQosMonParams {
		if err := AssertRequestedQosMonitoringParameterConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RepFreqs {
		if err := AssertReportingFrequencyConstraints(el); err != nil {
			return err
		}
	}
	if obj.RepThreshDl != nil && *obj.RepThreshDl < 0 {
		return &ParsingError{Param: "RepThreshDl", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.RepThreshUl != nil && *obj.RepThreshUl < 0 {
		return &ParsingError{Param: "RepThreshUl", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.RepThreshRp != nil && *obj.RepThreshRp < 0 {
		return &ParsingError{Param: "RepThreshRp", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ConThreshDl != nil && *obj.ConThreshDl < 0 {
		return &ParsingError{Param: "ConThreshDl", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ConThreshUl != nil && *obj.ConThreshUl < 0 {
		return &ParsingError{Param: "ConThreshUl", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
