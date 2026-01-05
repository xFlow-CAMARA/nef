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

// MonitoringEventReports - Represents a set of event monitoring reports.
type MonitoringEventReports struct {
	MonitoringEventReports []MonitoringEventReport `json:"monitoringEventReports"`
}

// AssertMonitoringEventReportsRequired checks if the required fields are not zero-ed
func AssertMonitoringEventReportsRequired(obj MonitoringEventReports) error {
	elements := map[string]interface{}{
		"monitoringEventReports": obj.MonitoringEventReports,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.MonitoringEventReports {
		if err := AssertMonitoringEventReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMonitoringEventReportsConstraints checks if the values respects the defined constraints
func AssertMonitoringEventReportsConstraints(obj MonitoringEventReports) error {
	for _, el := range obj.MonitoringEventReports {
		if err := AssertMonitoringEventReportConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
