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

// AlternativeServiceRequirementsData - Contains an alternative QoS related parameter set.
type AlternativeServiceRequirementsData struct {

	// Reference to this alternative QoS related parameter set.
	AltQosParamSetRef string `json:"altQosParamSetRef"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrUl string `json:"gbrUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrDl string `json:"gbrDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Pdb int32 `json:"pdb,omitempty"`

	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	Per string `json:"per,omitempty" validate:"regexp=^([0-9]E-[0-9])$"`
}

// AssertAlternativeServiceRequirementsDataRequired checks if the required fields are not zero-ed
func AssertAlternativeServiceRequirementsDataRequired(obj AlternativeServiceRequirementsData) error {
	elements := map[string]interface{}{
		"altQosParamSetRef": obj.AltQosParamSetRef,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertAlternativeServiceRequirementsDataConstraints checks if the values respects the defined constraints
func AssertAlternativeServiceRequirementsDataConstraints(obj AlternativeServiceRequirementsData) error {
	if obj.Pdb < 1 {
		return &ParsingError{Param: "Pdb", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
