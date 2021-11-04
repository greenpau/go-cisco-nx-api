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
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	notAvailable = "-"
	falseIndicator = "F"
)

type macAddressResponse struct {
	ID      uint64                   `json:"id" xml:"id"`
	Version string                   `json:"jsonrpc" xml:"jsonrpc"`
	Result  macAddressResponseResult `json:"result" xml:"result"`
}

type macAddressResponseResult struct {
	Body macAddressResponseResultBody `json:"body" xml:"body"`
}

type macAddressResponseResultBody struct {
	MacAddressTable macAddressResponseResultBodyMacAddressTable `json:"TABLE_mac_address" xml:"TABLE_mac_address"`
}

type macAddressResponseResultBodyMacAddressTable struct {
	MacAddressRow []macAddressResponseResultBodyMacAddressRow `json:"ROW_mac_address" xml:"ROW_mac_address"`
}

type macAddressResponseResultBodyMacAddressRow struct {
	MacAddress string `json:"disp_mac_addr" xml:"disp_mac_addr"`
	Type       string `json:"disp_type" xml:"disp_type"`
	Vlan       string `json:"disp_vlan" xml:"disp_vlan"`
	Age        string `json:"disp_age" xml:"disp_age"`
	Secure     string `json:"disp_secure" xml:"disp_secure"`
	Notify     string `json:"disp_notify" xml:"disp_notify"`
	Port       string `json:"disp_port" xml:"disp_port"`
}

// MacAddressTable contains MAC address table information.
// The information in the structure is from the output of "show mac address-table" command.
type MacAddressTable struct {
	Item []MacAddressItem
}

type MacAddressItem struct {
	Address string `json:"address" xml:"address"`
	Type    string `json:"type" xml:"type"`
	VlanID  int    `json:"vlan" xml:"vlan"`
	Age     int    `json:"age" xml:"age"`
	Port    string `json:"port" xml:"port"`
	Secured bool   `json:"secured" xml:"secured"`
	Notify  bool   `json:"notify" xml:"notify"`
}

// NewMacAddressTableFromBytes returns an MacAddressTable instance from an input byte array.
func NewMacAddressTableFromBytes(s []byte) (*MacAddressTable, error) {
	var table *MacAddressTable
	resp := &JSONRPCResponse{}
	err := json.Unmarshal(s, resp)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("command returned failure: %v", resp.Error)
	}
	var body JSONRPCResponseBody
	err = json.Unmarshal(resp.Result, &body)
	if err != nil {
		return nil, fmt.Errorf("parsing MAC address table result body error: %v", err)
	}
	var macAddressTableResult macAddressResponseResultBody
	err = json.Unmarshal(body.Body, &macAddressTableResult)
	if err != nil {
		return nil, fmt.Errorf("parsing MAC address table result error: %v", err)
	}

	table = new(MacAddressTable)
	for _, row := range macAddressTableResult.MacAddressTable.MacAddressRow {
		var item MacAddressItem
		item.Address = row.MacAddress
		item.Type = row.Type
		if row.Age == notAvailable {
			item.Age = -1
		} else {
			// error is ignored, assume switch will not return invalid value.
			// even if value is invalid, age is set to zero.
			item.Age, _ = strconv.Atoi(row.Age)
		}
		item.Port = row.Port
		if row.Notify == falseIndicator {
			item.Notify = false
		} else {
			item.Notify = true
		}

		if row.Secure == falseIndicator {
			item.Secured = false
		} else {
			item.Secured = true
		}
		if row.Vlan == notAvailable {
			item.VlanID = -1
		} else {
			item.VlanID, _ = strconv.Atoi(row.Vlan)
		}
		table.Item = append(table.Item, item)
	}
	return table, nil
}
