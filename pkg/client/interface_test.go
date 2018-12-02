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

func TestParseShowInterfaceJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"
	for i, test := range []struct {
		input      string
		count      int
		shouldFail bool
		shouldErr  bool
	}{
		{
			input:      "show.interfaces.1",
			count:      129,
			shouldFail: false,
			shouldErr:  false,
		},
		{
			input:      "show.interfaces.2",
			count:      14,
			shouldFail: false,
			shouldErr:  false,
		},
		{
			input:      "show.interfaces.3",
			count:      10,
			shouldFail: false,
			shouldErr:  false,
		},
		{
			input:      "show.interfaces.4",
			count:      153,
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
		interfaces, err := NewInterfacesFromBytes(content)
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, interfaces)
				testFailed++
				continue
			}
		}

		if interfaces != nil {
			if (len(interfaces) != test.count) && !test.shouldFail {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to len(interfaces) [%d] != %d", i, test.input, len(interfaces), test.count)
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

func TestParseShowInterfaceEthernet(t *testing.T) {
	outputDir := "../../assets/requests"
	fn := "resp.show.interfaces.5.json"
	fp := fmt.Sprintf("%s/%s", outputDir, fn)
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("FAIL: failed reading '%s', error: %v", fp, err)
	}
	interfaces, err := NewInterfacesFromBytes(content)
	if err != nil {
		t.Fatalf("FAIL: in '%s', error: %v", fp, err)
	}
	if len(interfaces) != 1 {
		t.Fatalf("FAIL: in '%s', error: len(interfaces) [%d] != 1", fp, len(interfaces))
	}
	t.Logf("PASS: ok")
}

func TestParseShowInterfaceSvi(t *testing.T) {
	outputDir := "../../assets/requests"
	fn := "resp.show.interfaces.6.json"
	fp := fmt.Sprintf("%s/%s", outputDir, fn)
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("FAIL: failed reading '%s', error: %v", fp, err)
	}
	interfaces, err := NewInterfacesFromBytes(content)
	if err != nil {
		t.Fatalf("FAIL: in '%s', error: %v", fp, err)
	}
	if len(interfaces) != 1 {
		t.Fatalf("FAIL: in '%s', error: len(interfaces) [%d] != 1", fp, len(interfaces))
	}
	t.Logf("PASS: ok")
}

func TestParseShowInterfaceMgmt(t *testing.T) {
	outputDir := "../../assets/requests"
	fn := "resp.show.interfaces.7.json"
	fp := fmt.Sprintf("%s/%s", outputDir, fn)
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("FAIL: failed reading '%s', error: %v", fp, err)
	}
	interfaces, err := NewInterfacesFromBytes(content)
	if err != nil {
		t.Fatalf("FAIL: in '%s', error: %v", fp, err)
	}
	if len(interfaces) != 1 {
		t.Fatalf("FAIL: in '%s', error: len(interfaces) [%d] != 1", fp, len(interfaces))
	}
	t.Logf("PASS: ok")
}
