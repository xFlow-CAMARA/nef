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

// EutraLocation - Contains the E-UTRA user location.
type EutraLocation struct {
	Tai Tai `json:"tai"`

	IgnoreTai bool `json:"ignoreTai,omitempty"`

	Ecgi Ecgi `json:"ecgi"`

	// This flag when present shall indicate that the Ecgi shall be ignored When present, it shall be set as follows: - true: ecgi shall be ignored. - false (default): ecgi shall not be ignored.
	IgnoreEcgi bool `json:"ignoreEcgi,omitempty"`

	// The value represents the elapsed time in minutes since the last network contact of the mobile station.  Value \"0\" indicates that the location information was obtained after a successful paging procedure for Active Location Retrieval when the UE is in idle mode or after a successful NG-RAN location reporting procedure with the eNB when the UE is in connected mode.  Any other value than \"0\" indicates that the location information is the last known one.  See 3GPP TS 29.002 clause 17.7.8.
	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`

	// Refer to geographical Information. See 3GPP TS 23.032 clause 7.3.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used.
	GeographicalInformation string `json:"geographicalInformation,omitempty" validate:"regexp=^[0-9A-F]{16}$"`

	// Refers to Calling Geodetic Location. See ITU-T Recommendation Q.763 (1999) [24] clause 3.88.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used.
	GeodeticInformation string `json:"geodeticInformation,omitempty" validate:"regexp=^[0-9A-F]{20}$"`

	GlobalNgenbId *GlobalRanNodeId `json:"globalNgenbId,omitempty"`

	GlobalENbId *GlobalRanNodeId `json:"globalENbId,omitempty"`
}

// AssertEutraLocationRequired checks if the required fields are not zero-ed
func AssertEutraLocationRequired(obj EutraLocation) error {
	elements := map[string]interface{}{
		"tai":  obj.Tai,
		"ecgi": obj.Ecgi,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertTaiRequired(obj.Tai); err != nil {
		return err
	}
	if err := AssertEcgiRequired(obj.Ecgi); err != nil {
		return err
	}
	if obj.GlobalNgenbId != nil {
		if err := AssertGlobalRanNodeIdRequired(*obj.GlobalNgenbId); err != nil {
			return err
		}
	}
	if obj.GlobalENbId != nil {
		if err := AssertGlobalRanNodeIdRequired(*obj.GlobalENbId); err != nil {
			return err
		}
	}
	return nil
}

// AssertEutraLocationConstraints checks if the values respects the defined constraints
func AssertEutraLocationConstraints(obj EutraLocation) error {
	if err := AssertTaiConstraints(obj.Tai); err != nil {
		return err
	}
	if err := AssertEcgiConstraints(obj.Ecgi); err != nil {
		return err
	}
	if obj.AgeOfLocationInformation < 0 {
		return &ParsingError{Param: "AgeOfLocationInformation", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AgeOfLocationInformation > 32767 {
		return &ParsingError{Param: "AgeOfLocationInformation", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.GlobalNgenbId != nil {
		if err := AssertGlobalRanNodeIdConstraints(*obj.GlobalNgenbId); err != nil {
			return err
		}
	}
	if obj.GlobalENbId != nil {
		if err := AssertGlobalRanNodeIdConstraints(*obj.GlobalENbId); err != nil {
			return err
		}
	}
	return nil
}
