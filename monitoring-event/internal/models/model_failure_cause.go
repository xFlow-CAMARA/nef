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

// FailureCause - Represents the reason of communication failure.
type FailureCause struct {

	// Identifies a non-transparent copy of the BSSGP cause code. Refer to 3GPP TS 29.128.
	BssgpCause int32 `json:"bssgpCause,omitempty"`

	// Identify the type of the S1AP-Cause. Refer to 3GPP TS 29.128.
	CauseType int32 `json:"causeType,omitempty"`

	// Identifies a non-transparent copy of the GMM cause code. Refer to 3GPP TS 29.128.
	GmmCause int32 `json:"gmmCause,omitempty"`

	// Identifies a non-transparent copy of the RANAP cause code. Refer to 3GPP TS 29.128.
	RanapCause int32 `json:"ranapCause,omitempty"`

	// Indicates RAN and/or NAS release cause code information, TWAN release cause code information or untrusted WLAN release cause code information. Refer to 3GPP TS 29.214.
	RanNasCause string `json:"ranNasCause,omitempty"`

	// Identifies a non-transparent copy of the S1AP cause code. Refer to 3GPP TS 29.128.
	S1ApCause int32 `json:"s1ApCause,omitempty"`

	// Identifies a non-transparent copy of the SM cause code. Refer to 3GPP TS 29.128.
	SmCause int32 `json:"smCause,omitempty"`
}

// AssertFailureCauseRequired checks if the required fields are not zero-ed
func AssertFailureCauseRequired(obj FailureCause) error {
	return nil
}

// AssertFailureCauseConstraints checks if the values respects the defined constraints
func AssertFailureCauseConstraints(obj FailureCause) error {
	return nil
}
