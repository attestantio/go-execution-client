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
	"time"

	"github.com/attestantio/go-execution-client/types"
	"github.com/pkg/errors"
)

// Block is a struct that covers versioned blocks.
type Block struct {
	Fork     Fork
	Berlin   *BerlinBlock
	London   *LondonBlock
	Shanghai *ShanghaiBlock
}

// blockTypeJSON is a struct that helps us identify the block type.
type blockTypeJSON struct {
	BaseFeePerGas string                   `json:"baseFeePerGas"`
	Withdrawals   []map[string]interface{} `json:"withdrawals"`
}

// MarshalJSON marshals a typed transaction.
func (b *Block) MarshalJSON() ([]byte, error) {
	switch b.Fork {
	case ForkBerlin:
		return json.Marshal(b.Berlin)
	case ForkLondon:
		return json.Marshal(b.London)
	case ForkShanghai:
		return json.Marshal(b.Shanghai)
	default:
		return nil, fmt.Errorf("unhandled block version %v", b.Fork)
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Block) UnmarshalJSON(input []byte) error {
	var data blockTypeJSON
	err := json.Unmarshal(input, &data)
	if err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	switch {
	case len(data.Withdrawals) > 0:
		b.Fork = ForkShanghai
		b.Shanghai = &ShanghaiBlock{}
		err = json.Unmarshal(input, b.Shanghai)
	case len(data.BaseFeePerGas) > 0:
		b.Fork = ForkLondon
		b.London = &LondonBlock{}
		err = json.Unmarshal(input, b.London)
	default:
		b.Fork = ForkBerlin
		b.Berlin = &BerlinBlock{}
		err = json.Unmarshal(input, b.Berlin)
	}

	return err
}

// BaseFeePerGas returns the base fee per gas of the block.
// This value will be 0 if the block does not use base fee (e.g. pre-London).
func (b *Block) BaseFeePerGas() uint64 {
	switch b.Fork {
	case ForkBerlin:
		return 0
	case ForkLondon:
		return b.London.BaseFeePerGas
	case ForkShanghai:
		return b.Shanghai.BaseFeePerGas
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Difficulty returns the difficulty of the block.
func (b *Block) Difficulty() uint64 {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Difficulty
	case ForkLondon:
		return b.London.Difficulty
	case ForkShanghai:
		return b.Shanghai.Difficulty
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// ExtraData returns the extra data of the block.
func (b *Block) ExtraData() []byte {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.ExtraData
	case ForkLondon:
		return b.London.ExtraData
	case ForkShanghai:
		return b.Shanghai.ExtraData
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// FeeRecipient returns the fee recipient of the block.
// This will return the miner for pre-paris blocks.
func (b *Block) FeeRecipient() types.Address {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Miner
	case ForkLondon:
		return b.London.Miner
	case ForkShanghai:
		return b.Shanghai.Miner
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// GasLimit returns the gas limit of the block.
func (b *Block) GasLimit() uint32 {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.GasLimit
	case ForkLondon:
		return b.London.GasLimit
	case ForkShanghai:
		return b.Shanghai.GasLimit
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// GasUsed returns the gas used of the block.
func (b *Block) GasUsed() uint32 {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.GasUsed
	case ForkLondon:
		return b.London.GasUsed
	case ForkShanghai:
		return b.Shanghai.GasUsed
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Hash returns the hash of the block.
func (b *Block) Hash() types.Hash {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Hash
	case ForkLondon:
		return b.London.Hash
	case ForkShanghai:
		return b.Shanghai.Hash
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// LogsBloom returns the logs bloom of the block.
func (b *Block) LogsBloom() []byte {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.LogsBloom
	case ForkLondon:
		return b.London.LogsBloom
	case ForkShanghai:
		return b.Shanghai.LogsBloom
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Miner returns the miner of the block.
// This will return fee recipient for post-london blocks.
func (b *Block) Miner() types.Address {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Miner
	case ForkLondon:
		return b.London.Miner
	case ForkShanghai:
		return b.Shanghai.Miner
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// MixHash returns the mix hash of the block.
func (b *Block) MixHash() types.Hash {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.MixHash
	case ForkLondon:
		return b.London.MixHash
	case ForkShanghai:
		return b.Shanghai.MixHash
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Nonce returns the nonce of the block.
func (b *Block) Nonce() []byte {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Nonce
	case ForkLondon:
		return b.London.Nonce
	case ForkShanghai:
		return b.Shanghai.Nonce
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Number returns the number of the block.
func (b *Block) Number() uint32 {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Number
	case ForkLondon:
		return b.London.Number
	case ForkShanghai:
		return b.Shanghai.Number
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// ParentHash returns the parent hash of the block.
func (b *Block) ParentHash() types.Hash {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.ParentHash
	case ForkLondon:
		return b.London.ParentHash
	case ForkShanghai:
		return b.Shanghai.ParentHash
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// ReceiptsRoot returns the receipts root of the block.
func (b *Block) ReceiptsRoot() types.Root {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.ReceiptsRoot
	case ForkLondon:
		return b.London.ReceiptsRoot
	case ForkShanghai:
		return b.Shanghai.ReceiptsRoot
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// SHA3Uncles returns the SHA3 hash of the uncles of the block.
func (b *Block) SHA3Uncles() []byte {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.SHA3Uncles
	case ForkLondon:
		return b.London.SHA3Uncles
	case ForkShanghai:
		return b.Shanghai.SHA3Uncles
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Size returns the size of the block.
func (b *Block) Size() uint32 {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Size
	case ForkLondon:
		return b.London.Size
	case ForkShanghai:
		return b.Shanghai.Size
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// StateRoot returns the state root of the block.
func (b *Block) StateRoot() types.Root {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.StateRoot
	case ForkLondon:
		return b.London.StateRoot
	case ForkShanghai:
		return b.Shanghai.StateRoot
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Timestamp returns the timestamp of the block.
func (b *Block) Timestamp() time.Time {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Timestamp
	case ForkLondon:
		return b.London.Timestamp
	case ForkShanghai:
		return b.Shanghai.Timestamp
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// TotalDifficulty returns the total difficulty of the block.
func (b *Block) TotalDifficulty() *big.Int {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.TotalDifficulty
	case ForkLondon:
		return b.London.TotalDifficulty
	case ForkShanghai:
		return b.Shanghai.TotalDifficulty
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Transactions returns the transactions of the block.
func (b *Block) Transactions() []*Transaction {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Transactions
	case ForkLondon:
		return b.London.Transactions
	case ForkShanghai:
		return b.Shanghai.Transactions
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// TransactionsRoot returns the transactions root of the block.
func (b *Block) TransactionsRoot() types.Root {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.TransactionsRoot
	case ForkLondon:
		return b.London.TransactionsRoot
	case ForkShanghai:
		return b.Shanghai.TransactionsRoot
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// Uncles returns the hashes of the uncles of the block.
func (b *Block) Uncles() []types.Hash {
	switch b.Fork {
	case ForkBerlin:
		return b.Berlin.Uncles
	case ForkLondon:
		return b.London.Uncles
	case ForkShanghai:
		return b.Shanghai.Uncles
	default:
		panic(fmt.Sprintf("unhandled block version %v", b.Fork))
	}
}

// String returns a string version of the structure.
func (b *Block) String() string {
	data, err := json.Marshal(b)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
