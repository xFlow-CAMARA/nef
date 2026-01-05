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
 * 3gpp-ueid
 *
 * API for UE ID service. © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.1.0-alpha.2
 */

package models

import (
	"errors"
)

// UeIdReq - Represents the parameters to request the retrieval of AF specific UE ID.
type UeIdReq struct {
	AfId string `json:"afId"`

	// Unsigned integer with valid values between 0 and 65535.
	AppPortId int32 `json:"appPortId,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	IpDomain string `json:"ipDomain,omitempty"`

	// String uniquely identifying MTC provider information.
	MtcProviderId string `json:"mtcProviderId,omitempty"`

	// Unsigned integer with valid values between 0 and 65535.
	PortNumber int32 `json:"portNumber,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	UeIpAddr *IpAddr `json:"ueIpAddr,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	UeMacAddr string `json:"ueMacAddr,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat,omitempty"`
}

// AssertUeIdReqRequired checks if the required fields are not zero-ed
func AssertUeIdReqRequired(obj UeIdReq) error {
	elements := map[string]interface{}{
		"afId": obj.AfId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	if obj.UeIpAddr != nil {
		if err := AssertIpAddrRequired(*obj.UeIpAddr); err != nil {
			return err
		}
	}
	return nil
}

// AssertUeIdReqConstraints checks if the values respects the defined constraints
func AssertUeIdReqConstraints(obj UeIdReq) error {
	if obj.AppPortId < 0 {
		return &ParsingError{Param: "AppPortId", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AppPortId > 65535 {
		return &ParsingError{Param: "AppPortId", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.PortNumber < 0 {
		return &ParsingError{Param: "PortNumber", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PortNumber > 65535 {
		return &ParsingError{Param: "PortNumber", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertSnssaiConstraints(obj.Snssai); err != nil {
		return err
	}
	if obj.UeIpAddr != nil {
		if err := AssertIpAddrConstraints(*obj.UeIpAddr); err != nil {
			return err
		}
	}
	return nil
}
