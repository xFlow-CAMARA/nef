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

// Ipv4AddressRange struct for Ipv4AddressRange
type Ipv4AddressRange struct {
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

// NewIpv4AddressRange instantiates a new Ipv4AddressRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv4AddressRange() *Ipv4AddressRange {
	this := Ipv4AddressRange{}
	return &this
}

// NewIpv4AddressRangeWithDefaults instantiates a new Ipv4AddressRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv4AddressRangeWithDefaults() *Ipv4AddressRange {
	this := Ipv4AddressRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Ipv4AddressRange) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4AddressRange) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Ipv4AddressRange) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *Ipv4AddressRange) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Ipv4AddressRange) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4AddressRange) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Ipv4AddressRange) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *Ipv4AddressRange) SetEnd(v string) {
	o.End = &v
}

func (o Ipv4AddressRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableIpv4AddressRange struct {
	value *Ipv4AddressRange
	isSet bool
}

func (v NullableIpv4AddressRange) Get() *Ipv4AddressRange {
	return v.value
}

func (v *NullableIpv4AddressRange) Set(val *Ipv4AddressRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv4AddressRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv4AddressRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv4AddressRange(val *Ipv4AddressRange) *NullableIpv4AddressRange {
	return &NullableIpv4AddressRange{value: val, isSet: true}
}

func (v NullableIpv4AddressRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv4AddressRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
