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
)

type InterfaceQuickResponse struct {
	InsAPI struct {
		Outputs struct {
			Output InterfaceQuickResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type InterfaceQuickResponseResult struct {
	Body  InterfaceQuickResultBody `json:"body" xml:"body"`
	Code  string                   `json:"code" xml:"code"`
	Input string                   `json:"input" xml:"input"`
	Msg   string                   `json:"msg" xml:"msg"`
}

type InterfaceQuickResultBody struct {
	TableInterface []struct {
		RowInterface []struct {
			Interface         string   `json:"interface" xml:"interface"`
			State             string   `json:"state" xml:"state"`
			StateRsnDesc      string   `json:"state_rsn_desc" xml:"state_rsn_desc"`
			AdminState        string   `json:"admin_state" xml:"admin_state"`
			ShareState        string   `json:"share_state" xml:"share_state"`
			EthHwDesc         string   `json:"eth_hw_desc" xml:"eth_hw_desc"`
			EthHwAddr         string   `json:"eth_hw_addr" xml:"eth_hw_addr"`
			EthBiaAddr        string   `json:"eth_bia_addr" xml:"eth_bia_addr"`
			Desc              string   `json:"desc" xml:"desc"`
			EthIPAddr         string   `json:"eth_ip_addr" xml:"eth_ip_addr"`
			EthIPMask         int      `json:"eth_ip_mask" xml:"eth_ip_mask"`
			EthMTU            int      `json:"eth_mtu" xml:"eth_mtu"`
			EthBW             []string `json:"eth_bw" xml:"eth_bw"`
			EthDly            int      `json:"eth_dly" xml:"eth_dly"`
			EthReliability    int      `json:"eth_reliability" xml:"eth_reliability"`
			EthTxLoad         int      `json:"eth_txload" xml:"eth_txload"`
			EthRxLoad         int      `json:"eth_rxload" xml:"eth_rxload"`
			Medium            string   `json:"medium" xml:"medium"`
			EthDuplex         string   `json:"eth_duplex" xml:"eth_duplex"`
			EthSpeed          string   `json:"eth_speed" xml:"eth_speed"`
			EthMedia          string   `json:"eth_media" xml:"eth_media"`
			EthBeacon         string   `json:"eth_beacon" xml:"eth_beacon"`
			EthAutoNeg        string   `json:"eth_autoneg" xml:"eth_autoneg"`
			EthInFlowCtrl     string   `json:"eth_in_flowctrl" xml:"eth_in_flowctrl"`
			EthOutFlowCtrl    string   `json:"eth_out_flowctrl" xml:"eth_out_flowctrl"`
			EthMdix           string   `json:"eth_mdix" xml:"eth_mdix"`
			EthRateMode       string   `json:"eth_ratemode" xml:"eth_ratemode"`
			EthSwrMonitor     string   `json:"eth_swr_monitor" xml:"eth_swr_monitor"`
			EthEthertype      string   `json:"eth_ethertype" xml:"eth_ethertype"`
			EthEeeState       string   `json:"eth_eee_state" xml:"eth_eee_state"`
			Switchport        string   `json:"switchport" xml:"switchport"`
			VdcLvlInPkts      int      `json:"vdc_lvl_in_pkts" xml:"vdc_lvl_in_pkts"`
			VdcLvlInBytes     int      `json:"vdc_lvl_in_bytes" xml:"vdc_lvl_in_bytes"`
			VdcLvlInUCast     int      `json:"vdc_lvl_in_ucast" xml:"vdc_lvl_in_ucast"`
			VdcLvlInMCast     int      `json:"vdc_lvl_in_mcast" xml:"vdc_lvl_in_mcast"`
			VdcLvlInBCast     int      `json:"vdc_lvl_in_bcast" xml:"vdc_lvl_in_bcast"`
			VdcLvlInAvgPkts   int      `json:"vdc_lvl_in_avg_pkts" xml:"vdc_lvl_in_avg_pkts"`
			VdcLvlInAvgBytes  int      `json:"vdc_lvl_in_avg_bytes" xml:"vdc_lvl_in_avg_bytes"`
			VdcLvlOutPkts     int      `json:"vdc_lvl_out_pkts" xml:"vdc_lvl_out_pkts"`
			VdcLvlOutBytes    int      `json:"vdc_lvl_out_bytes" xml:"vdc_lvl_out_bytes"`
			VdcLvlOutUCast    int      `json:"vdc_lvl_out_ucast" xml:"vdc_lvl_out_ucast"`
			VdcLvlOutMCast    int      `json:"vdc_lvl_out_mcast" xml:"vdc_lvl_out_mcast"`
			VdcLvlOutBCast    int      `json:"vdc_lvl_out_bcast" xml:"vdc_lvl_out_bcast"`
			VdcLvlOutAvgPkts  int      `json:"vdc_lvl_out_avg_pkts" xml:"vdc_lvl_out_avg_pkts"`
			VdcLvlOutAvgBytes int      `json:"vdc_lvl_out_avg_bytes" xml:"vdc_lvl_out_avg_bytes"`
		} `json:"ROW_interface" xml:"ROW_interface"`
	} `json:"TABLE_interface" xml:"TABLE_interface"`
}

type InterfaceQuickResultFlat struct {
	Interface         string   `json:"interface" xml:"interface"`
	State             string   `json:"state" xml:"state"`
	StateRsnDesc      string   `json:"state_rsn_desc" xml:"state_rsn_desc"`
	AdminState        string   `json:"admin_state" xml:"admin_state"`
	ShareState        string   `json:"share_state" xml:"share_state"`
	EthHwDesc         string   `json:"eth_hw_desc" xml:"eth_hw_desc"`
	EthHwAddr         string   `json:"eth_hw_addr" xml:"eth_hw_addr"`
	EthBiaAddr        string   `json:"eth_bia_addr" xml:"eth_bia_addr"`
	Desc              string   `json:"desc" xml:"desc"`
	EthIPAddr         string   `json:"eth_ip_addr" xml:"eth_ip_addr"`
	EthIPMask         int      `json:"eth_ip_mask" xml:"eth_ip_mask"`
	EthMTU            int      `json:"eth_mtu" xml:"eth_mtu"`
	EthBW             []string `json:"eth_bw" xml:"eth_bw"`
	EthDly            int      `json:"eth_dly" xml:"eth_dly"`
	EthReliability    int      `json:"eth_reliability" xml:"eth_reliability"`
	EthTxLoad         int      `json:"eth_txload" xml:"eth_txload"`
	EthRxLoad         int      `json:"eth_rxload" xml:"eth_rxload"`
	Medium            string   `json:"medium" xml:"medium"`
	EthDuplex         string   `json:"eth_duplex" xml:"eth_duplex"`
	EthSpeed          string   `json:"eth_speed" xml:"eth_speed"`
	EthMedia          string   `json:"eth_media" xml:"eth_media"`
	EthBeacon         string   `json:"eth_beacon" xml:"eth_beacon"`
	EthAutoNeg        string   `json:"eth_autoneg" xml:"eth_autoneg"`
	EthInFlowCtrl     string   `json:"eth_in_flowctrl" xml:"eth_in_flowctrl"`
	EthOutFlowCtrl    string   `json:"eth_out_flowctrl" xml:"eth_out_flowctrl"`
	EthMdix           string   `json:"eth_mdix" xml:"eth_mdix"`
	EthRateMode       string   `json:"eth_ratemode" xml:"eth_ratemode"`
	EthSwrMonitor     string   `json:"eth_swr_monitor" xml:"eth_swr_monitor"`
	EthEthertype      string   `json:"eth_ethertype" xml:"eth_ethertype"`
	EthEeeState       string   `json:"eth_eee_state" xml:"eth_eee_state"`
	Switchport        string   `json:"switchport" xml:"switchport"`
	VdcLvlInPkts      int      `json:"vdc_lvl_in_pkts" xml:"vdc_lvl_in_pkts"`
	VdcLvlInBytes     int      `json:"vdc_lvl_in_bytes" xml:"vdc_lvl_in_bytes"`
	VdcLvlInUCast     int      `json:"vdc_lvl_in_ucast" xml:"vdc_lvl_in_ucast"`
	VdcLvlInMCast     int      `json:"vdc_lvl_in_mcast" xml:"vdc_lvl_in_mcast"`
	VdcLvlInBCast     int      `json:"vdc_lvl_in_bcast" xml:"vdc_lvl_in_bcast"`
	VdcLvlInAvgPkts   int      `json:"vdc_lvl_in_avg_pkts" xml:"vdc_lvl_in_avg_pkts"`
	VdcLvlInAvgBytes  int      `json:"vdc_lvl_in_avg_bytes" xml:"vdc_lvl_in_avg_bytes"`
	VdcLvlOutPkts     int      `json:"vdc_lvl_out_pkts" xml:"vdc_lvl_out_pkts"`
	VdcLvlOutBytes    int      `json:"vdc_lvl_out_bytes" xml:"vdc_lvl_out_bytes"`
	VdcLvlOutUCast    int      `json:"vdc_lvl_out_ucast" xml:"vdc_lvl_out_ucast"`
	VdcLvlOutMCast    int      `json:"vdc_lvl_out_mcast" xml:"vdc_lvl_out_mcast"`
	VdcLvlOutBCast    int      `json:"vdc_lvl_out_bcast" xml:"vdc_lvl_out_bcast"`
	VdcLvlOutAvgPkts  int      `json:"vdc_lvl_out_avg_pkts" xml:"vdc_lvl_out_avg_pkts"`
	VdcLvlOutAvgBytes int      `json:"vdc_lvl_out_avg_bytes" xml:"vdc_lvl_out_avg_bytes"`
}

func (d *InterfaceQuickResponse) Flat() (out []InterfaceQuickResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}
func (d *InterfaceQuickResponseResult) Flat() (out []InterfaceQuickResultFlat) {
	for _, Ti := range d.Body.TableInterface {
		for _, Ri := range Ti.RowInterface {
			out = append(out, InterfaceQuickResultFlat{
				Interface:         Ri.Interface,
				State:             Ri.State,
				StateRsnDesc:      Ri.StateRsnDesc,
				AdminState:        Ri.AdminState,
				ShareState:        Ri.ShareState,
				EthHwDesc:         Ri.EthHwDesc,
				EthHwAddr:         Ri.EthHwAddr,
				EthBiaAddr:        Ri.EthBiaAddr,
				Desc:              Ri.Desc,
				EthIPAddr:         Ri.EthIPAddr,
				EthIPMask:         Ri.EthIPMask,
				EthMTU:            Ri.EthMTU,
				EthBW:             Ri.EthBW,
				EthDly:            Ri.EthDly,
				EthReliability:    Ri.EthReliability,
				EthTxLoad:         Ri.EthTxLoad,
				EthRxLoad:         Ri.EthRxLoad,
				Medium:            Ri.Medium,
				EthDuplex:         Ri.EthDuplex,
				EthSpeed:          Ri.EthSpeed,
				EthMedia:          Ri.EthMedia,
				EthBeacon:         Ri.EthBeacon,
				EthAutoNeg:        Ri.EthAutoNeg,
				EthInFlowCtrl:     Ri.EthInFlowCtrl,
				EthOutFlowCtrl:    Ri.EthOutFlowCtrl,
				EthMdix:           Ri.EthMdix,
				EthRateMode:       Ri.EthRateMode,
				EthSwrMonitor:     Ri.EthSwrMonitor,
				EthEthertype:      Ri.EthEthertype,
				EthEeeState:       Ri.EthEeeState,
				Switchport:        Ri.Switchport,
				VdcLvlInPkts:      Ri.VdcLvlInPkts,
				VdcLvlInBytes:     Ri.VdcLvlInBytes,
				VdcLvlInUCast:     Ri.VdcLvlInUCast,
				VdcLvlInMCast:     Ri.VdcLvlInMCast,
				VdcLvlInBCast:     Ri.VdcLvlInBCast,
				VdcLvlInAvgPkts:   Ri.VdcLvlInAvgPkts,
				VdcLvlInAvgBytes:  Ri.VdcLvlInAvgBytes,
				VdcLvlOutPkts:     Ri.VdcLvlOutPkts,
				VdcLvlOutBytes:    Ri.VdcLvlOutBytes,
				VdcLvlOutUCast:    Ri.VdcLvlOutUCast,
				VdcLvlOutMCast:    Ri.VdcLvlOutMCast,
				VdcLvlOutBCast:    Ri.VdcLvlOutBCast,
				VdcLvlOutAvgPkts:  Ri.VdcLvlOutAvgPkts,
				VdcLvlOutAvgBytes: Ri.VdcLvlOutAvgBytes,
			})
		}
	}
	return
}

// NewInterfaceQuickFromString returns instance from an input string.
func NewInterfaceQuickFromString(s string) (*InterfaceQuickResponse, error) {
	return NewInterfaceQuickFromReader(strings.NewReader(s))
}

// NewInterfaceQuickFromBytes returns instance from an input byte array.
func NewInterfaceQuickFromBytes(s []byte) (*InterfaceQuickResponse, error) {
	return NewInterfaceQuickFromReader(bytes.NewReader(s))
}

// NewInterfaceQuickFromReader returns instance from an input reader.
func NewInterfaceQuickFromReader(s io.Reader) (*InterfaceQuickResponse, error) {
	//si := &InterfaceQuick{}
	InterfaceQuickResponseDat := &InterfaceQuickResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(InterfaceQuickResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return InterfaceQuickResponseDat, nil
}

// NewInterfaceQuickResultFromString returns instance from an input string.
func NewInterfaceQuickResultFromString(s string) (*InterfaceQuickResponseResult, error) {
	return NewInterfaceQuickResultFromReader(strings.NewReader(s))
}

// NewInterfaceQuickResultFromBytes returns instance from an input byte array.
func NewInterfaceQuickResultFromBytes(s []byte) (*InterfaceQuickResponseResult, error) {
	return NewInterfaceQuickResultFromReader(bytes.NewReader(s))
}

// NewInterfaceQuickResultFromReader returns instance from an input reader.
func NewInterfaceQuickResultFromReader(s io.Reader) (*InterfaceQuickResponseResult, error) {
	//si := &InterfaceQuickResponseResult{}
	InterfaceQuickResponseResultDat := &InterfaceQuickResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(InterfaceQuickResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return InterfaceQuickResponseResultDat, nil
}
