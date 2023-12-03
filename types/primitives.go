// Copyright © 2021 - 2023 Attestant Limited.
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
)

// Hash is a 32-byte hash.
type Hash [32]byte

// String returns the string representation of the hash.
func (h Hash) String() string {
	return fmt.Sprintf("%#x", h)
}

// Format formats the hash.
func (h Hash) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, h.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, h[:])
	default:
		fmt.Fprintf(state, "%"+format, h[:])
	}
}

// VersionedHash is a 32-byte hash with the first byte being a version.
type VersionedHash [32]byte

// VersionedHashLength is the length of a versioned hash.
const VersionedHashLength = 32

// String returns the string representation of the versioned hash.
func (h VersionedHash) String() string {
	return fmt.Sprintf("%#x", h)
}

// Format formats the versioned hash.
func (h VersionedHash) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, h.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, h[:])
	default:
		fmt.Fprintf(state, "%"+format, h[:])
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (h *VersionedHash) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("input missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid suffix")
	}
	if len(input) != 1+2+VersionedHashLength*2+1 {
		return errors.New("incorrect length")
	}

	length, err := hex.Decode(h[:], input[3:3+VersionedHashLength*2])
	if err != nil {
		return errors.Wrapf(err, "invalid value %s", string(input[3:3+VersionedHashLength*2]))
	}

	if length != VersionedHashLength {
		return errors.New("incorrect length")
	}

	return nil
}

// MarshalJSON implements json.Marshaler.
func (h VersionedHash) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%#x"`, h)), nil
}

// MarshalYAML implements yaml.Marshaler.
func (r Root) MarshalYAML() ([]byte, error) {
	return []byte(fmt.Sprintf(`'%#x'`, r)), nil
}

// Root is a 32-byte merkle root.
type Root [32]byte

// String returns the string representation of the root.
func (r Root) String() string {
	return fmt.Sprintf("%#x", r)
}

// Format formats the root.
func (r Root) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, r.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, r[:])
	default:
		fmt.Fprintf(state, "%"+format, r[:])
	}
}
