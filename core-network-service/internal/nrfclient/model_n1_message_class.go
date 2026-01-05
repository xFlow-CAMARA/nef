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

// N1MessageClass struct for N1MessageClass
type N1MessageClass struct {
	N1MessageClassAnyOf *N1MessageClassAnyOf
	string              *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N1MessageClass) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into N1MessageClassAnyOf
	err = json.Unmarshal(data, &dst.N1MessageClassAnyOf)
	if err == nil {
		jsonN1MessageClassAnyOf, _ := json.Marshal(dst.N1MessageClassAnyOf)
		if string(jsonN1MessageClassAnyOf) == "{}" { // empty struct
			dst.N1MessageClassAnyOf = nil
		} else {
			return nil // data stored in dst.N1MessageClassAnyOf, return on the first match
		}
	} else {
		dst.N1MessageClassAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(N1MessageClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N1MessageClass) MarshalJSON() ([]byte, error) {
	if src.N1MessageClassAnyOf != nil {
		return json.Marshal(&src.N1MessageClassAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN1MessageClass struct {
	value *N1MessageClass
	isSet bool
}

func (v NullableN1MessageClass) Get() *N1MessageClass {
	return v.value
}

func (v *NullableN1MessageClass) Set(val *N1MessageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableN1MessageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableN1MessageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1MessageClass(val *N1MessageClass) *NullableN1MessageClass {
	return &NullableN1MessageClass{value: val, isSet: true}
}

func (v NullableN1MessageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1MessageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
