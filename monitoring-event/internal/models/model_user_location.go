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

// UserLocation - At least one of eutraLocation, nrLocation and n3gaLocation shall be present. Several of them may be present.
type UserLocation struct {
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"`

	NrLocation *NrLocation `json:"nrLocation,omitempty"`

	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`

	UtraLocation *UtraLocation `json:"utraLocation,omitempty"`

	GeraLocation *GeraLocation `json:"geraLocation,omitempty"`
}

// AssertUserLocationRequired checks if the required fields are not zero-ed
func AssertUserLocationRequired(obj UserLocation) error {

	return nil
}

// AssertUserLocationConstraints checks if the values respects the defined constraints
func AssertUserLocationConstraints(obj UserLocation) error {
	if err := AssertEutraLocationConstraints(*obj.EutraLocation); err != nil {
		return err
	}
	if err := AssertNrLocationConstraints(*obj.NrLocation); err != nil {
		return err
	}
	if err := AssertN3gaLocationConstraints(*obj.N3gaLocation); err != nil {
		return err
	}
	if obj.UtraLocation != nil {
		if err := AssertUtraLocationConstraints(*obj.UtraLocation); err != nil {
			return err
		}
	}
	if obj.GeraLocation != nil {
		if err := AssertGeraLocationConstraints(*obj.GeraLocation); err != nil {
			return err
		}
	}
	return nil
}
