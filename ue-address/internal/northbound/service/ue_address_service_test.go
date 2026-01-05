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
	"testing"
)

/*
func parseGpsi(gpsi string) (bool, string, error) {
	if strings.HasPrefix(gpsi, "extid-") {
		return true, strings.TrimPrefix(gpsi, "extid-"), nil
	}
	if strings.HasPrefix(gpsi, "msisdn-") {
		return false, strings.TrimPrefix(gpsi, "msisdn-"), nil
	}
	return false, "", fmt.Errorf("invalid gpsi type, must be msisdn or extid")
}
*/

func TestParseGpsi(t *testing.T) {
	gpsiMsidn := "msisdn-1234567890"
	isExtId, identifier, err := parseGpsi(gpsiMsidn)
	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}
	if isExtId != false {
		t.Errorf("expected isExtId to be false, got true")
	}
	if identifier != "1234567890" {
		t.Errorf("expected identifier to be '1234567890', got '%s'", identifier)
	}

	gpsiExtId := "extid-1234567890"
	isExtId, identifier, err = parseGpsi(gpsiExtId)
	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}
	if isExtId != true {
		t.Errorf("expected isExtId to be true, got false")
	}
	if identifier != "1234567890" {
		t.Errorf("expected identifier to be '1234567890', got '%s'", identifier)
	}

	gpsiMalformed1 := "msisd-1234567890"
	isExtId, identifier, err = parseGpsi(gpsiMalformed1)
	if err == nil {
		t.Errorf("expected error, got no error")
	}
	if isExtId != false {
		t.Errorf("expected isExtId to be false, got true")
	}
	if identifier != "" {
		t.Errorf("expected identifier to be blank, got '%s'", identifier)
	}

	gpsiMalformed2 := "msisd1234567890"
	isExtId, identifier, err = parseGpsi(gpsiMalformed2)
	if err == nil {
		t.Errorf("expected error, got no error")
	}
	if isExtId != false {
		t.Errorf("expected isExtId to be false, got true")
	}
	if identifier != "" {
		t.Errorf("expected identifier to be blank, got '%s'", identifier)
	}

	gpsiMalformed3 := "1234567890-msisdn"
	isExtId, identifier, err = parseGpsi(gpsiMalformed3)
	if err == nil {
		t.Errorf("expected error, got no error")
	}
	if isExtId != false {
		t.Errorf("expected isExtId to be false, got true")
	}
	if identifier != "" {
		t.Errorf("expected identifier to be blank, got '%s'", identifier)
	}

}
