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
)

// Configuration contains device configuration. The information in the
// structure is from the output of "show running-config" or
// "show startup-config" commands.
type Configuration struct {
	Text string `json:"text" xml:"text"`
}

// NewConfigurationFromString returns Configuration instance from an input string.
func NewConfigurationFromString(s string) (*Configuration, error) {
	return NewConfigurationFromBytes([]byte(s))
}

// NewConfigurationFromBytes returns Configuration instance from an input byte array.
func NewConfigurationFromBytes(s []byte) (*Configuration, error) {
	c := &Configuration{}
	resp := &insAPIResponse{}
	err := json.Unmarshal(s, resp)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	if resp.Result.Outputs.Output.Code != "200" {
		return nil, fmt.Errorf("error: %s, %s, server response: %s",
			resp.Result.Outputs.Output.Code, resp.Result.Outputs.Output.Message, string(s[:]))
	}
	c.Text = resp.Result.Outputs.Output.Body
	return c, nil
}
