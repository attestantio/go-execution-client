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

	"github.com/pkg/errors"
)

func (s *Service) blockIDToHeight(ctx context.Context, blockID string) (int64, error) {
	var height int64
	switch {
	case blockID == "":
		height = -1
	case strings.HasPrefix(blockID, "0x"):
		block, err := s.blockAtHash(ctx, blockID)
		if err != nil {
			return -1, err
		}
		if block == nil {
			return -1, errors.New("block not found")
		}
		height = int64(block.London.Number)
	default:
		var err error
		height, err = strconv.ParseInt(blockID, 10, 64)
		if err != nil {
			return -1, errors.Wrap(err, "unhandled block ID")
		}
	}
	return height, nil
}
