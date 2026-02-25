package client

import (
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/transport/socket/unix/client"
)

type MultiConfig = config.MultiConfig[client.Config]

var suffixes = []string{
	"UNIX_SOCKET_CLIENT_PORT",
	"UNIX_SCOKET_CLIENT_WRITE_TIMEOUT",
	"UNIX_SOCKET_CLIENT_IDLE_TIMEOUT",
}

// LoadConfig loads client config values from environment variables.
// The env var keys are prefixed with the optional prefix argument.
func LoadConfig(prefix ...string) (*client.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	path, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	writeTimeout, err := env.ParseTimeDuration(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	idleTimeout, err := env.ParseTimeDuration(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	return client.NewConfig(path, writeTimeout, idleTimeout), nil
}

// LoadMultiConfig scans env vars and builds Client configs based on their prefix.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
