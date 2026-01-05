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

// TscQosRequirementRm - Represents the same as the TscQosRequirement data type but with the nullable:true property.
type TscQosRequirementRm struct {

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	ReqGbrDl *string `json:"reqGbrDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	ReqGbrUl *string `json:"reqGbrUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	ReqMbrDl *string `json:"reqMbrDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	ReqMbrUl *string `json:"reqMbrUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'ExtMaxDataBurstVol' data type, but with the OpenAPI 'nullable: true' property.
	MaxTscBurstSize *int32 `json:"maxTscBurstSize,omitempty"`

	// This data type is defined in the same way as the 'PacketDelBudget' data type, but with the OpenAPI 'nullable: true' property.
	Req5Gsdelay *int32 `json:"req5Gsdelay,omitempty"`

	// This data type is defined in the same way as the 'PacketErrRate' data type, but with the OpenAPI 'nullable: true' property.
	ReqPer *string `json:"reqPer,omitempty" validate:"regexp=^([0-9]E-[0-9])$"`

	// This data type is defined in the same way as the TscPriorityLevel data type, but with the OpenAPI nullable property set to true.
	Priority *int32 `json:"priority,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	TscaiTimeDom *int32 `json:"tscaiTimeDom,omitempty"`

	TscaiInputDl *TscaiInputContainer `json:"tscaiInputDl,omitempty"`

	TscaiInputUl *TscaiInputContainer `json:"tscaiInputUl,omitempty"`

	// Indicates the capability for AF to adjust the burst sending time, when it is supported and set to \"true\". The default value is \"false\" if omitted.
	CapBatAdaptation *bool `json:"capBatAdaptation,omitempty"`
}

// AssertTscQosRequirementRmRequired checks if the required fields are not zero-ed
func AssertTscQosRequirementRmRequired(obj TscQosRequirementRm) error {
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

// AssertTscQosRequirementRmConstraints checks if the values respects the defined constraints
func AssertTscQosRequirementRmConstraints(obj TscQosRequirementRm) error {
	if obj.MaxTscBurstSize != nil && *obj.MaxTscBurstSize < 4096 {
		return &ParsingError{Param: "MaxTscBurstSize", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaxTscBurstSize != nil && *obj.MaxTscBurstSize > 2000000 {
		return &ParsingError{Param: "MaxTscBurstSize", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.Req5Gsdelay != nil && *obj.Req5Gsdelay < 1 {
		return &ParsingError{Param: "Req5Gsdelay", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Priority != nil && *obj.Priority < 1 {
		return &ParsingError{Param: "Priority", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Priority != nil && *obj.Priority > 8 {
		return &ParsingError{Param: "Priority", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.TscaiTimeDom != nil && *obj.TscaiTimeDom < 0 {
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
