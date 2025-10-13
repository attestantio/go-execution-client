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
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/attestantio/go-execution-client/types"
	"github.com/pkg/errors"
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

// MarshalInt64 marshals an int64 as per the Ethereum standard.
func MarshalInt64(input int64) string {
	if input == 0 {
		return "0x0"
	}

	return fmt.Sprintf("0x%s", strings.TrimPrefix(fmt.Sprintf("%x", input), "0"))
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

// MarshalNullableUint32 marshals a uint32 as per the Ethereum standard, with 0 as null.
func MarshalNullableUint32(input uint32) string {
	if input == 0 {
		return ""
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

var zeroAddress = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

// MarshalAddress marshals an address as per the Ethereum standard.
func MarshalAddress(input []byte) string {
	return fmt.Sprintf("%#x", input)
}

// MarshalNullableAddress marshals an address as per the Ethereum standard.
func MarshalNullableAddress(input []byte) string {
	if bytes.Equal(input, zeroAddress) {
		return ""
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

// StrToAddress turns a string in to an address.
func StrToAddress(name string, input string) (types.Address, error) {
	var res types.Address
	if input == "" {
		return res, fmt.Errorf("%s missing", name)
	}

	val, err := hex.DecodeString(PreUnmarshalHexString(input))
	if err != nil {
		return res, errors.Wrap(err, fmt.Sprintf("%s invalid", name))
	}

	if len(val) != len(res) {
		return res, fmt.Errorf("%s incorrect length", name)
	}

	copy(res[:], val)

	return res, nil
}

// StrToBigInt turns a string in to a big.Int.
func StrToBigInt(name string, input string) (*big.Int, error) {
	if input == "" {
		return nil, fmt.Errorf("%s missing", name)
	}

	res, success := new(big.Int).SetString(PreUnmarshalHexString(input), 16)
	if !success {
		return nil, fmt.Errorf("%s invalid", name)
	}

	return res, nil
}

// StrToByteArray turns a string in to a byte array.
func StrToByteArray(name string, input string) ([]byte, error) {
	if input == "" {
		return nil, fmt.Errorf("%s missing", name)
	}

	res, err := hex.DecodeString(PreUnmarshalHexString(input))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s invalid", name))
	}

	return res, nil
}

// StrToHash turns a string in to a hash.
func StrToHash(name string, input string) (types.Hash, error) {
	var res types.Hash
	if input == "" {
		return res, fmt.Errorf("%s missing", name)
	}

	val, err := hex.DecodeString(PreUnmarshalHexString(input))
	if err != nil {
		return res, errors.Wrap(err, fmt.Sprintf("%s invalid", name))
	}

	if len(val) != len(res) {
		return res, fmt.Errorf("%s incorrect length", name)
	}

	copy(res[:], val)

	return res, nil
}

// StrToRoot turns a string in to a root.
func StrToRoot(name string, input string) (types.Root, error) {
	var res types.Root
	if input == "" {
		return res, fmt.Errorf("%s missing", name)
	}

	val, err := hex.DecodeString(PreUnmarshalHexString(input))
	if err != nil {
		return res, errors.Wrap(err, fmt.Sprintf("%s invalid", name))
	}

	if len(val) != len(res) {
		return res, fmt.Errorf("%s incorrect length", name)
	}

	copy(res[:], val)

	return res, nil
}

// StrToTime turns a string in to a time.Time.
func StrToTime(name string, input string) (time.Time, error) {
	var (
		res time.Time
		err error
	)

	if input == "" {
		return res, fmt.Errorf("%s missing", name)
	}

	val, err := strconv.ParseUint(PreUnmarshalHexString(input), 16, 64)
	if err != nil {
		return res, errors.Wrap(err, fmt.Sprintf("%s invalid", name))
	}

	res = time.Unix(int64(val), 0)

	return res, nil
}

// StrToUint64 turns a string in to a uint64.
func StrToUint64(name string, input string) (uint64, error) {
	var (
		res uint64
		err error
	)

	if input == "" {
		return res, fmt.Errorf("%s missing", name)
	}

	res, err = strconv.ParseUint(PreUnmarshalHexString(input), 16, 64)
	if err != nil {
		return res, errors.Wrap(err, fmt.Sprintf("%s invalid", name))
	}

	return res, nil
}

// StrToUint32 turns a string in to a uint32.
func StrToUint32(name string, input string) (uint32, error) {
	var (
		res uint64
		err error
	)

	if input == "" {
		return 0, fmt.Errorf("%s missing", name)
	}

	res, err = strconv.ParseUint(PreUnmarshalHexString(input), 16, 32)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("%s invalid", name))
	}

	return uint32(res), nil
}
