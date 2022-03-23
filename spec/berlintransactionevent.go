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
	"github.com/pkg/errors"
)

// BerlinTransactionEvent contains a transaction event.
type BerlinTransactionEvent struct {
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

// berlinTransactionEventJSON is the spec representation of the struct.
// Non-alphabetical ordering to match JSON-RPC output.
type berlinTransactionEventJSON struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	Index            string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}

// MarshalJSON implements json.Marshaler.
func (t *BerlinTransactionEvent) MarshalJSON() ([]byte, error) {
	topics := make([]string, 0, len(t.Topics))
	for _, topic := range t.Topics {
		topics = append(topics, util.MarshalByteArray(topic[:]))
	}

	return json.Marshal(&berlinTransactionEventJSON{
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
func (t *BerlinTransactionEvent) UnmarshalJSON(input []byte) error {
	var data berlinTransactionEventJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	return t.unpack(&data)
}

func (t *BerlinTransactionEvent) unpack(data *berlinTransactionEventJSON) error {
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

	if len(data.Data) > 0 {
		t.Data, err = hex.DecodeString(util.PreUnmarshalHexString(data.Data))
		if err != nil {
			return errors.Wrap(err, "data invalid")
		}
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

// String returns a string version of the structure.
func (t *BerlinTransactionEvent) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(bytes.TrimSuffix(data, []byte("\n")))
}
