// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package app

import (
	"github.com/Deirror/servette/env"
)

// Config holds basic environment configuration like mode and domain.
type Config struct {
	Mode   env.Mode // Application mode: development, staging, production
	Domain string   // Public-facing domain, e.g., example.com
}

func NewConfig(m env.Mode, d string) *Config {
	return &Config{
		Mode:   m,
		Domain: d,
	}
}

// WithMode sets the mode and returns the updated Config.
func (c *Config) WithMode(m env.Mode) *Config {
	c.Mode = m
	return c
}

// WithDomain sets the domain and returns the updated Config.
func (c *Config) WithDomain(d string) *Config {
	c.Domain = d
	return c
}
