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
	"fmt"
)

// NullValue : JSON's null value.
type NullValue string

// List of NullValue
const ()

// AllowedNullValueEnumValues is all the allowed values of NullValue enum
var AllowedNullValueEnumValues = []NullValue{}

// validNullValueEnumValue provides a map of NullValues for fast verification of use input
var validNullValueEnumValues = map[NullValue]struct{}{}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NullValue) IsValid() bool {
	_, ok := validNullValueEnumValues[v]
	return ok
}

// NewNullValueFromValue returns a pointer to a valid NullValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNullValueFromValue(v string) (NullValue, error) {
	ev := NullValue(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for NullValue: valid values are %v", v, AllowedNullValueEnumValues)
}

// AssertNullValueRequired checks if the required fields are not zero-ed
func AssertNullValueRequired(obj NullValue) error {
	return nil
}

// AssertNullValueConstraints checks if the values respects the defined constraints
func AssertNullValueConstraints(obj NullValue) error {
	return nil
}
