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

package mock

import (
	"context"
	"math/big"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/attestantio/go-execution-client/api"
	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
)

// Service is a mock execution client service.
type Service struct{}

// New creates a new mock.
func New() (*Service, error) {
	return &Service{}, nil
}

// Name returns the name of the client implementation.
func (*Service) Name() string { return "mock" }

// Address returns the address of the client.
func (*Service) Address() string { return "mock" }

// Balance obtains the balance for the given address at the given block ID.
func (*Service) Balance(_ context.Context, _ types.Address, _ string) (*big.Int, error) {
	return big.NewInt(0), nil
}

// ReplayBlockTransactions obtains traces for all transactions in a block.
func (*Service) ReplayBlockTransactions(_ context.Context, _ string) ([]*api.TransactionResult, error) {
	return []*api.TransactionResult{}, nil
}

// Block returns the block with the given ID.
func (*Service) Block(_ context.Context, _ string) (*spec.Block, error) {
	return &spec.Block{}, nil
}

// ChainHeight returns the height of the chain as understood by the node.
func (*Service) ChainHeight(_ context.Context) (uint32, error) {
	return 0, nil
}

// ChainID returns the chain ID.
func (*Service) ChainID(_ context.Context) (uint64, error) {
	return 0, nil
}

// Events returns the events matching the filter.
func (*Service) Events(_ context.Context, _ *api.EventsFilter) ([]*spec.BerlinTransactionEvent, error) {
	return []*spec.BerlinTransactionEvent{}, nil
}

// Issuance returns the issuance of a block.
func (*Service) Issuance(_ context.Context, _ string) (*api.Issuance, error) {
	return &api.Issuance{}, nil
}

// NetworkID returns the network ID.
func (*Service) NetworkID(_ context.Context) (uint64, error) {
	return 0, nil
}

// NewPendingTransactions subscribes to new pending transactions.
func (*Service) NewPendingTransactions(_ context.Context, _ chan *spec.Transaction) (*util.Subscription, error) {
	return &util.Subscription{}, nil
}

// Syncing obtains information about the sync state of the node.
func (*Service) Syncing(_ context.Context) (*api.SyncState, error) {
	return &api.SyncState{}, nil
}

// Transaction returns the transaction for the given transaction hash.
func (*Service) Transaction(_ context.Context, _ types.Hash) (*spec.Transaction, error) {
	return &spec.Transaction{}, nil
}

// TransactionInBlock returns the transaction for the given transaction in a block at the given index.
func (*Service) TransactionInBlock(_ context.Context, _ types.Hash, _ uint32) (*spec.Transaction, error) {
	return &spec.Transaction{}, nil
}

// TransactionReceipt returns the transaction receipt for the given transaction hash.
func (*Service) TransactionReceipt(_ context.Context, _ types.Hash) (*spec.TransactionReceipt, error) {
	return &spec.TransactionReceipt{}, nil
}

// Call makes a call to the execution client.
func (*Service) Call(_ context.Context, _ *execclient.CallOpts) ([]byte, error) {
	return []byte{}, nil
}
