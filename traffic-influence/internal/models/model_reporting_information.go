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

import (
	"errors"
	"time"
)

// ReportingInformation - Represents the type of reporting that the subscription requires.
type ReportingInformation struct {
	ImmRep bool `json:"immRep,omitempty"`

	NotifMethod NotificationMethod `json:"notifMethod,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr int32 `json:"maxReportNbr,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	MonDur time.Time `json:"monDur,omitempty"`

	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	SampRatio int32 `json:"sampRatio,omitempty"`

	// Criteria for partitioning the UEs before applying the sampling ratio.
	PartitionCriteria []PartitioningCriteria `json:"partitionCriteria,omitempty"`

	// indicating a time in seconds.
	GrpRepTime int32 `json:"grpRepTime,omitempty"`

	NotifFlag NotificationFlag `json:"notifFlag,omitempty"`

	NotifFlagInstruct MutingExceptionInstructions `json:"notifFlagInstruct,omitempty"`

	MutingSetting MutingNotificationsSettings `json:"mutingSetting,omitempty"`
}

// AssertReportingInformationRequired checks if the required fields are not zero-ed
func AssertReportingInformationRequired(obj ReportingInformation) error {
	if err := AssertNotificationMethodRequired(obj.NotifMethod); err != nil {
		return err
	}
	for _, el := range obj.PartitionCriteria {
		if err := AssertPartitioningCriteriaRequired(el); err != nil {
			return err
		}
	}
	if err := AssertNotificationFlagRequired(obj.NotifFlag); err != nil {
		return err
	}
	if err := AssertMutingExceptionInstructionsRequired(obj.NotifFlagInstruct); err != nil {
		return err
	}
	if err := AssertMutingNotificationsSettingsRequired(obj.MutingSetting); err != nil {
		return err
	}
	return nil
}

// AssertReportingInformationConstraints checks if the values respects the defined constraints
func AssertReportingInformationConstraints(obj ReportingInformation) error {
	if err := AssertNotificationMethodConstraints(obj.NotifMethod); err != nil {
		return err
	}
	if obj.MaxReportNbr < 0 {
		return &ParsingError{Param: "MaxReportNbr", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SampRatio < 1 {
		return &ParsingError{Param: "SampRatio", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SampRatio > 100 {
		return &ParsingError{Param: "SampRatio", Err: errors.New(errMsgMaxValueConstraint)}
	}
	for _, el := range obj.PartitionCriteria {
		if err := AssertPartitioningCriteriaConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertNotificationFlagConstraints(obj.NotifFlag); err != nil {
		return err
	}
	if err := AssertMutingExceptionInstructionsConstraints(obj.NotifFlagInstruct); err != nil {
		return err
	}
	if err := AssertMutingNotificationsSettingsConstraints(obj.MutingSetting); err != nil {
		return err
	}
	return nil
}
