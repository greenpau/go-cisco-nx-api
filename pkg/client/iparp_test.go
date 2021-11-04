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
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseShowIPARPJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *IpArpResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.ip.arp",
			exp: &IpArpResponse{
				InsAPI: struct {
					Outputs struct {
						Output IpArpResponseResult "json:\"output\""
					} "json:\"outputs\""
					Sid     string "json:\"sid\""
					Type    string "json:\"type\""
					Version string "json:\"version\""
				}{Outputs: struct {
					Output IpArpResponseResult "json:\"output\""
				}{Output: IpArpResponseResult{Body: IpArpResultBody{TableVrf: []struct {
					RowVrf []struct {
						TableAdj []struct {
							RowAdj []struct {
								Flags      string "json:\"flags\""
								IntfOut    string "json:\"intf-out\""
								IPAddrOut  string "json:\"ip-addr-out\""
								Mac        string "json:\"mac,omitempty\""
								TimeStamp  string "json:\"time-stamp\""
								Incomplete string "json:\"incomplete,omitempty\""
							} "json:\"ROW_adj\""
						} "json:\"TABLE_adj\""
						CntTotal   int    "json:\"cnt-total\""
						VrfNameOut string "json:\"vrf-name-out\""
					} "json:\"ROW_vrf\""
				}{struct {
					RowVrf []struct {
						TableAdj []struct {
							RowAdj []struct {
								Flags      string "json:\"flags\""
								IntfOut    string "json:\"intf-out\""
								IPAddrOut  string "json:\"ip-addr-out\""
								Mac        string "json:\"mac,omitempty\""
								TimeStamp  string "json:\"time-stamp\""
								Incomplete string "json:\"incomplete,omitempty\""
							} "json:\"ROW_adj\""
						} "json:\"TABLE_adj\""
						CntTotal   int    "json:\"cnt-total\""
						VrfNameOut string "json:\"vrf-name-out\""
					} "json:\"ROW_vrf\""
				}{RowVrf: []struct {
					TableAdj []struct {
						RowAdj []struct {
							Flags      string "json:\"flags\""
							IntfOut    string "json:\"intf-out\""
							IPAddrOut  string "json:\"ip-addr-out\""
							Mac        string "json:\"mac,omitempty\""
							TimeStamp  string "json:\"time-stamp\""
							Incomplete string "json:\"incomplete,omitempty\""
						} "json:\"ROW_adj\""
					} "json:\"TABLE_adj\""
					CntTotal   int    "json:\"cnt-total\""
					VrfNameOut string "json:\"vrf-name-out\""
				}{struct {
					TableAdj []struct {
						RowAdj []struct {
							Flags      string "json:\"flags\""
							IntfOut    string "json:\"intf-out\""
							IPAddrOut  string "json:\"ip-addr-out\""
							Mac        string "json:\"mac,omitempty\""
							TimeStamp  string "json:\"time-stamp\""
							Incomplete string "json:\"incomplete,omitempty\""
						} "json:\"ROW_adj\""
					} "json:\"TABLE_adj\""
					CntTotal   int    "json:\"cnt-total\""
					VrfNameOut string "json:\"vrf-name-out\""
				}{TableAdj: []struct {
					RowAdj []struct {
						Flags      string "json:\"flags\""
						IntfOut    string "json:\"intf-out\""
						IPAddrOut  string "json:\"ip-addr-out\""
						Mac        string "json:\"mac,omitempty\""
						TimeStamp  string "json:\"time-stamp\""
						Incomplete string "json:\"incomplete,omitempty\""
					} "json:\"ROW_adj\""
				}{struct {
					RowAdj []struct {
						Flags      string "json:\"flags\""
						IntfOut    string "json:\"intf-out\""
						IPAddrOut  string "json:\"ip-addr-out\""
						Mac        string "json:\"mac,omitempty\""
						TimeStamp  string "json:\"time-stamp\""
						Incomplete string "json:\"incomplete,omitempty\""
					} "json:\"ROW_adj\""
				}{RowAdj: []struct {
					Flags      string "json:\"flags\""
					IntfOut    string "json:\"intf-out\""
					IPAddrOut  string "json:\"ip-addr-out\""
					Mac        string "json:\"mac,omitempty\""
					TimeStamp  string "json:\"time-stamp\""
					Incomplete string "json:\"incomplete,omitempty\""
				}{struct {
					Flags      string "json:\"flags\""
					IntfOut    string "json:\"intf-out\""
					IPAddrOut  string "json:\"ip-addr-out\""
					Mac        string "json:\"mac,omitempty\""
					TimeStamp  string "json:\"time-stamp\""
					Incomplete string "json:\"incomplete,omitempty\""
				}{Flags: "", IntfOut: "Vlan253", IPAddrOut: "7.57.253.1", Mac: "f44e.0584.7ffc", TimeStamp: "PT18M21S", Incomplete: ""}, struct {
					Flags      string "json:\"flags\""
					IntfOut    string "json:\"intf-out\""
					IPAddrOut  string "json:\"ip-addr-out\""
					Mac        string "json:\"mac,omitempty\""
					TimeStamp  string "json:\"time-stamp\""
					Incomplete string "json:\"incomplete,omitempty\""
				}{Flags: "", IntfOut: "Vlan50", IPAddrOut: "10.57.50.3", Mac: "f44e.0584.7ffc", TimeStamp: "PT4M31S", Incomplete: ""}, struct {
					Flags      string "json:\"flags\""
					IntfOut    string "json:\"intf-out\""
					IPAddrOut  string "json:\"ip-addr-out\""
					Mac        string "json:\"mac,omitempty\""
					TimeStamp  string "json:\"time-stamp\""
					Incomplete string "json:\"incomplete,omitempty\""
				}{Flags: "", IntfOut: "Ethernet1/2", IPAddrOut: "10.100.157.1", Mac: "f44e.0584.7ffc", TimeStamp: "PT18M34S", Incomplete: ""}, struct {
					Flags      string "json:\"flags\""
					IntfOut    string "json:\"intf-out\""
					IPAddrOut  string "json:\"ip-addr-out\""
					Mac        string "json:\"mac,omitempty\""
					TimeStamp  string "json:\"time-stamp\""
					Incomplete string "json:\"incomplete,omitempty\""
				}{Flags: "", IntfOut: "Ethernet1/4", IPAddrOut: "10.100.157.9", Mac: "f44e.0584.7ffc", TimeStamp: "PT18M35S", Incomplete: ""}, struct {
					Flags      string "json:\"flags\""
					IntfOut    string "json:\"intf-out\""
					IPAddrOut  string "json:\"ip-addr-out\""
					Mac        string "json:\"mac,omitempty\""
					TimeStamp  string "json:\"time-stamp\""
					Incomplete string "json:\"incomplete,omitempty\""
				}{Flags: "", IntfOut: "Ethernet1/7", IPAddrOut: "192.168.161.1", Mac: "f44e.0584.7ffc", TimeStamp: "PT17M56S", Incomplete: ""}, struct {
					Flags      string "json:\"flags\""
					IntfOut    string "json:\"intf-out\""
					IPAddrOut  string "json:\"ip-addr-out\""
					Mac        string "json:\"mac,omitempty\""
					TimeStamp  string "json:\"time-stamp\""
					Incomplete string "json:\"incomplete,omitempty\""
				}{Flags: "", IntfOut: "Ethernet1/1.50", IPAddrOut: "89.1.1.10", Mac: "", TimeStamp: "PT28S", Incomplete: "true"}, struct {
					Flags      string "json:\"flags\""
					IntfOut    string "json:\"intf-out\""
					IPAddrOut  string "json:\"ip-addr-out\""
					Mac        string "json:\"mac,omitempty\""
					TimeStamp  string "json:\"time-stamp\""
					Incomplete string "json:\"incomplete,omitempty\""
				}{Flags: "", IntfOut: "Ethernet1/1.52", IPAddrOut: "89.1.3.10", Mac: "", TimeStamp: "PT4S", Incomplete: "true"}}}}, CntTotal: 7, VrfNameOut: "default"}}}}}, Code: "200", Input: "show ip arp ", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"},
			},
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
		iparp, err := NewIpArpFromBytes(content)
		//fmt.Printf("%#v\n", iparp) //DEBUG
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *iparp)
				testFailed++
				continue
			}
		}

		if iparp != nil {
			if !reflect.DeepEqual(test.exp, iparp) {
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
