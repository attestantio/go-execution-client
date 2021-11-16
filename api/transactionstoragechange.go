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
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/attestantio/go-execution-client/util"
	"github.com/goccy/go-yaml"
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

// transactionStorageChangeYAML is the spec representation of the struct.
type transactionStorageChangeYAML struct {
	Creation   string                                  `yaml:"+,omitempty"`
	Alteration *transactionStorageChangeAlterationYAML `yaml:"*,omitempty"`
	Deletion   string                                  `yaml:"-,omitempty"`
}

type transactionStorageChangeAlterationYAML struct {
	From string `yaml:"from,omitempty"`
	To   string `yaml:"to,omitempty"`
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
	var transactionStateBalanceDiffJSON transactionStorageChangeJSON
	if err := json.Unmarshal(input, &transactionStateBalanceDiffJSON); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&transactionStateBalanceDiffJSON)
}

func (t *TransactionStorageChange) unpack(transactionStorageChangeJSON *transactionStorageChangeJSON) error {
	var err error

	if transactionStorageChangeJSON.Creation != "" {
		t.To, err = hex.DecodeString(util.PreUnmarshalHexString(transactionStorageChangeJSON.Creation))
		if err != nil {
			return errors.Wrap(err, "creation invalid")
		}
	}

	if transactionStorageChangeJSON.Deletion != "" {
		t.From, err = hex.DecodeString(util.PreUnmarshalHexString(transactionStorageChangeJSON.Deletion))
		if err != nil {
			return errors.Wrap(err, "deletion invalid")
		}
	}

	if transactionStorageChangeJSON.Alteration != nil {
		t.From, err = hex.DecodeString(util.PreUnmarshalHexString(transactionStorageChangeJSON.Alteration.From))
		if err != nil {
			return errors.Wrap(err, "from invalid")
		}

		t.To, err = hex.DecodeString(util.PreUnmarshalHexString(transactionStorageChangeJSON.Alteration.To))
		if err != nil {
			return errors.Wrap(err, "to invalid")
		}
	}

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (t *TransactionStorageChange) MarshalYAML() ([]byte, error) {
	if t.From == nil {
		return yaml.Marshal(&transactionStorageChangeYAML{
			Creation: util.MarshalByteArray(t.To),
		})
	}

	if t.To == nil {
		return yaml.Marshal(&transactionStorageChangeYAML{
			Deletion: util.MarshalByteArray(t.From),
		})
	}

	return yaml.Marshal(&transactionStorageChangeYAML{
		Alteration: &transactionStorageChangeAlterationYAML{
			From: util.MarshalByteArray(t.From),
			To:   util.MarshalByteArray(t.To),
		},
	})
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *TransactionStorageChange) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var transactionStateBalanceDiffJSON transactionStorageChangeJSON
	if err := yaml.Unmarshal(input, &transactionStateBalanceDiffJSON); err != nil {
		return err
	}
	return t.unpack(&transactionStateBalanceDiffJSON)
}

// String returns a string version of the structure.
func (t *TransactionStorageChange) String() string {
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
