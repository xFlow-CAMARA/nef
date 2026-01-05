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

// NetworkAreaInfo - Describes a network area information in which the NF service consumer requests the number of UEs.
type NetworkAreaInfo struct {

	// Contains a list of E-UTRA cell identities.
	Ecgis []Ecgi `json:"ecgis,omitempty"`

	// Contains a list of NR cell identities.
	Ncgis []Ncgi `json:"ncgis,omitempty"`

	// Contains a list of NG RAN nodes.
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`

	// Contains a list of tracking area identities.
	Tais []Tai `json:"tais,omitempty"`
}

// AssertNetworkAreaInfoRequired checks if the required fields are not zero-ed
func AssertNetworkAreaInfoRequired(obj NetworkAreaInfo) error {
	for _, el := range obj.Ecgis {
		if err := AssertEcgiRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ncgis {
		if err := AssertNcgiRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.GRanNodeIds {
		if err := AssertGlobalRanNodeIdRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Tais {
		if err := AssertTaiRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertNetworkAreaInfoConstraints checks if the values respects the defined constraints
func AssertNetworkAreaInfoConstraints(obj NetworkAreaInfo) error {
	for _, el := range obj.Ecgis {
		if err := AssertEcgiConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ncgis {
		if err := AssertNcgiConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.GRanNodeIds {
		if err := AssertGlobalRanNodeIdConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Tais {
		if err := AssertTaiConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
