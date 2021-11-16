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

package spec

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Transaction defines a transaction.
type Transaction interface {
	// Type returns the transaction type.
	Type() uint64
}

type transactionTypeJSON struct {
	Type string `json:"type"`
}

// UnmarshalTransactionJSON unmarshals a transaction.
func UnmarshalTransactionJSON(input []byte) (Transaction, error) {
	var data transactionTypeJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, errors.Wrap(err, "invalid JSON")
	}

	switch data.Type {
	case "0x", "0x0", "0x00":
		var type0Transaction Type0Transaction
		if err := json.Unmarshal(input, &type0Transaction); err != nil {
			return nil, errors.Wrap(err, "invalid JSON")
		}
		return &type0Transaction, nil
	case "0x1", "0x01":
		var type1Transaction Type1Transaction
		if err := json.Unmarshal(input, &type1Transaction); err != nil {
			return nil, errors.Wrap(err, "invalid JSON")
		}
		return &type1Transaction, nil
	case "0x2", "0x02":
		var type2Transaction Type2Transaction
		if err := json.Unmarshal(input, &type2Transaction); err != nil {
			return nil, errors.Wrap(err, "invalid JSON")
		}
		return &type2Transaction, nil
	default:
		return nil, fmt.Errorf("unhandled transaction type %s", data.Type)
	}
}
