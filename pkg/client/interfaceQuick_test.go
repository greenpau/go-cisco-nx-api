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
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseShowInterfaceQuickJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *InterfaceQuickResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.interface.quick",
			exp: &InterfaceQuickResponse{
				InsAPI: struct {
					Outputs struct {
						Output InterfaceQuickResponseResult "json:\"output\" xml:\"output\""
					} "json:\"outputs\" xml:\"outputs\""
					Sid     string "json:\"sid\" xml:\"sid\""
					Type    string "json:\"type\" xml:\"type\""
					Version string "json:\"version\" xml:\"version\""
				}{Outputs: struct {
					Output InterfaceQuickResponseResult "json:\"output\" xml:\"output\""
				}{Output: InterfaceQuickResponseResult{Body: InterfaceQuickResultBody{TableInterface: []struct {
					RowInterface []struct {
						Interface         string   "json:\"interface\" xml:\"interface\""
						State             string   "json:\"state\" xml:\"state\""
						StateRsnDesc      string   "json:\"state_rsn_desc\" xml:\"state_rsn_desc\""
						AdminState        string   "json:\"admin_state\" xml:\"admin_state\""
						ShareState        string   "json:\"share_state\" xml:\"share_state\""
						EthHwDesc         string   "json:\"eth_hw_desc\" xml:\"eth_hw_desc\""
						EthHwAddr         string   "json:\"eth_hw_addr\" xml:\"eth_hw_addr\""
						EthBiaAddr        string   "json:\"eth_bia_addr\" xml:\"eth_bia_addr\""
						Desc              string   "json:\"desc\" xml:\"desc\""
						EthIPAddr         string   "json:\"eth_ip_addr\" xml:\"eth_ip_addr\""
						EthIPMask         int      "json:\"eth_ip_mask\" xml:\"eth_ip_mask\""
						EthMTU            int      "json:\"eth_mtu\" xml:\"eth_mtu\""
						EthBW             []string "json:\"eth_bw\" xml:\"eth_bw\""
						EthDly            int      "json:\"eth_dly\" xml:\"eth_dly\""
						EthReliability    int      "json:\"eth_reliability\" xml:\"eth_reliability\""
						EthTxLoad         int      "json:\"eth_txload\" xml:\"eth_txload\""
						EthRxLoad         int      "json:\"eth_rxload\" xml:\"eth_rxload\""
						Medium            string   "json:\"medium\" xml:\"medium\""
						EthDuplex         string   "json:\"eth_duplex\" xml:\"eth_duplex\""
						EthSpeed          string   "json:\"eth_speed\" xml:\"eth_speed\""
						EthMedia          string   "json:\"eth_media\" xml:\"eth_media\""
						EthBeacon         string   "json:\"eth_beacon\" xml:\"eth_beacon\""
						EthAutoNeg        string   "json:\"eth_autoneg\" xml:\"eth_autoneg\""
						EthInFlowCtrl     string   "json:\"eth_in_flowctrl\" xml:\"eth_in_flowctrl\""
						EthOutFlowCtrl    string   "json:\"eth_out_flowctrl\" xml:\"eth_out_flowctrl\""
						EthMdix           string   "json:\"eth_mdix\" xml:\"eth_mdix\""
						EthRateMode       string   "json:\"eth_ratemode\" xml:\"eth_ratemode\""
						EthSwrMonitor     string   "json:\"eth_swr_monitor\" xml:\"eth_swr_monitor\""
						EthEthertype      string   "json:\"eth_ethertype\" xml:\"eth_ethertype\""
						EthEeeState       string   "json:\"eth_eee_state\" xml:\"eth_eee_state\""
						Switchport        string   "json:\"switchport\" xml:\"switchport\""
						VdcLvlInPkts      int      "json:\"vdc_lvl_in_pkts\" xml:\"vdc_lvl_in_pkts\""
						VdcLvlInBytes     int      "json:\"vdc_lvl_in_bytes\" xml:\"vdc_lvl_in_bytes\""
						VdcLvlInUCast     int      "json:\"vdc_lvl_in_ucast\" xml:\"vdc_lvl_in_ucast\""
						VdcLvlInMCast     int      "json:\"vdc_lvl_in_mcast\" xml:\"vdc_lvl_in_mcast\""
						VdcLvlInBCast     int      "json:\"vdc_lvl_in_bcast\" xml:\"vdc_lvl_in_bcast\""
						VdcLvlInAvgPkts   int      "json:\"vdc_lvl_in_avg_pkts\" xml:\"vdc_lvl_in_avg_pkts\""
						VdcLvlInAvgBytes  int      "json:\"vdc_lvl_in_avg_bytes\" xml:\"vdc_lvl_in_avg_bytes\""
						VdcLvlOutPkts     int      "json:\"vdc_lvl_out_pkts\" xml:\"vdc_lvl_out_pkts\""
						VdcLvlOutBytes    int      "json:\"vdc_lvl_out_bytes\" xml:\"vdc_lvl_out_bytes\""
						VdcLvlOutUCast    int      "json:\"vdc_lvl_out_ucast\" xml:\"vdc_lvl_out_ucast\""
						VdcLvlOutMCast    int      "json:\"vdc_lvl_out_mcast\" xml:\"vdc_lvl_out_mcast\""
						VdcLvlOutBCast    int      "json:\"vdc_lvl_out_bcast\" xml:\"vdc_lvl_out_bcast\""
						VdcLvlOutAvgPkts  int      "json:\"vdc_lvl_out_avg_pkts\" xml:\"vdc_lvl_out_avg_pkts\""
						VdcLvlOutAvgBytes int      "json:\"vdc_lvl_out_avg_bytes\" xml:\"vdc_lvl_out_avg_bytes\""
					} "json:\"ROW_interface\" xml:\"ROW_interface\""
				}{struct {
					RowInterface []struct {
						Interface         string   "json:\"interface\" xml:\"interface\""
						State             string   "json:\"state\" xml:\"state\""
						StateRsnDesc      string   "json:\"state_rsn_desc\" xml:\"state_rsn_desc\""
						AdminState        string   "json:\"admin_state\" xml:\"admin_state\""
						ShareState        string   "json:\"share_state\" xml:\"share_state\""
						EthHwDesc         string   "json:\"eth_hw_desc\" xml:\"eth_hw_desc\""
						EthHwAddr         string   "json:\"eth_hw_addr\" xml:\"eth_hw_addr\""
						EthBiaAddr        string   "json:\"eth_bia_addr\" xml:\"eth_bia_addr\""
						Desc              string   "json:\"desc\" xml:\"desc\""
						EthIPAddr         string   "json:\"eth_ip_addr\" xml:\"eth_ip_addr\""
						EthIPMask         int      "json:\"eth_ip_mask\" xml:\"eth_ip_mask\""
						EthMTU            int      "json:\"eth_mtu\" xml:\"eth_mtu\""
						EthBW             []string "json:\"eth_bw\" xml:\"eth_bw\""
						EthDly            int      "json:\"eth_dly\" xml:\"eth_dly\""
						EthReliability    int      "json:\"eth_reliability\" xml:\"eth_reliability\""
						EthTxLoad         int      "json:\"eth_txload\" xml:\"eth_txload\""
						EthRxLoad         int      "json:\"eth_rxload\" xml:\"eth_rxload\""
						Medium            string   "json:\"medium\" xml:\"medium\""
						EthDuplex         string   "json:\"eth_duplex\" xml:\"eth_duplex\""
						EthSpeed          string   "json:\"eth_speed\" xml:\"eth_speed\""
						EthMedia          string   "json:\"eth_media\" xml:\"eth_media\""
						EthBeacon         string   "json:\"eth_beacon\" xml:\"eth_beacon\""
						EthAutoNeg        string   "json:\"eth_autoneg\" xml:\"eth_autoneg\""
						EthInFlowCtrl     string   "json:\"eth_in_flowctrl\" xml:\"eth_in_flowctrl\""
						EthOutFlowCtrl    string   "json:\"eth_out_flowctrl\" xml:\"eth_out_flowctrl\""
						EthMdix           string   "json:\"eth_mdix\" xml:\"eth_mdix\""
						EthRateMode       string   "json:\"eth_ratemode\" xml:\"eth_ratemode\""
						EthSwrMonitor     string   "json:\"eth_swr_monitor\" xml:\"eth_swr_monitor\""
						EthEthertype      string   "json:\"eth_ethertype\" xml:\"eth_ethertype\""
						EthEeeState       string   "json:\"eth_eee_state\" xml:\"eth_eee_state\""
						Switchport        string   "json:\"switchport\" xml:\"switchport\""
						VdcLvlInPkts      int      "json:\"vdc_lvl_in_pkts\" xml:\"vdc_lvl_in_pkts\""
						VdcLvlInBytes     int      "json:\"vdc_lvl_in_bytes\" xml:\"vdc_lvl_in_bytes\""
						VdcLvlInUCast     int      "json:\"vdc_lvl_in_ucast\" xml:\"vdc_lvl_in_ucast\""
						VdcLvlInMCast     int      "json:\"vdc_lvl_in_mcast\" xml:\"vdc_lvl_in_mcast\""
						VdcLvlInBCast     int      "json:\"vdc_lvl_in_bcast\" xml:\"vdc_lvl_in_bcast\""
						VdcLvlInAvgPkts   int      "json:\"vdc_lvl_in_avg_pkts\" xml:\"vdc_lvl_in_avg_pkts\""
						VdcLvlInAvgBytes  int      "json:\"vdc_lvl_in_avg_bytes\" xml:\"vdc_lvl_in_avg_bytes\""
						VdcLvlOutPkts     int      "json:\"vdc_lvl_out_pkts\" xml:\"vdc_lvl_out_pkts\""
						VdcLvlOutBytes    int      "json:\"vdc_lvl_out_bytes\" xml:\"vdc_lvl_out_bytes\""
						VdcLvlOutUCast    int      "json:\"vdc_lvl_out_ucast\" xml:\"vdc_lvl_out_ucast\""
						VdcLvlOutMCast    int      "json:\"vdc_lvl_out_mcast\" xml:\"vdc_lvl_out_mcast\""
						VdcLvlOutBCast    int      "json:\"vdc_lvl_out_bcast\" xml:\"vdc_lvl_out_bcast\""
						VdcLvlOutAvgPkts  int      "json:\"vdc_lvl_out_avg_pkts\" xml:\"vdc_lvl_out_avg_pkts\""
						VdcLvlOutAvgBytes int      "json:\"vdc_lvl_out_avg_bytes\" xml:\"vdc_lvl_out_avg_bytes\""
					} "json:\"ROW_interface\" xml:\"ROW_interface\""
				}{RowInterface: []struct {
					Interface         string   "json:\"interface\" xml:\"interface\""
					State             string   "json:\"state\" xml:\"state\""
					StateRsnDesc      string   "json:\"state_rsn_desc\" xml:\"state_rsn_desc\""
					AdminState        string   "json:\"admin_state\" xml:\"admin_state\""
					ShareState        string   "json:\"share_state\" xml:\"share_state\""
					EthHwDesc         string   "json:\"eth_hw_desc\" xml:\"eth_hw_desc\""
					EthHwAddr         string   "json:\"eth_hw_addr\" xml:\"eth_hw_addr\""
					EthBiaAddr        string   "json:\"eth_bia_addr\" xml:\"eth_bia_addr\""
					Desc              string   "json:\"desc\" xml:\"desc\""
					EthIPAddr         string   "json:\"eth_ip_addr\" xml:\"eth_ip_addr\""
					EthIPMask         int      "json:\"eth_ip_mask\" xml:\"eth_ip_mask\""
					EthMTU            int      "json:\"eth_mtu\" xml:\"eth_mtu\""
					EthBW             []string "json:\"eth_bw\" xml:\"eth_bw\""
					EthDly            int      "json:\"eth_dly\" xml:\"eth_dly\""
					EthReliability    int      "json:\"eth_reliability\" xml:\"eth_reliability\""
					EthTxLoad         int      "json:\"eth_txload\" xml:\"eth_txload\""
					EthRxLoad         int      "json:\"eth_rxload\" xml:\"eth_rxload\""
					Medium            string   "json:\"medium\" xml:\"medium\""
					EthDuplex         string   "json:\"eth_duplex\" xml:\"eth_duplex\""
					EthSpeed          string   "json:\"eth_speed\" xml:\"eth_speed\""
					EthMedia          string   "json:\"eth_media\" xml:\"eth_media\""
					EthBeacon         string   "json:\"eth_beacon\" xml:\"eth_beacon\""
					EthAutoNeg        string   "json:\"eth_autoneg\" xml:\"eth_autoneg\""
					EthInFlowCtrl     string   "json:\"eth_in_flowctrl\" xml:\"eth_in_flowctrl\""
					EthOutFlowCtrl    string   "json:\"eth_out_flowctrl\" xml:\"eth_out_flowctrl\""
					EthMdix           string   "json:\"eth_mdix\" xml:\"eth_mdix\""
					EthRateMode       string   "json:\"eth_ratemode\" xml:\"eth_ratemode\""
					EthSwrMonitor     string   "json:\"eth_swr_monitor\" xml:\"eth_swr_monitor\""
					EthEthertype      string   "json:\"eth_ethertype\" xml:\"eth_ethertype\""
					EthEeeState       string   "json:\"eth_eee_state\" xml:\"eth_eee_state\""
					Switchport        string   "json:\"switchport\" xml:\"switchport\""
					VdcLvlInPkts      int      "json:\"vdc_lvl_in_pkts\" xml:\"vdc_lvl_in_pkts\""
					VdcLvlInBytes     int      "json:\"vdc_lvl_in_bytes\" xml:\"vdc_lvl_in_bytes\""
					VdcLvlInUCast     int      "json:\"vdc_lvl_in_ucast\" xml:\"vdc_lvl_in_ucast\""
					VdcLvlInMCast     int      "json:\"vdc_lvl_in_mcast\" xml:\"vdc_lvl_in_mcast\""
					VdcLvlInBCast     int      "json:\"vdc_lvl_in_bcast\" xml:\"vdc_lvl_in_bcast\""
					VdcLvlInAvgPkts   int      "json:\"vdc_lvl_in_avg_pkts\" xml:\"vdc_lvl_in_avg_pkts\""
					VdcLvlInAvgBytes  int      "json:\"vdc_lvl_in_avg_bytes\" xml:\"vdc_lvl_in_avg_bytes\""
					VdcLvlOutPkts     int      "json:\"vdc_lvl_out_pkts\" xml:\"vdc_lvl_out_pkts\""
					VdcLvlOutBytes    int      "json:\"vdc_lvl_out_bytes\" xml:\"vdc_lvl_out_bytes\""
					VdcLvlOutUCast    int      "json:\"vdc_lvl_out_ucast\" xml:\"vdc_lvl_out_ucast\""
					VdcLvlOutMCast    int      "json:\"vdc_lvl_out_mcast\" xml:\"vdc_lvl_out_mcast\""
					VdcLvlOutBCast    int      "json:\"vdc_lvl_out_bcast\" xml:\"vdc_lvl_out_bcast\""
					VdcLvlOutAvgPkts  int      "json:\"vdc_lvl_out_avg_pkts\" xml:\"vdc_lvl_out_avg_pkts\""
					VdcLvlOutAvgBytes int      "json:\"vdc_lvl_out_avg_bytes\" xml:\"vdc_lvl_out_avg_bytes\""
				}{struct {
					Interface         string   "json:\"interface\" xml:\"interface\""
					State             string   "json:\"state\" xml:\"state\""
					StateRsnDesc      string   "json:\"state_rsn_desc\" xml:\"state_rsn_desc\""
					AdminState        string   "json:\"admin_state\" xml:\"admin_state\""
					ShareState        string   "json:\"share_state\" xml:\"share_state\""
					EthHwDesc         string   "json:\"eth_hw_desc\" xml:\"eth_hw_desc\""
					EthHwAddr         string   "json:\"eth_hw_addr\" xml:\"eth_hw_addr\""
					EthBiaAddr        string   "json:\"eth_bia_addr\" xml:\"eth_bia_addr\""
					Desc              string   "json:\"desc\" xml:\"desc\""
					EthIPAddr         string   "json:\"eth_ip_addr\" xml:\"eth_ip_addr\""
					EthIPMask         int      "json:\"eth_ip_mask\" xml:\"eth_ip_mask\""
					EthMTU            int      "json:\"eth_mtu\" xml:\"eth_mtu\""
					EthBW             []string "json:\"eth_bw\" xml:\"eth_bw\""
					EthDly            int      "json:\"eth_dly\" xml:\"eth_dly\""
					EthReliability    int      "json:\"eth_reliability\" xml:\"eth_reliability\""
					EthTxLoad         int      "json:\"eth_txload\" xml:\"eth_txload\""
					EthRxLoad         int      "json:\"eth_rxload\" xml:\"eth_rxload\""
					Medium            string   "json:\"medium\" xml:\"medium\""
					EthDuplex         string   "json:\"eth_duplex\" xml:\"eth_duplex\""
					EthSpeed          string   "json:\"eth_speed\" xml:\"eth_speed\""
					EthMedia          string   "json:\"eth_media\" xml:\"eth_media\""
					EthBeacon         string   "json:\"eth_beacon\" xml:\"eth_beacon\""
					EthAutoNeg        string   "json:\"eth_autoneg\" xml:\"eth_autoneg\""
					EthInFlowCtrl     string   "json:\"eth_in_flowctrl\" xml:\"eth_in_flowctrl\""
					EthOutFlowCtrl    string   "json:\"eth_out_flowctrl\" xml:\"eth_out_flowctrl\""
					EthMdix           string   "json:\"eth_mdix\" xml:\"eth_mdix\""
					EthRateMode       string   "json:\"eth_ratemode\" xml:\"eth_ratemode\""
					EthSwrMonitor     string   "json:\"eth_swr_monitor\" xml:\"eth_swr_monitor\""
					EthEthertype      string   "json:\"eth_ethertype\" xml:\"eth_ethertype\""
					EthEeeState       string   "json:\"eth_eee_state\" xml:\"eth_eee_state\""
					Switchport        string   "json:\"switchport\" xml:\"switchport\""
					VdcLvlInPkts      int      "json:\"vdc_lvl_in_pkts\" xml:\"vdc_lvl_in_pkts\""
					VdcLvlInBytes     int      "json:\"vdc_lvl_in_bytes\" xml:\"vdc_lvl_in_bytes\""
					VdcLvlInUCast     int      "json:\"vdc_lvl_in_ucast\" xml:\"vdc_lvl_in_ucast\""
					VdcLvlInMCast     int      "json:\"vdc_lvl_in_mcast\" xml:\"vdc_lvl_in_mcast\""
					VdcLvlInBCast     int      "json:\"vdc_lvl_in_bcast\" xml:\"vdc_lvl_in_bcast\""
					VdcLvlInAvgPkts   int      "json:\"vdc_lvl_in_avg_pkts\" xml:\"vdc_lvl_in_avg_pkts\""
					VdcLvlInAvgBytes  int      "json:\"vdc_lvl_in_avg_bytes\" xml:\"vdc_lvl_in_avg_bytes\""
					VdcLvlOutPkts     int      "json:\"vdc_lvl_out_pkts\" xml:\"vdc_lvl_out_pkts\""
					VdcLvlOutBytes    int      "json:\"vdc_lvl_out_bytes\" xml:\"vdc_lvl_out_bytes\""
					VdcLvlOutUCast    int      "json:\"vdc_lvl_out_ucast\" xml:\"vdc_lvl_out_ucast\""
					VdcLvlOutMCast    int      "json:\"vdc_lvl_out_mcast\" xml:\"vdc_lvl_out_mcast\""
					VdcLvlOutBCast    int      "json:\"vdc_lvl_out_bcast\" xml:\"vdc_lvl_out_bcast\""
					VdcLvlOutAvgPkts  int      "json:\"vdc_lvl_out_avg_pkts\" xml:\"vdc_lvl_out_avg_pkts\""
					VdcLvlOutAvgBytes int      "json:\"vdc_lvl_out_avg_bytes\" xml:\"vdc_lvl_out_avg_bytes\""
				}{Interface: "mgmt0", State: "up", StateRsnDesc: "Link not connected", AdminState: "up", ShareState: "Dedicated", EthHwDesc: "GigabitEthernet", EthHwAddr: "<Redacted>", EthBiaAddr: "<Redacted>", Desc: "<Redacted>", EthIPAddr: "<Redacted>", EthIPMask: 24, EthMTU: 1500, EthBW: []string{"1000000", "1000000"}, EthDly: 10, EthReliability: 255, EthTxLoad: 1, EthRxLoad: 1, Medium: "broadcast", EthDuplex: "full", EthSpeed: "auto-speed", EthMedia: "1G", EthBeacon: "off", EthAutoNeg: "on", EthInFlowCtrl: "off", EthOutFlowCtrl: "off", EthMdix: "off", EthRateMode: "dedicated", EthSwrMonitor: "off", EthEthertype: "0x0000", EthEeeState: "n/a", Switchport: "Disabled", VdcLvlInPkts: 354499, VdcLvlInBytes: 32251676, VdcLvlInUCast: 337649, VdcLvlInMCast: 16842, VdcLvlInBCast: 8, VdcLvlInAvgPkts: 1, VdcLvlInAvgBytes: 744, VdcLvlOutPkts: 354493, VdcLvlOutBytes: 32218753, VdcLvlOutUCast: 337641, VdcLvlOutMCast: 16850, VdcLvlOutBCast: 2, VdcLvlOutAvgPkts: 1, VdcLvlOutAvgBytes: 768}}}}}, Code: "200", Input: "show interface status", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.2"}},
			shouldFail: false,
			shouldErr:  false,
		},
	} {
		fp := fmt.Sprintf("%s/resp.%s.json", outputDir, test.input)
		content, err := ioutil.ReadFile(fp)
		if err != nil {
			t.Logf("FAIL: Test %d: failed reading '%s', error: %v", i, fp, err)
			testFailed++
			continue
		}
		dat, err := NewInterfaceQuickFromBytes(content)
		//fmt.Printf("%#v\n", dat) //DEBUG
		//fmt.Printf("%#v\n", dat.Flat()) //DEBUG
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *dat)
				testFailed++
				continue
			}
		}

		if dat != nil {
			if !reflect.DeepEqual(test.exp, dat) {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to mismatch", i, test.input)
				testFailed++
			}
		}

		if test.shouldFail {
			t.Logf("PASS: Test %d: input '%s', expected to fail, failed", i, test.input)
		} else {
			t.Logf("PASS: Test %d: input '%s', expected to pass, passed", i, test.input)
		}
	}
	if testFailed > 0 {
		t.Fatalf("Failed %d tests", testFailed)
	}
}
func TestParseShowInterfaceQuickResultJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *InterfaceQuickResponseResult
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "result.show.interface.quick",
			exp: &InterfaceQuickResponseResult{
				Body: client.InterfaceQuickResultBody{TableInterface: []struct {
					RowInterface []struct {
						Interface         string   "json:\"interface\" xml:\"interface\""
						State             string   "json:\"state\" xml:\"state\""
						StateRsnDesc      string   "json:\"state_rsn_desc\" xml:\"state_rsn_desc\""
						AdminState        string   "json:\"admin_state\" xml:\"admin_state\""
						ShareState        string   "json:\"share_state\" xml:\"share_state\""
						EthHwDesc         string   "json:\"eth_hw_desc\" xml:\"eth_hw_desc\""
						EthHwAddr         string   "json:\"eth_hw_addr\" xml:\"eth_hw_addr\""
						EthBiaAddr        string   "json:\"eth_bia_addr\" xml:\"eth_bia_addr\""
						Desc              string   "json:\"desc\" xml:\"desc\""
						EthIPAddr         string   "json:\"eth_ip_addr\" xml:\"eth_ip_addr\""
						EthIPMask         int      "json:\"eth_ip_mask\" xml:\"eth_ip_mask\""
						EthMTU            int      "json:\"eth_mtu\" xml:\"eth_mtu\""
						EthBW             []string "json:\"eth_bw\" xml:\"eth_bw\""
						EthDly            int      "json:\"eth_dly\" xml:\"eth_dly\""
						EthReliability    int      "json:\"eth_reliability\" xml:\"eth_reliability\""
						EthTxLoad         int      "json:\"eth_txload\" xml:\"eth_txload\""
						EthRxLoad         int      "json:\"eth_rxload\" xml:\"eth_rxload\""
						Medium            string   "json:\"medium\" xml:\"medium\""
						EthDuplex         string   "json:\"eth_duplex\" xml:\"eth_duplex\""
						EthSpeed          string   "json:\"eth_speed\" xml:\"eth_speed\""
						EthMedia          string   "json:\"eth_media\" xml:\"eth_media\""
						EthBeacon         string   "json:\"eth_beacon\" xml:\"eth_beacon\""
						EthAutoNeg        string   "json:\"eth_autoneg\" xml:\"eth_autoneg\""
						EthInFlowCtrl     string   "json:\"eth_in_flowctrl\" xml:\"eth_in_flowctrl\""
						EthOutFlowCtrl    string   "json:\"eth_out_flowctrl\" xml:\"eth_out_flowctrl\""
						EthMdix           string   "json:\"eth_mdix\" xml:\"eth_mdix\""
						EthRateMode       string   "json:\"eth_ratemode\" xml:\"eth_ratemode\""
						EthSwrMonitor     string   "json:\"eth_swr_monitor\" xml:\"eth_swr_monitor\""
						EthEthertype      string   "json:\"eth_ethertype\" xml:\"eth_ethertype\""
						EthEeeState       string   "json:\"eth_eee_state\" xml:\"eth_eee_state\""
						Switchport        string   "json:\"switchport\" xml:\"switchport\""
						VdcLvlInPkts      int      "json:\"vdc_lvl_in_pkts\" xml:\"vdc_lvl_in_pkts\""
						VdcLvlInBytes     int      "json:\"vdc_lvl_in_bytes\" xml:\"vdc_lvl_in_bytes\""
						VdcLvlInUCast     int      "json:\"vdc_lvl_in_ucast\" xml:\"vdc_lvl_in_ucast\""
						VdcLvlInMCast     int      "json:\"vdc_lvl_in_mcast\" xml:\"vdc_lvl_in_mcast\""
						VdcLvlInBCast     int      "json:\"vdc_lvl_in_bcast\" xml:\"vdc_lvl_in_bcast\""
						VdcLvlInAvgPkts   int      "json:\"vdc_lvl_in_avg_pkts\" xml:\"vdc_lvl_in_avg_pkts\""
						VdcLvlInAvgBytes  int      "json:\"vdc_lvl_in_avg_bytes\" xml:\"vdc_lvl_in_avg_bytes\""
						VdcLvlOutPkts     int      "json:\"vdc_lvl_out_pkts\" xml:\"vdc_lvl_out_pkts\""
						VdcLvlOutBytes    int      "json:\"vdc_lvl_out_bytes\" xml:\"vdc_lvl_out_bytes\""
						VdcLvlOutUCast    int      "json:\"vdc_lvl_out_ucast\" xml:\"vdc_lvl_out_ucast\""
						VdcLvlOutMCast    int      "json:\"vdc_lvl_out_mcast\" xml:\"vdc_lvl_out_mcast\""
						VdcLvlOutBCast    int      "json:\"vdc_lvl_out_bcast\" xml:\"vdc_lvl_out_bcast\""
						VdcLvlOutAvgPkts  int      "json:\"vdc_lvl_out_avg_pkts\" xml:\"vdc_lvl_out_avg_pkts\""
						VdcLvlOutAvgBytes int      "json:\"vdc_lvl_out_avg_bytes\" xml:\"vdc_lvl_out_avg_bytes\""
					} "json:\"ROW_interface\" xml:\"ROW_interface\""
				}{struct {
					RowInterface []struct {
						Interface         string   "json:\"interface\" xml:\"interface\""
						State             string   "json:\"state\" xml:\"state\""
						StateRsnDesc      string   "json:\"state_rsn_desc\" xml:\"state_rsn_desc\""
						AdminState        string   "json:\"admin_state\" xml:\"admin_state\""
						ShareState        string   "json:\"share_state\" xml:\"share_state\""
						EthHwDesc         string   "json:\"eth_hw_desc\" xml:\"eth_hw_desc\""
						EthHwAddr         string   "json:\"eth_hw_addr\" xml:\"eth_hw_addr\""
						EthBiaAddr        string   "json:\"eth_bia_addr\" xml:\"eth_bia_addr\""
						Desc              string   "json:\"desc\" xml:\"desc\""
						EthIPAddr         string   "json:\"eth_ip_addr\" xml:\"eth_ip_addr\""
						EthIPMask         int      "json:\"eth_ip_mask\" xml:\"eth_ip_mask\""
						EthMTU            int      "json:\"eth_mtu\" xml:\"eth_mtu\""
						EthBW             []string "json:\"eth_bw\" xml:\"eth_bw\""
						EthDly            int      "json:\"eth_dly\" xml:\"eth_dly\""
						EthReliability    int      "json:\"eth_reliability\" xml:\"eth_reliability\""
						EthTxLoad         int      "json:\"eth_txload\" xml:\"eth_txload\""
						EthRxLoad         int      "json:\"eth_rxload\" xml:\"eth_rxload\""
						Medium            string   "json:\"medium\" xml:\"medium\""
						EthDuplex         string   "json:\"eth_duplex\" xml:\"eth_duplex\""
						EthSpeed          string   "json:\"eth_speed\" xml:\"eth_speed\""
						EthMedia          string   "json:\"eth_media\" xml:\"eth_media\""
						EthBeacon         string   "json:\"eth_beacon\" xml:\"eth_beacon\""
						EthAutoNeg        string   "json:\"eth_autoneg\" xml:\"eth_autoneg\""
						EthInFlowCtrl     string   "json:\"eth_in_flowctrl\" xml:\"eth_in_flowctrl\""
						EthOutFlowCtrl    string   "json:\"eth_out_flowctrl\" xml:\"eth_out_flowctrl\""
						EthMdix           string   "json:\"eth_mdix\" xml:\"eth_mdix\""
						EthRateMode       string   "json:\"eth_ratemode\" xml:\"eth_ratemode\""
						EthSwrMonitor     string   "json:\"eth_swr_monitor\" xml:\"eth_swr_monitor\""
						EthEthertype      string   "json:\"eth_ethertype\" xml:\"eth_ethertype\""
						EthEeeState       string   "json:\"eth_eee_state\" xml:\"eth_eee_state\""
						Switchport        string   "json:\"switchport\" xml:\"switchport\""
						VdcLvlInPkts      int      "json:\"vdc_lvl_in_pkts\" xml:\"vdc_lvl_in_pkts\""
						VdcLvlInBytes     int      "json:\"vdc_lvl_in_bytes\" xml:\"vdc_lvl_in_bytes\""
						VdcLvlInUCast     int      "json:\"vdc_lvl_in_ucast\" xml:\"vdc_lvl_in_ucast\""
						VdcLvlInMCast     int      "json:\"vdc_lvl_in_mcast\" xml:\"vdc_lvl_in_mcast\""
						VdcLvlInBCast     int      "json:\"vdc_lvl_in_bcast\" xml:\"vdc_lvl_in_bcast\""
						VdcLvlInAvgPkts   int      "json:\"vdc_lvl_in_avg_pkts\" xml:\"vdc_lvl_in_avg_pkts\""
						VdcLvlInAvgBytes  int      "json:\"vdc_lvl_in_avg_bytes\" xml:\"vdc_lvl_in_avg_bytes\""
						VdcLvlOutPkts     int      "json:\"vdc_lvl_out_pkts\" xml:\"vdc_lvl_out_pkts\""
						VdcLvlOutBytes    int      "json:\"vdc_lvl_out_bytes\" xml:\"vdc_lvl_out_bytes\""
						VdcLvlOutUCast    int      "json:\"vdc_lvl_out_ucast\" xml:\"vdc_lvl_out_ucast\""
						VdcLvlOutMCast    int      "json:\"vdc_lvl_out_mcast\" xml:\"vdc_lvl_out_mcast\""
						VdcLvlOutBCast    int      "json:\"vdc_lvl_out_bcast\" xml:\"vdc_lvl_out_bcast\""
						VdcLvlOutAvgPkts  int      "json:\"vdc_lvl_out_avg_pkts\" xml:\"vdc_lvl_out_avg_pkts\""
						VdcLvlOutAvgBytes int      "json:\"vdc_lvl_out_avg_bytes\" xml:\"vdc_lvl_out_avg_bytes\""
					} "json:\"ROW_interface\" xml:\"ROW_interface\""
				}{RowInterface: []struct {
					Interface         string   "json:\"interface\" xml:\"interface\""
					State             string   "json:\"state\" xml:\"state\""
					StateRsnDesc      string   "json:\"state_rsn_desc\" xml:\"state_rsn_desc\""
					AdminState        string   "json:\"admin_state\" xml:\"admin_state\""
					ShareState        string   "json:\"share_state\" xml:\"share_state\""
					EthHwDesc         string   "json:\"eth_hw_desc\" xml:\"eth_hw_desc\""
					EthHwAddr         string   "json:\"eth_hw_addr\" xml:\"eth_hw_addr\""
					EthBiaAddr        string   "json:\"eth_bia_addr\" xml:\"eth_bia_addr\""
					Desc              string   "json:\"desc\" xml:\"desc\""
					EthIPAddr         string   "json:\"eth_ip_addr\" xml:\"eth_ip_addr\""
					EthIPMask         int      "json:\"eth_ip_mask\" xml:\"eth_ip_mask\""
					EthMTU            int      "json:\"eth_mtu\" xml:\"eth_mtu\""
					EthBW             []string "json:\"eth_bw\" xml:\"eth_bw\""
					EthDly            int      "json:\"eth_dly\" xml:\"eth_dly\""
					EthReliability    int      "json:\"eth_reliability\" xml:\"eth_reliability\""
					EthTxLoad         int      "json:\"eth_txload\" xml:\"eth_txload\""
					EthRxLoad         int      "json:\"eth_rxload\" xml:\"eth_rxload\""
					Medium            string   "json:\"medium\" xml:\"medium\""
					EthDuplex         string   "json:\"eth_duplex\" xml:\"eth_duplex\""
					EthSpeed          string   "json:\"eth_speed\" xml:\"eth_speed\""
					EthMedia          string   "json:\"eth_media\" xml:\"eth_media\""
					EthBeacon         string   "json:\"eth_beacon\" xml:\"eth_beacon\""
					EthAutoNeg        string   "json:\"eth_autoneg\" xml:\"eth_autoneg\""
					EthInFlowCtrl     string   "json:\"eth_in_flowctrl\" xml:\"eth_in_flowctrl\""
					EthOutFlowCtrl    string   "json:\"eth_out_flowctrl\" xml:\"eth_out_flowctrl\""
					EthMdix           string   "json:\"eth_mdix\" xml:\"eth_mdix\""
					EthRateMode       string   "json:\"eth_ratemode\" xml:\"eth_ratemode\""
					EthSwrMonitor     string   "json:\"eth_swr_monitor\" xml:\"eth_swr_monitor\""
					EthEthertype      string   "json:\"eth_ethertype\" xml:\"eth_ethertype\""
					EthEeeState       string   "json:\"eth_eee_state\" xml:\"eth_eee_state\""
					Switchport        string   "json:\"switchport\" xml:\"switchport\""
					VdcLvlInPkts      int      "json:\"vdc_lvl_in_pkts\" xml:\"vdc_lvl_in_pkts\""
					VdcLvlInBytes     int      "json:\"vdc_lvl_in_bytes\" xml:\"vdc_lvl_in_bytes\""
					VdcLvlInUCast     int      "json:\"vdc_lvl_in_ucast\" xml:\"vdc_lvl_in_ucast\""
					VdcLvlInMCast     int      "json:\"vdc_lvl_in_mcast\" xml:\"vdc_lvl_in_mcast\""
					VdcLvlInBCast     int      "json:\"vdc_lvl_in_bcast\" xml:\"vdc_lvl_in_bcast\""
					VdcLvlInAvgPkts   int      "json:\"vdc_lvl_in_avg_pkts\" xml:\"vdc_lvl_in_avg_pkts\""
					VdcLvlInAvgBytes  int      "json:\"vdc_lvl_in_avg_bytes\" xml:\"vdc_lvl_in_avg_bytes\""
					VdcLvlOutPkts     int      "json:\"vdc_lvl_out_pkts\" xml:\"vdc_lvl_out_pkts\""
					VdcLvlOutBytes    int      "json:\"vdc_lvl_out_bytes\" xml:\"vdc_lvl_out_bytes\""
					VdcLvlOutUCast    int      "json:\"vdc_lvl_out_ucast\" xml:\"vdc_lvl_out_ucast\""
					VdcLvlOutMCast    int      "json:\"vdc_lvl_out_mcast\" xml:\"vdc_lvl_out_mcast\""
					VdcLvlOutBCast    int      "json:\"vdc_lvl_out_bcast\" xml:\"vdc_lvl_out_bcast\""
					VdcLvlOutAvgPkts  int      "json:\"vdc_lvl_out_avg_pkts\" xml:\"vdc_lvl_out_avg_pkts\""
					VdcLvlOutAvgBytes int      "json:\"vdc_lvl_out_avg_bytes\" xml:\"vdc_lvl_out_avg_bytes\""
				}{struct {
					Interface         string   "json:\"interface\" xml:\"interface\""
					State             string   "json:\"state\" xml:\"state\""
					StateRsnDesc      string   "json:\"state_rsn_desc\" xml:\"state_rsn_desc\""
					AdminState        string   "json:\"admin_state\" xml:\"admin_state\""
					ShareState        string   "json:\"share_state\" xml:\"share_state\""
					EthHwDesc         string   "json:\"eth_hw_desc\" xml:\"eth_hw_desc\""
					EthHwAddr         string   "json:\"eth_hw_addr\" xml:\"eth_hw_addr\""
					EthBiaAddr        string   "json:\"eth_bia_addr\" xml:\"eth_bia_addr\""
					Desc              string   "json:\"desc\" xml:\"desc\""
					EthIPAddr         string   "json:\"eth_ip_addr\" xml:\"eth_ip_addr\""
					EthIPMask         int      "json:\"eth_ip_mask\" xml:\"eth_ip_mask\""
					EthMTU            int      "json:\"eth_mtu\" xml:\"eth_mtu\""
					EthBW             []string "json:\"eth_bw\" xml:\"eth_bw\""
					EthDly            int      "json:\"eth_dly\" xml:\"eth_dly\""
					EthReliability    int      "json:\"eth_reliability\" xml:\"eth_reliability\""
					EthTxLoad         int      "json:\"eth_txload\" xml:\"eth_txload\""
					EthRxLoad         int      "json:\"eth_rxload\" xml:\"eth_rxload\""
					Medium            string   "json:\"medium\" xml:\"medium\""
					EthDuplex         string   "json:\"eth_duplex\" xml:\"eth_duplex\""
					EthSpeed          string   "json:\"eth_speed\" xml:\"eth_speed\""
					EthMedia          string   "json:\"eth_media\" xml:\"eth_media\""
					EthBeacon         string   "json:\"eth_beacon\" xml:\"eth_beacon\""
					EthAutoNeg        string   "json:\"eth_autoneg\" xml:\"eth_autoneg\""
					EthInFlowCtrl     string   "json:\"eth_in_flowctrl\" xml:\"eth_in_flowctrl\""
					EthOutFlowCtrl    string   "json:\"eth_out_flowctrl\" xml:\"eth_out_flowctrl\""
					EthMdix           string   "json:\"eth_mdix\" xml:\"eth_mdix\""
					EthRateMode       string   "json:\"eth_ratemode\" xml:\"eth_ratemode\""
					EthSwrMonitor     string   "json:\"eth_swr_monitor\" xml:\"eth_swr_monitor\""
					EthEthertype      string   "json:\"eth_ethertype\" xml:\"eth_ethertype\""
					EthEeeState       string   "json:\"eth_eee_state\" xml:\"eth_eee_state\""
					Switchport        string   "json:\"switchport\" xml:\"switchport\""
					VdcLvlInPkts      int      "json:\"vdc_lvl_in_pkts\" xml:\"vdc_lvl_in_pkts\""
					VdcLvlInBytes     int      "json:\"vdc_lvl_in_bytes\" xml:\"vdc_lvl_in_bytes\""
					VdcLvlInUCast     int      "json:\"vdc_lvl_in_ucast\" xml:\"vdc_lvl_in_ucast\""
					VdcLvlInMCast     int      "json:\"vdc_lvl_in_mcast\" xml:\"vdc_lvl_in_mcast\""
					VdcLvlInBCast     int      "json:\"vdc_lvl_in_bcast\" xml:\"vdc_lvl_in_bcast\""
					VdcLvlInAvgPkts   int      "json:\"vdc_lvl_in_avg_pkts\" xml:\"vdc_lvl_in_avg_pkts\""
					VdcLvlInAvgBytes  int      "json:\"vdc_lvl_in_avg_bytes\" xml:\"vdc_lvl_in_avg_bytes\""
					VdcLvlOutPkts     int      "json:\"vdc_lvl_out_pkts\" xml:\"vdc_lvl_out_pkts\""
					VdcLvlOutBytes    int      "json:\"vdc_lvl_out_bytes\" xml:\"vdc_lvl_out_bytes\""
					VdcLvlOutUCast    int      "json:\"vdc_lvl_out_ucast\" xml:\"vdc_lvl_out_ucast\""
					VdcLvlOutMCast    int      "json:\"vdc_lvl_out_mcast\" xml:\"vdc_lvl_out_mcast\""
					VdcLvlOutBCast    int      "json:\"vdc_lvl_out_bcast\" xml:\"vdc_lvl_out_bcast\""
					VdcLvlOutAvgPkts  int      "json:\"vdc_lvl_out_avg_pkts\" xml:\"vdc_lvl_out_avg_pkts\""
					VdcLvlOutAvgBytes int      "json:\"vdc_lvl_out_avg_bytes\" xml:\"vdc_lvl_out_avg_bytes\""
				}{Interface: "mgmt0", State: "up", StateRsnDesc: "Link not connected", AdminState: "up", ShareState: "Dedicated", EthHwDesc: "GigabitEthernet", EthHwAddr: "<Redacted>", EthBiaAddr: "<Redacted>", Desc: "<Redacted>", EthIPAddr: "<Redacted>", EthIPMask: 24, EthMTU: 1500, EthBW: []string{"1000000", "1000000"}, EthDly: 10, EthReliability: 255, EthTxLoad: 1, EthRxLoad: 1, Medium: "broadcast", EthDuplex: "full", EthSpeed: "auto-speed", EthMedia: "1G", EthBeacon: "off", EthAutoNeg: "on", EthInFlowCtrl: "off", EthOutFlowCtrl: "off", EthMdix: "off", EthRateMode: "dedicated", EthSwrMonitor: "off", EthEthertype: "0x0000", EthEeeState: "n/a", Switchport: "Disabled", VdcLvlInPkts: 354499, VdcLvlInBytes: 32251676, VdcLvlInUCast: 337649, VdcLvlInMCast: 16842, VdcLvlInBCast: 8, VdcLvlInAvgPkts: 1, VdcLvlInAvgBytes: 744, VdcLvlOutPkts: 354493, VdcLvlOutBytes: 32218753, VdcLvlOutUCast: 337641, VdcLvlOutMCast: 16850, VdcLvlOutBCast: 2, VdcLvlOutAvgPkts: 1, VdcLvlOutAvgBytes: 768}}}}}, Code: "200", Input: "show interface status", Msg: "Success"},
			shouldFail: false,
			shouldErr:  false,
		},
	} {
		fp := fmt.Sprintf("%s/resp.%s.json", outputDir, test.input)
		content, err := ioutil.ReadFile(fp)
		if err != nil {
			t.Logf("FAIL: Test %d: failed reading '%s', error: %v", i, fp, err)
			testFailed++
			continue
		}
		dat, err := NewInterfaceQuickResultFromBytes(content)
		fmt.Printf("%#v\n", dat) //DEBUG
		//fmt.Printf("%#v\n", dat.Flat()) //DEBUG
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *dat)
				testFailed++
				continue
			}
		}

		if dat != nil {
			if !reflect.DeepEqual(test.exp, dat) {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to mismatch", i, test.input)
				testFailed++
			}
		}

		if test.shouldFail {
			t.Logf("PASS: Test %d: input '%s', expected to fail, failed", i, test.input)
		} else {
			t.Logf("PASS: Test %d: input '%s', expected to pass, passed", i, test.input)
		}
	}
	if testFailed > 0 {
		t.Fatalf("Failed %d tests", testFailed)
	}
}
