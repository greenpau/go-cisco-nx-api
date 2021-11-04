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

func TestParseShowInterfaceBriefJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *InterfaceBriefResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.interface.brief",
			exp: &InterfaceBriefResponse{
				TableInterface: struct {
					RowInterface []struct {
						Interface    string "json:\"interface\""
						State        string "json:\"state,omitempty\""
						IPAddr       string "json:\"ip_addr,omitempty\""
						Speed        string "json:\"speed,omitempty\""
						MTU          int    "json:\"mtu,omitempty\""
						Vlan         string "json:\"vlan,omitempty\""
						Type         string "json:\"type,omitempty\""
						PortMode     string "json:\"portmode,omitempty\""
						StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
						RateMode     string "json:\"ratemode,omitempty\""
						PortChan     int    "json:\"portchan,omitempty\""
						Proto        string "json:\"proto,omitempty\""
						OperState    string "json:\"oper_state,omitempty\""
						SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
					} "json:\"ROW_interface\""
				}{RowInterface: []struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "mgmt0", State: "up", IPAddr: "172.22.199.43", Speed: "1000", MTU: 1500, Vlan: "", Type: "", PortMode: "", StateRsnDesc: "", RateMode: "", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet2/33", State: "down", IPAddr: "", Speed: "auto", MTU: 0, Vlan: "--", Type: "eth", PortMode: "routed", StateRsnDesc: "SFP not inserted", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet2/34", State: "down", IPAddr: "", Speed: "auto", MTU: 0, Vlan: "--", Type: "eth", PortMode: "routed", StateRsnDesc: "SFP not inserted", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet2/35", State: "down", IPAddr: "", Speed: "auto", MTU: 0, Vlan: "--", Type: "eth", PortMode: "routed", StateRsnDesc: "SFP not inserted", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet2/36", State: "down", IPAddr: "", Speed: "auto", MTU: 0, Vlan: "--", Type: "eth", PortMode: "routed", StateRsnDesc: "SFP not inserted", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet2/37", State: "down", IPAddr: "", Speed: "auto", MTU: 0, Vlan: "--", Type: "eth", PortMode: "routed", StateRsnDesc: "SFP not inserted", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet2/38", State: "down", IPAddr: "", Speed: "auto", MTU: 0, Vlan: "--", Type: "eth", PortMode: "routed", StateRsnDesc: "SFP not inserted", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet2/39", State: "down", IPAddr: "", Speed: "auto", MTU: 0, Vlan: "--", Type: "eth", PortMode: "routed", StateRsnDesc: "SFP not inserted", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet2/40", State: "down", IPAddr: "", Speed: "auto", MTU: 0, Vlan: "--", Type: "eth", PortMode: "routed", StateRsnDesc: "SFP not inserted", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/1", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "1", Type: "eth", PortMode: "trunk", StateRsnDesc: "none", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/2", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "70", Type: "eth", PortMode: "access", StateRsnDesc: "none", RateMode: "D", PortChan: 128, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/3", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "1", Type: "eth", PortMode: "trunk", StateRsnDesc: "none", RateMode: "D", PortChan: 256, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/10", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "10", Type: "eth", PortMode: "access", StateRsnDesc: "none", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/11", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "70", Type: "eth", PortMode: "access", StateRsnDesc: "none", RateMode: "D", PortChan: 128, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/12", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "1", Type: "eth", PortMode: "trunk", StateRsnDesc: "none", RateMode: "D", PortChan: 256, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/25", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "1", Type: "eth", PortMode: "trunk", StateRsnDesc: "none", RateMode: "D", PortChan: 700, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/26", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "1", Type: "eth", PortMode: "trunk", StateRsnDesc: "none", RateMode: "D", PortChan: 0, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Ethernet10/27", State: "up", IPAddr: "", Speed: "100G", MTU: 0, Vlan: "40", Type: "eth", PortMode: "access", StateRsnDesc: "none", RateMode: "D", PortChan: 800, Proto: "", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "port-channel128", State: "up", IPAddr: "", Speed: "a-100G", MTU: 0, Vlan: "70", Type: "eth", PortMode: "access", StateRsnDesc: "none", RateMode: "D", PortChan: 0, Proto: "lacp", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "port-channel256", State: "up", IPAddr: "", Speed: "a-100G", MTU: 0, Vlan: "1", Type: "eth", PortMode: "trunk", StateRsnDesc: "none", RateMode: "D", PortChan: 0, Proto: "lacp", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "port-channel700", State: "up", IPAddr: "", Speed: "a-100G", MTU: 0, Vlan: "1", Type: "eth", PortMode: "trunk", StateRsnDesc: "none", RateMode: "D", PortChan: 0, Proto: "lacp", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "port-channel800", State: "up", IPAddr: "", Speed: "a-100G", MTU: 0, Vlan: "40", Type: "eth", PortMode: "access", StateRsnDesc: "none", RateMode: "D", PortChan: 0, Proto: "lacp", OperState: "", SviRsnDesc: ""}, struct {
					Interface    string "json:\"interface\""
					State        string "json:\"state,omitempty\""
					IPAddr       string "json:\"ip_addr,omitempty\""
					Speed        string "json:\"speed,omitempty\""
					MTU          int    "json:\"mtu,omitempty\""
					Vlan         string "json:\"vlan,omitempty\""
					Type         string "json:\"type,omitempty\""
					PortMode     string "json:\"portmode,omitempty\""
					StateRsnDesc string "json:\"state_rsn_desc,omitempty\""
					RateMode     string "json:\"ratemode,omitempty\""
					PortChan     int    "json:\"portchan,omitempty\""
					Proto        string "json:\"proto,omitempty\""
					OperState    string "json:\"oper_state,omitempty\""
					SviRsnDesc   string "json:\"svi_rsn_desc,omitempty\""
				}{Interface: "Vlan1", State: "", IPAddr: "", Speed: "", MTU: 0, Vlan: "", Type: "svi", PortMode: "", StateRsnDesc: "", RateMode: "", PortChan: 0, Proto: "", OperState: "down", SviRsnDesc: "Administratively down"}}}},
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
		dat, err := NewInterfaceBriefFromBytes(content)
		//fmt.Printf("%#v\n", dat) //DEBUG
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
