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

// NFStatus struct for NFStatus
type NFStatus struct {
	NFStatusAnyOf *NFStatusAnyOf
	string        *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NFStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NFStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NFStatus) MarshalJSON() ([]byte, error) {
	if src.NFStatusAnyOf != nil {
		return json.Marshal(&src.NFStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNFStatus struct {
	value *NFStatus
	isSet bool
}

func (v NullableNFStatus) Get() *NFStatus {
	return v.value
}

func (v *NullableNFStatus) Set(val *NFStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNFStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNFStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFStatus(val *NFStatus) *NullableNFStatus {
	return &NullableNFStatus{value: val, isSet: true}
}

func (v NullableNFStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
