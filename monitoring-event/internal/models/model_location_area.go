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

// LocationArea - Represents a user location area.
type LocationArea struct {

	// Indicates a list of Cell Global Identities of the user which identifies the cell the UE is registered.
	CellIds []string `json:"cellIds,omitempty"`

	// Indicates a list of eNodeB identities in which the UE is currently located.
	EnodeBIds []string `json:"enodeBIds,omitempty"`

	// Identifies a list of Routing Area Identities of the user where the UE is located.
	RoutingAreaIds []string `json:"routingAreaIds,omitempty"`

	// Identifies a list of Tracking Area Identities of the user where the UE is located.
	TrackingAreaIds []string `json:"trackingAreaIds,omitempty"`

	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`

	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress `json:"civicAddresses,omitempty"`
}

// AssertLocationAreaRequired checks if the required fields are not zero-ed
func AssertLocationAreaRequired(obj LocationArea) error {
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
	return nil
}

// AssertLocationAreaConstraints checks if the values respects the defined constraints
func AssertLocationAreaConstraints(obj LocationArea) error {
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
	return nil
}
