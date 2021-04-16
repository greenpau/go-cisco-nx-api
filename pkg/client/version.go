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
	"bytes"
	"fmt"
	"github.com/pschou/go-json"
	"io"
	"strings"
	"time"
)

type VersionResponse struct {
	InsAPI struct {
		Outputs struct {
			Output VersionResponseResult `json:"output"`
		} `json:"outputs"`
		Sid     string `json:"sid"`
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"ins_api"`
}

type VersionResponseResult struct {
	Body  VersionResultBody `json:"body"`
	Code  string            `json:"code"`
	Input string            `json:"input"`
	Msg   string            `json:"msg"`
}

type VersionResultBody struct {
	TablePackageList []struct {
		RowPackageList []struct {
			PackageID []struct {
			} `json:"package_id"`
		} `json:"ROW_package_list"`
	} `json:"TABLE_package_list"`
	BiosCmplTime    string `json:"bios_cmpl_time"`
	BiosVerStr      string `json:"bios_ver_str"`
	BootflashSize   int    `json:"bootflash_size"`
	ChassisID       string `json:"chassis_id"`
	CPUName         string `json:"cpu_name"`
	HeaderStr       string `json:"header_str"`
	HostName        string `json:"host_name"`
	KernUptmDays    int    `json:"kern_uptm_days"`
	KernUptmHrs     int    `json:"kern_uptm_hrs"`
	KernUptmMins    int    `json:"kern_uptm_mins"`
	KernUptmSecs    int    `json:"kern_uptm_secs"`
	KernUpTime      time.Duration
	KickCmplTime    string `json:"kick_cmpl_time"`
	KickFileName    string `json:"kick_file_name"`
	KickTmstmp      string `json:"kick_tmstmp"`
	KickstartVerStr string `json:"kickstart_ver_str"`
	Manufacturer    string `json:"manufacturer"`
	MemType         string `json:"mem_type"`
	Memory          int    `json:"memory"`
	ModuleID        string `json:"module_id"`
	ProcBoardID     string `json:"proc_board_id"`
	RrCtime         string `json:"rr_ctime"`
	RrReason        string `json:"rr_reason"`
	RrService       string `json:"rr_service"`
	RrSysVer        string `json:"rr_sys_ver"`
	RrUsecs         int    `json:"rr_usecs"`
}

// NewVersionFromString returns SysInfo instance from an input string.
func NewVersionFromString(s string) (*VersionResponseResult, error) {
	return NewVersionFromReader(strings.NewReader(s))
}

// NewSysInfoFromBytes returns SysInfo instance from an input byte array.
func NewVersionFromBytes(s []byte) (*VersionResponseResult, error) {
	return NewVersionFromReader(bytes.NewReader(s))
}
func NewVersionFromReader(s io.Reader) (*VersionResponseResult, error) {
	//si := &Version{}
	VersionResponseDat := &VersionResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseNumber()
	jsonDec.UseSlice()
	err := jsonDec.Decode(VersionResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	//var Body *VersionResultBody
	Body := VersionResponseDat.InsAPI.Outputs.Output.Body
	VersionResponseDat.InsAPI.Outputs.Output.Body.KernUpTime =
		time.Second*time.Duration(Body.KernUptmSecs) +
			time.Minute*time.Duration(Body.KernUptmMins) +
			time.Hour*time.Duration(Body.KernUptmHrs+24*Body.KernUptmDays)
	return &VersionResponseDat.InsAPI.Outputs.Output, nil
}
