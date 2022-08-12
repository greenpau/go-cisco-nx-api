package client

import (
	"encoding/json"
	"fmt"
)

type vlanCountersResponse struct {
	ID      uint64                     `json:"id" xml:"id"`
	Version string                     `json:"jsonrpc" xml:"jsonrpc"`
	Result  vlanCountersResponseResult `json:"result" xml:"result"`
}

type vlanCountersResponseResult struct {
	Body vlanCountersResponseResultBody `json:"body" xml:"body"`
}

type vlanCountersResponseResultBody struct {
	VlanCountersTable vlanCountersResponseResultBodyVlanCountersTable `json:"TABLE_vlancounters" xml:"TABLE_vlancounters"`
}

type vlanCountersResponseResultBodyVlanCountersTable struct {
	VlanCountersRow []vlanCountersResponseResultBodyVlanCountersRow `json:"ROW_vlancounters" xml:"ROW_vlancounters"`
}

type vlanCountersResponseResultBodyVlanCountersRow struct {
	ID                    uint64 `json:"vlanshowbr-vlanid" xml:"vlanshowbr-vlanid"`
	InputUnicastBytes     uint64 `json:"l2_ing_ucast_b" xml:"l2_ing_ucast_b"`
	InputUnicastPackets   uint64 `json:"l2_ing_ucast_p" xml:"l2_ing_ucast_p"`
	InputMulticastBytes   uint64 `json:"l2_ing_mcast_b" xml:"l2_ing_mcast_b"`
	InputMulticastPackets uint64 `json:"l2_ing_mcast_p" xml:"l2_ing_mcast_p"`
	InputBroadcastBytes   uint64 `json:"l2_ing_bcast_b" xml:"l2_ing_bcast_b"`
	InputBroadcastPackets uint64 `json:"l2_ing_bcast_p" xml:"l2_ing_bcast_p"`
	OutputUnicastBytes    uint64 `json:"l2_egr_ucast_b" xml:"l2_egr_ucast_b"`
	OutputUnicastPackets  uint64 `json:"l2_egr_ucast_p" xml:"l2_egr_ucast_p"`
	L3InputUnicastBytes   uint64 `json:"l3_ucast_rcv_b" xml:"l3_ucast_rcv_b"`
	L3InputUnicastPackets uint64 `json:"l3_ucast_rcv_p" xml:"l3_ucast_rcv_p"`
}

type VlanCounters struct {
	ID                    uint64 `json:"id" xml:"id"`
	InputUnicastBytes     uint64 `json:"input_ucast_bytes" xml:"input_ucast_bytes"`
	InputUnicastBytesL3   uint64 `json:"input_ucast_bytes_l3" xml:"input_ucast_bytes_l3"`
	InputMulticastBytes   uint64 `json:"input_mcast_bytes" xml:"input_mcast_bytes"`
	InputBroadcastBytes   uint64 `json:"input_bcast_bytes" xml:"input_bcast_bytes"`
	InputUnicastPackets   uint64 `json:"input_ucast_packets" xml:"input_ucast_packets"`
	InputUnicastPacketsL3 uint64 `json:"input_ucast_packets_l3" xml:"input_ucast_packets_l3"`
	InputMulticastPackets uint64 `json:"input_mcast_packets" xml:"input_mcast_packets"`
	InputBroadcastPackets uint64 `json:"input_bcast_packets" xml:"input_bcast_packets"`
	OutputUnicastBytes    uint64 `json:"output_ucast_bytes" xml:"output_ucast_bytes"`
	OutputUnicastPackets  uint64 `json:"output_ucast_packets" xml:"output_ucast_packets"`
}

// NewVlanCountersFromString returns VlanCounter instance from an input string.
func NewVlanCountersFromString(s string) ([]*VlanCounters, error) {
	return NewVlanCountersFromBytes([]byte(s))
}

// NewVlanCountersFromBytes returns VlanCounter instance from an input byte array.
func NewVlanCountersFromBytes(s []byte) ([]*VlanCounters, error) {
	var vlanCounters []*VlanCounters
	vCountersResponse := &vlanCountersResponse{}
	err := json.Unmarshal(s, vCountersResponse)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	if len(vCountersResponse.Result.Body.VlanCountersTable.VlanCountersRow) < 1 {
		return nil, fmt.Errorf("Error parsing the received response: %s", s)
	}
	for _, v := range vCountersResponse.Result.Body.VlanCountersTable.VlanCountersRow {
		vlanCounter := &VlanCounters{}
		vlanCounter.ID = v.ID
		vlanCounter.InputUnicastBytes = v.InputUnicastBytes
		vlanCounter.InputUnicastBytesL3 = v.L3InputUnicastBytes
		vlanCounter.InputMulticastBytes = v.InputMulticastBytes
		vlanCounter.InputBroadcastBytes = v.InputBroadcastBytes
		vlanCounter.InputUnicastPackets = v.InputUnicastPackets
		vlanCounter.InputUnicastPacketsL3 = v.L3InputUnicastPackets
		vlanCounter.InputMulticastPackets = v.InputMulticastPackets
		vlanCounter.InputBroadcastPackets = v.InputBroadcastPackets
		vlanCounter.OutputUnicastBytes = v.OutputUnicastBytes
		vlanCounter.OutputUnicastPackets = v.OutputUnicastPackets
		vlanCounters = append(vlanCounters, vlanCounter)
	}
	return vlanCounters, nil
}
