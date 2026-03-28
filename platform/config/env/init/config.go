// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package initx

import (
	"github.com/Deirror/servette/config/env/config"
)

// Config is a special config which has the app paths for
// the actual configs to be loaded.
type Config struct {
	Cfgs config.MultiConfig
}

func NewConfig(cs config.MultiConfig) *Config {
	return &Config{
		Cfgs: cs,
	}
}

// LoadConfig uses OS env vars and gets app paths.
// It is used to identify how your app should load the actual configs.
func LoadConfig() (*Config, error) {
	cfgs, err := config.LoadMultiConfig()
	if err != nil {
		return nil, err
	}

	return NewConfig(cfgs), nil
}
