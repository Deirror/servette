// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package kv

// Config holds configuration details for connecting to a key-value store.
type Config struct {
	StoreURL string // URL of the key-value store
}

func NewConfig(storeURL string) *Config {
	return &Config{
		StoreURL: storeURL,
	}
}

// WithStoreURL sets a new key-value store URL and returns the updated KVConfig.
func (c *Config) WithKVStoreURL(url string) *Config {
	c.StoreURL = url
	return c
}
