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

// LcsQosClass - Specifies LCS QoS class.
type LcsQosClass struct {
}

// AssertLcsQosClassRequired checks if the required fields are not zero-ed
func AssertLcsQosClassRequired(obj LcsQosClass) error {
	return nil
}

// AssertLcsQosClassConstraints checks if the values respects the defined constraints
func AssertLcsQosClassConstraints(obj LcsQosClass) error {
	return nil
}
