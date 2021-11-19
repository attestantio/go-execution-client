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
	"strconv"

	"github.com/attestantio/go-execution-client/util"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// SyncState contains the sync state.
type SyncState struct {
	CurrentBlock  uint32
	HighestBlock  uint32
	StartingBlock uint32
	Syncing       bool
}

// syncStateJSON is the spec representation of the struct.
type syncStateJSON struct {
	CurrentBlock  string `json:"currentBlock,omitempty"`
	HighestBlock  string `json:"highestBlock,omitempty"`
	StartingBlock string `json:"startingBlock,omitempty"`
	Syncing       bool   `json:"syncing"`
}

// syncStateYAML is the spec representation of the struct.
type syncStateYAML struct {
	CurrentBlock  uint32 `yaml:"currentBlock,omitempty"`
	HighestBlock  uint32 `yaml:"highestBlock,omitempty"`
	StartingBlock uint32 `yaml:"startingBlock,omitempty"`
	Syncing       bool   `yaml:"syncing"`
}

// MarshalJSON implements json.Marshaler.
func (s *SyncState) MarshalJSON() ([]byte, error) {
	return json.Marshal(&syncStateJSON{
		CurrentBlock:  util.MarshalNullableUint32(s.CurrentBlock),
		HighestBlock:  util.MarshalNullableUint32(s.HighestBlock),
		StartingBlock: util.MarshalNullableUint32(s.StartingBlock),
		Syncing:       s.Syncing,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *SyncState) UnmarshalJSON(input []byte) error {
	// May be a simple bool, or an object.
	if bytes.HasPrefix(input, []byte("{")) {
		// It's an object.
		s.Syncing = true
		var data syncStateJSON
		if err := json.Unmarshal(input, &data); err != nil {
			return errors.Wrap(err, "invalid JSON")
		}
		return s.unpack(&data)
	}

	// It's a simple bool.
	if string(input) == "false" {
		s.CurrentBlock = 0
		s.HighestBlock = 0
		s.StartingBlock = 0
		s.Syncing = false
		return nil
	}

	return fmt.Errorf("invalid sync state JSON %q", string(input))
}

func (s *SyncState) unpack(data *syncStateJSON) error {
	var err error

	if data.CurrentBlock == "" {
		return errors.New("current block missing")
	}
	tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(data.CurrentBlock), 16, 32)
	if err != nil {
		return errors.Wrap(err, "current block invalid")
	}
	s.CurrentBlock = uint32(tmp)

	if data.HighestBlock == "" {
		return errors.New("highest block missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.HighestBlock), 16, 32)
	if err != nil {
		return errors.Wrap(err, "highest block invalid")
	}
	s.HighestBlock = uint32(tmp)

	if data.StartingBlock == "" {
		return errors.New("starting block missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.StartingBlock), 16, 32)
	if err != nil {
		return errors.Wrap(err, "starting block invalid")
	}
	s.StartingBlock = uint32(tmp)

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (s *SyncState) MarshalYAML() ([]byte, error) {
	yamlBytes, err := yaml.MarshalWithOptions(&syncStateYAML{
		CurrentBlock:  s.CurrentBlock,
		HighestBlock:  s.HighestBlock,
		StartingBlock: s.StartingBlock,
		Syncing:       s.Syncing,
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (s *SyncState) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var data syncStateJSON
	if err := yaml.Unmarshal(input, &data); err != nil {
		return err
	}
	return s.unpack(&data)
}

// String returns a string version of the structure.
func (s *SyncState) String() string {
	data, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
