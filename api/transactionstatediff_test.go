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

func TestTransactionStateDiff(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type api.transactionStateDiffJSON",
		},
		{
			name:  "BalanceInvalid",
			input: []byte(`{"balance":{"*":"true"},"code":"=","nonce":{"*":{"from":"0x397","to":"0x398"}},"storage":{}}`),
			err:   "invalid balance JSON: invalid JSON: json: cannot unmarshal string into Go struct field transactionStateChangeJSON.* of type api.transactionStateChangeAlterationJSON",
		},
		{
			name:  "NonceInvalid",
			input: []byte(`{"balance":{"*":{"from":"0x1d53ae02be969d05","to":"0x1d0ddce00eb7500a"}},"code":"=","nonce":{"*":true},"storage":{}}`),
			err:   "invalid nonce JSON: invalid JSON: json: cannot unmarshal bool into Go struct field transactionStateChangeJSON.* of type api.transactionStateChangeAlterationJSON",
		},
		{
			name:  "StorageInvalid",
			input: []byte(`{"balance":{"*":{"from":"0x1d53ae02be969d05","to":"0x1d0ddce00eb7500a"}},"code":"=","nonce":{"*":{"from":"0x397","to":"0x398"}},"storage":{"true":"true"}}`),
			err:   "invalid storage JSON: invalid JSON: json: cannot unmarshal string into Go value of type api.transactionStorageChangeJSON",
		},
		{
			name:  "StorageKeyInvalid",
			input: []byte(`{"balance":{"*":{"from":"0x1d53ae02be969d05","to":"0x1d0ddce00eb7500a"}},"code":"=","nonce":{"*":{"from":"0x397","to":"0x398"}},"storage":{"true":{}}}`),
			err:   "storage key invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:     "Good",
			input:    []byte(`{"balance":{"*":{"from":"0x1d53ae02be969d05","to":"0x1d0ddce00eb7500a"}},"code":"=","nonce":{"*":{"from":"0x397","to":"0x398"}},"storage":{}}`),
			expected: []byte(`{"balance":{"*":{"from":"0x1d53ae02be969d05","to":"0x1d0ddce00eb7500a"}},"nonce":{"*":{"from":"0x397","to":"0x398"}}}`),
		},
		{
			name:     "GoodNewBalance",
			input:    []byte(`{"balance":{"+":"0x1d53ae02be969d05"},"code":"=","nonce":{"+":"0x0"}}`),
			expected: []byte(`{"balance":{"+":"0x1d53ae02be969d05"},"nonce":{"+":"0x0"}}`),
		},
		{
			name:     "GoodNoChanges",
			input:    []byte(`{"balance":"=","code":"=","nonce":"="}`),
			expected: []byte(`{}`),
		},
		{
			name:     "GoodSelfDestruct",
			input:    []byte(`{"balance":{"*":{"from":"0x1","to":null}},"code":{"*":{"from":"0x608060405260","to":null}},"nonce":{"*":{"from":"0x1","to":null}}}`),
			expected: []byte(`{"balance":{"-":"0x1"},"nonce":{"-":"0x1"}}`),
		},
		{
			name:     "GoodSelfDestructCanonical",
			input:    []byte(`{"balance":{"-":"0x1"},"code":{"-":"0x608060405260"},"nonce":{"-":"0x1"}}`),
			expected: []byte(`{"balance":{"-":"0x1"},"nonce":{"-":"0x1"}}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res api.TransactionStateDiff
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
