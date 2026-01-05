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

package capif

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"

	"gitlab.eurecom.fr/open-exposure/nef/capif-service/internal/utils"
)

func retrieveAuthorizationDetails(resp_json *AuthProviderResponse) error {

	url := fmt.Sprintf("https://%s:%s/getauth", utils.Config.CapifFqdn, utils.Config.CapifAuthPort)
	// Create a HTTP post request
	r, err := http.NewRequest("GET", url, nil)
	r.SetBasicAuth(utils.Config.CapifUsername, utils.Config.CapifPassword)
	if err != nil {
		return err

	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	res, err := client.Do(r)
	if err != nil {
		return err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	err = json.NewDecoder(res.Body).Decode(resp_json)
	if err != nil {
		return err
	}

	return nil
}

func retrieveAccessToken() (string, error) {
	res := AuthProviderResponse{}
	err := retrieveAuthorizationDetails(&res)
	if err != nil {
		return "", err
	}
	return res.AccessToken, nil
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

func CreateCsr(name string) ([]byte, []byte, error) {
	log.Printf("Generating certificate signing request for %s", name)

	// step: generate a keypair
	keys, err := generatePrivateKey(2048)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to genarate private keys, error: %s", err)
	}

	// step: generate a csr template
	var csrTemplate = x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   name,
			Organization: []string{"EURECOM"},
			Country:      []string{"FR"},
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrCertificate, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, keys)
	if err != nil {
		return nil, nil, err
	}
	csr := pem.EncodeToMemory(&pem.Block{
		Type: "CERTIFICATE REQUEST", Bytes: csrCertificate,
	})

	priv := encodePrivateKeyToPEM(keys)
	err = os.WriteFile(utils.Config.CertsPath+""+name+".pem", priv, 0755)
	if err != nil {
		return nil, nil, fmt.Errorf("could not store private key on file: %v", err)
	}
	return csr, priv, nil

}

func httpClientAMFCerts() *http.Client {

	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	certs, err := os.ReadFile(utils.Config.CertsPath + "ca.crt")
	if err != nil {
		log.Fatalf("Failed to append ca_root.crt to RootCAs: %v", err)
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}
	//	cert, _ := tls.LoadX509KeyPair("APF.crt", "client.key")
	cert, _ := tls.LoadX509KeyPair(utils.Config.CertsPath+"AMF.crt", utils.Config.CertsPath+"AMF.pem")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      rootCAs,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	return client
}
