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
	//"time"
)

type IsisAdjDetailResponse struct {
	InsAPI struct {
		Outputs struct {
			Output IsisAdjDetailResponseResult `json:"output"`
		} `json:"outputs"`
		Sid     string `json:"sid"`
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"ins_api"`
}

type IsisAdjDetailResponseResult struct {
	Body  IsisAdjDetailResultBody `json:"body"`
	Code  string                  `json:"code"`
	Input string                  `json:"input"`
	Msg   string                  `json:"msg"`
}

type IsisAdjDetailResultBody struct {
	TableProcessTag []struct {
		RowProcessTag []struct {
			ProcessTagOut string `json:"process-tag-out"`
			TableVrf      []struct {
				RowVrf []struct {
					VrfNameOut      string `json:"vrf-name-out"`
					AdjSummaryOut   string `json:"adj-summary-out"`
					AdjInterfaceOut string `json:"adj-interface-out"`
					TableProcessAdj []struct {
						RowProcessAdj []struct {
							AdjSysNameOut              string `json:"adj-sys-name-out"`
							AdjSysIDOut                string `json:"adj-sys-id-out"`
							AdjUsageOut                string `json:"adj-usage-out"`
							AdjStateOut                string `json:"adj-state-out"`
							AdjHoldTimeOut             string `json:"adj-hold-time-out"`
							AdjIntfNameOut             string `json:"adj-intf-name-out"`
							AdjDetailSetOut            string `json:"adj-detail-set-out"`
							AdjTransitionsOut          string `json:"adj-transitions-out"`
							AdjFlapOut                 string `json:"adj-flap-out"`
							AdjFlapTimeOut             string `json:"adj-flap-time-out"`
							AdjCktTypeOut              string `json:"adj-ckt-type-out"`
							AdjIpv4AddrOut             string `json:"adj-ipv4-addr-out"`
							AdjIpv6AddrOut             string `json:"adj-ipv6-addr-out"`
							AdjBcastOut                string `json:"adj-bcast-out"`
							AdjBfdIpv4EstablishOut     string `json:"adj-bfd-ipv4-establish-out"`
							AdjBfdIpv6EstablishOut     string `json:"adj-bfd-ipv6-establish-out"`
							AdjResurrectOut            string `json:"adj-resurrect-out"`
							AdjRestartCapableOut       string `json:"adj-restart-capable-out"`
							AdjRestartAckOut           string `json:"adj-restart-ack-out"`
							AdjRestartModeOut          string `json:"adj-restart-mode-out"`
							AdjRestartAdjSeenRaOut     string `json:"adj-restart-adj-seen-ra-out"`
							AdjRestartAdjSeenCsnpOut   string `json:"adj-restart-adj-seen-csnp-out"`
							AdjRestartAdjSeenL1CsnpOut string `json:"adj-restart-adj-seen-l1-csnp-out"`
							AdjRestartAdjSeenL2CsnpOut string `json:"adj-restart-adj-seen-l2-csnp-out"`
							AdjRestartSuppressAdjOut   string `json:"adj-restart-suppress-adj-out"`
							TableAdjSid                struct {
								RowAdjSid []struct {
									AdjSidValue  string `json:"adj-sid-value"`
									AdjSidFFlag  string `json:"adj-sid-f-flag"`
									AdjSidBFlag  string `json:"adj-sid-b-flag"`
									AdjSidVFlag  string `json:"adj-sid-v-flag"`
									AdjSidLFlag  string `json:"adj-sid-l-flag"`
									AdjSidSFlag  string `json:"adj-sid-s-flag"`
									AdjSidPFlag  string `json:"adj-sid-p-flag"`
									AdjSidWeight string `json:"adj-sid-weight"`
								} `json:"ROW_adj_sid"`
							} `json:"TABLE_adj_sid"`
						} `json:"ROW_process_adj"`
					} `json:"TABLE_process_adj"`
				} `json:"ROW_vrf"`
			} `json:"TABLE_vrf"`
		} `json:"ROW_process_tag"`
	} `json:"TABLE_process_tag"`
}

func (d *IsisAdjDetailResponse) Flat() (out []IsisAdjDetailResultFlat) {
	for _, Tpt := range d.InsAPI.Outputs.Output.Body.TableProcessTag {
		for _, Rpt := range Tpt.RowProcessTag {
			for _, Tv := range Rpt.TableVrf {
				for _, Rv := range Tv.RowVrf {
					for _, Tp := range Rv.TableProcessAdj {
						for _, Rp := range Tp.RowProcessAdj {
							out = append(out, IsisAdjDetailResultFlat{
								//Rv
								VrfNameOut:      Rv.VrfNameOut,
								AdjSummaryOut:   Rv.AdjSummaryOut,
								AdjInterfaceOut: Rv.AdjInterfaceOut,
								//Rp
								AdjSysNameOut:              Rp.AdjSysNameOut,
								AdjSysIDOut:                Rp.AdjSysIDOut,
								AdjUsageOut:                Rp.AdjUsageOut,
								AdjStateOut:                Rp.AdjStateOut,
								AdjHoldTimeOut:             Rp.AdjHoldTimeOut,
								AdjIntfNameOut:             Rp.AdjIntfNameOut,
								AdjDetailSetOut:            Rp.AdjDetailSetOut,
								AdjTransitionsOut:          Rp.AdjTransitionsOut,
								AdjFlapOut:                 Rp.AdjFlapOut,
								AdjFlapTimeOut:             Rp.AdjFlapTimeOut,
								AdjCktTypeOut:              Rp.AdjCktTypeOut,
								AdjIpv4AddrOut:             Rp.AdjIpv4AddrOut,
								AdjIpv6AddrOut:             Rp.AdjIpv6AddrOut,
								AdjBcastOut:                Rp.AdjBcastOut,
								AdjBfdIpv4EstablishOut:     Rp.AdjBfdIpv4EstablishOut,
								AdjBfdIpv6EstablishOut:     Rp.AdjBfdIpv6EstablishOut,
								AdjResurrectOut:            Rp.AdjResurrectOut,
								AdjRestartCapableOut:       Rp.AdjRestartCapableOut,
								AdjRestartAckOut:           Rp.AdjRestartAckOut,
								AdjRestartModeOut:          Rp.AdjRestartModeOut,
								AdjRestartAdjSeenRaOut:     Rp.AdjRestartAdjSeenRaOut,
								AdjRestartAdjSeenCsnpOut:   Rp.AdjRestartAdjSeenCsnpOut,
								AdjRestartAdjSeenL1CsnpOut: Rp.AdjRestartAdjSeenL1CsnpOut,
								AdjRestartAdjSeenL2CsnpOut: Rp.AdjRestartAdjSeenL2CsnpOut,
								AdjRestartSuppressAdjOut:   Rp.AdjRestartSuppressAdjOut,
							})
						}
					}
				}
			}
		}
	}
	return
}

type IsisAdjDetailResultFlat struct {
	//Rv
	VrfNameOut      string `json:"vrf-name-out"`
	AdjSummaryOut   string `json:"adj-summary-out"`
	AdjInterfaceOut string `json:"adj-interface-out"`
	//Rp
	AdjSysNameOut              string `json:"adj-sys-name-out"`
	AdjSysIDOut                string `json:"adj-sys-id-out"`
	AdjUsageOut                string `json:"adj-usage-out"`
	AdjStateOut                string `json:"adj-state-out"`
	AdjHoldTimeOut             string `json:"adj-hold-time-out"`
	AdjIntfNameOut             string `json:"adj-intf-name-out"`
	AdjDetailSetOut            string `json:"adj-detail-set-out"`
	AdjTransitionsOut          string `json:"adj-transitions-out"`
	AdjFlapOut                 string `json:"adj-flap-out"`
	AdjFlapTimeOut             string `json:"adj-flap-time-out"`
	AdjCktTypeOut              string `json:"adj-ckt-type-out"`
	AdjIpv4AddrOut             string `json:"adj-ipv4-addr-out"`
	AdjIpv6AddrOut             string `json:"adj-ipv6-addr-out"`
	AdjBcastOut                string `json:"adj-bcast-out"`
	AdjBfdIpv4EstablishOut     string `json:"adj-bfd-ipv4-establish-out"`
	AdjBfdIpv6EstablishOut     string `json:"adj-bfd-ipv6-establish-out"`
	AdjResurrectOut            string `json:"adj-resurrect-out"`
	AdjRestartCapableOut       string `json:"adj-restart-capable-out"`
	AdjRestartAckOut           string `json:"adj-restart-ack-out"`
	AdjRestartModeOut          string `json:"adj-restart-mode-out"`
	AdjRestartAdjSeenRaOut     string `json:"adj-restart-adj-seen-ra-out"`
	AdjRestartAdjSeenCsnpOut   string `json:"adj-restart-adj-seen-csnp-out"`
	AdjRestartAdjSeenL1CsnpOut string `json:"adj-restart-adj-seen-l1-csnp-out"`
	AdjRestartAdjSeenL2CsnpOut string `json:"adj-restart-adj-seen-l2-csnp-out"`
	AdjRestartSuppressAdjOut   string `json:"adj-restart-suppress-adj-out"`
}

// NewIsisAdjDetailFromString returns SysInfo instance from an input string.
func NewIsisAdjDetailFromString(s string) (*IsisAdjDetailResponse, error) {
	return NewIsisAdjDetailFromReader(strings.NewReader(s))
}

// NewSysInfoFromBytes returns SysInfo instance from an input byte array.
func NewIsisAdjDetailFromBytes(s []byte) (*IsisAdjDetailResponse, error) {
	return NewIsisAdjDetailFromReader(bytes.NewReader(s))
}
func NewIsisAdjDetailFromReader(s io.Reader) (*IsisAdjDetailResponse, error) {
	//si := &IsisAdjDetail{}
	IsisAdjDetailResponseDat := &IsisAdjDetailResponse{}
	jsonDec := json.NewDecoder(s)
	//jsonDec.UseNumber()
	jsonDec.UseSlice()
	err := jsonDec.Decode(IsisAdjDetailResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return IsisAdjDetailResponseDat, nil
}
