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

// LocationArea5G - Represents a user location area when the UE is attached to 5G.
type LocationArea5G struct {

	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`

	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress `json:"civicAddresses,omitempty"`

	NwAreaInfo NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
}

// AssertLocationArea5GRequired checks if the required fields are not zero-ed
func AssertLocationArea5GRequired(obj LocationArea5G) error {
	for _, el := range obj.GeographicAreas {
		if err := AssertGeographicAreaRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.CivicAddresses {
		if err := AssertCivicAddressRequired(el); err != nil {
			return err
		}
	}
	if err := AssertNetworkAreaInfoRequired(obj.NwAreaInfo); err != nil {
		return err
	}
	return nil
}

// AssertLocationArea5GConstraints checks if the values respects the defined constraints
func AssertLocationArea5GConstraints(obj LocationArea5G) error {
	for _, el := range obj.GeographicAreas {
		if err := AssertGeographicAreaConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.CivicAddresses {
		if err := AssertCivicAddressConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertNetworkAreaInfoConstraints(obj.NwAreaInfo); err != nil {
		return err
	}
	return nil
}
