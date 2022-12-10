// Copyright © 2021, 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// Withdrawal is the spec representation of a withdrawal.
type Withdrawal struct {
	Index          uint64
	ValidatorIndex uint64
	Address        types.Address
	Amount         *big.Int
}

// withdrawalJSON is the spec representation of a type 0 transaction.
type withdrawalJSON struct {
	Index          string `json:"index"`
	ValidatorIndex string `json:"validatorIndex"`
	Address        string `json:"address"`
	Amount         string `json:"amount"`
}

// MarshalJSON marshals a type 0 transaction.
func (w *Withdrawal) MarshalJSON() ([]byte, error) {
	return json.Marshal(&withdrawalJSON{
		Index:          util.MarshalUint64(w.Index),
		ValidatorIndex: util.MarshalUint64(w.ValidatorIndex),
		Address:        util.MarshalAddress(w.Address[:]),
		Amount:         util.MarshalBigInt(w.Amount),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (w *Withdrawal) UnmarshalJSON(input []byte) error {
	var data withdrawalJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return w.unpack(&data)
}

// nolint:gocyclo
func (w *Withdrawal) unpack(data *withdrawalJSON) error {
	var err error
	var success bool

	if data.Index == "" {
		return errors.New("index missing")
	}
	w.Index, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Index), 16, 64)
	if err != nil {
		return errors.Wrap(err, "index invalid")
	}

	if data.ValidatorIndex == "" {
		return errors.New("validator index missing")
	}
	w.ValidatorIndex, err = strconv.ParseUint(util.PreUnmarshalHexString(data.ValidatorIndex), 16, 64)
	if err != nil {
		return errors.Wrap(err, "validator index invalid")
	}

	if data.Address == "" {
		return errors.New("address missing")
	}
	address, err := hex.DecodeString(util.PreUnmarshalHexString(data.Address))
	if err != nil {
		return errors.Wrap(err, "address invalid")
	}
	if len(address) != len(w.Address) {
		return fmt.Errorf("incorrect length %d for address", len(address))
	}
	copy(w.Address[:], address)

	if data.Amount == "" {
		return errors.New("amount missing")
	}
	w.Amount, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.Amount), 16)
	if !success {
		return errors.New("amount invalid")
	}

	return nil
}

// String returns a string version of the structure.
func (w *Withdrawal) String() string {
	data, err := json.Marshal(w)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
