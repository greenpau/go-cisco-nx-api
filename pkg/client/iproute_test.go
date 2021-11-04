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

func TestParseShowIPRouteJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *IpRouteResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.ip.route",
			exp: &IpRouteResponse{InsAPI: struct {
				Outputs struct {
					Output IpRouteResponseResult "json:\"output\""
				} "json:\"outputs\""
				Sid     string "json:\"sid\""
				Type    string "json:\"type\""
				Version string "json:\"version\""
			}{Outputs: struct {
				Output IpRouteResponseResult "json:\"output\""
			}{Output: IpRouteResponseResult{Body: IpRouteResultBody{TableVrf: []struct {
				RowVrf []struct {
					TableAddrf []struct {
						RowAddrf []struct {
							TablePrefix []struct {
								RowPrefix []struct {
									TablePath []struct {
										RowPath []struct {
											Clientname string "json:\"clientname\""
											Ifname     string "json:\"ifname\""
											Metric     int    "json:\"metric\""
											Pref       int    "json:\"pref\""
											UBest      string "json:\"ubest\""
											UpTime     string "json:\"uptime\""
										} "json:\"ROW_path\""
									} "json:\"TABLE_path\""
									Attached   string "json:\"attached\""
									IPPrefix   string "json:\"ipprefix\""
									McastNhops int    "json:\"mcast-nhops\""
									UcastNhops int    "json:\"ucast-nhops\""
								} "json:\"ROW_prefix\""
							} "json:\"TABLE_prefix\""
							AddRf string "json:\"addrf\""
						} "json:\"ROW_addrf\""
					} "json:\"TABLE_addrf\""
					VrfNameOut string "json:\"vrf-name-out\""
				} "json:\"ROW_vrf\""
			}{struct {
				RowVrf []struct {
					TableAddrf []struct {
						RowAddrf []struct {
							TablePrefix []struct {
								RowPrefix []struct {
									TablePath []struct {
										RowPath []struct {
											Clientname string "json:\"clientname\""
											Ifname     string "json:\"ifname\""
											Metric     int    "json:\"metric\""
											Pref       int    "json:\"pref\""
											UBest      string "json:\"ubest\""
											UpTime     string "json:\"uptime\""
										} "json:\"ROW_path\""
									} "json:\"TABLE_path\""
									Attached   string "json:\"attached\""
									IPPrefix   string "json:\"ipprefix\""
									McastNhops int    "json:\"mcast-nhops\""
									UcastNhops int    "json:\"ucast-nhops\""
								} "json:\"ROW_prefix\""
							} "json:\"TABLE_prefix\""
							AddRf string "json:\"addrf\""
						} "json:\"ROW_addrf\""
					} "json:\"TABLE_addrf\""
					VrfNameOut string "json:\"vrf-name-out\""
				} "json:\"ROW_vrf\""
			}{RowVrf: []struct {
				TableAddrf []struct {
					RowAddrf []struct {
						TablePrefix []struct {
							RowPrefix []struct {
								TablePath []struct {
									RowPath []struct {
										Clientname string "json:\"clientname\""
										Ifname     string "json:\"ifname\""
										Metric     int    "json:\"metric\""
										Pref       int    "json:\"pref\""
										UBest      string "json:\"ubest\""
										UpTime     string "json:\"uptime\""
									} "json:\"ROW_path\""
								} "json:\"TABLE_path\""
								Attached   string "json:\"attached\""
								IPPrefix   string "json:\"ipprefix\""
								McastNhops int    "json:\"mcast-nhops\""
								UcastNhops int    "json:\"ucast-nhops\""
							} "json:\"ROW_prefix\""
						} "json:\"TABLE_prefix\""
						AddRf string "json:\"addrf\""
					} "json:\"ROW_addrf\""
				} "json:\"TABLE_addrf\""
				VrfNameOut string "json:\"vrf-name-out\""
			}{struct {
				TableAddrf []struct {
					RowAddrf []struct {
						TablePrefix []struct {
							RowPrefix []struct {
								TablePath []struct {
									RowPath []struct {
										Clientname string "json:\"clientname\""
										Ifname     string "json:\"ifname\""
										Metric     int    "json:\"metric\""
										Pref       int    "json:\"pref\""
										UBest      string "json:\"ubest\""
										UpTime     string "json:\"uptime\""
									} "json:\"ROW_path\""
								} "json:\"TABLE_path\""
								Attached   string "json:\"attached\""
								IPPrefix   string "json:\"ipprefix\""
								McastNhops int    "json:\"mcast-nhops\""
								UcastNhops int    "json:\"ucast-nhops\""
							} "json:\"ROW_prefix\""
						} "json:\"TABLE_prefix\""
						AddRf string "json:\"addrf\""
					} "json:\"ROW_addrf\""
				} "json:\"TABLE_addrf\""
				VrfNameOut string "json:\"vrf-name-out\""
			}{TableAddrf: []struct {
				RowAddrf []struct {
					TablePrefix []struct {
						RowPrefix []struct {
							TablePath []struct {
								RowPath []struct {
									Clientname string "json:\"clientname\""
									Ifname     string "json:\"ifname\""
									Metric     int    "json:\"metric\""
									Pref       int    "json:\"pref\""
									UBest      string "json:\"ubest\""
									UpTime     string "json:\"uptime\""
								} "json:\"ROW_path\""
							} "json:\"TABLE_path\""
							Attached   string "json:\"attached\""
							IPPrefix   string "json:\"ipprefix\""
							McastNhops int    "json:\"mcast-nhops\""
							UcastNhops int    "json:\"ucast-nhops\""
						} "json:\"ROW_prefix\""
					} "json:\"TABLE_prefix\""
					AddRf string "json:\"addrf\""
				} "json:\"ROW_addrf\""
			}{struct {
				RowAddrf []struct {
					TablePrefix []struct {
						RowPrefix []struct {
							TablePath []struct {
								RowPath []struct {
									Clientname string "json:\"clientname\""
									Ifname     string "json:\"ifname\""
									Metric     int    "json:\"metric\""
									Pref       int    "json:\"pref\""
									UBest      string "json:\"ubest\""
									UpTime     string "json:\"uptime\""
								} "json:\"ROW_path\""
							} "json:\"TABLE_path\""
							Attached   string "json:\"attached\""
							IPPrefix   string "json:\"ipprefix\""
							McastNhops int    "json:\"mcast-nhops\""
							UcastNhops int    "json:\"ucast-nhops\""
						} "json:\"ROW_prefix\""
					} "json:\"TABLE_prefix\""
					AddRf string "json:\"addrf\""
				} "json:\"ROW_addrf\""
			}{RowAddrf: []struct {
				TablePrefix []struct {
					RowPrefix []struct {
						TablePath []struct {
							RowPath []struct {
								Clientname string "json:\"clientname\""
								Ifname     string "json:\"ifname\""
								Metric     int    "json:\"metric\""
								Pref       int    "json:\"pref\""
								UBest      string "json:\"ubest\""
								UpTime     string "json:\"uptime\""
							} "json:\"ROW_path\""
						} "json:\"TABLE_path\""
						Attached   string "json:\"attached\""
						IPPrefix   string "json:\"ipprefix\""
						McastNhops int    "json:\"mcast-nhops\""
						UcastNhops int    "json:\"ucast-nhops\""
					} "json:\"ROW_prefix\""
				} "json:\"TABLE_prefix\""
				AddRf string "json:\"addrf\""
			}{struct {
				TablePrefix []struct {
					RowPrefix []struct {
						TablePath []struct {
							RowPath []struct {
								Clientname string "json:\"clientname\""
								Ifname     string "json:\"ifname\""
								Metric     int    "json:\"metric\""
								Pref       int    "json:\"pref\""
								UBest      string "json:\"ubest\""
								UpTime     string "json:\"uptime\""
							} "json:\"ROW_path\""
						} "json:\"TABLE_path\""
						Attached   string "json:\"attached\""
						IPPrefix   string "json:\"ipprefix\""
						McastNhops int    "json:\"mcast-nhops\""
						UcastNhops int    "json:\"ucast-nhops\""
					} "json:\"ROW_prefix\""
				} "json:\"TABLE_prefix\""
				AddRf string "json:\"addrf\""
			}{TablePrefix: []struct {
				RowPrefix []struct {
					TablePath []struct {
						RowPath []struct {
							Clientname string "json:\"clientname\""
							Ifname     string "json:\"ifname\""
							Metric     int    "json:\"metric\""
							Pref       int    "json:\"pref\""
							UBest      string "json:\"ubest\""
							UpTime     string "json:\"uptime\""
						} "json:\"ROW_path\""
					} "json:\"TABLE_path\""
					Attached   string "json:\"attached\""
					IPPrefix   string "json:\"ipprefix\""
					McastNhops int    "json:\"mcast-nhops\""
					UcastNhops int    "json:\"ucast-nhops\""
				} "json:\"ROW_prefix\""
			}{struct {
				RowPrefix []struct {
					TablePath []struct {
						RowPath []struct {
							Clientname string "json:\"clientname\""
							Ifname     string "json:\"ifname\""
							Metric     int    "json:\"metric\""
							Pref       int    "json:\"pref\""
							UBest      string "json:\"ubest\""
							UpTime     string "json:\"uptime\""
						} "json:\"ROW_path\""
					} "json:\"TABLE_path\""
					Attached   string "json:\"attached\""
					IPPrefix   string "json:\"ipprefix\""
					McastNhops int    "json:\"mcast-nhops\""
					UcastNhops int    "json:\"ucast-nhops\""
				} "json:\"ROW_prefix\""
			}{RowPrefix: []struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "static", Ifname: "Null0", Metric: 0, Pref: 1, UBest: "true", UpTime: "P7DT12H2M5S"}}}}, Attached: "false", IPPrefix: "7.57.0.0/16", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Vlan253", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H34S"}}}}, Attached: "true", IPPrefix: "7.57.253.0/30", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Vlan253", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H34S"}}}}, Attached: "true", IPPrefix: "7.57.253.2/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "bgp-65057", Ifname: "", Metric: 0, Pref: 20, UBest: "true", UpTime: "P7DT11H58M55S"}, struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "bgp-65057", Ifname: "", Metric: 0, Pref: 20, UBest: "true", UpTime: "P7DT11H58M55S"}}}}, Attached: "false", IPPrefix: "7.57.255.1/32", McastNhops: 0, UcastNhops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Lo0", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H2M2S"}, struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Lo0", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H2M2S"}}}}, Attached: "true", IPPrefix: "7.57.255.2/32", McastNhops: 0, UcastNhops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "ospf-1", Ifname: "Vlan253", Metric: 42, Pref: 110, UBest: "true", UpTime: "P7DT12H22S"}}}}, Attached: "false", IPPrefix: "8.8.8.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "ospf-1", Ifname: "Vlan253", Metric: 42, Pref: 110, UBest: "true", UpTime: "P7DT12H22S"}}}}, Attached: "false", IPPrefix: "9.9.9.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "static", Ifname: "Null0", Metric: 0, Pref: 1, UBest: "true", UpTime: "P7DT12H2M5S"}}}}, Attached: "false", IPPrefix: "10.0.0.0/8", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "static", Ifname: "Null0", Metric: 0, Pref: 1, UBest: "true", UpTime: "P7DT12H2M5S"}}}}, Attached: "false", IPPrefix: "10.57.0.0/16", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Vlan8", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H44S"}}}}, Attached: "true", IPPrefix: "10.57.8.0/22", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Vlan8", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H44S"}}}}, Attached: "true", IPPrefix: "10.57.8.3/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Vlan50", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H34S"}}}}, Attached: "true", IPPrefix: "10.57.50.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Vlan50", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H34S"}}}}, Attached: "true", IPPrefix: "10.57.50.4/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/1", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "10.100.90.128/30", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/1", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "10.100.90.129/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/2", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H48S"}}}}, Attached: "true", IPPrefix: "10.100.157.0/30", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/2", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H48S"}}}}, Attached: "true", IPPrefix: "10.100.157.2/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/4", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H48S"}}}}, Attached: "true", IPPrefix: "10.100.157.8/30", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/4", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H48S"}}}}, Attached: "true", IPPrefix: "10.100.157.10/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "static", Ifname: "Eth1/5", Metric: 0, Pref: 1, UBest: "true", UpTime: "P7DT12H1M37S"}}}}, Attached: "false", IPPrefix: "17.0.1.1/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "static", Ifname: "Eth1/5", Metric: 0, Pref: 1, UBest: "true", UpTime: "P7DT12H1M37S"}}}}, Attached: "false", IPPrefix: "17.0.1.2/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "static", Ifname: "Eth1/5", Metric: 0, Pref: 1, UBest: "true", UpTime: "P7DT12H1M37S"}}}}, Attached: "false", IPPrefix: "17.0.1.3/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "bgp-65057", Ifname: "", Metric: 0, Pref: 20, UBest: "true", UpTime: "P7DT11H58M55S"}, struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "bgp-65057", Ifname: "", Metric: 0, Pref: 20, UBest: "true", UpTime: "P7DT11H58M55S"}}}}, Attached: "false", IPPrefix: "33.33.33.33/32", McastNhops: 0, UcastNhops: 2}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/1.50", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "89.1.1.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/1.50", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "89.1.1.1/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/1.51", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "89.1.2.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/1.51", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "89.1.2.1/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/1.52", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "89.1.3.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/1.52", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "89.1.3.1/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/1.53", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "89.1.4.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/1.53", Metric: 0, Pref: 0, UBest: "true", UpTime: "P6DT12H2M16S"}}}}, Attached: "true", IPPrefix: "89.1.4.1/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/5", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H1M37S"}}}}, Attached: "true", IPPrefix: "94.1.1.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/5", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H1M37S"}}}}, Attached: "true", IPPrefix: "94.1.1.96/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "direct", Ifname: "Eth1/7", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H45S"}}}}, Attached: "true", IPPrefix: "192.168.161.0/24", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{Clientname: "local", Ifname: "Eth1/7", Metric: 0, Pref: 0, UBest: "true", UpTime: "P7DT12H45S"}}}}, Attached: "true", IPPrefix: "192.168.161.2/32", McastNhops: 0, UcastNhops: 1}, struct {
				TablePath []struct {
					RowPath []struct {
						Clientname string "json:\"clientname\""
						Ifname     string "json:\"ifname\""
						Metric     int    "json:\"metric\""
						Pref       int    "json:\"pref\""
						UBest      string "json:\"ubest\""
						UpTime     string "json:\"uptime\""
					} "json:\"ROW_path\""
				} "json:\"TABLE_path\""
				Attached   string "json:\"attached\""
				IPPrefix   string "json:\"ipprefix\""
				McastNhops int    "json:\"mcast-nhops\""
				UcastNhops int    "json:\"ucast-nhops\""
			}{TablePath: []struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{struct {
				RowPath []struct {
					Clientname string "json:\"clientname\""
					Ifname     string "json:\"ifname\""
					Metric     int    "json:\"metric\""
					Pref       int    "json:\"pref\""
					UBest      string "json:\"ubest\""
					UpTime     string "json:\"uptime\""
				} "json:\"ROW_path\""
			}{RowPath: []struct {
				Clientname string "json:\"clientname\""
				Ifname     string "json:\"ifname\""
				Metric     int    "json:\"metric\""
				Pref       int    "json:\"pref\""
				UBest      string "json:\"ubest\""
				UpTime     string "json:\"uptime\""
			}{}}}, Attached: "true", IPPrefix: "192.168.161.2/32", McastNhops: 0, UcastNhops: 1}}}}, AddRf: "ipv4"}}}}, VrfNameOut: "default"}}}}}, Code: "200", Input: "show ip route ", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		dat, err := NewIpRouteFromBytes(content)
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
