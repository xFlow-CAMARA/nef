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

import "fmt"

// SupportedGadShapes - Indicates supported GAD shapes.
type SupportedGadShapes string

const (
	SupportedGadShapesPOINT                                      SupportedGadShapes = "POINT"
	SupportedGadShapesPOINT_UNCERTAINTY_CIRCLE                   SupportedGadShapes = "POINT_UNCERTAINTY_CIRCLE"
	SupportedGadShapesPOINT_UNCERTAINTY_ELLIPSE                  SupportedGadShapes = "POINT_UNCERTAINTY_ELLIPSE"
	SupportedGadShapesPOLYGON                                    SupportedGadShapes = "POLYGON"
	SupportedGadShapesPOINT_ALTITUDE                             SupportedGadShapes = "POINT_ALTITUDE"
	SupportedGadShapesPOINT_ALTITUDE_UNCERTAINTY                 SupportedGadShapes = "POINT_ALTITUDE_UNCERTAINTY"
	SupportedGadShapesELLIPSOID_ARC                              SupportedGadShapes = "ELLIPSOID_ARC"
	SupportedGadShapesLOCAL_2D_POINT_UNCERTAINTY_ELLIPSE         SupportedGadShapes = "LOCAL_2D_POINT_UNCERTAINTY_ELLIPSE"
	SupportedGadShapesLOCAL_3D_POINT_UNCERTAINTY_ELLIPSOID       SupportedGadShapes = "LOCAL_3D_POINT_UNCERTAINTY_ELLIPSOID"
	SupportedGadShapesDISTANCE_DIRECTION                         SupportedGadShapes = "DISTANCE_DIRECTION"
	SupportedGadShapesRELATIVE_2D_LOCATION_UNCERTAINTY_ELLIPSE   SupportedGadShapes = "RELATIVE_2D_LOCATION_UNCERTAINTY_ELLIPSE"
	SupportedGadShapesRELATIVE_3D_LOCATION_UNCERTAINTY_ELLIPSOID SupportedGadShapes = "RELATIVE_3D_LOCATION_UNCERTAINTY_ELLIPSOID"
)

// validSupportedGadShapes is the set of allowed enum values.
// We keep it as a map for O(1) lookup in validation.
var validSupportedGadShapes = map[SupportedGadShapes]struct{}{
	SupportedGadShapesPOINT:                                      {},
	SupportedGadShapesPOINT_UNCERTAINTY_CIRCLE:                   {},
	SupportedGadShapesPOINT_UNCERTAINTY_ELLIPSE:                  {},
	SupportedGadShapesPOLYGON:                                    {},
	SupportedGadShapesPOINT_ALTITUDE:                             {},
	SupportedGadShapesPOINT_ALTITUDE_UNCERTAINTY:                 {},
	SupportedGadShapesELLIPSOID_ARC:                              {},
	SupportedGadShapesLOCAL_2D_POINT_UNCERTAINTY_ELLIPSE:         {},
	SupportedGadShapesLOCAL_3D_POINT_UNCERTAINTY_ELLIPSOID:       {},
	SupportedGadShapesDISTANCE_DIRECTION:                         {},
	SupportedGadShapesRELATIVE_2D_LOCATION_UNCERTAINTY_ELLIPSE:   {},
	SupportedGadShapesRELATIVE_3D_LOCATION_UNCERTAINTY_ELLIPSOID: {},
}

// AssertSupportedGadShapesRequired checks if the required fields are not zero-ed.
// For an enum-like string type, "zero-ed" means "" (empty string).
func AssertSupportedGadShapesRequired(obj SupportedGadShapes) error {
	if obj == "" {
		return fmt.Errorf("shape is required and cannot be empty")
	}
	return nil
}

// AssertSupportedGadShapesConstraints checks if the value respects the defined constraints
// (i.e. it must be one of the known constants).
func AssertSupportedGadShapesConstraints(obj SupportedGadShapes) error {
	if _, ok := validSupportedGadShapes[obj]; !ok {
		return fmt.Errorf("unsupported GAD shape: %q", obj)
	}
	return nil
}

// IsValid is a convenience method you can call directly.
func (s SupportedGadShapes) IsValid() bool {
	_, ok := validSupportedGadShapes[s]
	return ok
}
