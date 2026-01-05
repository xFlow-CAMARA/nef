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

// RangeDirection - Represents a distance and direction from a point A to a point B.
type RangeDirection struct {
	Distance float32 `json:"distance,omitempty"`

	// Indicates value of angle.
	AzimuthDirection int32 `json:"azimuthDirection,omitempty"`

	// Indicates value of angle.
	ElevationDirection int32 `json:"elevationDirection,omitempty"`
}

// AssertRangeDirectionRequired checks if the required fields are not zero-ed
func AssertRangeDirectionRequired(obj RangeDirection) error {
	return nil
}

// AssertRangeDirectionConstraints checks if the values respects the defined constraints
func AssertRangeDirectionConstraints(obj RangeDirection) error {
	if obj.AzimuthDirection < 0 {
		return &ParsingError{Param: "AzimuthDirection", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AzimuthDirection > 360 {
		return &ParsingError{Param: "AzimuthDirection", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.ElevationDirection < 0 {
		return &ParsingError{Param: "ElevationDirection", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ElevationDirection > 360 {
		return &ParsingError{Param: "ElevationDirection", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
