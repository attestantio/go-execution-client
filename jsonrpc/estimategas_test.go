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

package jsonrpc_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/attestantio/go-execution-client/jsonrpc"
	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestEstimateGas(t *testing.T) {
	ctx := context.Background()
	s, err := jsonrpc.New(ctx,
		jsonrpc.WithLogLevel(zerolog.Disabled),
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	tests := []struct {
		name string
		tx   *spec.TransactionSubmission
		err  string
	}{
		{
			name: "Type0Transfer",
			tx: &spec.TransactionSubmission{
				Type: spec.TransactionType0,
				From: types.Address{},
				To:   &types.Address{},
			},
		},
		{
			name: "Type0ContractCreation",
			tx: &spec.TransactionSubmission{
				Type: spec.TransactionType0,
				From: types.Address{},
				Gas:  big.NewInt(100000),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := s.(execclient.GasEstimationProvide).EstimateGas(ctx, test.tx)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
