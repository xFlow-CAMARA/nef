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
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.4
 */

package models

import (
	"errors"
)

// UncertaintyEllipsoid - Ellipsoid with uncertainty
type UncertaintyEllipsoid struct {

	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor"`

	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor"`

	// Indicates value of uncertainty.
	Vertical float32 `json:"vertical"`

	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor"`
}

// AssertUncertaintyEllipsoidRequired checks if the required fields are not zero-ed
func AssertUncertaintyEllipsoidRequired(obj UncertaintyEllipsoid) error {
	elements := map[string]interface{}{
		"semiMajor":        obj.SemiMajor,
		"semiMinor":        obj.SemiMinor,
		"vertical":         obj.Vertical,
		"orientationMajor": obj.OrientationMajor,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertUncertaintyEllipsoidConstraints checks if the values respects the defined constraints
func AssertUncertaintyEllipsoidConstraints(obj UncertaintyEllipsoid) error {
	if obj.SemiMajor < 0 {
		return &ParsingError{Param: "SemiMajor", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SemiMinor < 0 {
		return &ParsingError{Param: "SemiMinor", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Vertical < 0 {
		return &ParsingError{Param: "Vertical", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.OrientationMajor < 0 {
		return &ParsingError{Param: "OrientationMajor", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.OrientationMajor > 180 {
		return &ParsingError{Param: "OrientationMajor", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
