// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package kv

import (
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/domain/kv"
	"github.com/Deirror/servette/env"
)

type MultiConfig = config.MultiConfig[kv.Config]

var suffixes = []string{
	"KV_STORE_URL",
}

// LoadConfig loads the key-value store configuration from environment variables,
// supporting an optional prefix.
func LoadConfig(prefix ...string) (*kv.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	url, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	return kv.NewConfig(url), nil
}

// LoadMultiConfig loads multiple KV Config instances by scanning env vars with suffixes.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
