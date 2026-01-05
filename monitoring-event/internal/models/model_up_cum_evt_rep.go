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

import (
	"errors"
)

// UpCumEvtRep - Represents the cumulative event report.
type UpCumEvtRep struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	UpLocRepStat int32 `json:"upLocRepStat,omitempty"`
}

// AssertUpCumEvtRepRequired checks if the required fields are not zero-ed
func AssertUpCumEvtRepRequired(obj UpCumEvtRep) error {
	return nil
}

// AssertUpCumEvtRepConstraints checks if the values respects the defined constraints
func AssertUpCumEvtRepConstraints(obj UpCumEvtRep) error {
	if obj.UpLocRepStat < 0 {
		return &ParsingError{Param: "UpLocRepStat", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
