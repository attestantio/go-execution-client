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
	"time"

	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/pkg/errors"
)

// PragueBlock contains a block after the Prague hardfork.
type PragueBlock struct {
	BaseFeePerGas         uint64
	BlobGasUsed           uint64
	Difficulty            uint64
	ExcessBlobGas         uint64
	ExtraData             []byte
	GasLimit              uint32
	GasUsed               uint32
	Hash                  types.Hash
	LogsBloom             []byte
	Miner                 types.Address
	MixHash               types.Hash
	Nonce                 []byte
	Number                uint32
	ParentBeaconBlockRoot types.Root
	ParentHash            types.Hash
	ReceiptsRoot          types.Root
	RequestsHash          types.Hash
	SHA3Uncles            []byte
	Size                  uint32
	StateRoot             types.Root
	Timestamp             time.Time
	TotalDifficulty       *big.Int
	Transactions          []*Transaction
	TransactionsRoot      types.Root
	Uncles                []types.Hash
	Withdrawals           []*Withdrawal
	WithdrawalsRoot       types.Root
}

// pragueBlockJSON is the spec representation of the struct.
type pragueBlockJSON struct {
	BaseFeePerGas         string         `json:"baseFeePerGas"`
	BlobGasUsed           string         `json:"blobGasUsed"`
	Difficulty            string         `json:"difficulty"`
	ExcessBlobGas         string         `json:"excessBlobGas"`
	ExtraData             string         `json:"extraData"`
	GasLimit              string         `json:"gasLimit"`
	GasUsed               string         `json:"gasUsed"`
	Hash                  string         `json:"hash"`
	LogsBloom             string         `json:"logsBloom"`
	Miner                 string         `json:"miner"`
	MixHash               string         `json:"mixHash"`
	Nonce                 string         `json:"nonce"`
	Number                string         `json:"number"`
	ParentBeaconBlockRoot string         `json:"parentBeaconBlockRoot"`
	ParentHash            string         `json:"parentHash"`
	ReceiptsRoot          string         `json:"receiptsRoot"`
	RequestsHash          string         `json:"requestsHash"`
	SHA3Uncles            string         `json:"sha3Uncles"`
	Size                  string         `json:"size"`
	StateRoot             string         `json:"stateRoot"`
	Timestamp             string         `json:"timestamp"`
	TotalDifficulty       string         `json:"totalDifficulty"`
	Transactions          []*Transaction `json:"transactions"`
	TransactionsRoot      string         `json:"transactionsRoot"`
	Uncles                []string       `json:"uncles"`
	Withdrawals           []*Withdrawal  `json:"withdrawals"`
	WithdrawalsRoot       string         `json:"withdrawalsRoot"`
}

// MarshalJSON implements json.Marshaler.
func (b *PragueBlock) MarshalJSON() ([]byte, error) {
	uncles := make([]string, 0, len(b.Uncles))
	for _, uncle := range b.Uncles {
		uncles = append(uncles, fmt.Sprintf("%#x", uncle))
	}

	return json.Marshal(&pragueBlockJSON{
		BaseFeePerGas:         util.MarshalUint64(b.BaseFeePerGas),
		BlobGasUsed:           util.MarshalUint64(b.BlobGasUsed),
		Difficulty:            util.MarshalUint64(b.Difficulty),
		ExcessBlobGas:         util.MarshalUint64(b.ExcessBlobGas),
		ExtraData:             util.MarshalByteArray(b.ExtraData),
		GasLimit:              util.MarshalUint32(b.GasLimit),
		GasUsed:               util.MarshalUint32(b.GasUsed),
		Hash:                  util.MarshalByteArray(b.Hash[:]),
		LogsBloom:             util.MarshalByteArray(b.LogsBloom),
		Miner:                 util.MarshalByteArray(b.Miner[:]),
		MixHash:               util.MarshalByteArray(b.MixHash[:]),
		Nonce:                 util.MarshalByteArray(b.Nonce),
		Number:                util.MarshalUint32(b.Number),
		ParentBeaconBlockRoot: util.MarshalByteArray(b.ParentBeaconBlockRoot[:]),
		ParentHash:            util.MarshalByteArray(b.ParentHash[:]),
		ReceiptsRoot:          util.MarshalByteArray(b.ReceiptsRoot[:]),
		RequestsHash:          util.MarshalByteArray(b.RequestsHash[:]),
		SHA3Uncles:            util.MarshalByteArray(b.SHA3Uncles),
		Size:                  util.MarshalUint32(b.Size),
		StateRoot:             util.MarshalByteArray(b.StateRoot[:]),
		Timestamp:             fmt.Sprintf("%#x", b.Timestamp.Unix()),
		TotalDifficulty:       util.MarshalBigInt(b.TotalDifficulty),
		Transactions:          b.Transactions,
		TransactionsRoot:      util.MarshalByteArray(b.TransactionsRoot[:]),
		Uncles:                uncles,
		Withdrawals:           b.Withdrawals,
		WithdrawalsRoot:       util.MarshalByteArray(b.WithdrawalsRoot[:]),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
//
//nolint:gocyclo
func (b *PragueBlock) UnmarshalJSON(input []byte) error {
	var data pragueBlockJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	var (
		success bool
		err     error
	)

	// Although base fee per gas is required in Prague, this also covers pre-Prague blocks so it is considered optional.

	if data.BaseFeePerGas != "" {
		b.BaseFeePerGas, err = strconv.ParseUint(util.PreUnmarshalHexString(data.BaseFeePerGas), 16, 64)
		if err != nil {
			return errors.Wrap(err, "base fee per gas invalid")
		}
	}

	if data.BlobGasUsed == "" {
		return errors.New("blob gas used missing")
	}

	tmp, err := strconv.ParseUint(util.PreUnmarshalHexString(data.BlobGasUsed), 16, 32)
	if err != nil {
		return errors.Wrap(err, "blob gas used invalid")
	}

	b.BlobGasUsed = tmp

	if data.Difficulty == "" {
		return errors.New("difficulty missing")
	}

	b.Difficulty, err = strconv.ParseUint(util.PreUnmarshalHexString(data.Difficulty), 16, 64)
	if err != nil {
		return errors.Wrap(err, "difficulty invalid")
	}

	if data.ExcessBlobGas == "" {
		return errors.New("excess blob gas missing")
	}

	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.ExcessBlobGas), 16, 32)
	if err != nil {
		return errors.Wrap(err, "excess blob gas invalid")
	}

	b.ExcessBlobGas = tmp

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

	tmp, err = strconv.ParseUint(util.PreUnmarshalHexString(data.GasUsed), 16, 32)
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

	if data.ParentBeaconBlockRoot == "" {
		return errors.New("parent beacon block root missing")
	}

	root, err := hex.DecodeString(util.PreUnmarshalHexString(data.ParentBeaconBlockRoot))
	if err != nil {
		return errors.Wrap(err, "parent beacon block root invalid")
	}

	copy(b.ParentBeaconBlockRoot[:], root)

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

	root, err = hex.DecodeString(util.PreUnmarshalHexString(data.ReceiptsRoot))
	if err != nil {
		return errors.Wrap(err, "receipts root invalid")
	}

	copy(b.ReceiptsRoot[:], root)

	if data.RequestsHash == "" {
		return errors.New("requests hash missing")
	}

	hash, err = hex.DecodeString(util.PreUnmarshalHexString(data.RequestsHash))
	if err != nil {
		return errors.Wrap(err, "requests hash invalid")
	}

	copy(b.RequestsHash[:], hash)

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

	if data.TotalDifficulty != "" {
		b.TotalDifficulty, success = new(big.Int).SetString(util.PreUnmarshalHexString(data.TotalDifficulty), 16)
		if !success {
			return errors.New("total difficulty invalid")
		}
	}

	b.Transactions = data.Transactions

	if data.TransactionsRoot == "" {
		return errors.New("transactions root missing")
	}

	root, err = hex.DecodeString(util.PreUnmarshalHexString(data.TransactionsRoot))
	if err != nil {
		return errors.Wrap(err, "transactions root invalid")
	}

	copy(b.TransactionsRoot[:], root)

	b.Uncles = make([]types.Hash, len(data.Uncles))
	for i, uncleStr := range data.Uncles {
		if uncleStr == "" {
			return errors.New("uncle missing")
		}

		hash, err := hex.DecodeString(util.PreUnmarshalHexString(uncleStr))
		if err != nil {
			return errors.Wrap(err, "uncle invalid")
		}

		copy(b.Uncles[i][:], hash)
	}

	b.Withdrawals = data.Withdrawals

	if data.WithdrawalsRoot == "" {
		return errors.New("withdrawals root missing")
	}

	root, err = hex.DecodeString(util.PreUnmarshalHexString(data.WithdrawalsRoot))
	if err != nil {
		return errors.Wrap(err, "withdrawals root invalid")
	}

	copy(b.WithdrawalsRoot[:], root)

	return nil
}

// String returns a string version of the structure.
func (b *PragueBlock) String() string {
	data, err := json.Marshal(b)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(bytes.TrimSuffix(data, []byte("\n")))
}
