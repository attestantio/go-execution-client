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

// TestIssuanceJSON tests JSON for Issuance.
func TestIssuanceJSON(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type api.issuanceJSON",
		},
		{
			name:     "AllEmpty",
			input:    []byte("{}"),
			expected: []byte(`{"blockReward":"0x0","uncleReward":"0x0","issuance":"0x0"}`),
		},
		{
			name:  "BlockRewardInvalid",
			input: []byte(`{"blockReward":"true","uncleReward":"0x234","issuance":"0x357"}`),
			err:   "block reward invalid",
		},
		{
			name:  "UncleRewardInvalid",
			input: []byte(`{"blockReward":"0x123","uncleReward":"true","issuance":"0x357"}`),
			err:   "uncle reward invalid",
		},
		{
			name:  "IssuanceInvalid",
			input: []byte(`{"blockReward":"0x123","uncleReward":"0x234","issuance":"true"}`),
			err:   "issuance invalid",
		},
		{
			name:  "Good",
			input: []byte(`{"blockReward":"0x123","uncleReward":"0x234","issuance":"0x357"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res api.Issuance
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
