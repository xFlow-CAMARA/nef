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

import (
	"errors"
)

// AfRoutingRequirement - Describes AF requirements on routing traffic.
type AfRoutingRequirement struct {
	AppReloc bool `json:"appReloc,omitempty"`

	RouteToLocs []RouteToLocation `json:"routeToLocs,omitempty"`

	SpVal SpatialValidity `json:"spVal,omitempty"`

	TempVals []TemporalValidity `json:"tempVals,omitempty"`

	UpPathChgSub *UpPathChgEvent `json:"upPathChgSub,omitempty"`

	AddrPreserInd bool `json:"addrPreserInd,omitempty"`

	// Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.
	SimConnInd bool `json:"simConnInd,omitempty"`

	// indicating a time in seconds.
	SimConnTerm int32 `json:"simConnTerm,omitempty"`

	// Contains EAS IP replacement information.
	EasIpReplaceInfos []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`

	// Indicates the EAS rediscovery is required.
	EasRedisInd bool `json:"easRedisInd,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxAllowedUpLat int32 `json:"maxAllowedUpLat,omitempty"`

	TfcCorreInfo *TrafficCorrelationInfo `json:"tfcCorreInfo,omitempty"`
}

// AssertAfRoutingRequirementRequired checks if the required fields are not zero-ed
func AssertAfRoutingRequirementRequired(obj AfRoutingRequirement) error {
	for _, el := range obj.RouteToLocs {
		if err := AssertRouteToLocationRequired(el); err != nil {
			return err
		}
	}
	if err := AssertSpatialValidityRequired(obj.SpVal); err != nil {
		return err
	}
	for _, el := range obj.TempVals {
		if err := AssertTemporalValidityRequired(el); err != nil {
			return err
		}
	}
	if obj.UpPathChgSub != nil {
		if err := AssertUpPathChgEventRequired(*obj.UpPathChgSub); err != nil {
			return err
		}
	}
	for _, el := range obj.EasIpReplaceInfos {
		if err := AssertEasIpReplacementInfoRequired(el); err != nil {
			return err
		}
	}
	if obj.TfcCorreInfo != nil {
		if err := AssertTrafficCorrelationInfoRequired(*obj.TfcCorreInfo); err != nil {
			return err
		}
	}
	return nil
}

// AssertAfRoutingRequirementConstraints checks if the values respects the defined constraints
func AssertAfRoutingRequirementConstraints(obj AfRoutingRequirement) error {
	for _, el := range obj.RouteToLocs {
		if err := AssertRouteToLocationConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertSpatialValidityConstraints(obj.SpVal); err != nil {
		return err
	}
	for _, el := range obj.TempVals {
		if err := AssertTemporalValidityConstraints(el); err != nil {
			return err
		}
	}
	if obj.UpPathChgSub != nil {
		if err := AssertUpPathChgEventConstraints(*obj.UpPathChgSub); err != nil {
			return err
		}
	}
	for _, el := range obj.EasIpReplaceInfos {
		if err := AssertEasIpReplacementInfoConstraints(el); err != nil {
			return err
		}
	}
	if obj.MaxAllowedUpLat < 0 {
		return &ParsingError{Param: "MaxAllowedUpLat", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TfcCorreInfo != nil {
		if err := AssertTrafficCorrelationInfoConstraints(*obj.TfcCorreInfo); err != nil {
			return err
		}
	}
	return nil
}
