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
	"time"
)

// checks if the Log type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Log{}

// Log Represents an individual service API invocation log entry.
type Log struct {
	// String identifying the API invoked.
	ApiId string `json:"apiId"`
	// Name of the API which was invoked, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122.
	ApiName string `json:"apiName"`
	// Version of the API which was invoked
	ApiVersion string `json:"apiVersion"`
	// Name of the specific resource invoked
	ResourceName string `json:"resourceName"`
	// string providing an URI formatted according to IETF RFC 3986.
	Uri       *string `json:"uri,omitempty"`
	Protocol  string  `json:"protocol"`
	Operation string  `json:"operation,omitempty"`
	// For HTTP protocol, it contains HTTP status code of the invocation
	Result string `json:"result"`
	// string with format \"date-time\" as defined in OpenAPI.
	InvocationTime *time.Time `json:"invocationTime,omitempty"`
	// Represents a period of time in units of milliseconds.
	InvocationLatency *int32 `json:"invocationLatency,omitempty"`
	// List of input parameters. Can be any value - string, number, boolean, array or object.
	InputParameters interface{} `json:"inputParameters,omitempty"`
	// List of output parameters. Can be any value - string, number, boolean, array or object.
	OutputParameters interface{}                  `json:"outputParameters,omitempty"`
	SrcInterface     NullableInterfaceDescription `json:"srcInterface,omitempty"`
	DestInterface    NullableInterfaceDescription `json:"destInterface,omitempty"`
	// It includes the node identifier (as defined in IETF RFC 7239 of all forwarding entities between the API invoker and the AEF, concatenated with comma and space, e.g. 192.0.2.43:80, unknown:_OBFport, 203.0.113.60
	FwdInterface *string `json:"fwdInterface,omitempty"`
}

type _Log Log

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLog(apiId string, apiName string, apiVersion string, resourceName string, protocol string, result string) *Log {
	this := Log{}
	this.ApiId = apiId
	this.ApiName = apiName
	this.ApiVersion = apiVersion
	this.ResourceName = resourceName
	this.Protocol = protocol
	this.Result = result
	return &this
}

// NewLogWithDefaults instantiates a new Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWithDefaults() *Log {
	this := Log{}
	return &this
}

// GetApiId returns the ApiId field value
func (o *Log) GetApiId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value
// and a boolean to check if the value has been set.
func (o *Log) GetApiIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiId, true
}

// SetApiId sets field value
func (o *Log) SetApiId(v string) {
	o.ApiId = v
}

// GetApiName returns the ApiName field value
func (o *Log) GetApiName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value
// and a boolean to check if the value has been set.
func (o *Log) GetApiNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiName, true
}

// SetApiName sets field value
func (o *Log) SetApiName(v string) {
	o.ApiName = v
}

// GetApiVersion returns the ApiVersion field value
func (o *Log) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *Log) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *Log) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetResourceName returns the ResourceName field value
func (o *Log) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *Log) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *Log) SetResourceName(v string) {
	o.ResourceName = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Log) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Log) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Log) SetUri(v string) {
	o.Uri = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *Log) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return o.Operation
}

// HasOperation returns a boolean if a field has been set.
func (o *Log) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given Operation and assigns it to the Operation field.
func (o *Log) SetOperation(v string) {
	o.Operation = v
}

// GetResult returns the Result field value
func (o *Log) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *Log) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *Log) SetResult(v string) {
	o.Result = v
}

// GetInvocationTime returns the InvocationTime field value if set, zero value otherwise.
func (o *Log) GetInvocationTime() time.Time {
	if o == nil || IsNil(o.InvocationTime) {
		var ret time.Time
		return ret
	}
	return *o.InvocationTime
}

// GetInvocationTimeOk returns a tuple with the InvocationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetInvocationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InvocationTime) {
		return nil, false
	}
	return o.InvocationTime, true
}

// HasInvocationTime returns a boolean if a field has been set.
func (o *Log) HasInvocationTime() bool {
	if o != nil && !IsNil(o.InvocationTime) {
		return true
	}

	return false
}

// SetInvocationTime gets a reference to the given time.Time and assigns it to the InvocationTime field.
func (o *Log) SetInvocationTime(v time.Time) {
	o.InvocationTime = &v
}

// GetInvocationLatency returns the InvocationLatency field value if set, zero value otherwise.
func (o *Log) GetInvocationLatency() int32 {
	if o == nil || IsNil(o.InvocationLatency) {
		var ret int32
		return ret
	}
	return *o.InvocationLatency
}

// GetInvocationLatencyOk returns a tuple with the InvocationLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetInvocationLatencyOk() (*int32, bool) {
	if o == nil || IsNil(o.InvocationLatency) {
		return nil, false
	}
	return o.InvocationLatency, true
}

// HasInvocationLatency returns a boolean if a field has been set.
func (o *Log) HasInvocationLatency() bool {
	if o != nil && !IsNil(o.InvocationLatency) {
		return true
	}

	return false
}

// SetInvocationLatency gets a reference to the given int32 and assigns it to the InvocationLatency field.
func (o *Log) SetInvocationLatency(v int32) {
	o.InvocationLatency = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.InputParameters) {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *Log) HasInputParameters() bool {
	if o != nil && !IsNil(o.InputParameters) {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *Log) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

// GetOutputParameters returns the OutputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetOutputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OutputParameters
}

// GetOutputParametersOk returns a tuple with the OutputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetOutputParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OutputParameters) {
		return nil, false
	}
	return &o.OutputParameters, true
}

// HasOutputParameters returns a boolean if a field has been set.
func (o *Log) HasOutputParameters() bool {
	if o != nil && !IsNil(o.OutputParameters) {
		return true
	}

	return false
}

// SetOutputParameters gets a reference to the given interface{} and assigns it to the OutputParameters field.
func (o *Log) SetOutputParameters(v interface{}) {
	o.OutputParameters = v
}

// GetSrcInterface returns the SrcInterface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetSrcInterface() InterfaceDescription {
	if o == nil || IsNil(o.SrcInterface.Get()) {
		var ret InterfaceDescription
		return ret
	}
	return *o.SrcInterface.Get()
}

// GetSrcInterfaceOk returns a tuple with the SrcInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetSrcInterfaceOk() (*InterfaceDescription, bool) {
	if o == nil {
		return nil, false
	}
	return o.SrcInterface.Get(), o.SrcInterface.IsSet()
}

// HasSrcInterface returns a boolean if a field has been set.
func (o *Log) HasSrcInterface() bool {
	if o != nil && o.SrcInterface.IsSet() {
		return true
	}

	return false
}

// SetSrcInterface gets a reference to the given NullableInterfaceDescription and assigns it to the SrcInterface field.
func (o *Log) SetSrcInterface(v InterfaceDescription) {
	o.SrcInterface.Set(&v)
}

// SetSrcInterfaceNil sets the value for SrcInterface to be an explicit nil
func (o *Log) SetSrcInterfaceNil() {
	o.SrcInterface.Set(nil)
}

// UnsetSrcInterface ensures that no value is present for SrcInterface, not even an explicit nil
func (o *Log) UnsetSrcInterface() {
	o.SrcInterface.Unset()
}

// GetDestInterface returns the DestInterface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetDestInterface() InterfaceDescription {
	if o == nil || IsNil(o.DestInterface.Get()) {
		var ret InterfaceDescription
		return ret
	}
	return *o.DestInterface.Get()
}

// GetDestInterfaceOk returns a tuple with the DestInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetDestInterfaceOk() (*InterfaceDescription, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestInterface.Get(), o.DestInterface.IsSet()
}

// HasDestInterface returns a boolean if a field has been set.
func (o *Log) HasDestInterface() bool {
	if o != nil && o.DestInterface.IsSet() {
		return true
	}

	return false
}

// SetDestInterface gets a reference to the given NullableInterfaceDescription and assigns it to the DestInterface field.
func (o *Log) SetDestInterface(v InterfaceDescription) {
	o.DestInterface.Set(&v)
}

// SetDestInterfaceNil sets the value for DestInterface to be an explicit nil
func (o *Log) SetDestInterfaceNil() {
	o.DestInterface.Set(nil)
}

// UnsetDestInterface ensures that no value is present for DestInterface, not even an explicit nil
func (o *Log) UnsetDestInterface() {
	o.DestInterface.Unset()
}

// GetFwdInterface returns the FwdInterface field value if set, zero value otherwise.
func (o *Log) GetFwdInterface() string {
	if o == nil || IsNil(o.FwdInterface) {
		var ret string
		return ret
	}
	return *o.FwdInterface
}

// GetFwdInterfaceOk returns a tuple with the FwdInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetFwdInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.FwdInterface) {
		return nil, false
	}
	return o.FwdInterface, true
}

// HasFwdInterface returns a boolean if a field has been set.
func (o *Log) HasFwdInterface() bool {
	if o != nil && !IsNil(o.FwdInterface) {
		return true
	}

	return false
}

// SetFwdInterface gets a reference to the given string and assigns it to the FwdInterface field.
func (o *Log) SetFwdInterface(v string) {
	o.FwdInterface = &v
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Log) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiId"] = o.ApiId
	toSerialize["apiName"] = o.ApiName
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["resourceName"] = o.ResourceName
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	toSerialize["protocol"] = o.Protocol
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	toSerialize["result"] = o.Result
	if !IsNil(o.InvocationTime) {
		toSerialize["invocationTime"] = o.InvocationTime
	}
	if !IsNil(o.InvocationLatency) {
		toSerialize["invocationLatency"] = o.InvocationLatency
	}
	if o.InputParameters != nil {
		toSerialize["inputParameters"] = o.InputParameters
	}
	if o.OutputParameters != nil {
		toSerialize["outputParameters"] = o.OutputParameters
	}
	if o.SrcInterface.IsSet() {
		toSerialize["srcInterface"] = o.SrcInterface.Get()
	}
	if o.DestInterface.IsSet() {
		toSerialize["destInterface"] = o.DestInterface.Get()
	}
	if !IsNil(o.FwdInterface) {
		toSerialize["fwdInterface"] = o.FwdInterface
	}
	return toSerialize, nil
}

func (o *Log) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiId",
		"apiName",
		"apiVersion",
		"resourceName",
		"protocol",
		"result",
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

	varLog := _Log{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLog)

	if err != nil {
		return err
	}

	*o = Log(varLog)

	return err
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
