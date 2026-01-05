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

// PduSetQosPara - Represents the PDU Set level QoS parameters.
type PduSetQosPara struct {

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds.
	PduSetDelayBudget int32 `json:"pduSetDelayBudget,omitempty"`

	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	PduSetErrRate string `json:"pduSetErrRate,omitempty" validate:"regexp=^([0-9]E-[0-9])$"`

	PduSetHandlingInfo PduSetHandlingInfo `json:"pduSetHandlingInfo,omitempty"`
}

// AssertPduSetQosParaRequired checks if the required fields are not zero-ed
func AssertPduSetQosParaRequired(obj PduSetQosPara) error {
	if err := AssertPduSetHandlingInfoRequired(obj.PduSetHandlingInfo); err != nil {
		return err
	}
	return nil
}

// AssertPduSetQosParaConstraints checks if the values respects the defined constraints
func AssertPduSetQosParaConstraints(obj PduSetQosPara) error {
	if obj.PduSetDelayBudget < 1 {
		return &ParsingError{Param: "PduSetDelayBudget", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertPduSetHandlingInfoConstraints(obj.PduSetHandlingInfo); err != nil {
		return err
	}
	return nil
}
