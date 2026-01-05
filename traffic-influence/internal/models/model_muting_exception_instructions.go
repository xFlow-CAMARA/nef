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

// MutingExceptionInstructions - Indicates to an Event producer NF instructions for the subscription and stored events when an exception (e.g. full buffer) occurs at the Event producer NF while the event is muted.
type MutingExceptionInstructions struct {
	BufferedNotifs BufferedNotificationsAction `json:"bufferedNotifs,omitempty"`

	Subscription SubscriptionAction `json:"subscription,omitempty"`
}

// AssertMutingExceptionInstructionsRequired checks if the required fields are not zero-ed
func AssertMutingExceptionInstructionsRequired(obj MutingExceptionInstructions) error {
	if err := AssertBufferedNotificationsActionRequired(obj.BufferedNotifs); err != nil {
		return err
	}
	if err := AssertSubscriptionActionRequired(obj.Subscription); err != nil {
		return err
	}
	return nil
}

// AssertMutingExceptionInstructionsConstraints checks if the values respects the defined constraints
func AssertMutingExceptionInstructionsConstraints(obj MutingExceptionInstructions) error {
	if err := AssertBufferedNotificationsActionConstraints(obj.BufferedNotifs); err != nil {
		return err
	}
	if err := AssertSubscriptionActionConstraints(obj.Subscription); err != nil {
		return err
	}
	return nil
}
