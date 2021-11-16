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

type cdpNeighborResponse struct {
	ID      uint64                    `json:"id" xml:"id"`
	Version string                    `json:"jsonrpc" xml:"jsonrpc"`
	Result  cdpNeighborResponseResult `json:"result" xml:"result"`
}

type cdpNeighborResponseResult struct {
	Body cdpNeighborResponseResultBody `json:"body" xml:"body"`
}

type cdpNeighborResponseResultBody struct {
	NeighborCount    int                                           `json:"neigh_count" xml:"neigh_count"`
	CDPNeighborTable cdpNeighborResponseResultBodycdpNeighborTable `json:"TABLE_cdp_neighbor_brief_info" xml:"TABLE_cdp_neighbor_brief_info"`
}

type cdpNeighborResponseResultBodycdpNeighborTable struct {
	CdpNeighborRow []cdpNeighborResponseResultBodyCdpNeighborRow `json:"ROW_cdp_neighbor_brief_info" xml:"ROW_cdp_neighbor_brief_info"`
}

type cdpNeighborResponseResultBodyCdpNeighborRow struct {
	IfIndex    int64       `json:"ifindex" xml:"ifindex"`
	DeviceID   string      `json:"device_id" xml:"device_id"`
	IntfID     string      `json:"intf_id" xml:"intf_id"`
	TTL        interface{} `json:"ttl" xml:"ttl"`
	Capability interface{} `json:"capability" xml:"capability"`
	PlatformID string      `json:"platform_id" xml:"platform_id"`
	PortID     string      `json:"port_id" xml:"port_id"`
}

// CDPNeighborTable contains CDP neighbor table information.
// The information in the structure is from the output of "show cdp neighbor" command.
type CDPNeighborTable struct {
	Item []CDPNeighborItem
}

type CDPNeighborItem struct {
	IfIndex    int64    `json:"ifindex" xml:"ifindex"`
	DeviceID   string   `json:"device_id" xml:"device_id"`
	IntfID     string   `json:"intf_id" xml:"intf_id"`
	TTL        int      `json:"ttl" xml:"ttl"`
	Capability []string `json:"capability" xml:"capability"`
	PlatformID string   `json:"platform_id" xml:"platform_id"`
	PortID     string   `json:"port_id" xml:"port_id"`
}

// NewCDPNeighborTableFromBytes returns an CDPNeighborTable instance from an input byte array.
func NewCDPNeighborTableFromBytes(s []byte) (*CDPNeighborTable, error) {
	var table *CDPNeighborTable
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
		return nil, fmt.Errorf("parsing CDP neighbor table result body error: %v", err)
	}
	var cdpNeighborTableResult cdpNeighborResponseResultBody
	err = json.Unmarshal(body.Body, &cdpNeighborTableResult)
	if err != nil {
		return nil, fmt.Errorf("parsing CDP neighbor table result error: %v", err)
	}

	table = new(CDPNeighborTable)
	for _, row := range cdpNeighborTableResult.CDPNeighborTable.CdpNeighborRow {
		var item CDPNeighborItem
		item.IfIndex = row.IfIndex
		item.DeviceID = row.DeviceID
		item.Capability = make([]string, 0)
		if row.Capability != nil {
			switch row.Capability.(type) {
			case string:
				item.Capability = append(item.Capability, row.Capability.(string))
			case []interface{}:
				vals := row.Capability.([]interface{})
				for _, val := range vals {
					item.Capability = append(item.Capability, val.(string))
				}
			}
		}
		item.IntfID = row.IntfID
		item.PortID = row.PortID
		switch row.TTL.(type) {
		case string:
			item.TTL, err = strconv.Atoi(row.TTL.(string))
			if err != nil {
				// skip error and continue.
				item.TTL = 0
			}
		case int:
			item.TTL = row.TTL.(int)
		}

		item.PlatformID = row.PlatformID
		table.Item = append(table.Item, item)
	}
	return table, nil
}
