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
	"strconv"
	"strings"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

var blockIdentifiers = map[string]struct{}{
	"earliest":  {},
	"latest":    {},
	"pending":   {},
	"safe":      {},
	"finalized": {},
}

// Block returns the block given an ID.
func (s *Service) Block(ctx context.Context, blockID string) (*spec.Block, error) {
	if blockID == "" ||
		blockID == "latest" {
		return s.blockAtHeight(ctx, -1)
	}
	if _, isIdentifier := blockIdentifiers[blockID]; isIdentifier {
		return s.blockAtIdentifier(ctx, blockID)
	}
	if strings.HasPrefix(blockID, "0x") {
		return s.blockAtHash(ctx, blockID)
	}
	height, err := strconv.ParseInt(blockID, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "unhandled block ID")
	}

	return s.blockAtHeight(ctx, height)
}

func (s *Service) blockAtHash(_ context.Context, hash string) (*spec.Block, error) {
	var block spec.Block

	if err := s.client.CallFor(&block, "eth_getBlockByHash", hash, true); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("eth_getBlockByHash for %#x failed", hash))
	}

	return &block, nil
}

func (s *Service) blockAtIdentifier(_ context.Context, id string) (*spec.Block, error) {
	var block spec.Block

	if err := s.client.CallFor(&block, "eth_getBlockByNumber", id, true); err != nil {
		return nil, errors.Wrapf(err, "eth_getBlockByNumber for %s failed", id)
	}

	return &block, nil
}

func (s *Service) blockAtHeight(ctx context.Context, height int64) (*spec.Block, error) {
	if height == -1 {
		return s.blockAtIdentifier(ctx, "latest")
	}

	return s.blockAtIdentifier(ctx, util.MarshalInt64(height))
}
