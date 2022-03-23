// Copyright © 2022 Attestant Limited.
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
	"github.com/attestantio/go-execution-client/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionReceipt(t *testing.T) {
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
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type spec.transactionReceiptJSON",
		},
		{
			name:  "Berlin",
			input: []byte(`{"blockHash":"0x249ea54eada07708b29d7c424b8466dec9f1d98067b0be1b89c7ee660cca858d","blockNumber":"0xb542","contractAddress":"0x9a049f5d18c239efaa258af9f3e7002949a977a0","cumulativeGasUsed":"0x5dc0","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gasUsed":"0x5dc0","logs":[],"logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","status":"0x0","to":null,"transactionHash":"0x6c929e1c3d860ee225d7f3a7addf9e3f740603d243260536dfa2f3cf02b51de4","transactionIndex":"0x0","type":"0x0"}`),
		},
		{
			name:  "London",
			input: []byte(`{"blockHash":"0x8b3b6628fdc8861f23da314edd29d4ea88eea42b82fff6708a4be4b6ef87d296","blockNumber":"0xdc5deb","contractAddress":null,"cumulativeGasUsed":"0x10b8b8","effectiveGasPrice":"0x40eb6e398","from":"0x3e74d7a29db9c136fff150ac61a3ff6e56818774","gasUsed":"0x168c0","logs":[{"address":"0x40875223d61a688954263892d0f76c94fd6b3d4a","topics":["0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62","0x0000000000000000000000000000000005756b5a03e751bd0280e3a55bc05b6e","0x0000000000000000000000003e74d7a29db9c136fff150ac61a3ff6e56818774","0x0000000000000000000000000000000000000000000000000000000000000000"],"data":"0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001","blockNumber":"0xdc5deb","transactionHash":"0x27d0399f69000e21f2709cbba95582c0dbcab40cb71e4b3cab7ee8f6b1febfef","transactionIndex":"0x9","blockHash":"0x8b3b6628fdc8861f23da314edd29d4ea88eea42b82fff6708a4be4b6ef87d296","logIndex":"0x13","removed":false},{"address":"0x0000000005756b5a03e751bd0280e3a55bc05b6e","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000003e74d7a29db9c136fff150ac61a3ff6e56818774","0x0000000000000000000000000000000000000000000000000000000000003041"],"data":"0x","blockNumber":"0xdc5deb","transactionHash":"0x27d0399f69000e21f2709cbba95582c0dbcab40cb71e4b3cab7ee8f6b1febfef","transactionIndex":"0x9","blockHash":"0x8b3b6628fdc8861f23da314edd29d4ea88eea42b82fff6708a4be4b6ef87d296","logIndex":"0x14","removed":false}],"logsBloom":"0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400001000000000000008000000020000000000000000000000000080008000000000000000000000000000000000000000000000208000000000000000008000000000000000000000000900000000000000000000000000000000000000000000000000000000080000000000002000000000000000000000000000000000000000000000080000000000000000000000000020000000000000008000000000000000000080000028000000000200000000004000000000000000000000020000000000000000000000a0000000000","status":"0x1","to":"0x0000000005756b5a03e751bd0280e3a55bc05b6e","transactionHash":"0x27d0399f69000e21f2709cbba95582c0dbcab40cb71e4b3cab7ee8f6b1febfef","transactionIndex":"0x9","type":"0x2"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.TransactionReceipt
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

func TestTransactionReceiptBerlinFuncs(t *testing.T) {
	input := []byte(`{"blockHash":"0x249ea54eada07708b29d7c424b8466dec9f1d98067b0be1b89c7ee660cca858d","blockNumber":"0xb542","contractAddress":"0x9a049f5d18c239efaa258af9f3e7002949a977a0","cumulativeGasUsed":"0x5dc0","from":"0xa1e4380a3b1f749673e270229993ee55f35663b4","gasUsed":"0x5dc0","logs":[],"logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","status":"0x0","to":null,"transactionHash":"0x6c929e1c3d860ee225d7f3a7addf9e3f740603d243260536dfa2f3cf02b51de4","transactionIndex":"0x0","type":"0x0"}`)
	var receipt spec.TransactionReceipt
	require.NoError(t, json.Unmarshal(input, &receipt))

	assert.Equal(t, types.Hash{0x24, 0x9e, 0xa5, 0x4e, 0xad, 0xa0, 0x77, 0x8, 0xb2, 0x9d, 0x7c, 0x42, 0x4b, 0x84, 0x66, 0xde, 0xc9, 0xf1, 0xd9, 0x80, 0x67, 0xb0, 0xbe, 0x1b, 0x89, 0xc7, 0xee, 0x66, 0xc, 0xca, 0x85, 0x8d}, receipt.BlockHash())
	assert.Equal(t, uint32(0xb542), receipt.BlockNumber())
	assert.Equal(t, types.Address{0x9a, 0x4, 0x9f, 0x5d, 0x18, 0xc2, 0x39, 0xef, 0xaa, 0x25, 0x8a, 0xf9, 0xf3, 0xe7, 0x0, 0x29, 0x49, 0xa9, 0x77, 0xa0}, *receipt.ContractAddress())
	assert.Equal(t, uint32(0x5dc0), receipt.CumulativeGasUsed())
	assert.Equal(t, uint64(0), receipt.EffectiveGasPrice())
	assert.Equal(t, types.Address{0xa1, 0xe4, 0x38, 0xa, 0x3b, 0x1f, 0x74, 0x96, 0x73, 0xe2, 0x70, 0x22, 0x99, 0x93, 0xee, 0x55, 0xf3, 0x56, 0x63, 0xb4}, receipt.From())
	assert.Equal(t, uint32(0x5dc0), receipt.GasUsed())
	assert.Len(t, receipt.Logs(), 0)
	assert.Equal(t, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, receipt.LogsBloom())
	assert.Equal(t, uint32(0), receipt.Status())
	assert.Nil(t, receipt.To())
	assert.Equal(t, types.Hash{0x6c, 0x92, 0x9e, 0x1c, 0x3d, 0x86, 0xe, 0xe2, 0x25, 0xd7, 0xf3, 0xa7, 0xad, 0xdf, 0x9e, 0x3f, 0x74, 0x6, 0x3, 0xd2, 0x43, 0x26, 0x5, 0x36, 0xdf, 0xa2, 0xf3, 0xcf, 0x2, 0xb5, 0x1d, 0xe4}, receipt.TransactionHash())
	assert.Equal(t, uint32(0), receipt.TransactionIndex())
	assert.Equal(t, spec.TransactionType0, receipt.Type())
}

func TestTransactionReceiptLondonFuncs(t *testing.T) {
	input := []byte(`{"blockHash":"0x8b3b6628fdc8861f23da314edd29d4ea88eea42b82fff6708a4be4b6ef87d296","blockNumber":"0xdc5deb","contractAddress":null,"cumulativeGasUsed":"0x10b8b8","effectiveGasPrice":"0x40eb6e398","from":"0x3e74d7a29db9c136fff150ac61a3ff6e56818774","gasUsed":"0x168c0","logs":[{"address":"0x40875223d61a688954263892d0f76c94fd6b3d4a","topics":["0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62","0x0000000000000000000000000000000005756b5a03e751bd0280e3a55bc05b6e","0x0000000000000000000000003e74d7a29db9c136fff150ac61a3ff6e56818774","0x0000000000000000000000000000000000000000000000000000000000000000"],"data":"0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001","blockNumber":"0xdc5deb","transactionHash":"0x27d0399f69000e21f2709cbba95582c0dbcab40cb71e4b3cab7ee8f6b1febfef","transactionIndex":"0x9","blockHash":"0x8b3b6628fdc8861f23da314edd29d4ea88eea42b82fff6708a4be4b6ef87d296","logIndex":"0x13","removed":false},{"address":"0x0000000005756b5a03e751bd0280e3a55bc05b6e","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000003e74d7a29db9c136fff150ac61a3ff6e56818774","0x0000000000000000000000000000000000000000000000000000000000003041"],"data":"0x","blockNumber":"0xdc5deb","transactionHash":"0x27d0399f69000e21f2709cbba95582c0dbcab40cb71e4b3cab7ee8f6b1febfef","transactionIndex":"0x9","blockHash":"0x8b3b6628fdc8861f23da314edd29d4ea88eea42b82fff6708a4be4b6ef87d296","logIndex":"0x14","removed":false}],"logsBloom":"0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400001000000000000008000000020000000000000000000000000080008000000000000000000000000000000000000000000000208000000000000000008000000000000000000000000900000000000000000000000000000000000000000000000000000000080000000000002000000000000000000000000000000000000000000000080000000000000000000000000020000000000000008000000000000000000080000028000000000200000000004000000000000000000000020000000000000000000000a0000000000","status":"0x1","to":"0x0000000005756b5a03e751bd0280e3a55bc05b6e","transactionHash":"0x27d0399f69000e21f2709cbba95582c0dbcab40cb71e4b3cab7ee8f6b1febfef","transactionIndex":"0x9","type":"0x2"}`)
	var receipt spec.TransactionReceipt
	require.NoError(t, json.Unmarshal(input, &receipt))

	assert.Equal(t, types.Hash{0x8b, 0x3b, 0x66, 0x28, 0xfd, 0xc8, 0x86, 0x1f, 0x23, 0xda, 0x31, 0x4e, 0xdd, 0x29, 0xd4, 0xea, 0x88, 0xee, 0xa4, 0x2b, 0x82, 0xff, 0xf6, 0x70, 0x8a, 0x4b, 0xe4, 0xb6, 0xef, 0x87, 0xd2, 0x96}, receipt.BlockHash())
	assert.Equal(t, uint32(0xdc5deb), receipt.BlockNumber())
	assert.Nil(t, receipt.ContractAddress())
	assert.Equal(t, uint32(0x10b8b8), receipt.CumulativeGasUsed())
	assert.Equal(t, uint64(0x40eb6e398), receipt.EffectiveGasPrice())
	assert.Equal(t, types.Address{0x3e, 0x74, 0xd7, 0xa2, 0x9d, 0xb9, 0xc1, 0x36, 0xff, 0xf1, 0x50, 0xac, 0x61, 0xa3, 0xff, 0x6e, 0x56, 0x81, 0x87, 0x74}, receipt.From())
	assert.Equal(t, uint32(0x168c0), receipt.GasUsed())
	assert.Len(t, receipt.Logs(), 2)
	assert.Equal(t, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x90, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x2, 0x80, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0}, receipt.LogsBloom())
	assert.Equal(t, uint32(0x1), receipt.Status())
	assert.Equal(t, types.Address{0x0, 0x0, 0x0, 0x0, 0x5, 0x75, 0x6b, 0x5a, 0x3, 0xe7, 0x51, 0xbd, 0x2, 0x80, 0xe3, 0xa5, 0x5b, 0xc0, 0x5b, 0x6e}, *receipt.To())
	assert.Equal(t, types.Hash{0x27, 0xd0, 0x39, 0x9f, 0x69, 0x0, 0xe, 0x21, 0xf2, 0x70, 0x9c, 0xbb, 0xa9, 0x55, 0x82, 0xc0, 0xdb, 0xca, 0xb4, 0xc, 0xb7, 0x1e, 0x4b, 0x3c, 0xab, 0x7e, 0xe8, 0xf6, 0xb1, 0xfe, 0xbf, 0xef}, receipt.TransactionHash())
	assert.Equal(t, uint32(0x9), receipt.TransactionIndex())
	assert.Equal(t, spec.TransactionType2, receipt.Type())
}

func TestTransactionReceiptUnknownFuncs(t *testing.T) {
	receipt := spec.TransactionReceipt{}

	assert.Panics(t, func() { receipt.BlockHash() })
	assert.Panics(t, func() { receipt.BlockNumber() })
	assert.Panics(t, func() { receipt.ContractAddress() })
	assert.Panics(t, func() { receipt.CumulativeGasUsed() })
	assert.Panics(t, func() { receipt.EffectiveGasPrice() })
	assert.Panics(t, func() { receipt.From() })
	assert.Panics(t, func() { receipt.GasUsed() })
	assert.Panics(t, func() { receipt.Logs() })
	assert.Panics(t, func() { receipt.LogsBloom() })
	assert.Panics(t, func() { receipt.Status() })
	assert.Panics(t, func() { receipt.To() })
	assert.Panics(t, func() { receipt.TransactionHash() })
	assert.Panics(t, func() { receipt.TransactionIndex() })
	assert.Panics(t, func() { receipt.Type() })
}
