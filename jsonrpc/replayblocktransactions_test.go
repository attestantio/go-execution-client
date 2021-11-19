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

package jsonrpc_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/attestantio/go-execution-client/jsonrpc"
	"github.com/stretchr/testify/require"
)

func TestReplayBlockTransactions(t *testing.T) {
	ctx := context.Background()
	s, err := jsonrpc.New(ctx,
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	tests := []struct {
		name     string
		blockID  string
		err      string
		expected string
	}{
		{
			name:     "NoTransactions",
			blockID:  "12345",
			expected: `[]`,
		},
		{
			name:     "FirstTransaction",
			blockID:  "46147",
			expected: `[{"stateDiff":{"0x5df9b87991262f6ba471f09758cde1c0fc1de734":{"balance":{"+":"0x7a69"},"nonce":{"+":"0x0"}},"0xa1e4380a3b1f749673e270229993ee55f35663b4":{"balance":{"*":{"from":"0x6c6b935b8bbd400000","to":"0x6c5d01021be7168597"}},"nonce":{"*":{"from":"0x0","to":"0x1"}}},"0xe6a7a1d47ff21b6321162aea7c6cb457d5476bca":{"balance":{"*":{"from":"0xf3426785a8ab466000","to":"0xf350f9df18816f6000"}}}},"transactionHash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060"}]`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			replay, err := s.(execclient.BlockReplaysProvider).ReplayBlockTransactions(ctx, test.blockID)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, replay)
				res, err := json.Marshal(replay)
				require.NoError(t, err)
				require.Equal(t, test.expected, string(res))
			}
		})
	}
}
