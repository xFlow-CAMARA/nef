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

// TsnQosContainer - Indicates TSC Traffic QoS.
type TsnQosContainer struct {

	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.
	MaxTscBurstSize int32 `json:"maxTscBurstSize,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	TscPackDelay int32 `json:"tscPackDelay,omitempty"`

	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	MaxPer string `json:"maxPer,omitempty" validate:"regexp=^([0-9]E-[0-9])$"`

	// Represents the priority level of TSC Flows.
	TscPrioLevel int32 `json:"tscPrioLevel,omitempty"`
}

// AssertTsnQosContainerRequired checks if the required fields are not zero-ed
func AssertTsnQosContainerRequired(obj TsnQosContainer) error {
	return nil
}

// AssertTsnQosContainerConstraints checks if the values respects the defined constraints
func AssertTsnQosContainerConstraints(obj TsnQosContainer) error {
	if obj.MaxTscBurstSize < 4096 {
		return &ParsingError{Param: "MaxTscBurstSize", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaxTscBurstSize > 2000000 {
		return &ParsingError{Param: "MaxTscBurstSize", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.TscPackDelay < 1 {
		return &ParsingError{Param: "TscPackDelay", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TscPrioLevel < 1 {
		return &ParsingError{Param: "TscPrioLevel", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TscPrioLevel > 8 {
		return &ParsingError{Param: "TscPrioLevel", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
