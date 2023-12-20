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
	"math/big"

	"github.com/attestantio/go-execution-client/types"
)

// TransactionSubmission provides the partial details of a transaction to submit to the network.
type TransactionSubmission struct {
	Type                 TransactionType
	ChainID              *big.Int
	From                 types.Address
	Nonce                *uint64
	To                   *types.Address
	Input                []byte
	Gas                  *big.Int
	GasPrice             *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	MaxFeePerBlobGas     *big.Int
	Value                *big.Int
	AccessList           []*AccessListEntry
	BlobVersionedHashes  []types.VersionedHash
	Blobs                []types.Blob
}
