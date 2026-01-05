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

// PduSessionInformation - Represents PDU session identification information.
type PduSessionInformation struct {
	Snssai Snssai `json:"snssai"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	UeIpv4 string `json:"ueIpv4,omitempty"`

	UeIpv6 Ipv6Prefix `json:"ueIpv6,omitempty"`

	IpDomain string `json:"ipDomain,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	UeMac string `json:"ueMac,omitempty"`
}

// AssertPduSessionInformationRequired checks if the required fields are not zero-ed
func AssertPduSessionInformationRequired(obj PduSessionInformation) error {
	elements := map[string]interface{}{
		"snssai": obj.Snssai,
		"dnn":    obj.Dnn,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	if err := AssertIpv6PrefixRequired(obj.UeIpv6); err != nil {
		return err
	}
	return nil
}

// AssertPduSessionInformationConstraints checks if the values respects the defined constraints
func AssertPduSessionInformationConstraints(obj PduSessionInformation) error {
	if err := AssertSnssaiConstraints(obj.Snssai); err != nil {
		return err
	}
	if err := AssertIpv6PrefixConstraints(obj.UeIpv6); err != nil {
		return err
	}
	return nil
}
