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
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

// JSONRPCRequest is the payload of JSON RPC request to the API.
type JSONRPCRequest struct {
	ID      uint64                   `json:"id" xml:"id"`
	Version string                   `json:"jsonrpc" xml:"jsonrpc"`
	Method  string                   `json:"method" xml:"method"`
	Params  JSONRPCRequestParameters `json:"params" xml:"params"`
}

// JSONRPCRequestParameters are the parameters for JSONRPCRequest.
type JSONRPCRequestParameters struct {
	Command string `json:"cmd" xml:"cmd"`
	Version uint64 `json:"version" xml:"version"`
}

// JSONRPCResponseErrorData defines error message in error response
type JSONRPCResponseErrorData struct {
	Msg string `json:msg`
}

// JSONRPCResponseError defines JSON RPC error response
type JSONRPCResponseError struct {
	Code    int64                    `json:code`
	Message string                   `json:message`
	Data    JSONRPCResponseErrorData `json:data`
}

// JSONRPCResponseBody defines JSON RPC normal response body
type JSONRPCResponseBody struct {
	Body json.RawMessage `json:"body"`
}

// JSONRPCResponse is the payload of JSON RPC response to the API.
type JSONRPCResponse struct {
	Version string                `json:"jsonrpc" xml:"jsonrpc"`
	Result  json.RawMessage       `json:"result, omitempty" xml:"result"`
	Error   *JSONRPCResponseError `json:"error, omitempty" xml:"error"`
	ID      uint64                `json:"id" xml:"id"`
}

// NewJSONRPCRequest returns an instance of JSONRPCRequest.
func NewJSONRPCRequest(cmds []string) []*JSONRPCRequest {
	var arr []*JSONRPCRequest
	for id, cmd := range cmds {
		r := &JSONRPCRequest{
			ID:      uint64(id + 1),
			Version: "2.0",
			Method:  "cli",
			Params: JSONRPCRequestParameters{
				Command: cmd,
				Version: 1,
			},
		}
		arr = append(arr, r)
	}
	return arr
}

// InsAPIRequest is the payload of NX-OS API request to the API.
type InsAPIRequest struct {
	Params InsAPIRequestParameters `json:"ins_api" xml:"ins_api"`
}

// InsAPIRequestParameters are the parameters for InsAPIRequest
type InsAPIRequestParameters struct {
	Version string `json:"version" xml:"version"`
	Type    string `json:"type" xml:"type"`
	Chunk   string `json:"chunk" xml:"chunk"`
	ID      string `json:"sid" xml:"sid"`
	Input   string `json:"input" xml:"input"`
	Format  string `json:"output_format" xml:"output_format"`
}

// NewInsAPIRequest returns an instance of InsAPIRequest based on the provided
// input and request type.
func NewInsAPIRequest(s, t string) *InsAPIRequest {
	r := &InsAPIRequest{}
	r.Params.Version = "1.0"
	r.Params.Type = t
	r.Params.Chunk = "0"
	r.Params.ID = "1"
	r.Params.Input = s
	r.Params.Format = "json"
	return r
}

// NewInsAPICliShowASCIIRequest returns an instance of InsAPIRequest for
// cli_show_ascii type of NX-OS API request.
func NewInsAPICliShowASCIIRequest(s string) *InsAPIRequest {
	return NewInsAPIRequest(s, "cli_show_ascii")
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

func callAPI(contentType string, url string, payload []byte, username, password string, secure bool) ([]byte, error) {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	if !secure {
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	cli := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}
	var reqContentType string
	switch contentType {
	case "jsonrpc":
		reqContentType = "application/json-rpc"
	case "json":
		reqContentType = "application/json"
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", reqContentType)
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
		if err.Error() != "EOF" {
			return nil, err
		}
	}
	if len(body) < 500 {
		if bytes.Contains(body, []byte("401 Authorization Required")) {
			return nil, fmt.Errorf("401 Authorization Required")
		}
		if bytes.Contains(body, []byte("405 Not Allowed")) {
			return nil, fmt.Errorf("405 Not Allowed")
		}
		if bytes.Contains(body, []byte("Server internal error")) {
			return nil, fmt.Errorf("500 Server Internal Error")
		}
	}
	return body, nil
}

// GetSystemInfo returns information about the system ("show version").
func (cli *Client) GetSystemInfo() (*SysInfo, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest([]string{"show version"})
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewSysInfoFromBytes(resp)
}

// GetVlans returns vlan information ("show vlan").
func (cli *Client) GetVlans() ([]*Vlan, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest([]string{"show vlan"})
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewVlansFromBytes(resp)
}

// GetInterfaces returns interface information ("show interface").
func (cli *Client) GetInterfaces() ([]*Interface, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest([]string{"show interface"})
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewInterfacesFromBytes(resp)
}

// GetInterface returns interface information ("show interface <name>").
func (cli *Client) GetInterface(name string) (*Interface, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest([]string{"show interface " + name})
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}

	intf, err := NewInterfaceFromBytes(resp)

	return intf, err
}

// GetSystemResources returns SystemResources instance ("show system resources").
func (cli *Client) GetSystemResources() (*SystemResources, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest([]string{"show system resources"})
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewSystemResourcesFromBytes(resp)
}

// GetSystemEnvironment returns SystemEnvironment instance ("show environment").
func (cli *Client) GetSystemEnvironment() (*SystemEnvironment, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest([]string{"show environment"})
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewSystemEnvironmentFromBytes(resp)
}

// GetGeneric returns the output of a particular command.
func (cli *Client) GetGeneric(s string) ([]byte, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest([]string{s})
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBgpSummary returns BgpSummary instance ("show ip bgp summary vrf all").
func (cli *Client) GetBgpSummary() (*BgpSummary, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewInsAPICliShowASCIIRequest("show ip bgp summary vrf all")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("json", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewBgpSummaryFromBytes(resp)
}

// GetRunningConfiguration returns Configuration instance for running
// configuration ("show running-config").
func (cli *Client) GetRunningConfiguration() (*Configuration, error) {
	return cli.getConfiguration("running")
}

// GetStartupConfiguration returns Configuration instance for startup
// configuration ("show startup-config").
func (cli *Client) GetStartupConfiguration() (*Configuration, error) {
	return cli.getConfiguration("startup")
}

func (cli *Client) getConfiguration(s string) (*Configuration, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewInsAPICliShowASCIIRequest("show " + s + "-config")
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("json", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewConfigurationFromBytes(resp)
}

// GetTransceivers returns data about transceivers attached to Interface
// ("show interface transceiver details").
func (cli *Client) GetTransceivers() ([]*Transceiver, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)
	req := NewJSONRPCRequest([]string{"show interface transceiver details"})
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}
	return NewTransceiversFromBytes(resp)
}

// Configure execute a batch of configuration commands
func (cli *Client) Configure(cmds []string) ([]JSONRPCResponse, error) {
	url := fmt.Sprintf("%s://%s:%d/ins", cli.protocol, cli.host, cli.port)

	req := NewJSONRPCRequest(cmds)
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := callAPI("jsonrpc", url, payload, cli.username, cli.password, cli.secure)
	if err != nil {
		return nil, err
	}

	var respJSON []JSONRPCResponse
	err = json.Unmarshal(resp, &respJSON)

	return respJSON, err
}
