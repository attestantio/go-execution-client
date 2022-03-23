// Copyright © 2021, 2022 Attestant Limited.
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

	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// AccessListEntry contains a single entry in an access list.
type AccessListEntry struct {
	Address     []byte
	StorageKeys [][]byte
}

// accessListEntryJSON is the spec representation of the struct.
type accessListEntryJSON struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

// MarshalJSON implements json.Marshaler.
func (a *AccessListEntry) MarshalJSON() ([]byte, error) {
	storageKeys := make([]string, 0, len(a.StorageKeys))
	for _, storageKey := range a.StorageKeys {
		storageKeys = append(storageKeys, util.MarshalByteArray(storageKey))
	}

	return json.Marshal(&accessListEntryJSON{
		Address:     util.MarshalByteArray(a.Address),
		StorageKeys: storageKeys,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *AccessListEntry) UnmarshalJSON(input []byte) error {
	var data accessListEntryJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return a.unpack(&data)
}

func (a *AccessListEntry) unpack(data *accessListEntryJSON) error {
	var err error

	if data.Address == "" {
		return errors.New("address missing")
	}
	a.Address, err = hex.DecodeString(util.PreUnmarshalHexString(data.Address))
	if err != nil {
		return errors.Wrap(err, "address invalid")
	}

	a.StorageKeys = make([][]byte, 0, len(data.StorageKeys))
	for _, storageKey := range data.StorageKeys {
		key, err := hex.DecodeString(util.PreUnmarshalHexString(storageKey))
		if err != nil {
			return errors.Wrap(err, "storage key invalid")
		}
		a.StorageKeys = append(a.StorageKeys, key)
	}

	return nil
}

// String returns a string version of the structure.
func (a *AccessListEntry) String() string {
	data, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
