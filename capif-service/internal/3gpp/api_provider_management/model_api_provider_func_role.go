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
CAPIF_API_Provider_Management_API

API for API provider domain functions management.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.3
*/

package api_provider_management

import (
	"encoding/json"
	"fmt"
)

// ApiProviderFuncRole Indicates the role (e.g. AEF, APF, etc.) of an API provider domain function.   Possible values are: - AEF: API provider function is API Exposing Function. - APF: API provider function is API Publishing Function. - AMF: API Provider function is API Management Function.
type ApiProviderFuncRole string

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ApiProviderFuncRole) UnmarshalJSON(data []byte) error {

	// try to unmarshal JSON data into String
	err := json.Unmarshal(data, &dst)
	if err == nil {
		jsonString, _ := json.Marshal(dst)
		if string(jsonString) != "{}" { // empty struct
			return nil // data stored in dst.String, return on the first match
		}
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ApiProviderFuncRole)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ApiProviderFuncRole) MarshalJSON() ([]byte, error) {
	if src != nil {
		return json.Marshal(&src)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableApiProviderFuncRole struct {
	value *ApiProviderFuncRole
	isSet bool
}

func (v NullableApiProviderFuncRole) Get() *ApiProviderFuncRole {
	return v.value
}

func (v *NullableApiProviderFuncRole) Set(val *ApiProviderFuncRole) {
	v.value = val
	v.isSet = true
}

func (v NullableApiProviderFuncRole) IsSet() bool {
	return v.isSet
}

func (v *NullableApiProviderFuncRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiProviderFuncRole(val *ApiProviderFuncRole) *NullableApiProviderFuncRole {
	return &NullableApiProviderFuncRole{value: val, isSet: true}
}

func (v NullableApiProviderFuncRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiProviderFuncRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
