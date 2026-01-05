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

// UserPlaneEventReport - Represents an event report for user plane.
type UserPlaneEventReport struct {
	Event UserPlaneEvent `json:"event"`

	AccumulatedUsage AccumulatedUsage `json:"accumulatedUsage,omitempty"`

	// Identifies the affected flows that were sent during event subscription. It might be omitted when the reported event applies to all the flows sent during the subscription.
	FlowIds []int32 `json:"flowIds,omitempty"`

	// Identifies the the flow filters for the single-modal data flows thatwere sent during event subscription. It may be omitted when the reported event applies to all the single-modal data flows sent during the subscription.
	MultiModFlows []MultiModalFlows `json:"multiModFlows,omitempty"`

	// The currently applied QoS reference. Applicable for event QOS_NOT_GUARANTEED or SUCCESSFUL_RESOURCES_ALLOCATION.
	AppliedQosRef string `json:"appliedQosRef,omitempty"`

	// When present and set to true it indicates that the Alternative QoS profiles are not supported by the access network. Applicable for event QOS_NOT_GUARANTEED.
	AltQosNotSuppInd bool `json:"altQosNotSuppInd,omitempty"`

	PlmnId PlmnIdNid `json:"plmnId,omitempty"`

	// Contains the QoS Monitoring Reporting information
	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`

	// Contains the PDV Monitoring Reporting information
	PdvMonReports []PdvMonitoringReport `json:"pdvMonReports,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	BatOffsetInfo BatOffsetInfo `json:"batOffsetInfo,omitempty"`

	// Contains the round trip delay over two SDFs reporting information
	RttMonReports []QosMonitoringReport `json:"rttMonReports,omitempty"`

	// Contains QoS Monitoring for data rate information. It shall be present when the notified event is \"QOS_MONITORING\" and data rate measurements are available.
	QosMonDatRateReps []QosMonitoringReport1 `json:"qosMonDatRateReps,omitempty"`

	// Contains QoS Monitoring for aggregated data rate information. It shall be present when the notified event is \"QOS_MONITORING\" and data rate measurements are available.
	AggrDataRateRpts []QosMonitoringReport1 `json:"aggrDataRateRpts,omitempty"`

	// Contains QoS Monitoring for congestion information. It shall be present when the notified event is \"QOS_MONITORING\" and congestion measurements are available.
	QosMonConInfoReps []QosMonitoringReport `json:"qosMonConInfoReps,omitempty"`
}

// AssertUserPlaneEventReportRequired checks if the required fields are not zero-ed
func AssertUserPlaneEventReportRequired(obj UserPlaneEventReport) error {
	elements := map[string]interface{}{
		"event": obj.Event,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUserPlaneEventRequired(obj.Event); err != nil {
		return err
	}
	if err := AssertAccumulatedUsageRequired(obj.AccumulatedUsage); err != nil {
		return err
	}
	for _, el := range obj.MultiModFlows {
		if err := AssertMultiModalFlowsRequired(el); err != nil {
			return err
		}
	}
	if err := AssertPlmnIdNidRequired(obj.PlmnId); err != nil {
		return err
	}
	for _, el := range obj.QosMonReports {
		if err := AssertQosMonitoringReportRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.PdvMonReports {
		if err := AssertPdvMonitoringReportRequired(el); err != nil {
			return err
		}
	}
	if err := AssertRatTypeRequired(obj.RatType); err != nil {
		return err
	}
	if err := AssertBatOffsetInfoRequired(obj.BatOffsetInfo); err != nil {
		return err
	}
	for _, el := range obj.RttMonReports {
		if err := AssertQosMonitoringReportRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.QosMonDatRateReps {
		if err := AssertQosMonitoringReport1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AggrDataRateRpts {
		if err := AssertQosMonitoringReport1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.QosMonConInfoReps {
		if err := AssertQosMonitoringReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertUserPlaneEventReportConstraints checks if the values respects the defined constraints
func AssertUserPlaneEventReportConstraints(obj UserPlaneEventReport) error {
	if err := AssertUserPlaneEventConstraints(obj.Event); err != nil {
		return err
	}
	if err := AssertAccumulatedUsageConstraints(obj.AccumulatedUsage); err != nil {
		return err
	}
	for _, el := range obj.MultiModFlows {
		if err := AssertMultiModalFlowsConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertPlmnIdNidConstraints(obj.PlmnId); err != nil {
		return err
	}
	for _, el := range obj.QosMonReports {
		if err := AssertQosMonitoringReportConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.PdvMonReports {
		if err := AssertPdvMonitoringReportConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertRatTypeConstraints(obj.RatType); err != nil {
		return err
	}
	if err := AssertBatOffsetInfoConstraints(obj.BatOffsetInfo); err != nil {
		return err
	}
	for _, el := range obj.RttMonReports {
		if err := AssertQosMonitoringReportConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.QosMonDatRateReps {
		if err := AssertQosMonitoringReport1Constraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AggrDataRateRpts {
		if err := AssertQosMonitoringReport1Constraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.QosMonConInfoReps {
		if err := AssertQosMonitoringReportConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
