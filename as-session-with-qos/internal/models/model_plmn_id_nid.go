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

// PlmnIdNid - Contains the serving core network operator PLMN ID and, for an SNPN, the NID that together with the PLMN ID identifies the SNPN.
type PlmnIdNid struct {

	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mcc string `json:"mcc" validate:"regexp=^\\\\d{3}$"`

	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in  clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc" validate:"regexp=^\\\\d{2,3}$"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid string `json:"nid,omitempty" validate:"regexp=^[A-Fa-f0-9]{11}$"`
}

// AssertPlmnIdNidRequired checks if the required fields are not zero-ed
func AssertPlmnIdNidRequired(obj PlmnIdNid) error {
	elements := map[string]interface{}{
		"mcc": obj.Mcc,
		"mnc": obj.Mnc,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPlmnIdNidConstraints checks if the values respects the defined constraints
func AssertPlmnIdNidConstraints(obj PlmnIdNid) error {
	return nil
}
