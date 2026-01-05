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

// PdvMonitoringReport - Packet Delay Variation reporting information.
type PdvMonitoringReport struct {

	// Identification of the flows.
	Flows []Flows `json:"flows,omitempty"`

	// Uplink packet delay variation in units of milliseconds.
	UlPdv int32 `json:"ulPdv,omitempty"`

	// Downlink packet delay variation in units of milliseconds.
	DlPdv int32 `json:"dlPdv,omitempty"`

	// Round trip packet delay variation in units of milliseconds.
	RtPdv int32 `json:"rtPdv,omitempty"`
}

// AssertPdvMonitoringReportRequired checks if the required fields are not zero-ed
func AssertPdvMonitoringReportRequired(obj PdvMonitoringReport) error {
	for _, el := range obj.Flows {
		if err := AssertFlowsRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPdvMonitoringReportConstraints checks if the values respects the defined constraints
func AssertPdvMonitoringReportConstraints(obj PdvMonitoringReport) error {
	for _, el := range obj.Flows {
		if err := AssertFlowsConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
