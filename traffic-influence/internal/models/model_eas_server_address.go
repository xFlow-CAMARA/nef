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
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.4
 */

package models

import (
	"errors"
)

// EasServerAddress - Represents the IP address and port of an EAS server.
type EasServerAddress struct {
	Ip *IpAddr `json:"ip"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Port int32 `json:"port"`
}

// AssertEasServerAddressRequired checks if the required fields are not zero-ed
func AssertEasServerAddressRequired(obj EasServerAddress) error {
	elements := map[string]interface{}{
		"ip":   obj.Ip,
		"port": obj.Port,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.Ip != nil {
		if err := AssertIpAddrRequired(*obj.Ip); err != nil {
			return err
		}
	}
	return nil
}

// AssertEasServerAddressConstraints checks if the values respects the defined constraints
func AssertEasServerAddressConstraints(obj EasServerAddress) error {
	if obj.Ip != nil {
		if err := AssertIpAddrConstraints(*obj.Ip); err != nil {
			return err
		}
	}
	if obj.Port < 0 {
		return &ParsingError{Param: "Port", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
