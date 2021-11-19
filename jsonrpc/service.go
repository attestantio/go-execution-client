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
	// Hold the initialising context to use for streams.
	ctx context.Context

	address string
	client  jsonrpc.RPCClient
	timeout time.Duration
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

	rpcClient := jsonrpc.NewClientWithOpts(address, &jsonrpc.RPCClientOpts{
		HTTPClient: client,
	})

	s := &Service{
		ctx:     ctx,
		client:  rpcClient,
		address: parameters.address,
		timeout: parameters.timeout,
	}

	// Fetch static values to confirm the connection is good.
	if err := s.fetchStaticValues(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to confirm node connection")
	}

	//	// Handle flags for API versioning.
	//	if err := s.checkAPIVersioning(ctx); err != nil {
	//		return nil, errors.Wrap(err, "failed to check API versioning")
	//	}

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
func (s *Service) fetchStaticValues(ctx context.Context) error {
	if _, err := s.NetworkID(ctx); err != nil {
		return errors.Wrap(err, "failed to fetch network ID")
	}

	return nil
}

// // checkAPIVersioning checks the versions of some APIs and sets
// // internal flags appropriately.
// func (s *Service) checkAPIVersioning(ctx context.Context) error {
// 	// Start by setting the API v2 flag for blocks and fetching block 0.
// 	s.supportsV2BeaconBlocks = true
// 	_, err := s.SignedBeaconBlock(ctx, "0")
// 	if err == nil {
// 		// It's good.  Assume that other V2 APIs introduced with Altair
// 		// are present.
// 		s.supportsV2BeaconState = true
// 		s.supportsV2ValidatorBlocks = true
// 	} else {
// 		// Assume this is down to the V2 endpoint missing rather than
// 		// some other failure.
// 		s.supportsV2BeaconBlocks = false
// 	}
// 	return nil
// }

// Name provides the name of the service.
func (s *Service) Name() string {
	return "json-rpc"
}

// Address provides the address for the connection.
func (s *Service) Address() string {
	return s.address
}

// close closes the service, freeing up resources.
func (s *Service) close() {
}
