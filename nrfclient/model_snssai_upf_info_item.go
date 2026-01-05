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

// SnssaiUpfInfoItem struct for SnssaiUpfInfoItem
type SnssaiUpfInfoItem struct {
	SNssai         Snssai           `json:"sNssai"`
	DnnUpfInfoList []DnnUpfInfoItem `json:"dnnUpfInfoList"`
}

// NewSnssaiUpfInfoItem instantiates a new SnssaiUpfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiUpfInfoItem(sNssai Snssai, dnnUpfInfoList []DnnUpfInfoItem) *SnssaiUpfInfoItem {
	this := SnssaiUpfInfoItem{}
	this.SNssai = sNssai
	this.DnnUpfInfoList = dnnUpfInfoList
	return &this
}

// NewSnssaiUpfInfoItemWithDefaults instantiates a new SnssaiUpfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiUpfInfoItemWithDefaults() *SnssaiUpfInfoItem {
	this := SnssaiUpfInfoItem{}
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiUpfInfoItem) GetSNssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiUpfInfoItem) GetSNssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiUpfInfoItem) SetSNssai(v Snssai) {
	o.SNssai = v
}

// GetDnnUpfInfoList returns the DnnUpfInfoList field value
func (o *SnssaiUpfInfoItem) GetDnnUpfInfoList() []DnnUpfInfoItem {
	if o == nil {
		var ret []DnnUpfInfoItem
		return ret
	}

	return o.DnnUpfInfoList
}

// GetDnnUpfInfoListOk returns a tuple with the DnnUpfInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiUpfInfoItem) GetDnnUpfInfoListOk() ([]DnnUpfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnUpfInfoList, true
}

// SetDnnUpfInfoList sets field value
func (o *SnssaiUpfInfoItem) SetDnnUpfInfoList(v []DnnUpfInfoItem) {
	o.DnnUpfInfoList = v
}

func (o SnssaiUpfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sNssai"] = o.SNssai
	}
	if true {
		toSerialize["dnnUpfInfoList"] = o.DnnUpfInfoList
	}
	return json.Marshal(toSerialize)
}

type NullableSnssaiUpfInfoItem struct {
	value *SnssaiUpfInfoItem
	isSet bool
}

func (v NullableSnssaiUpfInfoItem) Get() *SnssaiUpfInfoItem {
	return v.value
}

func (v *NullableSnssaiUpfInfoItem) Set(val *SnssaiUpfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiUpfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiUpfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiUpfInfoItem(val *SnssaiUpfInfoItem) *NullableSnssaiUpfInfoItem {
	return &NullableSnssaiUpfInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiUpfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiUpfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
