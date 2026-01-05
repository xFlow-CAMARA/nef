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

// PcfInfo struct for PcfInfo
type PcfInfo struct {
	GroupId     *string         `json:"groupId,omitempty"`
	DnnList     []string        `json:"dnnList,omitempty"`
	SupiRanges  []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges  []IdentityRange `json:"gpsiRanges,omitempty"`
	RxDiamHost  *string         `json:"rxDiamHost,omitempty"`
	RxDiamRealm *string         `json:"rxDiamRealm,omitempty"`
}

// NewPcfInfo instantiates a new PcfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfInfo() *PcfInfo {
	this := PcfInfo{}
	return &this
}

// NewPcfInfoWithDefaults instantiates a new PcfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfInfoWithDefaults() *PcfInfo {
	this := PcfInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *PcfInfo) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PcfInfo) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PcfInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *PcfInfo) GetDnnList() []string {
	if o == nil || o.DnnList == nil {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetDnnListOk() ([]string, bool) {
	if o == nil || o.DnnList == nil {
		return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *PcfInfo) HasDnnList() bool {
	if o != nil && o.DnnList != nil {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *PcfInfo) SetDnnList(v []string) {
	o.DnnList = v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *PcfInfo) GetSupiRanges() []SupiRange {
	if o == nil || o.SupiRanges == nil {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || o.SupiRanges == nil {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *PcfInfo) HasSupiRanges() bool {
	if o != nil && o.SupiRanges != nil {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *PcfInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *PcfInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || o.GpsiRanges == nil {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || o.GpsiRanges == nil {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *PcfInfo) HasGpsiRanges() bool {
	if o != nil && o.GpsiRanges != nil {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *PcfInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetRxDiamHost returns the RxDiamHost field value if set, zero value otherwise.
func (o *PcfInfo) GetRxDiamHost() string {
	if o == nil || o.RxDiamHost == nil {
		var ret string
		return ret
	}
	return *o.RxDiamHost
}

// GetRxDiamHostOk returns a tuple with the RxDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetRxDiamHostOk() (*string, bool) {
	if o == nil || o.RxDiamHost == nil {
		return nil, false
	}
	return o.RxDiamHost, true
}

// HasRxDiamHost returns a boolean if a field has been set.
func (o *PcfInfo) HasRxDiamHost() bool {
	if o != nil && o.RxDiamHost != nil {
		return true
	}

	return false
}

// SetRxDiamHost gets a reference to the given string and assigns it to the RxDiamHost field.
func (o *PcfInfo) SetRxDiamHost(v string) {
	o.RxDiamHost = &v
}

// GetRxDiamRealm returns the RxDiamRealm field value if set, zero value otherwise.
func (o *PcfInfo) GetRxDiamRealm() string {
	if o == nil || o.RxDiamRealm == nil {
		var ret string
		return ret
	}
	return *o.RxDiamRealm
}

// GetRxDiamRealmOk returns a tuple with the RxDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetRxDiamRealmOk() (*string, bool) {
	if o == nil || o.RxDiamRealm == nil {
		return nil, false
	}
	return o.RxDiamRealm, true
}

// HasRxDiamRealm returns a boolean if a field has been set.
func (o *PcfInfo) HasRxDiamRealm() bool {
	if o != nil && o.RxDiamRealm != nil {
		return true
	}

	return false
}

// SetRxDiamRealm gets a reference to the given string and assigns it to the RxDiamRealm field.
func (o *PcfInfo) SetRxDiamRealm(v string) {
	o.RxDiamRealm = &v
}

func (o PcfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.DnnList != nil {
		toSerialize["dnnList"] = o.DnnList
	}
	if o.SupiRanges != nil {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if o.GpsiRanges != nil {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if o.RxDiamHost != nil {
		toSerialize["rxDiamHost"] = o.RxDiamHost
	}
	if o.RxDiamRealm != nil {
		toSerialize["rxDiamRealm"] = o.RxDiamRealm
	}
	return json.Marshal(toSerialize)
}

type NullablePcfInfo struct {
	value *PcfInfo
	isSet bool
}

func (v NullablePcfInfo) Get() *PcfInfo {
	return v.value
}

func (v *NullablePcfInfo) Set(val *PcfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfInfo(val *PcfInfo) *NullablePcfInfo {
	return &NullablePcfInfo{value: val, isSet: true}
}

func (v NullablePcfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
