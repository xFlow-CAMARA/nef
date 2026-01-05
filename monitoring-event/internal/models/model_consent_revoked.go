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

// ConsentRevoked - Represents the information related to a revoked user consent.
type ConsentRevoked struct {
	UcPurpose UcPurpose `json:"ucPurpose"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn string `json:"msisdn,omitempty"`
}

// AssertConsentRevokedRequired checks if the required fields are not zero-ed
func AssertConsentRevokedRequired(obj ConsentRevoked) error {
	elements := map[string]interface{}{
		"ucPurpose": obj.UcPurpose,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUcPurposeRequired(obj.UcPurpose); err != nil {
		return err
	}
	return nil
}

// AssertConsentRevokedConstraints checks if the values respects the defined constraints
func AssertConsentRevokedConstraints(obj ConsentRevoked) error {
	if err := AssertUcPurposeConstraints(obj.UcPurpose); err != nil {
		return err
	}
	return nil
}
