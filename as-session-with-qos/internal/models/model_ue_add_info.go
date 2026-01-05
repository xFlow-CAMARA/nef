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

// UeAddInfo - Represent the UE address information.
type UeAddInfo struct {
	UeIpAddr *IpAddr `json:"ueIpAddr,omitempty"`

	// Unsigned integer with valid values between 0 and 65535.
	PortNumber int32 `json:"portNumber,omitempty"`
}

// AssertUeAddInfoRequired checks if the required fields are not zero-ed
func AssertUeAddInfoRequired(obj UeAddInfo) error {
	if obj.UeIpAddr != nil {
		if err := AssertIpAddrRequired(*obj.UeIpAddr); err != nil {
			return err
		}
	}
	return nil
}

// AssertUeAddInfoConstraints checks if the values respects the defined constraints
func AssertUeAddInfoConstraints(obj UeAddInfo) error {
	if obj.UeIpAddr != nil {
		if err := AssertIpAddrConstraints(*obj.UeIpAddr); err != nil {
			return err
		}
	}
	if obj.PortNumber < 0 {
		return &ParsingError{Param: "PortNumber", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PortNumber > 65535 {
		return &ParsingError{Param: "PortNumber", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
