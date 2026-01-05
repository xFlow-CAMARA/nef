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

// UdmInfo struct for UdmInfo
type UdmInfo struct {
	GroupId                        *string         `json:"groupId,omitempty"`
	SupiRanges                     []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges                     []IdentityRange `json:"gpsiRanges,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`
	RoutingIndicators              []string        `json:"routingIndicators,omitempty"`
}

// NewUdmInfo instantiates a new UdmInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdmInfo() *UdmInfo {
	this := UdmInfo{}
	return &this
}

// NewUdmInfoWithDefaults instantiates a new UdmInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdmInfoWithDefaults() *UdmInfo {
	this := UdmInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *UdmInfo) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *UdmInfo) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *UdmInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *UdmInfo) GetSupiRanges() []SupiRange {
	if o == nil || o.SupiRanges == nil {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || o.SupiRanges == nil {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *UdmInfo) HasSupiRanges() bool {
	if o != nil && o.SupiRanges != nil {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *UdmInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *UdmInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || o.GpsiRanges == nil {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || o.GpsiRanges == nil {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *UdmInfo) HasGpsiRanges() bool {
	if o != nil && o.GpsiRanges != nil {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *UdmInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *UdmInfo) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || o.ExternalGroupIdentifiersRanges == nil {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || o.ExternalGroupIdentifiersRanges == nil {
		return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *UdmInfo) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && o.ExternalGroupIdentifiersRanges != nil {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *UdmInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetRoutingIndicators returns the RoutingIndicators field value if set, zero value otherwise.
func (o *UdmInfo) GetRoutingIndicators() []string {
	if o == nil || o.RoutingIndicators == nil {
		var ret []string
		return ret
	}
	return o.RoutingIndicators
}

// GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetRoutingIndicatorsOk() ([]string, bool) {
	if o == nil || o.RoutingIndicators == nil {
		return nil, false
	}
	return o.RoutingIndicators, true
}

// HasRoutingIndicators returns a boolean if a field has been set.
func (o *UdmInfo) HasRoutingIndicators() bool {
	if o != nil && o.RoutingIndicators != nil {
		return true
	}

	return false
}

// SetRoutingIndicators gets a reference to the given []string and assigns it to the RoutingIndicators field.
func (o *UdmInfo) SetRoutingIndicators(v []string) {
	o.RoutingIndicators = v
}

func (o UdmInfo) MarshalJSON() ([]byte, error) {
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
	if o.RoutingIndicators != nil {
		toSerialize["routingIndicators"] = o.RoutingIndicators
	}
	return json.Marshal(toSerialize)
}

type NullableUdmInfo struct {
	value *UdmInfo
	isSet bool
}

func (v NullableUdmInfo) Get() *UdmInfo {
	return v.value
}

func (v *NullableUdmInfo) Set(val *UdmInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUdmInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUdmInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdmInfo(val *UdmInfo) *NullableUdmInfo {
	return &NullableUdmInfo{value: val, isSet: true}
}

func (v NullableUdmInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdmInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
