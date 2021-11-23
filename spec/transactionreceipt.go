// Copyright © 2021 Attestant Limited.
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

	"github.com/attestantio/go-execution-client/util"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// TransactionReceipt contains a transaction receipt.
type TransactionReceipt struct {
	BlockHash         Hash
	BlockNumber       uint32
	ContractAddress   *Address
	CumulativeGasUsed uint32
	From              Address
	GasUsed           uint32
	Logs              []*TransactionEvent
	LogsBloom         []byte
	Status            uint32
	To                *Address
	TransactionHash   Hash
	TransactionIndex  uint32
	Type              uint32
}

// transactionReceiptJSON is the spec representation of the struct.
type transactionReceiptJSON struct {
	BlockHash         string              `json:"blockHash"`
	BlockNumber       string              `json:"blockNumber"`
	ContractAddress   string              `json:"contractAddress,omitempty"`
	CumulativeGasUsed string              `json:"cumulativeGasUsed"`
	From              string              `json:"from"`
	GasUsed           string              `json:"gasUsed"`
	Logs              []*TransactionEvent `json:"logs"`
	LogsBloom         string              `json:"logsBloom"`
	Status            string              `json:"status"`
	To                string              `json:"to"`
	TransactionHash   string              `json:"transactionHash"`
	TransactionIndex  string              `json:"transactionIndex"`
	Type              string              `json:"type"`
}

// transactionReceiptYAML is the spec representation of the struct.
type transactionReceiptYAML struct {
	BlockHash         string              `yaml:"blockHash"`
	BlockNumber       uint32              `yaml:"blockNumber"`
	ContractAddress   string              `yaml:"contractAddress,omitempty"`
	CumulativeGasUsed uint32              `yaml:"cumulativeGasUsed"`
	From              string              `yaml:"from"`
	GasUsed           uint32              `yaml:"gasUsed"`
	Logs              []*TransactionEvent `yaml:"logs"`
	LogsBloom         string              `yaml:"logsBloom"`
	Status            uint32              `yaml:"status"`
	To                string              `yaml:"to"`
	TransactionHash   string              `yaml:"transactionHash"`
	TransactionIndex  uint32              `yaml:"transactionNumber"`
	Type              uint32              `yaml:"type"`
}

// MarshalJSON implements json.Marshaler.
func (t *TransactionReceipt) MarshalJSON() ([]byte, error) {
	contractAddress := ""
	if t.ContractAddress != nil {
		contractAddress = util.MarshalNullableAddress((*t.ContractAddress)[:])
	}
	to := ""
	if t.To != nil {
		to = util.MarshalNullableAddress(t.To[:])
	}
	return json.Marshal(&transactionReceiptJSON{
		BlockHash:         util.MarshalByteArray(t.BlockHash[:]),
		BlockNumber:       util.MarshalUint32(t.BlockNumber),
		ContractAddress:   contractAddress,
		CumulativeGasUsed: util.MarshalUint32(t.CumulativeGasUsed),
		From:              util.MarshalByteArray(t.From[:]),
		GasUsed:           util.MarshalUint32(t.GasUsed),
		Logs:              t.Logs,
		LogsBloom:         util.MarshalByteArray(t.LogsBloom),
		Status:            util.MarshalUint32(t.Status),
		To:                to,
		TransactionHash:   util.MarshalByteArray(t.TransactionHash[:]),
		TransactionIndex:  util.MarshalUint32(t.TransactionIndex),
		Type:              util.MarshalUint32(t.Type),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionReceipt) UnmarshalJSON(input []byte) error {
	var data transactionReceiptJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

func (t *TransactionReceipt) unpack(data *transactionReceiptJSON) error {
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

	if data.ContractAddress != "" {
		address, err := hex.DecodeString(util.PreUnmarshalHexString(data.ContractAddress))
		if err != nil {
			return errors.Wrap(err, "contract address invalid")
		}
		var contractAddress Address
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

	if data.To != "" {
		address, err = hex.DecodeString(util.PreUnmarshalHexString(data.To))
		if err != nil {
			return errors.Wrap(err, "to invalid")
		}
		var to Address
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

	if data.Type == "" {
		return errors.New("type missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Type), 16, 32)
	if err != nil {
		return errors.Wrap(err, "type invalid")
	}
	t.Type = uint32(tmp)

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (t *TransactionReceipt) MarshalYAML() ([]byte, error) {
	contractAddress := ""
	if t.ContractAddress != nil {
		contractAddress = util.MarshalNullableAddress((*t.ContractAddress)[:])
	}
	to := ""
	if t.To != nil {
		to = util.MarshalNullableAddress(t.To[:])
	}
	yamlBytes, err := yaml.MarshalWithOptions(&transactionReceiptYAML{
		BlockHash:         fmt.Sprintf("%#x", t.BlockHash),
		BlockNumber:       t.BlockNumber,
		ContractAddress:   contractAddress,
		CumulativeGasUsed: t.CumulativeGasUsed,
		From:              fmt.Sprintf("%#x", t.From),
		GasUsed:           t.GasUsed,
		Logs:              t.Logs,
		LogsBloom:         fmt.Sprintf("%#x", t.LogsBloom),
		Status:            t.Status,
		To:                to,
		TransactionHash:   fmt.Sprintf("%#x", t.TransactionHash),
		TransactionIndex:  t.TransactionIndex,
		Type:              t.Type,
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *TransactionReceipt) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var transactionReceiptJSON transactionReceiptJSON
	if err := yaml.Unmarshal(input, &transactionReceiptJSON); err != nil {
		return err
	}
	return t.unpack(&transactionReceiptJSON)
}

// String returns a string version of the structure.
func (t *TransactionReceipt) String() string {
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
