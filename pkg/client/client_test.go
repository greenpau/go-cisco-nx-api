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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGetSystemInfo(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ins", func(w http.ResponseWriter, req *http.Request) {
		var err error
		var fp string
		var fc []byte
		dataDir := "../../assets/requests"
		if req.Method != "POST" {
			http.Error(w, "Bad Request, expecting POST", http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad Request, ioutil.ReadAll: %s", err), http.StatusBadRequest)
			return
		}
		var j []*JSONRPCRequest
		err = json.Unmarshal(body, &j)
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad Request, json.Unmarshal: %s", err), http.StatusBadRequest)
			return
		}
		if len(j) != 1 {
			http.Error(w, fmt.Sprintf("Bad Request, expecting a single query, got %d", len(j)), http.StatusBadRequest)
		}
		cmd := j[0].Params.Command
		t.Logf("server: received command: %s", cmd)
		switch cmd {
		case "show version":
			fp = fmt.Sprintf("%s/%s", dataDir, "resp.show.version.1.json")
		case "show vlan":
			fp = fmt.Sprintf("%s/%s", dataDir, "resp.show.vlans.2.json")
		case "show interface":
			fp = fmt.Sprintf("%s/%s", dataDir, "resp.show.interfaces.4.json")
		default:
			http.Error(w, fmt.Sprintf("Bad Request, unsupported command: %s", cmd), http.StatusBadRequest)
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
}
