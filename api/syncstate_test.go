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

func TestSyncState(t *testing.T) {
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
			err:   "invalid sync state JSON \"[]\"",
		},
		{
			name:     "GoodNotSyncing",
			input:    []byte(`false`),
			expected: []byte(`{"syncing":false}`),
		},
		{
			name:  "StartingBlockInvalid",
			input: []byte(`{"startingBlock":"true","currentBlock":"0x123","highestBlock":"0x2345"}`),
			err:   "starting block invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "CurrentBlockInvalid",
			input: []byte(`{"startingBlock":"0x2","currentBlock":"true","highestBlock":"0x2345"}`),
			err:   "current block invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "HighestBlockInvalid",
			input: []byte(`{"startingBlock":"0x2","currentBlock":"0x123","highestBlock":"true"}`),
			err:   "highest block invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:     "GoodSyncing",
			input:    []byte(`{"startingBlock":"0x2","currentBlock":"0x123","highestBlock":"0x2345"}`),
			expected: []byte(`{"currentBlock":"0x123","highestBlock":"0x2345","startingBlock":"0x2","syncing":true}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res api.SyncState
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
