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

// checks if the ServiceKpis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceKpis{}

// ServiceKpis Represents information about the service characteristics provided by a service API.
type ServiceKpis struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReqRate *int32 `json:"maxReqRate,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	MaxRestime *int32 `json:"maxRestime,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Availability *int32 `json:"availability,omitempty"`
	// The maximum compute resource available in FLOPS for the API Invoker.
	AvalComp *string `json:"avalComp,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (kFLOPS|MFLOPS|GFLOPS|TFLOPS|PFLOPS|EFLOPS|ZFLOPS)$"`
	// The maximum graphical compute resource in FLOPS available for the API Invoker.
	AvalGraComp *string `json:"avalGraComp,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (kFLOPS|MFLOPS|GFLOPS|TFLOPS|PFLOPS|EFLOPS|ZFLOPS)$"`
	// The maximum memory resource available for the API Invoker.
	AvalMem *string `json:"avalMem,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (KB|MB|GB|TB|PB|EB|ZB|YB)$"`
	// The maximum storage resource available for the API Invoker.
	AvalStor *string `json:"avalStor,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (KB|MB|GB|TB|PB|EB|ZB|YB)$"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ConBand *int32 `json:"conBand,omitempty"`
}

// NewServiceKpis instantiates a new ServiceKpis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceKpis() *ServiceKpis {
	this := ServiceKpis{}
	return &this
}

// NewServiceKpisWithDefaults instantiates a new ServiceKpis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceKpisWithDefaults() *ServiceKpis {
	this := ServiceKpis{}
	return &this
}

// GetMaxReqRate returns the MaxReqRate field value if set, zero value otherwise.
func (o *ServiceKpis) GetMaxReqRate() int32 {
	if o == nil || IsNil(o.MaxReqRate) {
		var ret int32
		return ret
	}
	return *o.MaxReqRate
}

// GetMaxReqRateOk returns a tuple with the MaxReqRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKpis) GetMaxReqRateOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxReqRate) {
		return nil, false
	}
	return o.MaxReqRate, true
}

// HasMaxReqRate returns a boolean if a field has been set.
func (o *ServiceKpis) HasMaxReqRate() bool {
	if o != nil && !IsNil(o.MaxReqRate) {
		return true
	}

	return false
}

// SetMaxReqRate gets a reference to the given int32 and assigns it to the MaxReqRate field.
func (o *ServiceKpis) SetMaxReqRate(v int32) {
	o.MaxReqRate = &v
}

// GetMaxRestime returns the MaxRestime field value if set, zero value otherwise.
func (o *ServiceKpis) GetMaxRestime() int32 {
	if o == nil || IsNil(o.MaxRestime) {
		var ret int32
		return ret
	}
	return *o.MaxRestime
}

// GetMaxRestimeOk returns a tuple with the MaxRestime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKpis) GetMaxRestimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRestime) {
		return nil, false
	}
	return o.MaxRestime, true
}

// HasMaxRestime returns a boolean if a field has been set.
func (o *ServiceKpis) HasMaxRestime() bool {
	if o != nil && !IsNil(o.MaxRestime) {
		return true
	}

	return false
}

// SetMaxRestime gets a reference to the given int32 and assigns it to the MaxRestime field.
func (o *ServiceKpis) SetMaxRestime(v int32) {
	o.MaxRestime = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *ServiceKpis) GetAvailability() int32 {
	if o == nil || IsNil(o.Availability) {
		var ret int32
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKpis) GetAvailabilityOk() (*int32, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *ServiceKpis) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given int32 and assigns it to the Availability field.
func (o *ServiceKpis) SetAvailability(v int32) {
	o.Availability = &v
}

// GetAvalComp returns the AvalComp field value if set, zero value otherwise.
func (o *ServiceKpis) GetAvalComp() string {
	if o == nil || IsNil(o.AvalComp) {
		var ret string
		return ret
	}
	return *o.AvalComp
}

// GetAvalCompOk returns a tuple with the AvalComp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKpis) GetAvalCompOk() (*string, bool) {
	if o == nil || IsNil(o.AvalComp) {
		return nil, false
	}
	return o.AvalComp, true
}

// HasAvalComp returns a boolean if a field has been set.
func (o *ServiceKpis) HasAvalComp() bool {
	if o != nil && !IsNil(o.AvalComp) {
		return true
	}

	return false
}

// SetAvalComp gets a reference to the given string and assigns it to the AvalComp field.
func (o *ServiceKpis) SetAvalComp(v string) {
	o.AvalComp = &v
}

// GetAvalGraComp returns the AvalGraComp field value if set, zero value otherwise.
func (o *ServiceKpis) GetAvalGraComp() string {
	if o == nil || IsNil(o.AvalGraComp) {
		var ret string
		return ret
	}
	return *o.AvalGraComp
}

// GetAvalGraCompOk returns a tuple with the AvalGraComp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKpis) GetAvalGraCompOk() (*string, bool) {
	if o == nil || IsNil(o.AvalGraComp) {
		return nil, false
	}
	return o.AvalGraComp, true
}

// HasAvalGraComp returns a boolean if a field has been set.
func (o *ServiceKpis) HasAvalGraComp() bool {
	if o != nil && !IsNil(o.AvalGraComp) {
		return true
	}

	return false
}

// SetAvalGraComp gets a reference to the given string and assigns it to the AvalGraComp field.
func (o *ServiceKpis) SetAvalGraComp(v string) {
	o.AvalGraComp = &v
}

// GetAvalMem returns the AvalMem field value if set, zero value otherwise.
func (o *ServiceKpis) GetAvalMem() string {
	if o == nil || IsNil(o.AvalMem) {
		var ret string
		return ret
	}
	return *o.AvalMem
}

// GetAvalMemOk returns a tuple with the AvalMem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKpis) GetAvalMemOk() (*string, bool) {
	if o == nil || IsNil(o.AvalMem) {
		return nil, false
	}
	return o.AvalMem, true
}

// HasAvalMem returns a boolean if a field has been set.
func (o *ServiceKpis) HasAvalMem() bool {
	if o != nil && !IsNil(o.AvalMem) {
		return true
	}

	return false
}

// SetAvalMem gets a reference to the given string and assigns it to the AvalMem field.
func (o *ServiceKpis) SetAvalMem(v string) {
	o.AvalMem = &v
}

// GetAvalStor returns the AvalStor field value if set, zero value otherwise.
func (o *ServiceKpis) GetAvalStor() string {
	if o == nil || IsNil(o.AvalStor) {
		var ret string
		return ret
	}
	return *o.AvalStor
}

// GetAvalStorOk returns a tuple with the AvalStor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKpis) GetAvalStorOk() (*string, bool) {
	if o == nil || IsNil(o.AvalStor) {
		return nil, false
	}
	return o.AvalStor, true
}

// HasAvalStor returns a boolean if a field has been set.
func (o *ServiceKpis) HasAvalStor() bool {
	if o != nil && !IsNil(o.AvalStor) {
		return true
	}

	return false
}

// SetAvalStor gets a reference to the given string and assigns it to the AvalStor field.
func (o *ServiceKpis) SetAvalStor(v string) {
	o.AvalStor = &v
}

// GetConBand returns the ConBand field value if set, zero value otherwise.
func (o *ServiceKpis) GetConBand() int32 {
	if o == nil || IsNil(o.ConBand) {
		var ret int32
		return ret
	}
	return *o.ConBand
}

// GetConBandOk returns a tuple with the ConBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKpis) GetConBandOk() (*int32, bool) {
	if o == nil || IsNil(o.ConBand) {
		return nil, false
	}
	return o.ConBand, true
}

// HasConBand returns a boolean if a field has been set.
func (o *ServiceKpis) HasConBand() bool {
	if o != nil && !IsNil(o.ConBand) {
		return true
	}

	return false
}

// SetConBand gets a reference to the given int32 and assigns it to the ConBand field.
func (o *ServiceKpis) SetConBand(v int32) {
	o.ConBand = &v
}

func (o ServiceKpis) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceKpis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxReqRate) {
		toSerialize["maxReqRate"] = o.MaxReqRate
	}
	if !IsNil(o.MaxRestime) {
		toSerialize["maxRestime"] = o.MaxRestime
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.AvalComp) {
		toSerialize["avalComp"] = o.AvalComp
	}
	if !IsNil(o.AvalGraComp) {
		toSerialize["avalGraComp"] = o.AvalGraComp
	}
	if !IsNil(o.AvalMem) {
		toSerialize["avalMem"] = o.AvalMem
	}
	if !IsNil(o.AvalStor) {
		toSerialize["avalStor"] = o.AvalStor
	}
	if !IsNil(o.ConBand) {
		toSerialize["conBand"] = o.ConBand
	}
	return toSerialize, nil
}

type NullableServiceKpis struct {
	value *ServiceKpis
	isSet bool
}

func (v NullableServiceKpis) Get() *ServiceKpis {
	return v.value
}

func (v *NullableServiceKpis) Set(val *ServiceKpis) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceKpis) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceKpis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceKpis(val *ServiceKpis) *NullableServiceKpis {
	return &NullableServiceKpis{value: val, isSet: true}
}

func (v NullableServiceKpis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceKpis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
