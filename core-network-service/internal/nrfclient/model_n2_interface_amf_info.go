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

// N2InterfaceAmfInfo struct for N2InterfaceAmfInfo
type N2InterfaceAmfInfo struct {
	Ipv4EndpointAddress []string   `json:"ipv4EndpointAddress,omitempty"`
	Ipv6EndpointAddress []Ipv6Addr `json:"ipv6EndpointAddress,omitempty"`
	AmfName             *string    `json:"amfName,omitempty"`
}

// NewN2InterfaceAmfInfo instantiates a new N2InterfaceAmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN2InterfaceAmfInfo() *N2InterfaceAmfInfo {
	this := N2InterfaceAmfInfo{}
	return &this
}

// NewN2InterfaceAmfInfoWithDefaults instantiates a new N2InterfaceAmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN2InterfaceAmfInfoWithDefaults() *N2InterfaceAmfInfo {
	this := N2InterfaceAmfInfo{}
	return &this
}

// GetIpv4EndpointAddress returns the Ipv4EndpointAddress field value if set, zero value otherwise.
func (o *N2InterfaceAmfInfo) GetIpv4EndpointAddress() []string {
	if o == nil || o.Ipv4EndpointAddress == nil {
		var ret []string
		return ret
	}
	return o.Ipv4EndpointAddress
}

// GetIpv4EndpointAddressOk returns a tuple with the Ipv4EndpointAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InterfaceAmfInfo) GetIpv4EndpointAddressOk() ([]string, bool) {
	if o == nil || o.Ipv4EndpointAddress == nil {
		return nil, false
	}
	return o.Ipv4EndpointAddress, true
}

// HasIpv4EndpointAddress returns a boolean if a field has been set.
func (o *N2InterfaceAmfInfo) HasIpv4EndpointAddress() bool {
	if o != nil && o.Ipv4EndpointAddress != nil {
		return true
	}

	return false
}

// SetIpv4EndpointAddress gets a reference to the given []string and assigns it to the Ipv4EndpointAddress field.
func (o *N2InterfaceAmfInfo) SetIpv4EndpointAddress(v []string) {
	o.Ipv4EndpointAddress = v
}

// GetIpv6EndpointAddress returns the Ipv6EndpointAddress field value if set, zero value otherwise.
func (o *N2InterfaceAmfInfo) GetIpv6EndpointAddress() []Ipv6Addr {
	if o == nil || o.Ipv6EndpointAddress == nil {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6EndpointAddress
}

// GetIpv6EndpointAddressOk returns a tuple with the Ipv6EndpointAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InterfaceAmfInfo) GetIpv6EndpointAddressOk() ([]Ipv6Addr, bool) {
	if o == nil || o.Ipv6EndpointAddress == nil {
		return nil, false
	}
	return o.Ipv6EndpointAddress, true
}

// HasIpv6EndpointAddress returns a boolean if a field has been set.
func (o *N2InterfaceAmfInfo) HasIpv6EndpointAddress() bool {
	if o != nil && o.Ipv6EndpointAddress != nil {
		return true
	}

	return false
}

// SetIpv6EndpointAddress gets a reference to the given []Ipv6Addr and assigns it to the Ipv6EndpointAddress field.
func (o *N2InterfaceAmfInfo) SetIpv6EndpointAddress(v []Ipv6Addr) {
	o.Ipv6EndpointAddress = v
}

// GetAmfName returns the AmfName field value if set, zero value otherwise.
func (o *N2InterfaceAmfInfo) GetAmfName() string {
	if o == nil || o.AmfName == nil {
		var ret string
		return ret
	}
	return *o.AmfName
}

// GetAmfNameOk returns a tuple with the AmfName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InterfaceAmfInfo) GetAmfNameOk() (*string, bool) {
	if o == nil || o.AmfName == nil {
		return nil, false
	}
	return o.AmfName, true
}

// HasAmfName returns a boolean if a field has been set.
func (o *N2InterfaceAmfInfo) HasAmfName() bool {
	if o != nil && o.AmfName != nil {
		return true
	}

	return false
}

// SetAmfName gets a reference to the given string and assigns it to the AmfName field.
func (o *N2InterfaceAmfInfo) SetAmfName(v string) {
	o.AmfName = &v
}

func (o N2InterfaceAmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ipv4EndpointAddress != nil {
		toSerialize["ipv4EndpointAddress"] = o.Ipv4EndpointAddress
	}
	if o.Ipv6EndpointAddress != nil {
		toSerialize["ipv6EndpointAddress"] = o.Ipv6EndpointAddress
	}
	if o.AmfName != nil {
		toSerialize["amfName"] = o.AmfName
	}
	return json.Marshal(toSerialize)
}

type NullableN2InterfaceAmfInfo struct {
	value *N2InterfaceAmfInfo
	isSet bool
}

func (v NullableN2InterfaceAmfInfo) Get() *N2InterfaceAmfInfo {
	return v.value
}

func (v *NullableN2InterfaceAmfInfo) Set(val *N2InterfaceAmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InterfaceAmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InterfaceAmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InterfaceAmfInfo(val *N2InterfaceAmfInfo) *NullableN2InterfaceAmfInfo {
	return &NullableN2InterfaceAmfInfo{value: val, isSet: true}
}

func (v NullableN2InterfaceAmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InterfaceAmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
