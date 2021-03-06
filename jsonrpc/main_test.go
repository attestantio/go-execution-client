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

package jsonrpc_test

import (
	"encoding/hex"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/attestantio/go-execution-client/types"
	"github.com/rs/zerolog"
)

// timeout for tests.
var timeout = 60 * time.Second

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	if os.Getenv("JSONRPC_ADDRESS") != "" {
		os.Exit(m.Run())
	}
}

// strToHash is a helper to create a hash given a string representation.
func strToHash(input string) types.Hash {
	bytes, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}

	var hash types.Hash
	copy(hash[:], bytes)

	return hash
}
