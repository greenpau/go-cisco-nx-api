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
	"bytes"
	"fmt"
	"github.com/pschou/go-json"
	"io"
	"strings"
	"time"
)

type BGPSessionsResponse struct {
	InsAPI struct {
		Outputs struct {
			Output BGPSessionsResponseResult `json:"output" xml:"output"`
		} `json:"outputs" xml:"outputs"`
		Sid     string `json:"sid" xml:"sid"`
		Type    string `json:"type" xml:"type"`
		Version string `json:"version" xml:"version"`
	} `json:"ins_api" xml:"ins_api"`
}

type BGPSessionsResponseResult struct {
	Body  BGPSessionsResultBody `json:"body" xml:"body"`
	Code  string                `json:"code" xml:"code"`
	Input string                `json:"input" xml:"input"`
	Msg   string                `json:"msg" xml:"msg"`
}

type BGPSessionsResultBody struct {
	TableVrf []struct {
		RowVrf []struct {
			TableNeighbor []struct {
				RowNeighbor []struct {
					ConnectionsDropped    int    `json:"connectionsdropped" xml:"connectionsdropped"`
					LastFlap              string `json:"lastflap" xml:"lastflap"`
					LastRead              string `json:"lastread,omitempty" xml:"lastread,omitempty"`
					LastWrite             string `json:"lastwrite,omitempty" xml:"lastwrite,omitempty"`
					LocalPort             int    `json:"localport" xml:"localport"`
					NeighborID            string `json:"neighbor-id" xml:"neighbor-id"`
					NotificationsReceived int    `json:"notificationsreceived" xml:"notificationsreceived"`
					NotificationsSent     int    `json:"notificationssent" xml:"notificationssent"`
					RemoteAS              int    `json:"remoteas" xml:"remoteas"`
					RemotePort            int    `json:"remoteport" xml:"remoteport"`
					State                 string `json:"state" xml:"state"`
				} `json:"ROW_neighbor" xml:"ROW_neighbor"`
			} `json:"TABLE_neighbor" xml:"TABLE_neighbor"`
			LocalAS             int    `json:"local-as" xml:"local-as"`
			RouterID            string `json:"router-id" xml:"router-id"`
			VrfNameOut          string `json:"vrf-name-out" xml:"vrf-name-out"`
			VrfEstablishedPeers int    `json:"vrfestablishedpeers" xml:"vrfestablishedpeers"`
			VrfPeers            int    `json:"vrfpeers" xml:"vrfpeers"`
		} `json:"ROW_vrf" xml:"ROW_vrf"`
	} `json:"TABLE_vrf" xml:"TABLE_vrf"`
	LocalAS               int `json:"localas" xml:"localas"`
	TotalEstablishedPeers int `json:"totalestablishedpeers" xml:"totalestablishedpeers"`
	TotalPeers            int `json:"totalpeers" xml:"totalpeers"`
}

type BGPSessionsResultFlat struct {
	ConnectionsDropped    int           `json:"connectionsdropped" xml:"connectionsdropped"`
	LastFlap              time.Duration `json:"lastflap" xml:"lastflap"`
	LastRead              time.Duration `json:"lastread,omitempty" xml:"lastread,omitempty"`
	LastWrite             time.Duration `json:"lastwrite,omitempty" xml:"lastwrite,omitempty"`
	LocalPort             int           `json:"localport" xml:"localport"`
	NeighborID            string        `json:"neighbor-id" xml:"neighbor-id"`
	NotificationsReceived int           `json:"notificationsreceived" xml:"notificationsreceived"`
	NotificationsSent     int           `json:"notificationssent" xml:"notificationssent"`
	RemoteAS              int           `json:"remoteas" xml:"remoteas"`
	RemotePort            int           `json:"remoteport" xml:"remoteport"`
	State                 string        `json:"state" xml:"state"`
	LocalAS               int           `json:"local-as" xml:"local-as"`
	RouterID              string        `json:"router-id" xml:"router-id"`
	VrfNameOut            string        `json:"vrf-name-out" xml:"vrf-name-out"`
	VrfEstablishedPeers   int           `json:"vrfestablishedpeers" xml:"vrfestablishedpeers"`
	VrfPeers              int           `json:"vrfpeers" xml:"vrfpeers"`
}

func (d *BGPSessionsResponseResult) Flat() (out []BGPSessionsResultFlat) {
	for _, Tv := range d.Body.TableVrf {
		for _, Rv := range Tv.RowVrf {
			for _, Tn := range Rv.TableNeighbor {
				for _, Rn := range Tn.RowNeighbor {
					out = append(out, BGPSessionsResultFlat{
						ConnectionsDropped:    Rn.ConnectionsDropped,
						LastFlap:              ParseDuration(Rn.LastFlap),
						LastRead:              ParseDuration(Rn.LastRead),
						LastWrite:             ParseDuration(Rn.LastWrite),
						LocalPort:             Rn.LocalPort,
						NeighborID:            Rn.NeighborID,
						NotificationsReceived: Rn.NotificationsReceived,
						NotificationsSent:     Rn.NotificationsSent,
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

func (d *BGPSessionsResponse) Flat() (out []BGPSessionsResultFlat) {
	return d.InsAPI.Outputs.Output.Flat()
}

// NewBGPSessionsFromString returns instance from an input string.
func NewBGPSessionsFromString(s string) (*BGPSessionsResponse, error) {
	return NewBGPSessionsFromReader(strings.NewReader(s))
}

// NewBGPSessionsFromBytes returns instance from an input byte array.
func NewBGPSessionsFromBytes(s []byte) (*BGPSessionsResponse, error) {
	return NewBGPSessionsFromReader(bytes.NewReader(s))
}

// NewBGPSessionsFromReader returns instance from an input reader.
func NewBGPSessionsFromReader(s io.Reader) (*BGPSessionsResponse, error) {
	//si := &BGPSessions{}
	BGPSessionsResponseDat := &BGPSessionsResponse{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(BGPSessionsResponseDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return BGPSessionsResponseDat, nil
}

// NewBGPSessionsResultFromString returns instance from an input string.
func NewBGPSessionsResultFromString(s string) (*BGPSessionsResponseResult, error) {
	return NewBGPSessionsResultFromReader(strings.NewReader(s))
}

// NewBGPSessionsResultFromBytes returns instance from an input byte array.
func NewBGPSessionsResultFromBytes(s []byte) (*BGPSessionsResponseResult, error) {
	return NewBGPSessionsResultFromReader(bytes.NewReader(s))
}

// NewBGPSessionsResultFromReader returns instance from an input reader.
func NewBGPSessionsResultFromReader(s io.Reader) (*BGPSessionsResponseResult, error) {
	//si := &BGPSessionsResponseResult{}
	BGPSessionsResponseResultDat := &BGPSessionsResponseResult{}
	jsonDec := json.NewDecoder(s)
	jsonDec.UseAutoConvert()
	jsonDec.UseSlice()
	err := jsonDec.Decode(BGPSessionsResponseResultDat)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s", err)
	}
	return BGPSessionsResponseResultDat, nil
}
