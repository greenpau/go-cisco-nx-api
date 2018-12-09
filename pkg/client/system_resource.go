// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"encoding/json"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"strconv"
)

type systemResourcesResponse struct {
	ID      uint64                        `json:"id" xml:"id"`
	Version string                        `json:"jsonrpc" xml:"jsonrpc"`
	Result  systemResourcesResponseResult `json:"result" xml:"result"`
}

type systemResourcesResponseResult struct {
	Body systemResourcesResponseResultBody `json:"body" xml:"body"`
}

type systemResourcesResponseResultBody struct {
	CPUUsageTable    systemResourcesResponseResultBodyCPUUsageTable `json:"TABLE_cpu_usage" xml:"TABLE_cpu_usage"`
	CPUStateIdle     string                                         `json:"cpu_state_idle" xml:"cpu_state_idle"`
	CPUStateKernel   string                                         `json:"cpu_state_kernel" xml:"cpu_state_kernel"`
	CPUStateUser     string                                         `json:"cpu_state_user" xml:"cpu_state_user"`
	MemoryStatus     string                                         `json:"current_memory_status" xml:"current_memory_status"`
	MemoryUsageFree  uint64                                         `json:"memory_usage_free" xml:"memory_usage_free"`
	MemoryUsageTotal uint64                                         `json:"memory_usage_total" xml:"memory_usage_total"`
	MemoryUsageUsed  uint64                                         `json:"memory_usage_used" xml:"memory_usage_used"`
	LoadAverage1Min  string                                         `json:"load_avg_1min" xml:"load_avg_1min"`
	LoadAverage5Min  string                                         `json:"load_avg_5min" xml:"load_avg_5min"`
	LoadAverage15Min string                                         `json:"load_avg_15min" xml:"load_avg_15min"`
	ProcessesRunning uint64                                         `json:"processes_running" xml:"processes_running"`
	ProcessesTotal   uint64                                         `json:"processes_total" xml:"processes_total"`
}

type systemResourcesResponseResultBodyCPUUsageTable struct {
	CPUUsageRow []systemResourcesResponseResultBodyCPUUsageRow `json:"ROW_cpu_usage" xml:"ROW_cpu_usage"`
}

type systemResourcesResponseResultBodyCPUUsageRow struct {
	ID     int    `json:"cpuid" xml:"cpuid"`
	Idle   string `json:"idle" xml:"idle"`
	Kernel string `json:"kernel" xml:"kernel"`
	User   string `json:"user" xml:"user"`
}

// MemoryUsage is the memory utilization of SystemResources.
type MemoryUsage struct {
	Status string `json:"status" xml:"status"`
	Free   uint64 `json:"free" xml:"free"`
	Used   uint64 `json:"used" xml:"used"`
	Total  uint64 `json:"total" xml:"total"`
}

// CPU represents a CPU instance.
type CPU struct {
	ID    int      `json:"id" xml:"id"`
	Usage CPUUsage `json:"usage" xml:"usage"`
}

// CPUUsage represents CPU utilization.
type CPUUsage struct {
	Idle   float64 `json:"idle" xml:"idle"`
	Kernel float64 `json:"kernel" xml:"kernel"`
	User   float64 `json:"user" xml:"user"`
}

// ProcessUsage represents process stats.
type ProcessUsage struct {
	Running uint64 `json:"running" xml:"running"`
	Total   uint64 `json:"total" xml:"total"`
}

// SystemResources contains information about CPU and memory usage. The
// information in the structure is from the output of
// "show system resources" command.
type SystemResources struct {
	Memory         MemoryUsage        `json:"memory" xml:"memory"`
	CPUAverageLoad map[string]float64 `json:"cpu_avg_load" xml:"cpu_avg_load"`
	CPU            CPUUsage           `json:"cpu" xml:"cpu"`
	CPUs           []CPU              `json:"cpus" xml:"cpus"`
	Processes      ProcessUsage       `json:"processes" xml:"processes"`
}

// NewSystemResourcesFromString returns SystemResources instance from an input string.
func NewSystemResourcesFromString(s string) (*SystemResources, error) {
	return NewSystemResourcesFromBytes([]byte(s))
}

// NewSystemResourcesFromBytes returns SystemResources instance from an input byte array.
func NewSystemResourcesFromBytes(s []byte) (*SystemResources, error) {
	var systemResources *SystemResources
	vResponse := &systemResourcesResponse{}
	err := json.Unmarshal(s, vResponse)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	if len(vResponse.Result.Body.CPUUsageTable.CPUUsageRow) < 1 {
		return nil, fmt.Errorf("Error parsing the received response: %s", s)
	}

	systemResources = &SystemResources{}
	systemResources.CPUAverageLoad = make(map[string]float64)
	systemResources.Memory.Status = vResponse.Result.Body.MemoryStatus
	systemResources.Memory.Free = vResponse.Result.Body.MemoryUsageFree
	systemResources.Memory.Used = vResponse.Result.Body.MemoryUsageUsed
	systemResources.Memory.Total = vResponse.Result.Body.MemoryUsageTotal
	if i, err := strconv.ParseFloat(vResponse.Result.Body.CPUStateIdle, 64); err == nil {
		systemResources.CPU.Idle = i
	}
	if i, err := strconv.ParseFloat(vResponse.Result.Body.CPUStateKernel, 64); err == nil {
		systemResources.CPU.Kernel = i
	}
	if i, err := strconv.ParseFloat(vResponse.Result.Body.CPUStateUser, 64); err == nil {
		systemResources.CPU.User = i
	}
	if i, err := strconv.ParseFloat(vResponse.Result.Body.LoadAverage1Min, 64); err == nil {
		systemResources.CPUAverageLoad["1m"] = i
	}
	if i, err := strconv.ParseFloat(vResponse.Result.Body.LoadAverage5Min, 64); err == nil {
		systemResources.CPUAverageLoad["5m"] = i
	}
	if i, err := strconv.ParseFloat(vResponse.Result.Body.LoadAverage15Min, 64); err == nil {
		systemResources.CPUAverageLoad["15m"] = i
	}
	systemResources.Processes.Running = vResponse.Result.Body.ProcessesRunning
	systemResources.Processes.Total = vResponse.Result.Body.ProcessesTotal
	for _, r := range vResponse.Result.Body.CPUUsageTable.CPUUsageRow {
		c := CPU{
			ID: r.ID,
		}
		if i, err := strconv.ParseFloat(r.Idle, 64); err == nil {
			c.Usage.Idle = i
		}
		if i, err := strconv.ParseFloat(r.Kernel, 64); err == nil {
			c.Usage.Kernel = i
		}
		if i, err := strconv.ParseFloat(r.User, 64); err == nil {
			c.Usage.User = i
		}
		systemResources.CPUs = append(systemResources.CPUs, c)
	}
	return systemResources, nil
}
