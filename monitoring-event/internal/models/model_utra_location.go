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
	"time"
)

// UtraLocation - Exactly one of cgi, sai or lai shall be present.
type UtraLocation struct {
	Cgi CellGlobalId `json:"cgi,omitempty"`

	Sai ServiceAreaId `json:"sai,omitempty"`

	Lai LocationAreaId `json:"lai,omitempty"`

	Rai RoutingAreaId `json:"rai,omitempty"`

	// The value represents the elapsed time in minutes since the last network contact of the mobile station.  Value \"0\" indicates that the location information was obtained after a successful paging procedure for  Active Location Retrieval when the UE is in idle mode  or after a successful location reporting procedure  the UE is in connected mode. Any other value than \"0\" indicates that the location information is the last known one.  See 3GPP TS 29.002 clause 17.7.8.
	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`

	// Refer to geographical Information.See 3GPP TS 23.032 clause 7.3.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used.
	GeographicalInformation string `json:"geographicalInformation,omitempty"`

	// Refers to Calling Geodetic Location. See ITU-T Recommendation Q.763 (1999) clause 3.88.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used.
	GeodeticInformation string `json:"geodeticInformation,omitempty"`
}

// AssertUtraLocationRequired checks if the required fields are not zero-ed
func AssertUtraLocationRequired(obj UtraLocation) error {
	if err := AssertCellGlobalIdRequired(obj.Cgi); err != nil {
		return err
	}
	if err := AssertServiceAreaIdRequired(obj.Sai); err != nil {
		return err
	}
	if err := AssertLocationAreaIdRequired(obj.Lai); err != nil {
		return err
	}
	if err := AssertRoutingAreaIdRequired(obj.Rai); err != nil {
		return err
	}
	return nil
}

// AssertUtraLocationConstraints checks if the values respects the defined constraints
func AssertUtraLocationConstraints(obj UtraLocation) error {
	if err := AssertCellGlobalIdConstraints(obj.Cgi); err != nil {
		return err
	}
	if err := AssertServiceAreaIdConstraints(obj.Sai); err != nil {
		return err
	}
	if err := AssertLocationAreaIdConstraints(obj.Lai); err != nil {
		return err
	}
	if err := AssertRoutingAreaIdConstraints(obj.Rai); err != nil {
		return err
	}
	if obj.AgeOfLocationInformation < 0 {
		return &ParsingError{Param: "AgeOfLocationInformation", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AgeOfLocationInformation > 32767 {
		return &ParsingError{Param: "AgeOfLocationInformation", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
