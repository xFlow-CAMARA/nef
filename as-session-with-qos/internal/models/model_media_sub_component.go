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

// MediaSubComponent - Identifies a media subcomponent.
type MediaSubComponent struct {
	AfSigProtocol *AfSigProtocol `json:"afSigProtocol,omitempty"`

	EthfDescs []EthFlowDescription `json:"ethfDescs,omitempty"`

	FNum int32 `json:"fNum"`

	FDescs []string `json:"fDescs,omitempty"`

	// Represents additional flow description information (flow label and IPsec SPI) per Uplink and/or Downlink IP flows.
	AddInfoFlowDescs []AddFlowDescriptionInfo `json:"addInfoFlowDescs,omitempty"`

	FStatus string `json:"fStatus,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDl string `json:"marBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUl string `json:"marBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// 2-octet string, where each octet is encoded in hexadecimal representation. The first octet contains the IPv4 Type-of-Service or the IPv6 Traffic-Class field and the second octet contains the ToS/Traffic Class mask field.
	TosTrCl string `json:"tosTrCl,omitempty"`

	FlowUsage FlowUsage `json:"flowUsage,omitempty"`

	EvSubsc EventsSubscReqData `json:"evSubsc,omitempty"`
}

// AssertMediaSubComponentRequired checks if the required fields are not zero-ed
func AssertMediaSubComponentRequired(obj MediaSubComponent) error {
	elements := map[string]interface{}{
		"fNum": obj.FNum,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.AfSigProtocol != nil {
		if err := AssertAfSigProtocolRequired(*obj.AfSigProtocol); err != nil {
			return err
		}
	}
	for _, el := range obj.EthfDescs {
		if err := AssertEthFlowDescriptionRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AddInfoFlowDescs {
		if err := AssertAddFlowDescriptionInfoRequired(el); err != nil {
			return err
		}
	}
	if err := AssertFlowStatusRequired(obj.FStatus); err != nil {
		return err
	}
	if err := AssertFlowUsageRequired(obj.FlowUsage); err != nil {
		return err
	}
	if err := AssertEventsSubscReqDataRequired(obj.EvSubsc); err != nil {
		return err
	}
	return nil
}

// AssertMediaSubComponentConstraints checks if the values respects the defined constraints
func AssertMediaSubComponentConstraints(obj MediaSubComponent) error {
	if obj.AfSigProtocol != nil {
		if err := AssertAfSigProtocolConstraints(*obj.AfSigProtocol); err != nil {
			return err
		}
	}
	for _, el := range obj.EthfDescs {
		if err := AssertEthFlowDescriptionConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AddInfoFlowDescs {
		if err := AssertAddFlowDescriptionInfoConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertFlowStatusConstraints(obj.FStatus); err != nil {
		return err
	}
	if err := AssertFlowUsageConstraints(obj.FlowUsage); err != nil {
		return err
	}
	if err := AssertEventsSubscReqDataConstraints(obj.EvSubsc); err != nil {
		return err
	}
	return nil
}
