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
	"bytes"
	"encoding/binary"

	"github.com/attestantio/go-execution-client/types"
)

// Encodings are based upon the rules at https://eth.wiki/en/fundamentals/rlp

// RLPAddress returns the RLP encoding of the given address.
func RLPAddress(input types.Address) []byte {
	buf := bytes.Buffer{}
	buf.WriteByte(0x94)
	buf.Write(input[:])
	return buf.Bytes()
}

// RLPHash returns the RLP encoding of the given hash.
func RLPHash(input types.Hash) []byte {
	buf := bytes.Buffer{}
	buf.WriteByte(0xa0)
	buf.Write(input[:])
	return buf.Bytes()
}

// RLPUint64 returns the RLP encoding of the given uint64.
func RLPUint64(input uint64) []byte {
	if input == 0 {
		return []byte{0x80}
	}
	if input <= 0x7f {
		return []byte{uint8(input)}
	}
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, input)
	for loc := 0; loc < 8; loc++ {
		if res[loc] != 0 {
			return RLPBytes(res[loc:])
		}
	}
	return RLPBytes(res)
	// Find the byte position of the MSB.
	// msb := int(math.Log2(float64(input))) / 8
	// return util.RLPBytes(res[
}

// RLPList returns the RLP encoding of the list of bytes.
// It is expected that each entry in the list is already encoded.
func RLPList(items [][]byte) []byte {
	buf := bytes.Buffer{}
	for _, item := range items {
		buf.Write(item)
	}
	return append(encodeLength(uint64(buf.Len()), 0xc0), buf.Bytes()...)
}

// RLPBytes returns the RLP encoding of the given bytes.
func RLPBytes(input []byte) []byte {
	if len(input) == 1 && input[0] < 0x80 {
		return input
	}
	return append(encodeLength(uint64(len(input)), 0x80), input...)
}

func encodeLength(length uint64, offset byte) []byte {
	if length < 56 {
		return []byte{byte(length + uint64(offset))}
	}
	bl := toBinary(length)
	return append([]byte{byte(len(bl) + int(offset) + 55)}, bl...)
}

func toBinary(x uint64) []byte {
	if x == 0 {
		return nil
	}
	return append(toBinary(x/256), byte(x%256))
}
