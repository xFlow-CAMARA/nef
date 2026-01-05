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

// SearchResult struct for SearchResult
type SearchResult struct {
	ValidityPeriod       *int32      `json:"validityPeriod,omitempty"`
	NfInstances          []NFProfile `json:"nfInstances"`
	SearchId             *string     `json:"searchId,omitempty"`
	NumNfInstComplete    *int32      `json:"numNfInstComplete,omitempty"`
	NrfSupportedFeatures *string     `json:"nrfSupportedFeatures,omitempty"`
}

// NewSearchResult instantiates a new SearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResult(nfInstances []NFProfile) *SearchResult {
	this := SearchResult{}
	this.NfInstances = nfInstances
	return &this
}

// NewSearchResultWithDefaults instantiates a new SearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultWithDefaults() *SearchResult {
	this := SearchResult{}
	return &this
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *SearchResult) GetValidityPeriod() int32 {
	if o == nil || o.ValidityPeriod == nil {
		var ret int32
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetValidityPeriodOk() (*int32, bool) {
	if o == nil || o.ValidityPeriod == nil {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *SearchResult) HasValidityPeriod() bool {
	if o != nil && o.ValidityPeriod != nil {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given int32 and assigns it to the ValidityPeriod field.
func (o *SearchResult) SetValidityPeriod(v int32) {
	o.ValidityPeriod = &v
}

// GetNfInstances returns the NfInstances field value
func (o *SearchResult) GetNfInstances() []NFProfile {
	if o == nil {
		var ret []NFProfile
		return ret
	}

	return o.NfInstances
}

// GetNfInstancesOk returns a tuple with the NfInstances field value
// and a boolean to check if the value has been set.
func (o *SearchResult) GetNfInstancesOk() ([]NFProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.NfInstances, true
}

// SetNfInstances sets field value
func (o *SearchResult) SetNfInstances(v []NFProfile) {
	o.NfInstances = v
}

// GetSearchId returns the SearchId field value if set, zero value otherwise.
func (o *SearchResult) GetSearchId() string {
	if o == nil || o.SearchId == nil {
		var ret string
		return ret
	}
	return *o.SearchId
}

// GetSearchIdOk returns a tuple with the SearchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetSearchIdOk() (*string, bool) {
	if o == nil || o.SearchId == nil {
		return nil, false
	}
	return o.SearchId, true
}

// HasSearchId returns a boolean if a field has been set.
func (o *SearchResult) HasSearchId() bool {
	if o != nil && o.SearchId != nil {
		return true
	}

	return false
}

// SetSearchId gets a reference to the given string and assigns it to the SearchId field.
func (o *SearchResult) SetSearchId(v string) {
	o.SearchId = &v
}

// GetNumNfInstComplete returns the NumNfInstComplete field value if set, zero value otherwise.
func (o *SearchResult) GetNumNfInstComplete() int32 {
	if o == nil || o.NumNfInstComplete == nil {
		var ret int32
		return ret
	}
	return *o.NumNfInstComplete
}

// GetNumNfInstCompleteOk returns a tuple with the NumNfInstComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetNumNfInstCompleteOk() (*int32, bool) {
	if o == nil || o.NumNfInstComplete == nil {
		return nil, false
	}
	return o.NumNfInstComplete, true
}

// HasNumNfInstComplete returns a boolean if a field has been set.
func (o *SearchResult) HasNumNfInstComplete() bool {
	if o != nil && o.NumNfInstComplete != nil {
		return true
	}

	return false
}

// SetNumNfInstComplete gets a reference to the given int32 and assigns it to the NumNfInstComplete field.
func (o *SearchResult) SetNumNfInstComplete(v int32) {
	o.NumNfInstComplete = &v
}

// GetNrfSupportedFeatures returns the NrfSupportedFeatures field value if set, zero value otherwise.
func (o *SearchResult) GetNrfSupportedFeatures() string {
	if o == nil || o.NrfSupportedFeatures == nil {
		var ret string
		return ret
	}
	return *o.NrfSupportedFeatures
}

// GetNrfSupportedFeaturesOk returns a tuple with the NrfSupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetNrfSupportedFeaturesOk() (*string, bool) {
	if o == nil || o.NrfSupportedFeatures == nil {
		return nil, false
	}
	return o.NrfSupportedFeatures, true
}

// HasNrfSupportedFeatures returns a boolean if a field has been set.
func (o *SearchResult) HasNrfSupportedFeatures() bool {
	if o != nil && o.NrfSupportedFeatures != nil {
		return true
	}

	return false
}

// SetNrfSupportedFeatures gets a reference to the given string and assigns it to the NrfSupportedFeatures field.
func (o *SearchResult) SetNrfSupportedFeatures(v string) {
	o.NrfSupportedFeatures = &v
}

func (o SearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidityPeriod != nil {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}
	if true {
		toSerialize["nfInstances"] = o.NfInstances
	}
	if o.SearchId != nil {
		toSerialize["searchId"] = o.SearchId
	}
	if o.NumNfInstComplete != nil {
		toSerialize["numNfInstComplete"] = o.NumNfInstComplete
	}
	if o.NrfSupportedFeatures != nil {
		toSerialize["nrfSupportedFeatures"] = o.NrfSupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResult struct {
	value *SearchResult
	isSet bool
}

func (v NullableSearchResult) Get() *SearchResult {
	return v.value
}

func (v *NullableSearchResult) Set(val *SearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResult(val *SearchResult) *NullableSearchResult {
	return &NullableSearchResult{value: val, isSet: true}
}

func (v NullableSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
