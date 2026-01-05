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

// AusfInfo struct for AusfInfo
type AusfInfo struct {
	GroupId           *string     `json:"groupId,omitempty"`
	SupiRanges        []SupiRange `json:"supiRanges,omitempty"`
	RoutingIndicators []string    `json:"routingIndicators,omitempty"`
}

// NewAusfInfo instantiates a new AusfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAusfInfo() *AusfInfo {
	this := AusfInfo{}
	return &this
}

// NewAusfInfoWithDefaults instantiates a new AusfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAusfInfoWithDefaults() *AusfInfo {
	this := AusfInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AusfInfo) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AusfInfo) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AusfInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *AusfInfo) GetSupiRanges() []SupiRange {
	if o == nil || o.SupiRanges == nil {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || o.SupiRanges == nil {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *AusfInfo) HasSupiRanges() bool {
	if o != nil && o.SupiRanges != nil {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *AusfInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetRoutingIndicators returns the RoutingIndicators field value if set, zero value otherwise.
func (o *AusfInfo) GetRoutingIndicators() []string {
	if o == nil || o.RoutingIndicators == nil {
		var ret []string
		return ret
	}
	return o.RoutingIndicators
}

// GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfInfo) GetRoutingIndicatorsOk() ([]string, bool) {
	if o == nil || o.RoutingIndicators == nil {
		return nil, false
	}
	return o.RoutingIndicators, true
}

// HasRoutingIndicators returns a boolean if a field has been set.
func (o *AusfInfo) HasRoutingIndicators() bool {
	if o != nil && o.RoutingIndicators != nil {
		return true
	}

	return false
}

// SetRoutingIndicators gets a reference to the given []string and assigns it to the RoutingIndicators field.
func (o *AusfInfo) SetRoutingIndicators(v []string) {
	o.RoutingIndicators = v
}

func (o AusfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.SupiRanges != nil {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if o.RoutingIndicators != nil {
		toSerialize["routingIndicators"] = o.RoutingIndicators
	}
	return json.Marshal(toSerialize)
}

type NullableAusfInfo struct {
	value *AusfInfo
	isSet bool
}

func (v NullableAusfInfo) Get() *AusfInfo {
	return v.value
}

func (v *NullableAusfInfo) Set(val *AusfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAusfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAusfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAusfInfo(val *AusfInfo) *NullableAusfInfo {
	return &NullableAusfInfo{value: val, isSet: true}
}

func (v NullableAusfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAusfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
