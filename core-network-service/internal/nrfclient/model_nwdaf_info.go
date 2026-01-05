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

// NwdafInfo struct for NwdafInfo
type NwdafInfo struct {
	EventIds     []EventId    `json:"eventIds,omitempty"`
	NwdafEvents  []NwdafEvent `json:"nwdafEvents,omitempty"`
	TaiList      []Tai        `json:"taiList,omitempty"`
	TaiRangeList []TaiRange   `json:"taiRangeList,omitempty"`
}

// NewNwdafInfo instantiates a new NwdafInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafInfo() *NwdafInfo {
	this := NwdafInfo{}
	return &this
}

// NewNwdafInfoWithDefaults instantiates a new NwdafInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafInfoWithDefaults() *NwdafInfo {
	this := NwdafInfo{}
	return &this
}

// GetEventIds returns the EventIds field value if set, zero value otherwise.
func (o *NwdafInfo) GetEventIds() []EventId {
	if o == nil || o.EventIds == nil {
		var ret []EventId
		return ret
	}
	return o.EventIds
}

// GetEventIdsOk returns a tuple with the EventIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetEventIdsOk() ([]EventId, bool) {
	if o == nil || o.EventIds == nil {
		return nil, false
	}
	return o.EventIds, true
}

// HasEventIds returns a boolean if a field has been set.
func (o *NwdafInfo) HasEventIds() bool {
	if o != nil && o.EventIds != nil {
		return true
	}

	return false
}

// SetEventIds gets a reference to the given []EventId and assigns it to the EventIds field.
func (o *NwdafInfo) SetEventIds(v []EventId) {
	o.EventIds = v
}

// GetNwdafEvents returns the NwdafEvents field value if set, zero value otherwise.
func (o *NwdafInfo) GetNwdafEvents() []NwdafEvent {
	if o == nil || o.NwdafEvents == nil {
		var ret []NwdafEvent
		return ret
	}
	return o.NwdafEvents
}

// GetNwdafEventsOk returns a tuple with the NwdafEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetNwdafEventsOk() ([]NwdafEvent, bool) {
	if o == nil || o.NwdafEvents == nil {
		return nil, false
	}
	return o.NwdafEvents, true
}

// HasNwdafEvents returns a boolean if a field has been set.
func (o *NwdafInfo) HasNwdafEvents() bool {
	if o != nil && o.NwdafEvents != nil {
		return true
	}

	return false
}

// SetNwdafEvents gets a reference to the given []NwdafEvent and assigns it to the NwdafEvents field.
func (o *NwdafInfo) SetNwdafEvents(v []NwdafEvent) {
	o.NwdafEvents = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *NwdafInfo) GetTaiList() []Tai {
	if o == nil || o.TaiList == nil {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || o.TaiList == nil {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *NwdafInfo) HasTaiList() bool {
	if o != nil && o.TaiList != nil {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *NwdafInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NwdafInfo) GetTaiRangeList() []TaiRange {
	if o == nil || o.TaiRangeList == nil {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || o.TaiRangeList == nil {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NwdafInfo) HasTaiRangeList() bool {
	if o != nil && o.TaiRangeList != nil {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NwdafInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

func (o NwdafInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventIds != nil {
		toSerialize["eventIds"] = o.EventIds
	}
	if o.NwdafEvents != nil {
		toSerialize["nwdafEvents"] = o.NwdafEvents
	}
	if o.TaiList != nil {
		toSerialize["taiList"] = o.TaiList
	}
	if o.TaiRangeList != nil {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	return json.Marshal(toSerialize)
}

type NullableNwdafInfo struct {
	value *NwdafInfo
	isSet bool
}

func (v NullableNwdafInfo) Get() *NwdafInfo {
	return v.value
}

func (v *NullableNwdafInfo) Set(val *NwdafInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafInfo(val *NwdafInfo) *NullableNwdafInfo {
	return &NullableNwdafInfo{value: val, isSet: true}
}

func (v NullableNwdafInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
