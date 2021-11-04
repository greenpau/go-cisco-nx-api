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

type IpRouteResponse struct {
	InsAPI struct {
		Outputs struct {
			Output IpRouteResponseResult `json:"output"`
		} `json:"outputs"`
		Sid     string `json:"sid"`
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"ins_api"`
}

type IpRouteResponseResult struct {
	Body  IpRouteResultBody `json:"body"`
	Code  string            `json:"code"`
	Input string            `json:"input"`
	Msg   string            `json:"msg"`
}

type IpRouteResultBody struct {
	TableVrf []struct {
		RowVrf []struct {
			TableAddrf []struct {
				RowAddrf []struct {
					TablePrefix []struct {
						RowPrefix []struct {
							TablePath []struct {
								RowPath []struct {
									Clientname string `json:"clientname"`
									Ifname     string `json:"ifname"`
									Metric     int    `json:"metric"`
									Pref       int    `json:"pref"`
									UBest      string `json:"ubest"`
									UpTime     string `json:"uptime"`
								} `json:"ROW_path"`
							} `json:"TABLE_path"`
							Attached   string `json:"attached"`
							IPPrefix   string `json:"ipprefix"`
							McastNhops int    `json:"mcast-nhops"`
							UcastNhops int    `json:"ucast-nhops"`
						} `json:"ROW_prefix"`
					} `json:"TABLE_prefix"`
					AddRf string `json:"addrf"`
				} `json:"ROW_addrf"`
			} `json:"TABLE_addrf"`
			VrfNameOut string `json:"vrf-name-out"`
		} `json:"ROW_vrf"`
	} `json:"TABLE_vrf"`
}

func (d *IpRouteResponse) Flat() (out []IpRouteResultFlat) {
	for _, Tv := range d.InsAPI.Outputs.Output.Body.TableVrf {
		for _, Rv := range Tv.RowVrf {
			for _, Ta := range Rv.TableAddrf {
				for _, Ra := range Ta.RowAddrf {
					for _, Tpre := range Ra.TablePrefix {
						for _, Rpre := range Tpre.RowPrefix {
							for _, Tp := range Rpre.TablePath {
								for _, Rp := range Tp.RowPath {
									out = append(out, IpRouteResultFlat{
										Clientname: Rp.Clientname,
										Ifname:     Rp.Ifname,
										Metric:     Rp.Metric,
										Pref:       Rp.Pref,
										UBest:      Rp.UBest,
										UpTime:     ParseDuration(Rp.UpTime),
										Attached:   Rpre.Attached,
										IPPrefix:   Rpre.IPPrefix,
										McastNhops: Rpre.McastNhops,
										UcastNhops: Rpre.UcastNhops,
										AddRf:      Ra.AddRf,
										VrfNameOut: Rv.VrfNameOut,
									})
								}
							}
						}
					}
				}
			}
		}
	}
	return
}

type IpRouteResultFlat struct {
	Clientname string        `json:"clientname"`
	Ifname     string        `json:"ifname"`
	Metric     int           `json:"metric"`
	Pref       int           `json:"pref"`
	UBest      string        `json:"ubest"`
	UpTime     time.Duration `json:"uptime"`
	Attached   string        `json:"attached"`
	IPPrefix   string        `json:"ipprefix"`
	McastNhops int           `json:"mcast-nhops"`
	UcastNhops int           `json:"ucast-nhops"`
	AddRf      string        `json:"addrf"`
	VrfNameOut string        `json:"vrf-name-out"`
}

// NewIpRouteFromString returns SysInfo instance from an input string.
func NewIpRouteFromString(s string) (*IpRouteResponse, error) {
	return NewIpRouteFromReader(strings.NewReader(s))
}

// NewSysInfoFromBytes returns SysInfo instance from an input byte array.
func NewIpRouteFromBytes(s []byte) (*IpRouteResponse, error) {
	return NewIpRouteFromReader(bytes.NewReader(s))
}
func NewIpRouteFromReader(s io.Reader) (*IpRouteResponse, error) {
	//si := &IpRoute{}
	IpRouteResponseDat := &IpRouteResponse{}
	jsonDec := json.NewDecoder(s)
	//jsonDec.UseNumber()
	jsonDec.UseSlice()
	err := jsonDec.Decode(IpRouteResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return IpRouteResponseDat, nil
}
