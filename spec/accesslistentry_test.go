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

package spec_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/require"
)

func TestAccessListEntryJSON(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   string
	}{
		{
			name: "Empty",
			err:  "unexpected end of JSON input",
		},
		{
			name:  "JSONBad",
			input: []byte("[]"),
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type spec.accessListEntryJSON",
		},
		{
			name:  "AddressMissing",
			input: []byte(`{"storageKeys":["0x0000000000000000000000000000000000000000000000000000000000000008","0x0000000000000000000000000000000000000000000000000000000000000009","0x0000000000000000000000000000000000000000000000000000000000000007","0x000000000000000000000000000000000000000000000000000000000000000a","0x000000000000000000000000000000000000000000000000000000000000000c","0x0000000000000000000000000000000000000000000000000000000000000006"]}`),
			err:   "address missing",
		},
		{
			name:  "AddressWrongType",
			input: []byte(`{"address":true,"storageKeys":["0x0000000000000000000000000000000000000000000000000000000000000008","0x0000000000000000000000000000000000000000000000000000000000000009","0x0000000000000000000000000000000000000000000000000000000000000007","0x000000000000000000000000000000000000000000000000000000000000000a","0x000000000000000000000000000000000000000000000000000000000000000c","0x0000000000000000000000000000000000000000000000000000000000000006"]}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field accessListEntryJSON.address of type string",
		},
		{
			name:  "AddressInvalid",
			input: []byte(`{"address":"true","storageKeys":["0x0000000000000000000000000000000000000000000000000000000000000008","0x0000000000000000000000000000000000000000000000000000000000000009","0x0000000000000000000000000000000000000000000000000000000000000007","0x000000000000000000000000000000000000000000000000000000000000000a","0x000000000000000000000000000000000000000000000000000000000000000c","0x0000000000000000000000000000000000000000000000000000000000000006"]}`),
			err:   "address invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "StorageKeysWrongType",
			input: []byte(`{"address":"0xceff51756c56ceffca006cd410b03ffc46dd3a58","storageKeys":true}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field accessListEntryJSON.storageKeys of type []string",
		},
		{
			name:  "StorageKeyWrongType",
			input: []byte(`{"address":"0xceff51756c56ceffca006cd410b03ffc46dd3a58","storageKeys":[true]}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field accessListEntryJSON.storageKeys of type string",
		},
		{
			name:  "StorageKeyInvalid",
			input: []byte(`{"address":"0xceff51756c56ceffca006cd410b03ffc46dd3a58","storageKeys":["true"]}`),
			err:   "storage key invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "Good",
			input: []byte(`{"address":"0xceff51756c56ceffca006cd410b03ffc46dd3a58","storageKeys":["0x0000000000000000000000000000000000000000000000000000000000000008","0x0000000000000000000000000000000000000000000000000000000000000009","0x0000000000000000000000000000000000000000000000000000000000000007","0x000000000000000000000000000000000000000000000000000000000000000a","0x000000000000000000000000000000000000000000000000000000000000000c","0x0000000000000000000000000000000000000000000000000000000000000006"]}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.AccessListEntry
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				require.Equal(t, string(test.input), string(rt))
			}
		})
	}
}

func TestAccessListEntryYAML(t *testing.T) {
	input := `{"address":"0xceff51756c56ceffca006cd410b03ffc46dd3a58","storageKeys":["0x0000000000000000000000000000000000000000000000000000000000000008","0x0000000000000000000000000000000000000000000000000000000000000009","0x0000000000000000000000000000000000000000000000000000000000000007","0x000000000000000000000000000000000000000000000000000000000000000a","0x000000000000000000000000000000000000000000000000000000000000000c","0x0000000000000000000000000000000000000000000000000000000000000006"]}`
	expected := `{address: '0xceff51756c56ceffca006cd410b03ffc46dd3a58', storageKeys: ['0x0000000000000000000000000000000000000000000000000000000000000008', '0x0000000000000000000000000000000000000000000000000000000000000009', '0x0000000000000000000000000000000000000000000000000000000000000007', '0x000000000000000000000000000000000000000000000000000000000000000a', '0x000000000000000000000000000000000000000000000000000000000000000c', '0x0000000000000000000000000000000000000000000000000000000000000006']}`

	var res spec.AccessListEntry
	require.NoError(t, yaml.Unmarshal([]byte(input), &res))
	require.Equal(t, expected, res.String())
}
