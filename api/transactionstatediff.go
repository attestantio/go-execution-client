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
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// TransactionStateDiff contains the changes made to a balance as a result of a transaction.
type TransactionStateDiff struct {
	Output  []byte
	Balance *TransactionStateChange
	Nonce   *TransactionStateChange
	Storage map[spec.Hash]*TransactionStorageChange
}

// transactionStateDiffJSON is the spec representation of the struct.
type transactionStateDiffJSON struct {
	Balance *json.RawMessage            `json:"balance,omitempty"`
	Nonce   *json.RawMessage            `json:"nonce,omitempty"`
	Storage map[string]*json.RawMessage `json:"storage,omitempty"`
}

// transactionStateDiffYAML is the spec representation of the struct.
type transactionStateDiffYAML struct {
	Balance *TransactionStateChange              `yaml:"balance,omitempty"`
	Nonce   *TransactionStateChange              `yaml:"nonce,omitempty"`
	Storage map[string]*TransactionStorageChange `yaml:"storage,omitempty"`
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
	var transactionStateDiffJSON transactionStateDiffJSON
	if err := json.Unmarshal(input, &transactionStateDiffJSON); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&transactionStateDiffJSON)
}

func (t *TransactionStateDiff) unpack(transactionStateDiffJSON *transactionStateDiffJSON) error {
	if transactionStateDiffJSON.Balance != nil && (*transactionStateDiffJSON.Balance)[0] == '{' {
		var stateChange TransactionStateChange
		if err := json.Unmarshal(*transactionStateDiffJSON.Balance, &stateChange); err != nil {
			return errors.Wrap(err, "invalid balance JSON")
		}
		t.Balance = &stateChange
	}

	if transactionStateDiffJSON.Nonce != nil && (*transactionStateDiffJSON.Nonce)[0] == '{' {
		var stateChange TransactionStateChange
		if err := json.Unmarshal(*transactionStateDiffJSON.Nonce, &stateChange); err != nil {
			return errors.Wrap(err, "invalid nonce JSON")
		}
		t.Nonce = &stateChange
	}

	storage := make(map[spec.Hash]*TransactionStorageChange)
	for k, v := range transactionStateDiffJSON.Storage {
		var stateChange TransactionStorageChange
		if err := json.Unmarshal([]byte(*v), &stateChange); err != nil {
			return errors.Wrap(err, "invalid storage JSON")
		}
		hash, err := hex.DecodeString(strings.TrimPrefix(k, "0x"))
		if err != nil {
			return err
		}
		var key spec.Hash
		copy(key[:], hash)
		storage[key] = &stateChange
	}
	t.Storage = storage

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (t *TransactionStateDiff) MarshalYAML() ([]byte, error) {
	storage := make(map[string]*TransactionStorageChange)
	for k, v := range t.Storage {
		storage[fmt.Sprintf("%#x", k)] = v
	}

	yamlBytes, err := yaml.MarshalWithOptions(&transactionStateDiffYAML{
		Balance: t.Balance,
		Nonce:   t.Nonce,
		Storage: storage,
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *TransactionStateDiff) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var transactionStateDiffJSON transactionStateDiffJSON
	if err := yaml.Unmarshal(input, &transactionStateDiffJSON); err != nil {
		return err
	}
	return t.unpack(&transactionStateDiffJSON)
}

// String returns a string version of the structure.
func (t *TransactionStateDiff) String() string {
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
