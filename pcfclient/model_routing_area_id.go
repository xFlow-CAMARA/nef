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

// checks if the RoutingAreaId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingAreaId{}

// RoutingAreaId Contains a Routing Area Identification as defined in 3GPP TS 23.003, clause 4.2.
type RoutingAreaId struct {
	PlmnId PlmnId `json:"plmnId"`
	// Location Area Code
	Lac string `json:"lac" validate:"regexp=^[A-Fa-f0-9]{4}$"`
	// Routing Area Code
	Rac string `json:"rac" validate:"regexp=^[A-Fa-f0-9]{2}$"`
}

type _RoutingAreaId RoutingAreaId

// NewRoutingAreaId instantiates a new RoutingAreaId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingAreaId(plmnId PlmnId, lac string, rac string) *RoutingAreaId {
	this := RoutingAreaId{}
	this.PlmnId = plmnId
	this.Lac = lac
	this.Rac = rac
	return &this
}

// NewRoutingAreaIdWithDefaults instantiates a new RoutingAreaId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingAreaIdWithDefaults() *RoutingAreaId {
	this := RoutingAreaId{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *RoutingAreaId) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *RoutingAreaId) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *RoutingAreaId) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetLac returns the Lac field value
func (o *RoutingAreaId) GetLac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lac
}

// GetLacOk returns a tuple with the Lac field value
// and a boolean to check if the value has been set.
func (o *RoutingAreaId) GetLacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lac, true
}

// SetLac sets field value
func (o *RoutingAreaId) SetLac(v string) {
	o.Lac = v
}

// GetRac returns the Rac field value
func (o *RoutingAreaId) GetRac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rac
}

// GetRacOk returns a tuple with the Rac field value
// and a boolean to check if the value has been set.
func (o *RoutingAreaId) GetRacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rac, true
}

// SetRac sets field value
func (o *RoutingAreaId) SetRac(v string) {
	o.Rac = v
}

func (o RoutingAreaId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingAreaId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["lac"] = o.Lac
	toSerialize["rac"] = o.Rac
	return toSerialize, nil
}

func (o *RoutingAreaId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plmnId",
		"lac",
		"rac",
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

	varRoutingAreaId := _RoutingAreaId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoutingAreaId)

	if err != nil {
		return err
	}

	*o = RoutingAreaId(varRoutingAreaId)

	return err
}

type NullableRoutingAreaId struct {
	value *RoutingAreaId
	isSet bool
}

func (v NullableRoutingAreaId) Get() *RoutingAreaId {
	return v.value
}

func (v *NullableRoutingAreaId) Set(val *RoutingAreaId) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingAreaId) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingAreaId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingAreaId(val *RoutingAreaId) *NullableRoutingAreaId {
	return &NullableRoutingAreaId{value: val, isSet: true}
}

func (v NullableRoutingAreaId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingAreaId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
