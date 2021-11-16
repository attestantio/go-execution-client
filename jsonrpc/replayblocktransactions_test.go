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
			expected: "[]",
		},
		{
			name:     "FirstTransaction",
			blockID:  "46147",
			expected: "[]",
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
