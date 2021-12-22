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

func TestTransactionStorageChange(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type api.transactionStorageChangeJSON",
		},
		{
			name:  "CreationInvalid",
			input: []byte(`{"+":"true"}`),
			err:   "creation invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "AlterationFromInvalid",
			input: []byte(`{"*":{"from":"true","to":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"}}`),
			err:   "from invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "AlterationToInvalid",
			input: []byte(`{"*":{"from":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f","to":"true"}}`),
			err:   "to invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "DeletionInvalid",
			input: []byte(`{"-":"true"}`),
			err:   "deletion invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "GoodCreation",
			input: []byte(`{"+":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"}`),
		},
		{
			name:  "GoodAlteration",
			input: []byte(`{"*":{"from":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f","to":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"}}`),
		},
		{
			name:  "GoodAlterationToZero",
			input: []byte(`{"*":{"from":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20","to":"0x0000000000000000000000000000000000000000000000000000000000000000"}}`),
		},
		{
			name:  "GoodDeletion",
			input: []byte(`{"-":"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res api.TransactionStorageChange
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
