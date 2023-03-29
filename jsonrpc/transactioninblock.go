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

package jsonrpc

import (
	"context"
	"fmt"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
	"github.com/pkg/errors"
)

// TransactionInBlock returns the transaction for the given transaction in a block at the given index.
func (s *Service) TransactionInBlock(_ context.Context, blockHash types.Hash, index uint32) (*spec.Transaction, error) {
	if len(blockHash) == 0 {
		return nil, errors.New("hash nil")
	}

	var transaction spec.Transaction
	if err := s.client.CallFor(&transaction, "eth_getTransactionByBlockHashAndIndex", fmt.Sprintf("%#x", blockHash), fmt.Sprintf("%#x", index)); err != nil {
		return nil, err
	}

	return &transaction, nil
}
