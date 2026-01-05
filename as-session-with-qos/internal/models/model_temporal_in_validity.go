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
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.5
 */

package models

import (
	"time"
)

// TemporalInValidity - Indicates the time interval(s) during which the AF request is not to be applied.
type TemporalInValidity struct {

	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`

	// string with format 'date-time' as defined in OpenAPI.
	StopTime time.Time `json:"stopTime"`
}

// AssertTemporalInValidityRequired checks if the required fields are not zero-ed
func AssertTemporalInValidityRequired(obj TemporalInValidity) error {
	elements := map[string]interface{}{
		"startTime": obj.StartTime,
		"stopTime":  obj.StopTime,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertTemporalInValidityConstraints checks if the values respects the defined constraints
func AssertTemporalInValidityConstraints(obj TemporalInValidity) error {
	return nil
}
