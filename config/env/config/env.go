// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package config

import (
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/internal/utils/str"
	"github.com/Deirror/servette/path"
)

type MultiConfig = config.MultiConfig[config.Config]

var suffixes = []string{"CONFIG_READ_MODE", "RESOURCES"}

// LoadConfig loads Config from environment variables.
// Required vars: CONFIG_READ_MODE, RESOURCES
func LoadConfig(prefix ...string) (*config.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	readModeEnv, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	readMode, err := config.ParseReadMode(readModeEnv)
	if err != nil {
		return nil, err
	}

	readResourcesEnv, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	readResources := str.SplitAndTrim(readResourcesEnv)

	return config.NewConfig(readMode, StringsToResources(readResources)), nil
}

// LoadMultiConfig scans env vars and builds app configs based on their prefix.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}

// StringsToResources converts a slice of strings into []Resource
func StringsToResources(strs []string) []pathx.Resource {
	res := make([]pathx.Resource, len(strs))
	for i, s := range strs {
		res[i] = pathx.Resource(s)
	}
	return res
}
