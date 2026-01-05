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

// WebsockNotifConfig - Represents the configuration information for the delivery of notifications over Websockets.
type WebsockNotifConfig struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	WebsocketUri string `json:"websocketUri,omitempty"`

	// Set by the SCS/AS to indicate that the Websocket delivery is requested.
	RequestWebsocketUri bool `json:"requestWebsocketUri,omitempty"`
}

// AssertWebsockNotifConfigRequired checks if the required fields are not zero-ed
func AssertWebsockNotifConfigRequired(obj WebsockNotifConfig) error {
	return nil
}

// AssertWebsockNotifConfigConstraints checks if the values respects the defined constraints
func AssertWebsockNotifConfigConstraints(obj WebsockNotifConfig) error {
	return nil
}
