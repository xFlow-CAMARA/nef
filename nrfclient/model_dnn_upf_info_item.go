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

// DnnUpfInfoItem struct for DnnUpfInfoItem
type DnnUpfInfoItem struct {
	Dnn             string           `json:"dnn"`
	DnaiList        []string         `json:"dnaiList,omitempty"`
	PduSessionTypes []PduSessionType `json:"pduSessionTypes,omitempty"`
}

// NewDnnUpfInfoItem instantiates a new DnnUpfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnUpfInfoItem(dnn string) *DnnUpfInfoItem {
	this := DnnUpfInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnUpfInfoItemWithDefaults instantiates a new DnnUpfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnUpfInfoItemWithDefaults() *DnnUpfInfoItem {
	this := DnnUpfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnUpfInfoItem) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnUpfInfoItem) SetDnn(v string) {
	o.Dnn = v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetDnaiList() []string {
	if o == nil || o.DnaiList == nil {
		var ret []string
		return ret
	}
	return o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetDnaiListOk() ([]string, bool) {
	if o == nil || o.DnaiList == nil {
		return nil, false
	}
	return o.DnaiList, true
}

// HasDnaiList returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasDnaiList() bool {
	if o != nil && o.DnaiList != nil {
		return true
	}

	return false
}

// SetDnaiList gets a reference to the given []string and assigns it to the DnaiList field.
func (o *DnnUpfInfoItem) SetDnaiList(v []string) {
	o.DnaiList = v
}

// GetPduSessionTypes returns the PduSessionTypes field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetPduSessionTypes() []PduSessionType {
	if o == nil || o.PduSessionTypes == nil {
		var ret []PduSessionType
		return ret
	}
	return o.PduSessionTypes
}

// GetPduSessionTypesOk returns a tuple with the PduSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetPduSessionTypesOk() ([]PduSessionType, bool) {
	if o == nil || o.PduSessionTypes == nil {
		return nil, false
	}
	return o.PduSessionTypes, true
}

// HasPduSessionTypes returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasPduSessionTypes() bool {
	if o != nil && o.PduSessionTypes != nil {
		return true
	}

	return false
}

// SetPduSessionTypes gets a reference to the given []PduSessionType and assigns it to the PduSessionTypes field.
func (o *DnnUpfInfoItem) SetPduSessionTypes(v []PduSessionType) {
	o.PduSessionTypes = v
}

func (o DnnUpfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	if o.DnaiList != nil {
		toSerialize["dnaiList"] = o.DnaiList
	}
	if o.PduSessionTypes != nil {
		toSerialize["pduSessionTypes"] = o.PduSessionTypes
	}
	return json.Marshal(toSerialize)
}

type NullableDnnUpfInfoItem struct {
	value *DnnUpfInfoItem
	isSet bool
}

func (v NullableDnnUpfInfoItem) Get() *DnnUpfInfoItem {
	return v.value
}

func (v *NullableDnnUpfInfoItem) Set(val *DnnUpfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnUpfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnUpfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnUpfInfoItem(val *DnnUpfInfoItem) *NullableDnnUpfInfoItem {
	return &NullableDnnUpfInfoItem{value: val, isSet: true}
}

func (v NullableDnnUpfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnUpfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
