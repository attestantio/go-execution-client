// Copyright © 2021, 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
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
	"github.com/pkg/errors"
)

// Type0Transaction is the spec representation of a type 0 transaction.
type Type0Transaction struct {
	BlockHash        *types.Hash
	BlockNumber      *uint32
	From             types.Address
	Gas              uint32
	GasPrice         uint64
	Hash             types.Hash
	Input            []byte
	Nonce            uint64
	R                *big.Int
	S                *big.Int
	To               *types.Address
	TransactionIndex *uint32
	V                *big.Int
	Value            *big.Int
}

// type0TransactionJSON is the spec representation of a type 0 transaction.
type type0TransactionJSON struct {
	BlockHash        *string `json:"blockHash"`
	BlockNumber      *string `json:"blockNumber"`
	From             string  `json:"from"`
	Gas              string  `json:"gas"`
	GasPrice         string  `json:"gasPrice"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            string  `json:"nonce"`
	R                string  `json:"r"`
	S                string  `json:"s"`
	To               string  `json:"to"`
	TransactionIndex *string `json:"transactionIndex"`
	Type             string  `json:"type"`
	V                string  `json:"v"`
	Value            string  `json:"value"`
}

// MarshalJSON marshals a type 0 transaction.
func (t *Type0Transaction) MarshalJSON() ([]byte, error) {
	var blockHash *string
	if t.BlockHash != nil {
		tmp := fmt.Sprintf("%#x", *t.BlockHash)
		blockHash = &tmp
	}
	var blockNumber *string
	if t.BlockNumber != nil {
		tmp := util.MarshalUint32(*t.BlockNumber)
		blockNumber = &tmp
	}
	to := ""
	if t.To != nil {
		to = util.MarshalByteArray(t.To[:])
	}
	var transactionIndex *string
	if t.TransactionIndex != nil {
		tmp := util.MarshalUint32(*t.TransactionIndex)
		transactionIndex = &tmp
	}
	return json.Marshal(&type0TransactionJSON{
		BlockHash:        blockHash,
		BlockNumber:      blockNumber,
		From:             util.MarshalByteArray(t.From[:]),
		Gas:              util.MarshalUint32(t.Gas),
		GasPrice:         util.MarshalUint64(t.GasPrice),
		Hash:             util.MarshalByteArray(t.Hash[:]),
		Input:            util.MarshalByteArray(t.Input),
		Nonce:            util.MarshalUint64(t.Nonce),
		R:                util.MarshalBigInt(t.R),
		S:                util.MarshalBigInt(t.S),
		To:               to,
		TransactionIndex: transactionIndex,
		Type:             "0x0",
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

// nolint:gocyclo
func (t *Type0Transaction) unpack(data *type0TransactionJSON) error {
	var err error
	var success bool

	// Guard to ensure we are unpacking the correct transaction type.
	if data.Type == "" {
		// If type is empty it implies a type 0 transaction.
		data.Type = "0x0"
	}
	if data.Type != "0x0" {
		return errors.New("type incorrect")
	}

	if data.BlockHash != nil {
		hash, err := hex.DecodeString(util.PreUnmarshalHexString(*data.BlockHash))
		if err != nil {
			return errors.Wrap(err, "block hash invalid")
		}
		blockHash := types.Hash{}
		copy(blockHash[:], hash)
		t.BlockHash = &blockHash
		if data.BlockNumber == nil {
			return errors.New("block number missing")
		}
		tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(*data.BlockNumber), 16, 32)
		if err != nil {
			return errors.Wrap(err, "block number invalid")
		}
		blockNumber := uint32(tmp)
		t.BlockNumber = &blockNumber
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
	tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(data.Gas), 16, 32)
	if err != nil {
		return errors.Wrap(err, "gas invalid")
	}
	t.Gas = uint32(tmp)

	if data.GasPrice == "" {
		return errors.New("gas price missing")
	}
	t.GasPrice, err = strconv.ParseUint(util.PreUnmarshalHexString(data.GasPrice), 16, 64)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("gas price for %s invalid", data.Hash))
	}

	if data.Hash == "" {
		return errors.New("hash missing")
	}
	hash, err := hex.DecodeString(util.PreUnmarshalHexString(data.Hash))
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
		var to types.Address
		copy(to[:], address)
		t.To = &to
	}

	if data.TransactionIndex != nil {
		tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(*data.TransactionIndex), 16, 32)
		if err != nil {
			return errors.Wrap(err, "transaction index invalid")
		}
		transactionIndex := uint32(tmp)
		t.TransactionIndex = &transactionIndex
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

// MarshalRLP returns an RLP representation of the transaction.
func (t *Type0Transaction) MarshalRLP() ([]byte, error) {
	// Create generic buffers, to allow reuse.
	bufA := bytes.NewBuffer(make([]byte, 0, 1024))
	bufB := bytes.NewBuffer(make([]byte, 0, 1024))

	// Transaction data.
	util.RLPUint64(bufA, t.Nonce)
	util.RLPUint64(bufA, t.GasPrice)
	util.RLPUint64(bufA, uint64(t.Gas))
	if t.To != nil {
		util.RLPAddress(bufA, *t.To)
	} else {
		util.RLPNil(bufA)
	}
	if t.Value != nil {
		util.RLPBytes(bufA, t.Value.Bytes())
	} else {
		util.RLPNil(bufA)
	}
	util.RLPBytes(bufA, t.Input)

	// Signature.
	util.RLPBytes(bufA, t.V.Bytes())
	util.RLPBytes(bufA, t.R.Bytes())
	util.RLPBytes(bufA, t.S.Bytes())

	util.RLPList(bufB, bufA.Bytes())
	return bufB.Bytes(), nil
}

// String returns a string version of the structure.
func (t *Type0Transaction) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
