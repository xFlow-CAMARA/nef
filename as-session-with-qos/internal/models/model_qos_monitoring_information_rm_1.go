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

// QosMonitoringInformationRm1 - This data type is defined in the same way as the QosMonitoringInformation data type, but with the OpenAPI nullable property set to true.
type QosMonitoringInformationRm1 struct {
	RepThreshDl int32 `json:"repThreshDl,omitempty"`

	RepThreshUl int32 `json:"repThreshUl,omitempty"`

	RepThreshRp int32 `json:"repThreshRp,omitempty"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	RepThreshDatRateUl *string `json:"repThreshDatRateUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	RepThreshDatRateDl *string `json:"repThreshDatRateDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	ConThreshDl *int32 `json:"conThreshDl,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	ConThreshUl *int32 `json:"conThreshUl,omitempty"`
}

// AssertQosMonitoringInformationRm1Required checks if the required fields are not zero-ed
func AssertQosMonitoringInformationRm1Required(obj QosMonitoringInformationRm1) error {
	return nil
}

// AssertQosMonitoringInformationRm1Constraints checks if the values respects the defined constraints
func AssertQosMonitoringInformationRm1Constraints(obj QosMonitoringInformationRm1) error {
	if obj.ConThreshDl != nil && *obj.ConThreshDl < 0 {
		return &ParsingError{Param: "ConThreshDl", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ConThreshUl != nil && *obj.ConThreshUl < 0 {
		return &ParsingError{Param: "ConThreshUl", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
