// Copyright © 2021 - 2024 Attestant Limited.
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

package types

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

// AddressLength is the length of an execution layer address.
const AddressLength = 20

// Address is a 20-byte execution layer address.
type Address [AddressLength]byte

var emptyAddress = Address{}

// IsZero returns true if the address is zero.
func (a Address) IsZero() bool {
	return bytes.Equal(a[:], emptyAddress[:])
}

// String returns the EIP-55 string representation of the address.
func (a Address) String() string {
	data := []byte(hex.EncodeToString(a[:]))

	keccak := sha3.NewLegacyKeccak256()
	keccak.Write(data)
	hash := keccak.Sum(nil)

	for i := 0; i < len(data); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte >>= 4
		} else {
			hashByte &= 0xf
		}

		if data[i] > '9' && hashByte > 7 {
			data[i] -= 32
		}
	}

	return fmt.Sprintf("0x%s", string(data))
}

// Format formats the address.
func (a Address) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, a.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}

		fmt.Fprintf(state, "%"+format, a[:])
	default:
		fmt.Fprintf(state, "%"+format, a[:])
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *Address) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("input missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid prefix")
	}

	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid suffix")
	}

	if len(input) != 1+2+AddressLength*2+1 {
		return errors.New("incorrect length")
	}

	length, err := hex.Decode(a[:], input[3:3+AddressLength*2])
	if err != nil {
		return errors.Wrapf(err, "invalid value %s", string(input[3:3+AddressLength*2]))
	}

	if length != AddressLength {
		return errors.New("incorrect length")
	}

	return nil
}

// MarshalJSON implements json.Marshaler.
func (a Address) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, a.String())), nil
}
