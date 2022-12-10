// Copyright © 2022 Attestant Limited.
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

package spec_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/stretchr/testify/require"
)

// TestWithdrawalJSON tests the JSON encoding of withdrawals.
func TestWithdrawalJSON(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   string
	}{
		{
			name: "Empty",
			err:  "unexpected end of JSON input",
		},
		{
			name:  "JSONBad",
			input: []byte("[]"),
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type spec.withdrawalJSON",
		},
		{
			name:  "IndexMissing",
			input: []byte(`{"validatorIndex":"0x72a","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"0x365cc84731400"}`),
			err:   "index missing",
		},
		{
			name:  "IndexWrongType",
			input: []byte(`{"index":true,"validatorIndex":"0x72a","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"0x365cc84731400"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field withdrawalJSON.index of type string",
		},
		{
			name:  "IndexInvalid",
			input: []byte(`{"index":"true","validatorIndex":"0x72a","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"0x365cc84731400"}`),
			err:   "index invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "ValidatorIndexMissing",
			input: []byte(`{"index":"0x49fb","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"0x365cc84731400"}`),
			err:   "validator index missing",
		},
		{
			name:  "ValidatorIndexWrongType",
			input: []byte(`{"index":"0x49fb","validatorIndex":true,"address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"0x365cc84731400"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field withdrawalJSON.validatorIndex of type string",
		},
		{
			name:  "ValidatorIndexInvalid",
			input: []byte(`{"index":"0x49fb","validatorIndex":"true","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"0x365cc84731400"}`),
			err:   "validator index invalid: strconv.ParseUint: parsing \"true\": invalid syntax",
		},
		{
			name:  "AddressMissing",
			input: []byte(`{"index":"0x49fb","validatorIndex":"0x72a","amount":"0x365cc84731400"}`),
			err:   "address missing",
		},
		{
			name:  "AddressWrongType",
			input: []byte(`{"index":"0x49fb","validatorIndex":"0x72a","address":true,"amount":"0x365cc84731400"}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field withdrawalJSON.address of type string",
		},
		{
			name:  "AddressInvalid",
			input: []byte(`{"index":"0x49fb","validatorIndex":"0x72a","address":"true","amount":"0x365cc84731400"}`),
			err:   "address invalid: encoding/hex: invalid byte: U+0074 't'",
		},
		{
			name:  "AddressWrongLength",
			input: []byte(`{"index":"0x49fb","validatorIndex":"0x72a","address":"0x8ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"0x365cc84731400"}`),
			err:   "incorrect length 19 for address",
		},
		{
			name:  "AmountMissing",
			input: []byte(`{"index":"0x49fb","validatorIndex":"0x72a","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5"}`),
			err:   "amount missing",
		},
		{
			name:  "AmountWrongType",
			input: []byte(`{"index":"0x49fb","validatorIndex":"0x72a","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":true}`),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field withdrawalJSON.amount of type string",
		},
		{
			name:  "AmountInvalid",
			input: []byte(`{"index":"0x49fb","validatorIndex":"0x72a","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"true"}`),
			err:   "amount invalid",
		},
		{
			name:  "Good",
			input: []byte(`{"index":"0x49fb","validatorIndex":"0x72a","address":"0x388ea662ef2c223ec0b047d41bf3c0f362142ad5","amount":"0x365cc84731400"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.Withdrawal
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				require.Equal(t, string(test.input), string(rt))
				require.Equal(t, string(test.input), res.String())
			}
		})
	}
}
