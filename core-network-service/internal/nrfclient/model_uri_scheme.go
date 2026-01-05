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

// UriScheme struct for UriScheme
type UriScheme struct {
	UriSchemeAnyOf *UriSchemeAnyOf
	string         *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UriScheme) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UriSchemeAnyOf
	err = json.Unmarshal(data, &dst.UriSchemeAnyOf)
	if err == nil {
		jsonUriSchemeAnyOf, _ := json.Marshal(dst.UriSchemeAnyOf)
		if string(jsonUriSchemeAnyOf) == "{}" { // empty struct
			dst.UriSchemeAnyOf = nil
		} else {
			return nil // data stored in dst.UriSchemeAnyOf, return on the first match
		}
	} else {
		dst.UriSchemeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UriScheme)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UriScheme) MarshalJSON() ([]byte, error) {
	if src.UriSchemeAnyOf != nil {
		return json.Marshal(&src.UriSchemeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUriScheme struct {
	value *UriScheme
	isSet bool
}

func (v NullableUriScheme) Get() *UriScheme {
	return v.value
}

func (v *NullableUriScheme) Set(val *UriScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableUriScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableUriScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUriScheme(val *UriScheme) *NullableUriScheme {
	return &NullableUriScheme{value: val, isSet: true}
}

func (v NullableUriScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUriScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
