package client

import (
	"encoding/json"
	"fmt"
)

type errorCountersResponse struct {
	ID      uint64                      `json:"id" xml:"id"`
	Version string                      `json:"jsonrpc" xml:"jsonrpc"`
	Result  errorCountersResponseResult `json:"result" xml:"result"`
}

type errorCountersResponseResult struct {
	Body errorCountersResponseResultBody `json:"body" xml:"body"`
}

type errorCountersResponseResultBody struct {
	ErrorCountersTable errorCountersResponseResultBodyErrorCountersTable `json:"TABLE_interface" xml:"TABLE_interface"`
}

type errorCountersResponseResultBodyErrorCountersTable struct {
	ErrorCountersRow []errorCountersResponseResultBodyErrorCountersRow `json:"ROW_interface" xml:"ROW_interface"`
}

type errorCountersResponseResultBodyErrorCountersRow struct {
	Interface      string `json:"interface" xml:"interface"`
	EthAlignErr    uint64 `json:"eth_align_err" xml:"eth_align_err"`
	EthFCSErr      uint64 `json:"eth_fcs_err" xml:"eth_fcs_err"`
	EthOutDiscards uint64 `json:"eth_outdisc" xml:"eth_outdisc"`
	EthRcvErr      uint64 `json:"eth_rcv_err" xml:"eth_rcv_err"`
	EthUndersize   uint64 `json:"eth_undersize" xml:"eth_undersize"`
	EthXmitErr     uint64 `json:"eth_xmit_err" xml:"eth_xmit_err"`
	EthCarriSen    uint64 `json:"eth_carri_sen" xml:"eth_carri_sen"`
	EthExcessCol   uint64 `json:"eth_excess_col" xml:"eth_excess_col"`
	EthLateCol     uint64 `json:"eth_late_col" xml:"eth_late_col"`
	EthMultiCol    uint64 `json:"eth_multi_col" xml:"eth_multi_col"`
	EthRunts       uint64 `json:"eth_runts" xml:"eth_runts"`
	EthSingleCol   uint64 `json:"eth_single_col" xml:"eth_single_col"`
	EthDeferredTx  uint64 `json:"eth_deferred_tx" xml:"eth_deferred_tx"`
	EthGiants      uint64 `json:"eth_giants" xml:"eth_giants"`
	EthInMacRxErr  uint64 `json:"eth_inmacrx_err" xml:"eth_inmacrx_err"`
	EthInMacTxErr  uint64 `json:"eth_inmactx_err" xml:"eth_inmactx_err"`
	EthSymbolErr   uint64 `json:"eth_symbol_err" xml:"eth_symbol_err"`
	EthInDiscards  uint64 `json:"eth_indisc" xml:"eth_indisc"`
}

type ErrorCounters struct {
	Interface      string `json:"interface" xml:"interface"`
	EthAlignErr    uint64 `json:"eth_align_err" xml:"eth_align_err"`
	EthCarriSen    uint64 `json:"eth_carri_sen" xml:"eth_carri_sen"`
	EthDeferredTx  uint64 `json:"eth_deferred_tx" xml:"eth_deferred_tx"`
	EthExcessCol   uint64 `json:"eth_excess_col" xml:"eth_excess_col"`
	EthFCSErr      uint64 `json:"eth_fcs_err" xml:"eth_fcs_err"`
	EthGiants      uint64 `json:"eth_giants" xml:"eth_giants"`
	EthInDiscards  uint64 `json:"eth_indisc" xml:"eth_indisc"`
	EthInMacRxErr  uint64 `json:"eth_inmacrx_err" xml:"eth_inmacrx_err"`
	EthInMacTxErr  uint64 `json:"eth_inmactx_err" xml:"eth_inmactx_err"`
	EthLateCol     uint64 `json:"eth_late_col" xml:"eth_late_col"`
	EthMultiCol    uint64 `json:"eth_multi_col" xml:"eth_multi_col"`
	EthOutDiscards uint64 `json:"eth_outdisc" xml:"eth_outdisc"`
	EthRcvErr      uint64 `json:"eth_rcv_err" xml:"eth_rcv_err"`
	EthRunts       uint64 `json:"eth_runts" xml:"eth_runts"`
	EthSingleCol   uint64 `json:"eth_single_col" xml:"eth_single_col"`
	EthSymbolErr   uint64 `json:"eth_symbol_err" xml:"eth_symbol_err"`
	EthUndersize   uint64 `json:"eth_undersize" xml:"eth_undersize"`
	EthXmitErr     uint64 `json:"eth_xmit_err" xml:"eth_xmit_err"`
}

// NewErrorCountersFromString returns ErrorCounter instance from an input string.
func NewErrorCountersFromString(s string) ([]*ErrorCounters, error) {
	return NewErrorCountersFromBytes([]byte(s))
}

// NewErrorCountersFromBytes returns ErrorCounter instance from an input byte array.
func NewErrorCountersFromBytes(s []byte) ([]*ErrorCounters, error) {
	var errorCounters []*ErrorCounters
	eCountersResponse := &errorCountersResponse{}
	err := json.Unmarshal(s, eCountersResponse)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %s, server response: %s", err, string(s[:]))
	}
	if len(eCountersResponse.Result.Body.ErrorCountersTable.ErrorCountersRow) < 1 {
		return nil, fmt.Errorf("Error parsing the received response: %s", s)
	}
	i, size := 0, len(eCountersResponse.Result.Body.ErrorCountersTable.ErrorCountersRow)
	for _, e := range eCountersResponse.Result.Body.ErrorCountersTable.ErrorCountersRow {
		if i < size/4 {
			errorCounter := &ErrorCounters{}
			errorCounter.Interface = e.Interface
			errorCounter.EthAlignErr = e.EthAlignErr
			errorCounter.EthFCSErr = e.EthFCSErr
			errorCounter.EthOutDiscards = e.EthOutDiscards
			errorCounter.EthRcvErr = e.EthRcvErr
			errorCounter.EthUndersize = e.EthUndersize
			errorCounter.EthXmitErr = e.EthXmitErr
			errorCounters = append(errorCounters, errorCounter)
		} else if i < size/2 {
			errorCounter := errorCounters[i%(size/4)]
			errorCounter.EthCarriSen = e.EthCarriSen
			errorCounter.EthExcessCol = e.EthExcessCol
			errorCounter.EthLateCol = e.EthLateCol
			errorCounter.EthMultiCol = e.EthMultiCol
			errorCounter.EthRunts = e.EthRunts
			errorCounter.EthSingleCol = e.EthSingleCol
		} else if i < 3*size/4 {
			errorCounter := errorCounters[i%(size/4)]
			errorCounter.EthDeferredTx = e.EthDeferredTx
			errorCounter.EthGiants = e.EthGiants
			errorCounter.EthInMacRxErr = e.EthInMacRxErr
			errorCounter.EthInMacTxErr = e.EthInMacTxErr
			errorCounter.EthSymbolErr = e.EthSymbolErr
		} else {
			errorCounter := errorCounters[i%(size/4)]
			errorCounter.EthInDiscards = e.EthInDiscards
		}
		i += 1
	}
	return errorCounters, nil
}
