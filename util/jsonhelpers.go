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

package util

import (
	"fmt"
	"math/big"
	"strings"
)

var zero = big.NewInt(0)

// PreUnmarshalHexString tidies up an input hex string, removing any leading
// '0x' and ensuring that the input has an even number of hex digits.
func PreUnmarshalHexString(input string) string {
	res := strings.TrimPrefix(input, "0x")

	if len(res)%2 == 1 {
		res = fmt.Sprintf("0%s", res)
	}

	return res
}

// MarshalUint64 marshals a uint64 as per the Ethereum standard.
func MarshalUint64(input uint64) string {
	if input == 0 {
		return "0x0"
	}
	return fmt.Sprintf("0x%s", strings.TrimPrefix(fmt.Sprintf("%x", input), "0"))
}

// MarshalUint32 marshals a uint32 as per the Ethereum standard.
func MarshalUint32(input uint32) string {
	if input == 0 {
		return "0x0"
	}
	return fmt.Sprintf("0x%s", strings.TrimPrefix(fmt.Sprintf("%x", input), "0"))
}

// MarshalBigInt marshals a big.Int as per the Ethereum standard.
func MarshalBigInt(input *big.Int) string {
	if input == nil || input.Cmp(zero) == 0 {
		return "0x0"
	}
	return fmt.Sprintf("0x%s", strings.TrimPrefix(fmt.Sprintf("%x", input), "0"))
}

// MarshalByteArray marshals a byte array as per the Ethereum standard.
func MarshalByteArray(input []byte) string {
	if len(input) == 0 {
		return "0x"
	}
	return fmt.Sprintf("%#x", input)
}

// MarshalNullableByteArray marshals a byte array as per the Ethereum standard.
func MarshalNullableByteArray(input []byte) string {
	if len(input) == 0 {
		return ""
	}
	return fmt.Sprintf("%#x", input)
}
