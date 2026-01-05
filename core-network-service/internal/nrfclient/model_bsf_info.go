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

// BsfInfo struct for BsfInfo
type BsfInfo struct {
	DnnList           []string           `json:"dnnList,omitempty"`
	IpDomainList      []string           `json:"ipDomainList,omitempty"`
	Ipv4AddressRanges []Ipv4AddressRange `json:"ipv4AddressRanges,omitempty"`
	Ipv6PrefixRanges  []Ipv6PrefixRange  `json:"ipv6PrefixRanges,omitempty"`
}

// NewBsfInfo instantiates a new BsfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBsfInfo() *BsfInfo {
	this := BsfInfo{}
	return &this
}

// NewBsfInfoWithDefaults instantiates a new BsfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBsfInfoWithDefaults() *BsfInfo {
	this := BsfInfo{}
	return &this
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *BsfInfo) GetDnnList() []string {
	if o == nil || o.DnnList == nil {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetDnnListOk() ([]string, bool) {
	if o == nil || o.DnnList == nil {
		return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *BsfInfo) HasDnnList() bool {
	if o != nil && o.DnnList != nil {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *BsfInfo) SetDnnList(v []string) {
	o.DnnList = v
}

// GetIpDomainList returns the IpDomainList field value if set, zero value otherwise.
func (o *BsfInfo) GetIpDomainList() []string {
	if o == nil || o.IpDomainList == nil {
		var ret []string
		return ret
	}
	return o.IpDomainList
}

// GetIpDomainListOk returns a tuple with the IpDomainList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetIpDomainListOk() ([]string, bool) {
	if o == nil || o.IpDomainList == nil {
		return nil, false
	}
	return o.IpDomainList, true
}

// HasIpDomainList returns a boolean if a field has been set.
func (o *BsfInfo) HasIpDomainList() bool {
	if o != nil && o.IpDomainList != nil {
		return true
	}

	return false
}

// SetIpDomainList gets a reference to the given []string and assigns it to the IpDomainList field.
func (o *BsfInfo) SetIpDomainList(v []string) {
	o.IpDomainList = v
}

// GetIpv4AddressRanges returns the Ipv4AddressRanges field value if set, zero value otherwise.
func (o *BsfInfo) GetIpv4AddressRanges() []Ipv4AddressRange {
	if o == nil || o.Ipv4AddressRanges == nil {
		var ret []Ipv4AddressRange
		return ret
	}
	return o.Ipv4AddressRanges
}

// GetIpv4AddressRangesOk returns a tuple with the Ipv4AddressRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetIpv4AddressRangesOk() ([]Ipv4AddressRange, bool) {
	if o == nil || o.Ipv4AddressRanges == nil {
		return nil, false
	}
	return o.Ipv4AddressRanges, true
}

// HasIpv4AddressRanges returns a boolean if a field has been set.
func (o *BsfInfo) HasIpv4AddressRanges() bool {
	if o != nil && o.Ipv4AddressRanges != nil {
		return true
	}

	return false
}

// SetIpv4AddressRanges gets a reference to the given []Ipv4AddressRange and assigns it to the Ipv4AddressRanges field.
func (o *BsfInfo) SetIpv4AddressRanges(v []Ipv4AddressRange) {
	o.Ipv4AddressRanges = v
}

// GetIpv6PrefixRanges returns the Ipv6PrefixRanges field value if set, zero value otherwise.
func (o *BsfInfo) GetIpv6PrefixRanges() []Ipv6PrefixRange {
	if o == nil || o.Ipv6PrefixRanges == nil {
		var ret []Ipv6PrefixRange
		return ret
	}
	return o.Ipv6PrefixRanges
}

// GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetIpv6PrefixRangesOk() ([]Ipv6PrefixRange, bool) {
	if o == nil || o.Ipv6PrefixRanges == nil {
		return nil, false
	}
	return o.Ipv6PrefixRanges, true
}

// HasIpv6PrefixRanges returns a boolean if a field has been set.
func (o *BsfInfo) HasIpv6PrefixRanges() bool {
	if o != nil && o.Ipv6PrefixRanges != nil {
		return true
	}

	return false
}

// SetIpv6PrefixRanges gets a reference to the given []Ipv6PrefixRange and assigns it to the Ipv6PrefixRanges field.
func (o *BsfInfo) SetIpv6PrefixRanges(v []Ipv6PrefixRange) {
	o.Ipv6PrefixRanges = v
}

func (o BsfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DnnList != nil {
		toSerialize["dnnList"] = o.DnnList
	}
	if o.IpDomainList != nil {
		toSerialize["ipDomainList"] = o.IpDomainList
	}
	if o.Ipv4AddressRanges != nil {
		toSerialize["ipv4AddressRanges"] = o.Ipv4AddressRanges
	}
	if o.Ipv6PrefixRanges != nil {
		toSerialize["ipv6PrefixRanges"] = o.Ipv6PrefixRanges
	}
	return json.Marshal(toSerialize)
}

type NullableBsfInfo struct {
	value *BsfInfo
	isSet bool
}

func (v NullableBsfInfo) Get() *BsfInfo {
	return v.value
}

func (v *NullableBsfInfo) Set(val *BsfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBsfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBsfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsfInfo(val *BsfInfo) *NullableBsfInfo {
	return &NullableBsfInfo{value: val, isSet: true}
}

func (v NullableBsfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
