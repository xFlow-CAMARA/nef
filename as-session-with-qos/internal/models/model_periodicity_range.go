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

// PeriodicityRange - Contains the acceptable range (which is formulated as lower bound and upper bound of the periodicity of the start twobursts in reference to the external GM) or acceptable periodicity value(s) (which is formulated as a list of values for the periodicity).
type PeriodicityRange struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	LowerBound int32 `json:"lowerBound,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	UpperBound int32 `json:"upperBound,omitempty"`

	PeriodicVals []int32 `json:"periodicVals,omitempty"`
}

// AssertPeriodicityRangeRequired checks if the required fields are not zero-ed
func AssertPeriodicityRangeRequired(obj PeriodicityRange) error {
	return nil
}

// AssertPeriodicityRangeConstraints checks if the values respects the defined constraints
func AssertPeriodicityRangeConstraints(obj PeriodicityRange) error {
	if obj.LowerBound < 0 {
		return &ParsingError{Param: "LowerBound", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.UpperBound < 0 {
		return &ParsingError{Param: "UpperBound", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
