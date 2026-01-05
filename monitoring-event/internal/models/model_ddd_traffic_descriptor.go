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

import (
	"errors"
)

// DddTrafficDescriptor - Contains a Traffic Descriptor.
type DddTrafficDescriptor struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty" validate:"regexp=^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	MacAddr string `json:"macAddr,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`
}

// AssertDddTrafficDescriptorRequired checks if the required fields are not zero-ed
func AssertDddTrafficDescriptorRequired(obj DddTrafficDescriptor) error {
	if err := AssertIpv6AddrRequired(obj.Ipv6Addr); err != nil {
		return err
	}
	return nil
}

// AssertDddTrafficDescriptorConstraints checks if the values respects the defined constraints
func AssertDddTrafficDescriptorConstraints(obj DddTrafficDescriptor) error {
	if err := AssertIpv6AddrConstraints(obj.Ipv6Addr); err != nil {
		return err
	}
	if obj.PortNumber < 0 {
		return &ParsingError{Param: "PortNumber", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
