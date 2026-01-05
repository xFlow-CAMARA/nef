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

// NFTypeAnyOf the model 'NFTypeAnyOf'
type NFTypeAnyOf string

// List of NFType_anyOf
const (
	NRF      NFTypeAnyOf = "NRF"
	UDM      NFTypeAnyOf = "UDM"
	AMF      NFTypeAnyOf = "AMF"
	SMF      NFTypeAnyOf = "SMF"
	AUSF     NFTypeAnyOf = "AUSF"
	NEF      NFTypeAnyOf = "NEF"
	PCF      NFTypeAnyOf = "PCF"
	SMSF     NFTypeAnyOf = "SMSF"
	NSSF     NFTypeAnyOf = "NSSF"
	UDR      NFTypeAnyOf = "UDR"
	LMF      NFTypeAnyOf = "LMF"
	GMLC     NFTypeAnyOf = "GMLC"
	_5_G_EIR NFTypeAnyOf = "5G_EIR"
	SEPP     NFTypeAnyOf = "SEPP"
	UPF      NFTypeAnyOf = "UPF"
	N3_IWF   NFTypeAnyOf = "N3IWF"
	AF       NFTypeAnyOf = "AF"
	UDSF     NFTypeAnyOf = "UDSF"
	BSF      NFTypeAnyOf = "BSF"
	CHF      NFTypeAnyOf = "CHF"
	NWDAF    NFTypeAnyOf = "NWDAF"
)

// All allowed values of NFTypeAnyOf enum
var AllowedNFTypeAnyOfEnumValues = []NFTypeAnyOf{
	"NRF",
	"UDM",
	"AMF",
	"SMF",
	"AUSF",
	"NEF",
	"PCF",
	"SMSF",
	"NSSF",
	"UDR",
	"LMF",
	"GMLC",
	"5G_EIR",
	"SEPP",
	"UPF",
	"N3IWF",
	"AF",
	"UDSF",
	"BSF",
	"CHF",
	"NWDAF",
}

func (v *NFTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NFTypeAnyOf(value)
	for _, existing := range AllowedNFTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NFTypeAnyOf", value)
}

// NewNFTypeAnyOfFromValue returns a pointer to a valid NFTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNFTypeAnyOfFromValue(v string) (*NFTypeAnyOf, error) {
	ev := NFTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NFTypeAnyOf: valid values are %v", v, AllowedNFTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NFTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNFTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NFType_anyOf value
func (v NFTypeAnyOf) Ptr() *NFTypeAnyOf {
	return &v
}

type NullableNFTypeAnyOf struct {
	value *NFTypeAnyOf
	isSet bool
}

func (v NullableNFTypeAnyOf) Get() *NFTypeAnyOf {
	return v.value
}

func (v *NullableNFTypeAnyOf) Set(val *NFTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTypeAnyOf(val *NFTypeAnyOf) *NullableNFTypeAnyOf {
	return &NullableNFTypeAnyOf{value: val, isSet: true}
}

func (v NullableNFTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
