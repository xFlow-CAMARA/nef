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

// UPInterfaceTypeAnyOf the model 'UPInterfaceTypeAnyOf'
type UPInterfaceTypeAnyOf string

// List of UPInterfaceType_anyOf
const (
	N3 UPInterfaceTypeAnyOf = "N3"
	N6 UPInterfaceTypeAnyOf = "N6"
	N9 UPInterfaceTypeAnyOf = "N9"
)

// All allowed values of UPInterfaceTypeAnyOf enum
var AllowedUPInterfaceTypeAnyOfEnumValues = []UPInterfaceTypeAnyOf{
	"N3",
	"N6",
	"N9",
}

func (v *UPInterfaceTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UPInterfaceTypeAnyOf(value)
	for _, existing := range AllowedUPInterfaceTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UPInterfaceTypeAnyOf", value)
}

// NewUPInterfaceTypeAnyOfFromValue returns a pointer to a valid UPInterfaceTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUPInterfaceTypeAnyOfFromValue(v string) (*UPInterfaceTypeAnyOf, error) {
	ev := UPInterfaceTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UPInterfaceTypeAnyOf: valid values are %v", v, AllowedUPInterfaceTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UPInterfaceTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedUPInterfaceTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UPInterfaceType_anyOf value
func (v UPInterfaceTypeAnyOf) Ptr() *UPInterfaceTypeAnyOf {
	return &v
}

type NullableUPInterfaceTypeAnyOf struct {
	value *UPInterfaceTypeAnyOf
	isSet bool
}

func (v NullableUPInterfaceTypeAnyOf) Get() *UPInterfaceTypeAnyOf {
	return v.value
}

func (v *NullableUPInterfaceTypeAnyOf) Set(val *UPInterfaceTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUPInterfaceTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUPInterfaceTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUPInterfaceTypeAnyOf(val *UPInterfaceTypeAnyOf) *NullableUPInterfaceTypeAnyOf {
	return &NullableUPInterfaceTypeAnyOf{value: val, isSet: true}
}

func (v NullableUPInterfaceTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUPInterfaceTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
