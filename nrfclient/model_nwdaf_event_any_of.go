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
NRF NFDiscovery Service

NRF NFDiscovery Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0.alpha-1
*/

package nrfclient

import (
	"encoding/json"
	"fmt"
)

// NwdafEventAnyOf the model 'NwdafEventAnyOf'
type NwdafEventAnyOf string

// List of NwdafEvent_anyOf
const (
	SLICE_LOAD_LEVEL NwdafEventAnyOf = "SLICE_LOAD_LEVEL"

// NETWORK_PERFORMANCE NwdafEventAnyOf = "NETWORK_PERFORMANCE"
// NF_LOAD NwdafEventAnyOf = "NF_LOAD"
// SERVICE_EXPERIENCE NwdafEventAnyOf = "SERVICE_EXPERIENCE"
// UE_MOBILITY NwdafEventAnyOf = "UE_MOBILITY"
// UE_COMMUNICATION NwdafEventAnyOf = "UE_COMMUNICATION"
// QOS_SUSTAINABILITY NwdafEventAnyOf = "QOS_SUSTAINABILITY"
// ABNORMAL_BEHAVIOUR NwdafEventAnyOf = "ABNORMAL_BEHAVIOUR"
// USER_DATA_CONGESTION NwdafEventAnyOf = "USER_DATA_CONGESTION"
// NSI_LOAD_LEVEL NwdafEventAnyOf = "NSI_LOAD_LEVEL"
)

// All allowed values of NwdafEventAnyOf enum
var AllowedNwdafEventAnyOfEnumValues = []NwdafEventAnyOf{
	"SLICE_LOAD_LEVEL",
	"NETWORK_PERFORMANCE",
	"NF_LOAD",
	"SERVICE_EXPERIENCE",
	"UE_MOBILITY",
	"UE_COMMUNICATION",
	"QOS_SUSTAINABILITY",
	"ABNORMAL_BEHAVIOUR",
	"USER_DATA_CONGESTION",
	"NSI_LOAD_LEVEL",
}

func (v *NwdafEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NwdafEventAnyOf(value)
	for _, existing := range AllowedNwdafEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NwdafEventAnyOf", value)
}

// NewNwdafEventAnyOfFromValue returns a pointer to a valid NwdafEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNwdafEventAnyOfFromValue(v string) (*NwdafEventAnyOf, error) {
	ev := NwdafEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NwdafEventAnyOf: valid values are %v", v, AllowedNwdafEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NwdafEventAnyOf) IsValid() bool {
	for _, existing := range AllowedNwdafEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NwdafEvent_anyOf value
func (v NwdafEventAnyOf) Ptr() *NwdafEventAnyOf {
	return &v
}

type NullableNwdafEventAnyOf struct {
	value *NwdafEventAnyOf
	isSet bool
}

func (v NullableNwdafEventAnyOf) Get() *NwdafEventAnyOf {
	return v.value
}

func (v *NullableNwdafEventAnyOf) Set(val *NwdafEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafEventAnyOf(val *NwdafEventAnyOf) *NullableNwdafEventAnyOf {
	return &NullableNwdafEventAnyOf{value: val, isSet: true}
}

func (v NullableNwdafEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
