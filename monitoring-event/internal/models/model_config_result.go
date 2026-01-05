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

// ConfigResult - Represents one configuration processing result for a group's members.
type ConfigResult struct {

	// Each element indicates an external identifier of the UE.
	ExternalIds []string `json:"externalIds,omitempty"`

	// Each element identifies the MS internal PSTN/ISDN number allocated for the UE.
	Msisdns []string `json:"msisdns,omitempty"`

	ResultReason ResultReason `json:"resultReason"`
}

// AssertConfigResultRequired checks if the required fields are not zero-ed
func AssertConfigResultRequired(obj ConfigResult) error {
	elements := map[string]interface{}{
		"resultReason": obj.ResultReason,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertResultReasonRequired(obj.ResultReason); err != nil {
		return err
	}
	return nil
}

// AssertConfigResultConstraints checks if the values respects the defined constraints
func AssertConfigResultConstraints(obj ConfigResult) error {
	if err := AssertResultReasonConstraints(obj.ResultReason); err != nil {
		return err
	}
	return nil
}
