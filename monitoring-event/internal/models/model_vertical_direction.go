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
	"fmt"
)

// VerticalDirection : Indicates direction of vertical speed.
type VerticalDirection string

// List of VerticalDirection
const (
	UPWARD   VerticalDirection = "UPWARD"
	DOWNWARD VerticalDirection = "DOWNWARD"
)

// AllowedVerticalDirectionEnumValues is all the allowed values of VerticalDirection enum
var AllowedVerticalDirectionEnumValues = []VerticalDirection{
	"UPWARD",
	"DOWNWARD",
}

// validVerticalDirectionEnumValue provides a map of VerticalDirections for fast verification of use input
var validVerticalDirectionEnumValues = map[VerticalDirection]struct{}{
	"UPWARD":   {},
	"DOWNWARD": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerticalDirection) IsValid() bool {
	_, ok := validVerticalDirectionEnumValues[v]
	return ok
}

// NewVerticalDirectionFromValue returns a pointer to a valid VerticalDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerticalDirectionFromValue(v string) (VerticalDirection, error) {
	ev := VerticalDirection(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for VerticalDirection: valid values are %v", v, AllowedVerticalDirectionEnumValues)
}

// AssertVerticalDirectionRequired checks if the required fields are not zero-ed
func AssertVerticalDirectionRequired(obj VerticalDirection) error {
	return nil
}

// AssertVerticalDirectionConstraints checks if the values respects the defined constraints
func AssertVerticalDirectionConstraints(obj VerticalDirection) error {
	return nil
}
