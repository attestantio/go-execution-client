// Copyright © 2021, 2022 Attestant Limited.
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
	"math/big"
	"testing"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/stretchr/testify/require"
)

// TestType0TransactionJSON tests the JSON encoding of transactions.
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
			name:  "BlockHashWrongType",
			input: []byte(`{"blockHash":true,"blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.blockHash of type string",
		},
		{
			name:  "BlockHashInvalid",
			input: []byte(`{"blockHash":"true","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "block hash invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "BlockHashWithoutNumber",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "block number missing",
		},
		{
			name:  "BlockNumberWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":true,"from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.blockNumber of type string",
		},
		{
			name:  "BlockNumberInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"true","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "block number invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "FromMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "from missing",
		},
		{
			name:  "FromWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":true,"gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.from of type string",
		},
		{
			name:  "FromInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"true","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "from invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "GasMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "gas missing",
		},
		{
			name:  "GasWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":true,"gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.gas of type string",
		},
		{
			name:  "GasInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"true","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "gas invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "GasPriceMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "gas price missing",
		},
		{
			name:  "GasPriceWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":true,"hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.gasPrice of type string",
		},
		{
			name:  "GasPriceInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"true","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "gas price invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "HashMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "hash missing",
		},
		{
			name:  "HashWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":true,"input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.hash of type string",
		},
		{
			name:  "HashInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"true","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "hash invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "InputWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":true,"nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.input of type string",
		},
		{
			name:  "InputInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"true","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "input invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "NonceMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "nonce missing",
		},
		{
			name:  "NonceWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":true,"r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.nonce of type string",
		},
		{
			name:  "NonceInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"true","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "nonce invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "RMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "r missing",
		},
		{
			name:  "RWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":true,"s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.r of type string",
		},
		{
			name:  "RInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"true","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "r invalid",
		},
		{
			name:  "SMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "s missing",
		},
		{
			name:  "SWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":true,"to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.s of type string",
		},
		{
			name:  "SInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"true","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "s invalid",
		},
		{
			name:  "ToWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":true,"transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.to of type string",
		},
		{
			name:  "ToInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"true","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "to invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "TransactionIndexWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":true,"type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.transactionIndex of type string",
		},
		{
			name:  "TransactionIndexInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"true","type":"0x0","v":"0x1c","value":"0x7a69"}`),
			err:   "transaction index invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "TypeWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":true,"v":"0x1c","value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.type of type string",
		},
		{
			name:  "TypeInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"true","v":"0x1c","value":"0x7a69"}`),
			err:   "type incorrect",
		},
		{
			name:  "VMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","value":"0x7a69"}`),
			err:   "v missing",
		},
		{
			name:  "VWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":true,"value":"0x7a69"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.v of type string",
		},
		{
			name:  "VInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"true","value":"0x7a69"}`),
			err:   "v invalid",
		},
		{
			name:  "ValueMissing",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c"}`),
			err:   "value missing",
		},
		{
			name:  "ValueWrongType",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":true}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type0TransactionJSON.value of type string",
		},
		{
			name:  "ValueInvalid",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"true"}`),
			err:   "value invalid",
		},
		{
			name:  "Good",
			input: []byte(`{"blockHash":"0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd","blockNumber":"0xb443","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gas":"0x5208","gasPrice":"0x2d79883d2000","hash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060","input":"0x","nonce":"0x0","r":"0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0","s":"0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a","to":"0x5df9b87991262f6ba471f09758cde1c0fc1de734","transactionIndex":"0x0","type":"0x0","v":"0x1c","value":"0x7a69"}`),
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
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				require.Equal(t, string(test.input), string(rt))
				require.Equal(t, string(test.input), res.String())
			}
		})
	}
}

// TestType0TransactionRLP tests the RLP encoding of transactions.
func TestType0TransactionRLP(t *testing.T) {
	tests := []struct {
		name     string
		input    *spec.Type0Transaction
		expected []byte
		err      string
	}{
		{
			name: "Transfer",
			input: &spec.Type0Transaction{
				Nonce:    124,
				Gas:      21000,
				GasPrice: 71026000000,
				To:       address("0x9ad4c3844d43b21b1ab46a1c13fc9a935211b24b"),
				Value:    big.NewInt(4261560000000000),
				V:        new(big.Int).SetBytes(byteslice("0x25")),
				R:        new(big.Int).SetBytes(byteslice("0x716cce912eb8d2127408b2aefc083f0bd6f4dcee3b04b7db00cf82f93741e927")),
				S:        new(big.Int).SetBytes(byteslice("0x36770cdf4d54e67635aeafff71c5d2ce6b05294cc3c69ca4d161de0bf5a4224c")),
			},
			expected: byteslice("0xf86b7c8510897ac080825208949ad4c3844d43b21b1ab46a1c13fc9a935211b24b870f23ddc1fd30008025a0716cce912eb8d2127408b2aefc083f0bd6f4dcee3b04b7db00cf82f93741e927a036770cdf4d54e67635aeafff71c5d2ce6b05294cc3c69ca4d161de0bf5a4224c"),
		},
		{
			name: "ContractCreation",
			input: &spec.Type0Transaction{
				Nonce:    570212,
				Gas:      1600000,
				GasPrice: 182000000000,
				Value:    big.NewInt(0),
				Input:    byteslice("0x608060405234801561001057600080fd5b5060008054600160a060020a031916331790556102c7806100326000396000f30060806040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166344439209811461006e5780638da5cb5b1461008f578063d0679d34146100c0578063e6d66ac8146100e4578063e8edc8161461010e575b005b34801561007a57600080fd5b5061006c600160a060020a0360043516610123565b34801561009b57600080fd5b506100a4610172565b60408051600160a060020a039092168252519081900360200190f35b3480156100cc57600080fd5b5061006c600160a060020a036004351660243561018a565b3480156100f057600080fd5b5061006c600160a060020a03600435811690602435166044356101dc565b34801561011a57600080fd5b506100a461028c565b337365b0bf8ee4947edd2a500d74e50a3d757dc79de01461014357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b7365b0bf8ee4947edd2a500d74e50a3d757dc79de081565b600054600160a060020a031633146101a157600080fd5b604051600160a060020a0383169082156108fc029083906000818181858888f193505050501580156101d7573d6000803e3d6000fd5b505050565b600054600160a060020a031633146101f357600080fd5b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b15801561026f57600080fd5b505af1158015610283573d6000803e3d6000fd5b50505050505050565b600054600160a060020a0316815600a165627a7a7230582005c585170eb1ba497a4e0bc053a662a46f16fd200c85c37e4f8319d8ca9e93ab0029"),
				V:        new(big.Int).SetBytes(byteslice("0x25")),
				R:        new(big.Int).SetBytes(byteslice("0xa26334de11adf0f03baf0686ec6aa86e7a1cf9b9bf1b93d450a589d22cbfaae4")),
				S:        new(big.Int).SetBytes(byteslice("0x5aa1ed308d32aa0b266b57adc98cff3272dde08d13550b3dcd6920ca869bcb7a")),
			},
			expected: byteslice("0xf9034f8308b364852a600b9c0083186a008080b902f9608060405234801561001057600080fd5b5060008054600160a060020a031916331790556102c7806100326000396000f30060806040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166344439209811461006e5780638da5cb5b1461008f578063d0679d34146100c0578063e6d66ac8146100e4578063e8edc8161461010e575b005b34801561007a57600080fd5b5061006c600160a060020a0360043516610123565b34801561009b57600080fd5b506100a4610172565b60408051600160a060020a039092168252519081900360200190f35b3480156100cc57600080fd5b5061006c600160a060020a036004351660243561018a565b3480156100f057600080fd5b5061006c600160a060020a03600435811690602435166044356101dc565b34801561011a57600080fd5b506100a461028c565b337365b0bf8ee4947edd2a500d74e50a3d757dc79de01461014357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b7365b0bf8ee4947edd2a500d74e50a3d757dc79de081565b600054600160a060020a031633146101a157600080fd5b604051600160a060020a0383169082156108fc029083906000818181858888f193505050501580156101d7573d6000803e3d6000fd5b505050565b600054600160a060020a031633146101f357600080fd5b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b15801561026f57600080fd5b505af1158015610283573d6000803e3d6000fd5b50505050505050565b600054600160a060020a0316815600a165627a7a7230582005c585170eb1ba497a4e0bc053a662a46f16fd200c85c37e4f8319d8ca9e93ab002925a0a26334de11adf0f03baf0686ec6aa86e7a1cf9b9bf1b93d450a589d22cbfaae4a05aa1ed308d32aa0b266b57adc98cff3272dde08d13550b3dcd6920ca869bcb7a"),
		},
		{
			name: "Function",
			input: &spec.Type0Transaction{
				Nonce:    24,
				Gas:      50719,
				GasPrice: 93705977816,
				To:       address("0xdac17f958d2ee523a2206206994597c13d831ec7"),
				Value:    big.NewInt(0),
				Input:    byteslice("0xa9059cbb0000000000000000000000009a58e10992dbf47c7d222c19843a7b11af3f380500000000000000000000000000000000000000000000000000000004a817c800"),
				V:        new(big.Int).SetBytes(byteslice("0x26")),
				R:        new(big.Int).SetBytes(byteslice("0xa72c2173d25ea9490436b30f8371310ac0cb5d2073e92e71336edbffa7c2e526")),
				S:        new(big.Int).SetBytes(byteslice("349c8c5be3d1c1a4f82519db636262f932bf953ea8d7276dd33ccdee71a070af")),
			},
			expected: byteslice("0xf8a9188515d14fbfd882c61f94dac17f958d2ee523a2206206994597c13d831ec780b844a9059cbb0000000000000000000000009a58e10992dbf47c7d222c19843a7b11af3f380500000000000000000000000000000000000000000000000000000004a817c80026a0a72c2173d25ea9490436b30f8371310ac0cb5d2073e92e71336edbffa7c2e526a0349c8c5be3d1c1a4f82519db636262f932bf953ea8d7276dd33ccdee71a070af"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rlp, err := test.input.MarshalRLP()
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, rlp)
			}
		})
	}
}
