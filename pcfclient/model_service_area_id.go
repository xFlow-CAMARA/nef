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

// checks if the ServiceAreaId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAreaId{}

// ServiceAreaId Contains a Service Area Identifier as defined in 3GPP TS 23.003, clause 12.5.
type ServiceAreaId struct {
	PlmnId PlmnId `json:"plmnId"`
	// Location Area Code.
	Lac string `json:"lac" validate:"regexp=^[A-Fa-f0-9]{4}$"`
	// Service Area Code.
	Sac string `json:"sac" validate:"regexp=^[A-Fa-f0-9]{4}$"`
}

type _ServiceAreaId ServiceAreaId

// NewServiceAreaId instantiates a new ServiceAreaId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAreaId(plmnId PlmnId, lac string, sac string) *ServiceAreaId {
	this := ServiceAreaId{}
	this.PlmnId = plmnId
	this.Lac = lac
	this.Sac = sac
	return &this
}

// NewServiceAreaIdWithDefaults instantiates a new ServiceAreaId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAreaIdWithDefaults() *ServiceAreaId {
	this := ServiceAreaId{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *ServiceAreaId) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *ServiceAreaId) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *ServiceAreaId) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetLac returns the Lac field value
func (o *ServiceAreaId) GetLac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lac
}

// GetLacOk returns a tuple with the Lac field value
// and a boolean to check if the value has been set.
func (o *ServiceAreaId) GetLacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lac, true
}

// SetLac sets field value
func (o *ServiceAreaId) SetLac(v string) {
	o.Lac = v
}

// GetSac returns the Sac field value
func (o *ServiceAreaId) GetSac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sac
}

// GetSacOk returns a tuple with the Sac field value
// and a boolean to check if the value has been set.
func (o *ServiceAreaId) GetSacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sac, true
}

// SetSac sets field value
func (o *ServiceAreaId) SetSac(v string) {
	o.Sac = v
}

func (o ServiceAreaId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAreaId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["lac"] = o.Lac
	toSerialize["sac"] = o.Sac
	return toSerialize, nil
}

func (o *ServiceAreaId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plmnId",
		"lac",
		"sac",
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

	varServiceAreaId := _ServiceAreaId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceAreaId)

	if err != nil {
		return err
	}

	*o = ServiceAreaId(varServiceAreaId)

	return err
}

type NullableServiceAreaId struct {
	value *ServiceAreaId
	isSet bool
}

func (v NullableServiceAreaId) Get() *ServiceAreaId {
	return v.value
}

func (v *NullableServiceAreaId) Set(val *ServiceAreaId) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAreaId) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAreaId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAreaId(val *ServiceAreaId) *NullableServiceAreaId {
	return &NullableServiceAreaId{value: val, isSet: true}
}

func (v NullableServiceAreaId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAreaId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
