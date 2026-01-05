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
CAPIF_Publish_Service_API

API for publishing service APIs.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.5
*/

package publish_api_service

import (
	"encoding/json"
)

// checks if the IpAddrRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpAddrRange{}

// IpAddrRange Represents the list of public IP ranges
type IpAddrRange struct {
	// Represents the IPv4 Address ranges of the UE(s).
	UeIpv4AddrRanges []Ipv4AddressRange
	// Represents the Ipv6 Address ranges of the UE(s).
	UeIpv6AddrRanges []Ipv6AddressRange
}

// NewIpAddrRange instantiates a new IpAddrRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpAddrRange() *IpAddrRange {
	this := IpAddrRange{}
	return &this
}

// NewIpAddrRangeWithDefaults instantiates a new IpAddrRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpAddrRangeWithDefaults() *IpAddrRange {
	this := IpAddrRange{}
	return &this
}

// GetUeIpv4AddrRanges returns the UeIpv4AddrRanges field value if set, zero value otherwise.
func (o *IpAddrRange) GetUeIpv4AddrRanges() []Ipv4AddressRange {
	if o == nil || IsNil(o.UeIpv4AddrRanges) {
		var ret []Ipv4AddressRange
		return ret
	}
	return o.UeIpv4AddrRanges
}

// GetUeIpv4AddrRangesOk returns a tuple with the UeIpv4AddrRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddrRange) GetUeIpv4AddrRangesOk() ([]Ipv4AddressRange, bool) {
	if o == nil || IsNil(o.UeIpv4AddrRanges) {
		return nil, false
	}
	return o.UeIpv4AddrRanges, true
}

// HasUeIpv4AddrRanges returns a boolean if a field has been set.
func (o *IpAddrRange) HasUeIpv4AddrRanges() bool {
	if o != nil && !IsNil(o.UeIpv4AddrRanges) {
		return true
	}

	return false
}

// SetUeIpv4AddrRanges gets a reference to the given []Ipv4AddressRange and assigns it to the UeIpv4AddrRanges field.
func (o *IpAddrRange) SetUeIpv4AddrRanges(v []Ipv4AddressRange) {
	o.UeIpv4AddrRanges = v
}

// GetUeIpv6AddrRanges returns the UeIpv6AddrRanges field value if set, zero value otherwise.
func (o *IpAddrRange) GetUeIpv6AddrRanges() []Ipv6AddressRange {
	if o == nil || IsNil(o.UeIpv6AddrRanges) {
		var ret []Ipv6AddressRange
		return ret
	}
	return o.UeIpv6AddrRanges
}

// GetUeIpv6AddrRangesOk returns a tuple with the UeIpv6AddrRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAddrRange) GetUeIpv6AddrRangesOk() ([]Ipv6AddressRange, bool) {
	if o == nil || IsNil(o.UeIpv6AddrRanges) {
		return nil, false
	}
	return o.UeIpv6AddrRanges, true
}

// HasUeIpv6AddrRanges returns a boolean if a field has been set.
func (o *IpAddrRange) HasUeIpv6AddrRanges() bool {
	if o != nil && !IsNil(o.UeIpv6AddrRanges) {
		return true
	}

	return false
}

// SetUeIpv6AddrRanges gets a reference to the given []Ipv6AddressRange and assigns it to the UeIpv6AddrRanges field.
func (o *IpAddrRange) SetUeIpv6AddrRanges(v []Ipv6AddressRange) {
	o.UeIpv6AddrRanges = v
}

func (o IpAddrRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpAddrRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UeIpv4AddrRanges) {
		toSerialize["ueIpv4AddrRanges"] = o.UeIpv4AddrRanges
	}
	if !IsNil(o.UeIpv6AddrRanges) {
		toSerialize["ueIpv6AddrRanges"] = o.UeIpv6AddrRanges
	}
	return toSerialize, nil
}

type NullableIpAddrRange struct {
	value *IpAddrRange
	isSet bool
}

func (v NullableIpAddrRange) Get() *IpAddrRange {
	return v.value
}

func (v *NullableIpAddrRange) Set(val *IpAddrRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddrRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddrRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddrRange(val *IpAddrRange) *NullableIpAddrRange {
	return &NullableIpAddrRange{value: val, isSet: true}
}

func (v NullableIpAddrRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddrRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
