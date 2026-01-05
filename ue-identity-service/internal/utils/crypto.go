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

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

// EncodeIMSIWithAfID encrypts IMSI + AfId into an external ID
func EncodeIMSIWithAfID(imsi string, afId string, key []byte) (string, error) {
	// Concatenate IMSI + AF ID with separator (use a safe separator)
	payload := fmt.Sprintf("%s|%s", imsi, afId)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(payload), nil)
	final := append(nonce, ciphertext...)

	return base64.RawURLEncoding.EncodeToString(final), nil
}

// DecodeIMSIWithAfID decrypts an external ID to recover IMSI + AfId
func DecodeIMSIWithAfID(externalID string, key []byte) (string, string, error) {
	data, err := base64.RawURLEncoding.DecodeString(externalID)
	if err != nil {
		return "", "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", "", fmt.Errorf("invalid external ID")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", "", err
	}

	// Split payload back to IMSI + AfId
	parts := string(plaintext)

	decodedPayload := strings.Split(parts, "|")
	if len(decodedPayload) != 2 {
		return "", "", fmt.Errorf("invalid payload format")
	}

	imsi := decodedPayload[0]
	afId := decodedPayload[1]

	return imsi, afId, nil
}
