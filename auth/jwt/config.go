// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package jwt

import "time"

// Config holds the configuration related to JWT-based authentication.
type Config struct {
	CookieName string        // Name of the cookie that stores the JWT
	Secret     string        // Secret key used to sign JWTs
	TokenTTL   time.Duration // Time-to-live duration of the token
}

func NewConfig(name, secret string, ttl time.Duration) *Config {
	return &Config{
		CookieName: name,
		Secret:     secret,
		TokenTTL:   ttl,
	}
}

// WithCookieName sets the cookie name for the Config.
func (c *Config) WithCookieName(name string) *Config {
	c.CookieName = name
	return c
}

// WithTokenTTL sets the token time-to-live for the Config.
func (c *Config) WithTokenTTL(ttl time.Duration) *Config {
	c.TokenTTL = ttl
	return c
}

// WithSecret sets the secret key for signing JWTs in the Config.
func (c *Config) WithSecret(secret string) *Config {
	c.Secret = secret
	return c
}
