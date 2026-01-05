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

// LocationQoS - QoS of Location request.
type LocationQoS struct {

	// Indicates value of accuracy.
	HAccuracy float32 `json:"hAccuracy,omitempty"`

	// Indicates value of accuracy.
	VAccuracy float32 `json:"vAccuracy,omitempty"`

	VerticalRequested bool `json:"verticalRequested,omitempty"`

	ResponseTime ResponseTime `json:"responseTime,omitempty"`

	MinorLocQoses []MinorLocationQoS `json:"minorLocQoses,omitempty"`

	LcsQosClass LcsQosClass `json:"lcsQosClass,omitempty"`
}

// AssertLocationQoSRequired checks if the required fields are not zero-ed
func AssertLocationQoSRequired(obj LocationQoS) error {
	if err := AssertResponseTimeRequired(obj.ResponseTime); err != nil {
		return err
	}
	for _, el := range obj.MinorLocQoses {
		if err := AssertMinorLocationQoSRequired(el); err != nil {
			return err
		}
	}
	if err := AssertLcsQosClassRequired(obj.LcsQosClass); err != nil {
		return err
	}
	return nil
}

// AssertLocationQoSConstraints checks if the values respects the defined constraints
func AssertLocationQoSConstraints(obj LocationQoS) error {
	if obj.HAccuracy < 0 {
		return &ParsingError{Param: "HAccuracy", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.VAccuracy < 0 {
		return &ParsingError{Param: "VAccuracy", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertResponseTimeConstraints(obj.ResponseTime); err != nil {
		return err
	}
	for _, el := range obj.MinorLocQoses {
		if err := AssertMinorLocationQoSConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertLcsQosClassConstraints(obj.LcsQosClass); err != nil {
		return err
	}
	return nil
}
