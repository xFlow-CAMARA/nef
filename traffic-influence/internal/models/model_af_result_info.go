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

// AfResultInfo - Identifies the result of application layer handling.
type AfResultInfo struct {
	AfStatus AfResultStatus `json:"afStatus"`

	TrafficRoute *RouteToLocation `json:"trafficRoute,omitempty"`

	// If present and set to \"true\" it indicates that buffering of uplink traffic to the target DNAI is needed.
	UpBuffInd bool `json:"upBuffInd,omitempty"`

	// Contains EAS IP replacement information.
	EasIpReplaceInfos []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
}

// AssertAfResultInfoRequired checks if the required fields are not zero-ed
func AssertAfResultInfoRequired(obj AfResultInfo) error {
	elements := map[string]interface{}{
		"afStatus": obj.AfStatus,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAfResultStatusRequired(obj.AfStatus); err != nil {
		return err
	}
	if obj.TrafficRoute != nil {
		if err := AssertRouteToLocationRequired(*obj.TrafficRoute); err != nil {
			return err
		}
	}
	for _, el := range obj.EasIpReplaceInfos {
		if err := AssertEasIpReplacementInfoRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAfResultInfoConstraints checks if the values respects the defined constraints
func AssertAfResultInfoConstraints(obj AfResultInfo) error {
	if err := AssertAfResultStatusConstraints(obj.AfStatus); err != nil {
		return err
	}
	if obj.TrafficRoute != nil {
		if err := AssertRouteToLocationConstraints(*obj.TrafficRoute); err != nil {
			return err
		}
	}
	for _, el := range obj.EasIpReplaceInfos {
		if err := AssertEasIpReplacementInfoConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
