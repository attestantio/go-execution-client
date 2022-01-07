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
	"github.com/stretchr/testify/require"
)

// TestTransaction tests the Transaction() call.
func TestTransaction(t *testing.T) {
	ctx := context.Background()
	s, err := jsonrpc.New(ctx,
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	txHash := strToHash("0xc9e58c3c7cfd290033b37576feb1649b9609e1ad49fe1fd23148423bb2b44bd2")
	tx, err := s.(execclient.TransactionsProvider).Transaction(ctx, txHash)
	require.NoError(t, err)
	require.NotNil(t, tx)
	require.Equal(t, tx.Hash, txHash)
}
