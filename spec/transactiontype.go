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
)

// TransactionType defines the spec version of a transaction.
type TransactionType int

const (
	// TransactionType0 is a legacy transaction.
	TransactionType0 TransactionType = iota
	// TransactionType1 is an access list transaction.
	TransactionType1
	// TransactionType2 is an EIP-1559 transaction.
	TransactionType2
	// TransactionType3 is a data blob transaction.
	TransactionType3
)

var transactionTypeStrings = [...]string{
	"0x0",
	"0x1",
	"0x2",
	"0x3",
}

// MarshalJSON implements json.Marshaler.
func (d *TransactionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", transactionTypeStrings[*d])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *TransactionType) UnmarshalJSON(input []byte) error {
	var err error
	switch strings.ToLower(strings.Trim(string(input), `"`)) {
	case "", "0", "0x", "0x0":
		*d = TransactionType0
	case "1", "0x1":
		*d = TransactionType1
	case "2", "0x2":
		*d = TransactionType2
	case "3", "0x3":
		*d = TransactionType3
	default:
		err = fmt.Errorf("unrecognised transaction type %s", string(input))
	}

	return err
}

// String returns a string representation of the item.
func (d TransactionType) String() string {
	if int(d) < 0 || int(d) >= len(transactionTypeStrings) {
		return "unknown"
	}

	return transactionTypeStrings[d]
}
