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

package jsonrpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
	"github.com/attestantio/go-execution-client/util"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// NewPendingTransactions returns a subscription for pending transactions.
func (s *Service) NewPendingTransactions(ctx context.Context, ch chan *spec.Transaction) (*util.Subscription, error) {
	// This is closed in closeSocketOnCtxDone(), so...
	//nolint:bodyclose
	conn, _, err := websocket.DefaultDialer.Dial(s.webSocketAddress, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to server")
	}

	// Set up the subscription.
	if err := conn.WriteMessage(websocket.TextMessage,
		[]byte(`{"jsonrpc":"2.0", "id": 1, "method": "eth_subscribe", "params": ["newPendingTransactions"]}`),
	); err != nil {
		return nil, errors.Wrap(err, "failed to request subscription")
	}

	// Read the response to obtain the subscription ID.
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return nil, errors.Wrap(err, "failed to obtain subscription response")
	}
	log.Trace().Str("msg", string(msg)).Msg("Received subscription response")
	res := newPendingTransactionsResult{}
	if err := json.Unmarshal(msg, &res); err != nil {
		return nil, errors.Wrap(err, "failed to obtain subscription ID")
	}
	log.Trace().Str("subscription", fmt.Sprintf("%#x", res.Result)).Msg("Received subscription ID")

	// Handle incoming messages.
	go s.receiveNewPendingTransactionMsg(ctx, conn, ch)

	// Close the websocket when the context is done.
	go s.closeSocketOnCtxDone(ctx, conn)

	return &util.Subscription{
		ID: res.Result,
	}, nil
}

func (s *Service) receiveNewPendingTransactionMsg(ctx context.Context, conn *websocket.Conn, ch chan *spec.Transaction) {
	for {
		_, msg, err := conn.ReadMessage()
		if ctx.Err() != nil {
			// Context is done; leave.
			return
		}
		if err != nil {
			log.Error().Err(err).Msg("Failed to read received message")

			continue
		}
		log.Info().Str("msg", string(msg)).Msg("Received message")
		res := newPendingTransactionsEvent{}
		if err := json.Unmarshal(msg, &res); err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal message")

			continue
		}

		tx, err := s.Transaction(ctx, res.Params.Result)
		if err != nil {
			log.Error().Err(err).Str("tx_hash", fmt.Sprintf("%#x", res.Params.Result)).Msg("Failed to obtain transaction")

			continue
		}
		ch <- tx
	}
}

func (*Service) closeSocketOnCtxDone(ctx context.Context, conn *websocket.Conn) {
	<-ctx.Done()
	log.Trace().Msg("Context done; closing websocket connection")

	// Close the connection.
	err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Error().Err(err).Msg("Failed to send websocket close message")

		return
	}
	if err := conn.Close(); err != nil {
		log.Error().Err(err).Msg("Failed to close websocket")

		return
	}

	log.Trace().Msg("Websocket connection closed")
}

type newPendingTransactionsResult struct {
	Result []byte
}
type newPendingTransactionsResultJSON struct {
	Result string `json:"result"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (r *newPendingTransactionsResult) UnmarshalJSON(input []byte) error {
	var data newPendingTransactionsResultJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	var err error
	r.Result, err = util.StrToByteArray("result", data.Result)
	if err != nil {
		return err
	}

	return nil
}

type newPendingTransactionsEvent struct {
	Params *newPendingTransactionsEventParams `json:"params"`
}

type newPendingTransactionsEventParams struct {
	Subscription []byte
	Result       types.Hash
}

type newPendingTransactionsResultParamsJSON struct {
	Subscription string `json:"subscription"`
	Result       string `json:"result"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *newPendingTransactionsEventParams) UnmarshalJSON(input []byte) error {
	var data newPendingTransactionsResultParamsJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}

	var err error
	e.Subscription, err = util.StrToByteArray("subscription", data.Subscription)
	if err != nil {
		return err
	}
	e.Result, err = util.StrToHash("result", data.Result)
	if err != nil {
		return err
	}

	return nil
}
