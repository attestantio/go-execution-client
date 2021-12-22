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

package api_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-execution-client/api"
	"github.com/stretchr/testify/require"
)

func TestTransactionResult(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
		err      string
	}{
		{
			name: "Empty",
			err:  "unexpected end of JSON input",
		},
		{
			name:  "JSONBad",
			input: []byte("[]"),
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type api.transactionResultJSON",
		},
		{
			name:  "OutputInvalid",
			input: []byte(`{"output":"true","stateDiff":{"0x4581be18cd63c562cd28f4fb82ac6a4e51f7b93f":{"balance":{"*":{"from":"0x278f482045b836c3","to":"0x2769512f672330ed"}},"code":"=","nonce":{"*":{"from":"0xf4","to":"0xf5"}},"storage":{}},"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be":{"balance":"=","code":"=","nonce":"=","storage":{"0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f":{"*":{"from":"0x0000000000000000000000000000000000000000000000000000000000000000","to":"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"}}}},"0xea674fdde714fd979de3edf0f56aa9716b898ec8":{"balance":{"*":{"from":"0x361b7b53e39f6d12cc","to":"0x361b7b937f90ad50cc"}},"code":"=","nonce":"=","storage":{}}},"trace":[],"vmTrace":null,"transactionHash":"0xb6efb3d07c9c5f193903b75784dcb5e2dfc18ed35b1a61e7a08c24c9f5a50ea1"}`),
			err:   "output invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "StateDiffInvalid",
			input: []byte(`{"output":"0x0000000000000000000000000000000000000000000000000000000000000001","stateDiff":{"true"},"trace":[],"vmTrace":null,"transactionHash":"0xb6efb3d07c9c5f193903b75784dcb5e2dfc18ed35b1a61e7a08c24c9f5a50ea1"}`),
			err:   "invalid character '}' after object key",
		},
		{
			name:  "StateDiffKeyInvalid",
			input: []byte(`{"output":"0x0000000000000000000000000000000000000000000000000000000000000001","stateDiff":{"true":{"balance":{"*":{"from":"0x278f482045b836c3","to":"0x2769512f672330ed"}},"code":"=","nonce":{"*":{"from":"0xf4","to":"0xf5"}},"storage":{}},"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be":{"balance":"=","code":"=","nonce":"=","storage":{"0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f":{"*":{"from":"0x0000000000000000000000000000000000000000000000000000000000000000","to":"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"}}}},"0xea674fdde714fd979de3edf0f56aa9716b898ec8":{"balance":{"*":{"from":"0x361b7b53e39f6d12cc","to":"0x361b7b937f90ad50cc"}},"code":"=","nonce":"=","storage":{}}},"trace":[],"vmTrace":null,"transactionHash":"0xb6efb3d07c9c5f193903b75784dcb5e2dfc18ed35b1a61e7a08c24c9f5a50ea1"}`),
			err:   "address invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "TransactionHashInvalid",
			input: []byte(`{"output":"0x0000000000000000000000000000000000000000000000000000000000000001","stateDiff":{"0x4581be18cd63c562cd28f4fb82ac6a4e51f7b93f":{"balance":{"*":{"from":"0x278f482045b836c3","to":"0x2769512f672330ed"}},"code":"=","nonce":{"*":{"from":"0xf4","to":"0xf5"}},"storage":{}},"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be":{"balance":"=","code":"=","nonce":"=","storage":{"0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f":{"*":{"from":"0x0000000000000000000000000000000000000000000000000000000000000000","to":"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"}}}},"0xea674fdde714fd979de3edf0f56aa9716b898ec8":{"balance":{"*":{"from":"0x361b7b53e39f6d12cc","to":"0x361b7b937f90ad50cc"}},"code":"=","nonce":"=","storage":{}}},"trace":[],"vmTrace":null,"transactionHash":"true"}`),
			err:   "transaction hash invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:     "Good",
			input:    []byte(`{"output":"0x0000000000000000000000000000000000000000000000000000000000000001","stateDiff":{"0x4581be18cd63c562cd28f4fb82ac6a4e51f7b93f":{"balance":{"*":{"from":"0x278f482045b836c3","to":"0x2769512f672330ed"}},"code":"=","nonce":{"*":{"from":"0xf4","to":"0xf5"}},"storage":{}},"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be":{"balance":"=","code":"=","nonce":"=","storage":{"0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f":{"*":{"from":"0x0000000000000000000000000000000000000000000000000000000000000000","to":"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"}}}},"0xea674fdde714fd979de3edf0f56aa9716b898ec8":{"balance":{"*":{"from":"0x361b7b53e39f6d12cc","to":"0x361b7b937f90ad50cc"}},"code":"=","nonce":"=","storage":{}}},"trace":[],"vmTrace":null,"transactionHash":"0xb6efb3d07c9c5f193903b75784dcb5e2dfc18ed35b1a61e7a08c24c9f5a50ea1"}`),
			expected: []byte(`{"output":"0x0000000000000000000000000000000000000000000000000000000000000001","stateDiff":{"0x4581be18cd63c562cd28f4fb82ac6a4e51f7b93f":{"balance":{"*":{"from":"0x278f482045b836c3","to":"0x2769512f672330ed"}},"nonce":{"*":{"from":"0xf4","to":"0xf5"}}},"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be":{"storage":{"0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f":{"*":{"from":"0x0000000000000000000000000000000000000000000000000000000000000000","to":"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"}}}},"0xea674fdde714fd979de3edf0f56aa9716b898ec8":{"balance":{"*":{"from":"0x361b7b53e39f6d12cc","to":"0x361b7b937f90ad50cc"}}}},"transactionHash":"0xb6efb3d07c9c5f193903b75784dcb5e2dfc18ed35b1a61e7a08c24c9f5a50ea1"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res api.TransactionResult
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				if test.expected != nil {
					require.Equal(t, string(test.expected), string(rt))
				} else {
					require.Equal(t, string(test.input), string(rt))
				}
				require.Equal(t, string(rt), res.String())
			}
		})
	}
}
