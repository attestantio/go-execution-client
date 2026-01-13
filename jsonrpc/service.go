// Copyright © 2021, 2025 Attestant Limited.
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
	"errors"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/rs/zerolog"
	zerologger "github.com/rs/zerolog/log"
	"github.com/ybbus/jsonrpc/v2"
)

// Service is an Ethereum execution client service.
type Service struct {
	base             *url.URL
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
		return nil, err
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

	addrResult, err := parseAddress(parameters.address)
	if err != nil {
		return nil, err
	}

	base := addrResult.Base
	address := addrResult.Masked

	webSocketAddress := parameters.webSocketAddress

	// Protocol strings should be constants
	const (
		httpScheme  = "http://"
		httpsScheme = "https://"
		wsScheme    = "ws://"
		wssScheme   = "wss://"
	)

	if strings.HasPrefix(webSocketAddress, httpScheme) {
		webSocketAddress = wsScheme + webSocketAddress[len(httpScheme):]
	}

	if strings.HasPrefix(webSocketAddress, httpsScheme) {
		webSocketAddress = wssScheme + webSocketAddress[len(httpsScheme):]
	}

	if !strings.HasPrefix(webSocketAddress, "ws") {
		webSocketAddress = wsScheme + webSocketAddress
	}

	if webSocketAddress == "" {
		webSocketAddress = base.String()
	}

	log.Trace().Stringer("address", address).Str("web_socket_address", webSocketAddress).Msg("Addresses configured")

	rpcClient := jsonrpc.NewClientWithOpts(base.String(), &jsonrpc.RPCClientOpts{
		HTTPClient: client,
	})

	s := &Service{
		client:           rpcClient,
		base:             base,
		address:          address.String(),
		webSocketAddress: webSocketAddress,
		timeout:          parameters.timeout,
	}

	// Fetch static values to confirm the connection is good.
	if err := s.fetchStaticValues(ctx); err != nil {
		return nil, errors.Join(errors.New("failed to confirm node connection"), err)
	}

	// Handle flags for capabilities.
	if err := s.checkCapabilities(ctx); err != nil {
		return nil, errors.Join(errors.New("failed to check capabilities"), err)
	}

	// Close the service on context done.
	go func(s *Service) {
		<-ctx.Done()
		log.Trace().Msg("Context done; closing connection")
		s.close()
	}(s)

	return s, nil
}

// Name provides the name of the service.
func (*Service) Name() string {
	return "json-rpc"
}

// Address provides the address for the connection.
func (s *Service) Address() string {
	return s.address
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

// close closes the service, freeing up resources.
func (*Service) close() {
}

// AddressResult contains both the parsed base URL and a masked version for logging.
// Used as a clear named result to avoid unnamed returns of the same type (confusing-results warning).
type AddressResult struct {
	Base   *url.URL // Parsed base URL for the connection
	Masked *url.URL // Sanitized version of the base URL with masked passwords, paths, and query values for logging
}

func parseAddress(address string) (AddressResult, error) {
	// Protocol strings should be constants
	const httpPrefix = "http://"
	if !strings.HasPrefix(address, "http") {
		address = httpPrefix + address
	}

	base, err := url.Parse(address)
	if err != nil {
		return AddressResult{}, errors.Join(errors.New("invalid URL"), err)
	}
	// Remove any trailing slash from the path.
	base.Path = strings.TrimSuffix(base.Path, "/")

	// Attempt to mask any sensitive information in the URL, for logging purposes.
	baseAddress := *base
	if _, pwExists := baseAddress.User.Password(); pwExists {
		// Mask the password.
		user := baseAddress.User.Username()
		baseAddress.User = url.UserPassword(user, "xxxxx")
	}

	if baseAddress.Path != "" {
		// Mask the path.
		baseAddress.Path = "xxxxx"
	}

	if baseAddress.RawQuery != "" {
		// Mask all query values.
		sensitiveRegex := regexp.MustCompile("=([^&]*)(&)?")
		baseAddress.RawQuery = sensitiveRegex.ReplaceAllString(baseAddress.RawQuery, "=xxxxx$2")
	}

	return AddressResult{Base: base, Masked: &baseAddress}, nil
}
