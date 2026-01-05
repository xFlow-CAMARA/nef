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

// QosMonitoringReport1 - QoS Monitoring reporting information.
type QosMonitoringReport1 struct {
	Flows []Flows `json:"flows,omitempty"`

	UlDelays []int32 `json:"ulDelays,omitempty"`

	DlDelays []int32 `json:"dlDelays,omitempty"`

	RtDelays []int32 `json:"rtDelays,omitempty"`

	// Represents the packet delay measurement failure indicator.
	Pdmf bool `json:"pdmf,omitempty"`

	UlConInfo []int32 `json:"ulConInfo,omitempty"`

	DlConInfo []int32 `json:"dlConInfo,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	UlDataRate string `json:"ulDataRate,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	DlDataRate string `json:"dlDataRate,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
}

// AssertQosMonitoringReport1Required checks if the required fields are not zero-ed
func AssertQosMonitoringReport1Required(obj QosMonitoringReport1) error {
	for _, el := range obj.Flows {
		if err := AssertFlowsRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertQosMonitoringReport1Constraints checks if the values respects the defined constraints
func AssertQosMonitoringReport1Constraints(obj QosMonitoringReport1) error {
	for _, el := range obj.Flows {
		if err := AssertFlowsConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
