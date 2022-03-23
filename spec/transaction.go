// Copyright © 2021, 2022 Attestant Limited.
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
	"math/big"

	"github.com/attestantio/go-execution-client/types"
	"github.com/pkg/errors"
)

// Transaction is a struct that covers all transaction types.
type Transaction struct {
	Type             TransactionType
	Type0Transaction *Type0Transaction
	Type1Transaction *Type1Transaction
	Type2Transaction *Type2Transaction
}

// transactionTypeJSON is a simple struct to fetch the transaction type.
type transactionTypeJSON struct {
	Type TransactionType `json:"type"`
}

// MarshalJSON marshals a typed transaction.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	switch t.Type {
	case TransactionType0:
		return json.Marshal(t.Type0Transaction)
	case TransactionType1:
		return json.Marshal(t.Type1Transaction)
	case TransactionType2:
		return json.Marshal(t.Type2Transaction)
	default:
		return nil, fmt.Errorf("unhandled transaction type %v", t.Type)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	var data transactionTypeJSON
	err := json.Unmarshal(input, &data)
	if err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	t.Type = data.Type
	switch t.Type {
	case TransactionTypeUnknown:
		return errors.New("type missing")
	case TransactionType0:
		t.Type0Transaction = &Type0Transaction{}
		err = json.Unmarshal(input, t.Type0Transaction)
	case TransactionType1:
		t.Type1Transaction = &Type1Transaction{}
		err = json.Unmarshal(input, t.Type1Transaction)
	case TransactionType2:
		t.Type2Transaction = &Type2Transaction{}
		err = json.Unmarshal(input, t.Type2Transaction)
	default:
		err = fmt.Errorf("unhandled transaction type %v", data.Type)
	}

	return err
}

// AccessList returns the access list of the transaction.
// This value can be nil, if the transaction does not support access lists.
func (t *Transaction) AccessList() []*AccessListEntry {
	switch t.Type {
	case TransactionType0:
		return nil
	case TransactionType1:
		return t.Type1Transaction.AccessList
	case TransactionType2:
		return t.Type2Transaction.AccessList
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// BlockHash returns the block hash of the transaction.
// This value can be nil, if the transaction is not included in a block.
func (t *Transaction) BlockHash() *types.Hash {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.BlockHash
	case TransactionType1:
		return t.Type1Transaction.BlockHash
	case TransactionType2:
		return t.Type2Transaction.BlockHash
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// BlockNumber returns the block number of the transaction.
// This value can be nil, if the transaction is not included in a block.
func (t *Transaction) BlockNumber() *uint32 {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.BlockNumber
	case TransactionType1:
		return t.Type1Transaction.BlockNumber
	case TransactionType2:
		return t.Type2Transaction.BlockNumber
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// From returns the sender of the transaction.
func (t *Transaction) From() types.Address {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.From
	case TransactionType1:
		return t.Type1Transaction.From
	case TransactionType2:
		return t.Type2Transaction.From
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// Gas returns the gas limit of the transaction.
func (t *Transaction) Gas() uint32 {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.Gas
	case TransactionType1:
		return t.Type1Transaction.Gas
	case TransactionType2:
		return t.Type2Transaction.Gas
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// GasPrice returns the gas price of the transaction.
// This will be 0 for transactions that do not have an individual
// gas price, for example type 2 transactions.
func (t *Transaction) GasPrice() uint64 {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.GasPrice
	case TransactionType1:
		return t.Type1Transaction.GasPrice
	case TransactionType2:
		return 0
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// Hash returns the hash of the transaction.
func (t *Transaction) Hash() types.Hash {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.Hash
	case TransactionType1:
		return t.Type1Transaction.Hash
	case TransactionType2:
		return t.Type2Transaction.Hash
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// Input returns the input data of the transaction.
func (t *Transaction) Input() []byte {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.Input
	case TransactionType1:
		return t.Type1Transaction.Input
	case TransactionType2:
		return t.Type2Transaction.Input
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// MaxFeePerGas returns the maximum fee per gas paid by the transaction.
// This value can be 0, if the transaction does not support this (e.g. type 0 transactions).
func (t *Transaction) MaxFeePerGas() uint64 {
	switch t.Type {
	case TransactionType0:
		return 0
	case TransactionType1:
		return 0
	case TransactionType2:
		return t.Type2Transaction.MaxFeePerGas
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// MaxPriorityFeePerGas returns the maximum priority fee per gas paid by the transaction.
// This value can be 0, if the transaction does not support this (e.g. type 0 transactions).
func (t *Transaction) MaxPriorityFeePerGas() uint64 {
	switch t.Type {
	case TransactionType0:
		return 0
	case TransactionType1:
		return 0
	case TransactionType2:
		return t.Type2Transaction.MaxPriorityFeePerGas
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// Nonce returns the nonce of the transaction.
func (t *Transaction) Nonce() uint64 {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.Nonce
	case TransactionType1:
		return t.Type1Transaction.Nonce
	case TransactionType2:
		return t.Type2Transaction.Nonce
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// R returns the R portion of the signature of the transaction.
func (t *Transaction) R() *big.Int {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.R
	case TransactionType1:
		return t.Type1Transaction.R
	case TransactionType2:
		return t.Type2Transaction.R
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// S returns the S portion of the signature of the transaction.
func (t *Transaction) S() *big.Int {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.S
	case TransactionType1:
		return t.Type1Transaction.S
	case TransactionType2:
		return t.Type2Transaction.S
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// To returns the recipient of the transaction.
// This can be nil, for example on contract creation.
func (t *Transaction) To() *types.Address {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.To
	case TransactionType1:
		return t.Type1Transaction.To
	case TransactionType2:
		return t.Type2Transaction.To
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// TransactionIndex returns the index of the transaction in its block.
// This value can be nil, if the transaction is not included in a block.
func (t *Transaction) TransactionIndex() *uint32 {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.TransactionIndex
	case TransactionType1:
		return t.Type1Transaction.TransactionIndex
	case TransactionType2:
		return t.Type2Transaction.TransactionIndex
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// V returns the V portion of the signature of the transaction.
func (t *Transaction) V() *big.Int {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.V
	case TransactionType1:
		return t.Type1Transaction.V
	case TransactionType2:
		return t.Type2Transaction.V
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// Value returns the value of the transaction.
func (t *Transaction) Value() *big.Int {
	switch t.Type {
	case TransactionType0:
		return t.Type0Transaction.Value
	case TransactionType1:
		return t.Type1Transaction.Value
	case TransactionType2:
		return t.Type2Transaction.Value
	default:
		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
	}
}

// String returns a string version of the structure.
func (t *Transaction) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
