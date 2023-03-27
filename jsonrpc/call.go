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
	"encoding/hex"
	"fmt"
	"strings"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// Call makes a call to the execution client.
func (s *Service) Call(ctx context.Context, opts *execclient.CallOpts) ([]byte, error) {
	if opts == nil {
		return nil, errors.New("no options specified")
	}

	callOpts := make(map[string]string)
	if opts.From != nil {
		callOpts["from"] = opts.From.String()
	}
	if opts.To != nil {
		callOpts["to"] = opts.To.String()
	}
	if opts.Gas != nil {
		callOpts["gas"] = util.MarshalBigInt(opts.Gas)
	}
	if opts.GasPrice != nil {
		callOpts["gasPrice"] = util.MarshalBigInt(opts.GasPrice)
	}
	if opts.Value != nil {
		callOpts["value"] = util.MarshalBigInt(opts.Value)
	}
	if opts.Data != nil {
		callOpts["data"] = fmt.Sprintf("%#x", opts.Data)
	}

	block := "latest"
	if opts.Block != "" {
		block = opts.Block
	}

	var callResults string
	err := s.client.CallFor(&callResults, "eth_call", callOpts, block)
	if err != nil {
		return nil, errors.Wrap(err, "eth_call failed")
	}

	res, err := hex.DecodeString(strings.TrimPrefix(callResults, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "invalid response")
	}

	return res, nil
}
