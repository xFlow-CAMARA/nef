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

// DnfUnit struct for DnfUnit
type DnfUnit struct {
	DnfUnit []Atom `json:"dnfUnit"`
}

// NewDnfUnit instantiates a new DnfUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnfUnit(dnfUnit []Atom) *DnfUnit {
	this := DnfUnit{}
	this.DnfUnit = dnfUnit
	return &this
}

// NewDnfUnitWithDefaults instantiates a new DnfUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnfUnitWithDefaults() *DnfUnit {
	this := DnfUnit{}
	return &this
}

// GetDnfUnit returns the DnfUnit field value
func (o *DnfUnit) GetDnfUnit() []Atom {
	if o == nil {
		var ret []Atom
		return ret
	}

	return o.DnfUnit
}

// GetDnfUnitOk returns a tuple with the DnfUnit field value
// and a boolean to check if the value has been set.
func (o *DnfUnit) GetDnfUnitOk() ([]Atom, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnfUnit, true
}

// SetDnfUnit sets field value
func (o *DnfUnit) SetDnfUnit(v []Atom) {
	o.DnfUnit = v
}

func (o DnfUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnfUnit"] = o.DnfUnit
	}
	return json.Marshal(toSerialize)
}

type NullableDnfUnit struct {
	value *DnfUnit
	isSet bool
}

func (v NullableDnfUnit) Get() *DnfUnit {
	return v.value
}

func (v *NullableDnfUnit) Set(val *DnfUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableDnfUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableDnfUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnfUnit(val *DnfUnit) *NullableDnfUnit {
	return &NullableDnfUnit{value: val, isSet: true}
}

func (v NullableDnfUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnfUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
