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

// HorizontalWithVerticalVelocity - Horizontal and vertical velocity.
type HorizontalWithVerticalVelocity struct {

	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`

	// Indicates value of angle.
	Bearing int32 `json:"bearing"`

	// Indicates value of vertical speed.
	VSpeed float32 `json:"vSpeed"`

	VDirection VerticalDirection `json:"vDirection"`
}

// AssertHorizontalWithVerticalVelocityRequired checks if the required fields are not zero-ed
func AssertHorizontalWithVerticalVelocityRequired(obj HorizontalWithVerticalVelocity) error {
	elements := map[string]interface{}{
		"hSpeed":     obj.HSpeed,
		"bearing":    obj.Bearing,
		"vSpeed":     obj.VSpeed,
		"vDirection": obj.VDirection,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertHorizontalWithVerticalVelocityConstraints checks if the values respects the defined constraints
func AssertHorizontalWithVerticalVelocityConstraints(obj HorizontalWithVerticalVelocity) error {
	if obj.HSpeed < 0 {
		return &ParsingError{Param: "HSpeed", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.HSpeed > 2047 {
		return &ParsingError{Param: "HSpeed", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.Bearing < 0 {
		return &ParsingError{Param: "Bearing", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Bearing > 360 {
		return &ParsingError{Param: "Bearing", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.VSpeed < 0 {
		return &ParsingError{Param: "VSpeed", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.VSpeed > 255 {
		return &ParsingError{Param: "VSpeed", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
