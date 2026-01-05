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

// PointAltitudeUncertainty - Ellipsoid point with altitude and uncertainty ellipsoid.
type PointAltitudeUncertainty struct {
	GadShape

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`

	// Indicates value of confidence.
	VConfidence int32 `json:"vConfidence,omitempty"`
}

// AssertPointAltitudeUncertaintyRequired checks if the required fields are not zero-ed
func AssertPointAltitudeUncertaintyRequired(obj PointAltitudeUncertainty) error {
	elements := map[string]interface{}{
		"point":               obj.Point,
		"altitude":            obj.Altitude,
		"uncertaintyEllipse":  obj.UncertaintyEllipse,
		"uncertaintyAltitude": obj.UncertaintyAltitude,
		"confidence":          obj.Confidence,
		"shape":               obj.Shape,
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

// AssertPointAltitudeUncertaintyConstraints checks if the values respects the defined constraints
func AssertPointAltitudeUncertaintyConstraints(obj PointAltitudeUncertainty) error {
	if err := AssertGeographicalCoordinatesConstraints(obj.Point); err != nil {
		return err
	}
	if obj.Altitude < -32767 {
		return &ParsingError{Param: "Altitude", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Altitude > 32767 {
		return &ParsingError{Param: "Altitude", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertUncertaintyEllipseConstraints(obj.UncertaintyEllipse); err != nil {
		return err
	}
	if obj.UncertaintyAltitude < 0 {
		return &ParsingError{Param: "UncertaintyAltitude", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Confidence < 0 {
		return &ParsingError{Param: "Confidence", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Confidence > 100 {
		return &ParsingError{Param: "Confidence", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.VConfidence < 0 {
		return &ParsingError{Param: "VConfidence", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.VConfidence > 100 {
		return &ParsingError{Param: "VConfidence", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
