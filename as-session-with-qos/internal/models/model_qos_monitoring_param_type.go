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

// QosMonitoringParamType - Indicates the QoS monitoring parameter type.   Possible values are: - PACKET_DELAY: Indicates that the QoS monitoring parameter to be measured is packet delay. - CONGESTION: Indicates that the QoS monitoring parameter to be measured is congestion. - DATA_RATE: Indicates that the QoS monitoring parameter to be measured is data rate.
type QosMonitoringParamType struct {
}

// AssertQosMonitoringParamTypeRequired checks if the required fields are not zero-ed
func AssertQosMonitoringParamTypeRequired(obj QosMonitoringParamType) error {
	return nil
}

// AssertQosMonitoringParamTypeConstraints checks if the values respects the defined constraints
func AssertQosMonitoringParamTypeConstraints(obj QosMonitoringParamType) error {
	return nil
}
