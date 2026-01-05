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

// checks if the InvalidParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidParam{}

// InvalidParam Represents the description of invalid parameters, for a request rejected due to invalid parameters.
type InvalidParam struct {
	// Attribute's name encoded as a JSON Pointer, or header's name.
	Param string `json:"param"`
	// A human-readable reason, e.g. \"must be a positive integer\".
	Reason *string `json:"reason,omitempty"`
}

type _InvalidParam InvalidParam

// NewInvalidParam instantiates a new InvalidParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidParam(param string) *InvalidParam {
	this := InvalidParam{}
	this.Param = param
	return &this
}

// NewInvalidParamWithDefaults instantiates a new InvalidParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidParamWithDefaults() *InvalidParam {
	this := InvalidParam{}
	return &this
}

// GetParam returns the Param field value
func (o *InvalidParam) GetParam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Param
}

// GetParamOk returns a tuple with the Param field value
// and a boolean to check if the value has been set.
func (o *InvalidParam) GetParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Param, true
}

// SetParam sets field value
func (o *InvalidParam) SetParam(v string) {
	o.Param = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InvalidParam) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidParam) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InvalidParam) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InvalidParam) SetReason(v string) {
	o.Reason = &v
}

func (o InvalidParam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvalidParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["param"] = o.Param
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

func (o *InvalidParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"param",
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

	varInvalidParam := _InvalidParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvalidParam)

	if err != nil {
		return err
	}

	*o = InvalidParam(varInvalidParam)

	return err
}

type NullableInvalidParam struct {
	value *InvalidParam
	isSet bool
}

func (v NullableInvalidParam) Get() *InvalidParam {
	return v.value
}

func (v *NullableInvalidParam) Set(val *InvalidParam) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidParam) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidParam(val *InvalidParam) *NullableInvalidParam {
	return &NullableInvalidParam{value: val, isSet: true}
}

func (v NullableInvalidParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
