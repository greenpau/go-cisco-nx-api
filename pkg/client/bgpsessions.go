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
	"bytes"
	"fmt"
	"github.com/pschou/go-json"
	"io"
	"strings"
	"time"
)

// BGPSessionResponse is BGP Session Response.
type BGPSessionResponse struct {
	InsAPI struct {
		Outputs struct {
			Output BGPSessionResponseResult `json:"output"`
		} `json:"outputs"`
		Sid     string `json:"sid"`
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"ins_api"`
}

// BGPSessionResponseResult is the result of BGPSessionResponse.
type BGPSessionResponseResult struct {
	Body  BGPSessionResultBody `json:"body" xml:"body"`
	Code  string               `json:"code"`
	Input string               `json:"input"`
	Msg   string               `json:"msg"`
}

// BGPSessionResultBody is the body of the result of BGPSessionResponse.
type BGPSessionResultBody struct {
	TableVrf []struct {
		RowVrf []struct {
			TableNeighbor []struct {
				RowNeighbor []struct {
					ConnectionsDropped    string `json:"connectionsdropped"`
					LastFlap              string `json:"lastflap"`
					LastRead              string `json:"lastread,omitempty"`
					LastWrite             string `json:"lastwrite,omitempty"`
					LocalPort             string `json:"localport"`
					NeighborID            string `json:"neighbor-id"`
					NotificationsReceived string `json:"notificationsreceived"`
					NotificationsSent     string `json:"notificationssent"`
					RemoteAS              string `json:"remoteas"`
					RemotePort            string `json:"remoteport"`
					State                 string `json:"state"`
				} `json:"ROW_neighbor"`
			} `json:"TABLE_neighbor"`
			LocalAS             string `json:"local-as"`
			RouterID            string `json:"router-id"`
			VrfNameOut          string `json:"vrf-name-out"`
			VrfEstablishedPeers string `json:"vrfestablishedpeers"`
			VrfPeers            string `json:"vrfpeers"`
		} `json:"ROW_vrf"`
	} `json:"TABLE_vrf"`
	LocalAS               string `json:"localas"`
	TotalEstablishedPeers string `json:"totalestablishedpeers"`
	TotalPeers            string `json:"totalpeers"`
}

// Flat flattens BGPSessionResponse.
func (d *BGPSessionResponse) Flat() (out []BGPSessionResultFlat) {
	for _, Tv := range d.InsAPI.Outputs.Output.Body.TableVrf {
		for _, Rv := range Tv.RowVrf {
			for _, Tn := range Rv.TableNeighbor {
				for _, Rn := range Tn.RowNeighbor {
					out = append(out, BGPSessionResultFlat{
						ConnectionsDropped:    StrInt(Rn.ConnectionsDropped),
						LastFlap:              ParseDuration(Rn.LastFlap),
						LastRead:              ParseDuration(Rn.LastRead),
						LastWrite:             ParseDuration(Rn.LastWrite),
						LocalPort:             Rn.LocalPort,
						NeighborID:            Rn.NeighborID,
						NotificationsReceived: StrInt(Rn.NotificationsReceived),
						NotificationsSent:     StrInt(Rn.NotificationsSent),
						RemoteAS:              Rn.RemoteAS,
						RemotePort:            Rn.RemotePort,
						State:                 Rn.State,
						LocalAS:               Rv.LocalAS,
						RouterID:              Rv.RouterID,
						VrfNameOut:            Rv.VrfNameOut,
						VrfEstablishedPeers:   Rv.VrfEstablishedPeers,
						VrfPeers:              Rv.VrfPeers,
					})
				}
			}
		}
	}
	return
}

// BGPSessionResultFlat holds flat BGPSessionResult.
type BGPSessionResultFlat struct {
	ConnectionsDropped    int           `json:"connectionsdropped"`
	LastFlap              time.Duration `json:"lastflap"`
	LastRead              time.Duration `json:"lastread,omitempty"`
	LastWrite             time.Duration `json:"lastwrite,omitempty"`
	LocalPort             string        `json:"localport"`
	NeighborID            string        `json:"neighbor-id"`
	NotificationsReceived int           `json:"notificationsreceived"`
	NotificationsSent     int           `json:"notificationssent"`
	RemoteAS              string        `json:"remoteas"`
	RemotePort            string        `json:"remoteport"`
	State                 string        `json:"state"`
	LocalAS               string        `json:"local-as"`
	RouterID              string        `json:"router-id"`
	VrfNameOut            string        `json:"vrf-name-out"`
	VrfEstablishedPeers   string        `json:"vrfestablishedpeers"`
	VrfPeers              string        `json:"vrfpeers"`
}

// NewBGPSessionFromString returns SysInfo instance from an input string.
func NewBGPSessionFromString(s string) (*BGPSessionResponse, error) {
	return NewBGPSessionFromReader(strings.NewReader(s))
}

// NewSysInfoFromBytes returns SysInfo instance from an input byte array.
func NewBGPSessionFromBytes(s []byte) (*BGPSessionResponse, error) {
	return NewBGPSessionFromReader(bytes.NewReader(s))
}
func NewBGPSessionFromReader(s io.Reader) (*BGPSessionResponse, error) {
	//si := &BGPSession{}
	BGPSessionResponseDat := &BGPSessionResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseNumber()
	jsonDec.UseSlice()
	err := jsonDec.Decode(BGPSessionResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return BGPSessionResponseDat, nil
}
