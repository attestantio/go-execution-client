// Copyright © 2025 Attestant Limited.
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
	"github.com/stretchr/testify/require"
)

func TestType4TransactionJSON(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type spec.type4TransactionJSON",
		},
		{
			name:  "UnknownType",
			input: []byte(`{"type":"0xff"}`),
			err:   "type incorrect",
		},
		{
			name:  "Good",
			input: []byte(`{"accessList":[],"authorizationList":[{"address":"0x0a2df3ebaa5c4348c82c33412267bf91181a8670","chainId":"0x1","nonce":"0x0","r":"0xa44cc0481db7b192cbef80f9263a5f4b33d948cb4742d53ba0082f8d4d41e86c","s":"0x1847383e8224715b34807b7b46c548836177b71fec8401bd104d6e6db17039ee","yParity":"0x0"}],"blockHash":"0x57ffaa3a25f9139a69289348c679c342fbd56a282bd974110f23263c5825d9e1","blockNumber":"0x3f1","chainId":"0x1a5887710","from":"0x4c207c6a8d044dfd07a00c5c7c3196be6d23460f","gas":"0x41333","gasPrice":"0x3b9aca00","hash":"0x7a52896b992926914e65a03c3cfc99c19a88b31d498da78d04d1d6267247d304","input":"0x","maxFeePerGas":"0x3b9aca00","maxPriorityFeePerGas":"0x3b9aca00","nonce":"0x0","r":"0xea2a6a055f7c8a3247b0911e5d452849685a324edc506cbbccd3bc2ff5c8ee12","s":"0x1d95f7946803e57abef33f0db45faaa84a1db6a8321c06220b830def31a37928","to":"0xd1e6bd74dff7153072cb0d0669f2469d203bf771","transactionIndex":"0x0","type":"0x4","v":"0x1","value":"0x0","yParity":"0x1"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.Type4Transaction
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				require.JSONEq(t, string(test.input), string(rt))
				require.JSONEq(t, string(test.input), res.String())
			}
		})
	}
}
