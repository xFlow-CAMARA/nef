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

type LogInfo struct {
	Operation       string `json:"operation"`
	SourcePort      int32  `json:"sourcePort"`
	SourceIp        string `json:"sourceIp"`
	DestinationPort int32  `json:"destinationPort"`
	DestinationIp   string `json:"destinationIp"`
	Endpoint        string `json:"endpoint"`
	Result          string `json:"result"`
	TimeStamp       int64  `json:"timestamp"`
	Protocol        string `json:"protocol"`
	InvokerId       string `json:"invokerId"`
}
