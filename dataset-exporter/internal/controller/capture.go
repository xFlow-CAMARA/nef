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

package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.eurecom.fr/open-exposure/nef/dataset-exporter/internal/service"
)

type CaptureRequest struct {
	Duration string `json:"duration"` // e.g., "5m"
}

func StartCaptureHandler(svc *service.CaptureService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CaptureRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		dur, err := time.ParseDuration(req.Duration)
		if err != nil {
			http.Error(w, "Invalid duration format", http.StatusBadRequest)
			return
		}

		transactionID, err := svc.StartCaptureJob(dur)
		if err != nil {
			http.Error(w, "Could not start capture", http.StatusInternalServerError)
			return
		}

		resp := map[string]string{}
		resp["transactionId"] = transactionID
		resp["statusPath"] = "/dataset-exporter/v1/captures/status?id=" + transactionID
		resp["downloadPath"] = "/dataset-exporter/v1/captures/download?id=" + transactionID
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Could not encode response", http.StatusInternalServerError)
			return
		}
	}
}

func GetCaptureStatusHandler(svc *service.CaptureService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionID := r.URL.Query().Get("id")
		status, remaining, err := svc.GetCaptureStatus(transactionID)
		if err != nil {
			http.Error(w, "Capture not found", http.StatusNotFound)
			return
		}

		resp := map[string]interface{}{
			"status":        status,
			"remainingTime": remaining.String(),
			"downloadPath":  "/dataset-exporter/v1/captures/download?id=" + transactionID,
		}
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Could not encode response", http.StatusInternalServerError)
			return
		}
	}
}

func DownloadCaptureHandler(svc *service.CaptureService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionID := r.URL.Query().Get("id")
		file, err := svc.GetCaptureFile(transactionID)
		if err != nil {
			http.Error(w, "File not ready or not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=\"capture.csv\"")

		if _, err := w.Write(file); err != nil {
			http.Error(w, "Could not encode response", http.StatusInternalServerError)
			return
		}
	}
}
