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

// TnapId - Contain the TNAP Identifier see clause5.6.2 of 3GPP TS 23.501.
type TnapId struct {

	// This IE shall be present if the UE is accessing the 5GC via a trusted WLAN access network.When present, it shall contain the SSID of the access point to which the UE is attached, that is received over NGAP,  see IEEE Std 802.11-2012.
	SsId string `json:"ssId,omitempty"`

	// When present, it shall contain the BSSID of the access point to which the UE is attached, that is received over NGAP, see IEEE Std 802.11-2012.
	BssId string `json:"bssId,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	CivicAddress string `json:"civicAddress,omitempty"`
}

// AssertTnapIdRequired checks if the required fields are not zero-ed
func AssertTnapIdRequired(obj TnapId) error {
	return nil
}

// AssertTnapIdConstraints checks if the values respects the defined constraints
func AssertTnapIdConstraints(obj TnapId) error {
	return nil
}
