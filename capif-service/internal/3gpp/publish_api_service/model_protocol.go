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
CAPIF_Publish_Service_API

API for publishing service APIs.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.5
*/

package publish_api_service

// Protocol Indicates a protocol and protocol version used by the API.   Possible values are: - HTTP_1_1: Indicates that the protocol is HTTP version 1.1. - HTTP_2: Indicates that the protocol is HTTP version 2. - MQTT: Indicates that the protocol is Message Queuing Telemetry Transport. - WEBSOCKET: Indicates that the protocol is Websocket.
type Protocol string

const (
	HTTP_1_1 = "HTTP_1_1"
	HTTP_2   = "HTTP_2"
)
