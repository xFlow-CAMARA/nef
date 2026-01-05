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
	"time"
)

// checks if the Version type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Version{}

// Version Represents the API version information.
type Version struct {
	// API major version in URI (e.g. v1)
	ApiVersion string `json:"apiVersion"`
	// string with format \"date-time\" as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// Resources supported by the API.
	Resources []Resource `json:"resources,omitempty"`
	// Custom operations without resource association.
	CustOperations []CustomOperation `json:"custOperations,omitempty"`
}

type _Version Version

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion(apiVersion string) *Version {
	this := Version{}
	this.ApiVersion = apiVersion
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *Version) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *Version) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *Version) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *Version) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *Version) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *Version) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *Version) GetResources() []Resource {
	if o == nil || IsNil(o.Resources) {
		var ret []Resource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetResourcesOk() ([]Resource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *Version) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []Resource and assigns it to the Resources field.
func (o *Version) SetResources(v []Resource) {
	o.Resources = v
}

// GetCustOperations returns the CustOperations field value if set, zero value otherwise.
func (o *Version) GetCustOperations() []CustomOperation {
	if o == nil || IsNil(o.CustOperations) {
		var ret []CustomOperation
		return ret
	}
	return o.CustOperations
}

// GetCustOperationsOk returns a tuple with the CustOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetCustOperationsOk() ([]CustomOperation, bool) {
	if o == nil || IsNil(o.CustOperations) {
		return nil, false
	}
	return o.CustOperations, true
}

// HasCustOperations returns a boolean if a field has been set.
func (o *Version) HasCustOperations() bool {
	if o != nil && !IsNil(o.CustOperations) {
		return true
	}

	return false
}

// SetCustOperations gets a reference to the given []CustomOperation and assigns it to the CustOperations field.
func (o *Version) SetCustOperations(v []CustomOperation) {
	o.CustOperations = v
}

func (o Version) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Version) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.CustOperations) {
		toSerialize["custOperations"] = o.CustOperations
	}
	return toSerialize, nil
}

func (o *Version) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
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

	varVersion := _Version{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVersion)

	if err != nil {
		return err
	}

	*o = Version(varVersion)

	return err
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
