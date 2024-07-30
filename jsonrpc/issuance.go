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

	"github.com/attestantio/go-execution-client/api"
	"github.com/pkg/errors"
)

// Issuance returns the issuance of a block.
func (s *Service) Issuance(ctx context.Context, blockID string) (*api.Issuance, error) {
	if !s.isIssuanceProvider {
		return nil, errors.New("client does not provide issuance")
	}

	height, err := s.blockIDToHeight(ctx, blockID)
	if err != nil {
		return nil, err
	}

	return s.issuanceAtHeight(ctx, height)
}

func (s *Service) issuanceAtHeight(_ context.Context, height int64) (*api.Issuance, error) {
	var issuance api.Issuance

	if height == -1 {
		if err := s.client.CallFor(&issuance, "erigon_issuance", "latest"); err != nil {
			return nil, err
		}
	} else {
		if err := s.client.CallFor(&issuance, "erigon_issuance", fmt.Sprintf("0x%x", height)); err != nil {
			return nil, err
		}
	}

	return &issuance, nil
}
