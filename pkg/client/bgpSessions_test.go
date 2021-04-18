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
		exp        *BGPSessionsResponse
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "show.bgp.sessions",
			exp: &BGPSessionsResponse{
				InsAPI: struct {
					Outputs struct {
						Output BGPSessionsResponseResult "json:\"output\" xml:\"output\""
					} "json:\"outputs\" xml:\"outputs\""
					Sid     string "json:\"sid\" xml:\"sid\""
					Type    string "json:\"type\" xml:\"type\""
					Version string "json:\"version\" xml:\"version\""
				}{Outputs: struct {
					Output BGPSessionsResponseResult "json:\"output\" xml:\"output\""
				}{Output: BGPSessionsResponseResult{Body: BGPSessionsResultBody{TableVrf: []struct {
					RowVrf []struct {
						TableNeighbor []struct {
							RowNeighbor []struct {
								ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
								LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
								LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
								LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
								LocalPort             int    "json:\"localport\" xml:\"localport\""
								NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
								NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
								NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
								RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
								RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
								State                 string "json:\"state\" xml:\"state\""
							} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
						} "json:\"TABLE_neighbor\" xml:\"TABLE_neighbor\""
						LocalAS             int    "json:\"local-as\" xml:\"local-as\""
						RouterID            string "json:\"router-id\" xml:\"router-id\""
						VrfNameOut          string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
						VrfEstablishedPeers int    "json:\"vrfestablishedpeers\" xml:\"vrfestablishedpeers\""
						VrfPeers            int    "json:\"vrfpeers\" xml:\"vrfpeers\""
					} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
				}{struct {
					RowVrf []struct {
						TableNeighbor []struct {
							RowNeighbor []struct {
								ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
								LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
								LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
								LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
								LocalPort             int    "json:\"localport\" xml:\"localport\""
								NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
								NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
								NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
								RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
								RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
								State                 string "json:\"state\" xml:\"state\""
							} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
						} "json:\"TABLE_neighbor\" xml:\"TABLE_neighbor\""
						LocalAS             int    "json:\"local-as\" xml:\"local-as\""
						RouterID            string "json:\"router-id\" xml:\"router-id\""
						VrfNameOut          string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
						VrfEstablishedPeers int    "json:\"vrfestablishedpeers\" xml:\"vrfestablishedpeers\""
						VrfPeers            int    "json:\"vrfpeers\" xml:\"vrfpeers\""
					} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
				}{RowVrf: []struct {
					TableNeighbor []struct {
						RowNeighbor []struct {
							ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
							LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
							LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
							LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
							LocalPort             int    "json:\"localport\" xml:\"localport\""
							NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
							NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
							NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
							RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
							RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
							State                 string "json:\"state\" xml:\"state\""
						} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
					} "json:\"TABLE_neighbor\" xml:\"TABLE_neighbor\""
					LocalAS             int    "json:\"local-as\" xml:\"local-as\""
					RouterID            string "json:\"router-id\" xml:\"router-id\""
					VrfNameOut          string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
					VrfEstablishedPeers int    "json:\"vrfestablishedpeers\" xml:\"vrfestablishedpeers\""
					VrfPeers            int    "json:\"vrfpeers\" xml:\"vrfpeers\""
				}{struct {
					TableNeighbor []struct {
						RowNeighbor []struct {
							ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
							LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
							LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
							LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
							LocalPort             int    "json:\"localport\" xml:\"localport\""
							NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
							NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
							NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
							RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
							RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
							State                 string "json:\"state\" xml:\"state\""
						} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
					} "json:\"TABLE_neighbor\" xml:\"TABLE_neighbor\""
					LocalAS             int    "json:\"local-as\" xml:\"local-as\""
					RouterID            string "json:\"router-id\" xml:\"router-id\""
					VrfNameOut          string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
					VrfEstablishedPeers int    "json:\"vrfestablishedpeers\" xml:\"vrfestablishedpeers\""
					VrfPeers            int    "json:\"vrfpeers\" xml:\"vrfpeers\""
				}{TableNeighbor: []struct {
					RowNeighbor []struct {
						ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
						LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
						LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
						LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
						LocalPort             int    "json:\"localport\" xml:\"localport\""
						NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
						NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
						NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
						RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
						RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
						State                 string "json:\"state\" xml:\"state\""
					} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
				}{struct {
					RowNeighbor []struct {
						ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
						LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
						LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
						LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
						LocalPort             int    "json:\"localport\" xml:\"localport\""
						NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
						NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
						NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
						RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
						RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
						State                 string "json:\"state\" xml:\"state\""
					} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
				}{RowNeighbor: []struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: 179, NeighborID: "19.0.101.1", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 333, RemotePort: 37989, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M31S", LastRead: "PT4S", LastWrite: "PT34S", LocalPort: 179, NeighborID: "19.0.102.3", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 888, RemotePort: 46134, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: 179, NeighborID: "19.0.102.4", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 333, RemotePort: 43271, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: 179, NeighborID: "19.0.103.10", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 999, RemotePort: 35158, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M28S", LastRead: "PT48S", LastWrite: "PT31S", LocalPort: 179, NeighborID: "19.0.103.20", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 333, RemotePort: 41777, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M45S", LastRead: "", LastWrite: "", LocalPort: 0, NeighborID: "19.0.200.200", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 0, RemotePort: 0, State: "Idle"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M48S", LastRead: "", LastWrite: "", LocalPort: 0, NeighborID: "fec0::1002", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 333, RemotePort: 0, State: "Idle"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M48S", LastRead: "", LastWrite: "", LocalPort: 0, NeighborID: "fec0::2002", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 888, RemotePort: 0, State: "Idle"}}}}, LocalAS: 333, RouterID: "19.0.0.6", VrfNameOut: "default", VrfEstablishedPeers: 5, VrfPeers: 8}}}}, LocalAS: 333, TotalEstablishedPeers: 5, TotalPeers: 9}, Code: "200", Input: "show bgp sessions", Msg: "Success"}}, Sid: "eoc", Type: "cli_show", Version: "1.0"}},
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
		bgp, err := NewBGPSessionsFromBytes(content)
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

func TestParseBGPShowSessionsResultJsonOutput(t *testing.T) {
	testFailed := 0
	outputDir := "../../assets/requests"

	for i, test := range []struct {
		input      string
		exp        *BGPSessionsResponseResult
		shouldFail bool
		shouldErr  bool
	}{
		{
			input: "result.show.bgp.sessions",
			exp: &BGPSessionsResponseResult{
				Body: BGPSessionsResultBody{TableVrf: []struct {
					RowVrf []struct {
						TableNeighbor []struct {
							RowNeighbor []struct {
								ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
								LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
								LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
								LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
								LocalPort             int    "json:\"localport\" xml:\"localport\""
								NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
								NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
								NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
								RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
								RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
								State                 string "json:\"state\" xml:\"state\""
							} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
						} "json:\"TABLE_neighbor\" xml:\"TABLE_neighbor\""
						LocalAS             int    "json:\"local-as\" xml:\"local-as\""
						RouterID            string "json:\"router-id\" xml:\"router-id\""
						VrfNameOut          string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
						VrfEstablishedPeers int    "json:\"vrfestablishedpeers\" xml:\"vrfestablishedpeers\""
						VrfPeers            int    "json:\"vrfpeers\" xml:\"vrfpeers\""
					} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
				}{struct {
					RowVrf []struct {
						TableNeighbor []struct {
							RowNeighbor []struct {
								ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
								LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
								LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
								LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
								LocalPort             int    "json:\"localport\" xml:\"localport\""
								NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
								NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
								NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
								RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
								RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
								State                 string "json:\"state\" xml:\"state\""
							} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
						} "json:\"TABLE_neighbor\" xml:\"TABLE_neighbor\""
						LocalAS             int    "json:\"local-as\" xml:\"local-as\""
						RouterID            string "json:\"router-id\" xml:\"router-id\""
						VrfNameOut          string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
						VrfEstablishedPeers int    "json:\"vrfestablishedpeers\" xml:\"vrfestablishedpeers\""
						VrfPeers            int    "json:\"vrfpeers\" xml:\"vrfpeers\""
					} "json:\"ROW_vrf\" xml:\"ROW_vrf\""
				}{RowVrf: []struct {
					TableNeighbor []struct {
						RowNeighbor []struct {
							ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
							LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
							LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
							LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
							LocalPort             int    "json:\"localport\" xml:\"localport\""
							NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
							NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
							NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
							RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
							RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
							State                 string "json:\"state\" xml:\"state\""
						} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
					} "json:\"TABLE_neighbor\" xml:\"TABLE_neighbor\""
					LocalAS             int    "json:\"local-as\" xml:\"local-as\""
					RouterID            string "json:\"router-id\" xml:\"router-id\""
					VrfNameOut          string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
					VrfEstablishedPeers int    "json:\"vrfestablishedpeers\" xml:\"vrfestablishedpeers\""
					VrfPeers            int    "json:\"vrfpeers\" xml:\"vrfpeers\""
				}{struct {
					TableNeighbor []struct {
						RowNeighbor []struct {
							ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
							LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
							LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
							LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
							LocalPort             int    "json:\"localport\" xml:\"localport\""
							NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
							NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
							NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
							RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
							RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
							State                 string "json:\"state\" xml:\"state\""
						} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
					} "json:\"TABLE_neighbor\" xml:\"TABLE_neighbor\""
					LocalAS             int    "json:\"local-as\" xml:\"local-as\""
					RouterID            string "json:\"router-id\" xml:\"router-id\""
					VrfNameOut          string "json:\"vrf-name-out\" xml:\"vrf-name-out\""
					VrfEstablishedPeers int    "json:\"vrfestablishedpeers\" xml:\"vrfestablishedpeers\""
					VrfPeers            int    "json:\"vrfpeers\" xml:\"vrfpeers\""
				}{TableNeighbor: []struct {
					RowNeighbor []struct {
						ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
						LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
						LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
						LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
						LocalPort             int    "json:\"localport\" xml:\"localport\""
						NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
						NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
						NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
						RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
						RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
						State                 string "json:\"state\" xml:\"state\""
					} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
				}{struct {
					RowNeighbor []struct {
						ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
						LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
						LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
						LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
						LocalPort             int    "json:\"localport\" xml:\"localport\""
						NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
						NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
						NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
						RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
						RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
						State                 string "json:\"state\" xml:\"state\""
					} "json:\"ROW_neighbor\" xml:\"ROW_neighbor\""
				}{RowNeighbor: []struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: 179, NeighborID: "19.0.101.1", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 333, RemotePort: 37989, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M31S", LastRead: "PT4S", LastWrite: "PT34S", LocalPort: 179, NeighborID: "19.0.102.3", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 888, RemotePort: 46134, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: 179, NeighborID: "19.0.102.4", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 333, RemotePort: 43271, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M31S", LastRead: "PT51S", LastWrite: "PT34S", LocalPort: 179, NeighborID: "19.0.103.10", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 999, RemotePort: 35158, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M28S", LastRead: "PT48S", LastWrite: "PT31S", LocalPort: 179, NeighborID: "19.0.103.20", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 333, RemotePort: 41777, State: "Established"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M45S", LastRead: "", LastWrite: "", LocalPort: 0, NeighborID: "19.0.200.200", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 0, RemotePort: 0, State: "Idle"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M48S", LastRead: "", LastWrite: "", LocalPort: 0, NeighborID: "fec0::1002", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 333, RemotePort: 0, State: "Idle"}, struct {
					ConnectionsDropped    int    "json:\"connectionsdropped\" xml:\"connectionsdropped\""
					LastFlap              string "json:\"lastflap\" xml:\"lastflap\""
					LastRead              string "json:\"lastread,omitempty\" xml:\"lastread,omitempty\""
					LastWrite             string "json:\"lastwrite,omitempty\" xml:\"lastwrite,omitempty\""
					LocalPort             int    "json:\"localport\" xml:\"localport\""
					NeighborID            string "json:\"neighbor-id\" xml:\"neighbor-id\""
					NotificationsReceived int    "json:\"notificationsreceived\" xml:\"notificationsreceived\""
					NotificationsSent     int    "json:\"notificationssent\" xml:\"notificationssent\""
					RemoteAS              int    "json:\"remoteas\" xml:\"remoteas\""
					RemotePort            int    "json:\"remoteport\" xml:\"remoteport\""
					State                 string "json:\"state\" xml:\"state\""
				}{ConnectionsDropped: 0, LastFlap: "P6DT2H33M48S", LastRead: "", LastWrite: "", LocalPort: 0, NeighborID: "fec0::2002", NotificationsReceived: 0, NotificationsSent: 0, RemoteAS: 888, RemotePort: 0, State: "Idle"}}}}, LocalAS: 333, RouterID: "19.0.0.6", VrfNameOut: "default", VrfEstablishedPeers: 5, VrfPeers: 8}}}}, LocalAS: 333, TotalEstablishedPeers: 5, TotalPeers: 9}, Code: "200", Input: "show bgp sessions", Msg: "Success"},
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
		bgp, err := NewBGPSessionsResultFromBytes(content)
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
