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

// AddFlowDescriptionInfo - Contains additional flow description information.
type AddFlowDescriptionInfo struct {

	// 4-octet string representing the security parameter index of the IPSec packet in hexadecimal representation.
	Spi string `json:"spi,omitempty"`

	// 3-octet string representing the IPv6 flow label header field in hexadecimal representation.
	FlowLabel string `json:"flowLabel,omitempty"`

	FlowDir FlowDirection `json:"flowDir,omitempty"`
}

// AssertAddFlowDescriptionInfoRequired checks if the required fields are not zero-ed
func AssertAddFlowDescriptionInfoRequired(obj AddFlowDescriptionInfo) error {
	if err := AssertFlowDirectionRequired(obj.FlowDir); err != nil {
		return err
	}
	return nil
}

// AssertAddFlowDescriptionInfoConstraints checks if the values respects the defined constraints
func AssertAddFlowDescriptionInfoConstraints(obj AddFlowDescriptionInfo) error {
	if err := AssertFlowDirectionConstraints(obj.FlowDir); err != nil {
		return err
	}
	return nil
}
