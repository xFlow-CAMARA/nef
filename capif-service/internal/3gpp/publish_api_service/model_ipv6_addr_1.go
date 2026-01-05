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

// checks if the Ipv6Addr1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv6Addr1{}

// Ipv6Addr1 String identifying an IPv6 address formatted according to clause 4 of RFC5952. The mixed IPv4 IPv6 notation according to clause 5 of RFC5952 shall not be used.
type Ipv6Addr1 struct {
}

// NewIpv6Addr1 instantiates a new Ipv6Addr1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6Addr1() *Ipv6Addr1 {
	this := Ipv6Addr1{}
	return &this
}

// NewIpv6Addr1WithDefaults instantiates a new Ipv6Addr1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6Addr1WithDefaults() *Ipv6Addr1 {
	this := Ipv6Addr1{}
	return &this
}

func (o Ipv6Addr1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv6Addr1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableIpv6Addr1 struct {
	value *Ipv6Addr1
	isSet bool
}

func (v NullableIpv6Addr1) Get() *Ipv6Addr1 {
	return v.value
}

func (v *NullableIpv6Addr1) Set(val *Ipv6Addr1) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6Addr1) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6Addr1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6Addr1(val *Ipv6Addr1) *NullableIpv6Addr1 {
	return &NullableIpv6Addr1{value: val, isSet: true}
}

func (v NullableIpv6Addr1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6Addr1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
