// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package app

// Config holds basic environment configuration like mode and domain.
type Config struct {
	Mode   string // Application mode: development, staging, production
	Domain string // Public-facing domain, e.g., example.com
}

func NewConfig(mode, domain string) *Config {
	return &Config{
		Mode:   mode,
		Domain: domain,
	}
}

// WithMode sets the mode and returns the updated Config.
func (c *Config) WithMode(mode string) *Config {
	c.Mode = mode
	return c
}

// WithDomain sets the domain and returns the updated Config.
func (c *Config) WithDomain(domain string) *Config {
	c.Domain = domain
	return c
}
