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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
	"strings"
)

// JSONRPCRequest is the payload of JSON RPC request to the API.
type JSONRPCRequest struct {
	ID      uint64                   `json:"id",xml:"id"`
	Version string                   `json:"jsonrpc",xml:"jsonrpc"`
	Method  string                   `json:"method",xml:"method"`
	Params  JSONRPCRequestParameters `json:"params",xml:"params"`
}

// JSONRPCRequestParameters are the parameters for JSONRPCRequest.
type JSONRPCRequestParameters struct {
	Command string `json:"cmd",xml:"cmd"`
	Version uint64 `json:"version",xml:"version"`
}

// NewJSONRPCRequest returns an instance of JSONRPCRequest.
func NewJSONRPCRequest(s string) []*JSONRPCRequest {
	var arr []*JSONRPCRequest
	r := &JSONRPCRequest{
		ID:      1,
		Version: "2.0",
		Method:  "cli",
		Params: JSONRPCRequestParameters{
			Command: s,
			Version: 1,
		},
	}
	arr = append(arr, r)
	return arr
}

// Client is an instance of Cisco NX-OS API client.
type Client struct {
	host     string
	port     int
	protocol string
	username string
	password string
	secure   bool
}

// NewClient returns an instance of Client.
func NewClient() *Client {
	return &Client{
		port:     443,
		protocol: "https",
	}
}

// SetHost sets the target host for the API calls.
func (cli *Client) SetHost(s string) error {
	if s == "" {
		return fmt.Errorf("empty hostname or ip address")
	}
	cli.host = s
	return nil
}

// SetPort sets the port number for the API calls.
func (cli *Client) SetPort(p int) error {
	if p == 0 {
		return fmt.Errorf("invalid port: %d", p)
	}
	cli.port = p
	return nil
}

// SetUsername sets the username for the API calls.
func (cli *Client) SetUsername(s string) error {
	if s == "" {
		return fmt.Errorf("empty username")
	}
	cli.username = s
	return nil
}

// SetPassword sets the password for the API calls.
func (cli *Client) SetPassword(s string) error {
	if s == "" {
		return fmt.Errorf("empty password")
	}
	cli.password = s
	return nil
}

// SetProtocol sets the protocol for the API calls.
func (cli *Client) SetProtocol(s string) error {
	switch s {
	case "http":
		cli.protocol = s
	case "https":
		cli.protocol = s
	default:
		return fmt.Errorf("supported protocols: http, https; unsupported protocol: %s", s)
	}
	return nil
}

// SetSecure instructs the client to enforce the validation of certificates
// and check certificate errors.
func (cli *Client) SetSecure() error {
	cli.secure = true
	return nil
}

func callAPI(url string, payload []byte, username, password string, secure bool) ([]byte, error) {
	tr := &http.Transport{}
	if !secure {
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	cli := &http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json-rpc")
	req.Header.Add("Cache-Control", "no-cache")
	req.SetBasicAuth(username, password)

	res, err := cli.Do(req)
	if err != nil {
		if !strings.HasSuffix(err.Error(), "EOF") {
			return nil, err
		}
	}
	if res == nil {
		return nil, fmt.Errorf("response: <nil>, verify url: %s", url)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		spew.Dump(err)
		if err.Error() != "EOF" {
			return nil, err
		}
	}
	return body, nil
}

// GetSystemInfo returns information about the system ("show version").
func (cli *Client) GetSystemInfo() (*SysInfo, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest("show version")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	//spew.Dump(string(jreq[:]))
	//payload := []byte(`[{ "jsonrpc": "2.0", "method": "cli", "params": { "cmd": "show version", "version": 1 }, "id": 1 }]`)
	resp, err := callAPI(url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewSysInfoFromBytes(resp)
}

// GetVlans returns vlan information ("show vlan").
func (cli *Client) GetVlans() ([]*Vlan, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest("show vlan")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	//payload := []byte(`[{ "jsonrpc": "2.0", "method": "cli", "params": { "cmd": "show vlan", "version": 1 }, "id": 1 }]`)
	resp, err := callAPI(url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewVlansFromBytes(resp)
}

// GetInterfaces returns interface information ("show interface").
func (cli *Client) GetInterfaces() ([]*Interface, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest("show interface")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	//payload := []byte(`[{ "jsonrpc": "2.0", "method": "cli", "params": { "cmd": "show interface", "version": 1 }, "id": 1 }]`)
	resp, err := callAPI(url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewInterfacesFromBytes(resp)
}

// GetGeneric returns the output of a particular command.
func (cli *Client) GetGeneric(s string) ([]byte, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest(s)
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI(url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
