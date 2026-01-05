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

// AfAckInfo - Represents acknowledgement information of a traffic influence event notification.
type AfAckInfo struct {
	AfTransId string `json:"afTransId,omitempty"`

	AckResult AfResultInfo `json:"ackResult"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" validate:"regexp=^(msisdn-[0-9]{5,15}|extid-[^@]+@[^@]+|.+)$"`
}

// AssertAfAckInfoRequired checks if the required fields are not zero-ed
func AssertAfAckInfoRequired(obj AfAckInfo) error {
	elements := map[string]interface{}{
		"ackResult": obj.AckResult,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAfResultInfoRequired(obj.AckResult); err != nil {
		return err
	}
	return nil
}

// AssertAfAckInfoConstraints checks if the values respects the defined constraints
func AssertAfAckInfoConstraints(obj AfAckInfo) error {
	if err := AssertAfResultInfoConstraints(obj.AckResult); err != nil {
		return err
	}
	return nil
}
