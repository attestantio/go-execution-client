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
	"encoding/json"
	"fmt"

	"github.com/attestantio/go-execution-client/util"
)

// type2TransactionJSON is the spec representation of a type 2 transaction.
type type2TransactionJSON struct {
	AccessList           []*AccessListEntry `json:"accessList"`
	BlockHash            *string            `json:"blockHash"`
	BlockNumber          *string            `json:"blockNumber"`
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
	TransactionIndex     *string            `json:"transactionIndex"`
	Type                 string             `json:"type"`
	V                    string             `json:"v"`
	Value                string             `json:"value"`
}

// MarshalType2JSON marshals a type 2 transaction.
func (t *Transaction) MarshalType2JSON() ([]byte, error) {
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
	return json.Marshal(&type2TransactionJSON{
		AccessList:           t.AccessList,
		BlockHash:            blockHash,
		BlockNumber:          blockNumber,
		ChainID:              util.MarshalUint64(t.ChainID),
		From:                 util.MarshalByteArray(t.From[:]),
		Gas:                  util.MarshalUint32(t.Gas),
		GasPrice:             util.MarshalUint64(t.GasPrice),
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

// MarshalType2RLP returns an RLP representation of the transaction.
func (t *Transaction) MarshalType2RLP() ([]byte, error) {
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
