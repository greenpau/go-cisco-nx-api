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
	"reflect"
	"testing"
)

func TestParseShowVersion2JsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *VersionResponseResult
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.version",
			exp: &VersionResponseResult{
				Body: VersionResultBody{TablePackageList: []struct {
					RowPackageList []struct {
						PackageID []struct{} "json:\"package_id\""
					} "json:\"ROW_package_list\""
				}{struct {
					RowPackageList []struct {
						PackageID []struct{} "json:\"package_id\""
					} "json:\"ROW_package_list\""
				}{RowPackageList: []struct {
					PackageID []struct{} "json:\"package_id\""
				}{struct {
					PackageID []struct{} "json:\"package_id\""
				}{PackageID: []struct{}{}}}}}, BiosCmplTime: "10/18/2016", BiosVerStr: "08.32", BootflashSize: 21693714, ChassisID: "Nexus9000 C9508 (8 Slot) Chassis", CPUName: "Intel(R) Xeon(R) CPU E5-2403 0 @ 1.80GHz", HeaderStr: "Cisco Nexus Operating System (NX-OS) Software\nTAC support: http://www.cisco.com/tac\nCopyright (C) 2002-2018, Cisco and/or its affiliates.\nAll rights reserved.\nThe copyrights to certain works contained in this software are\nowned by other third parties and used and distributed under their own\nlicenses, such as open source.  This software is provided \"as is,\" and unless\notherwise stated, there is no warranty, express or implied, including but not\nlimited to warranties of merchantability and fitness for a particular purpose.\nCertain components of this software are licensed under\nthe GNU General Public License (GPL) version 2.0 or \nGNU General Public License (GPL) version 3.0  or the GNU\nLesser General Public License (LGPL) Version 2.1 or \nLesser General Public License (LGPL) Version 2.0. \nA copy of each such license is available at\nhttp://www.opensource.org/licenses/gpl-2.0.php and\nhttp://opensource.org/licenses/gpl-3.0.html and\nhttp://www.opensource.org/licenses/lgpl-2.1.php and\nhttp://www.gnu.org/licenses/old-licenses/library.txt.\n", HostName: "macsec2", KernUptmDays: 0, KernUptmHrs: 5, KernUptmMins: 3, KernUptmSecs: 24, KernUpTime: 18204000000000, KickCmplTime: " 5/22/2018 15:00:00", KickFileName: "bootflash:///nxos.7.0.3.I7.4.bin", KickTmstmp: "05/22/2018 15:26:08", KickstartVerStr: "7.0(3)I7(4)", Manufacturer: "Cisco Systems, Inc.", MemType: "kB", Memory: 16400780, ModuleID: "Supervisor Module", ProcBoardID: "SAL2015NQ3H", RrCtime: "Wed May 23 18:26:12 2018", RrReason: "Reset Requested by CLI command reload", RrService: "", RrSysVer: "7.0(3)I7(4)", RrUsecs: 681622}, Code: "200", Input: "show version", Msg: "Success"},
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
		dat, err := NewVersionFromBytes(content)
		//fmt.Printf("%#v\n", dat) //DEBUG
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *dat)
				testFailed++
				continue
			}
		}

		if dat != nil {
			if !reflect.DeepEqual(test.exp, dat) {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but failed due to mismatch", i, test.input)
				testFailed++
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
