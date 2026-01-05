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
)

// TrafficInfluSubPatch - Represents parameters to request the modification of a traffic influence subscription resource.
type TrafficInfluSubPatch struct {

	// Identifies whether an application can be relocated once a location of the application has been selected.
	AppReloInd *bool `json:"appReloInd,omitempty"`

	// Identifies IP packet filters.
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`

	// Identifies Ethernet packet filters.
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty"`

	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []RouteToLocation `json:"trafficRoutes,omitempty"`

	// Reference to a pre-configured steering of user traffic to service function chain in downlink.
	SfcIdDl *string `json:"sfcIdDl,omitempty"`

	// Reference to a pre-configured steering of user traffic to service function chain in uplink.
	SfcIdUl *string `json:"sfcIdUl,omitempty"`

	// A String which is transparently passed to the UPF to be applied for traffic to SFC.
	Metadata *string `json:"metadata,omitempty"`

	TfcCorrInd *bool `json:"tfcCorrInd,omitempty"`

	TempValidities *[]TemporalValidity `json:"tempValidities,omitempty"`

	// Identifies a geographic zone that the AF request applies only to the traffic of UE(s) located in this specific zone.
	// Deprecated
	ValidGeoZoneIds *[]string `json:"validGeoZoneIds,omitempty"`

	// Identifies geographical areas within which the AF request applies.
	GeoAreas *[]GeographicalArea `json:"geoAreas,omitempty"`

	AfAckInd *bool `json:"afAckInd,omitempty"`

	AddrPreserInd *bool `json:"addrPreserInd,omitempty"`

	// Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.
	SimConnInd bool `json:"simConnInd,omitempty"`

	// indicating a time in seconds.
	SimConnTerm int32 `json:"simConnTerm,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	MaxAllowedUpLat *int32 `json:"maxAllowedUpLat,omitempty"`

	// Contains EAS IP replacement information.
	EasIpReplaceInfos *[]EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`

	// Indicates the EAS rediscovery is required for the application if it is included and set to \"true\".
	EasRedisInd bool `json:"easRedisInd,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty"`

	EventReq ReportingInformation `json:"eventReq,omitempty"`

	TfcCorreInfo *TrafficCorrelationInfo `json:"tfcCorreInfo,omitempty"`
}

// AssertTrafficInfluSubPatchRequired checks if the required fields are not zero-ed
func AssertTrafficInfluSubPatchRequired(obj TrafficInfluSubPatch) error {
	for _, el := range obj.TrafficFilters {
		if err := AssertFlowInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EthTrafficFilters {
		if err := AssertEthFlowDescriptionRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TrafficRoutes {
		if err := AssertRouteToLocationRequired(el); err != nil {
			return err
		}
	}
	if obj.TempValidities != nil {
		for _, el := range *obj.TempValidities {
			if err := AssertTemporalValidityRequired(el); err != nil {
				return err
			}
		}
	}
	if obj.GeoAreas != nil {
		for _, el := range *obj.GeoAreas {
			if err := AssertGeographicalAreaRequired(el); err != nil {
				return err
			}
		}
	}
	if obj.EasIpReplaceInfos != nil {
		for _, el := range *obj.EasIpReplaceInfos {
			if err := AssertEasIpReplacementInfoRequired(el); err != nil {
				return err
			}
		}
	}
	if err := AssertReportingInformationRequired(obj.EventReq); err != nil {
		return err
	}
	if obj.TfcCorreInfo != nil {
		if err := AssertTrafficCorrelationInfoRequired(*obj.TfcCorreInfo); err != nil {
			return err
		}
	}
	return nil
}

// AssertTrafficInfluSubPatchConstraints checks if the values respects the defined constraints
func AssertTrafficInfluSubPatchConstraints(obj TrafficInfluSubPatch) error {
	for _, el := range obj.TrafficFilters {
		if err := AssertFlowInfoConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EthTrafficFilters {
		if err := AssertEthFlowDescriptionConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TrafficRoutes {
		if err := AssertRouteToLocationConstraints(el); err != nil {
			return err
		}
	}
	if obj.TempValidities != nil {
		for _, el := range *obj.TempValidities {
			if err := AssertTemporalValidityConstraints(el); err != nil {
				return err
			}
		}
	}
	if obj.GeoAreas != nil {
		for _, el := range *obj.GeoAreas {
			if err := AssertGeographicalAreaConstraints(el); err != nil {
				return err
			}
		}
	}
	if obj.MaxAllowedUpLat != nil && *obj.MaxAllowedUpLat < 0 {
		return &ParsingError{Param: "MaxAllowedUpLat", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.EasIpReplaceInfos != nil {
		for _, el := range *obj.EasIpReplaceInfos {
			if err := AssertEasIpReplacementInfoConstraints(el); err != nil {
				return err
			}
		}
	}
	if err := AssertReportingInformationConstraints(obj.EventReq); err != nil {
		return err
	}
	if obj.TfcCorreInfo != nil {
		if err := AssertTrafficCorrelationInfoConstraints(*obj.TfcCorreInfo); err != nil {
			return err
		}
	}
	return nil
}
