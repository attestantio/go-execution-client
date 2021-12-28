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
	"math/big"
	"strconv"

	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// Transaction is a struct that covers all transaction types.
type Transaction struct {
	Type                 uint64
	BlockHash            types.Hash
	BlockIndex           uint32
	ChainID              uint64
	BlockNumber          uint32
	From                 types.Address
	Gas                  uint32
	GasPrice             uint64
	MaxFeePerGas         uint64
	MaxPriorityFeePerGas uint64
	Hash                 types.Hash
	Input                []byte
	Nonce                uint64
	R                    *big.Int
	S                    *big.Int
	To                   *types.Address
	TransactionIndex     uint32
	V                    *big.Int
	Value                *big.Int
	AccessList           []*AccessListEntry
}

// transactionJSON is the spec representation of the struct.
type transactionJSON struct {
	AccessList           []*AccessListEntry `json:"accessList"`
	BlockHash            string             `json:"blockHash"`
	BlockNumber          string             `json:"blockNumber"`
	ChainID              string             `json:"chainId"`
	From                 string             `json:"from"`
	Gas                  string             `json:"gas"`
	GasPrice             string             `json:"gasPrice"`
	Hash                 string             `json:"hash"`
	Input                string             `json:"input"`
	MaxFeePerGas         string             `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string             `json:"maxPriorityFeePerGas"`
	Nonce                string             `json:"nonce"`
	R                    string             `json:"r"`
	S                    string             `json:"s"`
	To                   string             `json:"to"`
	TransactionIndex     string             `json:"transactionIndex"`
	Type                 string             `json:"type"`
	V                    string             `json:"v"`
	Value                string             `json:"value"`
}

// transactionYAML is the spec representation of the struct.
type transactionYAML struct {
	AccessList           []*AccessListEntry `yaml:"accessList"`
	BlockHash            string             `yaml:"blockHash"`
	BlockNumber          uint32             `yaml:"blockNumber"`
	ChainID              uint64             `yaml:"chainId"`
	From                 string             `yaml:"from"`
	Gas                  uint32             `yaml:"gas,omitempty"`
	GasPrice             uint64             `yaml:"gasPrice"`
	Hash                 string             `yaml:"hash"`
	Input                string             `yaml:"input"`
	MaxFeePerGas         uint64             `yaml:"maxFeePerGas"`
	MaxPriorityFeePerGas uint64             `yaml:"maxPriorityFeePerGas"`
	Nonce                uint64             `yaml:"nonce"`
	R                    *big.Int           `yaml:"r"`
	S                    *big.Int           `yaml:"s"`
	To                   string             `yaml:"to"`
	TransactionIndex     uint32             `yaml:"transactionIndex"`
	Type                 uint64             `yaml:"type"`
	V                    *big.Int           `yaml:"v"`
	Value                *big.Int           `yaml:"value"`
}

// MarshalJSON implements json.Marshaler.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	switch t.Type {
	case 0:
		return t.MarshalType0JSON()
	case 1:
		return t.MarshalType1JSON()
	case 2:
		return t.MarshalType2JSON()
	default:
		return nil, errors.New("unsupported type")
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	var data transactionJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

// nolint:gocyclo
func (t *Transaction) unpack(data *transactionJSON) error {
	var err error
	var success bool

	if data.Type == "" {
		return errors.New("type missing")
	}
	t.Type, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Type), 16, 64)
	if err != nil {
		return errors.Wrap(err, "type invalid")
	}
	if t.Type != 0 && t.Type != 1 && t.Type != 2 {
		return fmt.Errorf("unhandled transaction type %s", data.Type)
	}

	if t.Type == 1 || t.Type == 2 {
		t.AccessList = data.AccessList
		if t.AccessList == nil {
			t.AccessList = make([]*AccessListEntry, 0)
		}
	}

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

	if t.Type == 1 || t.Type == 2 {
		if data.ChainID == "" {
			return errors.New("chain id missing")
		}
		t.ChainID, err = strconv.ParseUint(util.PreUnmarshalHexString(data.ChainID), 16, 64)
		if err != nil {
			return errors.Wrap(err, "chain id invalid")
		}
	}

	if data.From == "" {
		return errors.New("from missing")
	}
	address, err := hex.DecodeString(util.PreUnmarshalHexString(data.From))
	if err != nil {
		return errors.Wrap(err, "from invalid")
	}
	copy(t.From[:], address)

	if data.Gas == "" {
		return errors.New("gas missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Gas), 16, 32)
	if err != nil {
		return errors.Wrap(err, "gas invalid")
	}
	t.Gas = uint32(tmp)

	if data.GasPrice == "" {
		return errors.New("gas price missing")
	}
	t.GasPrice, err = strconv.ParseUint(util.PreUnmarshalHexString(data.GasPrice), 16, 64)
	if err != nil {
		return errors.Wrap(err, "gas price invalid")
	}

	if data.Hash == "" {
		return errors.New("hash missing")
	}
	hash, err = hex.DecodeString(util.PreUnmarshalHexString(data.Hash))
	if err != nil {
		return errors.Wrap(err, "hash invalid")
	}
	copy(t.Hash[:], hash)

	t.Input, err = hex.DecodeString(util.PreUnmarshalHexString(data.Input))
	if err != nil {
		return errors.Wrap(err, "input invalid")
	}

	if t.Type == 2 {
		if data.MaxFeePerGas == "" {
			return errors.New("max fee per gas missing")
		}
		t.MaxFeePerGas, err = strconv.ParseUint(util.PreUnmarshalHexString(data.MaxFeePerGas), 16, 64)
		if err != nil {
			return errors.Wrap(err, "max fee per gas invalid")
		}

		if data.MaxPriorityFeePerGas == "" {
			return errors.New("max priority fee per gas missing")
		}
		t.MaxPriorityFeePerGas, err = strconv.ParseUint(util.PreUnmarshalHexString(data.MaxPriorityFeePerGas), 16, 64)
		if err != nil {
			return errors.Wrap(err, "max priority fee per gas invalid")
		}
	}

	if data.Nonce == "" {
		return errors.New("nonce missing")
	}
	t.Nonce, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Nonce), 16, 64)
	if err != nil {
		return errors.Wrap(err, "nonce invalid")
	}

	if data.To != "" {
		address, err = hex.DecodeString(util.PreUnmarshalHexString(data.To))
		if err != nil {
			return errors.Wrap(err, "to invalid")
		}
		var to types.Address
		copy(to[:], address)
		t.To = &to
	}

	if data.TransactionIndex == "" {
		return errors.New("transaction index missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.TransactionIndex), 16, 32)
	if err != nil {
		return errors.Wrap(err, "transaction index invalid")
	}
	t.TransactionIndex = uint32(tmp)

	if data.Value == "" {
		return errors.New("value missing")
	}
	t.Value, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.Value), 16)
	if !success {
		return errors.New("value invalid")
	}

	if data.V == "" {
		return errors.New("v missing")
	}
	t.V, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.V), 16)
	if !success {
		return errors.New("v invalid")
	}

	if data.R == "" {
		return errors.New("r missing")
	}
	t.R, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.R), 16)
	if !success {
		return errors.New("r invalid")
	}

	if data.S == "" {
		return errors.New("s missing")
	}
	t.S, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.S), 16)
	if !success {
		return errors.New("s invalid")
	}

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (t *Transaction) MarshalYAML() ([]byte, error) {
	to := ""
	if t.To != nil {
		to = util.MarshalByteArray(t.To[:])
	}
	yamlBytes, err := yaml.MarshalWithOptions(&transactionYAML{
		AccessList:           t.AccessList,
		BlockHash:            fmt.Sprintf("%#x", t.BlockHash),
		BlockNumber:          t.BlockNumber,
		ChainID:              t.ChainID,
		From:                 fmt.Sprintf("%#x", t.From),
		Gas:                  t.Gas,
		GasPrice:             t.GasPrice,
		Hash:                 fmt.Sprintf("%#x", t.Hash),
		Input:                fmt.Sprintf("%#x", t.Input),
		Nonce:                t.Nonce,
		MaxFeePerGas:         t.MaxFeePerGas,
		MaxPriorityFeePerGas: t.MaxPriorityFeePerGas,
		R:                    t.R,
		S:                    t.S,
		To:                   to,
		Type:                 t.Type,
		TransactionIndex:     t.TransactionIndex,
		V:                    t.V,
		Value:                t.Value,
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *Transaction) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var data transactionJSON
	if err := yaml.Unmarshal(input, &data); err != nil {
		return err
	}
	return t.unpack(&data)
}

// MarshalRLP returns an RLP representation of the transaction.
func (t *Transaction) MarshalRLP() ([]byte, error) {
	switch t.Type {
	case 0:
		return t.MarshalType0RLP()
	case 1:
		return t.MarshalType1RLP()
	case 2:
		return t.MarshalType2RLP()
	default:
		return nil, errors.New("unsupported type")
	}
}

// String returns a string version of the structure.
func (t *Transaction) String() string {
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
