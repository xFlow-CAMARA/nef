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

// NFServiceStatus struct for NFServiceStatus
type NFServiceStatus struct {
	NFStatusAnyOf *NFStatusAnyOf
	string        *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NFServiceStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NFStatusAnyOf
	err = json.Unmarshal(data, &dst.NFStatusAnyOf)
	if err == nil {
		jsonNFStatusAnyOf, _ := json.Marshal(dst.NFStatusAnyOf)
		if string(jsonNFStatusAnyOf) == "{}" { // empty struct
			dst.NFStatusAnyOf = nil
		} else {
			return nil // data stored in dst.NFStatusAnyOf, return on the first match
		}
	} else {
		dst.NFStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NFServiceStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NFServiceStatus) MarshalJSON() ([]byte, error) {
	if src.NFStatusAnyOf != nil {
		return json.Marshal(&src.NFStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNFServiceStatus struct {
	value *NFServiceStatus
	isSet bool
}

func (v NullableNFServiceStatus) Get() *NFServiceStatus {
	return v.value
}

func (v *NullableNFServiceStatus) Set(val *NFServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNFServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNFServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFServiceStatus(val *NFServiceStatus) *NullableNFServiceStatus {
	return &NullableNFServiceStatus{value: val, isSet: true}
}

func (v NullableNFServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
