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

// TrafficCorrelationInfo - Contains the information for traffic correlation.
type TrafficCorrelationInfo struct {
	CorrType CorrelationType `json:"corrType,omitempty"`

	// Identification of a set of UEs accessing the application identified by the  Application Identifier or traffic filtering information.
	TfcCorrId string `json:"tfcCorrId,omitempty"`

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166 with the OpenAPI defined 'nullable: true' property.
	ComEasIpv4Addr *string `json:"comEasIpv4Addr,omitempty" validate:"regexp=^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`

	ComEasIpv6Addr *Ipv6AddrRm `json:"comEasIpv6Addr,omitempty"`

	FqdnRange *[]FqdnPatternMatchingRule `json:"fqdnRange,omitempty"`

	// String providing an URI formatted according to RFC 3986 with the OpenAPI 'nullable: true' property.
	NotifUri *string `json:"notifUri,omitempty"`

	// Notification correlation identifier.
	NotifCorrId *string `json:"notifCorrId,omitempty"`
}

// AssertTrafficCorrelationInfoRequired checks if the required fields are not zero-ed
func AssertTrafficCorrelationInfoRequired(obj TrafficCorrelationInfo) error {
	if err := AssertCorrelationTypeRequired(obj.CorrType); err != nil {
		return err
	}
	if obj.ComEasIpv6Addr != nil {
		if err := AssertIpv6AddrRmRequired(*obj.ComEasIpv6Addr); err != nil {
			return err
		}
	}
	if obj.FqdnRange != nil {
		for _, el := range *obj.FqdnRange {
			if err := AssertFqdnPatternMatchingRuleRequired(el); err != nil {
				return err
			}
		}
	}
	return nil
}

// AssertTrafficCorrelationInfoConstraints checks if the values respects the defined constraints
func AssertTrafficCorrelationInfoConstraints(obj TrafficCorrelationInfo) error {
	if err := AssertCorrelationTypeConstraints(obj.CorrType); err != nil {
		return err
	}
	if obj.ComEasIpv6Addr != nil {
		if err := AssertIpv6AddrRmConstraints(*obj.ComEasIpv6Addr); err != nil {
			return err
		}
	}
	if obj.FqdnRange != nil {
		for _, el := range *obj.FqdnRange {
			if err := AssertFqdnPatternMatchingRuleConstraints(el); err != nil {
				return err
			}
		}
	}
	return nil
}
