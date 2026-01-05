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

// LocationInfo - Represents the user location information.
type LocationInfo struct {

	// Unsigned integer identifying a period of time in units of minutes.
	AgeOfLocationInfo int32 `json:"ageOfLocationInfo,omitempty"`

	// Indicates the Cell Global Identification of the user which identifies the cell the UE is registered.
	CellId string `json:"cellId,omitempty"`

	// Indicates the eNodeB in which the UE is currently located.
	EnodeBId *string `json:"enodeBId,omitempty"`

	// Identifies the Routing Area Identity of the user where the UE is located.
	RoutingAreaId *string `json:"routingAreaId,omitempty"`

	// Identifies the Tracking Area Identity of the user where the UE is located.
	TrackingAreaId *string `json:"trackingAreaId,omitempty"`

	// Identifies the PLMN Identity of the user where the UE is located.
	PlmnId *string `json:"plmnId,omitempty"`

	// Identifies the TWAN Identity of the user where the UE is located.
	TwanId *string `json:"twanId,omitempty"`

	UserLocation *UserLocation `json:"userLocation,omitempty"`

	GeographicArea *GeographicArea `json:"geographicArea,omitempty"`

	CivicAddress *CivicAddress `json:"civicAddress,omitempty"`

	PositionMethod PositioningMethod `json:"positionMethod,omitempty"`

	QosFulfilInd *AccuracyFulfilmentIndicator `json:"qosFulfilInd,omitempty"`

	UeVelocity *VelocityEstimate `json:"ueVelocity,omitempty"`

	LdrType *LdrType `json:"ldrType,omitempty"`

	AchievedQos *MinorLocationQoS `json:"achievedQos,omitempty"`

	// String identifying an UE with application layer ID. The format of the application  layer ID parameter is same as the Application layer ID defined in clause 11.3.4 of  3GPP TS 24.554.
	RelAppLayerId *string `json:"relAppLayerId,omitempty"`

	RangeDirection *RangeDirection `json:"rangeDirection,omitempty"`

	TwoDRelLoc *Model2DRelativeLocation `json:"twoDRelLoc,omitempty"`

	ThreeDRelLoc *Model3DRelativeLocation `json:"threeDRelLoc,omitempty"`

	RelVelocity *VelocityEstimate `json:"relVelocity,omitempty"`

	UpCumEvtRep *UpCumEvtRep `json:"upCumEvtRep,omitempty"`
}

// AssertLocationInfoRequired checks if the required fields are not zero-ed
func AssertLocationInfoRequired(obj LocationInfo) error {

	return nil
}

// AssertLocationInfoConstraints checks if the values respects the defined constraints
func AssertLocationInfoConstraints(obj LocationInfo) error {

	return nil
}
