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

package service

import (
	"fmt"
	"testing"

	"gitlab.eurecom.fr/open-exposure/nef/as-session-with-qos/internal/models"
)

func TestAsSessionWithQoSValidation(t *testing.T) {
	data := &models.AsSessionWithQoSSubscription{
		UeIpv4Addr:              "12.1.1.1",
		NotificationDestination: "http://notifications",
		Dnn:                     "internet",
		QosReference:            "QoS1",
		FlowInfo: []models.FlowInfo{{
			FlowId:           1,
			FlowDescriptions: []string{"permit ip 0.0.0.0 0.0.0.0"},
		},
		},
	}
	err := validateSubscriptionData(data)
	if err != nil {
		t.Errorf("expected no errors, got %s", err.Error())
	}

	dataNoIp := &models.AsSessionWithQoSSubscription{
		NotificationDestination: "http://notifications",
		Dnn:                     "internet",
		QosReference:            "QoS1",
		FlowInfo: []models.FlowInfo{{
			FlowId:           1,
			FlowDescriptions: []string{"permit ip 0.0.0.0 0.0.0.0"},
		},
		},
	}
	err = validateSubscriptionData(dataNoIp)
	expectedErr := fmt.Errorf("field UeIpv4Addr not provided")
	if err == nil {
		t.Errorf("expected %s, no error", expectedErr.Error())
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("expected %s, got %s", expectedErr.Error(), err.Error())
	}

	dataNoNotifDestination := &models.AsSessionWithQoSSubscription{
		UeIpv4Addr:   "12.1.1.1",
		Dnn:          "internet",
		QosReference: "QoS1",
		FlowInfo: []models.FlowInfo{{
			FlowId:           1,
			FlowDescriptions: []string{"permit ip 0.0.0.0 0.0.0.0"},
		},
		},
	}
	err = validateSubscriptionData(dataNoNotifDestination)
	expectedErr = fmt.Errorf("field NotificationDestination not provided")
	if err == nil {
		t.Errorf("expected %s, no error", expectedErr.Error())
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("expected %s, got %s", expectedErr.Error(), err.Error())
	}

	dataNoDnn := &models.AsSessionWithQoSSubscription{
		UeIpv4Addr:              "12.1.1.1",
		NotificationDestination: "http://notifications",
		QosReference:            "QoS1",
		FlowInfo: []models.FlowInfo{{
			FlowId:           1,
			FlowDescriptions: []string{"permit ip 0.0.0.0 0.0.0.0"},
		},
		},
	}
	err = validateSubscriptionData(dataNoDnn)
	expectedErr = fmt.Errorf("field Dnn not provided")
	if err == nil {
		t.Errorf("expected %s, no error", expectedErr.Error())
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("expected %s, got %s", expectedErr.Error(), err.Error())
	}

	dataNoQosRef := &models.AsSessionWithQoSSubscription{
		UeIpv4Addr:              "12.1.1.1",
		NotificationDestination: "http://notifications",
		Dnn:                     "internet",
		FlowInfo: []models.FlowInfo{{
			FlowId:           1,
			FlowDescriptions: []string{"permit ip 0.0.0.0 0.0.0.0"},
		},
		},
	}
	err = validateSubscriptionData(dataNoQosRef)
	expectedErr = fmt.Errorf("field QosReference not provided")
	if err == nil {
		t.Errorf("expected %s, no error", expectedErr.Error())
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("expected %s, got %s", expectedErr.Error(), err.Error())
	}

	dataNoFlowInfo := &models.AsSessionWithQoSSubscription{
		UeIpv4Addr:              "12.1.1.1",
		NotificationDestination: "http://notifications",
		Dnn:                     "internet",
		QosReference:            "QoS1",
	}
	err = validateSubscriptionData(dataNoFlowInfo)
	expectedErr = fmt.Errorf("field FlowInfo not provided")
	if err == nil {
		t.Errorf("expected %s, no error", expectedErr.Error())
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("expected %s, got %s", expectedErr.Error(), err.Error())
	}

	dataNoFlowDesc := &models.AsSessionWithQoSSubscription{
		UeIpv4Addr:              "12.1.1.1",
		NotificationDestination: "http://notifications",
		Dnn:                     "internet",
		QosReference:            "QoS1",
		FlowInfo: []models.FlowInfo{{
			FlowId: 1,
		},
		},
	}
	err = validateSubscriptionData(dataNoFlowDesc)
	expectedErr = fmt.Errorf("no FlowDescription provided")
	if err == nil {
		t.Errorf("expected %s, no error", expectedErr.Error())
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("expected %s, got %s", expectedErr.Error(), err.Error())
	}

	dataInvalidFlowDesc := &models.AsSessionWithQoSSubscription{
		UeIpv4Addr:              "12.1.1.1",
		NotificationDestination: "http://notifications",
		Dnn:                     "internet",
		QosReference:            "QoS1",
		FlowInfo: []models.FlowInfo{{
			FlowId:           1,
			FlowDescriptions: []string{},
		},
		},
	}
	err = validateSubscriptionData(dataInvalidFlowDesc)
	expectedErr = fmt.Errorf("no FlowDescription provided")
	if err == nil {
		t.Errorf("expected %s, no error", expectedErr.Error())
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("expected %s, got %s", expectedErr.Error(), err.Error())
	}

}
