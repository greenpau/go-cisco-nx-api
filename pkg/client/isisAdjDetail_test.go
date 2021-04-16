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
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseShowIsisAdjDetailJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *IsisAdjDetailResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.isis.2.adj.det",
			exp: &IsisAdjDetailResponse{
				InsAPI: struct {
					Outputs struct {
						Output IsisAdjDetailResponseResult "json:\"output\""
					} "json:\"outputs\""
					Sid     string "json:\"sid\""
					Type    string "json:\"type\""
					Version string "json:\"version\""
				}{Outputs: struct {
					Output IsisAdjDetailResponseResult "json:\"output\""
				}{Output: IsisAdjDetailResponseResult{Body: IsisAdjDetailResultBody{TableProcessTag: []struct {
					RowProcessTag []struct {
						ProcessTagOut string "json:\"process-tag-out\""
						TableVrf      []struct {
							RowVrf []struct {
								VrfNameOut      string "json:\"vrf-name-out\""
								AdjSummaryOut   string "json:\"adj-summary-out\""
								AdjInterfaceOut string "json:\"adj-interface-out\""
								TableProcessAdj []struct {
									RowProcessAdj []struct {
										AdjSysNameOut              string "json:\"adj-sys-name-out\""
										AdjSysIDOut                string "json:\"adj-sys-id-out\""
										AdjUsageOut                string "json:\"adj-usage-out\""
										AdjStateOut                string "json:\"adj-state-out\""
										AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
										AdjIntfNameOut             string "json:\"adj-intf-name-out\""
										AdjDetailSetOut            string "json:\"adj-detail-set-out\""
										AdjTransitionsOut          string "json:\"adj-transitions-out\""
										AdjFlapOut                 string "json:\"adj-flap-out\""
										AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
										AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
										AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
										AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
										AdjBcastOut                string "json:\"adj-bcast-out\""
										AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
										AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
										AdjResurrectOut            string "json:\"adj-resurrect-out\""
										AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
										AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
										AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
										AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
										AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
										AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
										AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
										AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
										TableAdjSid                struct {
											RowAdjSid []struct {
												AdjSidValue  string "json:\"adj-sid-value\""
												AdjSidFFlag  string "json:\"adj-sid-f-flag\""
												AdjSidBFlag  string "json:\"adj-sid-b-flag\""
												AdjSidVFlag  string "json:\"adj-sid-v-flag\""
												AdjSidLFlag  string "json:\"adj-sid-l-flag\""
												AdjSidSFlag  string "json:\"adj-sid-s-flag\""
												AdjSidPFlag  string "json:\"adj-sid-p-flag\""
												AdjSidWeight string "json:\"adj-sid-weight\""
											} "json:\"ROW_adj_sid\""
										} "json:\"TABLE_adj_sid\""
									} "json:\"ROW_process_adj\""
								} "json:\"TABLE_process_adj\""
							} "json:\"ROW_vrf\""
						} "json:\"TABLE_vrf\""
					} "json:\"ROW_process_tag\""
				}{struct {
					RowProcessTag []struct {
						ProcessTagOut string "json:\"process-tag-out\""
						TableVrf      []struct {
							RowVrf []struct {
								VrfNameOut      string "json:\"vrf-name-out\""
								AdjSummaryOut   string "json:\"adj-summary-out\""
								AdjInterfaceOut string "json:\"adj-interface-out\""
								TableProcessAdj []struct {
									RowProcessAdj []struct {
										AdjSysNameOut              string "json:\"adj-sys-name-out\""
										AdjSysIDOut                string "json:\"adj-sys-id-out\""
										AdjUsageOut                string "json:\"adj-usage-out\""
										AdjStateOut                string "json:\"adj-state-out\""
										AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
										AdjIntfNameOut             string "json:\"adj-intf-name-out\""
										AdjDetailSetOut            string "json:\"adj-detail-set-out\""
										AdjTransitionsOut          string "json:\"adj-transitions-out\""
										AdjFlapOut                 string "json:\"adj-flap-out\""
										AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
										AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
										AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
										AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
										AdjBcastOut                string "json:\"adj-bcast-out\""
										AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
										AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
										AdjResurrectOut            string "json:\"adj-resurrect-out\""
										AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
										AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
										AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
										AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
										AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
										AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
										AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
										AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
										TableAdjSid                struct {
											RowAdjSid []struct {
												AdjSidValue  string "json:\"adj-sid-value\""
												AdjSidFFlag  string "json:\"adj-sid-f-flag\""
												AdjSidBFlag  string "json:\"adj-sid-b-flag\""
												AdjSidVFlag  string "json:\"adj-sid-v-flag\""
												AdjSidLFlag  string "json:\"adj-sid-l-flag\""
												AdjSidSFlag  string "json:\"adj-sid-s-flag\""
												AdjSidPFlag  string "json:\"adj-sid-p-flag\""
												AdjSidWeight string "json:\"adj-sid-weight\""
											} "json:\"ROW_adj_sid\""
										} "json:\"TABLE_adj_sid\""
									} "json:\"ROW_process_adj\""
								} "json:\"TABLE_process_adj\""
							} "json:\"ROW_vrf\""
						} "json:\"TABLE_vrf\""
					} "json:\"ROW_process_tag\""
				}{RowProcessTag: []struct {
					ProcessTagOut string "json:\"process-tag-out\""
					TableVrf      []struct {
						RowVrf []struct {
							VrfNameOut      string "json:\"vrf-name-out\""
							AdjSummaryOut   string "json:\"adj-summary-out\""
							AdjInterfaceOut string "json:\"adj-interface-out\""
							TableProcessAdj []struct {
								RowProcessAdj []struct {
									AdjSysNameOut              string "json:\"adj-sys-name-out\""
									AdjSysIDOut                string "json:\"adj-sys-id-out\""
									AdjUsageOut                string "json:\"adj-usage-out\""
									AdjStateOut                string "json:\"adj-state-out\""
									AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
									AdjIntfNameOut             string "json:\"adj-intf-name-out\""
									AdjDetailSetOut            string "json:\"adj-detail-set-out\""
									AdjTransitionsOut          string "json:\"adj-transitions-out\""
									AdjFlapOut                 string "json:\"adj-flap-out\""
									AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
									AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
									AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
									AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
									AdjBcastOut                string "json:\"adj-bcast-out\""
									AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
									AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
									AdjResurrectOut            string "json:\"adj-resurrect-out\""
									AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
									AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
									AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
									AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
									AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
									AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
									AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
									AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
									TableAdjSid                struct {
										RowAdjSid []struct {
											AdjSidValue  string "json:\"adj-sid-value\""
											AdjSidFFlag  string "json:\"adj-sid-f-flag\""
											AdjSidBFlag  string "json:\"adj-sid-b-flag\""
											AdjSidVFlag  string "json:\"adj-sid-v-flag\""
											AdjSidLFlag  string "json:\"adj-sid-l-flag\""
											AdjSidSFlag  string "json:\"adj-sid-s-flag\""
											AdjSidPFlag  string "json:\"adj-sid-p-flag\""
											AdjSidWeight string "json:\"adj-sid-weight\""
										} "json:\"ROW_adj_sid\""
									} "json:\"TABLE_adj_sid\""
								} "json:\"ROW_process_adj\""
							} "json:\"TABLE_process_adj\""
						} "json:\"ROW_vrf\""
					} "json:\"TABLE_vrf\""
				}{struct {
					ProcessTagOut string "json:\"process-tag-out\""
					TableVrf      []struct {
						RowVrf []struct {
							VrfNameOut      string "json:\"vrf-name-out\""
							AdjSummaryOut   string "json:\"adj-summary-out\""
							AdjInterfaceOut string "json:\"adj-interface-out\""
							TableProcessAdj []struct {
								RowProcessAdj []struct {
									AdjSysNameOut              string "json:\"adj-sys-name-out\""
									AdjSysIDOut                string "json:\"adj-sys-id-out\""
									AdjUsageOut                string "json:\"adj-usage-out\""
									AdjStateOut                string "json:\"adj-state-out\""
									AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
									AdjIntfNameOut             string "json:\"adj-intf-name-out\""
									AdjDetailSetOut            string "json:\"adj-detail-set-out\""
									AdjTransitionsOut          string "json:\"adj-transitions-out\""
									AdjFlapOut                 string "json:\"adj-flap-out\""
									AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
									AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
									AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
									AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
									AdjBcastOut                string "json:\"adj-bcast-out\""
									AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
									AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
									AdjResurrectOut            string "json:\"adj-resurrect-out\""
									AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
									AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
									AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
									AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
									AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
									AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
									AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
									AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
									TableAdjSid                struct {
										RowAdjSid []struct {
											AdjSidValue  string "json:\"adj-sid-value\""
											AdjSidFFlag  string "json:\"adj-sid-f-flag\""
											AdjSidBFlag  string "json:\"adj-sid-b-flag\""
											AdjSidVFlag  string "json:\"adj-sid-v-flag\""
											AdjSidLFlag  string "json:\"adj-sid-l-flag\""
											AdjSidSFlag  string "json:\"adj-sid-s-flag\""
											AdjSidPFlag  string "json:\"adj-sid-p-flag\""
											AdjSidWeight string "json:\"adj-sid-weight\""
										} "json:\"ROW_adj_sid\""
									} "json:\"TABLE_adj_sid\""
								} "json:\"ROW_process_adj\""
							} "json:\"TABLE_process_adj\""
						} "json:\"ROW_vrf\""
					} "json:\"TABLE_vrf\""
				}{ProcessTagOut: "2", TableVrf: []struct {
					RowVrf []struct {
						VrfNameOut      string "json:\"vrf-name-out\""
						AdjSummaryOut   string "json:\"adj-summary-out\""
						AdjInterfaceOut string "json:\"adj-interface-out\""
						TableProcessAdj []struct {
							RowProcessAdj []struct {
								AdjSysNameOut              string "json:\"adj-sys-name-out\""
								AdjSysIDOut                string "json:\"adj-sys-id-out\""
								AdjUsageOut                string "json:\"adj-usage-out\""
								AdjStateOut                string "json:\"adj-state-out\""
								AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
								AdjIntfNameOut             string "json:\"adj-intf-name-out\""
								AdjDetailSetOut            string "json:\"adj-detail-set-out\""
								AdjTransitionsOut          string "json:\"adj-transitions-out\""
								AdjFlapOut                 string "json:\"adj-flap-out\""
								AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
								AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
								AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
								AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
								AdjBcastOut                string "json:\"adj-bcast-out\""
								AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
								AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
								AdjResurrectOut            string "json:\"adj-resurrect-out\""
								AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
								AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
								AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
								AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
								AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
								AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
								AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
								AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
								TableAdjSid                struct {
									RowAdjSid []struct {
										AdjSidValue  string "json:\"adj-sid-value\""
										AdjSidFFlag  string "json:\"adj-sid-f-flag\""
										AdjSidBFlag  string "json:\"adj-sid-b-flag\""
										AdjSidVFlag  string "json:\"adj-sid-v-flag\""
										AdjSidLFlag  string "json:\"adj-sid-l-flag\""
										AdjSidSFlag  string "json:\"adj-sid-s-flag\""
										AdjSidPFlag  string "json:\"adj-sid-p-flag\""
										AdjSidWeight string "json:\"adj-sid-weight\""
									} "json:\"ROW_adj_sid\""
								} "json:\"TABLE_adj_sid\""
							} "json:\"ROW_process_adj\""
						} "json:\"TABLE_process_adj\""
					} "json:\"ROW_vrf\""
				}{struct {
					RowVrf []struct {
						VrfNameOut      string "json:\"vrf-name-out\""
						AdjSummaryOut   string "json:\"adj-summary-out\""
						AdjInterfaceOut string "json:\"adj-interface-out\""
						TableProcessAdj []struct {
							RowProcessAdj []struct {
								AdjSysNameOut              string "json:\"adj-sys-name-out\""
								AdjSysIDOut                string "json:\"adj-sys-id-out\""
								AdjUsageOut                string "json:\"adj-usage-out\""
								AdjStateOut                string "json:\"adj-state-out\""
								AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
								AdjIntfNameOut             string "json:\"adj-intf-name-out\""
								AdjDetailSetOut            string "json:\"adj-detail-set-out\""
								AdjTransitionsOut          string "json:\"adj-transitions-out\""
								AdjFlapOut                 string "json:\"adj-flap-out\""
								AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
								AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
								AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
								AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
								AdjBcastOut                string "json:\"adj-bcast-out\""
								AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
								AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
								AdjResurrectOut            string "json:\"adj-resurrect-out\""
								AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
								AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
								AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
								AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
								AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
								AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
								AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
								AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
								TableAdjSid                struct {
									RowAdjSid []struct {
										AdjSidValue  string "json:\"adj-sid-value\""
										AdjSidFFlag  string "json:\"adj-sid-f-flag\""
										AdjSidBFlag  string "json:\"adj-sid-b-flag\""
										AdjSidVFlag  string "json:\"adj-sid-v-flag\""
										AdjSidLFlag  string "json:\"adj-sid-l-flag\""
										AdjSidSFlag  string "json:\"adj-sid-s-flag\""
										AdjSidPFlag  string "json:\"adj-sid-p-flag\""
										AdjSidWeight string "json:\"adj-sid-weight\""
									} "json:\"ROW_adj_sid\""
								} "json:\"TABLE_adj_sid\""
							} "json:\"ROW_process_adj\""
						} "json:\"TABLE_process_adj\""
					} "json:\"ROW_vrf\""
				}{RowVrf: []struct {
					VrfNameOut      string "json:\"vrf-name-out\""
					AdjSummaryOut   string "json:\"adj-summary-out\""
					AdjInterfaceOut string "json:\"adj-interface-out\""
					TableProcessAdj []struct {
						RowProcessAdj []struct {
							AdjSysNameOut              string "json:\"adj-sys-name-out\""
							AdjSysIDOut                string "json:\"adj-sys-id-out\""
							AdjUsageOut                string "json:\"adj-usage-out\""
							AdjStateOut                string "json:\"adj-state-out\""
							AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
							AdjIntfNameOut             string "json:\"adj-intf-name-out\""
							AdjDetailSetOut            string "json:\"adj-detail-set-out\""
							AdjTransitionsOut          string "json:\"adj-transitions-out\""
							AdjFlapOut                 string "json:\"adj-flap-out\""
							AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
							AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
							AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
							AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
							AdjBcastOut                string "json:\"adj-bcast-out\""
							AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
							AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
							AdjResurrectOut            string "json:\"adj-resurrect-out\""
							AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
							AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
							AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
							AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
							AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
							AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
							AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
							AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
							TableAdjSid                struct {
								RowAdjSid []struct {
									AdjSidValue  string "json:\"adj-sid-value\""
									AdjSidFFlag  string "json:\"adj-sid-f-flag\""
									AdjSidBFlag  string "json:\"adj-sid-b-flag\""
									AdjSidVFlag  string "json:\"adj-sid-v-flag\""
									AdjSidLFlag  string "json:\"adj-sid-l-flag\""
									AdjSidSFlag  string "json:\"adj-sid-s-flag\""
									AdjSidPFlag  string "json:\"adj-sid-p-flag\""
									AdjSidWeight string "json:\"adj-sid-weight\""
								} "json:\"ROW_adj_sid\""
							} "json:\"TABLE_adj_sid\""
						} "json:\"ROW_process_adj\""
					} "json:\"TABLE_process_adj\""
				}{struct {
					VrfNameOut      string "json:\"vrf-name-out\""
					AdjSummaryOut   string "json:\"adj-summary-out\""
					AdjInterfaceOut string "json:\"adj-interface-out\""
					TableProcessAdj []struct {
						RowProcessAdj []struct {
							AdjSysNameOut              string "json:\"adj-sys-name-out\""
							AdjSysIDOut                string "json:\"adj-sys-id-out\""
							AdjUsageOut                string "json:\"adj-usage-out\""
							AdjStateOut                string "json:\"adj-state-out\""
							AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
							AdjIntfNameOut             string "json:\"adj-intf-name-out\""
							AdjDetailSetOut            string "json:\"adj-detail-set-out\""
							AdjTransitionsOut          string "json:\"adj-transitions-out\""
							AdjFlapOut                 string "json:\"adj-flap-out\""
							AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
							AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
							AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
							AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
							AdjBcastOut                string "json:\"adj-bcast-out\""
							AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
							AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
							AdjResurrectOut            string "json:\"adj-resurrect-out\""
							AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
							AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
							AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
							AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
							AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
							AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
							AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
							AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
							TableAdjSid                struct {
								RowAdjSid []struct {
									AdjSidValue  string "json:\"adj-sid-value\""
									AdjSidFFlag  string "json:\"adj-sid-f-flag\""
									AdjSidBFlag  string "json:\"adj-sid-b-flag\""
									AdjSidVFlag  string "json:\"adj-sid-v-flag\""
									AdjSidLFlag  string "json:\"adj-sid-l-flag\""
									AdjSidSFlag  string "json:\"adj-sid-s-flag\""
									AdjSidPFlag  string "json:\"adj-sid-p-flag\""
									AdjSidWeight string "json:\"adj-sid-weight\""
								} "json:\"ROW_adj_sid\""
							} "json:\"TABLE_adj_sid\""
						} "json:\"ROW_process_adj\""
					} "json:\"TABLE_process_adj\""
				}{VrfNameOut: "default", AdjSummaryOut: "false", AdjInterfaceOut: "false", TableProcessAdj: []struct {
					RowProcessAdj []struct {
						AdjSysNameOut              string "json:\"adj-sys-name-out\""
						AdjSysIDOut                string "json:\"adj-sys-id-out\""
						AdjUsageOut                string "json:\"adj-usage-out\""
						AdjStateOut                string "json:\"adj-state-out\""
						AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
						AdjIntfNameOut             string "json:\"adj-intf-name-out\""
						AdjDetailSetOut            string "json:\"adj-detail-set-out\""
						AdjTransitionsOut          string "json:\"adj-transitions-out\""
						AdjFlapOut                 string "json:\"adj-flap-out\""
						AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
						AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
						AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
						AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
						AdjBcastOut                string "json:\"adj-bcast-out\""
						AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
						AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
						AdjResurrectOut            string "json:\"adj-resurrect-out\""
						AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
						AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
						AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
						AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
						AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
						AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
						AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
						AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
						TableAdjSid                struct {
							RowAdjSid []struct {
								AdjSidValue  string "json:\"adj-sid-value\""
								AdjSidFFlag  string "json:\"adj-sid-f-flag\""
								AdjSidBFlag  string "json:\"adj-sid-b-flag\""
								AdjSidVFlag  string "json:\"adj-sid-v-flag\""
								AdjSidLFlag  string "json:\"adj-sid-l-flag\""
								AdjSidSFlag  string "json:\"adj-sid-s-flag\""
								AdjSidPFlag  string "json:\"adj-sid-p-flag\""
								AdjSidWeight string "json:\"adj-sid-weight\""
							} "json:\"ROW_adj_sid\""
						} "json:\"TABLE_adj_sid\""
					} "json:\"ROW_process_adj\""
				}{struct {
					RowProcessAdj []struct {
						AdjSysNameOut              string "json:\"adj-sys-name-out\""
						AdjSysIDOut                string "json:\"adj-sys-id-out\""
						AdjUsageOut                string "json:\"adj-usage-out\""
						AdjStateOut                string "json:\"adj-state-out\""
						AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
						AdjIntfNameOut             string "json:\"adj-intf-name-out\""
						AdjDetailSetOut            string "json:\"adj-detail-set-out\""
						AdjTransitionsOut          string "json:\"adj-transitions-out\""
						AdjFlapOut                 string "json:\"adj-flap-out\""
						AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
						AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
						AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
						AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
						AdjBcastOut                string "json:\"adj-bcast-out\""
						AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
						AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
						AdjResurrectOut            string "json:\"adj-resurrect-out\""
						AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
						AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
						AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
						AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
						AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
						AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
						AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
						AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
						TableAdjSid                struct {
							RowAdjSid []struct {
								AdjSidValue  string "json:\"adj-sid-value\""
								AdjSidFFlag  string "json:\"adj-sid-f-flag\""
								AdjSidBFlag  string "json:\"adj-sid-b-flag\""
								AdjSidVFlag  string "json:\"adj-sid-v-flag\""
								AdjSidLFlag  string "json:\"adj-sid-l-flag\""
								AdjSidSFlag  string "json:\"adj-sid-s-flag\""
								AdjSidPFlag  string "json:\"adj-sid-p-flag\""
								AdjSidWeight string "json:\"adj-sid-weight\""
							} "json:\"ROW_adj_sid\""
						} "json:\"TABLE_adj_sid\""
					} "json:\"ROW_process_adj\""
				}{RowProcessAdj: []struct {
					AdjSysNameOut              string "json:\"adj-sys-name-out\""
					AdjSysIDOut                string "json:\"adj-sys-id-out\""
					AdjUsageOut                string "json:\"adj-usage-out\""
					AdjStateOut                string "json:\"adj-state-out\""
					AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
					AdjIntfNameOut             string "json:\"adj-intf-name-out\""
					AdjDetailSetOut            string "json:\"adj-detail-set-out\""
					AdjTransitionsOut          string "json:\"adj-transitions-out\""
					AdjFlapOut                 string "json:\"adj-flap-out\""
					AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
					AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
					AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
					AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
					AdjBcastOut                string "json:\"adj-bcast-out\""
					AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
					AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
					AdjResurrectOut            string "json:\"adj-resurrect-out\""
					AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
					AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
					AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
					AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
					AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
					AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
					AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
					AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
					TableAdjSid                struct {
						RowAdjSid []struct {
							AdjSidValue  string "json:\"adj-sid-value\""
							AdjSidFFlag  string "json:\"adj-sid-f-flag\""
							AdjSidBFlag  string "json:\"adj-sid-b-flag\""
							AdjSidVFlag  string "json:\"adj-sid-v-flag\""
							AdjSidLFlag  string "json:\"adj-sid-l-flag\""
							AdjSidSFlag  string "json:\"adj-sid-s-flag\""
							AdjSidPFlag  string "json:\"adj-sid-p-flag\""
							AdjSidWeight string "json:\"adj-sid-weight\""
						} "json:\"ROW_adj_sid\""
					} "json:\"TABLE_adj_sid\""
				}{struct {
					AdjSysNameOut              string "json:\"adj-sys-name-out\""
					AdjSysIDOut                string "json:\"adj-sys-id-out\""
					AdjUsageOut                string "json:\"adj-usage-out\""
					AdjStateOut                string "json:\"adj-state-out\""
					AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
					AdjIntfNameOut             string "json:\"adj-intf-name-out\""
					AdjDetailSetOut            string "json:\"adj-detail-set-out\""
					AdjTransitionsOut          string "json:\"adj-transitions-out\""
					AdjFlapOut                 string "json:\"adj-flap-out\""
					AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
					AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
					AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
					AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
					AdjBcastOut                string "json:\"adj-bcast-out\""
					AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
					AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
					AdjResurrectOut            string "json:\"adj-resurrect-out\""
					AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
					AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
					AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
					AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
					AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
					AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
					AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
					AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
					TableAdjSid                struct {
						RowAdjSid []struct {
							AdjSidValue  string "json:\"adj-sid-value\""
							AdjSidFFlag  string "json:\"adj-sid-f-flag\""
							AdjSidBFlag  string "json:\"adj-sid-b-flag\""
							AdjSidVFlag  string "json:\"adj-sid-v-flag\""
							AdjSidLFlag  string "json:\"adj-sid-l-flag\""
							AdjSidSFlag  string "json:\"adj-sid-s-flag\""
							AdjSidPFlag  string "json:\"adj-sid-p-flag\""
							AdjSidWeight string "json:\"adj-sid-weight\""
						} "json:\"ROW_adj_sid\""
					} "json:\"TABLE_adj_sid\""
				}{AdjSysNameOut: "n9k-reg-4", AdjSysIDOut: "N/A", AdjUsageOut: "2", AdjStateOut: "UP", AdjHoldTimeOut: "00:00:29", AdjIntfNameOut: "Ethernet1/21", AdjDetailSetOut: "true", AdjTransitionsOut: "1", AdjFlapOut: "true", AdjFlapTimeOut: "01:33:34", AdjCktTypeOut: "L2", AdjIpv4AddrOut: "45.1.1.1", AdjIpv6AddrOut: "0::", AdjBcastOut: "false", AdjBfdIpv4EstablishOut: "false", AdjBfdIpv6EstablishOut: "false", AdjResurrectOut: "false", AdjRestartCapableOut: "true", AdjRestartAckOut: "false", AdjRestartModeOut: "false", AdjRestartAdjSeenRaOut: "false", AdjRestartAdjSeenCsnpOut: "false", AdjRestartAdjSeenL1CsnpOut: "false", AdjRestartAdjSeenL2CsnpOut: "false", AdjRestartSuppressAdjOut: "false", TableAdjSid: struct {
					RowAdjSid []struct {
						AdjSidValue  string "json:\"adj-sid-value\""
						AdjSidFFlag  string "json:\"adj-sid-f-flag\""
						AdjSidBFlag  string "json:\"adj-sid-b-flag\""
						AdjSidVFlag  string "json:\"adj-sid-v-flag\""
						AdjSidLFlag  string "json:\"adj-sid-l-flag\""
						AdjSidSFlag  string "json:\"adj-sid-s-flag\""
						AdjSidPFlag  string "json:\"adj-sid-p-flag\""
						AdjSidWeight string "json:\"adj-sid-weight\""
					} "json:\"ROW_adj_sid\""
				}{RowAdjSid: []struct {
					AdjSidValue  string "json:\"adj-sid-value\""
					AdjSidFFlag  string "json:\"adj-sid-f-flag\""
					AdjSidBFlag  string "json:\"adj-sid-b-flag\""
					AdjSidVFlag  string "json:\"adj-sid-v-flag\""
					AdjSidLFlag  string "json:\"adj-sid-l-flag\""
					AdjSidSFlag  string "json:\"adj-sid-s-flag\""
					AdjSidPFlag  string "json:\"adj-sid-p-flag\""
					AdjSidWeight string "json:\"adj-sid-weight\""
				}{struct {
					AdjSidValue  string "json:\"adj-sid-value\""
					AdjSidFFlag  string "json:\"adj-sid-f-flag\""
					AdjSidBFlag  string "json:\"adj-sid-b-flag\""
					AdjSidVFlag  string "json:\"adj-sid-v-flag\""
					AdjSidLFlag  string "json:\"adj-sid-l-flag\""
					AdjSidSFlag  string "json:\"adj-sid-s-flag\""
					AdjSidPFlag  string "json:\"adj-sid-p-flag\""
					AdjSidWeight string "json:\"adj-sid-weight\""
				}{AdjSidValue: "16", AdjSidFFlag: "false", AdjSidBFlag: "false", AdjSidVFlag: "true", AdjSidLFlag: "true", AdjSidSFlag: "false", AdjSidPFlag: "false", AdjSidWeight: "1"}}}}, struct {
					AdjSysNameOut              string "json:\"adj-sys-name-out\""
					AdjSysIDOut                string "json:\"adj-sys-id-out\""
					AdjUsageOut                string "json:\"adj-usage-out\""
					AdjStateOut                string "json:\"adj-state-out\""
					AdjHoldTimeOut             string "json:\"adj-hold-time-out\""
					AdjIntfNameOut             string "json:\"adj-intf-name-out\""
					AdjDetailSetOut            string "json:\"adj-detail-set-out\""
					AdjTransitionsOut          string "json:\"adj-transitions-out\""
					AdjFlapOut                 string "json:\"adj-flap-out\""
					AdjFlapTimeOut             string "json:\"adj-flap-time-out\""
					AdjCktTypeOut              string "json:\"adj-ckt-type-out\""
					AdjIpv4AddrOut             string "json:\"adj-ipv4-addr-out\""
					AdjIpv6AddrOut             string "json:\"adj-ipv6-addr-out\""
					AdjBcastOut                string "json:\"adj-bcast-out\""
					AdjBfdIpv4EstablishOut     string "json:\"adj-bfd-ipv4-establish-out\""
					AdjBfdIpv6EstablishOut     string "json:\"adj-bfd-ipv6-establish-out\""
					AdjResurrectOut            string "json:\"adj-resurrect-out\""
					AdjRestartCapableOut       string "json:\"adj-restart-capable-out\""
					AdjRestartAckOut           string "json:\"adj-restart-ack-out\""
					AdjRestartModeOut          string "json:\"adj-restart-mode-out\""
					AdjRestartAdjSeenRaOut     string "json:\"adj-restart-adj-seen-ra-out\""
					AdjRestartAdjSeenCsnpOut   string "json:\"adj-restart-adj-seen-csnp-out\""
					AdjRestartAdjSeenL1CsnpOut string "json:\"adj-restart-adj-seen-l1-csnp-out\""
					AdjRestartAdjSeenL2CsnpOut string "json:\"adj-restart-adj-seen-l2-csnp-out\""
					AdjRestartSuppressAdjOut   string "json:\"adj-restart-suppress-adj-out\""
					TableAdjSid                struct {
						RowAdjSid []struct {
							AdjSidValue  string "json:\"adj-sid-value\""
							AdjSidFFlag  string "json:\"adj-sid-f-flag\""
							AdjSidBFlag  string "json:\"adj-sid-b-flag\""
							AdjSidVFlag  string "json:\"adj-sid-v-flag\""
							AdjSidLFlag  string "json:\"adj-sid-l-flag\""
							AdjSidSFlag  string "json:\"adj-sid-s-flag\""
							AdjSidPFlag  string "json:\"adj-sid-p-flag\""
							AdjSidWeight string "json:\"adj-sid-weight\""
						} "json:\"ROW_adj_sid\""
					} "json:\"TABLE_adj_sid\""
				}{AdjSysNameOut: "n9k-reg-2", AdjSysIDOut: "N/A", AdjUsageOut: "2", AdjStateOut: "UP", AdjHoldTimeOut: "00:00:28", AdjIntfNameOut: "Ethernet1/31", AdjDetailSetOut: "true", AdjTransitionsOut: "1", AdjFlapOut: "true", AdjFlapTimeOut: "01:33:30", AdjCktTypeOut: "L2", AdjIpv4AddrOut: "25.1.1.1", AdjIpv6AddrOut: "0::", AdjBcastOut: "false", AdjBfdIpv4EstablishOut: "false", AdjBfdIpv6EstablishOut: "false", AdjResurrectOut: "false", AdjRestartCapableOut: "true", AdjRestartAckOut: "false", AdjRestartModeOut: "false", AdjRestartAdjSeenRaOut: "false", AdjRestartAdjSeenCsnpOut: "false", AdjRestartAdjSeenL1CsnpOut: "false", AdjRestartAdjSeenL2CsnpOut: "false", AdjRestartSuppressAdjOut: "false", TableAdjSid: struct {
					RowAdjSid []struct {
						AdjSidValue  string "json:\"adj-sid-value\""
						AdjSidFFlag  string "json:\"adj-sid-f-flag\""
						AdjSidBFlag  string "json:\"adj-sid-b-flag\""
						AdjSidVFlag  string "json:\"adj-sid-v-flag\""
						AdjSidLFlag  string "json:\"adj-sid-l-flag\""
						AdjSidSFlag  string "json:\"adj-sid-s-flag\""
						AdjSidPFlag  string "json:\"adj-sid-p-flag\""
						AdjSidWeight string "json:\"adj-sid-weight\""
					} "json:\"ROW_adj_sid\""
				}{RowAdjSid: []struct {
					AdjSidValue  string "json:\"adj-sid-value\""
					AdjSidFFlag  string "json:\"adj-sid-f-flag\""
					AdjSidBFlag  string "json:\"adj-sid-b-flag\""
					AdjSidVFlag  string "json:\"adj-sid-v-flag\""
					AdjSidLFlag  string "json:\"adj-sid-l-flag\""
					AdjSidSFlag  string "json:\"adj-sid-s-flag\""
					AdjSidPFlag  string "json:\"adj-sid-p-flag\""
					AdjSidWeight string "json:\"adj-sid-weight\""
				}{struct {
					AdjSidValue  string "json:\"adj-sid-value\""
					AdjSidFFlag  string "json:\"adj-sid-f-flag\""
					AdjSidBFlag  string "json:\"adj-sid-b-flag\""
					AdjSidVFlag  string "json:\"adj-sid-v-flag\""
					AdjSidLFlag  string "json:\"adj-sid-l-flag\""
					AdjSidSFlag  string "json:\"adj-sid-s-flag\""
					AdjSidPFlag  string "json:\"adj-sid-p-flag\""
					AdjSidWeight string "json:\"adj-sid-weight\""
				}{AdjSidValue: "17", AdjSidFFlag: "false", AdjSidBFlag: "false", AdjSidVFlag: "true", AdjSidLFlag: "true", AdjSidSFlag: "false", AdjSidPFlag: "false", AdjSidWeight: "1"}}}}}}}}}}}}}}}}, Code: "200", Input: "show isis 2 adj det", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewIsisAdjDetailFromBytes(content)
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
