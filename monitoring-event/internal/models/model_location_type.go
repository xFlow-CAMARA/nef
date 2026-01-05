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

// LocationType - Represents a location type.   Possible values are - CURRENT_LOCATION: The SCS/AS requests to be notified for current location - LAST_KNOWN_LOCATION: The SCS/AS requests to be notified for last known location - CURRENT_OR_LAST_KNOWN_LOCATION: The AF requests the current or last known location - INITIAL_LOCATION: The AF requests the initial location
type LocationType struct {
}

// AssertLocationTypeRequired checks if the required fields are not zero-ed
func AssertLocationTypeRequired(obj LocationType) error {
	return nil
}

// AssertLocationTypeConstraints checks if the values respects the defined constraints
func AssertLocationTypeConstraints(obj LocationType) error {
	return nil
}
