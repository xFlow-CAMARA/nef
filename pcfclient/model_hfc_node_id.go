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
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.6
*/

package pcfclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HfcNodeId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HfcNodeId{}

// HfcNodeId REpresents the HFC Node Identifer received over NGAP.
type HfcNodeId struct {
	// This IE represents the identifier of the HFC node Id as specified in CableLabs WR-TR-5WWC-ARCH. It is provisioned by the wireline operator as part of wireline operations and may contain up to six characters.
	HfcNId string `json:"hfcNId"`
}

type _HfcNodeId HfcNodeId

// NewHfcNodeId instantiates a new HfcNodeId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHfcNodeId(hfcNId string) *HfcNodeId {
	this := HfcNodeId{}
	this.HfcNId = hfcNId
	return &this
}

// NewHfcNodeIdWithDefaults instantiates a new HfcNodeId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHfcNodeIdWithDefaults() *HfcNodeId {
	this := HfcNodeId{}
	return &this
}

// GetHfcNId returns the HfcNId field value
func (o *HfcNodeId) GetHfcNId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HfcNId
}

// GetHfcNIdOk returns a tuple with the HfcNId field value
// and a boolean to check if the value has been set.
func (o *HfcNodeId) GetHfcNIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HfcNId, true
}

// SetHfcNId sets field value
func (o *HfcNodeId) SetHfcNId(v string) {
	o.HfcNId = v
}

func (o HfcNodeId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HfcNodeId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hfcNId"] = o.HfcNId
	return toSerialize, nil
}

func (o *HfcNodeId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hfcNId",
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

	varHfcNodeId := _HfcNodeId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHfcNodeId)

	if err != nil {
		return err
	}

	*o = HfcNodeId(varHfcNodeId)

	return err
}

type NullableHfcNodeId struct {
	value *HfcNodeId
	isSet bool
}

func (v NullableHfcNodeId) Get() *HfcNodeId {
	return v.value
}

func (v *NullableHfcNodeId) Set(val *HfcNodeId) {
	v.value = val
	v.isSet = true
}

func (v NullableHfcNodeId) IsSet() bool {
	return v.isSet
}

func (v *NullableHfcNodeId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHfcNodeId(val *HfcNodeId) *NullableHfcNodeId {
	return &NullableHfcNodeId{value: val, isSet: true}
}

func (v NullableHfcNodeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHfcNodeId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
