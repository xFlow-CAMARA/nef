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

// MediaComponent - Identifies a media component.
type MediaComponent struct {

	// Contains an AF application identifier.
	AfAppId string `json:"afAppId,omitempty"`

	AfRoutReq AfRoutingRequirement `json:"afRoutReq,omitempty"`

	AfSfcReq *AfSfcRequirement `json:"afSfcReq,omitempty"`

	QosReference string `json:"qosReference,omitempty"`

	DisUeNotif bool `json:"disUeNotif,omitempty"`

	AltSerReqs []string `json:"altSerReqs,omitempty"`

	// Contains alternative service requirements that include individual QoS parameter sets.
	AltSerReqsData []AlternativeServiceRequirementsData `json:"altSerReqsData,omitempty"`

	// Represents the content version of some content.
	ContVer int32 `json:"contVer,omitempty"`

	Codecs []string `json:"codecs,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	DesMaxLatency float32 `json:"desMaxLatency,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	DesMaxLoss float32 `json:"desMaxLoss,omitempty"`

	FlusId string `json:"flusId,omitempty"`

	FStatus string `json:"fStatus,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDl string `json:"marBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUl string `json:"marBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxPacketLossRateDl *int32 `json:"maxPacketLossRateDl,omitempty"`

	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxPacketLossRateUl *int32 `json:"maxPacketLossRateUl,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxSuppBwDl string `json:"maxSuppBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxSuppBwUl string `json:"maxSuppBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	MedCompN int32 `json:"medCompN"`

	// Contains the requested bitrate and filters for the set of service data flows identified by their common flow identifier. The key of the map is the fNum attribute.
	MedSubComps map[string]MediaSubComponent `json:"medSubComps,omitempty"`

	MedType MediaType `json:"medType,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwDl string `json:"minDesBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwUl string `json:"minDesBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwDl string `json:"mirBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwUl string `json:"mirBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	PreemptCap PreemptionCapability `json:"preemptCap,omitempty"`

	PreemptVuln PreemptionVulnerability `json:"preemptVuln,omitempty"`

	PrioSharingInd PrioritySharingIndicator `json:"prioSharingInd,omitempty"`

	ResPrio ReservPriority `json:"resPrio,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RrBw string `json:"rrBw,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RsBw string `json:"rsBw,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	SharingKeyDl int32 `json:"sharingKeyDl,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	SharingKeyUl int32 `json:"sharingKeyUl,omitempty"`

	TsnQos TsnQosContainer `json:"tsnQos,omitempty"`

	TscaiInputDl *TscaiInputContainer `json:"tscaiInputDl,omitempty"`

	TscaiInputUl *TscaiInputContainer `json:"tscaiInputUl,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TscaiTimeDom int32 `json:"tscaiTimeDom,omitempty"`

	// Indicates the capability for AF to adjust the burst sending time, when it is supported and set to \"true\". The default value is \"false\" if omitted.
	CapBatAdaptation bool `json:"capBatAdaptation,omitempty"`

	// Indicates the service data flow needs to meet the Round-Trip (RT) latency requirement of the service, when it is included and set to \"true\". The default value is \"false\" if omitted.
	RTLatencyInd bool `json:"rTLatencyInd,omitempty"`

	PduSetQosDl PduSetQosPara `json:"pduSetQosDl,omitempty"`

	PduSetQosUl PduSetQosPara `json:"pduSetQosUl,omitempty"`

	ProtoDescDl ProtocolDescription `json:"protoDescDl,omitempty"`

	ProtoDescUl ProtocolDescription `json:"protoDescUl,omitempty"`

	// Indicates the time interval in units of milliseconds.
	PeriodUl int32 `json:"periodUl,omitempty"`

	// Indicates the time interval in units of milliseconds.
	PeriodDl int32 `json:"periodDl,omitempty"`

	L4sInd UplinkDownlinkSupport `json:"l4sInd,omitempty"`
}

// AssertMediaComponentRequired checks if the required fields are not zero-ed
func AssertMediaComponentRequired(obj MediaComponent) error {
	elements := map[string]interface{}{
		"medCompN": obj.MedCompN,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAfRoutingRequirementRequired(obj.AfRoutReq); err != nil {
		return err
	}
	if obj.AfSfcReq != nil {
		if err := AssertAfSfcRequirementRequired(*obj.AfSfcReq); err != nil {
			return err
		}
	}
	for _, el := range obj.AltSerReqsData {
		if err := AssertAlternativeServiceRequirementsDataRequired(el); err != nil {
			return err
		}
	}
	if err := AssertFlowStatusRequired(obj.FStatus); err != nil {
		return err
	}
	if err := AssertMediaTypeRequired(obj.MedType); err != nil {
		return err
	}
	if err := AssertPreemptionCapabilityRequired(obj.PreemptCap); err != nil {
		return err
	}
	if err := AssertPreemptionVulnerabilityRequired(obj.PreemptVuln); err != nil {
		return err
	}
	if err := AssertPrioritySharingIndicatorRequired(obj.PrioSharingInd); err != nil {
		return err
	}
	if err := AssertReservPriorityRequired(obj.ResPrio); err != nil {
		return err
	}
	if err := AssertTsnQosContainerRequired(obj.TsnQos); err != nil {
		return err
	}
	if obj.TscaiInputDl != nil {
		if err := AssertTscaiInputContainerRequired(*obj.TscaiInputDl); err != nil {
			return err
		}
	}
	if obj.TscaiInputUl != nil {
		if err := AssertTscaiInputContainerRequired(*obj.TscaiInputUl); err != nil {
			return err
		}
	}
	if err := AssertPduSetQosParaRequired(obj.PduSetQosDl); err != nil {
		return err
	}
	if err := AssertPduSetQosParaRequired(obj.PduSetQosUl); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionRequired(obj.ProtoDescDl); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionRequired(obj.ProtoDescUl); err != nil {
		return err
	}
	if err := AssertUplinkDownlinkSupportRequired(obj.L4sInd); err != nil {
		return err
	}
	return nil
}

// AssertMediaComponentConstraints checks if the values respects the defined constraints
func AssertMediaComponentConstraints(obj MediaComponent) error {

	return nil
}
