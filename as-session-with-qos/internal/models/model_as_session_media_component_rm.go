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

// AsSessionMediaComponentRm - Represents the AsSessionMediaComponent data type with nullable information.
type AsSessionMediaComponentRm struct {

	// Contains the IP data flow(s) description for a single-modal data flow.
	FlowInfos *[]FlowInfo `json:"flowInfos,omitempty"`

	QosReference *string `json:"qosReference,omitempty"`

	AltSerReqs *[]string `json:"altSerReqs,omitempty"`

	// Contains removable alternative service requirements that include individual QoS parameter sets.
	AltSerReqsData *[]AlternativeServiceRequirementsData `json:"altSerReqsData,omitempty"`

	DisUeNotif *bool `json:"disUeNotif,omitempty"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MarBwDl *string `json:"marBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MarBwUl *string `json:"marBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	MedCompN int32 `json:"medCompN"`

	MedType MediaType `json:"medType,omitempty"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MirBwDl *string `json:"mirBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MirBwUl *string `json:"mirBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	TsnQos *TsnQosContainerRm `json:"tsnQos,omitempty"`

	TscaiInputDl *TscaiInputContainer `json:"tscaiInputDl,omitempty"`

	TscaiInputUl *TscaiInputContainer `json:"tscaiInputUl,omitempty"`

	// Round-Trip latency requirement of the service data flow.
	RTLatencyReq bool `json:"rTLatencyReq,omitempty"`

	PduSetQosDl *PduSetQosPara `json:"pduSetQosDl,omitempty"`

	PduSetQosUl *PduSetQosPara `json:"pduSetQosUl,omitempty"`

	L4sInd UplinkDownlinkSupport `json:"l4sInd,omitempty"`

	ProtoDescDl ProtocolDescription `json:"protoDescDl,omitempty"`

	ProtoDescUl ProtocolDescription `json:"protoDescUl,omitempty"`

	// This data type is defined in the same way as the \"DurationMillisec\" data type, but with the OpenAPI nullable property set to true.
	PeriodUl int32 `json:"periodUl,omitempty"`

	// This data type is defined in the same way as the \"DurationMillisec\" data type, but with the OpenAPI nullable property set to true.
	PeriodDl int32 `json:"periodDl,omitempty"`

	EvSubsc *EventsSubscReqDataRm `json:"evSubsc,omitempty"`
}

// AssertAsSessionMediaComponentRmRequired checks if the required fields are not zero-ed
func AssertAsSessionMediaComponentRmRequired(obj AsSessionMediaComponentRm) error {
	elements := map[string]interface{}{
		"medCompN": obj.MedCompN,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.FlowInfos != nil {
		for _, el := range *obj.FlowInfos {
			if err := AssertFlowInfoRequired(el); err != nil {
				return err
			}
		}
	}
	if obj.AltSerReqsData != nil {
		for _, el := range *obj.AltSerReqsData {
			if err := AssertAlternativeServiceRequirementsDataRequired(el); err != nil {
				return err
			}
		}
	}
	if err := AssertMediaTypeRequired(obj.MedType); err != nil {
		return err
	}
	if obj.TsnQos != nil {
		if err := AssertTsnQosContainerRmRequired(*obj.TsnQos); err != nil {
			return err
		}
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
	if obj.PduSetQosDl != nil {
		if err := AssertPduSetQosParaRequired(*obj.PduSetQosDl); err != nil {
			return err
		}
	}
	if obj.PduSetQosUl != nil {
		if err := AssertPduSetQosParaRequired(*obj.PduSetQosUl); err != nil {
			return err
		}
	}
	if err := AssertUplinkDownlinkSupportRequired(obj.L4sInd); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionRequired(obj.ProtoDescDl); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionRequired(obj.ProtoDescUl); err != nil {
		return err
	}
	if obj.EvSubsc != nil {
		if err := AssertEventsSubscReqDataRmRequired(*obj.EvSubsc); err != nil {
			return err
		}
	}
	return nil
}

// AssertAsSessionMediaComponentRmConstraints checks if the values respects the defined constraints
func AssertAsSessionMediaComponentRmConstraints(obj AsSessionMediaComponentRm) error {
	if obj.FlowInfos != nil {
		for _, el := range *obj.FlowInfos {
			if err := AssertFlowInfoConstraints(el); err != nil {
				return err
			}
		}
	}
	if obj.AltSerReqsData != nil {
		for _, el := range *obj.AltSerReqsData {
			if err := AssertAlternativeServiceRequirementsDataConstraints(el); err != nil {
				return err
			}
		}
	}
	if err := AssertMediaTypeConstraints(obj.MedType); err != nil {
		return err
	}
	if obj.TsnQos != nil {
		if err := AssertTsnQosContainerRmConstraints(*obj.TsnQos); err != nil {
			return err
		}
	}
	if obj.TscaiInputDl != nil {
		if err := AssertTscaiInputContainerConstraints(*obj.TscaiInputDl); err != nil {
			return err
		}
	}
	if obj.TscaiInputUl != nil {
		if err := AssertTscaiInputContainerConstraints(*obj.TscaiInputUl); err != nil {
			return err
		}
	}
	if obj.PduSetQosDl != nil {
		if err := AssertPduSetQosParaConstraints(*obj.PduSetQosDl); err != nil {
			return err
		}
	}
	if obj.PduSetQosUl != nil {
		if err := AssertPduSetQosParaConstraints(*obj.PduSetQosUl); err != nil {
			return err
		}
	}
	if err := AssertUplinkDownlinkSupportConstraints(obj.L4sInd); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionConstraints(obj.ProtoDescDl); err != nil {
		return err
	}
	if err := AssertProtocolDescriptionConstraints(obj.ProtoDescUl); err != nil {
		return err
	}
	if obj.EvSubsc != nil {
		if err := AssertEventsSubscReqDataRmConstraints(*obj.EvSubsc); err != nil {
			return err
		}
	}
	return nil
}
