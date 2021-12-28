// Copyright © 2021 Attestant Limited.
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
	"encoding/json"

	"github.com/attestantio/go-execution-client/util"
)

// type0TransactionJSON is the spec representation of a type 0 transaction.
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

// MarshalType0JSON marshals a type 0 transaction.
func (t *Transaction) MarshalType0JSON() ([]byte, error) {
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
		TransactionIndex: util.MarshalUint32(t.TransactionIndex),
		V:                util.MarshalBigInt(t.V),
		Value:            util.MarshalBigInt(t.Value),
	})
}

// MarshalType0RLP returns an RLP representation of the transaction.
func (t *Transaction) MarshalType0RLP() ([]byte, error) {
	items := make([][]byte, 9)

	items[0] = util.RLPUint64(t.Nonce)
	items[1] = util.RLPUint64(t.GasPrice)
	items[2] = util.RLPUint64(uint64(t.Gas))
	if t.To != nil {
		items[3] = util.RLPAddress(*t.To)
	} else {
		items[3] = util.RLPBytes(nil)
	}
	if t.Value != nil {
		items[4] = util.RLPBytes(t.Value.Bytes())
	} else {
		items[4] = util.RLPBytes(nil)
	}
	items[5] = util.RLPBytes(t.Input)
	items[6] = util.RLPBytes([]byte{byte(int8(t.V.Uint64()))})
	items[7] = util.RLPBytes(t.R.Bytes())
	items[8] = util.RLPBytes(t.S.Bytes())

	return util.RLPList(items), nil
}
