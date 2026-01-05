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

// Ncgi - Contains the NCGI (NR Cell Global Identity), as described in 3GPP 23.003
type Ncgi struct {
	PlmnId PlmnId `json:"plmnId"`

	// 36-bit string identifying an NR Cell Id as specified in clause 9.3.1.7 of 3GPP TS 38.413,  in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character  representing the 4 most significant bits of the Cell Id shall appear first in the string, and  the character representing the 4 least significant bit of the Cell Id shall appear last in the  string.
	NrCellId string `json:"nrCellId" validate:"regexp=^[A-Fa-f0-9]{9}$"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid string `json:"nid,omitempty" validate:"regexp=^[A-Fa-f0-9]{11}$"`
}

// AssertNcgiRequired checks if the required fields are not zero-ed
func AssertNcgiRequired(obj Ncgi) error {
	elements := map[string]interface{}{
		"plmnId":   obj.PlmnId,
		"nrCellId": obj.NrCellId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnIdRequired(obj.PlmnId); err != nil {
		return err
	}
	return nil
}

// AssertNcgiConstraints checks if the values respects the defined constraints
func AssertNcgiConstraints(obj Ncgi) error {
	if err := AssertPlmnIdConstraints(obj.PlmnId); err != nil {
		return err
	}
	return nil
}
