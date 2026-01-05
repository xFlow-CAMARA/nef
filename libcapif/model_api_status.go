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
CAPIF_Publish_Service_API

API for publishing service APIs.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.5
*/

package libcapif

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ApiStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStatus{}

// ApiStatus Represents the API status.
type ApiStatus struct {
	// Indicates the list of AEF ID(s) where the API is active. If this attribute is omitted, the API is inactive at all AEF(s) defined in the \"aefProfiles\" attribute within the ServiceAPIDescription data structure.
	AefIds []string `json:"aefIds"`
}

type _ApiStatus ApiStatus

// NewApiStatus instantiates a new ApiStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStatus(aefIds []string) *ApiStatus {
	this := ApiStatus{}
	this.AefIds = aefIds
	return &this
}

// NewApiStatusWithDefaults instantiates a new ApiStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStatusWithDefaults() *ApiStatus {
	this := ApiStatus{}
	return &this
}

// GetAefIds returns the AefIds field value
func (o *ApiStatus) GetAefIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AefIds
}

// GetAefIdsOk returns a tuple with the AefIds field value
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetAefIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AefIds, true
}

// SetAefIds sets field value
func (o *ApiStatus) SetAefIds(v []string) {
	o.AefIds = v
}

func (o ApiStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aefIds"] = o.AefIds
	return toSerialize, nil
}

func (o *ApiStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aefIds",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiStatus := _ApiStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiStatus)

	if err != nil {
		return err
	}

	*o = ApiStatus(varApiStatus)

	return err
}

type NullableApiStatus struct {
	value *ApiStatus
	isSet bool
}

func (v NullableApiStatus) Get() *ApiStatus {
	return v.value
}

func (v *NullableApiStatus) Set(val *ApiStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStatus(val *ApiStatus) *NullableApiStatus {
	return &NullableApiStatus{value: val, isSet: true}
}

func (v NullableApiStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
