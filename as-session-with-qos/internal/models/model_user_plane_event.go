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

// UserPlaneEvent - Represents the user plane event.   Possible values are: - SESSION_TERMINATION: Indicates that Rx session is terminated. - LOSS_OF_BEARER : Indicates a loss of a bearer. - RECOVERY_OF_BEARER: Indicates a recovery of a bearer. - RELEASE_OF_BEARER: Indicates a release of a bearer. - USAGE_REPORT: Indicates the usage report event. - FAILED_RESOURCES_ALLOCATION: Indicates the resource allocation is failed. - QOS_GUARANTEED: The QoS targets of one or more SDFs are guaranteed again. - QOS_NOT_GUARANTEED: The QoS targets of one or more SDFs are not being guaranteed. - QOS_MONITORING: Indicates a QoS monitoring event. - SUCCESSFUL_RESOURCES_ALLOCATION: Indicates the resource allocation is successful. - ACCESS_TYPE_CHANGE: Indicates an Access type change. - PLMN_CHG: Indicates a PLMN change. - L4S_NOT_AVAILABLE: The ECN marking for L4S of one or more SDFs is not available. - L4S_AVAILABLE: The ECN marking for L4S of one or more SDFs is available again. - BAT_OFFSET_INFO: Indicates the network provided BAT offset and the optionally adjusted periodicity. - RT_DELAY_TWO_QOS_FLOWS: Indicates round-trip delay on UL and DL flows over two QoS flows. - PACK_DELAY_VAR: Indicates Packet Delay Variation is enabled for the SDF.
type UserPlaneEvent struct {
}

// AssertUserPlaneEventRequired checks if the required fields are not zero-ed
func AssertUserPlaneEventRequired(obj UserPlaneEvent) error {
	return nil
}

// AssertUserPlaneEventConstraints checks if the values respects the defined constraints
func AssertUserPlaneEventConstraints(obj UserPlaneEvent) error {
	return nil
}
