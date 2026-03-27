// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package urlx

// Config holds configuration for an URL.
type Config struct {
	URL string // any service URL
}

func NewConfig(url string) *Config {
	return &Config{
		URL: url,
	}
}

// WithURL sets the URL field and returns the updated config.
func (c *Config) WithURL(url string) *Config {
	c.URL = url
	return c
}
