// Copyright 2018 Paul Greenberg greenpau@outlook.com
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

package client

import (
	"bytes"
	"fmt"
	"github.com/pschou/go-json"
	"io"
	"strings"
	"time"
)

type IpArpResponse struct {
	InsAPI struct {
		Outputs struct {
			Output IpArpResponseResult `json:"output"`
		} `json:"outputs"`
		Sid     string `json:"sid"`
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"ins_api"`
}

type IpArpResponseResult struct {
	Body  IpArpResultBody `json:"body"`
	Code  string          `json:"code"`
	Input string          `json:"input"`
	Msg   string          `json:"msg"`
}

type IpArpResultBody struct {
	TableVrf []struct {
		RowVrf []struct {
			TableAdj []struct {
				RowAdj []struct {
					Flags      string `json:"flags"`
					IntfOut    string `json:"intf-out"`
					IPAddrOut  string `json:"ip-addr-out"`
					Mac        string `json:"mac,omitempty"`
					TimeStamp  string `json:"time-stamp"`
					Incomplete string `json:"incomplete,omitempty"`
				} `json:"ROW_adj"`
			} `json:"TABLE_adj"`
			CntTotal   int    `json:"cnt-total"`
			VrfNameOut string `json:"vrf-name-out"`
		} `json:"ROW_vrf"`
	} `json:"TABLE_vrf"`
}

func (d *IpArpResponse) Flat() (out []IpArpResultFlat) {
	for _, Tv := range d.InsAPI.Outputs.Output.Body.TableVrf {
		for _, Rv := range Tv.RowVrf {
			for _, Ta := range Rv.TableAdj {
				for _, Ra := range Ta.RowAdj {
					out = append(out, IpArpResultFlat{
						Flags:      Ra.Flags,
						IntfOut:    Ra.IntfOut,
						IPAddrOut:  Ra.IPAddrOut,
						Mac:        Ra.Mac,
						TimeStamp:  ParseDuration(Ra.TimeStamp),
						Incomplete: Ra.Incomplete,
						CntTotal:   Rv.CntTotal,
						VrfNameOut: Rv.VrfNameOut,
					})
				}
			}
		}
	}
	return
}

type IpArpResultFlat struct {
	Flags      string        `json:"flags"`
	IntfOut    string        `json:"intf-out"`
	IPAddrOut  string        `json:"ip-addr-out"`
	Mac        string        `json:"mac,omitempty"`
	TimeStamp  time.Duration `json:"time-stamp"`
	Incomplete string        `json:"incomplete,omitempty"`
	CntTotal   int           `json:"cnt-total"`
	VrfNameOut string        `json:"vrf-name-out"`
}

// NewIpArpFromString returns SysInfo instance from an input string.
func NewIpArpFromString(s string) (*IpArpResponse, error) {
	return NewIpArpFromReader(strings.NewReader(s))
}

// NewSysInfoFromBytes returns SysInfo instance from an input byte array.
func NewIpArpFromBytes(s []byte) (*IpArpResponse, error) {
	return NewIpArpFromReader(bytes.NewReader(s))
}
func NewIpArpFromReader(s io.Reader) (*IpArpResponse, error) {
	//si := &IpArp{}
	IpArpResponseDat := &IpArpResponse{}
	jsonDec := json.NewDecoder(s)
	//jsonDec.UseNumber()
	jsonDec.UseSlice()
	err := jsonDec.Decode(IpArpResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return IpArpResponseDat, nil
}
