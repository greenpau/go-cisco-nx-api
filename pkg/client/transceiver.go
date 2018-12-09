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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type transceiverResponse struct {
	ID      uint64                    `json:"id" xml:"id"`
	Version string                    `json:"jsonrpc" xml:"jsonrpc"`
	Result  transceiverResponseResult `json:"result" xml:"result"`
}

type transceiverResponseResult struct {
	Body transceiverResponseResultBody `json:"body" xml:"body"`
}

type transceiverResponseResultBody struct {
	TransceiverTable transceiverResponseResultBodyTransceiverTable `json:"TABLE_interface" xml:"TABLE_interface"`
}

type transceiverResponseResultBodyTransceiverTable struct {
	TransceiverRow []transceiverResponseResultBodyTransceiverRow `json:"ROW_interface" xml:"ROW_interface"`
}

type transceiverResponseResultBodyTransceiverLaneTable struct {
	TransceiverLaneRow []transceiverResponseResultBodyTransceiverLaneRow `json:"ROW_lane" xml:"ROW_lane"`
}

func (t *transceiverResponseResultBodyTransceiverLaneTable) UnmarshalJSON(b []byte) error {
	size := len(b)
	i := bytes.IndexByte(b, byte(':'))
	if i < 0 {
		return fmt.Errorf("Error unmarshalling transceiverResponseResultBodyTransceiverLaneTable")
	}
	j := bytes.IndexByte(b[i:], byte('['))
	switch {
	case j > 10 || j < 0:
		// single entry
		var r transceiverResponseResultBodyTransceiverLaneRow
		err := json.Unmarshal(b[(i+2):size-1], &r)
		if err != nil {
			return fmt.Errorf("Error unmarshalling transceiverResponseResultBodyTransceiverLaneTable: %s", err)
		}
		t.TransceiverLaneRow = append(t.TransceiverLaneRow, r)
	case j < 10 && j >= 0:
		// multiple entries
		var r []transceiverResponseResultBodyTransceiverLaneRow
		err := json.Unmarshal(b[(i+2):size-1], &r)
		if err != nil {
			return fmt.Errorf("Error unmarshalling transceiverResponseResultBodyTransceiverLaneTable: %s", err)
		}
		t.TransceiverLaneRow = r
	}
	return nil
}

type transceiverResponseResultBodyTransceiverLaneRow struct {
	LaneNumber             uint64 `json:"lane_number" xml:"lane_number"`
	Temperature            string `json:"temperature" xml:"temperature"`
	TemperatureAlarmHigh   string `json:"temp_alrm_hi" xml:"temp_alrm_hi"`
	TemperatureAlarmLow    string `json:"temp_alrm_lo" xml:"temp_alrm_lo"`
	TemperatureWarningHigh string `json:"temp_warn_hi" xml:"temp_warn_hi"`
	TemperatureWarningLow  string `json:"temp_warn_lo" xml:"temp_warn_lo"`
	Voltage                string `json:"voltage" xml:"voltage"`
	VoltageAlarmHigh       string `json:"volt_alrm_hi" xml:"volt_alrm_hi"`
	VoltageAlarmLow        string `json:"volt_alrm_lo" xml:"volt_alrm_lo"`
	VoltageWarningHigh     string `json:"volt_warn_hi" xml:"volt_warn_hi"`
	VoltageWarningLow      string `json:"volt_warn_lo" xml:"volt_warn_lo"`
	Current                string `json:"current" xml:"current"`
	CurrentAlarmHigh       string `json:"current_alrm_hi" xml:"current_alrm_hi"`
	CurrentAlarmLow        string `json:"current_alrm_lo" xml:"current_alrm_lo"`
	CurrentWarningHigh     string `json:"current_warn_hi" xml:"current_warn_hi"`
	CurrentWarningLow      string `json:"current_warn_lo" xml:"current_warn_lo"`
	TxPower                string `json:"tx_pwr" xml:"tx_pwr"`
	TxPowerAlarmHigh       string `json:"tx_pwr_alrm_hi" xml:"tx_pwr_alrm_hi"`
	TxPowerAlarmLow        string `json:"tx_pwr_alrm_lo" xml:"tx_pwr_alrm_lo"`
	TxPowerWarningHigh     string `json:"tx_pwr_warn_hi" xml:"tx_pwr_warn_hi"`
	TxPowerWarningLow      string `json:"tx_pwr_warn_lo" xml:"tx_pwr_warn_lo"`
	RxPower                string `json:"rx_pwr" xml:"rx_pwr"`
	RxPowerAlarmHigh       string `json:"rx_pwr_alrm_hi" xml:"rx_pwr_alrm_hi"`
	RxPowerAlarmLow        string `json:"rx_pwr_alrm_lo" xml:"rx_pwr_alrm_lo"`
	RxPowerWarningHigh     string `json:"rx_pwr_warn_hi" xml:"rx_pwr_warn_hi"`
	RxPowerWarningLow      string `json:"rx_pwr_warn_lo" xml:"rx_pwr_warn_lo"`
	Faults                 string `json:"xmit_faults" xml:"xmit_faults"`
}

type transceiverResponseResultBodyTransceiverRow struct {
	Name         string `json:"name" xml:"name"`
	Type         string `json:"type" xml:"type"`
	Transceiver  string `json:"transceiver" xml:"transceiver"`
	IsPresent    string `json:"sfp" xml:"sfp"`
	Interface    string `json:"interface" xml:"interface"`
	PartNumber   string `json:"partnum" xml:"partnum"`
	Revision     string `json:"rev" xml:"rev"`
	SerialNumber string `json:"serialnum" xml:"serialnum"`
	// Nominal bit rate in MBits/sec
	NominalBitRate uint64 `json:"nom_bitrate" xml:"nom_bitrate"`
	// Cisco Extended ID
	CiscoExtendedID string `json:"ciscoid" xml:"ciscoid"`
	// Cisco Extended ID Number
	CiscoExtendedIDNumber uint64                                            `json:"ciscoid_1" xml:"ciscoid_1"`
	CiscoVendorID         string                                            `json:"cisco_vendor_id" xml:"cisco_vendor_id"`
	CiscoProductID        string                                            `json:"cisco_product_id" xml:"cisco_product_id"`
	CiscoPartNumber       string                                            `json:"cisco_part_number" xml:"cisco_part_number"`
	Lanes                 transceiverResponseResultBodyTransceiverLaneTable `json:"TABLE_lane" xml:"TABLE_lane"`
}

// TransceiverLaneThresholds are the alarm and warning thresholds for a
// particular Transceiver metric, i.e. temperature, voltage, current, etc.
type TransceiverLaneThresholds struct {
	Temperature struct {
		Alarm struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
		Warning struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
	}
	Voltage struct {
		Alarm struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
		Warning struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
	}

	Current struct {
		Alarm struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
		Warning struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
	}

	TxPower struct {
		Alarm struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
		Warning struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
	}

	RxPower struct {
		Alarm struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
		Warning struct {
			High float64 `json:"high" xml:"high"`
			Low  float64 `json:"low" xml:"low"`
		}
	}
}

// TransceiverLane is one of the lanes of a Transceiver.
type TransceiverLane struct {
	ID          uint64                    `json:"id" xml:"id"`
	Temperature float64                   `json:"temperature" xml:"temperature"`
	Voltage     float64                   `json:"voltage" xml:"voltage"`
	Current     float64                   `json:"current" xml:"current"`
	TxPower     float64                   `json:"tx_power" xml:"tx_power"`
	RxPower     float64                   `json:"rx_power" xml:"rx_power"`
	Errors      float64                   `json:"errors" xml:"errors"`
	Thresholds  TransceiverLaneThresholds `json:"thresholds" xml:"thresholds"`
}

// Transceiver contains system information. The information in the structure
// is from the output of "show transceiver" command.
type Transceiver struct {
	Name             string             `json:"name" xml:"name"`
	Interface        string             `json:"interface" xml:"interface"`
	PartNumber       string             `json:"part_id" xml:"part_id"`
	PartID           string             `json:"part_number" xml:"part_number"`
	Revision         string             `json:"revision" xml:"revision"`
	SerialNumber     string             `json:"serial_number" xml:"serial_number"`
	NominalBitRate   uint64             `json:"nom_bitrate" xml:"nom_bitrate"`
	ExtendedID       string             `json:"extended_id" xml:"extended_id"`
	ExtendedIDNumber string             `json:"extended_id_number" xml:"extended_id_number"`
	VendorID         string             `json:"vendor_id" xml:"vendor_id"`
	ProductID        string             `json:"product_id" xml:"product_id"`
	ProductName      string             `json:"product_name" xml:"product_name"`
	Lanes            []*TransceiverLane `json:"lanes" xml:"lanes"`
}

// NewTransceiversFromString returns Transceiver instance from an input string.
func NewTransceiversFromString(s string) ([]*Transceiver, error) {
	return NewTransceiversFromBytes([]byte(s))
}

// NewTransceiversFromBytes returns Transceiver instance from an input byte array.
func NewTransceiversFromBytes(s []byte) ([]*Transceiver, error) {
	var transceivers []*Transceiver
	resp := &transceiverResponse{}
	err := json.Unmarshal(s, resp)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	if len(resp.Result.Body.TransceiverTable.TransceiverRow) == 0 {
		return nil, fmt.Errorf("no transceivers found")
	}
	for _, j := range resp.Result.Body.TransceiverTable.TransceiverRow {
		if j.IsPresent != "present" {
			continue
		}
		t := &Transceiver{}
		t.Lanes = []*TransceiverLane{}
		t.Interface = j.Interface
		t.PartID = j.PartNumber
		t.PartNumber = j.CiscoPartNumber
		t.Revision = j.Revision
		t.SerialNumber = j.SerialNumber
		t.NominalBitRate = j.NominalBitRate
		t.ExtendedID = j.CiscoExtendedID
		t.ExtendedIDNumber = strconv.FormatUint(j.CiscoExtendedIDNumber, 10)
		t.VendorID = j.CiscoVendorID
		t.Name = j.Name
		t.ProductID = j.CiscoProductID
		t.ProductName = j.Type
		var tc uint64
		for _, l := range j.Lanes.TransceiverLaneRow {
			lane := &TransceiverLane{}
			if l.LaneNumber == 0 {
				lane.ID = tc
			} else {
				lane.ID = l.LaneNumber
			}
			tc++
			if i, err := strconv.ParseFloat(l.Temperature, 64); err == nil {
				lane.Temperature = i
			}
			if i, err := strconv.ParseFloat(l.Voltage, 64); err == nil {
				lane.Voltage = i
			}
			if i, err := strconv.ParseFloat(l.Current, 64); err == nil {
				lane.Current = i
			}
			if i, err := strconv.ParseFloat(l.TxPower, 64); err == nil {
				lane.TxPower = i
			}
			if i, err := strconv.ParseFloat(l.RxPower, 64); err == nil {
				lane.RxPower = i
			}
			if i, err := strconv.ParseFloat(l.Faults, 64); err == nil {
				lane.Errors = i
			}
			// thresholds
			if i, err := strconv.ParseFloat(l.TemperatureAlarmHigh, 64); err == nil {
				lane.Thresholds.Temperature.Alarm.High = i
			}
			if i, err := strconv.ParseFloat(l.TemperatureAlarmLow, 64); err == nil {
				lane.Thresholds.Temperature.Alarm.Low = i
			}
			if i, err := strconv.ParseFloat(l.TemperatureWarningHigh, 64); err == nil {
				lane.Thresholds.Temperature.Warning.High = i
			}
			if i, err := strconv.ParseFloat(l.TemperatureWarningLow, 64); err == nil {
				lane.Thresholds.Temperature.Warning.Low = i
			}
			if i, err := strconv.ParseFloat(l.VoltageAlarmHigh, 64); err == nil {
				lane.Thresholds.Voltage.Alarm.High = i
			}
			if i, err := strconv.ParseFloat(l.VoltageAlarmLow, 64); err == nil {
				lane.Thresholds.Voltage.Alarm.Low = i
			}
			if i, err := strconv.ParseFloat(l.VoltageWarningHigh, 64); err == nil {
				lane.Thresholds.Voltage.Warning.High = i
			}
			if i, err := strconv.ParseFloat(l.VoltageWarningLow, 64); err == nil {
				lane.Thresholds.Voltage.Warning.Low = i
			}
			if i, err := strconv.ParseFloat(l.CurrentAlarmHigh, 64); err == nil {
				lane.Thresholds.Current.Alarm.High = i
			}
			if i, err := strconv.ParseFloat(l.CurrentAlarmLow, 64); err == nil {
				lane.Thresholds.Current.Alarm.Low = i
			}
			if i, err := strconv.ParseFloat(l.CurrentWarningHigh, 64); err == nil {
				lane.Thresholds.Current.Warning.High = i
			}
			if i, err := strconv.ParseFloat(l.CurrentWarningLow, 64); err == nil {
				lane.Thresholds.Current.Warning.Low = i
			}
			if i, err := strconv.ParseFloat(l.TxPowerAlarmHigh, 64); err == nil {
				lane.Thresholds.TxPower.Alarm.High = i
			}
			if i, err := strconv.ParseFloat(l.TxPowerAlarmLow, 64); err == nil {
				lane.Thresholds.TxPower.Alarm.Low = i
			}
			if i, err := strconv.ParseFloat(l.TxPowerWarningHigh, 64); err == nil {
				lane.Thresholds.TxPower.Warning.High = i
			}
			if i, err := strconv.ParseFloat(l.TxPowerWarningLow, 64); err == nil {
				lane.Thresholds.TxPower.Warning.Low = i
			}
			if i, err := strconv.ParseFloat(l.RxPowerAlarmHigh, 64); err == nil {
				lane.Thresholds.RxPower.Alarm.High = i
			}
			if i, err := strconv.ParseFloat(l.RxPowerAlarmLow, 64); err == nil {
				lane.Thresholds.RxPower.Alarm.Low = i
			}
			if i, err := strconv.ParseFloat(l.RxPowerWarningHigh, 64); err == nil {
				lane.Thresholds.RxPower.Warning.High = i
			}
			if i, err := strconv.ParseFloat(l.RxPowerWarningLow, 64); err == nil {
				lane.Thresholds.RxPower.Warning.Low = i
			}
			t.Lanes = append(t.Lanes, lane)
		}
		transceivers = append(transceivers, t)
	}
	return transceivers, nil
}

// String returns string representation of Transceiver.
func (t *Transceiver) String() string {
	var out strings.Builder
	out.WriteString(
		fmt.Sprintf(
			"interface: %s, name: %s, serial: %s, part_number: %s, part_id: %s, nominal_bit_rate: %d",
			t.Interface,
			t.Name,
			t.SerialNumber,
			t.PartNumber,
			t.PartID,
			t.NominalBitRate,
		),
	)
	for _, l := range t.Lanes {
		out.WriteString(
			fmt.Sprintf("\ninterface: %s, lane: %d, measurements: %0.2f,%0.2f,%0.2f,%0.2f,%0.2f,%.f",
				t.Interface,
				l.ID,
				l.Temperature,
				l.Voltage,
				l.Current,
				l.TxPower,
				l.RxPower,
				l.Errors,
			),
		)
	}
	out.WriteString("\n")
	return out.String()
}
