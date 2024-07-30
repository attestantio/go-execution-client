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

	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// TransactionStateDiff contains the changes made to a balance as a result of a transaction.
type TransactionStateDiff struct {
	Output  []byte
	Balance *TransactionStateChange
	Nonce   *TransactionStateChange
	Storage map[types.Hash]*TransactionStorageChange
}

// transactionStateDiffJSON is the spec representation of the struct.
type transactionStateDiffJSON struct {
	Balance *json.RawMessage            `json:"balance,omitempty"`
	Nonce   *json.RawMessage            `json:"nonce,omitempty"`
	Storage map[string]*json.RawMessage `json:"storage,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (t *TransactionStateDiff) MarshalJSON() ([]byte, error) {
	var balance *json.RawMessage
	if t.Balance != nil {
		tmp, err := json.Marshal(t.Balance)
		if err != nil {
			return nil, err
		}
		b := json.RawMessage(tmp)
		balance = &b
	}

	var nonce *json.RawMessage
	if t.Nonce != nil {
		tmp, err := json.Marshal(t.Nonce)
		if err != nil {
			return nil, err
		}
		n := json.RawMessage(tmp)
		nonce = &n
	}

	storage := make(map[string]*json.RawMessage)
	for k, v := range t.Storage {
		tmp, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		item := json.RawMessage(tmp)
		storage[fmt.Sprintf("%#x", k)] = &item
	}

	return json.Marshal(&transactionStateDiffJSON{
		Balance: balance,
		Nonce:   nonce,
		Storage: storage,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionStateDiff) UnmarshalJSON(input []byte) error {
	var data transactionStateDiffJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

func (t *TransactionStateDiff) unpack(data *transactionStateDiffJSON) error {
	if data.Balance != nil && (*data.Balance)[0] == '{' {
		var stateChange TransactionStateChange
		if err := json.Unmarshal(*data.Balance, &stateChange); err != nil {
			return errors.Wrap(err, "invalid balance JSON")
		}
		t.Balance = &stateChange
	}

	if data.Nonce != nil && (*data.Nonce)[0] == '{' {
		var stateChange TransactionStateChange
		if err := json.Unmarshal(*data.Nonce, &stateChange); err != nil {
			return errors.Wrap(err, "invalid nonce JSON")
		}
		t.Nonce = &stateChange
	}

	storage := make(map[types.Hash]*TransactionStorageChange)
	for k, v := range data.Storage {
		var stateChange TransactionStorageChange
		if err := json.Unmarshal([]byte(*v), &stateChange); err != nil {
			return errors.Wrap(err, "invalid storage JSON")
		}
		key, err := util.StrToHash("storage key", k)
		if err != nil {
			return err
		}
		storage[key] = &stateChange
	}
	t.Storage = storage

	return nil
}

// String returns a string version of the structure.
func (t *TransactionStateDiff) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
