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

// StoredSearchResult struct for StoredSearchResult
type StoredSearchResult struct {
	NfInstances []NFProfile `json:"nfInstances"`
}

// NewStoredSearchResult instantiates a new StoredSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredSearchResult(nfInstances []NFProfile) *StoredSearchResult {
	this := StoredSearchResult{}
	this.NfInstances = nfInstances
	return &this
}

// NewStoredSearchResultWithDefaults instantiates a new StoredSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredSearchResultWithDefaults() *StoredSearchResult {
	this := StoredSearchResult{}
	return &this
}

// GetNfInstances returns the NfInstances field value
func (o *StoredSearchResult) GetNfInstances() []NFProfile {
	if o == nil {
		var ret []NFProfile
		return ret
	}

	return o.NfInstances
}

// GetNfInstancesOk returns a tuple with the NfInstances field value
// and a boolean to check if the value has been set.
func (o *StoredSearchResult) GetNfInstancesOk() ([]NFProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.NfInstances, true
}

// SetNfInstances sets field value
func (o *StoredSearchResult) SetNfInstances(v []NFProfile) {
	o.NfInstances = v
}

func (o StoredSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfInstances"] = o.NfInstances
	}
	return json.Marshal(toSerialize)
}

type NullableStoredSearchResult struct {
	value *StoredSearchResult
	isSet bool
}

func (v NullableStoredSearchResult) Get() *StoredSearchResult {
	return v.value
}

func (v *NullableStoredSearchResult) Set(val *StoredSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredSearchResult(val *StoredSearchResult) *NullableStoredSearchResult {
	return &NullableStoredSearchResult{value: val, isSet: true}
}

func (v NullableStoredSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
