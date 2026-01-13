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

package jsonrpc

import (
	"context"
	"math/big"
	"strconv"
	"strings"

	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

type feeHistory struct {
	BaseFeePerGas []string `json:"baseFeePerGas"`
}

// BaseFee provides the base fee of the chain at the given block ID.
func (s *Service) BaseFee(_ context.Context,
	blockID string,
) (
	*big.Int,
	error,
) {
	if blockID == "" {
		return nil, errors.New("block ID not specified")
	}

	pending := false
	if blockID == "pending" {
		pending = true
		blockID = "latest"
	}

	// Convert decimal heights to hex.
	_, isIdentifier := blockIdentifiers[blockID]
	if !strings.HasPrefix(blockID, "0x") && !isIdentifier {
		tmp, err := strconv.ParseInt(blockID, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "unhandled block ID")
		}

		blockID = util.MarshalInt64(tmp)
	}

	res := feeHistory{}
	if err := s.client.CallFor(&res, "eth_feeHistory", "0x1", blockID, []float64{0}); err != nil {
		return nil, errors.Wrap(err, "call to eth_feeHistory failed")
	}

	if len(res.BaseFeePerGas) == 0 {
		return nil, errors.New("no data returned")
	}

	var baseFeePerGasStr string

	if pending {
		if len(res.BaseFeePerGas) < 2 {
			return nil, errors.New("no pending data returned")
		}

		baseFeePerGasStr = res.BaseFeePerGas[1]
	} else {
		baseFeePerGasStr = res.BaseFeePerGas[0]
	}

	baseFeePerGas, err := util.StrToBigInt("baseFeePerGas", baseFeePerGasStr)
	if err != nil {
		return nil, err
	}

	return baseFeePerGas, nil
}
