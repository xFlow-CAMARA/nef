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

package libcapif

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CustomOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomOperation{}

// CustomOperation Represents the description of a custom operation.
type CustomOperation struct {
	CommType string `json:"commType"`
	// it is set as {custOpName} part of the URI structure for a custom operation without resource association as defined in clause 5.2.4 of 3GPP TS 29.122.
	CustOpName string `json:"custOpName"`
	// Supported HTTP methods for the API resource. Only applicable when the protocol in AefProfile indicates HTTP.
	Operations []string `json:"operations,omitempty"`
	// Text description of the custom operation
	Description *string `json:"description,omitempty"`
}

type _CustomOperation CustomOperation

// NewCustomOperation instantiates a new CustomOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomOperation(commType string, custOpName string) *CustomOperation {
	this := CustomOperation{}
	this.CommType = commType
	this.CustOpName = custOpName
	return &this
}

// NewCustomOperationWithDefaults instantiates a new CustomOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomOperationWithDefaults() *CustomOperation {
	this := CustomOperation{}
	return &this
}

// GetCommType returns the CommType field value
func (o *CustomOperation) GetCommType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommType
}

// GetCommTypeOk returns a tuple with the CommType field value
// and a boolean to check if the value has been set.
func (o *CustomOperation) GetCommTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommType, true
}

// SetCommType sets field value
func (o *CustomOperation) SetCommType(v string) {
	o.CommType = v
}

// GetCustOpName returns the CustOpName field value
func (o *CustomOperation) GetCustOpName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustOpName
}

// GetCustOpNameOk returns a tuple with the CustOpName field value
// and a boolean to check if the value has been set.
func (o *CustomOperation) GetCustOpNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustOpName, true
}

// SetCustOpName sets field value
func (o *CustomOperation) SetCustOpName(v string) {
	o.CustOpName = v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *CustomOperation) GetOperations() []string {
	if o == nil || IsNil(o.Operations) {
		var ret []string
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOperation) GetOperationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *CustomOperation) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []string and assigns it to the Operations field.
func (o *CustomOperation) SetOperations(v []string) {
	o.Operations = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomOperation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOperation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomOperation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomOperation) SetDescription(v string) {
	o.Description = &v
}

func (o CustomOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commType"] = o.CommType
	toSerialize["custOpName"] = o.CustOpName
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *CustomOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commType",
		"custOpName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCustomOperation := _CustomOperation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomOperation)

	if err != nil {
		return err
	}

	*o = CustomOperation(varCustomOperation)

	return err
}

type NullableCustomOperation struct {
	value *CustomOperation
	isSet bool
}

func (v NullableCustomOperation) Get() *CustomOperation {
	return v.value
}

func (v *NullableCustomOperation) Set(val *CustomOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomOperation(val *CustomOperation) *NullableCustomOperation {
	return &NullableCustomOperation{value: val, isSet: true}
}

func (v NullableCustomOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
