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

	"github.com/attestantio/go-execution-client/util"
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

	s.CurrentBlock, err = util.StrToUint32("current block", data.CurrentBlock)
	if err != nil {
		return err
	}

	s.HighestBlock, err = util.StrToUint32("highest block", data.HighestBlock)
	if err != nil {
		return err
	}

	s.StartingBlock, err = util.StrToUint32("starting block", data.StartingBlock)
	if err != nil {
		return err
	}

	return nil
}

// String returns a string version of the structure.
func (s *SyncState) String() string {
	data, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
