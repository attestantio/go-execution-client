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

func TestType0TransactionJSON(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type spec.type0TransactionJSON",
		},
		{
			name:  "BlockHashMissing",
			input: []byte(`{"blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "block hash missing",
		},
		{
			name:  "BlockHashWrongType",
			input: []byte(`{"blockHash":true,"blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.blockHash of type string",
		},
		{
			name:  "BlockHashInvalid",
			input: []byte(`{"blockHash":"true","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "block hash invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "BlockNumberMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "block number missing",
		},
		{
			name:  "BlockNumberWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":true,"from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.blockNumber of type string",
		},
		{
			name:  "BlockNumberInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"true","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "block number invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "FromMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "from missing",
		},
		{
			name:  "FromWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":true,"gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.from of type string",
		},
		{
			name:  "FromInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"true","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "from invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "GasMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "gas missing",
		},
		{
			name:  "GasWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":true,"gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.gas of type string",
		},
		{
			name:  "GasInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"true","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "gas invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "GasPriceMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "gas price missing",
		},
		{
			name:  "GasPriceWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":true,"hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.gasPrice of type string",
		},
		{
			name:  "GasPriceInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"true","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "gas price invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "HashMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "hash missing",
		},
		{
			name:  "HashWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":true,"input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.hash of type string",
		},
		{
			name:  "HashInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"true","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "hash invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "InputWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":true,"nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.input of type string",
		},
		{
			name:  "InputInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"true","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "input invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "NonceMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "nonce missing",
		},
		{
			name:  "NonceWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":true,"to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.nonce of type string",
		},
		{
			name:  "NonceInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"true","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "nonce invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "ToWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":true,"transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.to of type string",
		},
		{
			name:  "ToInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"true","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "to invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "TransactionIndexMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "transaction index missing",
		},
		{
			name:  "TransactionIndexWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":true,"type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.transactionIndex of type string",
		},
		{
			name:  "TransactionIndexInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"true","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "transaction index invalid",
		},
		{
			name:  "ValueMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "value missing",
		},
		{
			name:  "ValueWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":true,"v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.value of type string",
		},
		{
			name:  "ValueInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"true","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "value invalid",
		},
		{
			name:  "VMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "v missing",
		},
		{
			name:  "VWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":true,"r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.v of type string",
		},
		{
			name:  "VInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"true","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "v invalid",
		},
		{
			name:  "RMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "r missing",
		},
		{
			name:  "RWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":true,"s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.r of type string",
		},
		{
			name:  "RInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"true","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e"}`),
			err:   "r invalid",
		},
		{
			name:  "SMissing",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251"}`),
			err:   "s missing",
		},
		{
			name:  "SWrongType",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":true}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.s of type string",
		},
		{
			name:  "SInvalid",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","value":"0x163c900f23774c0","v":"0x26","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"true"}`),
			err:   "s invalid",
		},
		{
			name:  "Good",
			input: []byte(`{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","v":"0x26","value":"0x163c900f23774c0"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.Type0Transaction
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, uint64(0), res.Type())
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				require.Equal(t, string(test.input), string(rt))
			}
		})
	}
}

func TestType0TransactionYAML(t *testing.T) {
	input := `{"blockHash":"0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1","blockNumber":"0xcf635b","from":"0x5c9261660637d09fde3f0d209b8acb79cf3e5124","gas":"0x5208","gasPrice":"0x735c01ab4d","hash":"0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1","input":"0x","nonce":"0x1","r":"0xc5073b102212f0285a7b0eb4b4bb402dd39f43b48f064499b58b71cd83ef9251","s":"0x525631949ac04b22f7ec255ad08c0f6439d15ea8f3520b32680417c9bad3878e","to":"0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2","transactionIndex":"0x2","type":"0x0","v":"0x26","value":"0x163c900f23774c0"}`
	expected := `{blockHash: '0x52d17ecd8dc98f16b97af3ab9e8d0aa07118b7beccb74dedd57c76b41b02c9a1', blockNumber: 13591387, from: '0x5c9261660637d09fde3f0d209b8acb79cf3e5124', gas: 21000, gasPrice: 495464852301, hash: '0xad4dcbc47172e012d40dddb05833805e52e6d2573db07e6ce3b4a07389be37b1', input: '', nonce: 1, r: 89118406738338943123197799493733697242572237422641854775350605826277816373841, s: 37241944623339082250440062366134361643821189374682197041789678185213238413198, to: '0x2faf487a4414fe77e2327f0bf4ae2a264a776ad2', transactionIndex: 2, type: 0, v: 38, value: 100144622633186496}`

	var res spec.Type0Transaction
	require.NoError(t, yaml.Unmarshal([]byte(input), &res))
	require.Equal(t, expected, res.String())
}
