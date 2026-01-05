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

// EthFlowDescription - Identifies an Ethernet flow.
type EthFlowDescription struct {

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	DestMacAddr string `json:"destMacAddr,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`

	EthType string `json:"ethType"`

	// Defines a packet filter of an IP flow.
	FDesc string `json:"fDesc,omitempty"`

	FDir FlowDirection `json:"fDir,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	SourceMacAddr string `json:"sourceMacAddr,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`

	VlanTags []string `json:"vlanTags,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	SrcMacAddrEnd string `json:"srcMacAddrEnd,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	DestMacAddrEnd string `json:"destMacAddrEnd,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`
}

// AssertEthFlowDescriptionRequired checks if the required fields are not zero-ed
func AssertEthFlowDescriptionRequired(obj EthFlowDescription) error {
	elements := map[string]interface{}{
		"ethType": obj.EthType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertFlowDirectionRequired(obj.FDir); err != nil {
		return err
	}
	return nil
}

// AssertEthFlowDescriptionConstraints checks if the values respects the defined constraints
func AssertEthFlowDescriptionConstraints(obj EthFlowDescription) error {
	if err := AssertFlowDirectionConstraints(obj.FDir); err != nil {
		return err
	}
	return nil
}
