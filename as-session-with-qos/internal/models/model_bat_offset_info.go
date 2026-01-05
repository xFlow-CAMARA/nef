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
	"errors"
)

// BatOffsetInfo - Indicates the offset of the BAT and the optionally adjusted periodicity.
type BatOffsetInfo struct {

	// Indicates the BAT offset of the arrival time of the data burst in units of milliseconds.
	RanBatOffsetNotif int32 `json:"ranBatOffsetNotif"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AdjPeriod int32 `json:"adjPeriod,omitempty"`

	// Identification of the flows. If no flows are provided, the BAT offset applies for all flows of the AF session.
	Flows []Flows `json:"flows,omitempty"`
}

// AssertBatOffsetInfoRequired checks if the required fields are not zero-ed
func AssertBatOffsetInfoRequired(obj BatOffsetInfo) error {
	elements := map[string]interface{}{
		"ranBatOffsetNotif": obj.RanBatOffsetNotif,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Flows {
		if err := AssertFlowsRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertBatOffsetInfoConstraints checks if the values respects the defined constraints
func AssertBatOffsetInfoConstraints(obj BatOffsetInfo) error {
	if obj.AdjPeriod < 0 {
		return &ParsingError{Param: "AdjPeriod", Err: errors.New(errMsgMinValueConstraint)}
	}
	for _, el := range obj.Flows {
		if err := AssertFlowsConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
