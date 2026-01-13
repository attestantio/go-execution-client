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

package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// EstimateGas estimates the gas required for a transaction.
func (s *Service) EstimateGas(_ context.Context,
	tx *spec.TransactionSubmission,
) (
	*big.Int,
	error,
) {
	if tx == nil {
		return nil, errors.New("no transaction specified")
	}

	opts := make(map[string]any)

	opts["type"] = tx.Type.String()
	if tx.ChainID != nil {
		opts["chainId"] = util.MarshalBigInt(tx.ChainID)
	}

	opts["from"] = tx.From.String()
	if tx.Nonce != nil {
		opts["nonce"] = util.MarshalUint64(*tx.Nonce)
	}

	if tx.To != nil {
		opts["to"] = tx.To.String()
	}

	if tx.Input != nil {
		opts["input"] = fmt.Sprintf("%#x", tx.Input)
	}

	if tx.Gas != nil {
		opts["gas"] = util.MarshalBigInt(tx.Gas)
	}

	if tx.GasPrice != nil {
		opts["gasPrice"] = util.MarshalBigInt(tx.GasPrice)
	}

	if tx.MaxPriorityFeePerGas != nil {
		opts["maxPriorityFeePerGas"] = util.MarshalBigInt(tx.MaxPriorityFeePerGas)
	}

	if tx.MaxFeePerGas != nil {
		opts["maxFeePerGas"] = util.MarshalBigInt(tx.MaxFeePerGas)
	}

	if tx.MaxFeePerBlobGas != nil {
		opts["maxFeePerBlobGas"] = util.MarshalBigInt(tx.MaxFeePerBlobGas)
	}

	if tx.Value != nil {
		opts["value"] = util.MarshalBigInt(tx.Value)
	}

	if tx.AccessList != nil {
		accessList, err := json.Marshal(tx.AccessList)
		if err != nil {
			return nil, err
		}

		opts["accessList"] = string(accessList)
	}

	if tx.BlobVersionedHashes != nil {
		blobVersionedHashes, err := json.Marshal(tx.BlobVersionedHashes)
		if err != nil {
			return nil, err
		}

		opts["blobVersionedHashes"] = string(blobVersionedHashes)
	}

	if tx.Blobs != nil {
		blobs, err := json.Marshal(tx.Blobs)
		if err != nil {
			return nil, err
		}

		opts["blobs"] = string(blobs)
	}

	var gasStr string
	if err := s.client.CallFor(&gasStr, "eth_estimateGas", opts, "latest"); err != nil {
		return nil, errors.Wrap(err, "call to eth_estimateGas failed")
	}

	gas, err := util.StrToBigInt("gas", gasStr)
	if err != nil {
		return nil, err
	}

	return gas, nil
}
