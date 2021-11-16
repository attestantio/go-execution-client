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
	"time"

	client "github.com/attestantio/go-execution-client"
	"github.com/attestantio/go-execution-client/jsonrpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		parameters []jsonrpc.Parameter
		location   string
		err        string
	}{
		{
			name: "Nil",
			err:  "problem with parameters: no address specified",
		},
		{
			name: "AddressNil",
			parameters: []jsonrpc.Parameter{
				jsonrpc.WithTimeout(5 * time.Second),
			},
			err: "problem with parameters: no address specified",
		},
		{
			name: "TimeoutZero",
			parameters: []jsonrpc.Parameter{
				jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
				jsonrpc.WithTimeout(0),
			},
			err: "problem with parameters: no timeout specified",
		},
		{
			name: "AddressInvalid",
			parameters: []jsonrpc.Parameter{
				jsonrpc.WithAddress(string([]byte{0x01})),
				jsonrpc.WithTimeout(5 * time.Second),
			},
			err: "failed to confirm node connection: failed to fetch network ID: rpc call net_version() on http://\x01: parse \"http://\\x01\": net/url: invalid control character in URL",
		},
		{
			name: "Good",
			parameters: []jsonrpc.Parameter{
				jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
				jsonrpc.WithTimeout(5 * time.Second),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := jsonrpc.New(ctx, test.parameters...)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInterfaces(t *testing.T) {
	ctx := context.Background()
	s, err := jsonrpc.New(ctx, jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")), jsonrpc.WithTimeout(5*time.Second))
	require.NoError(t, err)

	assert.Implements(t, (*client.NetworkIDProvider)(nil), s)
}
