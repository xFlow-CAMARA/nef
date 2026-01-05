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

// Model2DRelativeLocation - Represents a relative 2D location with uncertainty ellipse.
type Model2DRelativeLocation struct {

	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor,omitempty"`

	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor,omitempty"`

	// Indicates value of angle.
	OrientationAngle int32 `json:"orientationAngle,omitempty"`
}

// AssertModel2DRelativeLocationRequired checks if the required fields are not zero-ed
func AssertModel2DRelativeLocationRequired(obj Model2DRelativeLocation) error {
	return nil
}

// AssertModel2DRelativeLocationConstraints checks if the values respects the defined constraints
func AssertModel2DRelativeLocationConstraints(obj Model2DRelativeLocation) error {
	if obj.SemiMinor < 0 {
		return &ParsingError{Param: "SemiMinor", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SemiMajor < 0 {
		return &ParsingError{Param: "SemiMajor", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.OrientationAngle < 0 {
		return &ParsingError{Param: "OrientationAngle", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.OrientationAngle > 360 {
		return &ParsingError{Param: "OrientationAngle", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
