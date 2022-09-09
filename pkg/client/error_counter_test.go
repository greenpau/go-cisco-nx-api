package client

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseShowErrorCountersJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"
	for i, test := range []struct {
		input      string
		exp        *ErrorCounters
		count      int
		shouldFail bool
		shouldErr  bool
	}{
		{
			input:      "show.error.counters.1",
			exp:        &ErrorCounters{},
			count:      1,
			shouldFail: false,
			shouldErr:  false,
		},
		{
			input:      "show.error.counters.2",
			exp:        &ErrorCounters{},
			count:      2,
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
		errorCounters, err := NewErrorCountersFromBytes(content)
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, errorCounters)
				testFailed++
				continue
			}
		}

		if errorCounters != nil {
			if (len(errorCounters) != test.count) && !test.shouldFail {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to len(errorCounters) [%d] != %d", i, test.input, len(errorCounters), test.count)
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
