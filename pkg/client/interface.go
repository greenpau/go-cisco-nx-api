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
	//"github.com/davecgh/go-spew/spew"
	//"strings"
	"strconv"
	"strings"
)

type interfacesResponseResultBody struct {
	InterfaceTable interfacesResponseResultBodyInterfaceTable `json:"TABLE_interface" xml:"TABLE_interface"`
}

type interfacesResponseResultBodyInterfaceTable struct {
	InterfaceRow []interfaceResponseResultBodyInterfaceRow `json:"ROW_interface" xml:"ROW_interface"`
}

type interfaceOneResponseResultBody struct {
	InterfaceTable interfaceOneResponseResultBodyInterfaceTable `json:"TABLE_interface" xml:"TABLE_interface"`
}

type interfaceOneResponseResultBodyInterfaceTable struct {
	InterfaceRow interfaceResponseResultBodyInterfaceRow `json:"ROW_interface" xml:"ROW_interface"`
}

type interfaceResponseResultBodyInterfaceRow struct {
	AdminState    string `json:"admin_state" xml:"admin_state"`
	Desc          string `json:"desc" xml:"desc"`
	Encapsulation string `json:"encapsulation" xml:"encapsulation"`
	EthAutoneg    string `json:"eth_autoneg" xml:"eth_autoneg"`
	// WARN: eth_babbles is a string with numeric value
	EthBabbles string `json:"eth_babbles" xml:"eth_babbles"`
	// WARN: eth_bad_eth is a string with numeric value
	EthBadEth string `json:"eth_bad_eth" xml:"eth_bad_eth"`
	// WARN: eth_bad_proto is a string with numeric value
	EthBadProto      string `json:"eth_bad_proto" xml:"eth_bad_proto"`
	EthBeacon        string `json:"eth_beacon" xml:"eth_beacon"`
	EthBiaAddr       string `json:"eth_bia_addr" xml:"eth_bia_addr"`
	EthBundle        string `json:"eth_bundle" xml:"eth_bundle"`
	EthBw            uint64 `json:"eth_bw" xml:"eth_bw"`
	EthClearCounters string `json:"eth_clear_counters" xml:"eth_clear_counters"`
	// WARN: eth_coll is a string with numeric value
	EthColl string `json:"eth_coll" xml:"eth_coll"`
	// WARN: eth_crc is a string with numeric value
	EthCrc string `json:"eth_crc" xml:"eth_crc"`
	// WARN: eth_deferred is a string with numeric value
	EthDeferred string `json:"eth_deferred" xml:"eth_deferred"`
	EthDly      uint64 `json:"eth_dly" xml:"eth_dly"`
	// WARN: eth_dribble is a string with numeric value
	EthDribble   string `json:"eth_dribble" xml:"eth_dribble"`
	EthDuplex    string `json:"eth_duplex" xml:"eth_duplex"`
	EthEeeState  string `json:"eth_eee_state" xml:"eth_eee_state"`
	EthEncapVlan uint64 `json:"eth_encap_vlan" xml:"eth_encap_vlan"`
	EthEthertype string `json:"eth_ethertype" xml:"eth_ethertype"`
	// WARN: eth_frame is a string with numeric value
	EthFrame  string `json:"eth_frame" xml:"eth_frame"`
	EthGiants uint64 `json:"eth_giants" xml:"eth_giants"`
	EthHwAddr string `json:"eth_hw_addr" xml:"eth_hw_addr"`
	EthHwDesc string `json:"eth_hw_desc" xml:"eth_hw_desc"`
	// WARN: eth_ignored is a string with numeric value
	EthIgnored    string `json:"eth_ignored" xml:"eth_ignored"`
	EthInFlowctrl string `json:"eth_in_flowctrl" xml:"eth_in_flowctrl"`
	// WARN: eth_in_ifdown_drops is a string with numeric value
	EthInIfdownDrops string `json:"eth_in_ifdown_drops" xml:"eth_in_ifdown_drops"`
	EthInbcast       uint64 `json:"eth_inbcast" xml:"eth_inbcast"`
	EthInbytes       uint64 `json:"eth_inbytes" xml:"eth_inbytes"`
	// WARN: eth_indiscard is a string with numeric value
	EthIndiscard string `json:"eth_indiscard" xml:"eth_indiscard"`
	// WARN: eth_inerr is a string with numeric value
	EthInerr   string `json:"eth_inerr" xml:"eth_inerr"`
	EthInmcast uint64 `json:"eth_inmcast" xml:"eth_inmcast"`
	// WARN: eth_inpause is a string with numeric value
	EthInpause string `json:"eth_inpause" xml:"eth_inpause"`
	EthInpkts  uint64 `json:"eth_inpkts" xml:"eth_inpkts"`
	// WARN: eth_inrate1_bits is a string with numeric value
	EthInrate1Bits string `json:"eth_inrate1_bits" xml:"eth_inrate1_bits"`
	// WARN: eth_inrate1_pkts is a string with numeric value
	EthInrate1Pkts        string `json:"eth_inrate1_pkts" xml:"eth_inrate1_pkts"`
	EthInrate1SummaryBits string `json:"eth_inrate1_summary_bits" xml:"eth_inrate1_summary_bits"`
	EthInrate1SummaryPkts string `json:"eth_inrate1_summary_pkts" xml:"eth_inrate1_summary_pkts"`
	// WARN: eth_inrate2_bits is a string with numeric value
	EthInrate2Bits string `json:"eth_inrate2_bits" xml:"eth_inrate2_bits"`
	// WARN: eth_inrate2_pkts is a string with numeric value
	EthInrate2Pkts        string `json:"eth_inrate2_pkts" xml:"eth_inrate2_pkts"`
	EthInrate2SummaryBits string `json:"eth_inrate2_summary_bits" xml:"eth_inrate2_summary_bits"`
	EthInrate2SummaryPkts string `json:"eth_inrate2_summary_pkts" xml:"eth_inrate2_summary_pkts"`
	// WARN: eth_inrate3_bits is a string with numeric value
	EthInrate3Bits string `json:"eth_inrate3_bits" xml:"eth_inrate3_bits"`
	// WARN: eth_inrate3_pkts is a string with numeric value
	EthInrate3Pkts        string `json:"eth_inrate3_pkts" xml:"eth_inrate3_pkts"`
	EthInrate3SummaryBits string `json:"eth_inrate3_summary_bits" xml:"eth_inrate3_summary_bits"`
	EthInrate3SummaryPkts string `json:"eth_inrate3_summary_pkts" xml:"eth_inrate3_summary_pkts"`
	EthInucast            uint64 `json:"eth_inucast" xml:"eth_inucast"`
	EthIPAddr             string `json:"eth_ip_addr" xml:"eth_ip_addr"`
	EthIPMask             uint64 `json:"eth_ip_mask" xml:"eth_ip_mask"`
	EthIPPrefix           string `json:"eth_ip_prefix" xml:"eth_ip_prefix"`
	// WARN: eth_jumbo_inpkts is a string with numeric value
	EthJumboInpkts string `json:"eth_jumbo_inpkts" xml:"eth_jumbo_inpkts"`
	// WARN: eth_jumbo_outpkts is a string with numeric value
	EthJumboOutpkts string `json:"eth_jumbo_outpkts" xml:"eth_jumbo_outpkts"`
	// WARN: eth_latecoll is a string with numeric value
	EthLatecoll        string `json:"eth_latecoll" xml:"eth_latecoll"`
	EthLinkFlapped     string `json:"eth_link_flapped" xml:"eth_link_flapped"`
	EthLoadInterval1Rx uint64 `json:"eth_load_interval1_rx" xml:"eth_load_interval1_rx"`
	EthLoadInterval1Tx string `json:"eth_load_interval1_tx" xml:"eth_load_interval1_tx"`
	EthLoadInterval2Rx string `json:"eth_load_interval2_rx" xml:"eth_load_interval2_rx"`
	EthLoadInterval2Tx string `json:"eth_load_interval2_tx" xml:"eth_load_interval2_tx"`
	EthLoadInterval3Rx string `json:"eth_load_interval3_rx" xml:"eth_load_interval3_rx"`
	EthLoadInterval3Tx string `json:"eth_load_interval3_tx" xml:"eth_load_interval3_tx"`
	// WARN: eth_lostcarrier is a string with numeric value
	EthLostcarrier string `json:"eth_lostcarrier" xml:"eth_lostcarrier"`
	EthMdix        string `json:"eth_mdix" xml:"eth_mdix"`
	EthMedia       string `json:"eth_media" xml:"eth_media"`
	EthMembers     string `json:"eth_members" xml:"eth_members"`
	EthMode        string `json:"eth_mode" xml:"eth_mode"`
	EthMtu         string `json:"eth_mtu" xml:"eth_mtu"`
	EthNobuf       uint64 `json:"eth_nobuf" xml:"eth_nobuf"`
	// WARN: eth_nocarrier is a string with numeric value
	EthNocarrier   string `json:"eth_nocarrier" xml:"eth_nocarrier"`
	EthOutFlowctrl string `json:"eth_out_flowctrl" xml:"eth_out_flowctrl"`
	EthOutbcast    uint64 `json:"eth_outbcast" xml:"eth_outbcast"`
	EthOutbytes    uint64 `json:"eth_outbytes" xml:"eth_outbytes"`
	// WARN: eth_outdiscard is a string with numeric value
	EthOutdiscard string `json:"eth_outdiscard" xml:"eth_outdiscard"`
	// WARN: eth_outerr is a string with numeric value
	EthOuterr   string `json:"eth_outerr" xml:"eth_outerr"`
	EthOutmcast uint64 `json:"eth_outmcast" xml:"eth_outmcast"`
	// WARN: eth_outpause is a string with numeric value
	EthOutpause string `json:"eth_outpause" xml:"eth_outpause"`
	EthOutpkts  uint64 `json:"eth_outpkts" xml:"eth_outpkts"`
	// WARN: eth_outrate1_bits is a string with numeric value
	EthOutrate1Bits string `json:"eth_outrate1_bits" xml:"eth_outrate1_bits"`
	// WARN: eth_outrate1_pkts is a string with numeric value
	EthOutrate1Pkts        string `json:"eth_outrate1_pkts" xml:"eth_outrate1_pkts"`
	EthOutrate1SummaryBits string `json:"eth_outrate1_summary_bits" xml:"eth_outrate1_summary_bits"`
	EthOutrate1SummaryPkts string `json:"eth_outrate1_summary_pkts" xml:"eth_outrate1_summary_pkts"`
	// WARN: eth_outrate2_bits is a string with numeric value
	EthOutrate2Bits string `json:"eth_outrate2_bits" xml:"eth_outrate2_bits"`
	// WARN: eth_outrate2_pkts is a string with numeric value
	EthOutrate2Pkts        string `json:"eth_outrate2_pkts" xml:"eth_outrate2_pkts"`
	EthOutrate2SummaryBits string `json:"eth_outrate2_summary_bits" xml:"eth_outrate2_summary_bits"`
	EthOutrate2SummaryPkts string `json:"eth_outrate2_summary_pkts" xml:"eth_outrate2_summary_pkts"`
	// WARN: eth_outrate3_bits is a string with numeric value
	EthOutrate3Bits string `json:"eth_outrate3_bits" xml:"eth_outrate3_bits"`
	// WARN: eth_outrate3_pkts is a string with numeric value
	EthOutrate3Pkts        string `json:"eth_outrate3_pkts" xml:"eth_outrate3_pkts"`
	EthOutrate3SummaryBits string `json:"eth_outrate3_summary_bits" xml:"eth_outrate3_summary_bits"`
	EthOutrate3SummaryPkts string `json:"eth_outrate3_summary_pkts" xml:"eth_outrate3_summary_pkts"`
	EthOutucast            uint64 `json:"eth_outucast" xml:"eth_outucast"`
	// WARN: eth_overrun is a string with numeric value
	EthOverrun     string `json:"eth_overrun" xml:"eth_overrun"`
	EthRatemode    string `json:"eth_ratemode" xml:"eth_ratemode"`
	EthReliability string `json:"eth_reliability" xml:"eth_reliability"`
	EthResetCntr   uint64 `json:"eth_reset_cntr" xml:"eth_reset_cntr"`
	EthRunts       uint64 `json:"eth_runts" xml:"eth_runts"`
	EthRxload      string `json:"eth_rxload" xml:"eth_rxload"`
	EthSpeed       string `json:"eth_speed" xml:"eth_speed"`
	// WARN: eth_storm_supp is a string with numeric value
	EthStormSupp  string `json:"eth_storm_supp" xml:"eth_storm_supp"`
	EthSwtMonitor string `json:"eth_swt_monitor" xml:"eth_swt_monitor"`
	EthTxload     string `json:"eth_txload" xml:"eth_txload"`
	// WARN: eth_underrun is a string with numeric value
	EthUnderrun string `json:"eth_underrun" xml:"eth_underrun"`
	// WARN: eth_watchdog is a string with numeric value
	EthWatchdog string `json:"eth_watchdog" xml:"eth_watchdog"`
	Interface   string `json:"interface" xml:"interface"`
	// WARN: loop_in_bytes is a string with numeric value
	LoopInBytes string `json:"loop_in_bytes" xml:"loop_in_bytes"`
	// WARN: loop_in_compressed is a string with numeric value
	LoopInCompressed string `json:"loop_in_compressed" xml:"loop_in_compressed"`
	// WARN: loop_in_errors is a string with numeric value
	LoopInErrors string `json:"loop_in_errors" xml:"loop_in_errors"`
	// WARN: loop_in_fifo is a string with numeric value
	LoopInFifo string `json:"loop_in_fifo" xml:"loop_in_fifo"`
	// WARN: loop_in_frame is a string with numeric value
	LoopInFrame string `json:"loop_in_frame" xml:"loop_in_frame"`
	// WARN: loop_in_mcast is a string with numeric value
	LoopInMcast string `json:"loop_in_mcast" xml:"loop_in_mcast"`
	// WARN: loop_in_overrun is a string with numeric value
	LoopInOverrun string `json:"loop_in_overrun" xml:"loop_in_overrun"`
	LoopInPkts    uint64 `json:"loop_in_pkts" xml:"loop_in_pkts"`
	LoopOutBytes  string `json:"loop_out_bytes" xml:"loop_out_bytes"`
	// WARN: loop_out_carriers is a string with numeric value
	LoopOutCarriers string `json:"loop_out_carriers" xml:"loop_out_carriers"`
	// WARN: loop_out_collisions is a string with numeric value
	LoopOutCollisions string `json:"loop_out_collisions" xml:"loop_out_collisions"`
	// WARN: loop_out_errors is a string with numeric value
	LoopOutErrors string `json:"loop_out_errors" xml:"loop_out_errors"`
	// WARN: loop_out_fifo is a string with numeric value
	LoopOutFifo string `json:"loop_out_fifo" xml:"loop_out_fifo"`
	LoopOutPkts string `json:"loop_out_pkts" xml:"loop_out_pkts"`
	// WARN: loop_out_underruns is a string with numeric value
	LoopOutUnderruns string `json:"loop_out_underruns" xml:"loop_out_underruns"`
	Medium           string `json:"medium" xml:"medium"`
	ParentInterface  string `json:"parent_interface" xml:"parent_interface"`
	ShareState       string `json:"share_state" xml:"share_state"`
	// Interface state
	// - up       value: 0x0001
	// - down     value: 0x0002
	// - testing  value: 0x0004
	// - trunking value: 0x0008
	// - link-up  value: 0x0010
	State              string `json:"state" xml:"state"`
	StateRsnDesc       string `json:"state_rsn_desc" xml:"state_rsn_desc"`
	SviAdminState      string `json:"svi_admin_state" xml:"svi_admin_state"`
	SviArpType         string `json:"svi_arp_type" xml:"svi_arp_type"`
	SviBw              string `json:"svi_bw" xml:"svi_bw"`
	SviDelay           string `json:"svi_delay" xml:"svi_delay"`
	SviIPAddr          string `json:"svi_ip_addr" xml:"svi_ip_addr"`
	SviIPMask          string `json:"svi_ip_mask" xml:"svi_ip_mask"`
	SviLineProto       string `json:"svi_line_proto" xml:"svi_line_proto"`
	SviMac             string `json:"svi_mac" xml:"svi_mac"`
	SviMtu             string `json:"svi_mtu" xml:"svi_mtu"`
	SviRsnDesc         string `json:"svi_rsn_desc" xml:"svi_rsn_desc"`
	SviRxLoad          string `json:"svi_rx_load" xml:"svi_rx_load"`
	SviTimeLastCleared string `json:"svi_time_last_cleared" xml:"svi_time_last_cleared"`
	SviTxLoad          uint64 `json:"svi_tx_load" xml:"svi_tx_load"`
	SviReliability     string `json:"svi_reliability" xml:"svi_reliability"`
	// WARN: svi_ucast_bytes_in is a string with numeric value
	SviUcastBytesIn string `json:"svi_ucast_bytes_in" xml:"svi_ucast_bytes_in"`
	// WARN: svi_ucast_pkts_in is a string with numeric value
	SviUcastPktsIn  string `json:"svi_ucast_pkts_in" xml:"svi_ucast_pkts_in"`
	VdcLvlInAvgBits uint64 `json:"vdc_lvl_in_avg_bits" xml:"vdc_lvl_in_avg_bits"`
	VdcLvlInAvgPkts string `json:"vdc_lvl_in_avg_pkts" xml:"vdc_lvl_in_avg_pkts"`
	// WARN: vdc_lvl_in_bcast is a string with numeric value
	VdcLvlInBcast string `json:"vdc_lvl_in_bcast" xml:"vdc_lvl_in_bcast"`
	VdcLvlInBytes string `json:"vdc_lvl_in_bytes" xml:"vdc_lvl_in_bytes"`
	// WARN: vdc_lvl_in_mcast is a string with numeric value
	VdcLvlInMcast    string `json:"vdc_lvl_in_mcast" xml:"vdc_lvl_in_mcast"`
	VdcLvlInPkts     uint64 `json:"vdc_lvl_in_pkts" xml:"vdc_lvl_in_pkts"`
	VdcLvlInUcast    string `json:"vdc_lvl_in_ucast" xml:"vdc_lvl_in_ucast"`
	VdcLvlOutAvgBits string `json:"vdc_lvl_out_avg_bits" xml:"vdc_lvl_out_avg_bits"`
	VdcLvlOutAvgPkts string `json:"vdc_lvl_out_avg_pkts" xml:"vdc_lvl_out_avg_pkts"`
	VdcLvlOutBcast   string `json:"vdc_lvl_out_bcast" xml:"vdc_lvl_out_bcast"`
	VdcLvlOutBytes   string `json:"vdc_lvl_out_bytes" xml:"vdc_lvl_out_bytes"`
	VdcLvlOutMcast   string `json:"vdc_lvl_out_mcast" xml:"vdc_lvl_out_mcast"`
	VdcLvlOutPkts    string `json:"vdc_lvl_out_pkts" xml:"vdc_lvl_out_pkts"`
	VdcLvlOutUcast   string `json:"vdc_lvl_out_ucast" xml:"vdc_lvl_out_ucast"`
}

// Interface contains system information. The information in the structure
// is from the output of "show interface" command.
type Interface struct {
	Name        string `json:"name" xml:"name"`
	LocalIndex  int    `json:"local_index" xml:"local_index"`
	Description string `json:"description" xml:"description"`
	Counters    struct {
		Babbles           uint64 `json:"babbles" xml:"babbles"`
		BadEtherTypeDrops uint64 `json:"bad_ethtype_drops" xml:"bad_ethtype_drops"`
		BadProtocolDrops  uint64 `json:"bad_proto_drops" xml:"bad_proto_drops"`
		NoCarrier         uint64 `json:"no_carrier" xml:"no_carrier"`
		// Dribble are input packets with dribble condition
		Dribble          uint64 `json:"dribble" xml:"dribble"`
		InputFrameErrors uint64 `json:"input_frame_errors" xml:"input_frame_errors"`
		InputDiscards    uint64 `json:"input_discards" xml:"input_discards"`
		InputErrors      uint64 `json:"input_errors" xml:"input_errors"`
		InputPause       uint64 `json:"input_pause" xml:"input_pause"`
		InputOverruns    uint64 `json:"input_overruns" xml:"input_overruns"`
		// InputIfaceDownDrops are Input if-down drops
		InputIfaceDownDrops    uint64 `json:"input_iface_down_drops" xml:"input_iface_down_drops"`
		InputBytes             uint64 `json:"input_bytes" xml:"input_bytes"`
		InputUnicastBytes      uint64 `json:"input_ucast_bytes" xml:"input_ucast_bytes"`
		InputPackets           uint64 `json:"input_packets" xml:"input_packets"`
		InputUnicastPackets    uint64 `json:"input_ucast_packets" xml:"input_ucast_packets"`
		InputBroadcastPackets  uint64 `json:"input_bcast_packets" xml:"input_bcast_packets"`
		InputMulticastPackets  uint64 `json:"input_mcast_packets" xml:"input_mcast_packets"`
		InputJumboPackets      uint64 `json:"input_jumbo_packets" xml:"input_jumbo_packets"`
		InputCompressed        uint64 `json:"input_compressed" xml:"input_compressed"`
		InputFlowControl       string `json:"input_flow_control" xml:"input_flow_control"`
		InputFifo              uint64 `json:"input_fifo" xml:"input_fifo"`
		LateCollisions         uint64 `json:"late_collisions" xml:"late_collisions"`
		LostCarrier            uint64 `json:"lost_carrier" xml:"lost_carrier"`
		OutputDiscards         uint64 `json:"output_discards" xml:"output_discards"`
		OutputErrors           uint64 `json:"output_errors" xml:"output_errors"`
		OutputPause            uint64 `json:"output_pause" xml:"output_pause"`
		OutputUnderruns        uint64 `json:"output_underruns" xml:"output_underruns"`
		OutputBytes            uint64 `json:"output_bytes" xml:"output_bytes"`
		OutputUnicastBytes     uint64 `json:"output_ucast_bytes" xml:"output_ucast_bytes"`
		OutputPackets          uint64 `json:"output_packets" xml:"output_packets"`
		OutputUnicastPackets   uint64 `json:"output_ucast_packets" xml:"output_ucast_packets"`
		OutputBroadcastPackets uint64 `json:"output_bcast_packets" xml:"output_bcast_packets"`
		OutputMulticastPackets uint64 `json:"output_mcast_packets" xml:"output_mcast_packets"`
		OutputJumboPackets     uint64 `json:"output_jumbo_packets" xml:"output_jumbo_packets"`
		OutputCarrierErrors    uint64 `json:"output_carrier_errors" xml:"output_carrier_errors"`
		OutputFlowControl      string `json:"output_flow_control" xml:"output_flow_control"`
		Collisions             uint64 `json:"collisions" xml:"collisions"`
		OutputFifo             uint64 `json:"output_fifo" xml:"output_fifo"`
		Watchdog               uint64 `json:"watchdog" xml:"watchdog"`
		StormSuppression       uint64 `json:"storm_suppression" xml:"storm_suppression"`
		Ignored                uint64 `json:"ignored" xml:"ignored"`
		Runts                  uint64 `json:"runts" xml:"runts"`
		CrcErrors              uint64 `json:"crc_errors" xml:"crc_errors"`
		Deferred               uint64 `json:"deferred" xml:"deferred"`
		NoBufferReceivedErrors uint64 `json:"no_buffer" xml:"no_buffer"`
		Resets                 uint64 `json:"resets" xml:"resets"`
		Intervals              struct {
			Interval1 struct {
				InputRateBits            uint64 `json:"input_rate_bits" xml:"input_rate_bits"`
				InputRatePackets         uint64 `json:"input_rate_packets" xml:"input_rate_packets"`
				InputRateSummaryBits     string `json:"input_rate_summary_bits" xml:"input_rate_summary_bits"`
				InputRateSummaryPackets  string `json:"input_rate_summary_packets" xml:"input_rate_summary_packets"`
				OutputRateBits           uint64 `json:"output_rate_bits" xml:"output_rate_bits"`
				OutputRatePackets        uint64 `json:"output_rate_packets" xml:"output_rate_packets"`
				OutputRateSummaryBits    string `json:"output_rate_summary_bits" xml:"output_rate_summary_bits"`
				OutputRateSummaryPackets string `json:"output_rate_summary_packets" xml:"output_rate_summary_packets"`
				RxLoad                   uint64 `json:"rx_load" xml:"rx_load"`
				TxLoad                   uint64 `json:"tx_load" xml:"tx_load"`
			}
			Interval2 struct {
				InputRateBits            uint64 `json:"input_rate_bits" xml:"input_rate_bits"`
				InputRatePackets         uint64 `json:"input_rate_packets" xml:"input_rate_packets"`
				InputRateSummaryBits     string `json:"input_rate_summary_bits" xml:"input_rate_summary_bits"`
				InputRateSummaryPackets  string `json:"input_rate_summary_packets" xml:"input_rate_summary_packets"`
				OutputRateBits           uint64 `json:"output_rate_bits" xml:"output_rate_bits"`
				OutputRatePackets        uint64 `json:"output_rate_packets" xml:"output_rate_packets"`
				OutputRateSummaryBits    string `json:"output_rate_summary_bits" xml:"output_rate_summary_bits"`
				OutputRateSummaryPackets string `json:"output_rate_summary_packets" xml:"output_rate_summary_packets"`
				RxLoad                   uint64 `json:"rx_load" xml:"rx_load"`
				TxLoad                   uint64 `json:"tx_load" xml:"tx_load"`
			}
			Interval3 struct {
				InputRateBits            uint64 `json:"input_rate_bits" xml:"input_rate_bits"`
				InputRatePackets         uint64 `json:"input_rate_packets" xml:"input_rate_packets"`
				InputRateSummaryBits     string `json:"input_rate_summary_bits" xml:"input_rate_summary_bits"`
				InputRateSummaryPackets  string `json:"input_rate_summary_packets" xml:"input_rate_summary_packets"`
				OutputRateBits           uint64 `json:"output_rate_bits" xml:"output_rate_bits"`
				OutputRatePackets        uint64 `json:"output_rate_packets" xml:"output_rate_packets"`
				OutputRateSummaryBits    string `json:"output_rate_summary_bits" xml:"output_rate_summary_bits"`
				OutputRateSummaryPackets string `json:"output_rate_summary_packets" xml:"output_rate_summary_packets"`
				RxLoad                   uint64 `json:"rx_load" xml:"rx_load"`
				TxLoad                   uint64 `json:"tx_load" xml:"tx_load"`
			}
		}
	}
	LastClearCountersEvent string `json:"last_cleared_event" xml:"last_cleared_event"`
	LastLinkFlappedEvent   string `json:"link_flapped_event" xml:"link_flapped_event"`
	Props                  struct {
		BeaconEnabled          bool   `json:"beacon_enabled" xml:"beacon_enabled"`
		AutoNegotiationEnabled bool   `json:"auto_negotiation_enabled" xml:"auto_negotiation_enabled"`
		MdixEnabled            bool   `json:"mdix_enabled" xml:"mdix_enabled"`
		BiaHwAddr              string `json:"bia_hw_address" xml:"bia_hw_address"`
		HwAddr                 string `json:"hw_address" xml:"hw_address"`
		HwDescription          string `json:"hw_description" xml:"hw_description"`
		MTU                    uint64 `json:"mtu" xml:"mtu"`
		Medium                 string `json:"medium" xml:"medium"`
		ParentInterface        string `json:"parent_interface" xml:"parent_interface"`
		ParentBundle           string `json:"parent_bundle" xml:"parent_bundle"`
		ShareState             string `json:"share_state" xml:"share_state"`
		State                  string `json:"state" xml:"state"`
		AdminState             string `json:"admin_state" xml:"admin_state"`
		StateReasonDetailed    string `json:"state_reason_detailed" xml:"state_reason_detailed"`
		Encapsulation          string `json:"encapsulation" xml:"encapsulation"`
		EtherType              string `json:"ethertype" xml:"ethertype"`
		EncapsulatedVlan       uint64 `json:"encapsulated_vlan" xml:"encapsulated_vlan"`
		Media                  string `json:"media" xml:"media"`
		Duplex                 string `json:"duplex" xml:"duplex"`
		BundleMembers          string `json:"bundle_members" xml:"bundle_members"`
		EEEState               string `json:"eee_state" xml:"eee_state"`
		IPAddress              string `json:"ip_address" xml:"ip_address"`
		IPMask                 uint64 `json:"ip_mask" xml:"ip_mask"`
		RateMode               string `json:"rate_mode" xml:"rate_mode"`
		Mode                   string `json:"mode" xml:"mode"`
		Speed                  string `json:"speed" xml:"speed"`
		SwitchportMonitor      string `json:"switchport_monitor" xml:"switchport_monitor"`
	}
	Metrics struct {
		Bandwidth   uint64 `json:"bandwidth" xml:"bandwidth"`
		Delay       uint64 `json:"delay" xml:"delay"`
		Reliability uint64 `json:"reliability" xml:"reliability"`
		Rxload      uint64 `json:"rx_load" xml:"rx_load"`
		Txload      uint64 `json:"tx_load" xml:"tx_load"`
	}
}

// NewInterfacesFromString returns Interface instance from an input string.
func NewInterfacesFromString(s string) ([]*Interface, error) {
	return NewInterfacesFromBytes([]byte(s))
}

func parseInterfaceInfo(i int, j *interfaceResponseResultBodyInterfaceRow) *Interface {
	//spew.Dump(j)
	intf := &Interface{}
	intf.Name = j.Interface
	intf.Description = j.Desc
	intf.LocalIndex = i
	intf.Props.State = j.State
	intf.Props.AdminState = j.AdminState
	intf.Props.Encapsulation = j.Encapsulation

	// INFO: output packets and bytes
	intf.Counters.InputBytes = j.EthInbytes
	intf.Counters.InputPackets = j.EthInpkts
	intf.Counters.InputUnicastPackets = j.EthInucast
	intf.Counters.InputBroadcastPackets = j.EthInbcast
	intf.Counters.InputMulticastPackets = j.EthInmcast

	// INFO: input packets and bytes
	intf.Counters.OutputBytes = j.EthOutbytes
	intf.Counters.OutputPackets = j.EthOutpkts
	intf.Counters.OutputUnicastPackets = j.EthOutucast
	intf.Counters.OutputBroadcastPackets = j.EthOutbcast
	intf.Counters.OutputMulticastPackets = j.EthOutmcast

	// INFO: other ethernet counters
	intf.Counters.Runts = j.EthRunts
	intf.Counters.NoBufferReceivedErrors = j.EthNobuf
	intf.Counters.Resets = j.EthResetCntr

	// INFO: other ethernet props
	intf.LastLinkFlappedEvent = j.EthLinkFlapped
	intf.LastClearCountersEvent = j.EthClearCounters
	intf.Props.Medium = j.Medium
	intf.Props.ParentInterface = j.ParentInterface
	intf.Props.ShareState = j.ShareState
	intf.Props.StateReasonDetailed = j.StateRsnDesc
	intf.Props.HwDescription = j.EthHwDesc
	intf.Props.EtherType = j.EthEthertype
	intf.Props.EncapsulatedVlan = j.EthEncapVlan
	intf.Props.Media = j.EthMedia
	intf.Counters.InputFlowControl = j.EthInFlowctrl
	intf.Counters.OutputFlowControl = j.EthOutFlowctrl
	intf.Props.ParentBundle = j.EthBundle
	intf.Props.Duplex = j.EthDuplex
	intf.Props.BundleMembers = j.EthMembers
	intf.Props.EEEState = j.EthEeeState
	intf.Props.IPAddress = strings.TrimSpace(j.EthIPAddr)
	intf.Props.IPMask = j.EthIPMask
	intf.Props.RateMode = j.EthMode
	intf.Props.Mode = j.EthMode
	intf.Props.Speed = j.EthSpeed
	intf.Props.SwitchportMonitor = j.EthSwtMonitor

	// INFO: counters for intervals
	if i, err := strconv.ParseUint(j.EthInrate1Bits, 10, 64); err == nil {
		intf.Counters.Intervals.Interval1.InputRateBits = i
	}
	if i, err := strconv.ParseUint(j.EthInrate1Pkts, 10, 64); err == nil {
		intf.Counters.Intervals.Interval1.InputRatePackets = i
	}
	if i, err := strconv.ParseUint(j.EthInrate2Bits, 10, 64); err == nil {
		intf.Counters.Intervals.Interval2.InputRateBits = i
	}
	if i, err := strconv.ParseUint(j.EthInrate2Pkts, 10, 64); err == nil {
		intf.Counters.Intervals.Interval2.InputRatePackets = i
	}
	if i, err := strconv.ParseUint(j.EthInrate3Bits, 10, 64); err == nil {
		intf.Counters.Intervals.Interval3.InputRateBits = i
	}
	if i, err := strconv.ParseUint(j.EthInrate3Pkts, 10, 64); err == nil {
		intf.Counters.Intervals.Interval3.InputRatePackets = i
	}

	if i, err := strconv.ParseUint(j.EthOutrate1Bits, 10, 64); err == nil {
		intf.Counters.Intervals.Interval1.OutputRateBits = i
	}
	if i, err := strconv.ParseUint(j.EthOutrate1Pkts, 10, 64); err == nil {
		intf.Counters.Intervals.Interval1.OutputRatePackets = i
	}
	if i, err := strconv.ParseUint(j.EthOutrate2Bits, 10, 64); err == nil {
		intf.Counters.Intervals.Interval2.OutputRateBits = i
	}
	if i, err := strconv.ParseUint(j.EthOutrate2Pkts, 10, 64); err == nil {
		intf.Counters.Intervals.Interval2.OutputRatePackets = i
	}
	if i, err := strconv.ParseUint(j.EthOutrate3Bits, 10, 64); err == nil {
		intf.Counters.Intervals.Interval3.OutputRateBits = i
	}
	if i, err := strconv.ParseUint(j.EthOutrate3Pkts, 10, 64); err == nil {
		intf.Counters.Intervals.Interval3.OutputRatePackets = i
	}

	intf.Counters.Intervals.Interval1.InputRateSummaryBits = j.EthInrate1SummaryBits
	intf.Counters.Intervals.Interval1.InputRateSummaryPackets = j.EthInrate1SummaryPkts
	intf.Counters.Intervals.Interval2.InputRateSummaryBits = j.EthInrate2SummaryBits
	intf.Counters.Intervals.Interval2.InputRateSummaryPackets = j.EthInrate2SummaryPkts
	intf.Counters.Intervals.Interval3.InputRateSummaryBits = j.EthInrate3SummaryBits
	intf.Counters.Intervals.Interval3.InputRateSummaryPackets = j.EthInrate3SummaryPkts
	intf.Counters.Intervals.Interval1.OutputRateSummaryBits = j.EthOutrate1SummaryBits
	intf.Counters.Intervals.Interval1.OutputRateSummaryPackets = j.EthOutrate1SummaryPkts
	intf.Counters.Intervals.Interval2.OutputRateSummaryBits = j.EthOutrate2SummaryBits
	intf.Counters.Intervals.Interval2.OutputRateSummaryPackets = j.EthOutrate2SummaryPkts
	intf.Counters.Intervals.Interval3.OutputRateSummaryBits = j.EthOutrate3SummaryBits
	intf.Counters.Intervals.Interval3.OutputRateSummaryPackets = j.EthOutrate3SummaryPkts

	intf.Counters.Intervals.Interval1.RxLoad = j.EthLoadInterval1Rx
	if i, err := strconv.ParseUint(j.EthLoadInterval1Tx, 10, 64); err == nil {
		intf.Counters.Intervals.Interval1.TxLoad = i
	}
	if i, err := strconv.ParseUint(j.EthLoadInterval2Rx, 10, 64); err == nil {
		intf.Counters.Intervals.Interval2.RxLoad = i
	}
	if i, err := strconv.ParseUint(j.EthLoadInterval2Tx, 10, 64); err == nil {
		intf.Counters.Intervals.Interval2.TxLoad = i
	}
	if i, err := strconv.ParseUint(j.EthLoadInterval3Rx, 10, 64); err == nil {
		intf.Counters.Intervals.Interval3.RxLoad = i
	}
	if i, err := strconv.ParseUint(j.EthLoadInterval3Tx, 10, 64); err == nil {
		intf.Counters.Intervals.Interval3.TxLoad = i
	}

	// INFO: routing metrics
	intf.Metrics.Bandwidth = j.EthBw
	intf.Metrics.Delay = j.EthDly
	if i, err := strconv.ParseUint(j.EthReliability, 10, 64); err == nil {
		intf.Metrics.Reliability = i
	}
	if i, err := strconv.ParseUint(j.EthRxload, 10, 64); err == nil {
		intf.Metrics.Rxload = i
	}
	if i, err := strconv.ParseUint(j.EthTxload, 10, 64); err == nil {
		intf.Metrics.Txload = i
	}

	// INTO: MTU
	if i, err := strconv.ParseUint(j.EthMtu, 10, 64); err == nil {
		intf.Props.MTU = i
	}

	// INFO: eth_babbles is a string with numeric value
	if i, err := strconv.ParseUint(j.EthBabbles, 10, 64); err == nil {
		intf.Counters.Babbles = i
	}
	// INFO: eth_bad_eth (bad ether type drop) is a string with numeric value
	if i, err := strconv.ParseUint(j.EthBadEth, 10, 64); err == nil {
		intf.Counters.BadEtherTypeDrops = i
	}
	// INFO: eth_bad_proto is a string with numeric value
	if i, err := strconv.ParseUint(j.EthBadProto, 10, 64); err == nil {
		intf.Counters.BadProtocolDrops = i
	}
	// INFO: eth_nocarrier is a string with numeric value
	if i, err := strconv.ParseUint(j.EthNocarrier, 10, 64); err == nil {
		intf.Counters.NoCarrier = i
	}
	// INFO: eth_dribble is a string with numeric value
	if i, err := strconv.ParseUint(j.EthDribble, 10, 64); err == nil {
		intf.Counters.Dribble = i
	}
	// INFO: eth_frame is a string with numeric value
	if i, err := strconv.ParseUint(j.EthFrame, 10, 64); err == nil {
		intf.Counters.InputFrameErrors = i
	}
	// INFO: eth_in_ifdown_drops is a string with numeric value
	if i, err := strconv.ParseUint(j.EthInIfdownDrops, 10, 64); err == nil {
		intf.Counters.InputIfaceDownDrops = i
	}
	// INFO: eth_ignored is a string with numeric value
	if i, err := strconv.ParseUint(j.EthIgnored, 10, 64); err == nil {
		intf.Counters.Ignored = i
	}
	// INFO: eth_indiscard is a string with numeric value
	if i, err := strconv.ParseUint(j.EthIndiscard, 10, 64); err == nil {
		intf.Counters.InputDiscards = i
	}
	// INFO: eth_inerr is a string with numeric value
	if i, err := strconv.ParseUint(j.EthInerr, 10, 64); err == nil {
		intf.Counters.InputErrors = i
	}
	// INFO: eth_inpause is a string with numeric value
	if i, err := strconv.ParseUint(j.EthInpause, 10, 64); err == nil {
		intf.Counters.InputPause = i
	}
	// INFO: eth_latecoll is a string with numeric value
	if i, err := strconv.ParseUint(j.EthLatecoll, 10, 64); err == nil {
		intf.Counters.LateCollisions = i
	}
	// INFO: eth_lostcarrier is a string with numeric value
	if i, err := strconv.ParseUint(j.EthLostcarrier, 10, 64); err == nil {
		intf.Counters.LostCarrier = i
	}
	// INFO: eth_outdiscard is a string with numeric value
	if i, err := strconv.ParseUint(j.EthOutdiscard, 10, 64); err == nil {
		intf.Counters.OutputDiscards = i
	}
	// INFO: eth_outerr is a string with numeric value
	if i, err := strconv.ParseUint(j.EthOuterr, 10, 64); err == nil {
		intf.Counters.OutputErrors = i
	}
	// INFO: eth_outpause is a string with numeric value
	if i, err := strconv.ParseUint(j.EthOutpause, 10, 64); err == nil {
		intf.Counters.OutputPause = i
	}
	// INFO: eth_overrun is a string with numeric value
	if i, err := strconv.ParseUint(j.EthOverrun, 10, 64); err == nil {
		intf.Counters.InputOverruns = i
	}
	// INFO: eth_storm_supp is a string with numeric value
	if i, err := strconv.ParseUint(j.EthStormSupp, 10, 64); err == nil {
		intf.Counters.StormSuppression = i
	}
	// INFO: eth_underrun is a string with numeric value
	if i, err := strconv.ParseUint(j.EthUnderrun, 10, 64); err == nil {
		intf.Counters.OutputUnderruns = i
	}
	// INFO: eth_watchdog is a string with numeric value
	if i, err := strconv.ParseUint(j.EthWatchdog, 10, 64); err == nil {
		intf.Counters.Watchdog = i
	}
	// WARN: eth_jumbo_inpkts is a string with numeric value
	if i, err := strconv.ParseUint(j.EthJumboInpkts, 10, 64); err == nil {
		intf.Counters.InputJumboPackets = i
	}
	// WARN: eth_jumbo_outpkts is a string with numeric value
	if i, err := strconv.ParseUint(j.EthJumboOutpkts, 10, 64); err == nil {
		intf.Counters.OutputJumboPackets = i
	}
	// INFO: eth_coll is a string with numeric value
	if i, err := strconv.ParseUint(j.EthColl, 10, 64); err == nil {
		intf.Counters.Collisions = i
	}
	// INFO: eth_crc is a string with numeric value
	if i, err := strconv.ParseUint(j.EthCrc, 10, 64); err == nil {
		intf.Counters.CrcErrors = i
	}
	// INFO: eth_crc is a string with numeric value
	if i, err := strconv.ParseUint(j.EthCrc, 10, 64); err == nil {
		intf.Counters.CrcErrors = i
	}
	// INFO: eth_deferred is a string with numeric value
	if i, err := strconv.ParseUint(j.EthDeferred, 10, 64); err == nil {
		intf.Counters.Deferred = i
	}

	// Loopback-specific

	// INFO: loop_in_compressed is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopInCompressed, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.InputCompressed = i
		}
	}
	// INFO: loop_in_errors is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopInErrors, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.InputErrors = i
		}
	}
	// INFO: loop_in_fifo is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopInFifo, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.InputFifo = i
		}
	}
	// INFO: loop_in_frame is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopInFrame, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.InputFrameErrors = i
		}
	}
	// INFO: loop_in_mcast is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopInMcast, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.InputMulticastPackets = i
		}
	}
	// INFO: loop_in_overrun is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopInOverrun, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.InputOverruns = i
		}
	}
	// INFO: loop_out_carriers is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopOutCarriers, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.OutputCarrierErrors = i
		}
	}
	// INFO: loop_out_collisions is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopOutCollisions, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.Collisions = i
		}
	}
	// INFO: loop_out_errors is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopOutErrors, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.OutputErrors = i
		}
	}
	// INFO: loop_out_fifo is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopOutFifo, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.OutputFifo = i
		}
	}
	// WARN: loop_out_underruns is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopOutUnderruns, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.OutputUnderruns = i
		}
	}
	// WARN: loop_in_bytes is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopInBytes, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.InputBytes = i
		}
	}
	// WARN: loop_out_bytes is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopOutBytes, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.OutputBytes = i
		}
	}
	// WARN: loop_out_pkts is a string with numeric value
	if i, err := strconv.ParseUint(j.LoopOutPkts, 10, 64); err == nil {
		if i > 0 {
			intf.Counters.OutputPackets = i
		}
	}
	// WARN: inconsistency between loop_out_pkts and loop_in_pkts
	if j.LoopInPkts > 0 {
		intf.Counters.InputPackets = j.LoopInPkts
	}

	switch j.EthBeacon {
	case "on":
		intf.Props.BeaconEnabled = true
	default:
		intf.Props.BeaconEnabled = false
	}

	switch j.EthAutoneg {
	case "on":
		intf.Props.AutoNegotiationEnabled = true
	default:
		intf.Props.AutoNegotiationEnabled = false
	}

	switch j.EthMdix {
	case "on":
		intf.Props.MdixEnabled = true
	default:
		intf.Props.MdixEnabled = false
	}

	intf.Props.BiaHwAddr = strings.TrimSpace(j.EthBiaAddr)
	intf.Props.HwAddr = strings.TrimSpace(j.EthHwAddr)

	// INFO: svi/vlan interface
	if j.SviAdminState != "" {
		intf.Props.BiaHwAddr = strings.TrimSpace(j.SviMac)
		intf.Props.HwAddr = strings.TrimSpace(j.SviMac)
		intf.Props.AdminState = strings.TrimSpace(j.SviAdminState)
		if i, err := strconv.ParseUint(j.SviBw, 10, 64); err == nil {
			intf.Metrics.Bandwidth = i
		}
		if i, err := strconv.ParseUint(j.SviDelay, 10, 64); err == nil {
			intf.Metrics.Delay = i
		}
		if i, err := strconv.ParseUint(j.SviReliability, 10, 64); err == nil {
			intf.Metrics.Reliability = i
		}
		if i, err := strconv.ParseUint(j.SviRxLoad, 10, 64); err == nil {
			intf.Metrics.Rxload = i
		}
		intf.Metrics.Txload = j.SviTxLoad
		if i, err := strconv.ParseUint(j.SviMtu, 10, 64); err == nil {
			intf.Props.MTU = i
		}
		intf.LastClearCountersEvent = j.SviTimeLastCleared
		intf.Props.StateReasonDetailed = j.StateRsnDesc
		intf.Props.State = j.SviLineProto
		intf.Props.IPAddress = strings.TrimSpace(j.SviIPAddr)
		if i, err := strconv.ParseUint(j.SviIPMask, 10, 64); err == nil {
			intf.Props.IPMask = i
		}
		if i, err := strconv.ParseUint(j.SviUcastBytesIn, 10, 64); err == nil {
			intf.Counters.InputUnicastBytes = i
		}
		if i, err := strconv.ParseUint(j.SviUcastPktsIn, 10, 64); err == nil {
			intf.Counters.InputUnicastPackets = i
		}
	}

	// INFO: mgmt interface
	if j.VdcLvlInAvgBits > 0 {
		// WONT-DO: j.VdcLvlInAvgBits
		// WONT-DO: j.VdcLvlInAvgPkts
		// WONT-DO: j.VdcLvlOutAvgBits
		// WONT-DO: j.VdcLvlOutAvgPkts
		intf.Counters.InputPackets = j.VdcLvlInPkts
		if i, err := strconv.ParseUint(j.VdcLvlInBytes, 10, 64); err == nil {
			intf.Counters.InputBytes = i
		}
		if i, err := strconv.ParseUint(j.VdcLvlInUcast, 10, 64); err == nil {
			intf.Counters.InputUnicastPackets = i
		}
		if i, err := strconv.ParseUint(j.VdcLvlInBcast, 10, 64); err == nil {
			intf.Counters.InputBroadcastPackets = i
		}
		if i, err := strconv.ParseUint(j.VdcLvlInMcast, 10, 64); err == nil {
			intf.Counters.InputMulticastPackets = i
		}
		if i, err := strconv.ParseUint(j.VdcLvlOutBytes, 10, 64); err == nil {
			intf.Counters.OutputBytes = i
		}
		if i, err := strconv.ParseUint(j.VdcLvlOutPkts, 10, 64); err == nil {
			intf.Counters.OutputPackets = i
		}
		if i, err := strconv.ParseUint(j.VdcLvlOutUcast, 10, 64); err == nil {
			intf.Counters.OutputUnicastPackets = i
		}
		if i, err := strconv.ParseUint(j.VdcLvlOutBcast, 10, 64); err == nil {
			intf.Counters.OutputBroadcastPackets = i
		}
		if i, err := strconv.ParseUint(j.VdcLvlOutMcast, 10, 64); err == nil {
			intf.Counters.OutputMulticastPackets = i
		}
	}
	return intf
}

// NewInterfacesFromBytes returns a list of Interface instance from an input byte array.
func NewInterfacesFromBytes(s []byte) ([]*Interface, error) {
	var interfaces []*Interface
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
		return nil, fmt.Errorf("parsing body error: %v", err)
	}
	var intfResult interfacesResponseResultBody
	err = json.Unmarshal(body.Body, &intfResult)
	if err != nil {
		return nil, fmt.Errorf("parsing interface result error: %v", err)
	}

	if len(intfResult.InterfaceTable.InterfaceRow) == 0 {
		return nil, fmt.Errorf("no interfaces found")
	}
	for i, j := range intfResult.InterfaceTable.InterfaceRow {
		intf := parseInterfaceInfo(i, &j)
		interfaces = append(interfaces, intf)
		//spew.Dump(intf)
	}
	return interfaces, nil
}

// NewInterfaceFromBytes returns an Interface instance from an input byte array.
func NewInterfaceFromBytes(s []byte) (*Interface, error) {
	var interfaces *Interface
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
		return nil, fmt.Errorf("parsing body error: %v", err)
	}
	var intfResult interfaceOneResponseResultBody
	err = json.Unmarshal(body.Body, &intfResult)
	if err != nil {
		return nil, fmt.Errorf("parsing interface result error: %v", err)
	}

	interfaces = parseInterfaceInfo(0, &intfResult.InterfaceTable.InterfaceRow)
	return interfaces, nil
}
