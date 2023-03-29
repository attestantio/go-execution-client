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
)

// ChainHeight returns the height of the chain as understood by the node.
func (s *Service) ChainHeight(_ context.Context) (uint32, error) {
	res := ""
	if err := s.client.CallFor(&res, "eth_blockNumber"); err != nil {
		return 0, err
	}

	height, err := strconv.ParseUint(strings.TrimPrefix(res, "0x"), 16, 32)
	if err != nil {
		return 0, err
	}

	return uint32(height), nil
}
