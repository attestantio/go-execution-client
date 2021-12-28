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

package util_test

import (
	"testing"

	"github.com/attestantio/go-execution-client/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRLPBytes tests the RLP bytes encoding function.
func TestRLPBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Nil",
			input:    nil,
			expected: []byte{0x80},
		},
		{
			name:     "Dog",
			input:    []byte("dog"),
			expected: []byte{0x83, 'd', 'o', 'g'},
		},
		{
			name:     "0",
			input:    []byte("\x00"),
			expected: []byte{0x00},
		},
		{
			name:     "15",
			input:    []byte("\x0f"),
			expected: []byte{0x0f},
		},
		{
			name:     "1024",
			input:    []byte("\x04\x00"),
			expected: []byte{0x82, 0x04, 0x00},
		},
		{
			name:     "Lorum",
			input:    []byte("Lorem ipsum dolor sit amet, consectetur adipisicing elit"),
			expected: []byte{0xb8, 0x38, 'L', 'o', 'r', 'e', 'm', ' ', 'i', 'p', 's', 'u', 'm', ' ', 'd', 'o', 'l', 'o', 'r', ' ', 's', 'i', 't', ' ', 'a', 'm', 'e', 't', ',', ' ', 'c', 'o', 'n', 's', 'e', 'c', 't', 'e', 't', 'u', 'r', ' ', 'a', 'd', 'i', 'p', 'i', 's', 'i', 'c', 'i', 'n', 'g', ' ', 'e', 'l', 'i', 't'},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, util.RLPBytes(test.input))
		})
	}
}

// TestRLPList tests the RLP list encoding function.
func TestRLPList(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]byte
		expected []byte
	}{
		{
			name:     "Nil",
			input:    nil,
			expected: []byte{0xc0},
		},
		{
			name:     "Empty",
			input:    [][]byte{},
			expected: []byte{0xc0},
		},
		{
			name: "Catdog",
			input: [][]byte{
				{0x83, 'c', 'a', 't'},
				{0x83, 'd', 'o', 'g'},
			},
			expected: []byte{0xc8, 0x83, 'c', 'a', 't', 0x83, 'd', 'o', 'g'},
		},
		{
			name: "Nested",
			input: [][]byte{
				util.RLPBytes([]byte{0xfe}),
				util.RLPList([][]byte{
					util.RLPBytes([]byte{0x01}),
					util.RLPBytes([]byte{0x02}),
				}),
				util.RLPBytes([]byte{0xff}),
				util.RLPList([][]byte{
					util.RLPBytes([]byte{0x03}),
					util.RLPBytes([]byte{0x04}),
				}),
			},
			expected: []byte{0xca, 0x81, 0xfe, 0xc2, 0x01, 0x02, 0x81, 0xff, 0xc2, 0x03, 0x04},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, util.RLPList(test.input))
		})
	}
}

// TestRLPUint64 tests the RLP uint64 encoding function.
func TestRLPUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected []byte
	}{
		{
			name:     "0",
			input:    0,
			expected: []byte{0x80},
		},
		{
			name:     "1",
			input:    1,
			expected: []byte{0x01},
		},
		{
			name:     "127",
			input:    127,
			expected: []byte{0x7f},
		},
		{
			name:     "128",
			input:    128,
			expected: []byte{0x81, 0x80},
		},
		{
			name:     "182000000000",
			input:    182000000000,
			expected: []byte{0x85, 0x2a, 0x60, 0xb, 0x9c, 0x0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, util.RLPUint64(test.input))
		})
	}
}
