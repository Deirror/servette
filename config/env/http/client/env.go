// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package client

import (
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/transport"
	"github.com/Deirror/servette/transport/protocol/http/client"
)

type MultiConfig = config.MultiConfig[client.Config]

var suffixes = []string{
	"HTTP_CLIENT_TRANSPORT_TYPE",
	"HTTP_CLIENT_ENDPOINT",
	"HTTP_CLIENT_DIAL_TIMEOUT",
	"HTTP_CLIENT_REQUEST_TIMEOUT",
	"HTTP_CLIENT_RESPONSE_HEADER_TIMEOUT",
	"HTTP_CLIENT_IDLE_CONN_TIMEOUT",
	"HTTP_CLIENT_MAX_IDLE_CONNS",
	"HTTP_CLIENT_MAX_IDLE_CONNS_PER_HOST",
	"HTTP_CLIENT_MAX_CONNS_PER_HOST",
	"HTTP_CLIENT_MAX_REDIRECTS",
}

// LoadConfig loads client config values from environment variables.
// The env var keys are prefixed with the optional prefix argument.
func LoadConfig(prefix ...string) (*client.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	transTypeEnv, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}
	transType, err := transport.ParseType(transTypeEnv)
	if err != nil {
		return nil, err
	}

	endpoint, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	dialTimeout, err := env.ParseTimeDuration(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	reqTimeout, err := env.ParseTimeDuration(pfx + suffixes[3])
	if err != nil {
		return nil, err
	}

	respHeadTimeout, err := env.ParseTimeDuration(pfx + suffixes[4])
	if err != nil {
		return nil, err
	}

	idleConnTimeout, err := env.ParseTimeDuration(pfx + suffixes[5])
	if err != nil {
		return nil, err
	}

	maxIdleConns, err := env.ParseInt(pfx + suffixes[6])
	if err != nil {
		return nil, err
	}

	maxIdleConnsPerHost, err := env.ParseInt(pfx + suffixes[7])
	if err != nil {
		return nil, err
	}

	maxConnsPerHost, err := env.ParseInt(pfx + suffixes[8])
	if err != nil {
		return nil, err
	}

	maxRedirects, err := env.ParseInt(pfx + suffixes[9])
	if err != nil {
		return nil, err
	}

	// skip tls, since the config is loaded as shared
	return &client.Config{
		TransType:             transType,
		Endpoint:              endpoint,
		DialTimeout:           dialTimeout,
		RequestTimeout:        reqTimeout,
		ResponseHeaderTimeout: respHeadTimeout,
		IdleConnTimeout:       idleConnTimeout,
		MaxIdleConns:          maxIdleConns,
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		MaxConnsPerHost:       maxConnsPerHost,
		MaxRedirects:          maxRedirects,
	}, nil
}

// LoadMultiConfig scans env vars and builds Client configs based on their prefix.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
