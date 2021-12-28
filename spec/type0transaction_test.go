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
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
	"github.com/stretchr/testify/require"
)

func address(input string) *types.Address {
	tmp, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}
	res := types.Address{}
	copy(res[:], tmp)
	return &res
}

func byteslice(input string) []byte {
	res, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}
	return res
}

// TestType0TransactionRLP tests the RLP encoding of transactions.
func TestType0TransactionRLP(t *testing.T) {
	tests := []struct {
		name     string
		input    *spec.Transaction
		expected []byte
		err      string
	}{
		{
			name: "Transfer",
			input: &spec.Transaction{
				Type:     0,
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
			input: &spec.Transaction{
				Type:     0,
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
			input: &spec.Transaction{
				Type:     0,
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
