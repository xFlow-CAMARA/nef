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

// Atom struct for Atom
type Atom struct {
	Attr     string      `json:"attr"`
	Value    interface{} `json:"value"`
	Negative *bool       `json:"negative,omitempty"`
}

// NewAtom instantiates a new Atom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtom(attr string, value interface{}) *Atom {
	this := Atom{}
	this.Attr = attr
	this.Value = value
	return &this
}

// NewAtomWithDefaults instantiates a new Atom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtomWithDefaults() *Atom {
	this := Atom{}
	return &this
}

// GetAttr returns the Attr field value
func (o *Atom) GetAttr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attr
}

// GetAttrOk returns a tuple with the Attr field value
// and a boolean to check if the value has been set.
func (o *Atom) GetAttrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attr, true
}

// SetAttr sets field value
func (o *Atom) SetAttr(v string) {
	o.Attr = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Atom) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Atom) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Atom) SetValue(v interface{}) {
	o.Value = v
}

// GetNegative returns the Negative field value if set, zero value otherwise.
func (o *Atom) GetNegative() bool {
	if o == nil || o.Negative == nil {
		var ret bool
		return ret
	}
	return *o.Negative
}

// GetNegativeOk returns a tuple with the Negative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Atom) GetNegativeOk() (*bool, bool) {
	if o == nil || o.Negative == nil {
		return nil, false
	}
	return o.Negative, true
}

// HasNegative returns a boolean if a field has been set.
func (o *Atom) HasNegative() bool {
	if o != nil && o.Negative != nil {
		return true
	}

	return false
}

// SetNegative gets a reference to the given bool and assigns it to the Negative field.
func (o *Atom) SetNegative(v bool) {
	o.Negative = &v
}

func (o Atom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attr"] = o.Attr
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Negative != nil {
		toSerialize["negative"] = o.Negative
	}
	return json.Marshal(toSerialize)
}

type NullableAtom struct {
	value *Atom
	isSet bool
}

func (v NullableAtom) Get() *Atom {
	return v.value
}

func (v *NullableAtom) Set(val *Atom) {
	v.value = val
	v.isSet = true
}

func (v NullableAtom) IsSet() bool {
	return v.isSet
}

func (v *NullableAtom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtom(val *Atom) *NullableAtom {
	return &NullableAtom{value: val, isSet: true}
}

func (v NullableAtom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
