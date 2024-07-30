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

	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// TransactionStorageChange is a change to state, from one value to another.
type TransactionStorageChange struct {
	From []byte
	To   []byte
}

// transactionStorageChangeJSON is the spec representation of the struct.
type transactionStorageChangeJSON struct {
	Creation   string                                  `json:"+,omitempty"`
	Alteration *transactionStorageChangeAlterationJSON `json:"*,omitempty"`
	Deletion   string                                  `json:"-,omitempty"`
}

type transactionStorageChangeAlterationJSON struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (t *TransactionStorageChange) MarshalJSON() ([]byte, error) {
	if t.From == nil {
		return json.Marshal(&transactionStorageChangeJSON{
			Creation: util.MarshalByteArray(t.To),
		})
	}

	if t.To == nil {
		return json.Marshal(&transactionStorageChangeJSON{
			Deletion: util.MarshalByteArray(t.From),
		})
	}

	return json.Marshal(&transactionStorageChangeJSON{
		Alteration: &transactionStorageChangeAlterationJSON{
			From: util.MarshalByteArray(t.From),
			To:   util.MarshalByteArray(t.To),
		},
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionStorageChange) UnmarshalJSON(input []byte) error {
	var data transactionStorageChangeJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

func (t *TransactionStorageChange) unpack(data *transactionStorageChangeJSON) error {
	var err error
	if data.Creation != "" {
		t.To, err = util.StrToByteArray("creation", data.Creation)
		if err != nil {
			return err
		}
	}

	if data.Deletion != "" {
		t.From, err = util.StrToByteArray("deletion", data.Deletion)
		if err != nil {
			return err
		}
	}

	if data.Alteration != nil {
		t.From, err = util.StrToByteArray("from", data.Alteration.From)
		if err != nil {
			return err
		}

		t.To, err = util.StrToByteArray("to", data.Alteration.To)
		if err != nil {
			return err
		}
	}

	return nil
}

// String returns a string version of the structure.
func (t *TransactionStorageChange) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
