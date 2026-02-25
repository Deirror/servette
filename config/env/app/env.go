package app

import (
	"github.com/Deirror/servette/app"
	envcfg "github.com/Deirror/servette/config/env"
	config "github.com/Deirror/servette/config"
	"github.com/Deirror/servette/env"
)

type MultiConfig = config.MultiConfig[appx.Config]

var suffixes = []string{"APP_MODE", "APP_DOMAIN"}

// LoadConfig loads Config from environment variables.
// Required vars: APP_MODE, APP_DOMAIN
func LoadConfig(prefix ...string) (*appx.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	mode, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	domain, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	return appx.NewConfig(mode, domain), nil
}

// LoadMultiConfig scans env vars and builds app configs based on their prefix.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
