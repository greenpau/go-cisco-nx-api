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

type insAPIResponse struct {
	Result insAPIResponseResult `json:"ins_api" xml:"ins_api"`
}

type insAPIResponseResult struct {
	Type    string                      `json:"type" xml:"type"`
	Version string                      `json:"version" xml:"version"`
	Outputs insAPIResponseResultOutputs `json:"outputs" xml:"outputs"`
}

type insAPIResponseResultOutputs struct {
	Output insAPIResponseResultOutputsOutput `json:"output" xml:"output"`
}

type insAPIResponseResultOutputsOutput struct {
	Body    string `json:"body" xml:"body"`
	Code    string `json:"code" xml:"code"`
	Message string `json:"msg" xml:"msg"`
	Input   string `json:"input" xml:"input"`
}
