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

package spec

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// TransactionEvent contains a transaction event.
type TransactionEvent struct {
	Address          types.Address
	BlockHash        types.Hash
	BlockNumber      uint32
	Data             []byte
	Index            uint32
	Removed          bool
	Topics           []types.Hash
	TransactionHash  types.Hash
	TransactionIndex uint32
}

// transactionEventJSON is the spec representation of the struct.
type transactionEventJSON struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	Index            string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

// transactionEventYAML is the spec representation of the struct.
type transactionEventYAML struct {
	Address          string   `yaml:"address"`
	BlockHash        string   `yaml:"blockHash"`
	BlockNumber      uint32   `yaml:"blockNumber"`
	Data             string   `yaml:"data"`
	Index            uint32   `yaml:"index"`
	Removed          bool     `yaml:"removed"`
	Topics           []string `yaml:"topics"`
	TransactionHash  string   `yaml:"transactionHash"`
	TransactionIndex uint32   `yaml:"transactionNumber"`
}

// MarshalJSON implements json.Marshaler.
func (t *TransactionEvent) MarshalJSON() ([]byte, error) {
	topics := make([]string, 0, len(t.Topics))
	for _, topic := range t.Topics {
		topics = append(topics, util.MarshalByteArray(topic[:]))
	}

	return json.Marshal(&transactionEventJSON{
		Address:          util.MarshalAddress(t.Address[:]),
		BlockHash:        util.MarshalByteArray(t.BlockHash[:]),
		BlockNumber:      util.MarshalUint32(t.BlockNumber),
		Data:             util.MarshalByteArray(t.Data),
		Index:            util.MarshalUint32(t.Index),
		Removed:          t.Removed,
		Topics:           topics,
		TransactionHash:  util.MarshalByteArray(t.TransactionHash[:]),
		TransactionIndex: util.MarshalUint32(t.TransactionIndex),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionEvent) UnmarshalJSON(input []byte) error {
	var data transactionEventJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

func (t *TransactionEvent) unpack(data *transactionEventJSON) error {
	var err error

	if data.Address == "" {
		return errors.New("address missing")
	}
	address, err := hex.DecodeString(util.PreUnmarshalHexString(data.Address))
	if err != nil {
		return errors.Wrap(err, "address invalid")
	}
	copy(t.Address[:], address)

	if data.BlockHash == "" {
		return errors.New("block hash missing")
	}
	hash, err := hex.DecodeString(util.PreUnmarshalHexString(data.BlockHash))
	if err != nil {
		return errors.Wrap(err, "block hash invalid")
	}
	copy(t.BlockHash[:], hash)

	if data.BlockNumber == "" {
		return errors.New("block number missing")
	}
	tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(data.BlockNumber), 16, 32)
	if err != nil {
		return errors.Wrap(err, "block number invalid")
	}
	t.BlockNumber = uint32(tmp)

	if data.Data == "" {
		return errors.New("data missing")
	}
	t.Data, err = hex.DecodeString(util.PreUnmarshalHexString(data.Data))
	if err != nil {
		return errors.Wrap(err, "data invalid")
	}
	if len(t.Data) == 0 {
		t.Data = nil
	}

	if data.Index == "" {
		return errors.New("log index missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Index), 16, 32)
	if err != nil {
		return errors.Wrap(err, "log index invalid")
	}
	t.Index = uint32(tmp)

	t.Removed = data.Removed

	topics := make([]types.Hash, len(data.Topics))
	for i, topic := range data.Topics {
		hash, err := hex.DecodeString(util.PreUnmarshalHexString(topic))
		if err != nil {
			return errors.Wrap(err, "topic invalid")
		}
		copy(topics[i][:], hash)
	}
	t.Topics = topics

	if data.TransactionHash == "" {
		return errors.New("transaction hash missing")
	}
	hash, err = hex.DecodeString(util.PreUnmarshalHexString(data.TransactionHash))
	if err != nil {
		return errors.Wrap(err, "transaction hash invalid")
	}
	copy(t.TransactionHash[:], hash)

	if data.TransactionIndex == "" {
		return errors.New("transaction index missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.TransactionIndex), 16, 32)
	if err != nil {
		return errors.Wrap(err, "transaction index invalid")
	}
	t.TransactionIndex = uint32(tmp)

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (t *TransactionEvent) MarshalYAML() ([]byte, error) {
	topics := make([]string, 0, len(t.Topics))
	for _, topic := range t.Topics {
		topics = append(topics, fmt.Sprintf("%#x", topic))
	}

	yamlBytes, err := yaml.MarshalWithOptions(&transactionEventYAML{
		TransactionHash:  fmt.Sprintf("%#x", t.TransactionHash),
		TransactionIndex: t.TransactionIndex,
		BlockHash:        fmt.Sprintf("%#x", t.BlockHash),
		BlockNumber:      t.BlockNumber,
		Removed:          t.Removed,
		Index:            t.Index,
		Address:          fmt.Sprintf("%#x", t.Address),
		Data:             fmt.Sprintf("%#x", t.Data),
		Topics:           topics,
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (t *TransactionEvent) UnmarshalYAML(input []byte) error {
	// We unmarshal to the JSON struct to save on duplicate code.
	var transactionEventJSON transactionEventJSON
	if err := yaml.Unmarshal(input, &transactionEventJSON); err != nil {
		return err
	}
	return t.unpack(&transactionEventJSON)
}

// String returns a string version of the structure.
func (t *TransactionEvent) String() string {
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
