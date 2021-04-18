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

func TestParseShowInterfaceStatusJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *InterfaceStatusResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.interface.status",
			exp: &InterfaceStatusResponse{
				InsAPI: struct {
					Outputs struct {
						Output InterfaceStatusResponseResult "json:\"output\" xml:\"output\""
					} "json:\"outputs\" xml:\"outputs\""
					Sid     string "json:\"sid\" xml:\"sid\""
					Type    string "json:\"type\" xml:\"type\""
					Version string "json:\"version\" xml:\"version\""
				}{Outputs: struct {
					Output InterfaceStatusResponseResult "json:\"output\" xml:\"output\""
				}{Output: InterfaceStatusResponseResult{Body: InterfaceStatusResultBody{TableInterface: []struct {
					RowInterface []struct {
						Interface string "json:\"interface\" xml:\"interface\""
						State     string "json:\"state\" xml:\"state\""
						Vlan      string "json:\"vlan\" xml:\"vlan\""
						Duplex    string "json:\"duplex\" xml:\"duplex\""
						Speed     string "json:\"speed\" xml:\"speed\""
						Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
					} "json:\"ROW_interface\" xml:\"ROW_interface\""
				}{struct {
					RowInterface []struct {
						Interface string "json:\"interface\" xml:\"interface\""
						State     string "json:\"state\" xml:\"state\""
						Vlan      string "json:\"vlan\" xml:\"vlan\""
						Duplex    string "json:\"duplex\" xml:\"duplex\""
						Speed     string "json:\"speed\" xml:\"speed\""
						Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
					} "json:\"ROW_interface\" xml:\"ROW_interface\""
				}{RowInterface: []struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "mgmt0", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-1000", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/1", State: "notconnect", Vlan: "2", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/2", State: "connected", Vlan: "2", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/3", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/4", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/5", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU1M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/6", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/7", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/8", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/9", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/10", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/11", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/12", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/13", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/14", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/15", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/16", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU1M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/17", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/18", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/19", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/20", State: "connected", Vlan: "routed", Duplex: "full", Speed: "1000", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/21", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/22", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/23", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/24", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/25", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/26", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/27", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/28", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/29", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/30", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/31", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/32", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/33", State: "connected", Vlan: "2", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/34", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/35", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/36", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/37", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU1M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/38", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/39", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/40", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/41", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/42", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU1M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/43", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/44", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/45", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/46", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/47", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/48", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel12", State: "noOperMembers", Vlan: "1", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel22", State: "connected", Vlan: "2", Duplex: "full", Speed: "a-10G", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel63", State: "noOperMembers", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel123", State: "noOperMembers", Vlan: "1", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel154", State: "suspndByVpc", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "loopback22", State: "connected", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "loopback63", State: "connected", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Vlan1", State: "down", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: ""}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Vlan2", State: "connected", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: ""}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Vlan63", State: "down", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: ""}}}}}, Code: "200", Input: "show interface status", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.2"}},
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
		dat, err := NewInterfaceStatusFromBytes(content)
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

func TestParseShowInterfaceStatusResultJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *InterfaceStatusResponseResult
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "result.show.interface.status",
			exp: &InterfaceStatusResponseResult{
				Body: InterfaceStatusResultBody{TableInterface: []struct {
					RowInterface []struct {
						Interface string "json:\"interface\" xml:\"interface\""
						State     string "json:\"state\" xml:\"state\""
						Vlan      string "json:\"vlan\" xml:\"vlan\""
						Duplex    string "json:\"duplex\" xml:\"duplex\""
						Speed     string "json:\"speed\" xml:\"speed\""
						Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
					} "json:\"ROW_interface\" xml:\"ROW_interface\""
				}{struct {
					RowInterface []struct {
						Interface string "json:\"interface\" xml:\"interface\""
						State     string "json:\"state\" xml:\"state\""
						Vlan      string "json:\"vlan\" xml:\"vlan\""
						Duplex    string "json:\"duplex\" xml:\"duplex\""
						Speed     string "json:\"speed\" xml:\"speed\""
						Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
					} "json:\"ROW_interface\" xml:\"ROW_interface\""
				}{RowInterface: []struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "mgmt0", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-1000", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/1", State: "notconnect", Vlan: "2", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/2", State: "connected", Vlan: "2", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/3", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/4", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/5", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU1M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/6", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/7", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/8", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/9", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/10", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/11", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/12", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/13", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/14", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/15", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/16", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU1M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/17", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/18", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/19", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/20", State: "connected", Vlan: "routed", Duplex: "full", Speed: "1000", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/21", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/22", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/23", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/24", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/25", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/26", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/27", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/28", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/29", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/30", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/31", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/32", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/33", State: "connected", Vlan: "2", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/34", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/35", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/36", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU5M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/37", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU1M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/38", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/39", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/40", State: "sfpAbsent", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/41", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/42", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU1M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/43", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/44", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "SFP-H10GB-CU3M"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/45", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/46", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/47", State: "connected", Vlan: "routed", Duplex: "full", Speed: "a-10G", Type: "10Gbase-SR"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Ethernet9/48", State: "notconnect", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "1000base-T"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel12", State: "noOperMembers", Vlan: "1", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel22", State: "connected", Vlan: "2", Duplex: "full", Speed: "a-10G", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel63", State: "noOperMembers", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel123", State: "noOperMembers", Vlan: "1", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "port-channel154", State: "suspndByVpc", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "loopback22", State: "connected", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "loopback63", State: "connected", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: "--"}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Vlan1", State: "down", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: ""}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Vlan2", State: "connected", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: ""}, struct {
					Interface string "json:\"interface\" xml:\"interface\""
					State     string "json:\"state\" xml:\"state\""
					Vlan      string "json:\"vlan\" xml:\"vlan\""
					Duplex    string "json:\"duplex\" xml:\"duplex\""
					Speed     string "json:\"speed\" xml:\"speed\""
					Type      string "json:\"type,omitempty\" xml:\"type,omitempty\""
				}{Interface: "Vlan63", State: "down", Vlan: "routed", Duplex: "auto", Speed: "auto", Type: ""}}}}}, Code: "200", Input: "show interface status", Msg: "Success"},
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
		dat, err := NewInterfaceStatusResultFromBytes(content)
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
