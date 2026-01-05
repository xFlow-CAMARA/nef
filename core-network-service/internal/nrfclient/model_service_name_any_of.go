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
	"fmt"
)

// ServiceNameAnyOf the model 'ServiceNameAnyOf'
type ServiceNameAnyOf string

// List of ServiceName_anyOf
const (
	NNRF_NFM                  ServiceNameAnyOf = "nnrf-nfm"
	NNRF_DISC                 ServiceNameAnyOf = "nnrf-disc"
	NUDM_SDM                  ServiceNameAnyOf = "nudm-sdm"
	NUDM_UECM                 ServiceNameAnyOf = "nudm-uecm"
	NUDM_UEAU                 ServiceNameAnyOf = "nudm-ueau"
	NUDM_EE                   ServiceNameAnyOf = "nudm-ee"
	NUDM_PP                   ServiceNameAnyOf = "nudm-pp"
	NUDM_NIDDAU               ServiceNameAnyOf = "nudm-niddau"
	NAMF_COMM                 ServiceNameAnyOf = "namf_communication"
	NAMF_EVTS                 ServiceNameAnyOf = "namf-evts"
	NAMF_MT                   ServiceNameAnyOf = "namf-mt"
	NAMF_LOC                  ServiceNameAnyOf = "namf-loc"
	NSMF_PDUSESSION           ServiceNameAnyOf = "nsmf-pdusession"
	NSMF_EVENT_EXPOSURE       ServiceNameAnyOf = "nsmf-event-exposure"
	NAUSF_AUTH                ServiceNameAnyOf = "nausf-auth"
	NAUSF_SORPROTECTION       ServiceNameAnyOf = "nausf-sorprotection"
	NAUSF_UPUPROTECTION       ServiceNameAnyOf = "nausf-upuprotection"
	NNEF_PFDMANAGEMENT        ServiceNameAnyOf = "nnef-pfdmanagement"
	NPCF_AM_POLICY_CONTROL    ServiceNameAnyOf = "npcf-am-policy-control"
	NPCF_SMPOLICYCONTROL      ServiceNameAnyOf = "npcf-smpolicycontrol"
	NPCF_POLICYAUTHORIZATION  ServiceNameAnyOf = "npcf-policyauthorization"
	NPCF_BDTPOLICYCONTROL     ServiceNameAnyOf = "npcf-bdtpolicycontrol"
	NPCF_EVENTEXPOSURE        ServiceNameAnyOf = "npcf-eventexposure"
	NPCF_UE_POLICY_CONTROL    ServiceNameAnyOf = "npcf-ue-policy-control"
	NSMSF_SMS                 ServiceNameAnyOf = "nsmsf-sms"
	NNSSF_NSSELECTION         ServiceNameAnyOf = "nnssf-nsselection"
	NNSSF_NSSAIAVAILABILITY   ServiceNameAnyOf = "nnssf-nssaiavailability"
	NUDR_DR                   ServiceNameAnyOf = "nudr-dr"
	NLMF_LOC                  ServiceNameAnyOf = "nlmf-loc"
	N5G_EIR_EIC               ServiceNameAnyOf = "n5g-eir-eic"
	NBSF_MANAGEMENT           ServiceNameAnyOf = "nbsf-management"
	NCHF_SPENDINGLIMITCONTROL ServiceNameAnyOf = "nchf-spendinglimitcontrol"
	NCHF_CONVERGEDCHARGING    ServiceNameAnyOf = "nchf-convergedcharging"
	NNWDAF_EVENTSSUBSCRIPTION ServiceNameAnyOf = "nnwdaf-eventssubscription"
	NNWDAF_ANALYTICSINFO      ServiceNameAnyOf = "nnwdaf-analyticsinfo"
)

// All allowed values of ServiceNameAnyOf enum
var AllowedServiceNameAnyOfEnumValues = []ServiceNameAnyOf{
	"nnrf-nfm",
	"nnrf-disc",
	"nudm-sdm",
	"nudm-uecm",
	"nudm-ueau",
	"nudm-ee",
	"nudm-pp",
	"nudm-niddau",
	"namf_communication",
	"namf-evts",
	"namf-mt",
	"namf-loc",
	"nsmf-pdusession",
	"nsmf_event-exposure",
	"nausf-auth",
	"nausf-sorprotection",
	"nausf-upuprotection",
	"nnef-pfdmanagement",
	"npcf-am-policy-control",
	"npcf-smpolicycontrol",
	"npcf-policyauthorization",
	"npcf-bdtpolicycontrol",
	"npcf-eventexposure",
	"npcf-ue-policy-control",
	"nsmsf-sms",
	"nnssf-nsselection",
	"nnssf-nssaiavailability",
	"nudr-dr",
	"nlmf-loc",
	"n5g-eir-eic",
	"nbsf-management",
	"nchf-spendinglimitcontrol",
	"nchf-convergedcharging",
	"nnwdaf-eventssubscription",
	"nnwdaf-analyticsinfo",
}

func (v *ServiceNameAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceNameAnyOf(value)
	for _, existing := range AllowedServiceNameAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceNameAnyOf", value)
}

// NewServiceNameAnyOfFromValue returns a pointer to a valid ServiceNameAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceNameAnyOfFromValue(v string) (*ServiceNameAnyOf, error) {
	ev := ServiceNameAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceNameAnyOf: valid values are %v", v, AllowedServiceNameAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceNameAnyOf) IsValid() bool {
	for _, existing := range AllowedServiceNameAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceName_anyOf value
func (v ServiceNameAnyOf) Ptr() *ServiceNameAnyOf {
	return &v
}

type NullableServiceNameAnyOf struct {
	value *ServiceNameAnyOf
	isSet bool
}

func (v NullableServiceNameAnyOf) Get() *ServiceNameAnyOf {
	return v.value
}

func (v *NullableServiceNameAnyOf) Set(val *ServiceNameAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNameAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNameAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNameAnyOf(val *ServiceNameAnyOf) *NullableServiceNameAnyOf {
	return &NullableServiceNameAnyOf{value: val, isSet: true}
}

func (v NullableServiceNameAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNameAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
