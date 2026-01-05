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

// UpPathChgEvent - Contains the UP path change event subscription from the AF.
type UpPathChgEvent struct {

	// String providing an URI formatted according to RFC 3986.
	NotificationUri string `json:"notificationUri"`

	// It is used to set the value of Notification Correlation ID in the notification sent by the SMF.
	NotifCorreId string `json:"notifCorreId"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType"`

	AfAckInd bool `json:"afAckInd,omitempty"`
}

// AssertUpPathChgEventRequired checks if the required fields are not zero-ed
func AssertUpPathChgEventRequired(obj UpPathChgEvent) error {
	elements := map[string]interface{}{
		"notificationUri": obj.NotificationUri,
		"notifCorreId":    obj.NotifCorreId,
		"dnaiChgType":     obj.DnaiChgType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDnaiChangeTypeRequired(obj.DnaiChgType); err != nil {
		return err
	}
	return nil
}

// AssertUpPathChgEventConstraints checks if the values respects the defined constraints
func AssertUpPathChgEventConstraints(obj UpPathChgEvent) error {
	if err := AssertDnaiChangeTypeConstraints(obj.DnaiChgType); err != nil {
		return err
	}
	return nil
}
