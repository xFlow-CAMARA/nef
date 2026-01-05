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

package connector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gitlab.eurecom.fr/open-exposure/nef/ue-address/internal/models"
)

// ------------------------------------------------------------------------------
func (c *Connector) GetCurrentIps(supi string) ([]models.IpAddr, error) {

	/* Execute client code for the 3GPP target NF*/
	// The URL you want to GET
	url := c.Cfg().Sbi.ProfileSvc + fmt.Sprintf("/ue-profile/v1/profiles/%s", supi)

	// Create a custom HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []models.IpAddr{}, err
	}

	// Optionally, add headers
	req.Header.Set("Accept", "application/json")

	// Send the GET request
	resp, err := client.Do(req)
	if err != nil {
		return []models.IpAddr{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return []models.IpAddr{}, fmt.Errorf("the ue %s is not active", supi)
		}
		return []models.IpAddr{}, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.IpAddr{}, fmt.Errorf("error reading response body")
	}

	ueProfile := &models.UeProfile{}
	if err := json.Unmarshal(body, ueProfile); err != nil {
		return []models.IpAddr{}, fmt.Errorf("error parsing response body")
	}

	result := []models.IpAddr{}
	for _, session := range ueProfile.PduSessions {
		if session != nil {
			result = append(result, models.IpAddr{
				Ipv4Addr: session.Ipv4,
				//TODO add support for ipv6 sessions
			})
		}
	}

	return result, nil
}
