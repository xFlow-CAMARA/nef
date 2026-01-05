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
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.3.0-alpha.5
 */

package models

// RouteToLocation - At least one of the \"routeInfo\" attribute and the \"routeProfId\" attribute shall be included in the \"RouteToLocation\" data type.
type RouteToLocation struct {

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai"`

	RouteInfo *RouteInformation `json:"routeInfo,omitempty"`

	// Identifies the routing profile Id.
	RouteProfId *string `json:"routeProfId,omitempty"`
}

// AssertRouteToLocationRequired checks if the required fields are not zero-ed
func AssertRouteToLocationRequired(obj RouteToLocation) error {
	elements := map[string]interface{}{
		"dnai": obj.Dnai,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.RouteInfo != nil {
		if err := AssertRouteInformationRequired(*obj.RouteInfo); err != nil {
			return err
		}
	}
	return nil
}

// AssertRouteToLocationConstraints checks if the values respects the defined constraints
func AssertRouteToLocationConstraints(obj RouteToLocation) error {
	if obj.RouteInfo != nil {
		if err := AssertRouteInformationConstraints(*obj.RouteInfo); err != nil {
			return err
		}
	}
	return nil
}
