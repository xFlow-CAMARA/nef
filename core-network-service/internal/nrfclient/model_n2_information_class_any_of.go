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

// N2InformationClassAnyOf the model 'N2InformationClassAnyOf'
type N2InformationClassAnyOf string

// List of N2InformationClass_anyOf
const (
	//	SM N2InformationClassAnyOf = "SM"
	NRPPA    N2InformationClassAnyOf = "NRPPa"
	PWS      N2InformationClassAnyOf = "PWS"
	PWS_BCAL N2InformationClassAnyOf = "PWS-BCAL"
	PWS_RF   N2InformationClassAnyOf = "PWS-RF"
	RAN      N2InformationClassAnyOf = "RAN"
)

// All allowed values of N2InformationClassAnyOf enum
var AllowedN2InformationClassAnyOfEnumValues = []N2InformationClassAnyOf{
	"SM",
	"NRPPa",
	"PWS",
	"PWS-BCAL",
	"PWS-RF",
	"RAN",
}

func (v *N2InformationClassAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N2InformationClassAnyOf(value)
	for _, existing := range AllowedN2InformationClassAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N2InformationClassAnyOf", value)
}

// NewN2InformationClassAnyOfFromValue returns a pointer to a valid N2InformationClassAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN2InformationClassAnyOfFromValue(v string) (*N2InformationClassAnyOf, error) {
	ev := N2InformationClassAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N2InformationClassAnyOf: valid values are %v", v, AllowedN2InformationClassAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N2InformationClassAnyOf) IsValid() bool {
	for _, existing := range AllowedN2InformationClassAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N2InformationClass_anyOf value
func (v N2InformationClassAnyOf) Ptr() *N2InformationClassAnyOf {
	return &v
}

type NullableN2InformationClassAnyOf struct {
	value *N2InformationClassAnyOf
	isSet bool
}

func (v NullableN2InformationClassAnyOf) Get() *N2InformationClassAnyOf {
	return v.value
}

func (v *NullableN2InformationClassAnyOf) Set(val *N2InformationClassAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InformationClassAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InformationClassAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InformationClassAnyOf(val *N2InformationClassAnyOf) *NullableN2InformationClassAnyOf {
	return &NullableN2InformationClassAnyOf{value: val, isSet: true}
}

func (v NullableN2InformationClassAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InformationClassAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
