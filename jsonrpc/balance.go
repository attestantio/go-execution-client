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
	"math/big"
	"strconv"
	"strings"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// Balance obtains the balance for the given address at the given block ID.
func (s *Service) Balance(ctx context.Context, address types.Address, blockID string) (*big.Int, error) {
	if blockID == "" || blockID == "latest" {
		return s.balanceAtHeight(ctx, address, -1)
	}

	if strings.HasPrefix(blockID, "0x") {
		return s.balanceAtHash(ctx, address, blockID)
	}

	height, err := strconv.ParseInt(blockID, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "unhandled block ID")
	}

	return s.balanceAtHeight(ctx, address, height)
}

//nolint:unparam
func (s *Service) balanceAtHash(ctx context.Context,
	address types.Address,
	hash string,
) (
	*big.Int,
	error,
) {
	var block spec.Block

	if err := s.client.CallFor(&block, "eth_getBlockByHash", hash, false); err != nil {
		return nil, err
	}

	return s.balanceAtHash(ctx, address, fmt.Sprintf("%#x", block.Hash()))
}

func (s *Service) balanceAtHeight(_ context.Context,
	address types.Address,
	height int64,
) (
	*big.Int,
	error,
) {
	var (
		balanceStr string
		err        error
	)

	if height == -1 {
		err = s.client.CallFor(&balanceStr, "eth_getBalance", fmt.Sprintf("%#x", address), "latest")
	} else {
		err = s.client.CallFor(&balanceStr, "eth_getBalance", fmt.Sprintf("%#x", address), fmt.Sprintf("%#x", height))
	}

	if err != nil {
		return nil, err
	}

	balance, err := util.StrToBigInt("balance", balanceStr)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
