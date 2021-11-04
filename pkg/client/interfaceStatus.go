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
	"fmt"
	"github.com/pschou/go-json"
	"io"
	"strings"
)

type InterfaceStatusResponse struct {
	InsAPI struct {
		Outputs struct {
			Output InterfaceStatusResponseResult `json:"output"`
		} `json:"outputs"`
		Sid     string `json:"sid"`
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"ins_api"`
}

type InterfaceStatusResponseResult struct {
	Body  InterfaceStatusResultBody `json:"body"`
	Code  string                    `json:"code"`
	Input string                    `json:"input"`
	Msg   string                    `json:"msg"`
}

type InterfaceStatusResultBody struct {
	TableInterface []struct {
		RowInterface []struct {
			Interface string `json:"interface"`
			State     string `json:"state"`
			Vlan      string `json:"vlan"`
			Duplex    string `json:"duplex"`
			Speed     string `json:"speed"`
			Type      string `json:"type,omitempty"`
		} `json:"ROW_interface"`
	} `json:"TABLE_interface"`
}

func (d *InterfaceStatusResponse) Flat() (out []InterfaceStatusResultFlat) {
	for _, Ti := range d.InsAPI.Outputs.Output.Body.TableInterface {
		for _, Ri := range Ti.RowInterface {
			out = append(out, InterfaceStatusResultFlat{
				Interface: Ri.Interface,
				State:     Ri.State,
				Vlan:      Ri.Vlan,
				Duplex:    Ri.Duplex,
				Speed:     Ri.Speed,
				Type:      Ri.Type,
			})
		}
	}
	return
}

type InterfaceStatusResultFlat struct {
	Interface string `json:"interface"`
	State     string `json:"state"`
	Vlan      string `json:"vlan"`
	Duplex    string `json:"duplex"`
	Speed     string `json:"speed"`
	Type      string `json:"type,omitempty"`
}

// NewInterfaceStatusFromString returns SysInfo instance from an input string.
func NewInterfaceStatusFromString(s string) (*InterfaceStatusResponse, error) {
	return NewInterfaceStatusFromReader(strings.NewReader(s))
}

// NewSysInfoFromBytes returns SysInfo instance from an input byte array.
func NewInterfaceStatusFromBytes(s []byte) (*InterfaceStatusResponse, error) {
	return NewInterfaceStatusFromReader(bytes.NewReader(s))
}
func NewInterfaceStatusFromReader(s io.Reader) (*InterfaceStatusResponse, error) {
	//si := &InterfaceStatus{}
	InterfaceStatusResponseDat := &InterfaceStatusResponse{}
	jsonDec := json.NewDecoder(s)
	//jsonDec.UseNumber()
	jsonDec.UseSlice()
	err := jsonDec.Decode(InterfaceStatusResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return InterfaceStatusResponseDat, nil
}
