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
	"encoding/json"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"strings"
)

type vlanResponse struct {
	ID      uint64             `json:"id" xml:"id"`
	Version string             `json:"jsonrpc" xml:"jsonrpc"`
	Result  vlanResponseResult `json:"result" xml:"result"`
}

type vlanResponseResult struct {
	Body vlanResponseResultBody `json:"body" xml:"body"`
}

type vlanResponseResultBody struct {
	VlanBriefTable vlanResponseResultBodyVlanBriefTable `json:"TABLE_vlanbrief" xml:"TABLE_vlanbrief"`
	MtuInfoTable   vlanResponseResultBodyMtuInfoTable   `json:"TABLE_mtuinfo" xml:"TABLE_mtuinfo"`
}

type vlanResponseResultBodyVlanBriefTable struct {
	VlanBriefRow []vlanResponseResultBodyVlanBriefRow `json:"ROW_vlanbrief" xml:"ROW_vlanbrief"`
}

func (t *vlanResponseResultBodyVlanBriefTable) UnmarshalJSON(b []byte) error {
	size := len(b)
	i := bytes.IndexByte(b, byte(':'))
	if i < 0 {
		return fmt.Errorf("Error unmarshalling vlanResponseResultBodyVlanBriefTable")
	}
	j := bytes.IndexByte(b[i:], byte('['))
	switch {
	case j > 10 || j < 0:
		// single entry
		var r vlanResponseResultBodyVlanBriefRow
		err := json.Unmarshal(b[(i+2):size-1], &r)
		if err != nil {
			return fmt.Errorf("Error unmarshalling vlanResponseResultBodyVlanBriefTable: %s", err)
		}
		t.VlanBriefRow = append(t.VlanBriefRow, r)
	case j < 10 && j >= 0:
		// multiple entries
		var r []vlanResponseResultBodyVlanBriefRow
		err := json.Unmarshal(b[(i+2):size-1], &r)
		if err != nil {
			return fmt.Errorf("Error unmarshalling vlanResponseResultBodyVlanBriefTable: %s", err)
		}
		t.VlanBriefRow = r
	}
	return nil
}

type vlanResponseResultBodyMtuInfoTable struct {
	MtuInfoRow []vlanResponseResultBodyMtuInfoRow `json:"ROW_mtuinfo" xml:"ROW_mtuinfo"`
}

func (t *vlanResponseResultBodyMtuInfoTable) UnmarshalJSON(b []byte) error {
	size := len(b)
	i := bytes.IndexByte(b, byte(':'))
	if i < 0 {
		return fmt.Errorf("Error unmarshalling vlanResponseResultBodyMtuInfoTable")
	}
	j := bytes.IndexByte(b[i:], byte('['))
	switch {
	case j > 10 || j < 0:
		// single entry
		var r vlanResponseResultBodyMtuInfoRow
		err := json.Unmarshal(b[(i+2):size-1], &r)
		if err != nil {
			return fmt.Errorf("Error unmarshalling vlanResponseResultBodyMtuInfoTable: %s", err)
		}
		t.MtuInfoRow = append(t.MtuInfoRow, r)
	case j < 10 && j >= 0:
		// multiple entries
		var r []vlanResponseResultBodyMtuInfoRow
		err := json.Unmarshal(b[(i+2):size-1], &r)
		if err != nil {
			return fmt.Errorf("Error unmarshalling vlanResponseResultBodyMtuInfoTable: %s", err)
		}
		t.MtuInfoRow = r
	}
	return nil
}

type vlanResponseResultBodyVlanBriefRow struct {
	ID            string                                  `json:"vlanshowbr-vlanid" xml:"vlanshowbr-vlanid"`
	UtfID         string                                  `json:"vlanshowbr-vlanid-utf" xml:"vlanshowbr-vlanid-utf"`
	Name          string                                  `json:"vlanshowbr-vlanname" xml:"vlanshowbr-vlanname"`
	State         string                                  `json:"vlanshowbr-vlanstate" xml:"vlanshowbr-vlanstate"`
	ShutdownState string                                  `json:"vlanshowbr-shutstate" xml:"vlanshowbr-shutstate"`
	Ports         vlanResponseResultBodyVlanBriefRowPorts `json:"vlanshowplist-ifidx" xml:"vlanshowplist-ifidx"`
}

type vlanResponseResultBodyVlanBriefRowPorts struct {
	Entries []string
}

func (t *vlanResponseResultBodyVlanBriefRowPorts) UnmarshalJSON(b []byte) error {
	if bytes.HasPrefix(b, []byte("[")) {
		entries := []string{}
		err := json.Unmarshal(b, &entries)
		if err != nil {
			return fmt.Errorf("Error unmarshalling vlanResponseResultBodyVlanBriefRowPorts: %s", err)
		}
		for _, entry := range entries {
			for _, port := range strings.Split(entry, ",") {
				if t.Entries == nil {
					t.Entries = []string{}
				}
				t.Entries = append(t.Entries, port)
			}
		}
	} else {
		var entries string
		err := json.Unmarshal(b, &entries)
		if err != nil {
			return fmt.Errorf("Error unmarshalling vlanResponseResultBodyVlanBriefRowPorts: %s", err)
		}
		for _, port := range strings.Split(entries, ",") {
			if t.Entries == nil {
				t.Entries = []string{}
			}
			t.Entries = append(t.Entries, port)
		}
	}
	return nil
}

type vlanResponseResultBodyMtuInfoRow struct {
	ID        string `json:"vlanshowinfo-vlanid" xml:"vlanshowinfo-vlanid"`
	MediaType string `json:"vlanshowinfo-media-type" xml:"vlanshowinfo-media-type"`
	Mode      string `json:"vlanshowinfo-vlanmode" xml:"vlanshowinfo-vlanmode"`
}

// Vlan contains system information. The information in the structure
// is from the output of "show vlan" command.
type Vlan struct {
	ID            string   `json:"id" xml:"id"`
	Name          string   `json:"name" xml:"name"`
	State         string   `json:"state" xml:"state"`
	ShutdownState string   `json:"shutdown_state" xml:"shutdown_state"`
	Ports         []string `json:"ports" xml:"ports"`
	MediaType     string   `json:"media_type" xml:"media_type"`
	Mode          string   `json:"mode" xml:"mode"`
}

// NewVlansFromString returns Vlan instance from an input string.
func NewVlansFromString(s string) ([]*Vlan, error) {
	return NewVlansFromBytes([]byte(s))
}

// NewVlansFromBytes returns Vlan instance from an input byte array.
func NewVlansFromBytes(s []byte) ([]*Vlan, error) {
	var vlans []*Vlan
	vResponse := &vlanResponse{}
	err := json.Unmarshal(s, vResponse)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	if len(vResponse.Result.Body.VlanBriefTable.VlanBriefRow) < 1 {
		return nil, fmt.Errorf("Error parsing the received response: %s", s)
	}
	vlanRef := make(map[string]int)
	for i, v := range vResponse.Result.Body.VlanBriefTable.VlanBriefRow {
		vlan := &Vlan{}
		vlan.ID = v.ID
		vlan.Name = v.Name
		vlan.State = v.State
		vlan.ShutdownState = v.ShutdownState
		if v.Ports.Entries != nil {
			vlan.Ports = v.Ports.Entries
		} else {
			vlan.Ports = []string{}
		}
		vlanRef[vlan.ID] = i
		vlans = append(vlans, vlan)
	}
	for _, v := range vResponse.Result.Body.MtuInfoTable.MtuInfoRow {
		if i, ok := vlanRef[v.ID]; ok {
			vlans[i].MediaType = v.MediaType
			vlans[i].Mode = v.Mode
		}
	}
	return vlans, nil
}
