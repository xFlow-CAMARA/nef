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

package contexts

import (
	"testing"

	"gitlab.eurecom.fr/open-exposure/nef/monitoring-event/internal/models"
)

func TestAddAf(t *testing.T) {
	afId := "af-test"
	af := NewAf(afId)

	if af.afId != afId {
		t.Errorf("got afId %q, wanted %q", af.afId, afId)
	}

	subscriptions := af.GetAfSubscriptions()
	got := len(subscriptions)
	want := 0 // af is just initialized so no subscription must be there
	if len(subscriptions) != 0 {
		t.Errorf("got number of subscription %q, wanted %q", got, want)
	}
}

func TestAddNewSubscription(t *testing.T) {
	afId := "af-test"
	af := NewAf(afId)

	/* suscription data */
	invalidData := &models.MonitoringEventSubscription{}
	supi := "00106000000001"

	nilSubCtx := af.NewAfSubscription(supi, invalidData)

	if nilSubCtx != nil {
		t.Errorf("accepted invalid subscription data")
	}

	correctData := &models.MonitoringEventSubscription{
		NotificationDestination: "test.com",
		MonitoringType:          models.MonitoringTypeDownlinkDataDeliveryStatus,
	}

	subCtx := af.NewAfSubscription(supi, correctData)

	if subCtx != nil && subCtx.data == nil {
		t.Errorf("got nil data inside subscriptionContext")
	}

	subId := subCtx.subId
	loc := subCtx.GetLocation()
	expected_loc := "/3gpp-monitoring-event/v1/" + afId + "/subscriptions/" + subId

	if loc != expected_loc {
		t.Errorf("got location %q, wanted %q", loc, expected_loc)
	}
	if loc != subCtx.loc {
		t.Errorf("got location %q, wanted %q", loc, expected_loc)
	}
	if subCtx.data.Self != subId {
		t.Errorf("got selfId %s, wanted %s", subCtx.data.Self, subId)
	}

	subscriptions := af.GetAfSubscriptions()
	got := len(subscriptions)
	want := 1 // af is just initialized so no subscription must be there
	if got != 1 {
		t.Errorf("got number of subscription %d, wanted %d", got, want)
	}

	subCtxFromGet := af.GetAfSubscription(subId)
	if subCtxFromGet != subCtx {
		t.Errorf("got different subCtx reference, got %v, wanted %v", subCtxFromGet, subCtx)
	}
}

func TestGetSubscription(t *testing.T) {
	afId := "af-test"
	af := NewAf(afId)

	/* suscription data */
	data := &models.MonitoringEventSubscription{
		NotificationDestination: "test.com",
		MonitoringType:          models.MonitoringTypeDownlinkDataDeliveryStatus,
	}
	supi := "00106000000001"

	subCtx := af.NewAfSubscription(supi, data)

	if subCtx == nil {
		t.Errorf("got nil subscriptionContext")
		return
	}

	subId := subCtx.subId

	subCtxFromGet := af.GetAfSubscription(subId)
	if subCtxFromGet != subCtx {
		t.Errorf("got different subCtx reference, got %v, wanted %v", subCtxFromGet, subCtx)
	}

	wrongId := "fakeSubId"
	subCtxFromGet = af.GetAfSubscription(wrongId)
	if subCtxFromGet != nil {
		t.Errorf("got subCtx reference %v, should be nil", subCtxFromGet)
	}

}

func TestDeleteSubscription(t *testing.T) {
	afId := "af-test"
	af := NewAf(afId)

	/* suscription data */
	data := &models.MonitoringEventSubscription{
		NotificationDestination: "test.com",
		MonitoringType:          models.MonitoringTypeDownlinkDataDeliveryStatus,
	}
	supi := "00106000000001"

	subCtx := af.NewAfSubscription(supi, data)

	if subCtx == nil {
		t.Errorf("got nil subscriptionContext")
		return
	}

	subscriptions := af.GetAfSubscriptions()
	got := len(subscriptions)
	want := 1 // af is just initialized so no subscription must be there
	if got != 1 {
		t.Errorf("got number of subscription %d, wanted %d", got, want)
	}

	subId := subCtx.subId

	gotSubCtx := af.GetAfSubscription(subId)
	if gotSubCtx != subCtx {
		t.Errorf("got different subCtx reference, got %v, wanted %v", gotSubCtx, subCtx)
	}

	err := af.DeleteAfscription(subId)
	if err != nil {
		t.Errorf("could not delete subscription as expected")
	}

	gotSubCtx = af.GetAfSubscription(subId)
	if gotSubCtx != nil {
		t.Errorf("got different subCtx reference %v, wanted nil", gotSubCtx)
	}

	subscriptions = af.GetAfSubscriptions()
	got = len(subscriptions)
	want = 0 // af is just initialized so no subscription must be there
	if len(subscriptions) != 0 {
		t.Errorf("got number of subscription %q, wanted %q", got, want)
	}

	wrongSubId := "ssub"
	err = af.DeleteAfscription(wrongSubId)
	if err == nil {
		t.Errorf("wrong subId did not trigger any error when deleting")
	}

}
