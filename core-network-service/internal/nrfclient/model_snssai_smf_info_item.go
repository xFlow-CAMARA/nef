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

// SnssaiSmfInfoItem struct for SnssaiSmfInfoItem
type SnssaiSmfInfoItem struct {
	SNssai         Snssai           `json:"sNssai"`
	DnnSmfInfoList []DnnSmfInfoItem `json:"dnnSmfInfoList"`
}

// NewSnssaiSmfInfoItem instantiates a new SnssaiSmfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiSmfInfoItem(sNssai Snssai, dnnSmfInfoList []DnnSmfInfoItem) *SnssaiSmfInfoItem {
	this := SnssaiSmfInfoItem{}
	this.SNssai = sNssai
	this.DnnSmfInfoList = dnnSmfInfoList
	return &this
}

// NewSnssaiSmfInfoItemWithDefaults instantiates a new SnssaiSmfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiSmfInfoItemWithDefaults() *SnssaiSmfInfoItem {
	this := SnssaiSmfInfoItem{}
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiSmfInfoItem) GetSNssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiSmfInfoItem) GetSNssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiSmfInfoItem) SetSNssai(v Snssai) {
	o.SNssai = v
}

// GetDnnSmfInfoList returns the DnnSmfInfoList field value
func (o *SnssaiSmfInfoItem) GetDnnSmfInfoList() []DnnSmfInfoItem {
	if o == nil {
		var ret []DnnSmfInfoItem
		return ret
	}

	return o.DnnSmfInfoList
}

// GetDnnSmfInfoListOk returns a tuple with the DnnSmfInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiSmfInfoItem) GetDnnSmfInfoListOk() ([]DnnSmfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnSmfInfoList, true
}

// SetDnnSmfInfoList sets field value
func (o *SnssaiSmfInfoItem) SetDnnSmfInfoList(v []DnnSmfInfoItem) {
	o.DnnSmfInfoList = v
}

func (o SnssaiSmfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sNssai"] = o.SNssai
	}
	if true {
		toSerialize["dnnSmfInfoList"] = o.DnnSmfInfoList
	}
	return json.Marshal(toSerialize)
}

type NullableSnssaiSmfInfoItem struct {
	value *SnssaiSmfInfoItem
	isSet bool
}

func (v NullableSnssaiSmfInfoItem) Get() *SnssaiSmfInfoItem {
	return v.value
}

func (v *NullableSnssaiSmfInfoItem) Set(val *SnssaiSmfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiSmfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiSmfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiSmfInfoItem(val *SnssaiSmfInfoItem) *NullableSnssaiSmfInfoItem {
	return &NullableSnssaiSmfInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiSmfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiSmfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
