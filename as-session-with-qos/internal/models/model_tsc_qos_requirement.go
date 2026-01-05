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

// TscQosRequirement - Represents QoS requirements for time sensitive communication.
type TscQosRequirement struct {

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ReqGbrDl string `json:"reqGbrDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ReqGbrUl string `json:"reqGbrUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ReqMbrDl string `json:"reqMbrDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ReqMbrUl string `json:"reqMbrUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.
	MaxTscBurstSize int32 `json:"maxTscBurstSize,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Req5Gsdelay int32 `json:"req5Gsdelay,omitempty"`

	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	ReqPer string `json:"reqPer,omitempty" validate:"regexp=^([0-9]E-[0-9])$"`

	// Represents the priority level of TSC Flows.
	Priority int32 `json:"priority,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TscaiTimeDom int32 `json:"tscaiTimeDom,omitempty"`

	TscaiInputDl *TscaiInputContainer `json:"tscaiInputDl,omitempty"`

	TscaiInputUl *TscaiInputContainer `json:"tscaiInputUl,omitempty"`

	// Indicates the capability for AF to adjust the burst sending time, when it is supported and set to \"true\". The default value is \"false\" if omitted.
	CapBatAdaptation bool `json:"capBatAdaptation,omitempty"`
}

// AssertTscQosRequirementRequired checks if the required fields are not zero-ed
func AssertTscQosRequirementRequired(obj TscQosRequirement) error {
	if obj.TscaiInputDl != nil {
		if err := AssertTscaiInputContainerRequired(*obj.TscaiInputDl); err != nil {
			return err
		}
	}
	if obj.TscaiInputUl != nil {
		if err := AssertTscaiInputContainerRequired(*obj.TscaiInputUl); err != nil {
			return err
		}
	}
	return nil
}

// AssertTscQosRequirementConstraints checks if the values respects the defined constraints
func AssertTscQosRequirementConstraints(obj TscQosRequirement) error {
	if obj.MaxTscBurstSize < 4096 {
		return &ParsingError{Param: "MaxTscBurstSize", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaxTscBurstSize > 2000000 {
		return &ParsingError{Param: "MaxTscBurstSize", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.Req5Gsdelay < 1 {
		return &ParsingError{Param: "Req5Gsdelay", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Priority < 1 {
		return &ParsingError{Param: "Priority", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Priority > 8 {
		return &ParsingError{Param: "Priority", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.TscaiTimeDom < 0 {
		return &ParsingError{Param: "TscaiTimeDom", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TscaiInputDl != nil {
		if err := AssertTscaiInputContainerConstraints(*obj.TscaiInputDl); err != nil {
			return err
		}
	}
	if obj.TscaiInputUl != nil {
		if err := AssertTscaiInputContainerConstraints(*obj.TscaiInputUl); err != nil {
			return err
		}
	}
	return nil
}
