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

package jsonrpc

import (
	"context"
	"fmt"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
	"github.com/pkg/errors"
)

// TransactionReceipt returns the transaction receipt for the given transaction hash.
func (s *Service) TransactionReceipt(_ context.Context, hash types.Hash) (*spec.BerlinTransactionReceipt, error) {
	if len(hash) == 0 {
		return nil, errors.New("hash nil")
	}

	var receipt spec.BerlinTransactionReceipt
	if err := s.client.CallFor(&receipt, "eth_getTransactionReceipt", fmt.Sprintf("%#x", hash)); err != nil {
		return nil, err
	}

	return &receipt, nil
}
