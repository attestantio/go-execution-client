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
	"os"
	"testing"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/attestantio/go-execution-client/jsonrpc"
	"github.com/attestantio/go-execution-client/spec"
	"github.com/stretchr/testify/require"
)

// TestNewPendingTransactions tests the TestNewPendingTransactions function.
func TestNewPendingTransactions(t *testing.T) {
	ctx := context.Background()
	s, err := jsonrpc.New(ctx,
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	ch := make(chan *spec.Transaction)
	subscription, err := s.(execclient.NewPendingTransactionsProvider).NewPendingTransactions(ctx, ch)
	require.NoError(t, err)
	require.NotNil(t, subscription)

	// Wait to see a transaction.
	tx := <-ch
	require.NotNil(t, tx)
	//		// TODO remove.
	//		{
	//			data, err := json.Marshal(tx)
	//			require.NoError(t, err)
	//			fmt.Printf("%s\n", string(data))
	//		}
}
