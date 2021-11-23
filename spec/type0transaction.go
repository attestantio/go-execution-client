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

	"github.com/attestantio/go-execution-client/util"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// Type0Transaction contains a type 0 transaction,
// also known as a legacy transaction.
type Type0Transaction struct {
	BlockHash        Hash
	BlockNumber      uint32
	From             Address
	Gas              uint32
	GasPrice         uint64
	Hash             Hash
	Input            []byte
	Nonce            uint64
	R                *big.Int
	S                *big.Int
	To               *Address
	TransactionIndex *big.Int
	V                *big.Int
	Value            *big.Int
}

// type0TransactionJSON is the spec representation of the struct.
type type0TransactionJSON struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	R                string `json:"r"`
	S                string `json:"s"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Type             string `json:"type"`
	V                string `json:"v"`
	Value            string `json:"value"`
}

// type0TransactionYAML is the spec representation of the struct.
type type0TransactionYAML struct {
	BlockHash        string   `yaml:"blockHash"`
	BlockNumber      uint32   `yaml:"blockNumber"`
	From             string   `yaml:"from"`
	Gas              uint32   `yaml:"gas"`
	GasPrice         uint64   `yaml:"gasPrice"`
	Hash             string   `yaml:"hash"`
	Input            string   `yaml:"input"`
	Nonce            uint64   `yaml:"nonce"`
	R                *big.Int `yaml:"r"`
	S                *big.Int `yaml:"s"`
	To               string   `yaml:"to"`
	TransactionIndex *big.Int `yaml:"transactionIndex"`
	Type             uint32   `yaml:"type"`
	V                *big.Int `yaml:"v"`
	Value            *big.Int `yaml:"value"`
}

// MarshalJSON implements json.Marshaler.
func (t *Type0Transaction) MarshalJSON() ([]byte, error) {
	to := ""
	if t.To != nil {
		to = util.MarshalByteArray(t.To[:])
	}
	return json.Marshal(&type0TransactionJSON{
		BlockHash:        util.MarshalByteArray(t.BlockHash[:]),
		BlockNumber:      util.MarshalUint32(t.BlockNumber),
		From:             util.MarshalByteArray(t.From[:]),
		Gas:              util.MarshalUint32(t.Gas),
		GasPrice:         util.MarshalUint64(t.GasPrice),
		Hash:             util.MarshalByteArray(t.Hash[:]),
		Input:            util.MarshalByteArray(t.Input),
		Nonce:            util.MarshalUint64(t.Nonce),
		R:                util.MarshalBigInt(t.R),
		S:                util.MarshalBigInt(t.S),
		To:               to,
		Type:             "0x0",
		TransactionIndex: util.MarshalBigInt(t.TransactionIndex),
		V:                util.MarshalBigInt(t.V),
		Value:            util.MarshalBigInt(t.Value),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Type0Transaction) UnmarshalJSON(input []byte) error {
	var data type0TransactionJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

func (t *Type0Transaction) unpack(data *type0TransactionJSON) error {
	var err error
	var success bool

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
		var to Address
		copy(to[:], address)
		t.To = &to
	}

	if data.TransactionIndex == "" {
		return errors.New("transaction index missing")
	}
	t.TransactionIndex, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.TransactionIndex), 16)
	if !success {
		return errors.New("transaction index invalid")
	}

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
func (t *Type0Transaction) MarshalYAML() ([]byte, error) {
	to := ""
	if t.To != nil {
		to = util.MarshalByteArray(t.To[:])
	}
	yamlBytes, err := yaml.MarshalWithOptions(&type0TransactionYAML{
		BlockHash:        fmt.Sprintf("%#x", t.BlockHash),
		BlockNumber:      t.BlockNumber,
		From:             fmt.Sprintf("%#x", t.From),
		Gas:              t.Gas,
		GasPrice:         t.GasPrice,
		Hash:             fmt.Sprintf("%#x", t.Hash),
		Input:            fmt.Sprintf("%#x", t.Input),
		Nonce:            t.Nonce,
		R:                t.R,
		S:                t.S,
		To:               to,
		Type:             0,
		TransactionIndex: t.TransactionIndex,
		V:                t.V,
		Value:            t.Value,
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *Type0Transaction) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var type0TransactionJSON type0TransactionJSON
	if err := yaml.Unmarshal(input, &type0TransactionJSON); err != nil {
		return err
	}
	return t.unpack(&type0TransactionJSON)
}

// String returns a string version of the structure.
func (t *Type0Transaction) String() string {
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}

// Type returns the type of this transaction.
func (t *Type0Transaction) Type() uint64 {
	return 0
}
