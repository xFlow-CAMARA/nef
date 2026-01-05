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

// EllipsoidArc - Ellipsoid Arc.
type EllipsoidArc struct {
	GadShape

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of the inner radius.
	InnerRadius int32 `json:"innerRadius"`

	// Indicates value of uncertainty.
	UncertaintyRadius float32 `json:"uncertaintyRadius"`

	// Indicates value of angle.
	OffsetAngle int32 `json:"offsetAngle"`

	// Indicates value of angle.
	IncludedAngle int32 `json:"includedAngle"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// AssertEllipsoidArcRequired checks if the required fields are not zero-ed
func AssertEllipsoidArcRequired(obj EllipsoidArc) error {
	elements := map[string]interface{}{
		"point":             obj.Point,
		"innerRadius":       obj.InnerRadius,
		"uncertaintyRadius": obj.UncertaintyRadius,
		"offsetAngle":       obj.OffsetAngle,
		"includedAngle":     obj.IncludedAngle,
		"confidence":        obj.Confidence,
		"shape":             obj.Shape,
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
	return nil
}

// AssertEllipsoidArcConstraints checks if the values respects the defined constraints
func AssertEllipsoidArcConstraints(obj EllipsoidArc) error {
	if err := AssertGeographicalCoordinatesConstraints(obj.Point); err != nil {
		return err
	}
	if obj.InnerRadius < 0 {
		return &ParsingError{Param: "InnerRadius", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.InnerRadius > 327675 {
		return &ParsingError{Param: "InnerRadius", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.UncertaintyRadius < 0 {
		return &ParsingError{Param: "UncertaintyRadius", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.OffsetAngle < 0 {
		return &ParsingError{Param: "OffsetAngle", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.OffsetAngle > 360 {
		return &ParsingError{Param: "OffsetAngle", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.IncludedAngle < 0 {
		return &ParsingError{Param: "IncludedAngle", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.IncludedAngle > 360 {
		return &ParsingError{Param: "IncludedAngle", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.Confidence < 0 {
		return &ParsingError{Param: "Confidence", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Confidence > 100 {
		return &ParsingError{Param: "Confidence", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
