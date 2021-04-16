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

func TestParseBGPShowSessionsJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *BGPSessionResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.bgp.sessions",
			exp: &BGPSessionResponse{
				InsAPI: struct {
					Outputs struct {
						Output BGPSessionResponseResult "json:\"output\""
					} "json:\"outputs\""
					Sid     string "json:\"sid\""
					Type    string "json:\"type\""
					Version string "json:\"version\""
				}{Outputs: struct {
					Output BGPSessionResponseResult "json:\"output\""
				}{Output: BGPSessionResponseResult{Body: BGPSessionResultBody{TableVrf: []struct {
					RowVrf []struct {
						TableNeighbor []struct {
							RowNeighbor []struct {
								ConnectionsDropped    string "json:\"connectionsdropped\""
								LastFlap              string "json:\"lastflap\""
								LastRead              string "json:\"lastread,omitempty\""
								LastWrite             string "json:\"lastwrite,omitempty\""
								LocalPort             string "json:\"localport\""
								NeighborID            string "json:\"neighbor-id\""
								NotificationsReceived string "json:\"notificationsreceived\""
								NotificationsSent     string "json:\"notificationssent\""
								RemoteAS              string "json:\"remoteas\""
								RemotePort            string "json:\"remoteport\""
								State                 string "json:\"state\""
							} "json:\"ROW_neighbor\""
						} "json:\"TABLE_neighbor\""
						LocalAS             string "json:\"local-as\""
						RouterID            string "json:\"router-id\""
						VrfNameOut          string "json:\"vrf-name-out\""
						VrfEstablishedPeers string "json:\"vrfestablishedpeers\""
						VrfPeers            string "json:\"vrfpeers\""
					} "json:\"ROW_vrf\""
				}{struct {
					RowVrf []struct {
						TableNeighbor []struct {
							RowNeighbor []struct {
								ConnectionsDropped    string "json:\"connectionsdropped\""
								LastFlap              string "json:\"lastflap\""
								LastRead              string "json:\"lastread,omitempty\""
								LastWrite             string "json:\"lastwrite,omitempty\""
								LocalPort             string "json:\"localport\""
								NeighborID            string "json:\"neighbor-id\""
								NotificationsReceived string "json:\"notificationsreceived\""
								NotificationsSent     string "json:\"notificationssent\""
								RemoteAS              string "json:\"remoteas\""
								RemotePort            string "json:\"remoteport\""
								State                 string "json:\"state\""
							} "json:\"ROW_neighbor\""
						} "json:\"TABLE_neighbor\""
						LocalAS             string "json:\"local-as\""
						RouterID            string "json:\"router-id\""
						VrfNameOut          string "json:\"vrf-name-out\""
						VrfEstablishedPeers string "json:\"vrfestablishedpeers\""
						VrfPeers            string "json:\"vrfpeers\""
					} "json:\"ROW_vrf\""
				}{RowVrf: []struct {
					TableNeighbor []struct {
						RowNeighbor []struct {
							ConnectionsDropped    string "json:\"connectionsdropped\""
							LastFlap              string "json:\"lastflap\""
							LastRead              string "json:\"lastread,omitempty\""
							LastWrite             string "json:\"lastwrite,omitempty\""
							LocalPort             string "json:\"localport\""
							NeighborID            string "json:\"neighbor-id\""
							NotificationsReceived string "json:\"notificationsreceived\""
							NotificationsSent     string "json:\"notificationssent\""
							RemoteAS              string "json:\"remoteas\""
							RemotePort            string "json:\"remoteport\""
							State                 string "json:\"state\""
						} "json:\"ROW_neighbor\""
					} "json:\"TABLE_neighbor\""
					LocalAS             string "json:\"local-as\""
					RouterID            string "json:\"router-id\""
					VrfNameOut          string "json:\"vrf-name-out\""
					VrfEstablishedPeers string "json:\"vrfestablishedpeers\""
					VrfPeers            string "json:\"vrfpeers\""
				}{struct {
					TableNeighbor []struct {
						RowNeighbor []struct {
							ConnectionsDropped    string "json:\"connectionsdropped\""
							LastFlap              string "json:\"lastflap\""
							LastRead              string "json:\"lastread,omitempty\""
							LastWrite             string "json:\"lastwrite,omitempty\""
							LocalPort             string "json:\"localport\""
							NeighborID            string "json:\"neighbor-id\""
							NotificationsReceived string "json:\"notificationsreceived\""
							NotificationsSent     string "json:\"notificationssent\""
							RemoteAS              string "json:\"remoteas\""
							RemotePort            string "json:\"remoteport\""
							State                 string "json:\"state\""
						} "json:\"ROW_neighbor\""
					} "json:\"TABLE_neighbor\""
					LocalAS             string "json:\"local-as\""
					RouterID            string "json:\"router-id\""
					VrfNameOut          string "json:\"vrf-name-out\""
					VrfEstablishedPeers string "json:\"vrfestablishedpeers\""
					VrfPeers            string "json:\"vrfpeers\""
				}{TableNeighbor: []struct {
					RowNeighbor []struct {
						ConnectionsDropped    string "json:\"connectionsdropped\""
						LastFlap              string "json:\"lastflap\""
						LastRead              string "json:\"lastread,omitempty\""
						LastWrite             string "json:\"lastwrite,omitempty\""
						LocalPort             string "json:\"localport\""
						NeighborID            string "json:\"neighbor-id\""
						NotificationsReceived string "json:\"notificationsreceived\""
						NotificationsSent     string "json:\"notificationssent\""
						RemoteAS              string "json:\"remoteas\""
						RemotePort            string "json:\"remoteport\""
						State                 string "json:\"state\""
					} "json:\"ROW_neighbor\""
				}{struct {
					RowNeighbor []struct {
						ConnectionsDropped    string "json:\"connectionsdropped\""
						LastFlap              string "json:\"lastflap\""
						LastRead              string "json:\"lastread,omitempty\""
						LastWrite             string "json:\"lastwrite,omitempty\""
						LocalPort             string "json:\"localport\""
						NeighborID            string "json:\"neighbor-id\""
						NotificationsReceived string "json:\"notificationsreceived\""
						NotificationsSent     string "json:\"notificationssent\""
						RemoteAS              string "json:\"remoteas\""
						RemotePort            string "json:\"remoteport\""
						State                 string "json:\"state\""
					} "json:\"ROW_neighbor\""
				}{RowNeighbor: []struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{ConnectionsDropped: "0", LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: "179", NeighborID: "19.0.101.1", NotificationsReceived: "0", NotificationsSent: "0", RemoteAS: "333", RemotePort: "37989", State: "Established"}, struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{ConnectionsDropped: "0", LastFlap: "P6DT2H33M31S", LastRead: "PT4S", LastWrite: "PT34S", LocalPort: "179", NeighborID: "19.0.102.3", NotificationsReceived: "0", NotificationsSent: "0", RemoteAS: "888", RemotePort: "46134", State: "Established"}, struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{ConnectionsDropped: "0", LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: "179", NeighborID: "19.0.102.4", NotificationsReceived: "0", NotificationsSent: "0", RemoteAS: "333", RemotePort: "43271", State: "Established"}, struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{ConnectionsDropped: "0", LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: "179", NeighborID: "19.0.103.10", NotificationsReceived: "0", NotificationsSent: "0", RemoteAS: "999", RemotePort: "35158", State: "Established"}, struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{ConnectionsDropped: "0", LastFlap: "P6DT2H33M28S", LastRead: "PT48S", LastWrite: "PT31S", LocalPort: "179", NeighborID: "19.0.103.20", NotificationsReceived: "0", NotificationsSent: "0", RemoteAS: "333", RemotePort: "41777", State: "Established"}, struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{ConnectionsDropped: "0", LastFlap: "P6DT2H33M45S", LastRead: "", LastWrite: "", LocalPort: "0", NeighborID: "19.0.200.200", NotificationsReceived: "0", NotificationsSent: "0", RemoteAS: "0", RemotePort: "0", State: "Idle"}, struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{ConnectionsDropped: "0", LastFlap: "P6DT2H33M48S", LastRead: "", LastWrite: "", LocalPort: "0", NeighborID: "fec0::1002", NotificationsReceived: "0", NotificationsSent: "0", RemoteAS: "333", RemotePort: "0", State: "Idle"}, struct {
					ConnectionsDropped    string "json:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\""
					LocalPort             string "json:\"localport\""
					NeighborID            string "json:\"neighbor-id\""
					NotificationsReceived string "json:\"notificationsreceived\""
					NotificationsSent     string "json:\"notificationssent\""
					RemoteAS              string "json:\"remoteas\""
					RemotePort            string "json:\"remoteport\""
					State                 string "json:\"state\""
				}{ConnectionsDropped: "0", LastFlap: "P6DT2H33M48S", LastRead: "", LastWrite: "", LocalPort: "0", NeighborID: "fec0::2002", NotificationsReceived: "0", NotificationsSent: "0", RemoteAS: "888", RemotePort: "0", State: "Idle"}}}}, LocalAS: "333", RouterID: "19.0.0.6", VrfNameOut: "default", VrfEstablishedPeers: "5", VrfPeers: "8"}}}}, LocalAS: "333", TotalEstablishedPeers: "5", TotalPeers: "9"}, Code: "200", Input: "show bgp sessions", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"},
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
		bgp, err := NewBGPSessionFromBytes(content)
		//fmt.Printf("%#v", bgp) //DEBUG
		//fmt.Printf("%+v", bgp.Flat()) //DEBUG
		if err != nil {
			if !test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to pass, but threw error: %v", i, test.input, err)
				testFailed++
				continue
			}
		} else {
			if test.shouldErr {
				t.Logf("FAIL: Test %d: input '%s', expected to throw error, but passed: %v", i, test.input, *bgp)
				testFailed++
				continue
			}
		}

		if bgp != nil {
			if !reflect.DeepEqual(test.exp, bgp) {
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
