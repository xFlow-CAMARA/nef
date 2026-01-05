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

// IdleStatusInfo - Represents the information relevant to when the UE transitions into idle mode.
type IdleStatusInfo struct {

	// Unsigned integer identifying a period of time in units of seconds.
	ActiveTime int32 `json:"activeTime,omitempty"`

	EdrxCycleLength float32 `json:"edrxCycleLength,omitempty"`

	// Identifies the number of packets shall be buffered in the serving gateway. It shall be present if the idle status indication is requested by the SCS/AS with \"idleStatusIndication\" in the \"monitoringEventSubscription\" sets to \"true\".
	SuggestedNumberOfDlPackets int32 `json:"suggestedNumberOfDlPackets,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	IdleStatusTimestamp time.Time `json:"idleStatusTimestamp,omitempty"`

	// Unsigned integer identifying a period of time in units of seconds.
	PeriodicAUTimer int32 `json:"periodicAUTimer,omitempty"`
}

// AssertIdleStatusInfoRequired checks if the required fields are not zero-ed
func AssertIdleStatusInfoRequired(obj IdleStatusInfo) error {
	return nil
}

// AssertIdleStatusInfoConstraints checks if the values respects the defined constraints
func AssertIdleStatusInfoConstraints(obj IdleStatusInfo) error {
	if obj.ActiveTime < 0 {
		return &ParsingError{Param: "ActiveTime", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.EdrxCycleLength < 0 {
		return &ParsingError{Param: "EdrxCycleLength", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SuggestedNumberOfDlPackets < 0 {
		return &ParsingError{Param: "SuggestedNumberOfDlPackets", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PeriodicAUTimer < 0 {
		return &ParsingError{Param: "PeriodicAUTimer", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
