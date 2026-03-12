// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package urlx

import (
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/env"
	urlx "github.com/Deirror/servette/transport/url"
)

type MultiConfig = config.MultiConfig[urlx.Config]

// suffixes defines environment variable suffixes for URL Config.
// EXTERNAL prefix is used to identify properly the envs, since URL is too generic.
var suffixes = []string{
	"EXTERNAL_URL",
}

// LoadConfig loads URL Config from environment variables.
// Optionally accepts a prefix to prepend to suffixes.
func LoadConfig(prefix ...string) (*urlx.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	url, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	return urlx.NewConfig(url), nil
}

// LoadMultiConfig loads multiple URL Config from environment variables
// by scanning for all environment variable sets matching externalURLSuffixes.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
