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

package nrfclient

import (
	"encoding/json"
	"fmt"
)

// NFStatusAnyOf the model 'NFStatusAnyOf'
type NFStatusAnyOf string

// List of NFStatus_anyOf
const (
	REGISTERED     NFStatusAnyOf = "REGISTERED"
	SUSPENDED      NFStatusAnyOf = "SUSPENDED"
	UNDISCOVERABLE NFStatusAnyOf = "UNDISCOVERABLE"
)

// All allowed values of NFStatusAnyOf enum
var AllowedNFStatusAnyOfEnumValues = []NFStatusAnyOf{
	"REGISTERED",
	"SUSPENDED",
	"UNDISCOVERABLE",
}

func (v *NFStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NFStatusAnyOf(value)
	for _, existing := range AllowedNFStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NFStatusAnyOf", value)
}

// NewNFStatusAnyOfFromValue returns a pointer to a valid NFStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNFStatusAnyOfFromValue(v string) (*NFStatusAnyOf, error) {
	ev := NFStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NFStatusAnyOf: valid values are %v", v, AllowedNFStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NFStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedNFStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NFStatus_anyOf value
func (v NFStatusAnyOf) Ptr() *NFStatusAnyOf {
	return &v
}

type NullableNFStatusAnyOf struct {
	value *NFStatusAnyOf
	isSet bool
}

func (v NullableNFStatusAnyOf) Get() *NFStatusAnyOf {
	return v.value
}

func (v *NullableNFStatusAnyOf) Set(val *NFStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNFStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNFStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFStatusAnyOf(val *NFStatusAnyOf) *NullableNFStatusAnyOf {
	return &NullableNFStatusAnyOf{value: val, isSet: true}
}

func (v NullableNFStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
