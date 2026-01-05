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

// SacEventStatus - Contains the network slice status information in terms of the current number of UEs registered  with a network slice, the current number of PDU Sessions established on a network slice or both.
type SacEventStatus struct {
	ReachedNumUes SacInfo `json:"reachedNumUes,omitempty"`

	ReachedNumPduSess SacInfo `json:"reachedNumPduSess,omitempty"`
}

// AssertSacEventStatusRequired checks if the required fields are not zero-ed
func AssertSacEventStatusRequired(obj SacEventStatus) error {
	if err := AssertSacInfoRequired(obj.ReachedNumUes); err != nil {
		return err
	}
	if err := AssertSacInfoRequired(obj.ReachedNumPduSess); err != nil {
		return err
	}
	return nil
}

// AssertSacEventStatusConstraints checks if the values respects the defined constraints
func AssertSacEventStatusConstraints(obj SacEventStatus) error {
	if err := AssertSacInfoConstraints(obj.ReachedNumUes); err != nil {
		return err
	}
	if err := AssertSacInfoConstraints(obj.ReachedNumPduSess); err != nil {
		return err
	}
	return nil
}
