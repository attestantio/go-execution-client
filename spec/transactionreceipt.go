// Copyright © 2022 Attestant Limited.
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

// TransactionReceipt is a struct that covers all transaction receipt versions.
type TransactionReceipt struct {
	Fork                     Fork
	BerlinTransactionReceipt *BerlinTransactionReceipt
	LondonTransactionReceipt *LondonTransactionReceipt
	CancunTransactionReceipt *CancunTransactionReceipt
}

// transactionReceiptJSON is a simple struct to fetch the transaction type.
type transactionReceiptJSON struct {
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	BlobGasPrice      string `json:"blobGasPrice"`
}

// MarshalJSON marshals a typed transaction.
func (t *TransactionReceipt) MarshalJSON() ([]byte, error) {
	switch t.Fork {
	case ForkBerlin:
		return json.Marshal(t.BerlinTransactionReceipt)
	case ForkLondon, ForkShanghai, ForkCancun:
		return json.Marshal(t.LondonTransactionReceipt)
	default:
		return nil, fmt.Errorf("unhandled transaction receipt fork %v", t.Fork)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionReceipt) UnmarshalJSON(input []byte) error {
	var data transactionReceiptJSON
	err := json.Unmarshal(input, &data)
	if err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	switch {
	case data.BlobGasPrice != "":
		t.Fork = ForkCancun
		t.CancunTransactionReceipt = &CancunTransactionReceipt{}
		err = json.Unmarshal(input, t.CancunTransactionReceipt)
	case data.EffectiveGasPrice != "":
		t.Fork = ForkLondon
		t.LondonTransactionReceipt = &LondonTransactionReceipt{}
		err = json.Unmarshal(input, t.LondonTransactionReceipt)
	default:
		t.Fork = ForkBerlin
		t.BerlinTransactionReceipt = &BerlinTransactionReceipt{}
		err = json.Unmarshal(input, t.BerlinTransactionReceipt)
	}

	return err
}

// BlobGasPrice returns the blob gas price of the transaction receipt.
func (t *TransactionReceipt) BlobGasPrice() *big.Int {
	switch t.Fork {
	case ForkBerlin, ForkLondon, ForkShanghai:
		return nil
	case ForkCancun:
		return t.CancunTransactionReceipt.BlobGasPrice
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// BlobGasUsed returns the blob gas used of the transaction receipt.
func (t *TransactionReceipt) BlobGasUsed() uint32 {
	switch t.Fork {
	case ForkBerlin, ForkLondon, ForkShanghai:
		return 0
	case ForkCancun:
		return t.CancunTransactionReceipt.BlobGasUsed
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// BlockHash returns the block hash of the transaction receipt.
func (t *TransactionReceipt) BlockHash() types.Hash {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.BlockHash
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.BlockHash
	case ForkCancun:
		return t.CancunTransactionReceipt.BlockHash
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// BlockNumber returns the block number of the transaction receipt.
func (t *TransactionReceipt) BlockNumber() uint32 {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.BlockNumber
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.BlockNumber
	case ForkCancun:
		return t.CancunTransactionReceipt.BlockNumber
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// ContractAddress returns the contract address of the transaction receipt.
// This will be nil for transactions that did not create a contract.
func (t *TransactionReceipt) ContractAddress() *types.Address {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.ContractAddress
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.ContractAddress
	case ForkCancun:
		return t.CancunTransactionReceipt.ContractAddress
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// CumulativeGasUsed returns the cumulative gas used in the block up to this receipt.
func (t *TransactionReceipt) CumulativeGasUsed() uint32 {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.CumulativeGasUsed
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.CumulativeGasUsed
	case ForkCancun:
		return t.CancunTransactionReceipt.CumulativeGasUsed
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// EffectiveGasPrice returns the effective gas price of the transaction.
// This will return 0 for pre-London transactions.
func (t *TransactionReceipt) EffectiveGasPrice() uint64 {
	switch t.Fork {
	case ForkBerlin:
		return 0
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.EffectiveGasPrice
	case ForkCancun:
		return t.CancunTransactionReceipt.EffectiveGasPrice
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// From returns the sender of the transaction receipt.
func (t *TransactionReceipt) From() types.Address {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.From
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.From
	case ForkCancun:
		return t.CancunTransactionReceipt.From
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// GasUsed returns the gas used by the transaction.
func (t *TransactionReceipt) GasUsed() uint32 {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.GasUsed
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.GasUsed
	case ForkCancun:
		return t.CancunTransactionReceipt.GasUsed
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// Logs returns the logs generated by the transaction.
func (t *TransactionReceipt) Logs() []*BerlinTransactionEvent {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.Logs
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.Logs
	case ForkCancun:
		return t.CancunTransactionReceipt.Logs
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// LogsBloom returns the logs bloom generated by the transaction.
func (t *TransactionReceipt) LogsBloom() []byte {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.LogsBloom
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.LogsBloom
	case ForkCancun:
		return t.CancunTransactionReceipt.LogsBloom
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// Status returns the status returned by the transaction.
func (t *TransactionReceipt) Status() uint32 {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.Status
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.Status
	case ForkCancun:
		return t.CancunTransactionReceipt.Status
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// To returns the recipient of the transaction receipt.
// This value will be nil for contract creation.
func (t *TransactionReceipt) To() *types.Address {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.To
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.To
	case ForkCancun:
		return t.CancunTransactionReceipt.To
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// TransactionHash returns the hash of the transaction.
func (t *TransactionReceipt) TransactionHash() types.Hash {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.TransactionHash
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.TransactionHash
	case ForkCancun:
		return t.CancunTransactionReceipt.TransactionHash
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// TransactionIndex returns the index of the transaction in the block.
func (t *TransactionReceipt) TransactionIndex() uint32 {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.TransactionIndex
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.TransactionIndex
	case ForkCancun:
		return t.CancunTransactionReceipt.TransactionIndex
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// Type returns the type of the transaction in the block.
func (t *TransactionReceipt) Type() TransactionType {
	switch t.Fork {
	case ForkBerlin:
		return t.BerlinTransactionReceipt.Type
	case ForkLondon, ForkShanghai:
		return t.LondonTransactionReceipt.Type
	case ForkCancun:
		return t.CancunTransactionReceipt.Type
	default:
		panic(fmt.Errorf("unhandled transaction receipt fork %s", t.Fork))
	}
}

// String returns a string version of the structure.
func (t *TransactionReceipt) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
