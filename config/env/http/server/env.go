package server

import (
	"errors"

	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/protocol/http/server"
	"github.com/Deirror/servette/transport"
)

type MultiConfig = config.MultiConfig[server.Config]

var suffixes = []string{
	"HTTP_SERVER_TRANSPORT_TYPE",
	"HTTP_SERVER_ENDPOINT",
	"HTTP_SERVER_READ_TIMEOUT",
	"HTTP_SERVER_WRITE_TIMEOUT",
	"HTTP_SERVER_IDLE_TIMEOUT",
	"HTTP_SERVER_READ_HEADER_TIMEOUT",
	"HTTP_SERVER_MAX_HEADER_BYTES",
}

// LoadConfig loads server config values from environment variables.
// The env var keys are prefixed with the optional prefix argument.
func LoadConfig(prefix ...string) (*server.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	transType, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}
	if !transport.IsValidType(transType) {
		return nil, errors.New("transport type is unknown")
	}

	endpoint, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	readTimeout, err := env.ParseTimeDuration(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	writeTimeout, err := env.ParseTimeDuration(pfx + suffixes[3])
	if err != nil {
		return nil, err
	}

	idleTimeout, err := env.ParseTimeDuration(pfx + suffixes[4])
	if err != nil {
		return nil, err
	}

	readHeaderTimeout, err := env.ParseTimeDuration(pfx + suffixes[5])
	if err != nil {
		return nil, err
	}

	maxHeaderBytes, err := env.ParseInt(pfx + suffixes[6])
	if err != nil {
		return nil, err
	}

	return &server.Config{
		TransType:         transType,
		Endpoint:          endpoint,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}, nil
}

// LoadMultiConfig scans env vars and builds Server configs based on their prefix.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
