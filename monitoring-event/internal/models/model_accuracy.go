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

// Accuracy - Represents a desired granularity of accuracy of the requested location information.   Possible values are - CGI_ECGI: The SCS/AS requests to be notified using cell level location accuracy. - ENODEB: The SCS/AS requests to be notified using eNodeB level location accuracy. - TA_RA: The SCS/AS requests to be notified using TA/RA level location accuracy. - PLMN: The SCS/AS requests to be notified using PLMN level location accuracy. - TWAN_ID: The SCS/AS requests to be notified using TWAN identifier level location accuracy. - GEO_AREA: The SCS/AS requests to be notified using the geographical area accuracy. - CIVIC_ADDR: The SCS/AS requests to be notified using the civic address accuracy.
type Accuracy struct {
}

// AssertAccuracyRequired checks if the required fields are not zero-ed
func AssertAccuracyRequired(obj Accuracy) error {
	return nil
}

// AssertAccuracyConstraints checks if the values respects the defined constraints
func AssertAccuracyConstraints(obj Accuracy) error {
	return nil
}
