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
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

// Secret key for pseudonymization
// TODO decide whether to change salt for each transaction, will be safer but user capturing across multiple transaction won't be able to merge the data

var secretSalt = []byte(uuid.NewString()) // should come from config/env

type CaptureStatus string

const (
	StatusInProgress CaptureStatus = "IN_PROGRESS"
	StatusCompleted  CaptureStatus = "COMPLETED"
	StatusFailed     CaptureStatus = "FAILED"
)

type CaptureJob struct {
	ID        string
	Status    CaptureStatus
	ExpiresAt time.Time
	Buffer    bytes.Buffer
	Mutex     sync.Mutex
}

type CaptureService struct {
	RedisClient *redis.Client
	Jobs        map[string]*CaptureJob
	JobsMutex   sync.Mutex
}

func NewCaptureService(redisClient *redis.Client) *CaptureService {
	return &CaptureService{
		RedisClient: redisClient,
		Jobs:        map[string]*CaptureJob{},
		JobsMutex:   sync.Mutex{},
	}
}

func (cs *CaptureService) StartCaptureJob(duration time.Duration) (string, error) {
	id := uuid.New().String()
	job := &CaptureJob{
		ID:        id,
		Status:    StatusInProgress,
		ExpiresAt: time.Now().Add(duration),
	}

	cs.JobsMutex.Lock()
	cs.Jobs[id] = job
	cs.JobsMutex.Unlock()

	go cs.runCapture(job, duration)

	return id, nil
}

func (cs *CaptureService) runCapture(job *CaptureJob, duration time.Duration) {
	ctx := context.Background()
	sub := cs.RedisClient.PSubscribe(ctx, "broadcast:*")

	defer func() {
		_ = sub.Close()
	}()

	ch := sub.Channel()

	writer := csv.NewWriter(&job.Buffer)
	if err := writer.Write([]string{"timestamp", "event", "user", "data"}); err != nil {
		log.Printf("could not write CSV header: %v", err)
		writer.Flush()
		job.Mutex.Lock()
		job.Status = StatusFailed
		job.Mutex.Unlock()
		return
	}

	timer := time.NewTimer(duration)

	for {

		select {
		case msg := <-ch:
			var event map[string]interface{}
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				continue
			}

			data, err := json.Marshal(event["data"])
			if err != nil {
				continue
			}
			maskedUE := maskUEIdentity(event)
			if err := writer.Write([]string{
				time.Now().Format(time.RFC3339),
				event["type"].(string),
				maskedUE,
				string(data),
			}); err != nil {
				writer.Flush()
				job.Mutex.Lock()
				job.Status = StatusFailed
				job.Mutex.Unlock()
				return
			}
		case <-timer.C:
			writer.Flush()
			job.Mutex.Lock()
			job.Status = StatusCompleted
			job.Mutex.Unlock()
			return
		}
	}
}

func (cs *CaptureService) GetCaptureStatus(id string) (CaptureStatus, time.Duration, error) {
	cs.JobsMutex.Lock()
	job, exists := cs.Jobs[id]
	cs.JobsMutex.Unlock()
	if !exists {
		return "", 0, errors.New("capture not found")
	}

	remaining := time.Until(job.ExpiresAt)
	if remaining < 0 {
		remaining = 0
	}

	return job.Status, remaining, nil
}

func (cs *CaptureService) GetCaptureFile(id string) ([]byte, error) {
	cs.JobsMutex.Lock()
	job, exists := cs.Jobs[id]
	cs.JobsMutex.Unlock()
	if !exists || job.Status != StatusCompleted {
		return nil, errors.New("file not ready")
	}
	return job.Buffer.Bytes(), nil
}

func maskUEIdentity(event map[string]interface{}) string {
	if imsi, ok := event["imsi"].(string); ok {
		return pseudonymizeIMSI(imsi)
	}
	return "UE-UNKNOWN"
}

func pseudonymizeIMSI(imsi string) string {
	mac := hmac.New(sha256.New, secretSalt)
	mac.Write([]byte(imsi))
	hash := mac.Sum(nil)

	// Take first 6 bytes to keep pseudonym short
	return "UE-" + hex.EncodeToString(hash)[:12]
}
