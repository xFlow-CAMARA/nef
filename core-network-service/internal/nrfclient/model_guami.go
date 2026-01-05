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
)

// Guami struct for Guami
type Guami struct {
	PlmnId PlmnId `json:"plmnId"`
	AmfId  string `json:"amfId"`
}

// NewGuami instantiates a new Guami object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuami(plmnId PlmnId, amfId string) *Guami {
	this := Guami{}
	this.PlmnId = plmnId
	this.AmfId = amfId
	return &this
}

// NewGuamiWithDefaults instantiates a new Guami object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuamiWithDefaults() *Guami {
	this := Guami{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Guami) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Guami) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Guami) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetAmfId returns the AmfId field value
func (o *Guami) GetAmfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfId
}

// GetAmfIdOk returns a tuple with the AmfId field value
// and a boolean to check if the value has been set.
func (o *Guami) GetAmfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmfId, true
}

// SetAmfId sets field value
func (o *Guami) SetAmfId(v string) {
	o.AmfId = v
}

func (o Guami) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["amfId"] = o.AmfId
	}
	return json.Marshal(toSerialize)
}

type NullableGuami struct {
	value *Guami
	isSet bool
}

func (v NullableGuami) Get() *Guami {
	return v.value
}

func (v *NullableGuami) Set(val *Guami) {
	v.value = val
	v.isSet = true
}

func (v NullableGuami) IsSet() bool {
	return v.isSet
}

func (v *NullableGuami) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuami(val *Guami) *NullableGuami {
	return &NullableGuami{value: val, isSet: true}
}

func (v NullableGuami) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuami) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
