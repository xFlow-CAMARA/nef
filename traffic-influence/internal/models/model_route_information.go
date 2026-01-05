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

// RouteInformation - At least one of the \"ipv4Addr\" attribute and the \"ipv6Addr\" attribute shall be included in the \"RouteInformation\" data type.
type RouteInformation struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty" validate:"regexp=^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`

	Ipv6Addr Ipv6Addr1 `json:"ipv6Addr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber"`
}

// AssertRouteInformationRequired checks if the required fields are not zero-ed
func AssertRouteInformationRequired(obj RouteInformation) error {
	elements := map[string]interface{}{
		"portNumber": obj.PortNumber,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertIpv6Addr1Required(obj.Ipv6Addr); err != nil {
		return err
	}
	return nil
}

// AssertRouteInformationConstraints checks if the values respects the defined constraints
func AssertRouteInformationConstraints(obj RouteInformation) error {
	if err := AssertIpv6Addr1Constraints(obj.Ipv6Addr); err != nil {
		return err
	}
	if obj.PortNumber < 0 {
		return &ParsingError{Param: "PortNumber", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
