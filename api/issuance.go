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
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/attestantio/go-execution-client/util"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// Issuance contains issuance for a block.
type Issuance struct {
	BlockReward *big.Int
	UncleReward *big.Int
	Issuance    *big.Int
}

// issuanceJSON is the spec representation of the struct.
type issuanceJSON struct {
	BlockReward string `json:"blockReward"`
	UncleReward string `json:"uncleReward"`
	Issuance    string `json:"issuance"`
}

// issuanceYAML is the spec representation of the struct.
type issuanceYAML struct {
	BlockReward *big.Int `yaml:"blockReward"`
	UncleReward *big.Int `yaml:"uncleReward"`
	Issuance    *big.Int `yaml:"issuance"`
}

// MarshalJSON implements json.Marshaler.
func (i *Issuance) MarshalJSON() ([]byte, error) {
	return json.Marshal(&issuanceJSON{
		BlockReward: util.MarshalBigInt(i.BlockReward),
		UncleReward: util.MarshalBigInt(i.UncleReward),
		Issuance:    util.MarshalBigInt(i.Issuance),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Issuance) UnmarshalJSON(input []byte) error {
	var data issuanceJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}
	return i.unpack(&data)
}

func (i *Issuance) unpack(data *issuanceJSON) error {
	var success bool

	if data.BlockReward == "" {
		i.BlockReward = zero
	} else {
		i.BlockReward, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.BlockReward), 16)
		if !success {
			return errors.New("block reward invalid")
		}
	}

	if data.UncleReward == "" {
		i.UncleReward = zero
	} else {
		i.UncleReward, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.UncleReward), 16)
		if !success {
			return errors.New("uncle reward invalid")
		}
	}

	if data.Issuance == "" {
		i.Issuance = zero
	} else {
		i.Issuance, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.Issuance), 16)
		if !success {
			return errors.New("issuance invalid")
		}
	}

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (i *Issuance) MarshalYAML() ([]byte, error) {
	yamlBytes, err := yaml.MarshalWithOptions(&issuanceYAML{
		BlockReward: i.BlockReward,
		UncleReward: i.UncleReward,
		Issuance:    i.Issuance,
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (i *Issuance) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var data issuanceJSON
	if err := yaml.Unmarshal(input, &data); err != nil {
		return err
	}
	return i.unpack(&data)
}

// String returns a string version of the structure.
func (i *Issuance) String() string {
	data, err := yaml.Marshal(i)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
