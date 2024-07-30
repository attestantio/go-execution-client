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

package jsonrpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	zerologger "github.com/rs/zerolog/log"
	"github.com/ybbus/jsonrpc/v2"
)

// Service is an Ethereum execution client service.
type Service struct {
	address          string
	webSocketAddress string
	client           jsonrpc.RPCClient
	timeout          time.Duration

	// Client capability information.
	isIssuanceProvider bool
}

// log is a service-wide logger.
var log zerolog.Logger

// New creates a new execution client service, connecting with a standard HTTP.
func New(ctx context.Context, params ...Parameter) (execclient.Service, error) {
	parameters, err := parseAndCheckParameters(params...)
	if err != nil {
		return nil, errors.Wrap(err, "problem with parameters")
	}

	// Set logging.
	log = zerologger.With().Str("service", "client").Str("impl", "jsonrpc").Logger()
	if parameters.logLevel != log.GetLevel() {
		log = log.Level(parameters.logLevel)
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:        64,
			MaxIdleConnsPerHost: 64,
			IdleConnTimeout:     384 * time.Second,
		},
	}

	address := parameters.address
	if !strings.HasPrefix(address, "http") {
		address = fmt.Sprintf("http://%s", parameters.address)
	}

	webSocketAddress := parameters.webSocketAddress
	if strings.HasPrefix(webSocketAddress, "http://") {
		webSocketAddress = fmt.Sprintf("ws://%s", webSocketAddress[7:])
	}
	if strings.HasPrefix(webSocketAddress, "https://") {
		webSocketAddress = fmt.Sprintf("wss://%s", webSocketAddress[8:])
	}
	if !strings.HasPrefix(webSocketAddress, "ws") {
		webSocketAddress = fmt.Sprintf("ws://%s", webSocketAddress)
	}
	log.Trace().Str("address", address).Str("web_socket_address", webSocketAddress).Msg("Addresses configured")

	rpcClient := jsonrpc.NewClientWithOpts(address, &jsonrpc.RPCClientOpts{
		HTTPClient: client,
	})

	s := &Service{
		client:           rpcClient,
		address:          address,
		webSocketAddress: webSocketAddress,
		timeout:          parameters.timeout,
	}

	// Fetch static values to confirm the connection is good.
	if err := s.fetchStaticValues(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to confirm node connection")
	}

	// Handle flags for capabilities.
	if err := s.checkCapabilities(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to check capabilities")
	}

	// Close the service on context done.
	go func(s *Service) {
		<-ctx.Done()
		log.Trace().Msg("Context done; closing connection")
		s.close()
	}(s)

	return s, nil
}

// fetchStaticValues fetches values that never change.
// This caches the values, avoiding future API calls.
func (*Service) fetchStaticValues(_ context.Context) error {
	return nil
}

// checkCapabilites checks the capabilities of the client and sets
// internal flags appropriately.
//
//nolint:unparam
func (s *Service) checkCapabilities(ctx context.Context) error {
	_, err := s.issuanceAtHeight(ctx, 1)
	s.isIssuanceProvider = err == nil

	return nil
}

// Name provides the name of the service.
func (*Service) Name() string {
	return "json-rpc"
}

// Address provides the address for the connection.
func (s *Service) Address() string {
	return s.address
}

// close closes the service, freeing up resources.
func (*Service) close() {
}
