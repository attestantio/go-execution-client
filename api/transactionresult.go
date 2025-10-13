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

// TransactionResult contains the result of a transaction.
type TransactionResult struct {
	Output          []byte
	StateDiff       map[types.Address]*TransactionStateDiff
	TransactionHash types.Hash
}

// transactionResultJSON is the spec representation of the struct.
type transactionResultJSON struct {
	Output          string                           `json:"output,omitempty"`
	StateDiff       map[string]*TransactionStateDiff `json:"stateDiff,omitempty"`
	TransactionHash string                           `json:"transactionHash"`
}

// MarshalJSON implements json.Marshaler.
func (t *TransactionResult) MarshalJSON() ([]byte, error) {
	stateDiff := make(map[string]*TransactionStateDiff)
	for k, v := range t.StateDiff {
		stateDiff[util.MarshalAddress(k[:])] = v
	}

	return json.Marshal(&transactionResultJSON{
		Output:          fmt.Sprintf("%#x", t.Output),
		StateDiff:       stateDiff,
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

// String returns a string version of the structure.
func (t *TransactionResult) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}

func (t *TransactionResult) unpack(data *transactionResultJSON) error {
	var err error

	if data.Output != "" {
		t.Output, err = util.StrToByteArray("output", data.Output)
		if err != nil {
			return err
		}
	}

	stateDiff := make(map[types.Address]*TransactionStateDiff)

	for k, v := range data.StateDiff {
		address, err := util.StrToAddress("address", k)
		if err != nil {
			return err
		}

		stateDiff[address] = v
	}

	t.StateDiff = stateDiff

	t.TransactionHash, err = util.StrToHash("transaction hash", data.TransactionHash)
	if err != nil {
		return err
	}

	return nil
}
