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

// ConsentRevocNotif - Represents the user consent revocation information conveyed in a user consent revocation notification.
type ConsentRevocNotif struct {
	SubscriptionId string `json:"subscriptionId"`

	ConsentsRevoked []ConsentRevoked `json:"consentsRevoked"`
}

// AssertConsentRevocNotifRequired checks if the required fields are not zero-ed
func AssertConsentRevocNotifRequired(obj ConsentRevocNotif) error {
	elements := map[string]interface{}{
		"subscriptionId":  obj.SubscriptionId,
		"consentsRevoked": obj.ConsentsRevoked,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.ConsentsRevoked {
		if err := AssertConsentRevokedRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertConsentRevocNotifConstraints checks if the values respects the defined constraints
func AssertConsentRevocNotifConstraints(obj ConsentRevocNotif) error {
	for _, el := range obj.ConsentsRevoked {
		if err := AssertConsentRevokedConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
