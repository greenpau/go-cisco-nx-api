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

type IpRouteResponse struct {
	InsAPI struct {
		Outputs struct {
			Output IpRouteResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type IpRouteResponseResult struct {
	Body  IpRouteResultBody `json:"body" xml:"body"`
	Code  string            `json:"code" xml:"code"`
	Input string            `json:"input" xml:"input"`
	Msg   string            `json:"msg" xml:"msg"`
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
									Clientname string `json:"clientname" xml:"clientname"`
									Ifname     string `json:"ifname" xml:"ifname"`
									Metric     int    `json:"metric" xml:"metric"`
									Pref       int    `json:"pref" xml:"pref"`
									UBest      bool   `json:"ubest" xml:"ubest"`
									UpTime     string `json:"uptime" xml:"uptime"`
								} `json:"ROW_path" xml:"ROW_path"`
							} `json:"TABLE_path" xml:"TABLE_path"`
							Attached   bool   `json:"attached" xml:"attached"`
							IPPrefix   string `json:"ipprefix" xml:"ipprefix"`
							McastNhops int    `json:"mcast-nhops" xml:"mcast-nhops"`
							UcastNhops int    `json:"ucast-nhops" xml:"ucast-nhops"`
						} `json:"ROW_prefix" xml:"ROW_prefix"`
					} `json:"TABLE_prefix" xml:"TABLE_prefix"`
					AddRf string `json:"addrf" xml:"addrf"`
				} `json:"ROW_addrf" xml:"ROW_addrf"`
			} `json:"TABLE_addrf" xml:"TABLE_addrf"`
			VrfNameOut string `json:"vrf-name-out" xml:"vrf-name-out"`
		} `json:"ROW_vrf" xml:"ROW_vrf"`
	} `json:"TABLE_vrf" xml:"TABLE_vrf"`
}

type IpRouteResultFlat struct {
	Clientname string        `json:"clientname" xml:"clientname"`
	Ifname     string        `json:"ifname" xml:"ifname"`
	Metric     int           `json:"metric" xml:"metric"`
	Pref       int           `json:"pref" xml:"pref"`
	UBest      bool          `json:"ubest" xml:"ubest"`
	UpTime     time.Duration `json:"uptime" xml:"uptime"`
	Attached   bool          `json:"attached" xml:"attached"`
	IPPrefix   string        `json:"ipprefix" xml:"ipprefix"`
	McastNhops int           `json:"mcast-nhops" xml:"mcast-nhops"`
	UcastNhops int           `json:"ucast-nhops" xml:"ucast-nhops"`
	AddRf      string        `json:"addrf" xml:"addrf"`
	VrfNameOut string        `json:"vrf-name-out" xml:"vrf-name-out"`
}

func (d *IpRouteResponse) Flat() (out []IpRouteResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}
func (d *IpRouteResponseResult) Flat() (out []IpRouteResultFlat) {
	for _, Tv := range d.Body.TableVrf {
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

// NewIpRouteFromString returns instance from an input string.
func NewIpRouteFromString(s string) (*IpRouteResponse, error) {
	return NewIpRouteFromReader(strings.NewReader(s))
}

// NewIpRouteFromBytes returns instance from an input byte array.
func NewIpRouteFromBytes(s []byte) (*IpRouteResponse, error) {
	return NewIpRouteFromReader(bytes.NewReader(s))
}

// NewIpRouteFromReader returns instance from an input reader.
func NewIpRouteFromReader(s io.Reader) (*IpRouteResponse, error) {
	//si := &IpRoute{}
	IpRouteResponseDat := &IpRouteResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(IpRouteResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return IpRouteResponseDat, nil
}

// NewIpRouteResultFromString returns instance from an input string.
func NewIpRouteResultFromString(s string) (*IpRouteResponseResult, error) {
	return NewIpRouteResultFromReader(strings.NewReader(s))
}

// NewIpRouteResultFromBytes returns instance from an input byte array.
func NewIpRouteResultFromBytes(s []byte) (*IpRouteResponseResult, error) {
	return NewIpRouteResultFromReader(bytes.NewReader(s))
}

// NewIpRouteResultFromReader returns instance from an input reader.
func NewIpRouteResultFromReader(s io.Reader) (*IpRouteResponseResult, error) {
	//si := &IpRouteResponseResult{}
	IpRouteResponseResultDat := &IpRouteResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(IpRouteResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return IpRouteResponseResultDat, nil
}
