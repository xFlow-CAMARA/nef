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

// Snssai struct for Snssai
type Snssai struct {
	Sst int32   `json:"sst"`
	Sd  *string `json:"sd,omitempty"`
}

// NewSnssai instantiates a new Snssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssai(sst int32) *Snssai {
	this := Snssai{}
	this.Sst = sst
	return &this
}

// NewSnssaiWithDefaults instantiates a new Snssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiWithDefaults() *Snssai {
	this := Snssai{}
	return &this
}

// GetSst returns the Sst field value
func (o *Snssai) GetSst() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sst
}

// GetSstOk returns a tuple with the Sst field value
// and a boolean to check if the value has been set.
func (o *Snssai) GetSstOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sst, true
}

// SetSst sets field value
func (o *Snssai) SetSst(v int32) {
	o.Sst = v
}

// GetSd returns the Sd field value if set, zero value otherwise.
func (o *Snssai) GetSd() string {
	if o == nil || o.Sd == nil {
		var ret string
		return ret
	}
	return *o.Sd
}

// GetSdOk returns a tuple with the Sd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snssai) GetSdOk() (*string, bool) {
	if o == nil || o.Sd == nil {
		return nil, false
	}
	return o.Sd, true
}

// HasSd returns a boolean if a field has been set.
func (o *Snssai) HasSd() bool {
	if o != nil && o.Sd != nil {
		return true
	}

	return false
}

// SetSd gets a reference to the given string and assigns it to the Sd field.
func (o *Snssai) SetSd(v string) {
	o.Sd = &v
}

func (o Snssai) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sst"] = o.Sst
	}
	if o.Sd != nil {
		toSerialize["sd"] = o.Sd
	}
	return json.Marshal(toSerialize)
}

type NullableSnssai struct {
	value *Snssai
	isSet bool
}

func (v NullableSnssai) Get() *Snssai {
	return v.value
}

func (v *NullableSnssai) Set(val *Snssai) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssai(val *Snssai) *NullableSnssai {
	return &NullableSnssai{value: val, isSet: true}
}

func (v NullableSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
