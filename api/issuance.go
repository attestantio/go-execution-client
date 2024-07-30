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
	var err error

	if data.BlockReward == "" {
		i.BlockReward = zero
	} else {
		i.BlockReward, err = util.StrToBigInt("block reward", data.BlockReward)
		if err != nil {
			return err
		}
	}

	if data.UncleReward == "" {
		i.UncleReward = zero
	} else {
		i.UncleReward, err = util.StrToBigInt("uncle reward", data.UncleReward)
		if err != nil {
			return err
		}
	}

	if data.Issuance == "" {
		i.Issuance = zero
	} else {
		i.Issuance, err = util.StrToBigInt("issuance", data.Issuance)
		if err != nil {
			return err
		}
	}

	return nil
}

// String returns a string version of the structure.
func (i *Issuance) String() string {
	data, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
