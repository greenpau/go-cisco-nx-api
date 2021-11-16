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
	"testing"
)

func TestParseShowVersionJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"
	for i, test := range []struct {
		input      string
		exp        *SysInfo
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.version.1",
			exp: &SysInfo{
				ChassisID: "Nexus9000 9000v Chassis",
				CPUName:   "Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz",
			},
			shouldFail: false,
			shouldErr:  false,
		},
		{
			input:      "show.version.2",
			exp:        nil,
			shouldFail: false,
			shouldErr:  true,
		},
		{
			input: "show.version.3",
			exp: &SysInfo{
				ChassisID: "Nexus9000 C9236C chassis",
				CPUName:   "Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz",
			},
			shouldFail: false,
			shouldErr:  false,
		},
		{
			input: "show.version.4",
			exp: &SysInfo{
				ChassisID: "Nexus9000 C93180YC-FX Chassis",
				CPUName:   "Intel(R) Xeon(R) CPU D-1528 @ 1.90GHz",
			},
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
		sysinfo, err := NewSysInfoFromBytes(content)
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *sysinfo)
				testFailed++
				continue
			}
		}

		if sysinfo != nil {
			if (sysinfo.ChassisID != test.exp.ChassisID) && !test.shouldFail {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to mismatch in '%s' field: '%v' (actual) vs. '%v' (expected)",
					i, test.input, "ChassisID", sysinfo.ChassisID, test.exp.ChassisID)
				testFailed++
				continue
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
