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
	"log"
	"net/http"
	"time"
)

// ------------------------------------------------------------------------------
func (c *Connector) LookupExternalId(afId string, externalId string) (string, error) {
	/* Execute client code for the 3GPP target NF*/
	// The URL you want to GET
	url := c.Cfg().Sbi.IdentitySvc + "/resolve?externalId=" + externalId

	// Create a custom HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Optionally, add headers
	req.Header.Set("Accept", "application/json")

	// Send the GET request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Printf("could not close response body correctly")
		}
	}(resp.Body)

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return "", fmt.Errorf("the external ID %s was not found", externalId)
		}
		return "", fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body")
	}

	var val map[string]interface{}
	if err := json.Unmarshal(body, &val); err != nil {
		return "", fmt.Errorf("error parsing response body")
	}

	return val["Supi"].(string), nil
}
