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

// TsnQosContainerRm - Indicates removable TSC Traffic QoS.
type TsnQosContainerRm struct {

	// This data type is defined in the same way as the 'ExtMaxDataBurstVol' data type, but with the OpenAPI 'nullable: true' property.
	MaxTscBurstSize *int32 `json:"maxTscBurstSize,omitempty"`

	// This data type is defined in the same way as the 'PacketDelBudget' data type, but with the OpenAPI 'nullable: true' property.
	TscPackDelay *int32 `json:"tscPackDelay,omitempty"`

	// This data type is defined in the same way as the 'PacketErrRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxPer *string `json:"maxPer,omitempty" validate:"regexp=^([0-9]E-[0-9])$"`

	// This data type is defined in the same way as the TscPriorityLevel data type, but with the OpenAPI nullable property set to true.
	TscPrioLevel *int32 `json:"tscPrioLevel,omitempty"`
}

// AssertTsnQosContainerRmRequired checks if the required fields are not zero-ed
func AssertTsnQosContainerRmRequired(obj TsnQosContainerRm) error {
	return nil
}

// AssertTsnQosContainerRmConstraints checks if the values respects the defined constraints
func AssertTsnQosContainerRmConstraints(obj TsnQosContainerRm) error {
	if obj.MaxTscBurstSize != nil && *obj.MaxTscBurstSize < 4096 {
		return &ParsingError{Param: "MaxTscBurstSize", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaxTscBurstSize != nil && *obj.MaxTscBurstSize > 2000000 {
		return &ParsingError{Param: "MaxTscBurstSize", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.TscPackDelay != nil && *obj.TscPackDelay < 1 {
		return &ParsingError{Param: "TscPackDelay", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TscPrioLevel != nil && *obj.TscPrioLevel < 1 {
		return &ParsingError{Param: "TscPrioLevel", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TscPrioLevel != nil && *obj.TscPrioLevel > 8 {
		return &ParsingError{Param: "TscPrioLevel", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
