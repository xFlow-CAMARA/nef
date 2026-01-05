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

// DnnSmfInfoItem struct for DnnSmfInfoItem
type DnnSmfInfoItem struct {
	Dnn string `json:"dnn"`
}

// NewDnnSmfInfoItem instantiates a new DnnSmfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnSmfInfoItem(dnn string) *DnnSmfInfoItem {
	this := DnnSmfInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnSmfInfoItemWithDefaults instantiates a new DnnSmfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnSmfInfoItemWithDefaults() *DnnSmfInfoItem {
	this := DnnSmfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnSmfInfoItem) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnSmfInfoItem) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnSmfInfoItem) SetDnn(v string) {
	o.Dnn = v
}

func (o DnnSmfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	return json.Marshal(toSerialize)
}

type NullableDnnSmfInfoItem struct {
	value *DnnSmfInfoItem
	isSet bool
}

func (v NullableDnnSmfInfoItem) Get() *DnnSmfInfoItem {
	return v.value
}

func (v *NullableDnnSmfInfoItem) Set(val *DnnSmfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnSmfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnSmfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnSmfInfoItem(val *DnnSmfInfoItem) *NullableDnnSmfInfoItem {
	return &NullableDnnSmfInfoItem{value: val, isSet: true}
}

func (v NullableDnnSmfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnSmfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
