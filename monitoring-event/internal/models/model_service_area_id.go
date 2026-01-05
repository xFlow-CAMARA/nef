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

// ServiceAreaId - Contains a Service Area Identifier as defined in 3GPP TS 23.003, clause 12.5.
type ServiceAreaId struct {
	PlmnId PlmnId1 `json:"plmnId"`

	// Location Area Code.
	Lac string `json:"lac" validate:"regexp=^[A-Fa-f0-9]{4}$"`

	// Service Area Code.
	Sac string `json:"sac" validate:"regexp=^[A-Fa-f0-9]{4}$"`
}

// AssertServiceAreaIdRequired checks if the required fields are not zero-ed
func AssertServiceAreaIdRequired(obj ServiceAreaId) error {
	elements := map[string]interface{}{
		"plmnId": obj.PlmnId,
		"lac":    obj.Lac,
		"sac":    obj.Sac,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnId1Required(obj.PlmnId); err != nil {
		return err
	}
	return nil
}

// AssertServiceAreaIdConstraints checks if the values respects the defined constraints
func AssertServiceAreaIdConstraints(obj ServiceAreaId) error {
	if err := AssertPlmnId1Constraints(obj.PlmnId); err != nil {
		return err
	}
	return nil
}
