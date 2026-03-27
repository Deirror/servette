// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package oauth

import (
	"github.com/Deirror/servette/auth/oauth"
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/env"
)

type MultiConfig = config.MultiConfig[oauth.Config]

var suffixes = []string{
	"OAUTH_CLIENT_ID",
	"OAUTH_CLIENT_SECRET",
	"OAUTH_REDIRECT_URL",
}

// LoadConfig loads OAuth Config from environment variables,
// with an optional prefix.
func LoadConfig(prefix ...string) (*oauth.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	clientID, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	clientSecret, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	redirectURL, err := env.Get(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	return oauth.NewConfig(clientID, clientSecret, redirectURL), nil
}

// LoadMultiConfig loads multiple OAuth Configs by scanning env vars with oauth suffixes.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
