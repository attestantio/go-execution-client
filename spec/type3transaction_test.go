// Copyright © 2023 Attestant Limited.
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

func TestType3TransactionJSON(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type spec.type3TransactionJSON",
		},
		{
			name:  "UnknownType",
			input: []byte(`{"type":"0xff"}`),
			err:   "type incorrect",
		},
		{
			name:  "AccessListWrongType",
			input: []byte(`{"accessList":true,"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.accessList of type []*spec.AccessListEntry",
		},
		{
			name:  "AccessListEntryEmpty",
			input: []byte(`{"accessList":[""],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: invalid JSON: json: cannot unmarshal string into Go value of type spec.accessListEntryJSON",
		},
		{
			name:  "AccessListEntryWrongType",
			input: []byte(`{"accessList":[true],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: invalid JSON: json: cannot unmarshal bool into Go value of type spec.accessListEntryJSON",
		},
		{
			name:  "AccessListEntryInvalid",
			input: []byte(`{"accessList":["true"],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: invalid JSON: json: cannot unmarshal string into Go value of type spec.accessListEntryJSON",
		},
		{
			name:  "BlobGasPriceWrongType",
			input: []byte(`{"accessList":[],"blobGasPrice":true,"blobGasUsed":"0x20000","blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.blobGasPrice of type string",
		},
		{
			name:  "BlobGasPriceInvalid",
			input: []byte(`{"accessList":[],"blobGasPrice":"true","blobGasUsed":"0x20000","blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "blob gas price invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "BlobGasUsedWrongType",
			input: []byte(`{"accessList":[],"blobGasPrice":"0x1","blobGasUsed":true,"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.blobGasUsed of type string",
		},
		{
			name:  "BlobGasUsedInvalid",
			input: []byte(`{"accessList":[],"blobGasPrice":"0x1","blobGasUsed":"true","blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "blob gas used invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "BlobVersionedHashesWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":[true],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: invalid prefix",
		},
		{
			name:  "BlobVersionedHashWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":[true],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: invalid prefix",
		},
		{
			name:  "BlobVersionedHashInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["true"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: invalid prefix",
		},
		{
			name:  "BlockHashWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":true,"blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.blockHash of type string",
		},
		{
			name:  "BlockHashInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"true","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "block hash invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "BlockHashWithoutNumber",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "block number missing",
		},
		{
			name:  "BlockNumberWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":true,"chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.blockNumber of type string",
		},
		{
			name:  "BlockNumberInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"true","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "block number invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "ChainIDMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "chain id missing",
		},
		{
			name:  "ChainIDWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":true,"from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.chainId of type string",
		},
		{
			name:  "ChainIDInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"true","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "chain id invalid",
		},
		{
			name:  "FromMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "from missing",
		},
		{
			name:  "FromWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":true,"gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.from of type string",
		},
		{
			name:  "FromInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"true","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "from invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "GasMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "gas missing",
		},
		{
			name:  "GasWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":true,"gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.gas of type string",
		},
		{
			name:  "GasInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"true","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "gas invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "GasPriceWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":true,"hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.gasPrice of type string",
		},
		{
			name:  "GasPriceInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"true","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "gas price invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "HashMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "hash missing",
		},
		{
			name:  "HashWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":true,"input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.hash of type string",
		},
		{
			name:  "HashInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"true","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "hash invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "InputWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":true,"maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.input of type string",
		},
		{
			name:  "InputInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"true","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "input invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "MaxFeePerBlobGasMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "max fee per blob gas missing",
		},
		{
			name:  "MaxFeePerBlobGasWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":true,"maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.maxFeePerBlobGas of type string",
		},
		{
			name:  "MaxFeePerBlobGasInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"true","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "max fee per blob gas invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "MaxFeePerGasMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "max fee per gas missing",
		},
		{
			name:  "MaxFeePerGasWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":true,"maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.maxFeePerGas of type string",
		},
		{
			name:  "MaxFeePerGasInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"true","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "max fee per gas invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "MaxPriorityFeePerGasMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "max priority fee per gas missing",
		},
		{
			name:  "MaxPriorityFeePerGasWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":true,"nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.maxPriorityFeePerGas of type string",
		},
		{
			name:  "MaxPriorityFeePerGasInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"true","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "max priority fee per gas invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "NonceMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "nonce missing",
		},
		{
			name:  "NonceWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":true,"r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.nonce of type string",
		},
		{
			name:  "NonceInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"true","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "nonce invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "RMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "r missing",
		},
		{
			name:  "RWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":true,"s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.r of type string",
		},
		{
			name:  "RInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"true","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "r invalid",
		},
		{
			name:  "SMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "s missing",
		},
		{
			name:  "SWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":true,"to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.s of type string",
		},
		{
			name:  "SInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"true","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "s invalid",
		},
		{
			name:  "ToWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":true,"transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.to of type string",
		},
		{
			name:  "ToInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"true","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "to invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "TransactionIndexWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":true,"type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.transactionIndex of type string",
		},
		{
			name:  "TransactionIndexInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"true","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "transaction index invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "TypeMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "type missing for type 3 transaction",
		},
		{
			name:  "TypeWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":true,"v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.type of type string",
		},
		{
			name:  "TypeInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"true","v":"0x0","value":"0x0","yParity":"0x0"}`),
			err:   "type incorrect",
		},
		{
			name:  "YParityAndVMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","value":"0x0"}`),
			err:   "yParity and v missing",
		},
		{
			name:  "VWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":true,"value":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.v of type string",
		},
		{
			name:  "VInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"true","value":"0x0"}`),
			err:   "v invalid",
		},
		{
			name:  "YParityWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":true}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.yParity of type string",
		},
		{
			name:  "YParityInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"true"}`),
			err:   "yParity invalid",
		},
		{
			name:  "ValueMissing",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","yParity":"0x0"}`),
			err:   "value missing",
		},
		{
			name:  "ValueWrongType",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":true,"yParity":"0x0"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field type3TransactionJSON.value of type string",
		},
		{
			name:  "ValueInvalid",
			input: []byte(`{"accessList":[],"blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"true","yParity":"0x0"}`),
			err:   "value invalid",
		},
		{
			name:  "Good",
			input: []byte(`{"accessList":[],"blobGasPrice":"0x1","blobGasUsed":"0x20000","blobVersionedHashes":["0x010a8b6ba3ab54a1119c273b22a15c700be68aba00cc989903f0a53459d6daad"],"blockHash":"0xc313c53462bd60aff8637329ef2c554f3f730e158ad363b4dfb8bb23d3b9747d","blockNumber":"0x520","chainId":"0x1a1f0ff42","from":"0x7e454a14b8e7528465eef86f0dc1da4f235d9d79","gas":"0x5208","gasPrice":"0xee6b281c","hash":"0xa93f6ce94cd78c3344d2f639620e36f327f59bb8d583a65f4d851ce11e36e2d2","input":"0x","maxFeePerBlobGas":"0x3e8","maxFeePerGas":"0xee6b281c","maxPriorityFeePerGas":"0xee6b281c","nonce":"0x1","r":"0x8a8eb78331b968788c11b9c418f6979cc6dec32f4f4ad3ba398fec5b9058f0fa","s":"0x50225cb61b29987a032d6187ae3d4af7fadd6107e17f1adb9ad7f136417d3548","to":"0x000000000000000000000000000000000000f1c1","transactionIndex":"0x0","type":"0x3","v":"0x0","value":"0x0","yParity":"0x0"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.Type3Transaction
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

// TODO.
// // TestType3TransactionRLP tests the RLP encoding of transactions.
// func TestType3TransactionRLP(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    *spec.Type3Transaction
// 		expected []byte
// 		err      string
// 	}{
// 		{
// 			name: "Transfer",
// 			input: &spec.Type3Transaction{
// 				ChainID:              new(big.Int).SetBytes(byteslice("0x01")),
// 				Nonce:                2599,
// 				Gas:                  21000,
// 				MaxPriorityFeePerGas: 1250000000,
// 				MaxFeePerGas:         122135661622,
// 				To:                   address("0xD28085614D0CE92D98FDc1d0cFc10e5fd6da6fbc"),
// 				Value:                big.NewInt(283100000000000000),
// 				V:                    new(big.Int).SetBytes(byteslice("0x01")),
// 				R:                    new(big.Int).SetBytes(byteslice("0x6506a2afd6e0f57b6887d68d089c97adb7af316947e8033af6a29a206fb6bd1d")),
// 				S:                    new(big.Int).SetBytes(byteslice("0x347e78edb68ee405f42b818eabbc736d8d4723a403f7c391ea0f8abf780c32f5")),
// 			},
// 			expected: byteslice("0xb87802f87501820a27844a817c80851c6fda4c3682520894d28085614d0ce92d98fdc1d0cfc10e5fd6da6fbc8803edc5f337e9c00080c001a06506a2afd6e0f57b6887d68d089c97adb7af316947e8033af6a29a206fb6bd1da0347e78edb68ee405f42b818eabbc736d8d4723a403f7c391ea0f8abf780c32f5"),
// 		},
// 		{
// 			name: "ContractCreation",
// 			input: &spec.Type3Transaction{
// 				ChainID:              new(big.Int).SetBytes(byteslice("0x01")),
// 				Nonce:                2,
// 				Gas:                  1006654,
// 				MaxPriorityFeePerGas: 1500000000,
// 				MaxFeePerGas:         111376341985,
// 				Value:                big.NewInt(0),
// 				Input:                byteslice("0x608060405260405162000f1f38038062000f1f83398181016040526101408110156200002a57600080fd5b5080516020808301516040808501516060860151608087015160a088015160c089015160e08a01516101008b0151610120909b015187516001600160a01b03808916602483015280881660448301528087166064830152808616608483015280851660a483015280841660c48301528d1660e4820152610104808201839052895180830390910181526101249091018952998a0180516001600160e01b0316638830191160e01b17905287517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152975197889003601c01909720999a979995989497939692959194909392918b918b918390829060008051602062000e7c833981519152600019909101146200013d57fe5b62000151826001600160e01b03620001eb16565b80511562000172576200017082826200025160201b6200047e1760201c565b505b5050604080517f656970313936372e70726f78792e61646d696e000000000000000000000000008152905190819003601301902060008051602062000e5c83398151915260001990910114620001c457fe5b620001d8826001600160e01b036200028916565b505050505050505050505050506200046c565b62000201816200029c60201b620004aa1760201c565b6200023e5760405162461bcd60e51b815260040180806020018281038252603681526020018062000ec36036913960400191505060405180910390fd5b60008051602062000e7c83398151915255565b606062000282838360405180606001604052806027815260200162000e9c602791396001600160e01b03620002a216565b9392505050565b60008051602062000e5c83398151915255565b3b151590565b6060620002b8846001600160e01b036200029c16565b620002f55760405162461bcd60e51b815260040180806020018281038252602681526020018062000ef96026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b60208310620003355780518252601f19909201916020918201910162000314565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811462000397576040519150601f19603f3d011682016040523d82523d6000602084013e6200039c565b606091505b509092509050620003b88282866001600160e01b03620003c216565b9695505050505050565b60608315620003d357508162000282565b825115620003e45782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200043057818101518382015260200162000416565b50505050905090810190601f1680156200045e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b6109e0806200047c6000396000f3fe60806040526004361061005e5760003560e01c80635c60da1b116100435780635c60da1b146101425780638f28397014610180578063f851a440146101c05761006d565b80633659cfe6146100755780634f1ef286146100b55761006d565b3661006d5761006b6101d5565b005b61006b6101d5565b34801561008157600080fd5b5061006b6004803603602081101561009857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101ef565b61006b600480360360408110156100cb57600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516919081019060408101602082013564010000000081111561010357600080fd5b82018360208201111561011557600080fd5b8035906020019184600183028401116401000000008311171561013757600080fd5b509092509050610243565b34801561014e57600080fd5b506101576102da565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561018c57600080fd5b5061006b600480360360208110156101a357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610331565b3480156101cc57600080fd5b50610157610439565b6101dd6104b0565b6101ed6101e8610544565b610569565b565b6101f761058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561023857610233816105b2565b610240565b6102406101d5565b50565b61024b61058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102cd57610287836105b2565b6102c78383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061047e92505050565b506102d5565b6102d56101d5565b505050565b60006102e461058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103265761031f610544565b905061032e565b61032e6101d5565b90565b61033961058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102385773ffffffffffffffffffffffffffffffffffffffff81166103d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a8152602001806108ac603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61040161058d565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301528051918290030190a1610233816105ff565b600061044361058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103265761031f61058d565b60606104a383836040518060600160405280602781526020016108e660279139610623565b9392505050565b3b151590565b6104b861058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561053c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260428152602001806109696042913960600191505060405180910390fd5b6101ed6101ed565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015610588573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b6105bb8161076b565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b606061062e846104aa565b610683576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806109436026913960400191505060405180910390fd5b600060608573ffffffffffffffffffffffffffffffffffffffff16856040518082805190602001908083835b602083106106ec57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016106af565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461074c576040519150601f19603f3d011682016040523d82523d6000602084013e610751565b606091505b50915091506107618282866107ed565b9695505050505050565b610774816104aa565b6107c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603681526020018061090d6036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b606083156107fc5750816104a3565b82511561080c5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610870578181015183820152602001610858565b50505050905090810190601f16801561089d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe5472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f2061646472657373416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a2646970667358221220ca8b31ef68b5e180f5ea3d8cea3d2ad8c208482f6646937570877e1cbcdaaf1b64736f6c634300060b0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163740000000000000000000000007b9c993a41e0777f5ec9a2db09b8e6c2b344f33f000000000000000000000000e4b245fa9cb539aad125d4849b7c99cc0efcea2b000000000000000000000000b88460bb2696cab9d66013a05dff29a28330689d000000000000000000000000f418588522d5dd018b425e472991e52ebbeeeeee000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d00000000000000000000000024a42fd28c976a61df5d00d0599c34c4f90748c80000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000fc1e690f61efd961294b3e1ce3313fbd8aa4f85d0000000000000000000000000000000000000000000000000000000000000000"),
// 				V:                    new(big.Int).SetBytes(byteslice("0x00")),
// 				R:                    new(big.Int).SetBytes(byteslice("0x10dd5c6d8dbba5a1a19231072d30331972099750f4e979f96f80f9a8c7f8cc18")),
// 				S:                    new(big.Int).SetBytes(byteslice("0x3b7d091437d454553a27c6068ed5bf82592294d0ae27c94086407e0f8c3a8ba9")),
// 			},
// 			expected: byteslice("0xb910bd02f910b901028459682f008519ee8c1be1830f5c3e8080b9105f608060405260405162000f1f38038062000f1f83398181016040526101408110156200002a57600080fd5b5080516020808301516040808501516060860151608087015160a088015160c089015160e08a01516101008b0151610120909b015187516001600160a01b03808916602483015280881660448301528087166064830152808616608483015280851660a483015280841660c48301528d1660e4820152610104808201839052895180830390910181526101249091018952998a0180516001600160e01b0316638830191160e01b17905287517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152975197889003601c01909720999a979995989497939692959194909392918b918b918390829060008051602062000e7c833981519152600019909101146200013d57fe5b62000151826001600160e01b03620001eb16565b80511562000172576200017082826200025160201b6200047e1760201c565b505b5050604080517f656970313936372e70726f78792e61646d696e000000000000000000000000008152905190819003601301902060008051602062000e5c83398151915260001990910114620001c457fe5b620001d8826001600160e01b036200028916565b505050505050505050505050506200046c565b62000201816200029c60201b620004aa1760201c565b6200023e5760405162461bcd60e51b815260040180806020018281038252603681526020018062000ec36036913960400191505060405180910390fd5b60008051602062000e7c83398151915255565b606062000282838360405180606001604052806027815260200162000e9c602791396001600160e01b03620002a216565b9392505050565b60008051602062000e5c83398151915255565b3b151590565b6060620002b8846001600160e01b036200029c16565b620002f55760405162461bcd60e51b815260040180806020018281038252602681526020018062000ef96026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b60208310620003355780518252601f19909201916020918201910162000314565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811462000397576040519150601f19603f3d011682016040523d82523d6000602084013e6200039c565b606091505b509092509050620003b88282866001600160e01b03620003c216565b9695505050505050565b60608315620003d357508162000282565b825115620003e45782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200043057818101518382015260200162000416565b50505050905090810190601f1680156200045e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b6109e0806200047c6000396000f3fe60806040526004361061005e5760003560e01c80635c60da1b116100435780635c60da1b146101425780638f28397014610180578063f851a440146101c05761006d565b80633659cfe6146100755780634f1ef286146100b55761006d565b3661006d5761006b6101d5565b005b61006b6101d5565b34801561008157600080fd5b5061006b6004803603602081101561009857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101ef565b61006b600480360360408110156100cb57600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516919081019060408101602082013564010000000081111561010357600080fd5b82018360208201111561011557600080fd5b8035906020019184600183028401116401000000008311171561013757600080fd5b509092509050610243565b34801561014e57600080fd5b506101576102da565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561018c57600080fd5b5061006b600480360360208110156101a357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610331565b3480156101cc57600080fd5b50610157610439565b6101dd6104b0565b6101ed6101e8610544565b610569565b565b6101f761058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561023857610233816105b2565b610240565b6102406101d5565b50565b61024b61058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102cd57610287836105b2565b6102c78383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061047e92505050565b506102d5565b6102d56101d5565b505050565b60006102e461058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103265761031f610544565b905061032e565b61032e6101d5565b90565b61033961058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102385773ffffffffffffffffffffffffffffffffffffffff81166103d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a8152602001806108ac603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61040161058d565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301528051918290030190a1610233816105ff565b600061044361058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103265761031f61058d565b60606104a383836040518060600160405280602781526020016108e660279139610623565b9392505050565b3b151590565b6104b861058d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561053c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260428152602001806109696042913960600191505060405180910390fd5b6101ed6101ed565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e808015610588573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b6105bb8161076b565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b606061062e846104aa565b610683576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806109436026913960400191505060405180910390fd5b600060608573ffffffffffffffffffffffffffffffffffffffff16856040518082805190602001908083835b602083106106ec57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016106af565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461074c576040519150601f19603f3d011682016040523d82523d6000602084013e610751565b606091505b50915091506107618282866107ed565b9695505050505050565b610774816104aa565b6107c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603681526020018061090d6036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b606083156107fc5750816104a3565b82511561080c5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610870578181015183820152602001610858565b50505050905090810190601f16801561089d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe5472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f2061646472657373416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a2646970667358221220ca8b31ef68b5e180f5ea3d8cea3d2ad8c208482f6646937570877e1cbcdaaf1b64736f6c634300060b0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163740000000000000000000000007b9c993a41e0777f5ec9a2db09b8e6c2b344f33f000000000000000000000000e4b245fa9cb539aad125d4849b7c99cc0efcea2b000000000000000000000000b88460bb2696cab9d66013a05dff29a28330689d000000000000000000000000f418588522d5dd018b425e472991e52ebbeeeeee000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d00000000000000000000000024a42fd28c976a61df5d00d0599c34c4f90748c80000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000fc1e690f61efd961294b3e1ce3313fbd8aa4f85d0000000000000000000000000000000000000000000000000000000000000000c080a010dd5c6d8dbba5a1a19231072d30331972099750f4e979f96f80f9a8c7f8cc18a03b7d091437d454553a27c6068ed5bf82592294d0ae27c94086407e0f8c3a8ba9"),
// 		},
// 		{
// 			name: "Function",
// 			input: &spec.Type3Transaction{
// 				ChainID:              new(big.Int).SetBytes(byteslice("0x01")),
// 				Nonce:                177,
// 				Gas:                  359305,
// 				MaxPriorityFeePerGas: 1500000000,
// 				MaxFeePerGas:         122722795913,
// 				To:                   address("0x7a250d5630b4cf539739df2c5dacb4c659f2488d"),
// 				Value:                big.NewInt(0),
// 				Input:                byteslice("0x5c11d7950000000000000000000000000000000000000000000000008b68177f660420460000000000000000000000000000000000000000000005fe68def8df304461fc00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000b09ebc3ccfde499152c2d1fb423d4d83a9d6f2e60000000000000000000000000000000000000000000000000000000061c9ed6200000000000000000000000000000000000000000000000000000000000000030000000000000000000000008b3192f5eebd8579568a2ed41e6feb402f93f73f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000f2ee18d33bf5f9cd1ccb917c1ccecd8908231247"),
// 				V:                    new(big.Int).SetBytes(byteslice("0x00")),
// 				R:                    new(big.Int).SetBytes(byteslice("0xa01e352eb34f5a831b70e1ec7776ec1e4bb753a0fe45304e8116aa5de3808e86")),
// 				S:                    new(big.Int).SetBytes(byteslice("0x428b2ef19f872ee1d52ee648233ab83ecaa167d1bda6ad5d6a084c6e6e677d56")),
// 			},
// 			expected: byteslice("0xb9019702f901930181b18459682f00851c92d9418983057b89947a250d5630b4cf539739df2c5dacb4c659f2488d80b901245c11d7950000000000000000000000000000000000000000000000008b68177f660420460000000000000000000000000000000000000000000005fe68def8df304461fc00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000b09ebc3ccfde499152c2d1fb423d4d83a9d6f2e60000000000000000000000000000000000000000000000000000000061c9ed6200000000000000000000000000000000000000000000000000000000000000030000000000000000000000008b3192f5eebd8579568a2ed41e6feb402f93f73f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000f2ee18d33bf5f9cd1ccb917c1ccecd8908231247c080a0a01e352eb34f5a831b70e1ec7776ec1e4bb753a0fe45304e8116aa5de3808e86a0428b2ef19f872ee1d52ee648233ab83ecaa167d1bda6ad5d6a084c6e6e677d56"),
// 		},
// 		{
// 			name: "AccessList",
// 			input: &spec.Type3Transaction{
// 				ChainID:              new(big.Int).SetBytes(byteslice("0x01")),
// 				Nonce:                3157,
// 				Gas:                  330030,
// 				MaxPriorityFeePerGas: 0,
// 				MaxFeePerGas:         41715866680,
// 				To:                   address("0x00000000000123685885532dcB685c442Dc83126"),
// 				Value:                big.NewInt(0),
// 				AccessList: []*spec.AccessListEntry{
// 					{
// 						Address: byteslice("0xcf6daab95c476106eca715d48de4b13287ffdeaa"),
// 						StorageKeys: [][]byte{
// 							byteslice("0x000000000000000000000000000000000000000000000000000000000000000a"),
// 							byteslice("0x000000000000000000000000000000000000000000000000000000000000000f"),
// 							byteslice("0x0000000000000000000000000000000000000000000000000000000000000008"),
// 							byteslice("0x0000000000000000000000000000000000000000000000000000000000000006"),
// 							byteslice("0x0000000000000000000000000000000000000000000000000000000000000007"),
// 							byteslice("0x000000000000000000000000000000000000000000000000000000000000000c"),
// 							byteslice("0x0000000000000000000000000000000000000000000000000000000000000009"),
// 						},
// 					},
// 					{
// 						Address: byteslice("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
// 						StorageKeys: [][]byte{
// 							byteslice("0xbb609bf4e7a9c5741724a2ad55eabdc40913c8923f30251e38f2ea5017d9a032"),
// 							byteslice("0xf3ccbe01292f6eaa94374ba731a5fc9a7ced51e4be348a95f9a10feaa1df67b1"),
// 							byteslice("0xbc23203c048bd7d017695d81f6390933f6b9889cc354a109a941dc4b2a9d4799"),
// 						},
// 					},
// 					{
// 						Address: byteslice("0x2f62f2b4c5fcd7570a709dec05d68ea19c82a9ec"),
// 						StorageKeys: [][]byte{
// 							byteslice("0x0000000000000000000000000000000000000000000000000000000000000000"),
// 							byteslice("0x0000000000000000000000000000000000000000000000000000000000000004"),
// 							byteslice("0xad860a26b2adedd5a0c5d198c9503a420ba615f9041c00093858eff051edf0a0"),
// 							byteslice("0x000000000000000000000000000000000000000000000000000000000000002c"),
// 							byteslice("0x704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a7"),
// 							byteslice("0x000000000000000000000000000000000000000000000000000000000000002d"),
// 							byteslice("0x0000000000000000000000000000000000000000000000000000000000000002"),
// 							byteslice("0x0000000000000000000000000000000000000000000000000000000000000001"),
// 							byteslice("0x704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a5"),
// 							byteslice("0x704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a6"),
// 							byteslice("0x704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a4"),
// 						},
// 					},
// 					{
// 						Address: byteslice("0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce"),
// 						StorageKeys: [][]byte{
// 							byteslice("0x0547f50ddc56a26e1c9d9a497b4980b477f9e22a83efd7e8da74c6b128769bb2"),
// 							byteslice("0x8ca43d7619371d0834e06fa31b042aaaa680586ab3652a0f8c98e768b6443765"),
// 						},
// 					},
// 				},
// 				Input: byteslice("0x2976dcef000000000000000000000000cf6daab95c476106eca715d48de4b13287ffdeaa000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000b9850756e368da00000000000000000000000000000000000000000000000000000000000004c4022c0d9f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c4ade23268074c0b00000000000000000000000000000000000123685885532dcb685c442dc83126000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000004200000000000000000000000002f62f2b4c5fcd7570a709dec05d68ea19c82a9ec000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000003a4128acb08000000000000000000000000cf6daab95c476106eca715d48de4b13287ffdeaa0000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffec9c7b27bf045a724415a4f9000000000000000000000000fffd8963efd1fc6a506488495d951d5263988d2500000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000123685885532dcb685c442dc8312600000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000264a7996fa5000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb0000000000000000000000002f62f2b4c5fcd7570a709dec05d68ea19c82a9ec000000000000000000000000000000000000000000000000c3ad434b850200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000242e1a7d4d00000000000000000000000000000000000000000000000001009ee6e3054c0b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
// 				V:     new(big.Int).SetBytes(byteslice("0x01")),
// 				R:     new(big.Int).SetBytes(byteslice("0x3f3fce3c5da8e8c896209057348873d2295a05f725e662f3f0c0230ee637f6a9")),
// 				S:     new(big.Int).SetBytes(byteslice("0x46ad8ff3062b5d33d9bfdaa159d05409c66b1b1deb61247b7976fe8d9bc9f6c5")),
// 			},
// 			expected: byteslice("0xb9093302f9092f01820c55808509b6759c388305092e9400000000000123685885532dcb685c442dc8312680b905642976dcef000000000000000000000000cf6daab95c476106eca715d48de4b13287ffdeaa000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000b9850756e368da00000000000000000000000000000000000000000000000000000000000004c4022c0d9f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c4ade23268074c0b00000000000000000000000000000000000123685885532dcb685c442dc83126000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000004200000000000000000000000002f62f2b4c5fcd7570a709dec05d68ea19c82a9ec000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000003a4128acb08000000000000000000000000cf6daab95c476106eca715d48de4b13287ffdeaa0000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffec9c7b27bf045a724415a4f9000000000000000000000000fffd8963efd1fc6a506488495d951d5263988d2500000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000123685885532dcb685c442dc8312600000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000264a7996fa5000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb0000000000000000000000002f62f2b4c5fcd7570a709dec05d68ea19c82a9ec000000000000000000000000000000000000000000000000c3ad434b850200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000242e1a7d4d00000000000000000000000000000000000000000000000001009ee6e3054c0b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f9035df8fe94cf6daab95c476106eca715d48de4b13287ffdeaaf8e7a0000000000000000000000000000000000000000000000000000000000000000aa0000000000000000000000000000000000000000000000000000000000000000fa00000000000000000000000000000000000000000000000000000000000000008a00000000000000000000000000000000000000000000000000000000000000006a00000000000000000000000000000000000000000000000000000000000000007a0000000000000000000000000000000000000000000000000000000000000000ca00000000000000000000000000000000000000000000000000000000000000009f87a94c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2f863a0bb609bf4e7a9c5741724a2ad55eabdc40913c8923f30251e38f2ea5017d9a032a0f3ccbe01292f6eaa94374ba731a5fc9a7ced51e4be348a95f9a10feaa1df67b1a0bc23203c048bd7d017695d81f6390933f6b9889cc354a109a941dc4b2a9d4799f90183942f62f2b4c5fcd7570a709dec05d68ea19c82a9ecf9016ba00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000004a0ad860a26b2adedd5a0c5d198c9503a420ba615f9041c00093858eff051edf0a0a0000000000000000000000000000000000000000000000000000000000000002ca0704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a7a0000000000000000000000000000000000000000000000000000000000000002da00000000000000000000000000000000000000000000000000000000000000002a00000000000000000000000000000000000000000000000000000000000000001a0704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a5a0704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a6a0704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a4f8599495ad61b0a150d79219dcf64e1e6cc01f0b64c4cef842a00547f50ddc56a26e1c9d9a497b4980b477f9e22a83efd7e8da74c6b128769bb2a08ca43d7619371d0834e06fa31b042aaaa680586ab3652a0f8c98e768b644376501a03f3fce3c5da8e8c896209057348873d2295a05f725e662f3f0c0230ee637f6a9a046ad8ff3062b5d33d9bfdaa159d05409c66b1b1deb61247b7976fe8d9bc9f6c5"),
// 		},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			rlp, err := test.input.MarshalRLP()
// 			if test.err != "" {
// 				require.EqualError(t, err, test.err)
// 			} else {
// 				require.NoError(t, err)
// 				require.Equal(t, test.expected, rlp)
// 			}
// 		})
// 	}
// }
//
// func BenchmarkRLP(b *testing.B) {
// 	tx := &spec.Type3Transaction{
// 		ChainID:              new(big.Int).SetBytes(byteslice("0x01")),
// 		Nonce:                3157,
// 		Gas:                  330030,
// 		MaxPriorityFeePerGas: 0,
// 		MaxFeePerGas:         41715866680,
// 		To:                   address("0x00000000000123685885532dcB685c442Dc83126"),
// 		Value:                big.NewInt(0),
// 		AccessList: []*spec.AccessListEntry{
// 			{
// 				Address: byteslice("0xcf6daab95c476106eca715d48de4b13287ffdeaa"),
// 				StorageKeys: [][]byte{
// 					byteslice("0x000000000000000000000000000000000000000000000000000000000000000a"),
// 					byteslice("0x000000000000000000000000000000000000000000000000000000000000000f"),
// 					byteslice("0x0000000000000000000000000000000000000000000000000000000000000008"),
// 					byteslice("0x0000000000000000000000000000000000000000000000000000000000000006"),
// 					byteslice("0x0000000000000000000000000000000000000000000000000000000000000007"),
// 					byteslice("0x000000000000000000000000000000000000000000000000000000000000000c"),
// 					byteslice("0x0000000000000000000000000000000000000000000000000000000000000009"),
// 				},
// 			},
// 			{
// 				Address: byteslice("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
// 				StorageKeys: [][]byte{
// 					byteslice("0xbb609bf4e7a9c5741724a2ad55eabdc40913c8923f30251e38f2ea5017d9a032"),
// 					byteslice("0xf3ccbe01292f6eaa94374ba731a5fc9a7ced51e4be348a95f9a10feaa1df67b1"),
// 					byteslice("0xbc23203c048bd7d017695d81f6390933f6b9889cc354a109a941dc4b2a9d4799"),
// 				},
// 			},
// 			{
// 				Address: byteslice("0x2f62f2b4c5fcd7570a709dec05d68ea19c82a9ec"),
// 				StorageKeys: [][]byte{
// 					byteslice("0x0000000000000000000000000000000000000000000000000000000000000000"),
// 					byteslice("0x0000000000000000000000000000000000000000000000000000000000000004"),
// 					byteslice("0xad860a26b2adedd5a0c5d198c9503a420ba615f9041c00093858eff051edf0a0"),
// 					byteslice("0x000000000000000000000000000000000000000000000000000000000000002c"),
// 					byteslice("0x704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a7"),
// 					byteslice("0x000000000000000000000000000000000000000000000000000000000000002d"),
// 					byteslice("0x0000000000000000000000000000000000000000000000000000000000000002"),
// 					byteslice("0x0000000000000000000000000000000000000000000000000000000000000001"),
// 					byteslice("0x704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a5"),
// 					byteslice("0x704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a6"),
// 					byteslice("0x704a382402faf574444c81a64959d0b8a34f0f113130828ba8fccc54309419a4"),
// 				},
// 			},
// 			{
// 				Address: byteslice("0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce"),
// 				StorageKeys: [][]byte{
// 					byteslice("0x0547f50ddc56a26e1c9d9a497b4980b477f9e22a83efd7e8da74c6b128769bb2"),
// 					byteslice("0x8ca43d7619371d0834e06fa31b042aaaa680586ab3652a0f8c98e768b6443765"),
// 				},
// 			},
// 		},
// 		Input: byteslice("0x2976dcef000000000000000000000000cf6daab95c476106eca715d48de4b13287ffdeaa000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000b9850756e368da00000000000000000000000000000000000000000000000000000000000004c4022c0d9f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c4ade23268074c0b00000000000000000000000000000000000123685885532dcb685c442dc83126000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000004200000000000000000000000002f62f2b4c5fcd7570a709dec05d68ea19c82a9ec000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000003a4128acb08000000000000000000000000cf6daab95c476106eca715d48de4b13287ffdeaa0000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffec9c7b27bf045a724415a4f9000000000000000000000000fffd8963efd1fc6a506488495d951d5263988d2500000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000123685885532dcb685c442dc8312600000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000264a7996fa5000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb0000000000000000000000002f62f2b4c5fcd7570a709dec05d68ea19c82a9ec000000000000000000000000000000000000000000000000c3ad434b850200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000242e1a7d4d00000000000000000000000000000000000000000000000001009ee6e3054c0b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
// 		V:     new(big.Int).SetBytes(byteslice("0x01")),
// 		R:     new(big.Int).SetBytes(byteslice("0x3f3fce3c5da8e8c896209057348873d2295a05f725e662f3f0c0230ee637f6a9")),
// 		S:     new(big.Int).SetBytes(byteslice("0x46ad8ff3062b5d33d9bfdaa159d05409c66b1b1deb61247b7976fe8d9bc9f6c5")),
// 	}
//
// 	for i := 0; i < b.N; i++ {
// 		_, _ = tx.MarshalRLP()
// 	}
// }
