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

// UdrInfo struct for UdrInfo
type UdrInfo struct {
	GroupId                        *string         `json:"groupId,omitempty"`
	SupiRanges                     []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges                     []IdentityRange `json:"gpsiRanges,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`
	SupportedDataSets              []DataSetId     `json:"supportedDataSets,omitempty"`
}

// NewUdrInfo instantiates a new UdrInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdrInfo() *UdrInfo {
	this := UdrInfo{}
	return &this
}

// NewUdrInfoWithDefaults instantiates a new UdrInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdrInfoWithDefaults() *UdrInfo {
	this := UdrInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *UdrInfo) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *UdrInfo) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *UdrInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *UdrInfo) GetSupiRanges() []SupiRange {
	if o == nil || o.SupiRanges == nil {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || o.SupiRanges == nil {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *UdrInfo) HasSupiRanges() bool {
	if o != nil && o.SupiRanges != nil {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *UdrInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *UdrInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || o.GpsiRanges == nil {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || o.GpsiRanges == nil {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *UdrInfo) HasGpsiRanges() bool {
	if o != nil && o.GpsiRanges != nil {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *UdrInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *UdrInfo) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || o.ExternalGroupIdentifiersRanges == nil {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || o.ExternalGroupIdentifiersRanges == nil {
		return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *UdrInfo) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && o.ExternalGroupIdentifiersRanges != nil {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *UdrInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetSupportedDataSets returns the SupportedDataSets field value if set, zero value otherwise.
func (o *UdrInfo) GetSupportedDataSets() []DataSetId {
	if o == nil || o.SupportedDataSets == nil {
		var ret []DataSetId
		return ret
	}
	return o.SupportedDataSets
}

// GetSupportedDataSetsOk returns a tuple with the SupportedDataSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetSupportedDataSetsOk() ([]DataSetId, bool) {
	if o == nil || o.SupportedDataSets == nil {
		return nil, false
	}
	return o.SupportedDataSets, true
}

// HasSupportedDataSets returns a boolean if a field has been set.
func (o *UdrInfo) HasSupportedDataSets() bool {
	if o != nil && o.SupportedDataSets != nil {
		return true
	}

	return false
}

// SetSupportedDataSets gets a reference to the given []DataSetId and assigns it to the SupportedDataSets field.
func (o *UdrInfo) SetSupportedDataSets(v []DataSetId) {
	o.SupportedDataSets = v
}

func (o UdrInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.SupiRanges != nil {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if o.GpsiRanges != nil {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if o.ExternalGroupIdentifiersRanges != nil {
		toSerialize["externalGroupIdentifiersRanges"] = o.ExternalGroupIdentifiersRanges
	}
	if o.SupportedDataSets != nil {
		toSerialize["supportedDataSets"] = o.SupportedDataSets
	}
	return json.Marshal(toSerialize)
}

type NullableUdrInfo struct {
	value *UdrInfo
	isSet bool
}

func (v NullableUdrInfo) Get() *UdrInfo {
	return v.value
}

func (v *NullableUdrInfo) Set(val *UdrInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUdrInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUdrInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdrInfo(val *UdrInfo) *NullableUdrInfo {
	return &NullableUdrInfo{value: val, isSet: true}
}

func (v NullableUdrInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdrInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
