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

// RelatedUe - Related UE Information
type RelatedUe struct {

	// String identifying an UE with application layer ID. The format of the application  layer ID parameter is same as the Application layer ID defined in clause 11.3.4 of  3GPP TS 24.554.
	ApplicationlayerId string `json:"applicationlayerId"`

	RelatedUEType RelatedUeType `json:"relatedUEType"`
}

// AssertRelatedUeRequired checks if the required fields are not zero-ed
func AssertRelatedUeRequired(obj RelatedUe) error {
	elements := map[string]interface{}{
		"applicationlayerId": obj.ApplicationlayerId,
		"relatedUEType":      obj.RelatedUEType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertRelatedUeTypeRequired(obj.RelatedUEType); err != nil {
		return err
	}
	return nil
}

// AssertRelatedUeConstraints checks if the values respects the defined constraints
func AssertRelatedUeConstraints(obj RelatedUe) error {
	if err := AssertRelatedUeTypeConstraints(obj.RelatedUEType); err != nil {
		return err
	}
	return nil
}
