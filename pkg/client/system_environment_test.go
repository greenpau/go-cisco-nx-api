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
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseShowSystemEnvironmentJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"
	for i, test := range []struct {
		input       string
		exp         *SystemEnvironment
		fanCount    int
		psCount     int
		sensorCount int
		shouldFail  bool
		shouldErr   bool
	}{
		{
			input:       "show.environment.1",
			exp:         &SystemEnvironment{},
			fanCount:    6,
			psCount:     2,
			sensorCount: 5,
			shouldFail:  false,
			shouldErr:   false,
		},
	} {
		fp := fmt.Sprintf("%s/resp.%s.json", outputDir, test.input)
		content, err := ioutil.ReadFile(fp)
		if err != nil {
			t.Logf("FAIL: Test %d: failed reading '%s', error: %v", i, fp, err)
			testFailed++
			continue
		}
		environment, err := NewSystemEnvironmentFromBytes(content)
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, environment)
				testFailed++
				continue
			}
		}

		if (len(environment.Fans) != test.fanCount) && !test.shouldFail {
			t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to fanCount [%d] != %d", i, test.input, len(environment.Fans), test.fanCount)
			testFailed++
			continue
		}

		if (len(environment.Sensors) != test.sensorCount) && !test.shouldFail {
			t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to sensorCount [%d] != %d", i, test.input, len(environment.Sensors), test.sensorCount)
			testFailed++
			continue
		}

		if (len(environment.PowerSupplies) != test.psCount) && !test.shouldFail {
			t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to psCount [%d] != %d", i, test.input, len(environment.PowerSupplies), test.psCount)
			testFailed++
			continue
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
