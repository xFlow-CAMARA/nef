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
)

// AppliedParameterConfiguration - Represents the parameter configuration applied in the network.
type AppliedParameterConfiguration struct {

	// Each element uniquely identifies a user.
	ExternalIds []string `json:"externalIds,omitempty"`

	// Each element identifies the MS internal PSTN/ISDN number allocated for a UE.
	Msisdns []string `json:"msisdns,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumLatency int32 `json:"maximumLatency,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumResponseTime int32 `json:"maximumResponseTime,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	MaximumDetectionTime int32 `json:"maximumDetectionTime,omitempty"`
}

// AssertAppliedParameterConfigurationRequired checks if the required fields are not zero-ed
func AssertAppliedParameterConfigurationRequired(obj AppliedParameterConfiguration) error {
	return nil
}

// AssertAppliedParameterConfigurationConstraints checks if the values respects the defined constraints
func AssertAppliedParameterConfigurationConstraints(obj AppliedParameterConfiguration) error {
	if obj.MaximumLatency < 0 {
		return &ParsingError{Param: "MaximumLatency", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaximumResponseTime < 0 {
		return &ParsingError{Param: "MaximumResponseTime", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaximumDetectionTime < 0 {
		return &ParsingError{Param: "MaximumDetectionTime", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
