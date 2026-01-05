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
CAPIF_API_Provider_Management_API

API for API provider domain functions management.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.3
*/

package api_provider_management

import (
	"bytes"
	"encoding/json"
)

// checks if the APIProviderEnrolmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIProviderEnrolmentDetails{}

// APIProviderEnrolmentDetails Represents an API provider domain's enrolment details.
type APIProviderEnrolmentDetails struct {
	// API provider domain ID assigned by the CAPIF core function to the API management function while registering the API provider domain. Shall not be present in the HTTP POST request from the API Management function to the CAPIF core function, to on-board itself. Shall be present in all other HTTP requests and responses.
	ApiProvDomId *string `json:"apiProvDomId,omitempty"`
	// Security information necessary for the CAPIF core function to validate the registration of the API provider domain. Shall be present in HTTP POST request from API management function to CAPIF core function for API provider domain registration.
	RegSec string `json:"regSec"`
	// A list of individual API provider domain functions details. When included by the API management function in the HTTP request message, it lists the API provider domain functions that the API management function intends to register/update in registration or update registration procedure. When included by the CAPIF core function in the HTTP response message, it lists the API domain functions details that are registered or updated successfully.
	ApiProvFuncs []APIProviderFunctionDetails `json:"apiProvFuncs,omitempty"`
	// Generic information related to the API provider domain such as details of the API provider applications.
	ApiProvDomInfo *string `json:"apiProvDomInfo,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`
	// Registration or update specific failure information of failed API provider domain function registrations.Shall be present in the HTTP response body if atleast one of the API provider domain function registration or update registration fails.
	FailReason *string `json:"failReason,omitempty"`
	// Represents the API provider name.
	ApiProvName *string `json:"apiProvName,omitempty"`
}

type _APIProviderEnrolmentDetails APIProviderEnrolmentDetails

// NewAPIProviderEnrolmentDetails instantiates a new APIProviderEnrolmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIProviderEnrolmentDetails(regSec string) *APIProviderEnrolmentDetails {
	this := APIProviderEnrolmentDetails{}
	this.RegSec = regSec
	return &this
}

// NewAPIProviderEnrolmentDetailsWithDefaults instantiates a new APIProviderEnrolmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIProviderEnrolmentDetailsWithDefaults() *APIProviderEnrolmentDetails {
	this := APIProviderEnrolmentDetails{}
	return &this
}

// GetApiProvDomId returns the ApiProvDomId field value if set, zero value otherwise.
func (o *APIProviderEnrolmentDetails) GetApiProvDomId() string {
	if o == nil || IsNil(o.ApiProvDomId) {
		var ret string
		return ret
	}
	return *o.ApiProvDomId
}

// GetApiProvDomIdOk returns a tuple with the ApiProvDomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetails) GetApiProvDomIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiProvDomId) {
		return nil, false
	}
	return o.ApiProvDomId, true
}

// HasApiProvDomId returns a boolean if a field has been set.
func (o *APIProviderEnrolmentDetails) HasApiProvDomId() bool {
	if o != nil && !IsNil(o.ApiProvDomId) {
		return true
	}

	return false
}

// SetApiProvDomId gets a reference to the given string and assigns it to the ApiProvDomId field.
func (o *APIProviderEnrolmentDetails) SetApiProvDomId(v string) {
	o.ApiProvDomId = &v
}

// GetRegSec returns the RegSec field value
func (o *APIProviderEnrolmentDetails) GetRegSec() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegSec
}

// GetRegSecOk returns a tuple with the RegSec field value
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetails) GetRegSecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegSec, true
}

// SetRegSec sets field value
func (o *APIProviderEnrolmentDetails) SetRegSec(v string) {
	o.RegSec = v
}

// GetApiProvFuncs returns the ApiProvFuncs field value if set, zero value otherwise.
func (o *APIProviderEnrolmentDetails) GetApiProvFuncs() []APIProviderFunctionDetails {
	if o == nil || IsNil(o.ApiProvFuncs) {
		var ret []APIProviderFunctionDetails
		return ret
	}
	return o.ApiProvFuncs
}

// GetApiProvFuncsOk returns a tuple with the ApiProvFuncs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetails) GetApiProvFuncsOk() ([]APIProviderFunctionDetails, bool) {
	if o == nil || IsNil(o.ApiProvFuncs) {
		return nil, false
	}
	return o.ApiProvFuncs, true
}

// HasApiProvFuncs returns a boolean if a field has been set.
func (o *APIProviderEnrolmentDetails) HasApiProvFuncs() bool {
	if o != nil && !IsNil(o.ApiProvFuncs) {
		return true
	}

	return false
}

// SetApiProvFuncs gets a reference to the given []APIProviderFunctionDetails and assigns it to the ApiProvFuncs field.
func (o *APIProviderEnrolmentDetails) SetApiProvFuncs(v []APIProviderFunctionDetails) {
	o.ApiProvFuncs = v
}

// GetApiProvDomInfo returns the ApiProvDomInfo field value if set, zero value otherwise.
func (o *APIProviderEnrolmentDetails) GetApiProvDomInfo() string {
	if o == nil || IsNil(o.ApiProvDomInfo) {
		var ret string
		return ret
	}
	return *o.ApiProvDomInfo
}

// GetApiProvDomInfoOk returns a tuple with the ApiProvDomInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetails) GetApiProvDomInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ApiProvDomInfo) {
		return nil, false
	}
	return o.ApiProvDomInfo, true
}

// HasApiProvDomInfo returns a boolean if a field has been set.
func (o *APIProviderEnrolmentDetails) HasApiProvDomInfo() bool {
	if o != nil && !IsNil(o.ApiProvDomInfo) {
		return true
	}

	return false
}

// SetApiProvDomInfo gets a reference to the given string and assigns it to the ApiProvDomInfo field.
func (o *APIProviderEnrolmentDetails) SetApiProvDomInfo(v string) {
	o.ApiProvDomInfo = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *APIProviderEnrolmentDetails) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetails) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *APIProviderEnrolmentDetails) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *APIProviderEnrolmentDetails) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *APIProviderEnrolmentDetails) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetails) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *APIProviderEnrolmentDetails) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *APIProviderEnrolmentDetails) SetFailReason(v string) {
	o.FailReason = &v
}

// GetApiProvName returns the ApiProvName field value if set, zero value otherwise.
func (o *APIProviderEnrolmentDetails) GetApiProvName() string {
	if o == nil || IsNil(o.ApiProvName) {
		var ret string
		return ret
	}
	return *o.ApiProvName
}

// GetApiProvNameOk returns a tuple with the ApiProvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIProviderEnrolmentDetails) GetApiProvNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiProvName) {
		return nil, false
	}
	return o.ApiProvName, true
}

// HasApiProvName returns a boolean if a field has been set.
func (o *APIProviderEnrolmentDetails) HasApiProvName() bool {
	if o != nil && !IsNil(o.ApiProvName) {
		return true
	}

	return false
}

// SetApiProvName gets a reference to the given string and assigns it to the ApiProvName field.
func (o *APIProviderEnrolmentDetails) SetApiProvName(v string) {
	o.ApiProvName = &v
}

func (o APIProviderEnrolmentDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIProviderEnrolmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiProvDomId) {
		toSerialize["apiProvDomId"] = o.ApiProvDomId
	}
	toSerialize["regSec"] = o.RegSec
	if !IsNil(o.ApiProvFuncs) {
		toSerialize["apiProvFuncs"] = o.ApiProvFuncs
	}
	if !IsNil(o.ApiProvDomInfo) {
		toSerialize["apiProvDomInfo"] = o.ApiProvDomInfo
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !IsNil(o.FailReason) {
		toSerialize["failReason"] = o.FailReason
	}
	if !IsNil(o.ApiProvName) {
		toSerialize["apiProvName"] = o.ApiProvName
	}
	return toSerialize, nil
}

func (o *APIProviderEnrolmentDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"regSec",
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

	varAPIProviderEnrolmentDetails := _APIProviderEnrolmentDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAPIProviderEnrolmentDetails)

	if err != nil {
		return err
	}

	*o = APIProviderEnrolmentDetails(varAPIProviderEnrolmentDetails)

	return err
}

type NullableAPIProviderEnrolmentDetails struct {
	value *APIProviderEnrolmentDetails
	isSet bool
}

func (v NullableAPIProviderEnrolmentDetails) Get() *APIProviderEnrolmentDetails {
	return v.value
}

func (v *NullableAPIProviderEnrolmentDetails) Set(val *APIProviderEnrolmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIProviderEnrolmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIProviderEnrolmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIProviderEnrolmentDetails(val *APIProviderEnrolmentDetails) *NullableAPIProviderEnrolmentDetails {
	return &NullableAPIProviderEnrolmentDetails{value: val, isSet: true}
}

func (v NullableAPIProviderEnrolmentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIProviderEnrolmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
