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

// N1MessageClassAnyOf the model 'N1MessageClassAnyOf'
type N1MessageClassAnyOf string

// List of N1MessageClass_anyOf
const (
	_5_GMM N1MessageClassAnyOf = "5GMM"
	SM     N1MessageClassAnyOf = "SM"
	LPP    N1MessageClassAnyOf = "LPP"
	SMS    N1MessageClassAnyOf = "SMS"
	UPDP   N1MessageClassAnyOf = "UPDP"
)

// All allowed values of N1MessageClassAnyOf enum
var AllowedN1MessageClassAnyOfEnumValues = []N1MessageClassAnyOf{
	"5GMM",
	"SM",
	"LPP",
	"SMS",
	"UPDP",
}

func (v *N1MessageClassAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N1MessageClassAnyOf(value)
	for _, existing := range AllowedN1MessageClassAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N1MessageClassAnyOf", value)
}

// NewN1MessageClassAnyOfFromValue returns a pointer to a valid N1MessageClassAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN1MessageClassAnyOfFromValue(v string) (*N1MessageClassAnyOf, error) {
	ev := N1MessageClassAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N1MessageClassAnyOf: valid values are %v", v, AllowedN1MessageClassAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N1MessageClassAnyOf) IsValid() bool {
	for _, existing := range AllowedN1MessageClassAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N1MessageClass_anyOf value
func (v N1MessageClassAnyOf) Ptr() *N1MessageClassAnyOf {
	return &v
}

type NullableN1MessageClassAnyOf struct {
	value *N1MessageClassAnyOf
	isSet bool
}

func (v NullableN1MessageClassAnyOf) Get() *N1MessageClassAnyOf {
	return v.value
}

func (v *NullableN1MessageClassAnyOf) Set(val *N1MessageClassAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN1MessageClassAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN1MessageClassAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1MessageClassAnyOf(val *N1MessageClassAnyOf) *NullableN1MessageClassAnyOf {
	return &NullableN1MessageClassAnyOf{value: val, isSet: true}
}

func (v NullableN1MessageClassAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1MessageClassAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
