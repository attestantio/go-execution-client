// Copyright © 2025 Attestant Limited.
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

// AuthorizationListEntry contains a single entry in an authorization list.
type AuthorizationListEntry struct {
	Address types.Address
	ChainID *big.Int
	Nonce   uint64
	R       *big.Int
	S       *big.Int
	YParity *big.Int
}

// authorizationListEntryJSON is the spec representation of the struct.
type authorizationListEntryJSON struct {
	Address string `json:"address"`
	ChainID string `json:"chainId"`
	Nonce   string `json:"nonce"`
	R       string `json:"r"`
	S       string `json:"s"`
	YParity string `json:"yParity,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (a *AuthorizationListEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(&authorizationListEntryJSON{
		Address: util.MarshalByteArray(a.Address[:]),
		ChainID: util.MarshalBigInt(a.ChainID),
		Nonce:   util.MarshalUint64(a.Nonce),
		R:       util.MarshalBigInt(a.R),
		S:       util.MarshalBigInt(a.S),
		YParity: util.MarshalUint64(a.YParity.Uint64()),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *AuthorizationListEntry) UnmarshalJSON(input []byte) error {
	var data authorizationListEntryJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return a.unpack(&data)
}

func (a *AuthorizationListEntry) unpack(data *authorizationListEntryJSON) error {
	var success bool
	var err error

	if data.Address == "" {
		return errors.New("address missing")
	}
	address, err := hex.DecodeString(util.PreUnmarshalHexString(data.Address))
	if err != nil {
		return errors.Wrap(err, "address invalid")
	}
	copy(a.Address[:], address)

	if data.ChainID == "" {
		return errors.New("chain id missing")
	}
	a.ChainID, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.ChainID), 16)
	if !success {
		return errors.New("chain id invalid")
	}

	if data.Nonce == "" {
		return errors.New("nonce missing")
	}
	a.Nonce, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Nonce), 16, 64)
	if err != nil {
		return errors.Wrap(err, "nonce invalid")
	}

	if data.R == "" {
		return errors.New("r missing")
	}
	a.R, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.R), 16)
	if !success {
		return errors.New("r invalid")
	}

	if data.S == "" {
		return errors.New("s missing")
	}
	a.S, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.S), 16)
	if !success {
		return errors.New("s invalid")
	}

	if data.YParity == "" {
		return errors.New("yParity missing")
	}
	a.YParity, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.YParity), 16)
	if !success {
		return errors.New("yParity invalid")
	}

	return nil
}

// String returns a string version of the structure.
func (a *AuthorizationListEntry) String() string {
	data, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(bytes.TrimSuffix(data, []byte("\n")))
}
