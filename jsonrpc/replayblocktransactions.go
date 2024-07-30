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
	"strconv"
	"strings"

	"github.com/attestantio/go-execution-client/api"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// ReplayBlockTransactions obtains traces for all transactions in a block.
func (s *Service) ReplayBlockTransactions(ctx context.Context, blockID string) ([]*api.TransactionResult, error) {
	if strings.HasPrefix(blockID, "0x") {
		return nil, errors.New("fetch by block hash not implemented")
	}
	height, err := strconv.ParseInt(blockID, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "unhandled block ID")
	}

	return s.replayBlockTransactionsAtHeight(ctx, height)
}

func (s *Service) replayBlockTransactionsAtHeight(_ context.Context, height int64) ([]*api.TransactionResult, error) {
	var transactionResults []*api.TransactionResult

	log.Trace().Int64("height", height).Msg("Replaying block transactions")
	var err error
	switch {
	case height < 0:
		err = s.client.CallFor(&transactionResults, "trace_replayBlockTransactions", "latest", []string{"stateDiff"})
	case height == 0:
		// Block 0 is a special case, with no transactions.
		transactionResults = make([]*api.TransactionResult, 0)
	default:
		err = s.client.CallFor(&transactionResults,
			"trace_replayBlockTransactions",
			util.MarshalUint32(uint32(height)),
			[]string{"stateDiff"},
		)
	}
	if err != nil {
		return nil, err
	}

	return transactionResults, nil
}
