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
	"context"
	"encoding/json"
	"os"
	"testing"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/attestantio/go-execution-client/jsonrpc"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestReplayBlockTransactions(t *testing.T) {
	ctx := context.Background()
	s, err := jsonrpc.New(ctx,
		jsonrpc.WithLogLevel(zerolog.Disabled),
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	tests := []struct {
		name     string
		blockID  string
		err      string
		expected []string
	}{
		{
			name:     "NoTransactions",
			blockID:  "12345",
			expected: []string{`[]`},
		},
		{
			name:    "FirstTransaction",
			blockID: "46147",
			expected: []string{
				`[{"stateDiff":{"0x5df9b87991262f6ba471f09758cde1c0fc1de734":{"balance":{"+":"0x7a69"},"nonce":{"+":"0x0"}},"0xa1e4380a3b1f749673e270229993ee55f35663b4":{"balance":{"*":{"from":"0x6c6b935b8bbd400000","to":"0x6c5d01021be7168597"}},"nonce":{"*":{"from":"0x0","to":"0x1"}}},"0xe6a7a1d47ff21b6321162aea7c6cb457d5476bca":{"balance":{"*":{"from":"0xf3426785a8ab466000","to":"0xf350f9df18816f6000"}}}},"transactionHash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060"}]`,
				`[{"stateDiff":{"0x5df9b87991262f6ba471f09758cde1c0fc1de734":{"balance":{"*":{"from":"0x0","to":"0x7a69"}}},"0xa1e4380a3b1f749673e270229993ee55f35663b4":{"balance":{"*":{"from":"0x6c6b935b8bbd400000","to":"0x6c5d01021be7168597"}},"nonce":{"*":{"from":"0x0","to":"0x1"}}},"0xe6a7a1d47ff21b6321162aea7c6cb457d5476bca":{"balance":{"*":{"from":"0xf3426785a8ab466000","to":"0xf350f9df18816f6000"}}}},"transactionHash":"0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060"}]`,
			},
		},
		{
			name:     "StorageChange",
			blockID:  "11977076",
			expected: []string{`[{"stateDiff":{"0x0000000000004946c0e9f43f4dee607b0ef1fa1c":{"nonce":{"*":{"from":"0x6840d9","to":"0x684165"}},"storage":{"0x0000000000000000000000000000000000000000000000000000000000000002":{"*":{"from":"0x00000000000000000000000000000000000000000000000000000000006840d8","to":"0x0000000000000000000000000000000000000000000000000000000000684164"}},"0x2c7a74ce447c4ff7d241edaea041145a45daf4ba74ebab56a20d9041e365357f":{"*":{"from":"0x0000000000000000000000000000000000000000000000000000000000000119","to":"0x00000000000000000000000000000000000000000000000000000000000001a5"}}}},"0x0090705a9b9fbb680d9924c1394e702e073656a3":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x01242ab0f8e203c45beb2033347156dd5a5fb3ac":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x01deeceea40d4ba0a266dda79226954bd3b9a27d":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x02ec8474402a65d79538300b50d5182e5370ced9":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x0360cbf679536a607502768b8f1c80df848de220":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x0426e2e6507a3181d1c5be5ee04975bd386ecca8":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x05d9fbe029da4318d0c08c0196a1b9f505051227":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x05e014cb12845966501ab0d5228bd96184b7beec":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x06568e70e32ba2651955b472a8587857f6fdfdcf":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x084fb7caeca192da67c5e6dc7d1a8a286bd52b58":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x0879a9d31d396c0ff3e2bb38bbe093eef7c52838":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x0acfb865fcdb154a333391681457089b0b5419cc":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x0da1a682efbe9e3db02046c02ff1cfa86c90600c":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x12ec7af771de6dbd02097987672a9ca6dcace31f":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x177ab4eb51fc3c5169f4c7003f3d78d0218f16e9":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x1929c16cce87133b03eaf0fa8553f914ee85df62":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x1981c3ef2e12bb5f8a3d0a94d7cab1d59c8a1d2b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x1bef220f9f4ab3c34f0835fd1821c932ff215aeb":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x1ccf946a15ec54ab0cd3a5276caa5b0b859799a1":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x1f3e63110e5316e3b21015797ca4d51616d92f8f":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x246c35dfb612fabaee8e319ab13224c714bfed85":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x2561f5009e138d4d04857ec538a89b5cf50ce6b6":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x260005a315882c80ddde8e8ce7c84cc768383ba3":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x2c393b8835e36cd413bc62201e29fbfb8ef63f8d":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x2cfef9f13324700b8f35deea2be178565c216c76":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x2d761e1567737d84009ac06e076e7355ad4a905e":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x2e0e0224b663713a90e3d211ac0ccef362618266":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x2feff460997bf6b6314c83aa1411eff4d957c9fa":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x301ae1770958930d582c7e44dc81f8438ec46675":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x33b6245f4c5721b22fc6fa3b4c5d5f59fb2fe034":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x33f1bcb19b1bb1be1220e172cf2a64ab2fe4f66c":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x384ee7ad0b4127278f936eae3f4030adf7cf7c86":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x39f68a02e7638170da1438dfdb4e68e31e6d0661":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x3b7ded563c72ea6fde3da03c3a7db8d15109857b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x3ef606e887e53edc880c80d12d2201719c6c1140":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x3f7ec059cbdac7f61e5a268354de94a6dcf78451":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x3faf37c5ec476eec44effd0ccefbacaddab298e6":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x3fb384128dc386e4ae9b6c6122ad309ea44facdf":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x402bf9461e6bce7432dec587d221e6e16c0b150e":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x426d6434f8a3cbf66eb6947b34916f4eb0dcea97":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x4346dc07da017ebcf1d0293eab73f33f4e0b050e":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x45d07a2527515dd4238ec009729ceef7ec96b75c":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x479a74ed455a962c6d64294206f445eff3969be2":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x4e340c94f732cc8550b46215fc3b2b70b3770a65":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x512de3819c0497663f664553f1bf99057af8ac63":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x5464fc0019a242f4d7efab521bf764b15a97e4ca":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x5543d97b60a19e085b9e909333a146359af647ef":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x55939a2f26a00a426537bed29deb6ee2ff5ef251":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x56aa02b8ed6883c2902fdbaa74496c6bcf7a6d6b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x56e25af93c61e6d38d58330b453ae660433f8c99":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x575f6c490d24dc90321b7ed161d7526ae41c401f":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x5acc71a6155b63ae9f4370d31ddf72cb41f0dd39":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x5d17977190a1e8a9ade0e3cd4c19e7d9b90e12ee":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x6144a9571b505aa1a8900f1cc46e6cf69d6ca80c":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x63273eeb53b2a4d63ed5b26544a0c84de88f1336":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x6b781ca09adc95abd274337c4a992e72cca60daa":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x6ce594b0fcb41191ebee176881875bd06f3a274a":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x6fc26574a6f76bd098a347e2f35a56695e909910":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x6fd243947e3acd870067ec105b61601f94d1664f":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x71984ac608403ad67e9c29ae073d6dacb73a70eb":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x71fcbee77a380bde04b75637ccf28bc5fd709dbf":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x731eb7fbde475475d432f4b62d544e433f98ce43":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x73c3e56d48dcdee4da3397492c84728ac64e67c2":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x73fb6c1eada15b1c9b694a68f6627832e455d7a9":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x758ae98e346181e776f8dd77c0929892c058be1e":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x7820d9ee519d7d749ee1f4dfb46bbc654cc20c4b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x78fd71ed937a04cc4171fc37314ae0dedf4feafe":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x794879907134515b83e9c75d0c8a3a6a687923dd":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x7abb19bddea0ee8ddd5fcaa5012c6d5504192838":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x7de5e23705e9bfb5d8aef3436a65ade8913810ac":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x7dee31c4e9f01a98b758ab5b2a630b99f32ac7ec":{"nonce":{"*":{"from":"0x4","to":"0x5"}}},"0x7ee6f03274d848ad9a42268cfef846e42c5151a0":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x7f5b8d48daff24484cc2585bc8ab53798781a496":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x7fb35f1ff048bdf929a7ac05966ad8103c4b18c6":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x81a58e0db9d7cffa0826778b22c3eca5c729aa4b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x81a7b46789475fd429ca8ecc72341ac45e0991dd":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x83ba670f475ee271e1f9985438be1745167804cf":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x83c995ad9cec63a2e45a93a205a578cb897a420a":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x85aa08078bab2c1b120c6dfe9a2389ba94aa508e":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x8874f0f6d7c9d38df6862251465f32d028db545b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x88c4d4bb9fc1af32184d5a023196b713aa10acf7":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x8ab53f524347bea18a512e352891aaef2ed059d5":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x8b09e3c8f6fcf462f33986c232b36c4b2f1aa1e7":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x8befa1535b489ab731d4b156bb57ac51751ecb76":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x8e483c32b2028ea7bfbfa0223e5d361e74fda2b0":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x90ca629f2af60b0057b10afb8eca25293bfdf70a":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x91ccd4a171ff578adbeaf07b950bddedac128565":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x92dcbb9cf450ef5a7ce8fd3e4598c63bb663be61":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x93a76b7a94686657777b4a5e3586ccfe2ab76fed":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x94f9b074ad84737e32b9ce09738bb7d1a9be6454":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x97454ba76184b100573f5bf9d801cb3843623eb5":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0x9eff39ae758c7e80a4c0cda1a63a6f16fa6460e3":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xa183c96e30a47d738cf2d669329408e59956e0d9":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xa3604afa7193e36985736c8fdb0366c5c05a8932":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xa66c7efd4e8c0d78eeddf65d5c83e1d71cd89d1a":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xa71d23fb7a94412bf4f88efdd41f1d01a3475ad5":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xa82f88274474c21a22362de1f616caa9c2ac2702":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xa9d2757b2cf4168a1a3002b0eccb8d08507e0f93":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xacd8463fb707cbd1f2ed77cecdfe3ae18962f825":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xae9992b42a93789aa9bfa0d31b9d5ebe266ccfc0":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xb0eaa95e1ec98da997d84dca607f2dfc4b97f1f7":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xb298ace89f0e11d2a7810f6cca7288be1a2b199b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xb35e4d17d74e3e444f349093c896ff465483b9ed":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xb6d5fc8401aecf13fdfc4dd7afe2a5a96ae96073":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xbfb5f1e92edd9547bae098070c772b53c64b711c":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xc07e95a87657e058afac02011eb342d4adcd42a6":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xc16e0cfdfdf759f6fd574915c734df4b6cf00ac4":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xc3ce85415a12318ff23160d87bc3ae2cd5d4963b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xc4644474304d6f101af5e182833089c08dd105b5":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xc99e6f1e34851a1bf69a489202d27665bbe3a84e":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xcb191e2613bbfe8356df2539b7d161f00639dff1":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xcd4d8dd23b76f7f4d9fd8c991a10d593b4c1de99":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xcf896a8f2e29c801ebd0a79e4f83b137d09866be":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd0a9a9a8779fd9ab8ccf5ba1c150b9554e7968ab":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd14c78ac11de46bac51d64db0eef5adbb1dd003a":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd1a3dba63a1f7531941e51ebf7563a3abe218ca3":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd3af947389afe482b86d17243e7f0523a805be1f":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd52b7bda6db3c16eaead7c47ed757c401dbe33cc":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd53b27fe4648a42c47eb9e84097196d68f20cbd8":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd714c23ad29ce8f745d943167a4ae9eb731c4dec":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd8dc0c02e94d0f183fcb407999274ae4526297e5":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xd8e265a02ae7bd061e748a8a5e13786fb067f94f":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xdc14974e6aa6590a0447d6a69dde13f5d6ec03ce":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xdce284ddba5242a6c35b1638474fbc851c905dae":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xe093f46dc156ea7abb7d2314064a57f89bc080d9":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xe17376e6ed494ac7938e7b1ddb2caeb71f2f2b38":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xe4d9754732a3b0cfd89bcd980e013d86ac2835c0":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xe84f5435930312ac623f7133b15c535e40b457ba":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xea05f81aca8bf5762264a2afc4118819068be73d":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xedac2004a9f720bf864affde66a20ff825818536":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xefe2473cfb0b1924d516d15e57dc74c6d4b7700b":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xf1c59d32441db57147288e36afb8439a5f3cc3d0":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xf1ed30a307d007e4b120c9663fb57c30ab51496e":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xf206abcf6cb8af3eeedad1333dbd383aedbbe41e":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xf23adb1249bcb84a03ecb0ec268259266e95467f":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xf3bf43362923f982a7303cc2157fd26a365e118f":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xf5151c9847f6bd9feb67354727e96825348174b3":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xf64887c298b9cec5cbdc14ca37bf94d93b7fd972":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xfa5b232e1ace9b4ba26d4f5fb21bc81764c604fd":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xfa7072b713b7045cc270ecfe65a76f6d5d34112a":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}},"0xfd045524765fc3df77f5fb4170f0a251e13f8491":{"balance":{"+":"0x0"},"nonce":{"+":"0x1"}}},"transactionHash":"0xbaf59fe9b51d6611fca28bfe0371b51e30e71df8e66b173f0543ce48e88bd019"}]`},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			replay, err := s.(execclient.BlockReplaysProvider).ReplayBlockTransactions(ctx, test.blockID)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, replay)
				res, err := json.Marshal(replay)
				require.NoError(t, err)
				match := false
				for i := range test.expected {
					if test.expected[i] == string(res) {
						match = true
						break
					}
				}
				require.True(t, match, "output does not match any expected string")
				// require.Equal(t, test.expected, string(res))
			}
		})
	}
}
