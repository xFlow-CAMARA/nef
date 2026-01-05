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

// UPInterfaceType struct for UPInterfaceType
type UPInterfaceType struct {
	UPInterfaceTypeAnyOf *UPInterfaceTypeAnyOf
	string               *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UPInterfaceType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UPInterfaceTypeAnyOf
	err = json.Unmarshal(data, &dst.UPInterfaceTypeAnyOf)
	if err == nil {
		jsonUPInterfaceTypeAnyOf, _ := json.Marshal(dst.UPInterfaceTypeAnyOf)
		if string(jsonUPInterfaceTypeAnyOf) == "{}" { // empty struct
			dst.UPInterfaceTypeAnyOf = nil
		} else {
			return nil // data stored in dst.UPInterfaceTypeAnyOf, return on the first match
		}
	} else {
		dst.UPInterfaceTypeAnyOf = nil
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

	return fmt.Errorf("Data failed to match schemas in anyOf(UPInterfaceType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UPInterfaceType) MarshalJSON() ([]byte, error) {
	if src.UPInterfaceTypeAnyOf != nil {
		return json.Marshal(&src.UPInterfaceTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUPInterfaceType struct {
	value *UPInterfaceType
	isSet bool
}

func (v NullableUPInterfaceType) Get() *UPInterfaceType {
	return v.value
}

func (v *NullableUPInterfaceType) Set(val *UPInterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableUPInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableUPInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUPInterfaceType(val *UPInterfaceType) *NullableUPInterfaceType {
	return &NullableUPInterfaceType{value: val, isSet: true}
}

func (v NullableUPInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUPInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
