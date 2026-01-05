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

// AfEventSubscription - Describes the event information delivered in the subscription.
type AfEventSubscription struct {
	Event AfEvent `json:"event"`

	NotifMethod AfNotifMethod `json:"notifMethod,omitempty"`

	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty"`

	// indicating a time in seconds.
	WaitTime int32 `json:"waitTime,omitempty"`

	QosMonParamType QosMonitoringParamType `json:"qosMonParamType,omitempty"`
}

// AssertAfEventSubscriptionRequired checks if the required fields are not zero-ed
func AssertAfEventSubscriptionRequired(obj AfEventSubscription) error {
	elements := map[string]interface{}{
		"event": obj.Event,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAfEventRequired(obj.Event); err != nil {
		return err
	}
	if err := AssertAfNotifMethodRequired(obj.NotifMethod); err != nil {
		return err
	}
	if err := AssertQosMonitoringParamTypeRequired(obj.QosMonParamType); err != nil {
		return err
	}
	return nil
}

// AssertAfEventSubscriptionConstraints checks if the values respects the defined constraints
func AssertAfEventSubscriptionConstraints(obj AfEventSubscription) error {
	if err := AssertAfEventConstraints(obj.Event); err != nil {
		return err
	}
	if err := AssertAfNotifMethodConstraints(obj.NotifMethod); err != nil {
		return err
	}
	if err := AssertQosMonitoringParamTypeConstraints(obj.QosMonParamType); err != nil {
		return err
	}
	return nil
}
