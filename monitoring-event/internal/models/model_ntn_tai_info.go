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

type NtnTaiInfo struct {
	PlmnId PlmnIdNid `json:"plmnId"`

	TacList []string `json:"tacList"`

	// 2 or 3-octet string identifying a tracking area code as specified in clause 9.3.3.10  of 3GPP TS 38.413, in hexadecimal representation. Each character in the string shall  take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the TAC shall  appear first in the string, and the character representing the 4 least significant bit  of the TAC shall appear last in the string.
	DerivedTac string `json:"derivedTac,omitempty" validate:"regexp=(^[A-Fa-f0-9]{4}$)|(^[A-Fa-f0-9]{6}$)"`
}

// AssertNtnTaiInfoRequired checks if the required fields are not zero-ed
func AssertNtnTaiInfoRequired(obj NtnTaiInfo) error {
	elements := map[string]interface{}{
		"plmnId":  obj.PlmnId,
		"tacList": obj.TacList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnIdNidRequired(obj.PlmnId); err != nil {
		return err
	}
	return nil
}

// AssertNtnTaiInfoConstraints checks if the values respects the defined constraints
func AssertNtnTaiInfoConstraints(obj NtnTaiInfo) error {
	if err := AssertPlmnIdNidConstraints(obj.PlmnId); err != nil {
		return err
	}
	return nil
}
