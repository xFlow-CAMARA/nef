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
	"bytes"
	"encoding/json"
)

// checks if the APIProviderFunctionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIProviderFunctionDetails{}

// APIProviderFunctionDetails Represents an API provider domain function's details.
type APIProviderFunctionDetails struct {
	// API provider domain functionID assigned by the CAPIF core function to the API provider domain function while registering/updating the API provider domain. Shall not be present in the HTTP POST request from the API management function to the CAPIF core function, to register itself. Shall be present in all other HTTP requests and responses.
	ApiProvFuncId   *string                 `json:"apiProvFuncId,omitempty"`
	RegInfo         RegistrationInformation `json:"regInfo"`
	ApiProvFuncRole string                  `json:"apiProvFuncRole"`
	// Generic information related to the API provider domain function such as details of the API provider applications.
	ApiProvFuncInfo *string `json:"apiProvFuncInfo,omitempty"`
}

type _APIProviderFunctionDetails APIProviderFunctionDetails

// NewAPIProviderFunctionDetails instantiates a new APIProviderFunctionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIProviderFunctionDetails(regInfo RegistrationInformation, apiProvFuncRole string) *APIProviderFunctionDetails {
	this := APIProviderFunctionDetails{}
	this.RegInfo = regInfo
	this.ApiProvFuncRole = apiProvFuncRole
	return &this
}

// NewAPIProviderFunctionDetailsWithDefaults instantiates a new APIProviderFunctionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIProviderFunctionDetailsWithDefaults() *APIProviderFunctionDetails {
	this := APIProviderFunctionDetails{}
	return &this
}

// GetApiProvFuncId returns the ApiProvFuncId field value if set, zero value otherwise.
func (o *APIProviderFunctionDetails) GetApiProvFuncId() string {
	if o == nil || IsNil(o.ApiProvFuncId) {
		var ret string
		return ret
	}
	return *o.ApiProvFuncId
}

// GetApiProvFuncIdOk returns a tuple with the ApiProvFuncId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderFunctionDetails) GetApiProvFuncIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiProvFuncId) {
		return nil, false
	}
	return o.ApiProvFuncId, true
}

// HasApiProvFuncId returns a boolean if a field has been set.
func (o *APIProviderFunctionDetails) HasApiProvFuncId() bool {
	if o != nil && !IsNil(o.ApiProvFuncId) {
		return true
	}

	return false
}

// SetApiProvFuncId gets a reference to the given string and assigns it to the ApiProvFuncId field.
func (o *APIProviderFunctionDetails) SetApiProvFuncId(v string) {
	o.ApiProvFuncId = &v
}

// GetRegInfo returns the RegInfo field value
func (o *APIProviderFunctionDetails) GetRegInfo() RegistrationInformation {
	if o == nil {
		var ret RegistrationInformation
		return ret
	}

	return o.RegInfo
}

// GetRegInfoOk returns a tuple with the RegInfo field value
// and a boolean to check if the value has been set.
func (o *APIProviderFunctionDetails) GetRegInfoOk() (*RegistrationInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegInfo, true
}

// SetRegInfo sets field value
func (o *APIProviderFunctionDetails) SetRegInfo(v RegistrationInformation) {
	o.RegInfo = v
}

// GetApiProvFuncRole returns the ApiProvFuncRole field value
func (o *APIProviderFunctionDetails) GetApiProvFuncRole() string {
	if o == nil {
		return ""
	}

	return o.ApiProvFuncRole
}

// GetApiProvFuncRoleOk returns a tuple with the ApiProvFuncRole field value
// and a boolean to check if the value has been set.
func (o *APIProviderFunctionDetails) GetApiProvFuncRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiProvFuncRole, true
}

// SetApiProvFuncRole sets field value
func (o *APIProviderFunctionDetails) SetApiProvFuncRole(v string) {
	o.ApiProvFuncRole = v
}

// GetApiProvFuncInfo returns the ApiProvFuncInfo field value if set, zero value otherwise.
func (o *APIProviderFunctionDetails) GetApiProvFuncInfo() string {
	if o == nil || IsNil(o.ApiProvFuncInfo) {
		var ret string
		return ret
	}
	return *o.ApiProvFuncInfo
}

// GetApiProvFuncInfoOk returns a tuple with the ApiProvFuncInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderFunctionDetails) GetApiProvFuncInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ApiProvFuncInfo) {
		return nil, false
	}
	return o.ApiProvFuncInfo, true
}

// HasApiProvFuncInfo returns a boolean if a field has been set.
func (o *APIProviderFunctionDetails) HasApiProvFuncInfo() bool {
	if o != nil && !IsNil(o.ApiProvFuncInfo) {
		return true
	}

	return false
}

// SetApiProvFuncInfo gets a reference to the given string and assigns it to the ApiProvFuncInfo field.
func (o *APIProviderFunctionDetails) SetApiProvFuncInfo(v string) {
	o.ApiProvFuncInfo = &v
}

func (o APIProviderFunctionDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIProviderFunctionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiProvFuncId) {
		toSerialize["apiProvFuncId"] = o.ApiProvFuncId
	}
	toSerialize["regInfo"] = o.RegInfo
	toSerialize["apiProvFuncRole"] = o.ApiProvFuncRole
	if !IsNil(o.ApiProvFuncInfo) {
		toSerialize["apiProvFuncInfo"] = o.ApiProvFuncInfo
	}
	return toSerialize, nil
}

func (o *APIProviderFunctionDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"regInfo",
		"apiProvFuncRole",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return nil
		}
	}

	varAPIProviderFunctionDetails := _APIProviderFunctionDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAPIProviderFunctionDetails)

	if err != nil {
		return err
	}

	*o = APIProviderFunctionDetails(varAPIProviderFunctionDetails)

	return err
}

type NullableAPIProviderFunctionDetails struct {
	value *APIProviderFunctionDetails
	isSet bool
}

func (v NullableAPIProviderFunctionDetails) Get() *APIProviderFunctionDetails {
	return v.value
}

func (v *NullableAPIProviderFunctionDetails) Set(val *APIProviderFunctionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIProviderFunctionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIProviderFunctionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIProviderFunctionDetails(val *APIProviderFunctionDetails) *NullableAPIProviderFunctionDetails {
	return &NullableAPIProviderFunctionDetails{value: val, isSet: true}
}

func (v NullableAPIProviderFunctionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIProviderFunctionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
