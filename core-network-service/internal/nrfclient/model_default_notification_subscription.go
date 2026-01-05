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

// DefaultNotificationSubscription struct for DefaultNotificationSubscription
type DefaultNotificationSubscription struct {
	NotificationType   NotificationType    `json:"notificationType"`
	CallbackUri        string              `json:"callbackUri"`
	N1MessageClass     *N1MessageClass     `json:"n1MessageClass,omitempty"`
	N2InformationClass *N2InformationClass `json:"n2InformationClass,omitempty"`
}

// NewDefaultNotificationSubscription instantiates a new DefaultNotificationSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultNotificationSubscription(notificationType NotificationType, callbackUri string) *DefaultNotificationSubscription {
	this := DefaultNotificationSubscription{}
	this.NotificationType = notificationType
	this.CallbackUri = callbackUri
	return &this
}

// NewDefaultNotificationSubscriptionWithDefaults instantiates a new DefaultNotificationSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultNotificationSubscriptionWithDefaults() *DefaultNotificationSubscription {
	this := DefaultNotificationSubscription{}
	return &this
}

// GetNotificationType returns the NotificationType field value
func (o *DefaultNotificationSubscription) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *DefaultNotificationSubscription) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetCallbackUri returns the CallbackUri field value
func (o *DefaultNotificationSubscription) GetCallbackUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUri
}

// GetCallbackUriOk returns a tuple with the CallbackUri field value
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetCallbackUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUri, true
}

// SetCallbackUri sets field value
func (o *DefaultNotificationSubscription) SetCallbackUri(v string) {
	o.CallbackUri = v
}

// GetN1MessageClass returns the N1MessageClass field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetN1MessageClass() N1MessageClass {
	if o == nil || o.N1MessageClass == nil {
		var ret N1MessageClass
		return ret
	}
	return *o.N1MessageClass
}

// GetN1MessageClassOk returns a tuple with the N1MessageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetN1MessageClassOk() (*N1MessageClass, bool) {
	if o == nil || o.N1MessageClass == nil {
		return nil, false
	}
	return o.N1MessageClass, true
}

// HasN1MessageClass returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasN1MessageClass() bool {
	if o != nil && o.N1MessageClass != nil {
		return true
	}

	return false
}

// SetN1MessageClass gets a reference to the given N1MessageClass and assigns it to the N1MessageClass field.
func (o *DefaultNotificationSubscription) SetN1MessageClass(v N1MessageClass) {
	o.N1MessageClass = &v
}

// GetN2InformationClass returns the N2InformationClass field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetN2InformationClass() N2InformationClass {
	if o == nil || o.N2InformationClass == nil {
		var ret N2InformationClass
		return ret
	}
	return *o.N2InformationClass
}

// GetN2InformationClassOk returns a tuple with the N2InformationClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetN2InformationClassOk() (*N2InformationClass, bool) {
	if o == nil || o.N2InformationClass == nil {
		return nil, false
	}
	return o.N2InformationClass, true
}

// HasN2InformationClass returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasN2InformationClass() bool {
	if o != nil && o.N2InformationClass != nil {
		return true
	}

	return false
}

// SetN2InformationClass gets a reference to the given N2InformationClass and assigns it to the N2InformationClass field.
func (o *DefaultNotificationSubscription) SetN2InformationClass(v N2InformationClass) {
	o.N2InformationClass = &v
}

func (o DefaultNotificationSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["notificationType"] = o.NotificationType
	}
	if true {
		toSerialize["callbackUri"] = o.CallbackUri
	}
	if o.N1MessageClass != nil {
		toSerialize["n1MessageClass"] = o.N1MessageClass
	}
	if o.N2InformationClass != nil {
		toSerialize["n2InformationClass"] = o.N2InformationClass
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultNotificationSubscription struct {
	value *DefaultNotificationSubscription
	isSet bool
}

func (v NullableDefaultNotificationSubscription) Get() *DefaultNotificationSubscription {
	return v.value
}

func (v *NullableDefaultNotificationSubscription) Set(val *DefaultNotificationSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultNotificationSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultNotificationSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultNotificationSubscription(val *DefaultNotificationSubscription) *NullableDefaultNotificationSubscription {
	return &NullableDefaultNotificationSubscription{value: val, isSet: true}
}

func (v NullableDefaultNotificationSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultNotificationSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
