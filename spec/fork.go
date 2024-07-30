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

package spec

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// Fork defines the fork version of a response.
type Fork int

const (
	// ForkUnknown is an unknown fork.
	ForkUnknown Fork = iota
	// ForkBerlin is the Berlin fork.
	ForkBerlin
	// ForkLondon is the London fork.
	ForkLondon
	// ForkShanghai is the Shanghai fork.
	ForkShanghai
	// ForkCancun is the Cancun fork.
	ForkCancun
)

var forkStrings = [...]string{
	"unknown",
	"berlin",
	"london",
	"shanghai",
	"cancun",
}

// MarshalJSON implements json.Marshaler.
func (d *Fork) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", forkStrings[*d])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *Fork) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("fork missing")
	}

	var err error
	switch strings.ToLower(strings.Trim(string(input), `"`)) {
	case "berlin":
		*d = ForkBerlin
	case "london":
		*d = ForkLondon
	case "shanghai":
		*d = ForkShanghai
	case "cancun":
		*d = ForkCancun
	default:
		err = fmt.Errorf("unrecognised fork version %s", string(input))
	}
	return err
}

// String returns a string representation of the item.
func (d Fork) String() string {
	if int(d) < 0 || int(d) >= len(forkStrings) {
		return "unknown"
	}

	return forkStrings[d]
}
