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

// type1TransactionJSON is the spec representation of a type 1 transaction.
type type1TransactionJSON struct {
	AccessList       []*AccessListEntry `json:"accessList"`
	BlockHash        string             `json:"blockHash"`
	BlockNumber      string             `json:"blockNumber"`
	ChainID          string             `json:"chainId"`
	From             string             `json:"from"`
	Gas              string             `json:"gas"`
	GasPrice         string             `json:"gasPrice"`
	Hash             string             `json:"hash"`
	Input            string             `json:"input"`
	Nonce            string             `json:"nonce"`
	R                string             `json:"r"`
	S                string             `json:"s"`
	To               string             `json:"to"`
	TransactionIndex string             `json:"transactionIndex"`
	Type             string             `json:"type"`
	V                string             `json:"v"`
	Value            string             `json:"value"`
}

// MarshalType1JSON marshals a type 1 transaction.
func (t *Transaction) MarshalType1JSON() ([]byte, error) {
	to := ""
	if t.To != nil {
		to = util.MarshalByteArray(t.To[:])
	}
	return json.Marshal(&type1TransactionJSON{
		AccessList:       t.AccessList,
		BlockHash:        util.MarshalByteArray(t.BlockHash[:]),
		BlockNumber:      util.MarshalUint32(t.BlockNumber),
		ChainID:          util.MarshalUint64(t.ChainID),
		From:             util.MarshalByteArray(t.From[:]),
		Gas:              util.MarshalUint32(t.Gas),
		GasPrice:         util.MarshalUint64(t.GasPrice),
		Hash:             util.MarshalByteArray(t.Hash[:]),
		Input:            util.MarshalByteArray(t.Input),
		Nonce:            util.MarshalUint64(t.Nonce),
		R:                util.MarshalBigInt(t.R),
		S:                util.MarshalBigInt(t.S),
		To:               to,
		Type:             "0x1",
		TransactionIndex: util.MarshalUint32(t.TransactionIndex),
		V:                util.MarshalBigInt(t.V),
		Value:            util.MarshalBigInt(t.Value),
	})
}
