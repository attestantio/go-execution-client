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

	"github.com/attestantio/go-execution-client/util"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// TransactionResult contains the result of a transaction.
type TransactionResult struct {
	Output          []byte
	StateDiff       map[string]*TransactionStateDiff
	TransactionHash []byte
}

// transactionResultJSON is the spec representation of the struct.
type transactionResultJSON struct {
	Output          string                           `json:"output,omitempty"`
	StateDiff       map[string]*TransactionStateDiff `json:"stateDiff,omitempty"`
	TransactionHash string                           `json:"transactionHash"`
}

// transactionResultYAML is the spec representation of the struct.
type transactionResultYAML struct {
	Output          string                           `yaml:"output,omitempty"`
	StateDiff       map[string]*TransactionStateDiff `yaml:"stateDiff,omitempty"`
	TransactionHash string                           `yaml:"transactionHash"`
}

// MarshalJSON implements json.Marshaler.
func (t *TransactionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(&transactionResultJSON{
		Output:          fmt.Sprintf("%#x", t.Output),
		StateDiff:       t.StateDiff,
		TransactionHash: fmt.Sprintf("%#x", t.TransactionHash),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionResult) UnmarshalJSON(input []byte) error {
	var transactionResultJSON transactionResultJSON
	if err := json.Unmarshal(input, &transactionResultJSON); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&transactionResultJSON)
}

func (t *TransactionResult) unpack(transactionResultJSON *transactionResultJSON) error {
	var err error

	if transactionResultJSON.Output != "" {
		t.Output, err = hex.DecodeString(util.PreUnmarshalHexString(transactionResultJSON.Output))
		if err != nil {
			return errors.Wrap(err, "output invalid")
		}
	}

	t.StateDiff = transactionResultJSON.StateDiff

	if transactionResultJSON.TransactionHash == "" {
		return errors.New("transaction hash missing")
	}
	t.TransactionHash, err = hex.DecodeString(util.PreUnmarshalHexString(transactionResultJSON.TransactionHash))
	if err != nil {
		return errors.Wrap(err, "transaction hash invalid")
	}

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (t *TransactionResult) MarshalYAML() ([]byte, error) {
	yamlBytes, err := yaml.MarshalWithOptions(&transactionResultYAML{
		Output:          fmt.Sprintf("%#x", t.Output),
		StateDiff:       t.StateDiff,
		TransactionHash: fmt.Sprintf("%#x", t.TransactionHash),
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *TransactionResult) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var transactionResultJSON transactionResultJSON
	if err := yaml.Unmarshal(input, &transactionResultJSON); err != nil {
		return err
	}
	return t.unpack(&transactionResultJSON)
}

// String returns a string version of the structure.
func (t *TransactionResult) String() string {
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
