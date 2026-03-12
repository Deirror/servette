// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package app

import (
	appx "github.com/Deirror/servette/app"
	config "github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
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
