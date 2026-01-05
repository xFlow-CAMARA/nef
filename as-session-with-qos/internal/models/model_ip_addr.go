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

// IpAddr - Contains an IP adresse.
type IpAddr struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	Ipv6Prefix Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}

// AssertIpAddrRequired checks if the required fields are not zero-ed
func AssertIpAddrRequired(obj IpAddr) error {
	if err := AssertIpv6AddrRequired(obj.Ipv6Addr); err != nil {
		return err
	}
	if err := AssertIpv6PrefixRequired(obj.Ipv6Prefix); err != nil {
		return err
	}
	return nil
}

// AssertIpAddrConstraints checks if the values respects the defined constraints
func AssertIpAddrConstraints(obj IpAddr) error {
	if err := AssertIpv6AddrConstraints(obj.Ipv6Addr); err != nil {
		return err
	}
	if err := AssertIpv6PrefixConstraints(obj.Ipv6Prefix); err != nil {
		return err
	}
	return nil
}
