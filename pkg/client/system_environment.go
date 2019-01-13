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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type systemEnvironmentResponse struct {
	ID      uint64                          `json:"id" xml:"id"`
	Version string                          `json:"jsonrpc" xml:"jsonrpc"`
	Result  systemEnvironmentResponseResult `json:"result" xml:"result"`
}

type systemEnvironmentResponseResult struct {
	Body systemEnvironmentResponseResultBody `json:"body" xml:"body"`
}

type systemEnvironmentResponseResultBody struct {
	TempInfoTable systemEnvironmentResponseResultBodyTempInfoTable `json:"TABLE_tempinfo" xml:"TABLE_tempinfo"`
	FanDetails    systemEnvironmentResponseResultBodyFanDetails    `json:"fandetails" xml:"fandetails"`
	PsDetails     systemEnvironmentResponseResultBodyPsDetails     `json:"powersup" xml:"powersup"`
}

type systemEnvironmentResponseResultBodyTempInfoTable struct {
	TempInfoRow []systemEnvironmentResponseResultBodyTempInfoRow `json:"ROW_tempinfo" xml:"ROW_tempinfo"`
}

type systemEnvironmentResponseResultBodyTempInfoRow struct {
	AlarmStatus         string `json:"alarmstatus" xml:"alarmstatus"`
	Temperature         string `json:"curtemp" xml:"curtemp"`
	MaxTemperatureAlarm string `json:"majthres" xml:"majthres"`
	MinTemperatureAlarm string `json:"minthres" xml:"minthres"`
	Sensor              string `json:"sensor" xml:"sensor"`
	TemperatureModule   string `json:"tempmod" xml:"tempmod"`
}

type systemEnvironmentResponseResultBodyFanDetails struct {
	FanInfoTable systemEnvironmentResponseResultBodyFanInfoTable `json:"TABLE_faninfo" xml:"TABLE_faninfo"`
}

type systemEnvironmentResponseResultBodyFanInfoTable struct {
	FanInfoRow []systemEnvironmentResponseResultBodyFanInfoRow `json:"ROW_faninfo" xml:"ROW_faninfo"`
}

type systemEnvironmentResponseResultBodyFanInfoRow struct {
	Direction string `json:"fandir" xml:"fandir"`
	Model     string `json:"fanmodel" xml:"fanmodel"`
	Name      string `json:"fanname" xml:"fanname"`
	Status    string `json:"fanstatus" xml:"fanstatus"`
}

type systemEnvironmentResponseResultBodyPsDetails struct {
	PsInfoTable systemEnvironmentResponseResultBodyPsInfoTable `json:"TABLE_psinfo" xml:"TABLE_psinfo"`
}

type systemEnvironmentResponseResultBodyPsInfoTable struct {
	PsInfoRow []systemEnvironmentResponseResultBodyPsInfoRow `json:"ROW_psinfo" xml:"ROW_psinfo"`
}

type systemEnvironmentResponseResultBodyPsInfoRow struct {
	ID            int    `json:"psnum" xml:"psnum"`
	Model         string `json:"psmodel" xml:"psmodel"`
	PowerInput    string `json:"actual_input" xml:"actual_input"`
	PowerOutput   string `json:"actual_out" xml:"actual_out"`
	PowerCapacity string `json:"tot_capa" xml:"tot_capa"`
	Status        string `json:"ps_status" xml:"ps_status"`
}

// Fan holds information about a fan.
type Fan struct {
	Direction string `json:"direction" xml:"direction"`
	Model     string `json:"model" xml:"model"`
	Name      string `json:"name" xml:"name"`
	Status    string `json:"status" xml:"status"`
}

// PowerSupply holds information about a power supply.
type PowerSupply struct {
	ID            int     `json:"id" xml:"id"`
	Model         string  `json:"model" xml:"model"`
	PowerInput    float64 `json:"power_input" xml:"power_input"`
	PowerOutput   float64 `json:"power_output" xml:"power_output"`
	PowerCapacity float64 `json:"power_capacity" xml:"power_capacity"`
	Status        string  `json:"status" xml:"status"`
}

// Sensor holds information about a temperature sensor.
type Sensor struct {
	Name          string  `json:"name" xml:"name"`
	Module        uint64  `json:"module" xml:"module"`
	Temperature   float64 `json:"temperature" xml:"temperature"`
	ThresholdHigh float64 `json:"threshold_high" xml:"threshold_high"`
	ThresholdLow  float64 `json:"threshold_low" xml:"threshold_low"`
	Status        string  `json:"status" xml:"status"`
}

// SystemEnvironment contains information about fans, power supplies, and
// sensors. The information in the structure is from the output of
// "show environment" command.
type SystemEnvironment struct {
	Fans          []*Fan         `json:"fans" xml:"fans"`
	PowerSupplies []*PowerSupply `json:"power_supplies" xml:"power_supplies"`
	Sensors       []*Sensor      `json:"sensors" xml:"sensors"`
}

// NewSystemEnvironmentFromString returns SystemEnvironment instance from an input string.
func NewSystemEnvironmentFromString(s string) (*SystemEnvironment, error) {
	return NewSystemEnvironmentFromBytes([]byte(s))
}

// NewSystemEnvironmentFromBytes returns SystemEnvironment instance from an input byte array.
func NewSystemEnvironmentFromBytes(s []byte) (*SystemEnvironment, error) {
	var systemEnvironment *SystemEnvironment
	vResponse := &systemEnvironmentResponse{}
	err := json.Unmarshal(s, vResponse)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	if len(vResponse.Result.Body.TempInfoTable.TempInfoRow) < 1 && len(vResponse.Result.Body.FanDetails.FanInfoTable.FanInfoRow) < 1 {
		return nil, fmt.Errorf("Error parsing the received response: %s", s)
	}

	systemEnvironment = &SystemEnvironment{
		Fans:          []*Fan{},
		PowerSupplies: []*PowerSupply{},
		Sensors:       []*Sensor{},
	}

	for _, t := range vResponse.Result.Body.TempInfoTable.TempInfoRow {
		ts := &Sensor{}
		ts.Status = strings.TrimSpace(t.AlarmStatus)
		ts.Name = strings.TrimSpace(t.Sensor)
		if i, err := strconv.ParseUint(strings.TrimSpace(t.TemperatureModule), 10, 64); err == nil {
			ts.Module = i
		}
		if i, err := strconv.ParseFloat(strings.TrimSpace(t.Temperature), 64); err == nil {
			ts.Temperature = i
		}
		if i, err := strconv.ParseFloat(strings.TrimSpace(t.MaxTemperatureAlarm), 64); err == nil {
			ts.ThresholdHigh = i
		}
		if i, err := strconv.ParseFloat(strings.TrimSpace(t.MinTemperatureAlarm), 64); err == nil {
			ts.ThresholdLow = i
		}
		systemEnvironment.Sensors = append(systemEnvironment.Sensors, ts)
	}

	for _, f := range vResponse.Result.Body.FanDetails.FanInfoTable.FanInfoRow {
		fn := &Fan{}
		fn.Direction = strings.TrimSpace(f.Direction)
		fn.Model = strings.TrimSpace(f.Model)
		fn.Name = strings.TrimSpace(f.Name)
		fn.Status = strings.TrimSpace(f.Status)
		systemEnvironment.Fans = append(systemEnvironment.Fans, fn)
	}

	for _, p := range vResponse.Result.Body.PsDetails.PsInfoTable.PsInfoRow {
		ps := &PowerSupply{}
		ps.ID = p.ID
		ps.Model = strings.TrimSpace(p.Model)
		if i, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(p.PowerInput), " W", "", -1), 64); err == nil {
			ps.PowerInput = i
		}
		if i, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(p.PowerOutput), " W", "", -1), 64); err == nil {
			ps.PowerOutput = i
		}
		if i, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(p.PowerCapacity), " W", "", -1), 64); err == nil {
			ps.PowerCapacity = i
		}
		ps.Status = strings.TrimSpace(p.Status)
		systemEnvironment.PowerSupplies = append(systemEnvironment.PowerSupplies, ps)
	}

	return systemEnvironment, nil
}
