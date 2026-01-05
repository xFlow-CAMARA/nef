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

// UserPlaneNotificationData - Represents the parameters to be conveyed in a user plane event(s) notification.
type UserPlaneNotificationData struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Transaction string `json:"transaction"`

	// Contains the reported event and applicable information
	EventReports []UserPlaneEventReport `json:"eventReports"`
}

// AssertUserPlaneNotificationDataRequired checks if the required fields are not zero-ed
func AssertUserPlaneNotificationDataRequired(obj UserPlaneNotificationData) error {
	elements := map[string]interface{}{
		"transaction":  obj.Transaction,
		"eventReports": obj.EventReports,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.EventReports {
		if err := AssertUserPlaneEventReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertUserPlaneNotificationDataConstraints checks if the values respects the defined constraints
func AssertUserPlaneNotificationDataConstraints(obj UserPlaneNotificationData) error {
	for _, el := range obj.EventReports {
		if err := AssertUserPlaneEventReportConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
