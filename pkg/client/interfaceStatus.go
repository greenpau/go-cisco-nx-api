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

type InterfaceStatusResponse struct {
	InsAPI struct {
		Outputs struct {
			Output InterfaceStatusResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type InterfaceStatusResponseResult struct {
	Body  InterfaceStatusResultBody `json:"body" xml:"body"`
	Code  string                    `json:"code" xml:"code"`
	Input string                    `json:"input" xml:"input"`
	Msg   string                    `json:"msg" xml:"msg"`
}

type InterfaceStatusResultBody struct {
	TableInterface []struct {
		RowInterface []struct {
			Interface string `json:"interface" xml:"interface"`
			State     string `json:"state" xml:"state"`
			Vlan      string `json:"vlan" xml:"vlan"`
			Duplex    string `json:"duplex" xml:"duplex"`
			Speed     string `json:"speed" xml:"speed"`
			Type      string `json:"type,omitempty" xml:"type,omitempty"`
		} `json:"ROW_interface" xml:"ROW_interface"`
	} `json:"TABLE_interface" xml:"TABLE_interface"`
}

func (d *InterfaceStatusResponse) Flat() (out []InterfaceStatusResultFlat) {
	return d.InsAPI.Outputs.Output.Body.Flat()
}
func (d *InterfaceStatusResultBody) Flat() (out []InterfaceStatusResultFlat) {
	for _, Ti := range d.TableInterface {
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
	Interface string `json:"interface" xml:"interface"`
	State     string `json:"state" xml:"state"`
	Vlan      string `json:"vlan" xml:"vlan"`
	Duplex    string `json:"duplex" xml:"duplex"`
	Speed     string `json:"speed" xml:"speed"`
	Type      string `json:"type,omitempty" xml:"type,omitempty"`
}

// NewInterfaceStatusFromString returns instance from an input string.
func NewInterfaceStatusFromString(s string) (*InterfaceStatusResponse, error) {
	return NewInterfaceStatusFromReader(strings.NewReader(s))
}

// NewInterfaceStatusFromBytes returns instance from an input byte array.
func NewInterfaceStatusFromBytes(s []byte) (*InterfaceStatusResponse, error) {
	return NewInterfaceStatusFromReader(bytes.NewReader(s))
}

// NewInterfaceStatusFromReader returns instance from an input reader.
func NewInterfaceStatusFromReader(s io.Reader) (*InterfaceStatusResponse, error) {
	//si := &InterfaceStatus{}
	InterfaceStatusResponseDat := &InterfaceStatusResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(InterfaceStatusResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return InterfaceStatusResponseDat, nil
}

// NewInterfaceStatusResultFromString returns instance from an input string.
func NewInterfaceStatusResultFromString(s string) (*InterfaceStatusResponseResult, error) {
	return NewInterfaceStatusResultFromReader(strings.NewReader(s))
}

// NewInterfaceStatusResultFromBytes returns instance from an input byte array.
func NewInterfaceStatusResultFromBytes(s []byte) (*InterfaceStatusResponseResult, error) {
	return NewInterfaceStatusResultFromReader(bytes.NewReader(s))
}

// NewInterfaceStatusResultFromReader returns instance from an input reader.
func NewInterfaceStatusResultFromReader(s io.Reader) (*InterfaceStatusResponseResult, error) {
	//si := &InterfaceStatusResponseResult{}
	InterfaceStatusResponseResultDat := &InterfaceStatusResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(InterfaceStatusResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return InterfaceStatusResponseResultDat, nil
}
