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

// checks if the EasIpReplacementInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasIpReplacementInfo{}

// EasIpReplacementInfo Contains EAS IP replacement information for a Source and a Target EAS.
type EasIpReplacementInfo struct {
	Source EasServerAddress `json:"source"`
	Target EasServerAddress `json:"target"`
}

type _EasIpReplacementInfo EasIpReplacementInfo

// NewEasIpReplacementInfo instantiates a new EasIpReplacementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasIpReplacementInfo(source EasServerAddress, target EasServerAddress) *EasIpReplacementInfo {
	this := EasIpReplacementInfo{}
	this.Source = source
	this.Target = target
	return &this
}

// NewEasIpReplacementInfoWithDefaults instantiates a new EasIpReplacementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasIpReplacementInfoWithDefaults() *EasIpReplacementInfo {
	this := EasIpReplacementInfo{}
	return &this
}

// GetSource returns the Source field value
func (o *EasIpReplacementInfo) GetSource() EasServerAddress {
	if o == nil {
		var ret EasServerAddress
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EasIpReplacementInfo) GetSourceOk() (*EasServerAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EasIpReplacementInfo) SetSource(v EasServerAddress) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *EasIpReplacementInfo) GetTarget() EasServerAddress {
	if o == nil {
		var ret EasServerAddress
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *EasIpReplacementInfo) GetTargetOk() (*EasServerAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *EasIpReplacementInfo) SetTarget(v EasServerAddress) {
	o.Target = v
}

func (o EasIpReplacementInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasIpReplacementInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["target"] = o.Target
	return toSerialize, nil
}

func (o *EasIpReplacementInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"target",
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

	varEasIpReplacementInfo := _EasIpReplacementInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEasIpReplacementInfo)

	if err != nil {
		return err
	}

	*o = EasIpReplacementInfo(varEasIpReplacementInfo)

	return err
}

type NullableEasIpReplacementInfo struct {
	value *EasIpReplacementInfo
	isSet bool
}

func (v NullableEasIpReplacementInfo) Get() *EasIpReplacementInfo {
	return v.value
}

func (v *NullableEasIpReplacementInfo) Set(val *EasIpReplacementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEasIpReplacementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEasIpReplacementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasIpReplacementInfo(val *EasIpReplacementInfo) *NullableEasIpReplacementInfo {
	return &NullableEasIpReplacementInfo{value: val, isSet: true}
}

func (v NullableEasIpReplacementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasIpReplacementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
