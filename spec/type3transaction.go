// Copyright © 2023 Attestant Limited.
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
	"github.com/pkg/errors"
)

// Type3Transaction is a Cancun type 3 transaction.
type Type3Transaction struct {
	AccessList []*AccessListEntry
	// BlockHash is only available for transactions included in a block, so optional.
	BlockHash *types.Hash
	// BlockNumber is only available for transactions included in a block, so optional.
	BlockNumber *uint32
	ChainID     uint64
	From        types.Address
	Gas         uint32
	// GasPrice is only available for transactions included in a block, so optional.
	GasPrice             *uint64
	Hash                 types.Hash
	Input                []byte
	MaxFeePerGas         uint64
	MaxPriorityFeePerGas uint64
	Nonce                uint64
	R                    *big.Int
	S                    *big.Int
	To                   *types.Address
	// TransactionIndex is only available for transactions included in a block, so optional.
	TransactionIndex *uint32
	V                *big.Int
	Value            *big.Int
}

// type3TransactionJSON is the spec representation of a type 3 transaction.
type type3TransactionJSON struct {
	AccessList           []*AccessListEntry `json:"accessList,omitempty"`
	BlockHash            *string            `json:"blockHash,omitempty"`
	BlockNumber          *string            `json:"blockNumber,omitempty"`
	ChainID              string             `json:"chainId"`
	From                 string             `json:"from"`
	Gas                  string             `json:"gas"`
	GasPrice             *string            `json:"gasPrice,omitempty"`
	Hash                 string             `json:"hash"`
	Input                string             `json:"input"`
	MaxFeePerGas         string             `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string             `json:"maxPriorityFeePerGas"`
	Nonce                string             `json:"nonce"`
	R                    string             `json:"r"`
	S                    string             `json:"s"`
	To                   string             `json:"to"`
	TransactionIndex     *string            `json:"transactionIndex,omitempty"`
	Type                 string             `json:"type"`
	V                    string             `json:"v"`
	Value                string             `json:"value"`
}

// MarshalJSON marshals a type 3 transaction.
func (t *Type3Transaction) MarshalJSON() ([]byte, error) {
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
	var gasPrice *string
	if t.GasPrice != nil {
		tmp := util.MarshalUint64(*t.GasPrice)
		gasPrice = &tmp
	}

	return json.Marshal(&type3TransactionJSON{
		AccessList:           t.AccessList,
		BlockHash:            blockHash,
		BlockNumber:          blockNumber,
		ChainID:              util.MarshalUint64(t.ChainID),
		From:                 util.MarshalByteArray(t.From[:]),
		Gas:                  util.MarshalUint32(t.Gas),
		GasPrice:             gasPrice,
		Hash:                 util.MarshalByteArray(t.Hash[:]),
		Input:                util.MarshalByteArray(t.Input),
		MaxFeePerGas:         util.MarshalUint64(t.MaxFeePerGas),
		MaxPriorityFeePerGas: util.MarshalUint64(t.MaxPriorityFeePerGas),
		Nonce:                util.MarshalUint64(t.Nonce),
		R:                    util.MarshalBigInt(t.R),
		S:                    util.MarshalBigInt(t.S),
		To:                   to,
		Type:                 "0x2",
		TransactionIndex:     transactionIndex,
		V:                    util.MarshalBigInt(t.V),
		Value:                util.MarshalBigInt(t.Value),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Type3Transaction) UnmarshalJSON(input []byte) error {
	var data type3TransactionJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

// nolint:gocyclo
func (t *Type3Transaction) unpack(data *type3TransactionJSON) error {
	var err error
	var success bool

	// Guard to ensure we are unpacking the correct transaction type.
	if data.Type == "" {
		return errors.New("type missing for type 3 transaction")
	}
	if data.Type != "0x3" {
		return errors.New("type incorrect")
	}

	t.AccessList = data.AccessList
	if t.AccessList == nil {
		t.AccessList = make([]*AccessListEntry, 0)
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

	if data.ChainID == "" {
		return errors.New("chain id missing")
	}
	t.ChainID, err = strconv.ParseUint(util.PreUnmarshalHexString(data.ChainID), 16, 64)
	if err != nil {
		return errors.Wrap(err, "chain id invalid")
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

	if data.GasPrice != nil {
		tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(*data.GasPrice), 16, 64)
		if err != nil {
			return errors.Wrap(err, "gas price invalid")
		}
		t.GasPrice = &tmp
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
	fmt.Printf("%v\n", data.R)
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
func (t *Type3Transaction) MarshalRLP() ([]byte, error) {
	// Create generic buffers, to allow reuse.
	bufA := bytes.NewBuffer(make([]byte, 0, 1024))
	bufB := bytes.NewBuffer(make([]byte, 0, 1024))

	// Transaction data.
	util.RLPUint64(bufA, t.ChainID)
	util.RLPUint64(bufA, t.Nonce)
	util.RLPUint64(bufA, t.MaxPriorityFeePerGas)
	util.RLPUint64(bufA, t.MaxFeePerGas)
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
	if len(t.AccessList) != 0 {
		entryBuf := bytes.NewBuffer(make([]byte, 0, 1024))
		addressBuf := bytes.NewBuffer(make([]byte, 0, 20))
		for _, accessListEntry := range t.AccessList {
			util.RLPBytes(entryBuf, accessListEntry.Address)
			for _, key := range accessListEntry.StorageKeys {
				util.RLPBytes(addressBuf, key)
			}
			util.RLPList(entryBuf, addressBuf.Bytes())
			addressBuf.Reset()
			util.RLPList(bufB, entryBuf.Bytes())
			entryBuf.Reset()
		}
	}
	util.RLPList(bufA, bufB.Bytes())
	bufB.Reset()

	// Signature.
	util.RLPBytes(bufA, t.V.Bytes())
	util.RLPBytes(bufA, t.R.Bytes())
	util.RLPBytes(bufA, t.S.Bytes())

	// EIP-2718 definition.
	bufB.WriteByte(0x02)
	util.RLPList(bufB, bufA.Bytes())
	bufA.Reset()
	util.RLPBytes(bufA, bufB.Bytes())
	return bufA.Bytes(), nil
}

// String returns a string version of the structure.
func (t *Type3Transaction) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
