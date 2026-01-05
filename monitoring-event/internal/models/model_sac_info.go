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

// SacInfo - Represents threshold(s) to control the triggering of network slice reporting notifications or the information contained in the network slice reporting notification.
type SacInfo struct {
	NumericValNumUes int32 `json:"numericValNumUes,omitempty"`

	NumericValNumPduSess int32 `json:"numericValNumPduSess,omitempty"`

	PercValueNumUes int32 `json:"percValueNumUes,omitempty"`

	PercValueNumPduSess int32 `json:"percValueNumPduSess,omitempty"`

	UesWithPduSessionInd bool `json:"uesWithPduSessionInd,omitempty"`
}

// AssertSacInfoRequired checks if the required fields are not zero-ed
func AssertSacInfoRequired(obj SacInfo) error {
	return nil
}

// AssertSacInfoConstraints checks if the values respects the defined constraints
func AssertSacInfoConstraints(obj SacInfo) error {
	if obj.PercValueNumUes < 0 {
		return &ParsingError{Param: "PercValueNumUes", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PercValueNumUes > 100 {
		return &ParsingError{Param: "PercValueNumUes", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.PercValueNumPduSess < 0 {
		return &ParsingError{Param: "PercValueNumPduSess", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PercValueNumPduSess > 100 {
		return &ParsingError{Param: "PercValueNumPduSess", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
