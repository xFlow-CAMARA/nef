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

// ChfInfo struct for ChfInfo
type ChfInfo struct {
	SupiRangeList []SupiRange     `json:"supiRangeList,omitempty"`
	GpsiRangeList []IdentityRange `json:"gpsiRangeList,omitempty"`
	PlmnRangeList []PlmnRange     `json:"plmnRangeList,omitempty"`
}

// NewChfInfo instantiates a new ChfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChfInfo() *ChfInfo {
	this := ChfInfo{}
	return &this
}

// NewChfInfoWithDefaults instantiates a new ChfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChfInfoWithDefaults() *ChfInfo {
	this := ChfInfo{}
	return &this
}

// GetSupiRangeList returns the SupiRangeList field value if set, zero value otherwise.
func (o *ChfInfo) GetSupiRangeList() []SupiRange {
	if o == nil || o.SupiRangeList == nil {
		var ret []SupiRange
		return ret
	}
	return o.SupiRangeList
}

// GetSupiRangeListOk returns a tuple with the SupiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetSupiRangeListOk() ([]SupiRange, bool) {
	if o == nil || o.SupiRangeList == nil {
		return nil, false
	}
	return o.SupiRangeList, true
}

// HasSupiRangeList returns a boolean if a field has been set.
func (o *ChfInfo) HasSupiRangeList() bool {
	if o != nil && o.SupiRangeList != nil {
		return true
	}

	return false
}

// SetSupiRangeList gets a reference to the given []SupiRange and assigns it to the SupiRangeList field.
func (o *ChfInfo) SetSupiRangeList(v []SupiRange) {
	o.SupiRangeList = v
}

// GetGpsiRangeList returns the GpsiRangeList field value if set, zero value otherwise.
func (o *ChfInfo) GetGpsiRangeList() []IdentityRange {
	if o == nil || o.GpsiRangeList == nil {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRangeList
}

// GetGpsiRangeListOk returns a tuple with the GpsiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetGpsiRangeListOk() ([]IdentityRange, bool) {
	if o == nil || o.GpsiRangeList == nil {
		return nil, false
	}
	return o.GpsiRangeList, true
}

// HasGpsiRangeList returns a boolean if a field has been set.
func (o *ChfInfo) HasGpsiRangeList() bool {
	if o != nil && o.GpsiRangeList != nil {
		return true
	}

	return false
}

// SetGpsiRangeList gets a reference to the given []IdentityRange and assigns it to the GpsiRangeList field.
func (o *ChfInfo) SetGpsiRangeList(v []IdentityRange) {
	o.GpsiRangeList = v
}

// GetPlmnRangeList returns the PlmnRangeList field value if set, zero value otherwise.
func (o *ChfInfo) GetPlmnRangeList() []PlmnRange {
	if o == nil || o.PlmnRangeList == nil {
		var ret []PlmnRange
		return ret
	}
	return o.PlmnRangeList
}

// GetPlmnRangeListOk returns a tuple with the PlmnRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetPlmnRangeListOk() ([]PlmnRange, bool) {
	if o == nil || o.PlmnRangeList == nil {
		return nil, false
	}
	return o.PlmnRangeList, true
}

// HasPlmnRangeList returns a boolean if a field has been set.
func (o *ChfInfo) HasPlmnRangeList() bool {
	if o != nil && o.PlmnRangeList != nil {
		return true
	}

	return false
}

// SetPlmnRangeList gets a reference to the given []PlmnRange and assigns it to the PlmnRangeList field.
func (o *ChfInfo) SetPlmnRangeList(v []PlmnRange) {
	o.PlmnRangeList = v
}

func (o ChfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SupiRangeList != nil {
		toSerialize["supiRangeList"] = o.SupiRangeList
	}
	if o.GpsiRangeList != nil {
		toSerialize["gpsiRangeList"] = o.GpsiRangeList
	}
	if o.PlmnRangeList != nil {
		toSerialize["plmnRangeList"] = o.PlmnRangeList
	}
	return json.Marshal(toSerialize)
}

type NullableChfInfo struct {
	value *ChfInfo
	isSet bool
}

func (v NullableChfInfo) Get() *ChfInfo {
	return v.value
}

func (v *NullableChfInfo) Set(val *ChfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChfInfo(val *ChfInfo) *NullableChfInfo {
	return &NullableChfInfo{value: val, isSet: true}
}

func (v NullableChfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
