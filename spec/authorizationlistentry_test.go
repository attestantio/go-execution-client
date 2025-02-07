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

func TestAuthorizationListEntryJSON(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type spec.authorizationListEntryJSON",
		},
		{
			name:  "ChainIDMissing",
			input: []byte(`{"address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "chain id missing",
		},
		{
			name:  "ChainIDWrongType",
			input: []byte(`{"chainId":true,"address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field authorizationListEntryJSON.chainId of type string",
		},
		{
			name:  "ChainIDInvalid",
			input: []byte(`{"chainId":"true","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "chain id invalid",
		},
		{
			name:  "AddressMissing",
			input: []byte(`{"chainId":"0x1a5887710","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "address missing",
		},
		{
			name:  "AddressWrongType",
			input: []byte(`{"chainId":"0x1a5887710","address":true,"nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field authorizationListEntryJSON.address of type string",
		},
		{
			name:  "AddressInvalid",
			input: []byte(`{"chainId":"0x1a5887710","address":"true","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "address invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "NonceMissing",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "nonce missing",
		},
		{
			name:  "NonceWrongType",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":true,"yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field authorizationListEntryJSON.nonce of type string",
		},
		{
			name:  "NonceInvalid",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"true","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "nonce invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "YParityMissing",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "yParity missing",
		},
		{
			name:  "YParityWrongType",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":true,"r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field authorizationListEntryJSON.yParity of type string",
		},
		{
			name:  "YParityInvalid",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"true","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "yParity invalid",
		},
		{
			name:  "RMissing",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "r missing",
		},
		{
			name:  "RWrongType",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":true,"s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field authorizationListEntryJSON.r of type string",
		},
		{
			name:  "RInvalid",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":"true","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
			err:   "r invalid",
		},
		{
			name:  "SMissing",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd"}`),
			err:   "s missing",
		},
		{
			name:  "SWrongType",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":true}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field authorizationListEntryJSON.s of type string",
		},
		{
			name:  "SInvalid",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"true"}`),
			err:   "s invalid",
		},
		{
			name:  "Good",
			input: []byte(`{"chainId":"0x1a5887710","address":"0xfc278435313cc88059bf72b92952763e7fc511fb","nonce":"0x0","yParity":"0x0","r":"0xb46862488912b8b80cf1b962dac36ca81f58f7fb32b5365729e96fb73e3048cd","s":"0x6783eadde8364bb04c9c758c0c4cb6481de01330d7a7114ff8a00227a65133c7"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.AuthorizationListEntry
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
