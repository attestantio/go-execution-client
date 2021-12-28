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
	"encoding/json"

	"github.com/attestantio/go-execution-client/util"
)

// type2TransactionJSON is the spec representation of a type 2 transaction.
type type2TransactionJSON struct {
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

// MarshalType2JSON marshals a type 2 transaction.
func (t *Transaction) MarshalType2JSON() ([]byte, error) {
	to := ""
	if t.To != nil {
		to = util.MarshalByteArray(t.To[:])
	}
	return json.Marshal(&type2TransactionJSON{
		AccessList:           t.AccessList,
		BlockHash:            util.MarshalByteArray(t.BlockHash[:]),
		BlockNumber:          util.MarshalUint32(t.BlockNumber),
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
		TransactionIndex:     util.MarshalUint32(t.TransactionIndex),
		V:                    util.MarshalBigInt(t.V),
		Value:                util.MarshalBigInt(t.Value),
	})
}

// MarshalType2RLP returns an RLP representation of the transaction.
func (t *Transaction) MarshalType2RLP() ([]byte, error) {
	items := make([][]byte, 12)

	items[0] = util.RLPUint64(t.ChainID)
	items[1] = util.RLPUint64(t.Nonce)
	items[2] = util.RLPUint64(t.MaxPriorityFeePerGas)
	items[3] = util.RLPUint64(t.MaxFeePerGas)
	items[4] = util.RLPUint64(uint64(t.Gas))
	if t.To != nil {
		items[5] = util.RLPAddress(*t.To)
	} else {
		items[5] = util.RLPBytes(nil)
	}
	if t.Value != nil {
		items[6] = util.RLPBytes(t.Value.Bytes())
	} else {
		items[6] = util.RLPBytes(nil)
	}
	items[7] = util.RLPBytes(t.Input)
	accessList := make([][]byte, len(t.AccessList))
	for i, accessListEntry := range t.AccessList {
		list := make([][]byte, 2)
		list[0] = util.RLPBytes(accessListEntry.Address)
		keys := make([][]byte, len(accessListEntry.StorageKeys))
		for j, key := range accessListEntry.StorageKeys {
			keys[j] = util.RLPBytes(key)
		}
		list[1] = util.RLPList(keys)
		accessList[i] = util.RLPList(list)
	}
	items[8] = util.RLPList(accessList)

	if t.V.Uint64() != 0 {
		items[9] = util.RLPBytes([]byte{byte(int8(t.V.Uint64()))})
	} else {
		items[9] = util.RLPBytes(nil)
	}
	items[10] = util.RLPBytes(t.R.Bytes())
	items[11] = util.RLPBytes(t.S.Bytes())

	list := append([]byte{0x02}, util.RLPList(items)...)
	return util.RLPBytes(list), nil
}
