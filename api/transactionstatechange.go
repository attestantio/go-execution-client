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
	"github.com/goccy/go-yaml"
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
	Deletion   string                                `json:"-,omitempty"`
}

type transactionStateChangeAlterationJSON struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// transactionStateChangeYAML is the spec representation of the struct.
type transactionStateChangeYAML struct {
	Creation   string                                `yaml:"+,omitempty"`
	Alteration *transactionStateChangeAlterationYAML `yaml:"*,omitempty"`
	Deletion   string                                `yaml:"-,omitempty"`
}

type transactionStateChangeAlterationYAML struct {
	From *big.Int `json:"from"`
	To   *big.Int `json:"to"`
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
	var transactionStateBalanceDiffJSON transactionStateChangeJSON
	if err := json.Unmarshal(input, &transactionStateBalanceDiffJSON); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&transactionStateBalanceDiffJSON)
}

func (t *TransactionStateChange) unpack(transactionStateChangeJSON *transactionStateChangeJSON) error {
	if transactionStateChangeJSON.Creation != "" {
		value, success := new(big.Int).SetString(util.PreUnmarshalHexString(transactionStateChangeJSON.Creation), 16)
		if !success {
			return errors.New("creation invalid")
		}
		t.To = value
	}

	if transactionStateChangeJSON.Deletion != "" {
		value, success := new(big.Int).SetString(util.PreUnmarshalHexString(transactionStateChangeJSON.Deletion), 16)
		if !success {
			return errors.New("deletion invalid")
		}
		t.From = value
	}

	if transactionStateChangeJSON.Alteration != nil {
		from, success := new(big.Int).SetString(util.PreUnmarshalHexString(transactionStateChangeJSON.Alteration.From), 16)
		if !success {
			return errors.New("from invalid")
		}
		t.From = from

		to, success := new(big.Int).SetString(util.PreUnmarshalHexString(transactionStateChangeJSON.Alteration.To), 16)
		if !success {
			return errors.New("to invalid")
		}
		t.To = to
	}

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (t *TransactionStateChange) MarshalYAML() ([]byte, error) {
	if t.From == nil {
		return yaml.Marshal(&transactionStateChangeYAML{
			Creation: util.MarshalBigInt(t.To),
		})
	}

	if t.To == nil {
		return yaml.Marshal(&transactionStateChangeYAML{
			Deletion: util.MarshalBigInt(t.From),
		})
	}

	return yaml.Marshal(&transactionStateChangeYAML{
		Alteration: &transactionStateChangeAlterationYAML{
			From: t.From,
			To:   t.To,
		},
	})
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *TransactionStateChange) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var transactionStateBalanceDiffJSON transactionStateChangeJSON
	if err := yaml.Unmarshal(input, &transactionStateBalanceDiffJSON); err != nil {
		return err
	}
	return t.unpack(&transactionStateBalanceDiffJSON)
}

// String returns a string version of the structure.
func (t *TransactionStateChange) String() string {
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
