// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//            and Paul Schou     (github.com/pschou)
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
	"fmt"
	"github.com/pschou/go-json"
	"io"
	"strings"
)

type InterfaceBriefResponse struct {
	TableInterface struct {
		RowInterface []struct {
			Interface    string `json:"interface"`
			State        string `json:"state,omitempty"`
			IPAddr       string `json:"ip_addr,omitempty"`
			Speed        string `json:"speed,omitempty"`
			MTU          int    `json:"mtu,omitempty"`
			Vlan         string `json:"vlan,omitempty"`
			Type         string `json:"type,omitempty"`
			PortMode     string `json:"portmode,omitempty"`
			StateRsnDesc string `json:"state_rsn_desc,omitempty"`
			RateMode     string `json:"ratemode,omitempty"`
			PortChan     int    `json:"portchan,omitempty"`
			Proto        string `json:"proto,omitempty"`
			OperState    string `json:"oper_state,omitempty"`
			SviRsnDesc   string `json:"svi_rsn_desc,omitempty"`
		} `json:"ROW_interface"`
	} `json:"TABLE_interface"`
}

// NewInterfaceBriefFromString returns instance from an input string.
func NewInterfaceBriefFromString(s string) (*InterfaceBriefResponse, error) {
	return NewInterfaceBriefFromReader(strings.NewReader(s))
}

// NewInterfaceBriefFromBytes returns instance from an input byte array.
func NewInterfaceBriefFromBytes(s []byte) (*InterfaceBriefResponse, error) {
	return NewInterfaceBriefFromReader(bytes.NewReader(s))
}

// NewInterfaceBriefFromReader returns instance from an input reader.
func NewInterfaceBriefFromReader(s io.Reader) (*InterfaceBriefResponse, error) {
	//si := &InterfaceBrief{}
	InterfaceBriefResponseDat := &InterfaceBriefResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(InterfaceBriefResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return InterfaceBriefResponseDat, nil
}
