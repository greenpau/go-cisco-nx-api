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
)

type sysInfoResponse struct {
	ID      uint64                `json:"id" xml:"id"`
	Version string                `json:"jsonrpc" xml:"jsonrpc"`
	Result  sysInfoResponseResult `json:"result" xml:"result"`
}

type sysInfoResponseResult struct {
	Body sysInfoResponseResultBody `json:"body" xml:"body"`
}

type sysInfoResponseResultBody struct {
	BiosCmplTime       string `json:"bios_cmpl_time" xml:"bios_cmpl_time"`
	BiosVerStr         string `json:"bios_ver_str" xml:"bios_ver_str"`
	BootflashSize      int64  `json:"bootflash_size" xml:"bootflash_size"`
	ChassisID          string `json:"chassis_id" xml:"chassis_id"`
	CPUName            string `json:"cpu_name" xml:"cpu_name"`
	HostName           string `json:"host_name" xml:"host_name"`
	KernUptmDays       int64  `json:"kern_uptm_days" xml:"kern_uptm_days"`
	KernUptmHrs        int64  `json:"kern_uptm_hrs" xml:"kern_uptm_hrs"`
	KernUptmMins       int64  `json:"kern_uptm_mins" xml:"kern_uptm_mins"`
	KernUptmSecs       int64  `json:"kern_uptm_secs" xml:"kern_uptm_secs"`
	KickCmplTime       string `json:"kick_cmpl_time" xml:"kick_cmpl_time"`
	KickFileName       string `json:"kick_file_name" xml:"kick_file_name"`
	KickTmstmp         string `json:"kick_tmstmp" xml:"kick_tmstmp"`
	KickstartVerStr    string `json:"kickstart_ver_str" xml:"kickstart_ver_str"`
	Manufacturer       string `json:"manufacturer" xml:"manufacturer"`
	MemType            string `json:"mem_type" xml:"mem_type"`
	Memory             int64  `json:"memory" xml:"memory"`
	ProcBoardID        string `json:"proc_board_id" xml:"proc_board_id"`
	ResetTime          string `json:"rr_ctime" xml:"rr_ctime"`
	ResetReason        string `json:"rr_reason" xml:"rr_reason"`
	ResetSystemVersion string `json:"rr_sys_ver" xml:"rr_sys_ver"`
}

// SysInfo contains system information. The information in the structure
// is from the output of "show version" command.
type SysInfo struct {
	Bios struct {
		CompileTime string `json:"compile_time" xml:"compile_time"`
		Version      string `json:"version" xml:"version"`
	}
	Bootflash struct {
		Size int64 `json:"size" xml:"size"`
	}
	ChassisID      string `json:"chassis_id" xml:"chassis_id"`
	CPUName        string `json:"cpu_name" xml:"cpu_name"`
	Hostname       string `json:"hostname" xml:"hostname"`
	Uptime         int64  `json:"uptime" xml:"uptime"`
	KickstartImage struct {
		CompileTime string `json:"compile_time" xml:"compile_time"`
		FileName    string `json:"filename" xml:"filename"`
		Timestamp   string `json:"timestamp" xml:"timestamp"`
		Version     string `json:"version" xml:"version"`
	}
	Manufacturer string `json:"manufacturer" xml:"manufacturer"`
	Memory       struct {
		Unit string `json:"unit" xml:"unit"`
		Size int64  `json:"size" xml:"size"`
	}
	ProcessorBoardID string `json:"proc_board_id" xml:"proc_board_id"`
	Reset            struct {
		Time          string `json:"time" xml:"time"`
		Reason        string `json:"reason" xml:"reason"`
		SystemVersion string `json:"sys_ver" xml:"sys_ver"`
	}
}

// NewSysInfoFromString returns SysInfo instance from an input string.
func NewSysInfoFromString(s string) (*SysInfo, error) {
	return NewSysInfoFromBytes([]byte(s))
}

// NewSysInfoFromBytes returns SysInfo instance from an input byte array.
func NewSysInfoFromBytes(s []byte) (*SysInfo, error) {
	si := &SysInfo{}
	siResponse := &sysInfoResponse{}
	err := json.Unmarshal(s, siResponse)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	si.Bios.CompileTime = siResponse.Result.Body.BiosCmplTime
	si.Bios.Version = siResponse.Result.Body.BiosVerStr
	si.Bootflash.Size = siResponse.Result.Body.BootflashSize
	si.ChassisID = siResponse.Result.Body.ChassisID
	si.CPUName = siResponse.Result.Body.CPUName
	si.Hostname = siResponse.Result.Body.HostName
	si.Uptime = siResponse.Result.Body.KernUptmSecs + (siResponse.Result.Body.KernUptmMins * 60) +
		(siResponse.Result.Body.KernUptmHrs * 3600) + (siResponse.Result.Body.KernUptmDays * 86400)
	si.KickstartImage.CompileTime = siResponse.Result.Body.KickCmplTime
	si.KickstartImage.FileName = siResponse.Result.Body.KickFileName
	si.KickstartImage.Timestamp = siResponse.Result.Body.KickTmstmp
	si.KickstartImage.Version = siResponse.Result.Body.KickstartVerStr
	si.Manufacturer = siResponse.Result.Body.Manufacturer
	si.Memory.Size = siResponse.Result.Body.Memory
	si.Memory.Unit = siResponse.Result.Body.MemType
	si.ProcessorBoardID = siResponse.Result.Body.ProcBoardID
	si.Reset.Time = siResponse.Result.Body.ResetTime
	si.Reset.Reason = siResponse.Result.Body.ResetReason
	si.Reset.SystemVersion = siResponse.Result.Body.ResetSystemVersion
	if siResponse.Result.Body.ProcBoardID == "" {
		return nil, fmt.Errorf("Error parsing the received response: %s", s)
	}
	return si, nil
}
