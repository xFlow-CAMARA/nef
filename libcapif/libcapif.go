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

package libcapif

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CapifConnector struct {
	configuration *ServiceAPIDescription
	capifSvc      string
}

func NewConnector(url string) *CapifConnector {
	ctr := &CapifConnector{
		capifSvc:      url,
		configuration: nil,
	}
	return ctr
}

func (c *CapifConnector) InstantiateConnector(apiName string, apiVersion string, httpVer string) {

	c.configuration = &ServiceAPIDescription{
		ApiName: apiName,
		AefProfiles: []AefProfile{{
			AefId: "",
			Versions: []Version{{
				ApiVersion:     apiVersion,
				Resources:      []Resource{},
				CustOperations: []CustomOperation{},
			}},
			Protocol:              httpVer,
			DataFormat:            PtrString("JSON"),
			SecurityMethods:       []string{"OAUTH"},
			InterfaceDescriptions: []InterfaceDescription{},
		},
		},

		Description:        PtrString(apiName),
		ServiceAPICategory: PtrString("NEF"),
		SupportedFeatures:  PtrString("0"),
	}
}

func (c *CapifConnector) SetApiSuppFeatures(suppFeat string) {
	c.configuration.ApiSuppFeats = PtrString(suppFeat)
}

func (c *CapifConnector) AddInterface(fqdn string, port int32) {
	interf := &c.configuration.AefProfiles[0].InterfaceDescriptions
	*interf = append(*interf, InterfaceDescription{
		Fqdn: PtrString(fqdn),
		Port: PtrInt32(port),
	})
}

func (c *CapifConnector) AddEndpoint(uri string, commType string, methods []string, description string) {

	opName := uri + methods[0] + "ops"
	resources := &c.configuration.AefProfiles[0].Versions[0].Resources
	*resources = append(*resources, Resource{
		ResourceName: uri,
		CommType:     commType,
		Uri:          uri,
		CustOpName:   PtrString(opName),
		Description:  PtrString(description),
	})

	operations := &c.configuration.AefProfiles[0].Versions[0].CustOperations
	*operations = append(*operations, CustomOperation{
		CommType:   commType,
		CustOpName: opName,
		Operations: methods,
	})
}

func (c *CapifConnector) RegisterService() {
	/*Encode the body*/
	out := &bytes.Buffer{}
	json.NewEncoder(out).Encode(c.configuration)

	req, _ := http.NewRequest("POST", "http://"+c.capifSvc+"/services", out)
	/*Instantiate http clinet*/
	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil || res.StatusCode != http.StatusCreated {
		log.Default().Printf("could not register to capif, disabling feature")
		c.configuration = nil
		return
	}

	defer res.Body.Close()
	/* in v0.1.0 we encode the serviceId as a string*/
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Default().Printf("cannot retrieve capif service id: %s", err.Error())
		c.configuration = nil
		return
	}

	c.configuration.ApiId = PtrString(string(b))

	log.Printf("service correctly registered to capif")
}

func (c *CapifConnector) CapifLog(r *http.Request, statuscode int) error {
	if c.configuration == nil {
		return fmt.Errorf("could not log, capif service not registered")
	}

	invokerId := r.Header.Get("invokerId")
	logData := &LogInfo{}
	logData.InvokerId = invokerId
	logData.Endpoint = r.RequestURI
	logData.Operation = r.Method
	logData.Protocol = c.configuration.AefProfiles[0].Protocol
	logData.Result = fmt.Sprint(statuscode)
	logData.TimeStamp = time.Now().Unix()

	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		logData.SourceIp = ip
		i, err := strconv.ParseInt(port, 10, 32)
		if err != nil {
			logData.SourcePort = 0
		} else {
			logData.SourcePort = int32(i)
		}
		logData.SourceIp = ip
	}

	/* encode to json */
	out := &bytes.Buffer{}
	json.NewEncoder(out).Encode(logData)
	req, _ := http.NewRequest("POST", "http://"+c.capifSvc+"/services/"+*c.configuration.ApiId+"/logs", out)
	/*Instantiate http clinet*/
	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil || res.StatusCode != http.StatusCreated {
		log.Default().Printf("could not send logging request, ignoring")
		return e
	}
	defer res.Body.Close()

	log.Printf("CapifLog -> %s %s %s", logData.Operation, logData.Result, logData.Endpoint)
	return nil
}

func (c *CapifConnector) DeleteService() error {

	if c.capifSvc == "" {
		c.configuration = nil
		return nil
	}
	req, _ := http.NewRequest("DELETE", "http://"+c.capifSvc+"/services/"+*c.configuration.ApiId, nil)
	/*Instantiate http clinet*/
	client := &http.Client{}
	_, e := client.Do(req)
	if e != nil {
		log.Default().Printf("could not delete service")
		c.configuration = nil
		return e
	}
	log.Printf("service correctly deleted from capif")
	return nil
}

func (c *CapifConnector) ValidateToken(token string) (string, error) {
	authData := OauthInfo{}
	authData.Token = token

	/* encode to json */
	out := &bytes.Buffer{}
	json.NewEncoder(out).Encode(authData)
	req, _ := http.NewRequest("POST", "http://"+c.capifSvc+"/services/"+*c.configuration.ApiId+"/validateToken", out)
	/*Instantiate http clinet*/
	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil || res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not validate the token")
	}
	defer res.Body.Close()

	authRes := &OauthInfo{}
	json.NewDecoder(res.Body).Decode(authRes)
	log.Printf("%+v", authRes)
	if authRes.Valid {
		return authRes.InvokerId, nil
	}
	return "", fmt.Errorf("could not validate the token")
}

// middleware
// LoggingMiddleware logs the details of each incoming HTTP request and outgoing response.
func (c *CapifConnector) CapifMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c.configuration == nil {
			next.ServeHTTP(w, r)
			return
		}
		// Log the incoming request details
		// Check for a custom header and respond with an error if missing
		token_raw := r.Header.Get("Authorization")

		/*remove "Bearer" from token string */
		token := strings.Split(token_raw, " ")[1]

		if token != "" {
			invokerId, err := c.ValidateToken(token)
			if err == nil {
				r.Header.Set("InvokerId", invokerId)
			} else {
				log.Printf("token validation error: %s", err.Error())
				http.Error(w, "Could not validate user token", http.StatusForbidden)
				return
			}
		} else {
			http.Error(w, "Authorization Token not provided", http.StatusForbidden)
			return
		}

		// Wrap the ResponseWriter to capture the status code and other response details
		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rec, r)
		c.CapifLog(r, rec.statusCode)
	})
}

// responseRecorder is a wrapper around http.ResponseWriter that lets us capture
// the status code of the response.
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code for logging purposes
func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}
