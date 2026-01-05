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

// Polygon - Polygon.
type Polygon struct {
	GadShape

	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`
}

// AssertPolygonRequired checks if the required fields are not zero-ed
func AssertPolygonRequired(obj Polygon) error {
	elements := map[string]interface{}{
		"pointList": obj.PointList,
		"shape":     obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGadShapeRequired(obj.GadShape); err != nil {
		return err
	}

	for _, el := range obj.PointList {
		if err := AssertGeographicalCoordinatesRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPolygonConstraints checks if the values respects the defined constraints
func AssertPolygonConstraints(obj Polygon) error {
	for _, el := range obj.PointList {
		if err := AssertGeographicalCoordinatesConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
