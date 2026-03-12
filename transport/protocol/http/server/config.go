// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package server

import (
	"crypto/tls"
	"time"

	"github.com/Deirror/servette/transport"
)

// Config contains configuration settings for an HTTP server.
type Config struct {
	// Transport type: TCP or Unix
	TransType transport.Type

	// Endpoint for TCP (host:port) or base URL for Unix ("http://unix"), but can be used also for socket filepaths
	Endpoint string

	// Timeouts
	ReadTimeout       time.Duration // max time to read request headers + body
	WriteTimeout      time.Duration // max time to write response
	IdleTimeout       time.Duration // max time for idle connections
	ReadHeaderTimeout time.Duration // time to read request headers

	// TLS
	TLSConfig *tls.Config

	MaxHeaderBytes int
}

func DefaultConfigx() *Config {
	return &Config{
		TransType:         transport.TCP,
		Endpoint:          "",
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       90 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		TLSConfig:         nil,
		MaxHeaderBytes:    1 << 20, // 1 MB, reasonable default
	}
}

func NewConfigWithTimeouts(endpoint string, readTimeout, writeTimeout, readHeaderTimeout, idleTimeout time.Duration) *Config {
	c := DefaultConfigx()
	if endpoint != "" {
		c.Endpoint = endpoint
	}
	if readTimeout != 0 {
		c.ReadTimeout = readTimeout
	}
	if writeTimeout != 0 {
		c.WriteTimeout = writeTimeout
	}
	if readHeaderTimeout != 0 {
		c.ReadHeaderTimeout = readHeaderTimeout
	}
	if idleTimeout != 0 {
		c.IdleTimeout = idleTimeout
	}
	return c
}

func (c *Config) WithTransType(t transport.Type) *Config {
	c.TransType = t
	return c
}

func (c *Config) withendpoint(endpoint string) *Config {
	c.Endpoint = endpoint
	return c
}

func (c *Config) WithReadTimeout(d time.Duration) *Config {
	c.ReadTimeout = d
	return c
}

func (c *Config) WithWriteTimeout(d time.Duration) *Config {
	c.WriteTimeout = d
	return c
}

func (c *Config) WithIdleTimeout(d time.Duration) *Config {
	c.IdleTimeout = d
	return c
}

func (c *Config) WithReadHeaderTimeout(d time.Duration) *Config {
	c.ReadHeaderTimeout = d
	return c
}

func (c *Config) WithTLSConfig(tlsCfg *tls.Config) *Config {
	c.TLSConfig = tlsCfg
	return c
}

func (c *Config) WithMaxHeaderBytes(n int) *Config {
	c.MaxHeaderBytes = n
	return c
}
