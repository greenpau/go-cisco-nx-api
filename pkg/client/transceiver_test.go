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

func TestParseShowTransceiverJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"
	for i, test := range []struct {
		input      string
		count      int
		shouldFail bool
		shouldErr  bool
	}{
		{
			input:      "show.interface.transceiver.details.1",
			count:      2,
			shouldFail: false,
			shouldErr:  false,
		},
		{
			input:      "show.interface.transceiver.details.2",
			count:      1,
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
		transceivers, err := NewTransceiversFromBytes(content)
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed", i, test.input)
				testFailed++
				continue
			}
		}

		if (len(transceivers) != test.count) && !test.shouldFail {
			t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to count [%d] != %d", i, test.input, len(transceivers), test.count)
			testFailed++
			continue
		} else {
			for _, tr := range transceivers {
				t.Logf("INFO: Test %d: input '%s'\n%s", i, test.input, tr)
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
