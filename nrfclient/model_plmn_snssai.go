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
NRF NFDiscovery Service

NRF NFDiscovery Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0.alpha-1
*/

package nrfclient

import (
	"encoding/json"
)

// PlmnSnssai struct for PlmnSnssai
type PlmnSnssai struct {
	PlmnId     PlmnId   `json:"plmnId"`
	SNssaiList []Snssai `json:"sNssaiList"`
}

// NewPlmnSnssai instantiates a new PlmnSnssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnSnssai(plmnId PlmnId, sNssaiList []Snssai) *PlmnSnssai {
	this := PlmnSnssai{}
	this.PlmnId = plmnId
	this.SNssaiList = sNssaiList
	return &this
}

// NewPlmnSnssaiWithDefaults instantiates a new PlmnSnssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnSnssaiWithDefaults() *PlmnSnssai {
	this := PlmnSnssai{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *PlmnSnssai) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *PlmnSnssai) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *PlmnSnssai) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetSNssaiList returns the SNssaiList field value
func (o *PlmnSnssai) GetSNssaiList() []Snssai {
	if o == nil {
		var ret []Snssai
		return ret
	}

	return o.SNssaiList
}

// GetSNssaiListOk returns a tuple with the SNssaiList field value
// and a boolean to check if the value has been set.
func (o *PlmnSnssai) GetSNssaiListOk() ([]Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return o.SNssaiList, true
}

// SetSNssaiList sets field value
func (o *PlmnSnssai) SetSNssaiList(v []Snssai) {
	o.SNssaiList = v
}

func (o PlmnSnssai) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["sNssaiList"] = o.SNssaiList
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnSnssai struct {
	value *PlmnSnssai
	isSet bool
}

func (v NullablePlmnSnssai) Get() *PlmnSnssai {
	return v.value
}

func (v *NullablePlmnSnssai) Set(val *PlmnSnssai) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnSnssai(val *PlmnSnssai) *NullablePlmnSnssai {
	return &NullablePlmnSnssai{value: val, isSet: true}
}

func (v NullablePlmnSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
