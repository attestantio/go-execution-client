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

// TestEventsFilterJSON tests JSON for EventsFilter.
func TestEventsFilterJSON(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type api.eventsFilterJSON",
		},
		{
			name:  "FromBlockInvalid",
			input: []byte(`{"fromBlock":"true","toBlock":"0x7d0","address":"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be","topics":["0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f"]}`),
			err:   "from block invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "ToBlockInvalid",
			input: []byte(`{"fromBlock":"0x3e8","toBlock":"true","address":"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be","topics":["0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f"]}`),
			err:   "to block invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "AddressInvalid",
			input: []byte(`{"fromBlock":"0x3e8","toBlock":"0x7d0","address":"true","topics":["0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f"]}`),
			err:   "address invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "TopicsInvalid",
			input: []byte(`{"fromBlock":"0x3e8","toBlock":"0x7d0","address":"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be","topics":["true"]}`),
			err:   "topic invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "Good",
			input: []byte(`{"fromBlock":"0x3e8","toBlock":"0x7d0","address":"0xa700f2b3d8ebe35cef86fcc3c2105daff41617be","topics":["0x060ac38f43eaa9a6ea5d69fb296993be6072bafed76748ba5e27dd187da0b70f"]}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res api.EventsFilter
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
