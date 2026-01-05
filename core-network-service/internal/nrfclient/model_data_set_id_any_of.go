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

// DataSetIdAnyOf the model 'DataSetIdAnyOf'
type DataSetIdAnyOf string

// List of DataSetId_anyOf
const (
	SUBSCRIPTION DataSetIdAnyOf = "SUBSCRIPTION"
	POLICY       DataSetIdAnyOf = "POLICY"
	EXPOSURE     DataSetIdAnyOf = "EXPOSURE"
	APPLICATION  DataSetIdAnyOf = "APPLICATION"
)

// All allowed values of DataSetIdAnyOf enum
var AllowedDataSetIdAnyOfEnumValues = []DataSetIdAnyOf{
	"SUBSCRIPTION",
	"POLICY",
	"EXPOSURE",
	"APPLICATION",
}

func (v *DataSetIdAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSetIdAnyOf(value)
	for _, existing := range AllowedDataSetIdAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSetIdAnyOf", value)
}

// NewDataSetIdAnyOfFromValue returns a pointer to a valid DataSetIdAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSetIdAnyOfFromValue(v string) (*DataSetIdAnyOf, error) {
	ev := DataSetIdAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSetIdAnyOf: valid values are %v", v, AllowedDataSetIdAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSetIdAnyOf) IsValid() bool {
	for _, existing := range AllowedDataSetIdAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSetId_anyOf value
func (v DataSetIdAnyOf) Ptr() *DataSetIdAnyOf {
	return &v
}

type NullableDataSetIdAnyOf struct {
	value *DataSetIdAnyOf
	isSet bool
}

func (v NullableDataSetIdAnyOf) Get() *DataSetIdAnyOf {
	return v.value
}

func (v *NullableDataSetIdAnyOf) Set(val *DataSetIdAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSetIdAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSetIdAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSetIdAnyOf(val *DataSetIdAnyOf) *NullableDataSetIdAnyOf {
	return &NullableDataSetIdAnyOf{value: val, isSet: true}
}

func (v NullableDataSetIdAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSetIdAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
