package server

import (
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/transport/socket/unix/server"
)

type MultiConfig = envcfg.MultiConfig[server.Config]

var suffixes = []string{
	"UNIX_SOCKET_SERVER_PORT",
	"UNIX_SOCKET_SERVER_READ_TIMEOUT",
	"UNIX_SOCKET_SERVER_WRITE_TIMEOUT",
	"UNIX_SOCKET_SERVER_IDLE_TIMEOUT",
}

// LoadConfig loads server config values from environment variables.
// The env var keys are prefixed with the optional prefix argument.
func LoadConfig(prefix ...string) (*server.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	port, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	readTimeout, err := env.ParseTimeDuration(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	writeTimeout, err := env.ParseTimeDuration(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	idleTimeout, err := env.ParseTimeDuration(pfx + suffixes[3])
	if err != nil {
		return nil, err
	}

	return server.NewConfig(port, readTimeout, writeTimeout, idleTimeout), nil
}

// LoadMultiConfig scans env vars and builds Server configs based on their prefix.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
