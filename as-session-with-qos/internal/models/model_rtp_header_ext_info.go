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

// RtpHeaderExtInfo - RTP Header Extension information
type RtpHeaderExtInfo struct {
	RtpHeaderExtType RtpHeaderExtType `json:"rtpHeaderExtType,omitempty"`

	RtpHeaderExtId int32 `json:"rtpHeaderExtId,omitempty"`

	LongFormat bool `json:"longFormat,omitempty"`

	PduSetSizeActive bool `json:"pduSetSizeActive,omitempty"`
}

// AssertRtpHeaderExtInfoRequired checks if the required fields are not zero-ed
func AssertRtpHeaderExtInfoRequired(obj RtpHeaderExtInfo) error {
	if err := AssertRtpHeaderExtTypeRequired(obj.RtpHeaderExtType); err != nil {
		return err
	}
	return nil
}

// AssertRtpHeaderExtInfoConstraints checks if the values respects the defined constraints
func AssertRtpHeaderExtInfoConstraints(obj RtpHeaderExtInfo) error {
	if err := AssertRtpHeaderExtTypeConstraints(obj.RtpHeaderExtType); err != nil {
		return err
	}
	if obj.RtpHeaderExtId < 1 {
		return &ParsingError{Param: "RtpHeaderExtId", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.RtpHeaderExtId > 255 {
		return &ParsingError{Param: "RtpHeaderExtId", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
