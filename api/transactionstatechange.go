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

package api

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// TransactionStateChange is a change to state, from one value to another.
type TransactionStateChange struct {
	From *big.Int
	To   *big.Int
}

// transactionStateChangeJSON is the spec representation of the struct.
type transactionStateChangeJSON struct {
	Creation   string                                `json:"+,omitempty"`
	Alteration *transactionStateChangeAlterationJSON `json:"*,omitempty"`
	Deletion   string                                `json:"-,omitempty"` //nolint:revive // omitempty required for correct behavior
}

type transactionStateChangeAlterationJSON struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// MarshalJSON implements json.Marshaler.
func (t *TransactionStateChange) MarshalJSON() ([]byte, error) {
	if t.From == nil {
		return json.Marshal(&transactionStateChangeJSON{
			Creation: util.MarshalBigInt(t.To),
		})
	}

	if t.To == nil {
		return json.Marshal(&transactionStateChangeJSON{
			Deletion: util.MarshalBigInt(t.From),
		})
	}

	return json.Marshal(&transactionStateChangeJSON{
		Alteration: &transactionStateChangeAlterationJSON{
			From: util.MarshalBigInt(t.From),
			To:   util.MarshalBigInt(t.To),
		},
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionStateChange) UnmarshalJSON(input []byte) error {
	var data transactionStateChangeJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

// String returns a string version of the structure.
func (t *TransactionStateChange) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}

func (t *TransactionStateChange) unpack(data *transactionStateChangeJSON) error {
	var err error
	if data.Creation != "" {
		t.To, err = util.StrToBigInt("creation", data.Creation)
		if err != nil {
			return err
		}
	}

	if data.Deletion != "" {
		t.From, err = util.StrToBigInt("deletion", data.Deletion)
		if err != nil {
			return err
		}
	}

	if data.Alteration != nil && data.Alteration.From != "" {
		t.From, err = util.StrToBigInt("from", data.Alteration.From)
		if err != nil {
			return err
		}
	}

	if data.Alteration != nil && data.Alteration.To != "" {
		t.To, err = util.StrToBigInt("to", data.Alteration.To)
		if err != nil {
			return err
		}
	}

	return nil
}
