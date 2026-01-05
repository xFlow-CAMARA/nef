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

// MonitoringNotification - Represents an event monitoring notification.
type MonitoringNotification struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Subscription string `json:"subscription"`

	// Each element identifies a notification of grouping configuration result.
	ConfigResults []ConfigResult `json:"configResults,omitempty"`

	// Monitoring event reports.
	MonitoringEventReports []MonitoringEventReport `json:"monitoringEventReports,omitempty"`

	// Identifies the added external Identifier(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	AddedExternalIds []string `json:"addedExternalIds,omitempty"`

	// Identifies the added MSISDN(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	AddedMsisdns []string `json:"addedMsisdns,omitempty"`

	// Identifies the cancelled external Identifier(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	CancelExternalIds []string `json:"cancelExternalIds,omitempty"`

	// Identifies the cancelled MSISDN(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	CancelMsisdns []string `json:"cancelMsisdns,omitempty"`

	// Indicates whether to request to cancel the corresponding monitoring subscription. Set to false or omitted otherwise.
	CancelInd bool `json:"cancelInd,omitempty"`

	AppliedParam AppliedParameterConfiguration `json:"appliedParam,omitempty"`
}

// AssertMonitoringNotificationRequired checks if the required fields are not zero-ed
func AssertMonitoringNotificationRequired(obj MonitoringNotification) error {
	elements := map[string]interface{}{
		"subscription": obj.Subscription,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.ConfigResults {
		if err := AssertConfigResultRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.MonitoringEventReports {
		if err := AssertMonitoringEventReportRequired(el); err != nil {
			return err
		}
	}
	if err := AssertAppliedParameterConfigurationRequired(obj.AppliedParam); err != nil {
		return err
	}
	return nil
}

// AssertMonitoringNotificationConstraints checks if the values respects the defined constraints
func AssertMonitoringNotificationConstraints(obj MonitoringNotification) error {
	for _, el := range obj.ConfigResults {
		if err := AssertConfigResultConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.MonitoringEventReports {
		if err := AssertMonitoringEventReportConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertAppliedParameterConfigurationConstraints(obj.AppliedParam); err != nil {
		return err
	}
	return nil
}
