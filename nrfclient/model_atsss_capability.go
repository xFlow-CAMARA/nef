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

// AtsssCapability struct for AtsssCapability
type AtsssCapability struct {
	AtsssLL *bool `json:"atsssLL,omitempty"`
	Mptcp   *bool `json:"mptcp,omitempty"`
}

// NewAtsssCapability instantiates a new AtsssCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtsssCapability() *AtsssCapability {
	this := AtsssCapability{}
	var atsssLL bool = false
	this.AtsssLL = &atsssLL
	var mptcp bool = false
	this.Mptcp = &mptcp
	return &this
}

// NewAtsssCapabilityWithDefaults instantiates a new AtsssCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtsssCapabilityWithDefaults() *AtsssCapability {
	this := AtsssCapability{}
	var atsssLL bool = false
	this.AtsssLL = &atsssLL
	var mptcp bool = false
	this.Mptcp = &mptcp
	return &this
}

// GetAtsssLL returns the AtsssLL field value if set, zero value otherwise.
func (o *AtsssCapability) GetAtsssLL() bool {
	if o == nil || o.AtsssLL == nil {
		var ret bool
		return ret
	}
	return *o.AtsssLL
}

// GetAtsssLLOk returns a tuple with the AtsssLL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsssCapability) GetAtsssLLOk() (*bool, bool) {
	if o == nil || o.AtsssLL == nil {
		return nil, false
	}
	return o.AtsssLL, true
}

// HasAtsssLL returns a boolean if a field has been set.
func (o *AtsssCapability) HasAtsssLL() bool {
	if o != nil && o.AtsssLL != nil {
		return true
	}

	return false
}

// SetAtsssLL gets a reference to the given bool and assigns it to the AtsssLL field.
func (o *AtsssCapability) SetAtsssLL(v bool) {
	o.AtsssLL = &v
}

// GetMptcp returns the Mptcp field value if set, zero value otherwise.
func (o *AtsssCapability) GetMptcp() bool {
	if o == nil || o.Mptcp == nil {
		var ret bool
		return ret
	}
	return *o.Mptcp
}

// GetMptcpOk returns a tuple with the Mptcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtsssCapability) GetMptcpOk() (*bool, bool) {
	if o == nil || o.Mptcp == nil {
		return nil, false
	}
	return o.Mptcp, true
}

// HasMptcp returns a boolean if a field has been set.
func (o *AtsssCapability) HasMptcp() bool {
	if o != nil && o.Mptcp != nil {
		return true
	}

	return false
}

// SetMptcp gets a reference to the given bool and assigns it to the Mptcp field.
func (o *AtsssCapability) SetMptcp(v bool) {
	o.Mptcp = &v
}

func (o AtsssCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AtsssLL != nil {
		toSerialize["atsssLL"] = o.AtsssLL
	}
	if o.Mptcp != nil {
		toSerialize["mptcp"] = o.Mptcp
	}
	return json.Marshal(toSerialize)
}

type NullableAtsssCapability struct {
	value *AtsssCapability
	isSet bool
}

func (v NullableAtsssCapability) Get() *AtsssCapability {
	return v.value
}

func (v *NullableAtsssCapability) Set(val *AtsssCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableAtsssCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableAtsssCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtsssCapability(val *AtsssCapability) *NullableAtsssCapability {
	return &NullableAtsssCapability{value: val, isSet: true}
}

func (v NullableAtsssCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtsssCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
