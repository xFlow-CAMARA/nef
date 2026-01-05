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

// TaiRange struct for TaiRange
type TaiRange struct {
	PlmnId       PlmnId     `json:"plmnId"`
	TacRangeList []TacRange `json:"tacRangeList"`
}

// NewTaiRange instantiates a new TaiRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaiRange(plmnId PlmnId, tacRangeList []TacRange) *TaiRange {
	this := TaiRange{}
	this.PlmnId = plmnId
	this.TacRangeList = tacRangeList
	return &this
}

// NewTaiRangeWithDefaults instantiates a new TaiRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaiRangeWithDefaults() *TaiRange {
	this := TaiRange{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *TaiRange) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *TaiRange) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *TaiRange) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetTacRangeList returns the TacRangeList field value
func (o *TaiRange) GetTacRangeList() []TacRange {
	if o == nil {
		var ret []TacRange
		return ret
	}

	return o.TacRangeList
}

// GetTacRangeListOk returns a tuple with the TacRangeList field value
// and a boolean to check if the value has been set.
func (o *TaiRange) GetTacRangeListOk() ([]TacRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.TacRangeList, true
}

// SetTacRangeList sets field value
func (o *TaiRange) SetTacRangeList(v []TacRange) {
	o.TacRangeList = v
}

func (o TaiRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["tacRangeList"] = o.TacRangeList
	}
	return json.Marshal(toSerialize)
}

type NullableTaiRange struct {
	value *TaiRange
	isSet bool
}

func (v NullableTaiRange) Get() *TaiRange {
	return v.value
}

func (v *NullableTaiRange) Set(val *TaiRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTaiRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTaiRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaiRange(val *TaiRange) *NullableTaiRange {
	return &NullableTaiRange{value: val, isSet: true}
}

func (v NullableTaiRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaiRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
