// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//            and Paul Schou     (github.com/pschou)
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
			Output VersionResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type VersionResponseResult struct {
	Body  VersionResultBody `json:"body" xml:"body"`
	Code  string            `json:"code" xml:"code"`
	Input string            `json:"input" xml:"input"`
	Msg   string            `json:"msg" xml:"msg"`
}

type VersionResultBody struct {
	TablePackageList []struct {
		RowPackageList []struct {
			PackageID []struct {
			} `json:"package_id" xml:"package_id"`
		} `json:"ROW_package_list" xml:"ROW_package_list"`
	} `json:"TABLE_package_list" xml:"TABLE_package_list"`
	BiosCmplTime    string `json:"bios_cmpl_time" xml:"bios_cmpl_time"`
	BiosVerStr      string `json:"bios_ver_str" xml:"bios_ver_str"`
	BootflashSize   int    `json:"bootflash_size" xml:"bootflash_size"`
	ChassisID       string `json:"chassis_id" xml:"chassis_id"`
	CPUName         string `json:"cpu_name" xml:"cpu_name"`
	HeaderStr       string `json:"header_str" xml:"header_str"`
	HostName        string `json:"host_name" xml:"host_name"`
	KernUptmDays    int    `json:"kern_uptm_days" xml:"kern_uptm_days"`
	KernUptmHrs     int    `json:"kern_uptm_hrs" xml:"kern_uptm_hrs"`
	KernUptmMins    int    `json:"kern_uptm_mins" xml:"kern_uptm_mins"`
	KernUptmSecs    int    `json:"kern_uptm_secs" xml:"kern_uptm_secs"`
	KernUpTime      time.Duration
	KickCmplTime    string `json:"kick_cmpl_time" xml:"kick_cmpl_time"`
	KickFileName    string `json:"kick_file_name" xml:"kick_file_name"`
	KickTmstmp      string `json:"kick_tmstmp" xml:"kick_tmstmp"`
	KickstartVerStr string `json:"kickstart_ver_str" xml:"kickstart_ver_str"`
	Manufacturer    string `json:"manufacturer" xml:"manufacturer"`
	MemType         string `json:"mem_type" xml:"mem_type"`
	Memory          int    `json:"memory" xml:"memory"`
	ModuleID        string `json:"module_id" xml:"module_id"`
	ProcBoardID     string `json:"proc_board_id" xml:"proc_board_id"`
	RrCtime         string `json:"rr_ctime" xml:"rr_ctime"`
	RrReason        string `json:"rr_reason" xml:"rr_reason"`
	RrService       string `json:"rr_service" xml:"rr_service"`
	RrSysVer        string `json:"rr_sys_ver" xml:"rr_sys_ver"`
	RrUsecs         int    `json:"rr_usecs" xml:"rr_usecs"`
}

// NewVersionFromString returns instance from an input string.
func NewVersionFromString(s string) (*VersionResponse, error) {
	return NewVersionFromReader(strings.NewReader(s))
}

// NewVersionFromBytes returns instance from an input byte array.
func NewVersionFromBytes(s []byte) (*VersionResponse, error) {
	return NewVersionFromReader(bytes.NewReader(s))
}

// NewVersionFromReader returns instance from an input reader.
func NewVersionFromReader(s io.Reader) (*VersionResponse, error) {
	//si := &Version{}
	VersionResponseDat := &VersionResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(VersionResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return VersionResponseDat, nil
}

// NewVersionResultFromString returns instance from an input string.
func NewVersionResultFromString(s string) (*VersionResponseResult, error) {
	return NewVersionResultFromReader(strings.NewReader(s))
}

// NewVersionResultFromBytes returns instance from an input byte array.
func NewVersionResultFromBytes(s []byte) (*VersionResponseResult, error) {
	return NewVersionResultFromReader(bytes.NewReader(s))
}

// NewVersionResultFromReader returns instance from an input reader.
func NewVersionResultFromReader(s io.Reader) (*VersionResponseResult, error) {
	//si := &VersionResponseResult{}
	VersionResponseResultDat := &VersionResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(VersionResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return VersionResponseResultDat, nil
}
