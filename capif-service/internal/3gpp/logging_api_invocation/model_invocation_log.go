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
CAPIF_Logging_API_Invocation_API

API for invocation logs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

package logging_api_invocation

import (
	"bytes"
	"encoding/json"
)

// checks if the InvocationLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvocationLog{}

// InvocationLog Represents a set of Service API invocation logs to be stored in a CAPIF core function.
type InvocationLog struct {
	// Identity information of the API exposing function requesting logging of service API invocations
	AefId string `json:"aefId"`
	// Identity of the API invoker which invoked the service API
	ApiInvokerId string `json:"apiInvokerId"`
	// Service API invocation log
	Logs []Log `json:"logs"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`
}

type _InvocationLog InvocationLog

// NewInvocationLog instantiates a new InvocationLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvocationLog(aefId string, apiInvokerId string, logs []Log) *InvocationLog {
	this := InvocationLog{}
	this.AefId = aefId
	this.ApiInvokerId = apiInvokerId
	this.Logs = logs
	return &this
}

// NewInvocationLogWithDefaults instantiates a new InvocationLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvocationLogWithDefaults() *InvocationLog {
	this := InvocationLog{}
	return &this
}

// GetAefId returns the AefId field value
func (o *InvocationLog) GetAefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AefId
}

// GetAefIdOk returns a tuple with the AefId field value
// and a boolean to check if the value has been set.
func (o *InvocationLog) GetAefIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AefId, true
}

// SetAefId sets field value
func (o *InvocationLog) SetAefId(v string) {
	o.AefId = v
}

// GetApiInvokerId returns the ApiInvokerId field value
func (o *InvocationLog) GetApiInvokerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiInvokerId
}

// GetApiInvokerIdOk returns a tuple with the ApiInvokerId field value
// and a boolean to check if the value has been set.
func (o *InvocationLog) GetApiInvokerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiInvokerId, true
}

// SetApiInvokerId sets field value
func (o *InvocationLog) SetApiInvokerId(v string) {
	o.ApiInvokerId = v
}

// GetLogs returns the Logs field value
func (o *InvocationLog) GetLogs() []Log {
	if o == nil {
		var ret []Log
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *InvocationLog) GetLogsOk() ([]Log, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *InvocationLog) SetLogs(v []Log) {
	o.Logs = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *InvocationLog) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvocationLog) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *InvocationLog) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *InvocationLog) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o InvocationLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvocationLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aefId"] = o.AefId
	toSerialize["apiInvokerId"] = o.ApiInvokerId
	toSerialize["logs"] = o.Logs
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

func (o *InvocationLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aefId",
		"apiInvokerId",
		"logs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return nil
		}
	}

	varInvocationLog := _InvocationLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvocationLog)

	if err != nil {
		return err
	}

	*o = InvocationLog(varInvocationLog)

	return err
}

type NullableInvocationLog struct {
	value *InvocationLog
	isSet bool
}

func (v NullableInvocationLog) Get() *InvocationLog {
	return v.value
}

func (v *NullableInvocationLog) Set(val *InvocationLog) {
	v.value = val
	v.isSet = true
}

func (v NullableInvocationLog) IsSet() bool {
	return v.isSet
}

func (v *NullableInvocationLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvocationLog(val *InvocationLog) *NullableInvocationLog {
	return &NullableInvocationLog{value: val, isSet: true}
}

func (v NullableInvocationLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvocationLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
