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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// LondonBlock contains a block after the London hardfork.
type LondonBlock struct {
	BaseFeePerGas    uint64
	Difficulty       uint64
	ExtraData        []byte
	GasLimit         uint32
	GasUsed          uint32
	Hash             Hash
	LogsBloom        []byte
	Miner            Address
	MixHash          Hash
	Nonce            []byte
	Number           uint32
	ParentHash       Hash
	ReceiptsRoot     Root
	SHA3Uncles       []byte
	Size             uint32
	StateRoot        Root
	Timestamp        time.Time
	TotalDifficulty  *big.Int
	Transactions     []Transaction
	TransactionsRoot Root
	Uncles           []Hash
}

// londonBlockJSON is the spec representation of the struct.
type londonBlockJSON struct {
	BaseFeePerGas    string        `json:"baseFeePerGas"`
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	SHA3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	Transactions     []interface{} `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []string      `json:"uncles"`
}

// MarshalJSON implements json.Marshaler.
func (b *LondonBlock) MarshalJSON() ([]byte, error) {
	transactions := make([]interface{}, 0, len(b.Transactions))
	for _, transaction := range b.Transactions {
		transactions = append(transactions, transaction)
	}

	uncles := make([]string, 0, len(b.Uncles))
	for _, uncle := range b.Uncles {
		uncles = append(uncles, fmt.Sprintf("%#x", uncle))
	}

	return json.Marshal(&londonBlockJSON{
		BaseFeePerGas:    util.MarshalUint64(b.BaseFeePerGas),
		Difficulty:       util.MarshalUint64(b.Difficulty),
		ExtraData:        util.MarshalByteArray(b.ExtraData),
		GasLimit:         util.MarshalUint32(b.GasLimit),
		GasUsed:          util.MarshalUint32(b.GasUsed),
		Hash:             util.MarshalByteArray(b.Hash[:]),
		LogsBloom:        util.MarshalByteArray(b.LogsBloom),
		Miner:            util.MarshalByteArray(b.Miner[:]),
		MixHash:          util.MarshalByteArray(b.MixHash[:]),
		Nonce:            util.MarshalByteArray(b.Nonce),
		Number:           util.MarshalUint32(b.Number),
		ParentHash:       util.MarshalByteArray(b.ParentHash[:]),
		ReceiptsRoot:     util.MarshalByteArray(b.ReceiptsRoot[:]),
		SHA3Uncles:       util.MarshalByteArray(b.SHA3Uncles),
		Size:             util.MarshalUint32(b.Size),
		StateRoot:        util.MarshalByteArray(b.StateRoot[:]),
		Timestamp:        fmt.Sprintf("%#x", b.Timestamp.Unix()),
		TotalDifficulty:  util.MarshalBigInt(b.TotalDifficulty),
		Transactions:     transactions,
		TransactionsRoot: util.MarshalByteArray(b.TransactionsRoot[:]),
		Uncles:           uncles,
	})
}

// UnmarshalJSON implements json.Unmarshaler.
// nolint:gocyclo
func (b *LondonBlock) UnmarshalJSON(input []byte) error {
	var data londonBlockJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	var success bool
	var err error

	// Although base fee per gas is required in London, this also covers pre-London blocks so it is considered optional.
	if data.BaseFeePerGas != "" {
		b.BaseFeePerGas, err = strconv.ParseUint(util.PreUnmarshalHexString(data.BaseFeePerGas), 16, 64)
		if err != nil {
			return errors.Wrap(err, "base fee per gas invalid")
		}
	}

	if data.Difficulty == "" {
		return errors.New("difficulty missing")
	}
	b.Difficulty, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Difficulty), 16, 64)
	if err != nil {
		return errors.Wrap(err, "difficulty invalid")
	}

	if data.ExtraData == "" {
		return errors.New("extra data missing")
	}
	b.ExtraData, err = hex.DecodeString(util.PreUnmarshalHexString(data.ExtraData))
	if err != nil {
		return errors.Wrap(err, "extra data invalid")
	}

	if data.GasUsed == "" {
		return errors.New("gas used missing")
	}
	tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(data.GasUsed), 16, 32)
	if err != nil {
		return errors.Wrap(err, "gas used invalid")
	}
	b.GasUsed = uint32(tmp)

	if data.GasLimit == "" {
		return errors.New("gas limit missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.GasLimit), 16, 32)
	if err != nil {
		return errors.Wrap(err, "gas limit invalid")
	}
	b.GasLimit = uint32(tmp)

	if data.Hash == "" {
		return errors.New("hash missing")
	}
	hash, err := hex.DecodeString(util.PreUnmarshalHexString(data.Hash))
	if err != nil {
		return errors.Wrap(err, "hash invalid")
	}
	copy(b.Hash[:], hash)

	if data.LogsBloom == "" {
		return errors.New("logs bloom missing")
	}
	b.LogsBloom, err = hex.DecodeString(util.PreUnmarshalHexString(data.LogsBloom))
	if err != nil {
		return errors.Wrap(err, "logs bloom invalid")
	}

	if data.Miner == "" {
		return errors.New("miner missing")
	}
	address, err := hex.DecodeString(util.PreUnmarshalHexString(data.Miner))
	if err != nil {
		return errors.Wrap(err, "miner invalid")
	}
	copy(b.Miner[:], address)

	if data.MixHash == "" {
		return errors.New("mix hash missing")
	}
	hash, err = hex.DecodeString(util.PreUnmarshalHexString(data.MixHash))
	if err != nil {
		return errors.Wrap(err, "mix hash invalid")
	}
	copy(b.MixHash[:], hash)

	if data.Nonce == "" {
		return errors.New("nonce missing")
	}
	b.Nonce, err = hex.DecodeString(util.PreUnmarshalHexString(data.Nonce))
	if err != nil {
		return errors.Wrap(err, "nonce invalid")
	}

	if data.Number == "" {
		return errors.New("number missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Number), 16, 32)
	if err != nil {
		return errors.Wrap(err, "number invalid")
	}
	b.Number = uint32(tmp)

	if data.ParentHash == "" {
		return errors.New("parent hash missing")
	}
	hash, err = hex.DecodeString(util.PreUnmarshalHexString(data.ParentHash))
	if err != nil {
		return errors.Wrap(err, "parent hash invalid")
	}
	copy(b.ParentHash[:], hash)

	if data.ReceiptsRoot == "" {
		return errors.New("receipts root missing")
	}
	root, err := hex.DecodeString(util.PreUnmarshalHexString(data.ReceiptsRoot))
	if err != nil {
		return errors.Wrap(err, "receipts root invalid")
	}
	copy(b.ReceiptsRoot[:], root)

	if data.SHA3Uncles == "" {
		return errors.New("sha3 uncles missing")
	}
	b.SHA3Uncles, err = hex.DecodeString(util.PreUnmarshalHexString(data.SHA3Uncles))
	if err != nil {
		return errors.Wrap(err, "sha3 uncles invalid")
	}

	if data.Size == "" {
		return errors.New("size missing")
	}
	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Size), 16, 32)
	if err != nil {
		return errors.Wrap(err, "size invalid")
	}
	b.Size = uint32(tmp)

	if data.StateRoot == "" {
		return errors.New("state root missing")
	}
	root, err = hex.DecodeString(util.PreUnmarshalHexString(data.StateRoot))
	if err != nil {
		return errors.Wrap(err, "state root invalid")
	}
	copy(b.StateRoot[:], root)

	if data.Timestamp == "" {
		return errors.New("timestamp missing")
	}
	timestamp, err := strconv.ParseInt(util.PreUnmarshalHexString(data.Timestamp), 16, 64)
	if err != nil {
		return errors.Wrap(err, "timestamp invalid")
	}
	b.Timestamp = time.Unix(timestamp, 0)

	if data.TotalDifficulty == "" {
		return errors.New("total difficulty missing")
	}
	b.TotalDifficulty, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.TotalDifficulty), 16)
	if !success {
		return errors.New("total difficulty invalid")
	}

	b.Transactions = make([]Transaction, 0, len(data.Transactions))
	for _, tx := range data.Transactions {
		txJSON, err := json.Marshal(tx)
		if err != nil {
			return errors.Wrap(err, "failed to remarshal")
		}
		transaction, err := UnmarshalTransactionJSON(txJSON)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal remarshaled transaction")
		}
		b.Transactions = append(b.Transactions, transaction)
	}

	if data.TransactionsRoot == "" {
		return errors.New("transactions root missing")
	}
	root, err = hex.DecodeString(util.PreUnmarshalHexString(data.TransactionsRoot))
	if err != nil {
		return errors.Wrap(err, "transactions root invalid")
	}
	copy(b.TransactionsRoot[:], root)

	b.Uncles = make([]Hash, len(data.Uncles))
	for i, uncleStr := range data.Uncles {
		if uncleStr == "" {
			return errors.New("uncle invalid")
		}
		hash, err := hex.DecodeString(util.PreUnmarshalHexString(uncleStr))
		if err != nil {
			return errors.Wrap(err, "uncle invalid")
		}
		copy(b.Uncles[i][:], hash)
	}

	return nil
}
