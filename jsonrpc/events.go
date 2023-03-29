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

	"github.com/attestantio/go-execution-client/api"
	"github.com/attestantio/go-execution-client/spec"
	"github.com/pkg/errors"
)

// Events returns the events matching the filter.
func (s *Service) Events(_ context.Context, filter *api.EventsFilter) ([]*spec.BerlinTransactionEvent, error) {
	if filter == nil {
		return nil, errors.New("filter not specified")
	}

	var events []*spec.BerlinTransactionEvent

	if err := s.client.CallFor(&events, "eth_getLogs", []*api.EventsFilter{filter}); err != nil {
		return nil, err
	}

	return events, nil
}
