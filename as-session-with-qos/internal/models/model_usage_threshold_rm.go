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

// UsageThresholdRm - Represents the same as the UsageThreshold data type but with the nullable:true property.
type UsageThresholdRm struct {

	// Unsigned integer identifying a period of time in units of seconds with \"nullable=true\" property.
	Duration *int32 `json:"duration,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	TotalVolume *int64 `json:"totalVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	UplinkVolume *int64 `json:"uplinkVolume,omitempty"`
}

// AssertUsageThresholdRmRequired checks if the required fields are not zero-ed
func AssertUsageThresholdRmRequired(obj UsageThresholdRm) error {
	return nil
}

// AssertUsageThresholdRmConstraints checks if the values respects the defined constraints
func AssertUsageThresholdRmConstraints(obj UsageThresholdRm) error {
	if obj.Duration != nil && *obj.Duration < 0 {
		return &ParsingError{Param: "Duration", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TotalVolume != nil && *obj.TotalVolume < 0 {
		return &ParsingError{Param: "TotalVolume", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.DownlinkVolume != nil && *obj.DownlinkVolume < 0 {
		return &ParsingError{Param: "DownlinkVolume", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.UplinkVolume != nil && *obj.UplinkVolume < 0 {
		return &ParsingError{Param: "UplinkVolume", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
