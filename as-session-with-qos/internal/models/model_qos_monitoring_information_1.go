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

// QosMonitoringInformation1 - Indicates the QoS Monitoring information to report, i.e. UL and/or DL and or round trip delay.
type QosMonitoringInformation1 struct {
	RepThreshDl int32 `json:"repThreshDl,omitempty"`

	RepThreshUl int32 `json:"repThreshUl,omitempty"`

	RepThreshRp int32 `json:"repThreshRp,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RepThreshDatRateUl string `json:"repThreshDatRateUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RepThreshDatRateDl string `json:"repThreshDatRateDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ConThreshDl int32 `json:"conThreshDl,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ConThreshUl int32 `json:"conThreshUl,omitempty"`
}

// AssertQosMonitoringInformation1Required checks if the required fields are not zero-ed
func AssertQosMonitoringInformation1Required(obj QosMonitoringInformation1) error {
	return nil
}

// AssertQosMonitoringInformation1Constraints checks if the values respects the defined constraints
func AssertQosMonitoringInformation1Constraints(obj QosMonitoringInformation1) error {
	if obj.ConThreshDl < 0 {
		return &ParsingError{Param: "ConThreshDl", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ConThreshUl < 0 {
		return &ParsingError{Param: "ConThreshUl", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
