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

// PointUncertaintyEllipse - Ellipsoid point with uncertainty ellipse.
type PointUncertaintyEllipse struct {
	GadShape

	Point GeographicalCoordinates `json:"point"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// AssertPointUncertaintyEllipseRequired checks if the required fields are not zero-ed
func AssertPointUncertaintyEllipseRequired(obj PointUncertaintyEllipse) error {
	elements := map[string]interface{}{
		"point":              obj.Point,
		"uncertaintyEllipse": obj.UncertaintyEllipse,
		"confidence":         obj.Confidence,
		"shape":              obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGadShapeRequired(obj.GadShape); err != nil {
		return err
	}

	if err := AssertGeographicalCoordinatesRequired(obj.Point); err != nil {
		return err
	}
	if err := AssertUncertaintyEllipseRequired(obj.UncertaintyEllipse); err != nil {
		return err
	}
	return nil
}

// AssertPointUncertaintyEllipseConstraints checks if the values respects the defined constraints
func AssertPointUncertaintyEllipseConstraints(obj PointUncertaintyEllipse) error {
	if err := AssertGeographicalCoordinatesConstraints(obj.Point); err != nil {
		return err
	}
	if err := AssertUncertaintyEllipseConstraints(obj.UncertaintyEllipse); err != nil {
		return err
	}
	if obj.Confidence < 0 {
		return &ParsingError{Param: "Confidence", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Confidence > 100 {
		return &ParsingError{Param: "Confidence", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
