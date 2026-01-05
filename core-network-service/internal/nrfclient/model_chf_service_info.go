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

// ChfServiceInfo struct for ChfServiceInfo
type ChfServiceInfo struct {
	PrimaryChfServiceInstance   *string `json:"primaryChfServiceInstance,omitempty"`
	SecondaryChfServiceInstance *string `json:"secondaryChfServiceInstance,omitempty"`
}

// NewChfServiceInfo instantiates a new ChfServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChfServiceInfo() *ChfServiceInfo {
	this := ChfServiceInfo{}
	return &this
}

// NewChfServiceInfoWithDefaults instantiates a new ChfServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChfServiceInfoWithDefaults() *ChfServiceInfo {
	this := ChfServiceInfo{}
	return &this
}

// GetPrimaryChfServiceInstance returns the PrimaryChfServiceInstance field value if set, zero value otherwise.
func (o *ChfServiceInfo) GetPrimaryChfServiceInstance() string {
	if o == nil || o.PrimaryChfServiceInstance == nil {
		var ret string
		return ret
	}
	return *o.PrimaryChfServiceInstance
}

// GetPrimaryChfServiceInstanceOk returns a tuple with the PrimaryChfServiceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfServiceInfo) GetPrimaryChfServiceInstanceOk() (*string, bool) {
	if o == nil || o.PrimaryChfServiceInstance == nil {
		return nil, false
	}
	return o.PrimaryChfServiceInstance, true
}

// HasPrimaryChfServiceInstance returns a boolean if a field has been set.
func (o *ChfServiceInfo) HasPrimaryChfServiceInstance() bool {
	if o != nil && o.PrimaryChfServiceInstance != nil {
		return true
	}

	return false
}

// SetPrimaryChfServiceInstance gets a reference to the given string and assigns it to the PrimaryChfServiceInstance field.
func (o *ChfServiceInfo) SetPrimaryChfServiceInstance(v string) {
	o.PrimaryChfServiceInstance = &v
}

// GetSecondaryChfServiceInstance returns the SecondaryChfServiceInstance field value if set, zero value otherwise.
func (o *ChfServiceInfo) GetSecondaryChfServiceInstance() string {
	if o == nil || o.SecondaryChfServiceInstance == nil {
		var ret string
		return ret
	}
	return *o.SecondaryChfServiceInstance
}

// GetSecondaryChfServiceInstanceOk returns a tuple with the SecondaryChfServiceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfServiceInfo) GetSecondaryChfServiceInstanceOk() (*string, bool) {
	if o == nil || o.SecondaryChfServiceInstance == nil {
		return nil, false
	}
	return o.SecondaryChfServiceInstance, true
}

// HasSecondaryChfServiceInstance returns a boolean if a field has been set.
func (o *ChfServiceInfo) HasSecondaryChfServiceInstance() bool {
	if o != nil && o.SecondaryChfServiceInstance != nil {
		return true
	}

	return false
}

// SetSecondaryChfServiceInstance gets a reference to the given string and assigns it to the SecondaryChfServiceInstance field.
func (o *ChfServiceInfo) SetSecondaryChfServiceInstance(v string) {
	o.SecondaryChfServiceInstance = &v
}

func (o ChfServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrimaryChfServiceInstance != nil {
		toSerialize["primaryChfServiceInstance"] = o.PrimaryChfServiceInstance
	}
	if o.SecondaryChfServiceInstance != nil {
		toSerialize["secondaryChfServiceInstance"] = o.SecondaryChfServiceInstance
	}
	return json.Marshal(toSerialize)
}

type NullableChfServiceInfo struct {
	value *ChfServiceInfo
	isSet bool
}

func (v NullableChfServiceInfo) Get() *ChfServiceInfo {
	return v.value
}

func (v *NullableChfServiceInfo) Set(val *ChfServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChfServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChfServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChfServiceInfo(val *ChfServiceInfo) *NullableChfServiceInfo {
	return &NullableChfServiceInfo{value: val, isSet: true}
}

func (v NullableChfServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChfServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
