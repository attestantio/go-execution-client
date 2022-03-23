// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// LondonTransactionReceipt contains a transaction receipt.
type LondonTransactionReceipt struct {
	BlockHash         types.Hash
	BlockNumber       uint32
	ContractAddress   *types.Address
	CumulativeGasUsed uint32
	EffectiveGasPrice uint64
	From              types.Address
	GasUsed           uint32
	Logs              []*BerlinTransactionEvent
	LogsBloom         []byte
	Status            uint32
	To                *types.Address
	TransactionHash   types.Hash
	TransactionIndex  uint32
	Type              TransactionType
}

// londonTransactionReceiptJSON is the spec representation of the struct.
type londonTransactionReceiptJSON struct {
	BlockHash         string                    `json:"blockHash"`
	BlockNumber       string                    `json:"blockNumber"`
	ContractAddress   *string                   `json:"contractAddress"`
	CumulativeGasUsed string                    `json:"cumulativeGasUsed"`
	EffectiveGasPrice string                    `json:"effectiveGasPrice"`
	From              string                    `json:"from"`
	GasUsed           string                    `json:"gasUsed"`
	Logs              []*BerlinTransactionEvent `json:"logs"`
	LogsBloom         string                    `json:"logsBloom"`
	Status            string                    `json:"status"`
	To                *string                   `json:"to"`
	TransactionHash   string                    `json:"transactionHash"`
	TransactionIndex  string                    `json:"transactionIndex"`
	Type              TransactionType           `json:"type"`
}

// MarshalJSON implements json.Marshaler.
func (t *LondonTransactionReceipt) MarshalJSON() ([]byte, error) {
	var contractAddress *string
	if t.ContractAddress != nil {
		tmp := util.MarshalNullableAddress((*t.ContractAddress)[:])
		contractAddress = &tmp
	}
	var to *string
	if t.To != nil {
		tmp := util.MarshalNullableAddress(t.To[:])
		to = &tmp
	}
	return json.Marshal(&londonTransactionReceiptJSON{
		BlockHash:         util.MarshalByteArray(t.BlockHash[:]),
		BlockNumber:       util.MarshalUint32(t.BlockNumber),
		ContractAddress:   contractAddress,
		CumulativeGasUsed: util.MarshalUint32(t.CumulativeGasUsed),
		EffectiveGasPrice: util.MarshalUint64(t.EffectiveGasPrice),
		From:              util.MarshalByteArray(t.From[:]),
		GasUsed:           util.MarshalUint32(t.GasUsed),
		Logs:              t.Logs,
		LogsBloom:         util.MarshalByteArray(t.LogsBloom),
		Status:            util.MarshalUint32(t.Status),
		To:                to,
		TransactionHash:   util.MarshalByteArray(t.TransactionHash[:]),
		TransactionIndex:  util.MarshalUint32(t.TransactionIndex),
		Type:              t.Type,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *LondonTransactionReceipt) UnmarshalJSON(input []byte) error {
	var data londonTransactionReceiptJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

func (t *LondonTransactionReceipt) unpack(data *londonTransactionReceiptJSON) error {
	var err error

	if data.BlockHash == "" {
		return errors.New("block hash missing")
	}
	hash, err := hex.DecodeString(util.PreUnmarshalHexString(data.BlockHash))
	if err != nil {
		return errors.Wrap(err, "block hash invalid")
	}
	copy(t.BlockHash[:], hash)

	if data.BlockNumber == "" {
		return errors.New("block number missing")
	}
	tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(data.BlockNumber), 16, 32)
	if err != nil {
		return errors.Wrap(err, "block number invalid")
	}
	t.BlockNumber = uint32(tmp)

	if data.ContractAddress != nil {
		address, err := hex.DecodeString(util.PreUnmarshalHexString(*data.ContractAddress))
		if err != nil {
			return errors.Wrap(err, "contract address invalid")
		}
		var contractAddress types.Address
		copy(contractAddress[:], address)
		t.ContractAddress = &contractAddress
	}

	if data.CumulativeGasUsed == "" {
		return errors.New("cumulative gas used missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.CumulativeGasUsed), 16, 32)
	if err != nil {
		return errors.Wrap(err, "cumulative gas used invalid")
	}
	t.CumulativeGasUsed = uint32(tmp)

	if data.EffectiveGasPrice == "" {
		return errors.New("effective gas price missing")
	}
	t.EffectiveGasPrice, err = strconv.ParseUint(util.PreUnmarshalHexString(data.EffectiveGasPrice), 16, 64)
	if err != nil {
		return errors.Wrap(err, "effective gas price invalid")
	}

	if data.From == "" {
		return errors.New("from missing")
	}
	address, err := hex.DecodeString(util.PreUnmarshalHexString(data.From))
	if err != nil {
		return errors.Wrap(err, "from invalid")
	}
	copy(t.From[:], address)

	if data.GasUsed == "" {
		return errors.New("gas used missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.GasUsed), 16, 32)
	if err != nil {
		return errors.Wrap(err, "gas used invalid")
	}
	t.GasUsed = uint32(tmp)

	t.Logs = data.Logs

	if data.LogsBloom == "" {
		return errors.New("logs bloom missing")
	}
	t.LogsBloom, err = hex.DecodeString(util.PreUnmarshalHexString(data.LogsBloom))
	if err != nil {
		return errors.Wrap(err, "logs bloom invalid")
	}

	if data.Status == "" {
		return errors.New("status missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Status), 16, 32)
	if err != nil {
		return errors.Wrap(err, "status invalid")
	}
	t.Status = uint32(tmp)

	if data.To != nil {
		address, err = hex.DecodeString(util.PreUnmarshalHexString(*data.To))
		if err != nil {
			return errors.Wrap(err, "to invalid")
		}
		var to types.Address
		copy(to[:], address)
		t.To = &to
	}

	if data.TransactionHash == "" {
		return errors.New("transaction hash missing")
	}
	hash, err = hex.DecodeString(util.PreUnmarshalHexString(data.TransactionHash))
	if err != nil {
		return errors.Wrap(err, "transaction hash invalid")
	}
	copy(t.TransactionHash[:], hash)

	if data.TransactionIndex == "" {
		return errors.New("transaction index missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.TransactionIndex), 16, 32)
	if err != nil {
		return errors.Wrap(err, "transaction index invalid")
	}
	t.TransactionIndex = uint32(tmp)

	if data.Type == TransactionTypeUnknown {
		return errors.New("transaction type unrecognised")
	}
	t.Type = data.Type

	return nil
}

// String returns a string version of the structure.
func (t *LondonTransactionReceipt) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
