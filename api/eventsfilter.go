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
	"strings"

	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// EventsFilter contains the events filter.
type EventsFilter struct {
	FromBlock string
	ToBlock   string
	Address   *types.Address
	Topics    []types.Hash
}

// eventsFilterJSON is the spec representation of the struct.
type eventsFilterJSON struct {
	FromBlock string   `json:"fromBlock,omitempty"`
	ToBlock   string   `json:"toBlock,omitempty"`
	Address   string   `json:"address,omitempty"`
	Topics    []string `json:"topics,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (e *EventsFilter) MarshalJSON() ([]byte, error) {
	eventsFilterJSON := &eventsFilterJSON{
		FromBlock: e.FromBlock,
		ToBlock:   e.ToBlock,
	}

	if e.Address != nil {
		eventsFilterJSON.Address = util.MarshalAddress((*e.Address)[:])
	}

	topics := make([]string, 0, len(e.Topics))
	for _, topic := range e.Topics {
		topics = append(topics, util.MarshalByteArray(topic[:]))
	}

	eventsFilterJSON.Topics = topics

	return json.Marshal(eventsFilterJSON)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *EventsFilter) UnmarshalJSON(input []byte) error {
	var eventsFilterJSON eventsFilterJSON
	if err := json.Unmarshal(input, &eventsFilterJSON); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return e.unpack(&eventsFilterJSON)
}

// String returns a string version of the structure.
func (e *EventsFilter) String() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}

func (e *EventsFilter) unpack(data *eventsFilterJSON) error {
	switch strings.ToLower(data.FromBlock) {
	case "":
		// Nothing to do.
	case "pending", "latest", "safe", "finalized":
		// State name.
		e.FromBlock = data.FromBlock
	default:
		// Block number.
		if _, err := util.StrToUint32("from block", data.FromBlock); err != nil {
			return err
		}

		e.FromBlock = data.FromBlock
	}

	switch strings.ToLower(data.ToBlock) {
	case "":
		// Nothing to do.
	case "pending", "latest", "safe", "finalized":
		// State name.
		e.ToBlock = data.ToBlock
	default:
		// Block number.
		if _, err := util.StrToUint32("to block", data.ToBlock); err != nil {
			return err
		}

		e.ToBlock = data.ToBlock
	}

	if data.Address != "" {
		address, err := util.StrToAddress("address", data.Address)
		if err != nil {
			return err
		}

		e.Address = &address
	}

	if data.Topics != nil {
		var err error

		topics := make([]types.Hash, len(data.Topics))
		for i, topic := range data.Topics {
			topics[i], err = util.StrToHash("topic", topic)
			if err != nil {
				return err
			}
		}

		e.Topics = topics
	}

	return nil
}
