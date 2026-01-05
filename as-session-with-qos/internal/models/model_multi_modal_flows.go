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

// MultiModalFlows - Represents a flow information within a single-modal data flow.
type MultiModalFlows struct {

	// It contains the ordinal number of the single-modal data flow. Identifies the single-modal data flow.
	MedCompN int32 `json:"medCompN"`

	// Identifies the affected flow(s) within the single-modal data flow (identified by the medCompN attribute). It may be omitted when all flows are affected.
	FlowIds []int32 `json:"flowIds,omitempty"`
}

// AssertMultiModalFlowsRequired checks if the required fields are not zero-ed
func AssertMultiModalFlowsRequired(obj MultiModalFlows) error {
	elements := map[string]interface{}{
		"medCompN": obj.MedCompN,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertMultiModalFlowsConstraints checks if the values respects the defined constraints
func AssertMultiModalFlowsConstraints(obj MultiModalFlows) error {
	return nil
}
