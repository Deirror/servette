// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package config

import (
	"github.com/Deirror/servette/path"
)

// Config represents the configuration for the application.
// It combines a ReadMode, which determines how the configuration
// is loaded or interpreted, with a Resource, which points to
// the underlying source of the configuration (e.g., a local path
// or a remote URI).
type Config struct {
	ReadMode                   // Determines the mode in which configuration is read
	Resources []pathx.Resource // The resources representing the config source
}

func NewConfig(m ReadMode, rs []pathx.Resource) *Config {
	return &Config{
		ReadMode:  m,
		Resources: rs,
	}
}

// WithReadMode returns a copy of the Config with a new ReadMode.
func (c *Config) WithReadMode(m ReadMode) *Config {
	c.ReadMode = m
	return c
}

// WithResource returns a copy of the Config with a new Resource.
func (c *Config) WithResource(rs []pathx.Resource) *Config {
	c.Resources = rs
	return c
}
