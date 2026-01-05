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
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func ValidateInvokerToken(aefId string, apiName string, authData *OauthInfo) error {
	// Read the public key from the file
	pubKeyBytes, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		return err
	}

	// Parse the RSA public key
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		return err
	}

	claims := jwt.MapClaims{}

	_, err = jwt.ParseWithClaims(authData.Token, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the token's signing method is RSA
		//fmt.Printf("%+v", token)
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})

	log.Default().Printf("%+v", claims)
	for key, val := range claims {
		if key == "sub" {
			authData.InvokerId = val.(string)
			break
		}
	}

	if err != nil {
		authData.Valid = false
		return fmt.Errorf("error parsing jwt: %v", err)
	}

	authData.Valid = true

	for key, val := range claims {
		if key == "scope" {
			scope := strings.Split(val.(string), "#")[1]
			authData.AefId = strings.Split(scope, ":")[0]
			authData.ApiName = strings.Split(scope, ":")[1]
			break
		}
	}

	/*validate here*/
	if authData.AefId == aefId && authData.ApiName == apiName && authData.Valid {
		return nil
	} else {
		return fmt.Errorf("user not authorized")
	}

}
