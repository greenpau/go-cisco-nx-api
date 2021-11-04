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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ins", func(w http.ResponseWriter, req *http.Request) {
		var err error
		var fp string
		var fc []byte
		dataDir := "../../assets/requests"
		showCmdFileMap := map[string]string{
			"show version":                       "resp.show.version.1.json",
			"show vlan":                          "resp.show.vlans.2.json",
			"show interface":                     "resp.show.interfaces.4.json",
			"show system resources":              "resp.show.system.resources.1.json",
			"show environment":                   "resp.show.environment.1.json",
			"show running-config":                "resp.show.running.config.1.json",
			"show ip bgp summary vrf all":        "resp.show.ip.bgp.summary.vrf.all.1.json",
			"show interface transceiver details": "resp.show.interface.transceiver.details.1.json",
			"show clock":                         "resp.show.clock.json",
		}
		if req.Method != "POST" {
			http.Error(w, "Bad Request, expecting POST", http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad Request, ioutil.ReadAll: %s", err), http.StatusBadRequest)
			return
		}

		var cmd string

		if bytes.Contains(body, []byte("jsonrpc")) {
			var j []*JSONRPCRequest
			err = json.Unmarshal(body, &j)
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad Request, json.Unmarshal: %s", err), http.StatusBadRequest)
				return
			}
			if len(j) == 0 {
				http.Error(w, fmt.Sprintf("Bad Request, empty input"), http.StatusBadRequest)
				return
			}
			cmd = j[0].Params.Command
			if strings.HasPrefix(cmd, "show") {
				if len(j) != 1 {
					http.Error(w, fmt.Sprintf("Bad Request, expecting a single query, got %d", len(j)), http.StatusBadRequest)
					return
				}
				t.Logf("server: received command: %s", cmd)
				respFileName, isCmdSupported := showCmdFileMap[cmd]
				if !isCmdSupported {
					http.Error(w, fmt.Sprintf("Bad Request, unsupported command: %s", cmd), http.StatusBadRequest)
					return
				}

				fp = fmt.Sprintf("%s/%s", dataDir, respFileName)
			} else {
				// interface config commands
				var cmds []string
				for i := range j {
					cmds = append(cmds, j[i].Params.Command)
				}

				t.Logf("server: received commands: %s", cmds)
				if len(j) == 1 {
					if !strings.HasPrefix(cmds[0], "vlan") {
						http.Error(w, fmt.Sprintf("Wrong config command %s", cmds[0]), http.StatusBadRequest)
						return
					}
					fp = fmt.Sprintf("%s/%s", dataDir, "resp.vlan.json")
				} else {
					if !strings.HasPrefix(cmds[0], "interface") {
						http.Error(w, fmt.Sprintf("Wrong config command %s", cmds[0]), http.StatusBadRequest)
						return
					}
					fp = fmt.Sprintf("%s/%s", dataDir, "resp.shutdown.interface.json")
				}
			}
		} else if bytes.Contains(body, []byte(`"ins_api":`)) {
			var j *InsAPIRequest
			err = json.Unmarshal(body, &j)
			if err != nil {
				http.Error(w, fmt.Sprintf("Bad Request, json.Unmarshal: %s, Body: %s", err, body), http.StatusBadRequest)
				return
			}
			cmd = j.Params.Input
			t.Logf("server: received command: %s", cmd)
			respFileName, isCmdSupported := showCmdFileMap[cmd]
			if !isCmdSupported {
				http.Error(w, fmt.Sprintf("Bad Request, unsupported command: %s", cmd), http.StatusBadRequest)
				return
			}

			fp = fmt.Sprintf("%s/%s", dataDir, respFileName)
		} else {
			http.Error(w, fmt.Sprintf("Bad Request, unsupported payload %s", string(body[:])), http.StatusBadRequest)
			return
		}

		fc, err = ioutil.ReadFile(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(fc)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	srv := strings.Split(server.URL, ":")
	proto := srv[0]
	port, _ := strconv.Atoi(srv[2])

	cli := NewClient()
	cli.SetHost("127.0.0.1")
	cli.SetPort(port)
	cli.SetProtocol(proto)
	cli.SetUsername("admin")
	cli.SetPassword("cisco")

	start := time.Now()
	sysinfo, err := cli.GetSystemInfo()
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	t.Logf("client: Hostname: %s", sysinfo.Hostname)
	t.Logf("client: Processor Board ID: %s", sysinfo.ProcessorBoardID)
	t.Logf("client: Kickstart Image Version: %s", sysinfo.KickstartImage.Version)
	t.Logf("client: Uptime: %d", sysinfo.Uptime)
	t.Logf("client: took %s", time.Since(start))

	start = time.Now()
	ifaces, err := cli.GetInterfaces()
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	t.Logf("client: Interfaces: %d", len(ifaces))
	t.Logf("client: took %s", time.Since(start))

	start = time.Now()
	vlans, err := cli.GetVlans()
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	t.Logf("client: VLANs: %d", len(vlans))
	t.Logf("client: took %s", time.Since(start))

	start = time.Now()
	resources, err := cli.GetSystemResources()
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	t.Logf("client: CPUs: %d", len(resources.CPUs))
	t.Logf("client: Processes: %d", resources.Processes.Total)
	t.Logf("client: took %s", time.Since(start))

	start = time.Now()
	environment, err := cli.GetSystemEnvironment()
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	t.Logf("client: Fans: %d", len(environment.Fans))
	t.Logf("client: Power Supplies: %d", len(environment.PowerSupplies))
	t.Logf("client: Sensors: %d", len(environment.Sensors))
	t.Logf("client: took %s", time.Since(start))

	start = time.Now()
	conf, err := cli.GetRunningConfiguration()
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	t.Logf("client: Config output size (bytes): %d", len(conf.Text))
	t.Logf("client: took %s", time.Since(start))

	start = time.Now()
	bgp, err := cli.GetBgpSummary()
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	t.Logf("client: BGP summary output size (bytes): %d", len(bgp.Text))
	t.Logf("client: took %s", time.Since(start))

	start = time.Now()
	transceivers, err := cli.GetTransceivers()
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	t.Logf("client: Transceivers: %d", len(transceivers))

	output, err := cli.GetGeneric("show clock")
	if err != nil {
		t.Fatalf("client: %s", err)
	}

	var respJSON JSONRPCResponse
	err = json.Unmarshal(output, &respJSON)
	if err != nil {
		t.Fatalf("JSON parsing failed:%v\n Input: %s", err, string(output))
	}

	if respJSON.Error != nil {
		t.Fatalf("Command returned failure: %v", respJSON.Error)
	}

	var body JSONRPCResponseBody
	err = json.Unmarshal(respJSON.Result, &body)
	if err != nil {
		t.Fatalf("JSON parsing failed:%v\n Input: %s", err, string(respJSON.Result))
	}

	if !strings.Contains(string(body.Body), "simple_time") {
		t.Fatalf("client: returned unknown response from show clock")
	}

	resp, err := cli.Configure([]string{"interface e1/1", "shutdown"})
	if err != nil {
		t.Fatalf("client: %s", err)
	}

	for _, r := range resp {
		if r.Error != nil {
			t.Fatalf("failed to execute command %v:\n%v\n", r.ID, r.Error)
		}
	}

	resp, err = cli.Configure([]string{"vlan 1-2"})
	if err != nil {
		t.Fatalf("client: %s", err)
	}

	for _, r := range resp {
		if r.Error != nil {
			t.Fatalf("failed to execute command %v:\n%v\n", r.ID, r.Error)
		}
	}

	t.Logf("client: took %s", time.Since(start))
}
