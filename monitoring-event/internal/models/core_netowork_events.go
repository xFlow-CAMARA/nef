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

package models

type CoreNetworkEvent string

const (
	CORENETWORKEVENT_LOCATION_REPORT              CoreNetworkEvent = "LOCATION_REPORT"
	CORENETWORKEVENT_REGISTRATION_STATE_REPORT    CoreNetworkEvent = "REGISTRATION_STATE_REPORT"
	CORENETWORKEVENT_CONNECTIVITY_STATE_REPORT    CoreNetworkEvent = "CONNECTIVITY_STATE_REPORT"
	CORENETWORKEVENT_REACHABILITY_REPORT          CoreNetworkEvent = "REACHABILITY_REPORT"
	CORENETWORKEVENT_COMMUNICATION_FAILURE_REPORT CoreNetworkEvent = "COMMUNICATION_FAILURE_REPORT"
	CORENETWORKEVENT_LOSS_OF_CONNECTIVITY         CoreNetworkEvent = "LOSS_OF_CONNECTIVITY"
	CORENETWORKEVENT_UP_PATH_CH                   CoreNetworkEvent = "UP_PATH_CH"              // required
	CORENETWORKEVENT_PDU_SES_REL                  CoreNetworkEvent = "PDU_SES_REL"             // impl
	CORENETWORKEVENT_PLMN_CH                      CoreNetworkEvent = "PLMN_CH"                 // impl
	CORENETWORKEVENT_UE_IP_CH                     CoreNetworkEvent = "UE_IP_CH"                // impl
	CORENETWORKEVENT_DDDS                         CoreNetworkEvent = "DDDS"                    // impl
	CORENETWORKEVENT_PDU_SES_EST                  CoreNetworkEvent = "PDU_SES_EST"             // impl
	CORENETWORKEVENT_QOS_MON                      CoreNetworkEvent = "QOS_MON"                 // impl
	CORENETWORKEVENT_PDN_CONNECTIVITY_STATUS      CoreNetworkEvent = "PDN_CONNECTIVITY_STATUS" // impl (un umbrella event fo)

)
