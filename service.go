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

package client

import (
	"context"

	"github.com/attestantio/go-execution-client/api"
	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
)

// Service is the service providing a connection to an execution client.
type Service interface {
	// Name returns the name of the client implementation.
	Name() string

	// Address returns the address of the client.
	Address() string
}

// BlockReplaysProvider is the interface for providing block replays.
type BlockReplaysProvider interface {
	// ReplayBlockTransactions obtains traces for all transactions in a block.
	ReplayBlockTransactions(ctx context.Context, blockID string) ([]*api.TransactionResult, error)
}

// BlocksProvider is the interface for providing blocks.
type BlocksProvider interface {
	// Block returns the block with the given ID.
	Block(ctx context.Context, blockID string) (*spec.Block, error)
}

// ChainHeightProvider is the interface for providing chain height.
type ChainHeightProvider interface {
	// ChainHeight returns the height of the chain as understood by the node.
	ChainHeight(ctx context.Context) (uint32, error)
}

// EventsProvider is the interface for providing events.
type EventsProvider interface {
	// Events returns the events matching the filter.
	Events(ctx context.Context, filter *api.EventsFilter) ([]*spec.TransactionEvent, error)
}

// IssuanceProvider is the interface for providing issuance.
type IssuanceProvider interface {
	// Issuance returns the issuance of a block.
	Issuance(ctx context.Context, blockID string) (*api.Issuance, error)
}

// NetworkIDProvider is the interface for providing the network ID.
type NetworkIDProvider interface {
	// NetworkID returns the network ID.
	NetworkID(ctx context.Context) (uint64, error)
}

// SyncingProvider is the interface for providing syncing information.
type SyncingProvider interface {
	// Syncing obtains information about the sync state of the node.
	Syncing(ctx context.Context) (*api.SyncState, error)
}

// TransactionReceiptsProvider is the interface for providing transaction receipts.
type TransactionReceiptsProvider interface {
	// TransactionReceipt returns the transaction receipt for the given transaction hash.
	TransactionReceipt(ctx context.Context, hash types.Hash) (*spec.TransactionReceipt, error)
}
