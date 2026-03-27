// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package client

import (
	"crypto/tls"
	"time"

	"github.com/Deirror/servette/transport"
)

// Config contains configuration settings for an HTTP client.
type Config struct {
	// Transport type (tcp, unix, etc.)
	TransType transport.Type

	// Endpoint (scheme + host or "http://unix" for unix socket) or spcket path, or just "where to connect"
	Endpoint string

	// Timeouts
	DialTimeout           time.Duration // time to establish TCP/Unix connection
	RequestTimeout        time.Duration // overall request deadline (http.Client.Timeout)
	ResponseHeaderTimeout time.Duration // time to wait for response headers
	IdleConnTimeout       time.Duration // how long idle connections are kept

	// Pooling
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	MaxConnsPerHost     int

	// TLS
	TLSConfig *tls.Config

	// Redirect behavior
	MaxRedirects int
}

// DefaultCfg returns a production-friendly default configuration.
func DefaultConfigx() *Config {
	return &Config{
		TransType:             transport.TCP,
		Endpoint:              "",
		DialTimeout:           5 * time.Second,
		RequestTimeout:        30 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		IdleConnTimeout:       90 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		MaxConnsPerHost:       0, // unlimited
		TLSConfig:             nil,
		MaxRedirects:          10,
	}
}

func NewConfigWithTimeouts(endpoint string, dialTimeout, requestTimeout, responseHeaderTimeout, idleConnTimeout time.Duration) *Config {
	c := DefaultConfigx()
	if endpoint != "" {
		c.Endpoint = endpoint
	}
	if dialTimeout != 0 {
		c.DialTimeout = dialTimeout
	}
	if requestTimeout != 0 {
		c.RequestTimeout = requestTimeout
	}
	if responseHeaderTimeout != 0 {
		c.ResponseHeaderTimeout = responseHeaderTimeout
	}
	if idleConnTimeout != 0 {
		c.IdleConnTimeout = idleConnTimeout
	}
	return c
}

func (c *Config) WithTransport(t transport.Type) *Config {
	c.TransType = t
	return c
}

func (c *Config) WithEndpoint(e string) *Config {
	c.Endpoint = e
	return c
}

func (c *Config) WithDialTimeout(d time.Duration) *Config {
	c.DialTimeout = d
	return c
}

func (c *Config) WithRequestTimeout(d time.Duration) *Config {
	c.RequestTimeout = d
	return c
}

func (c *Config) WithResponseHeaderTimeout(d time.Duration) *Config {
	c.ResponseHeaderTimeout = d
	return c
}

func (c *Config) WithIdleConnTimeout(d time.Duration) *Config {
	c.IdleConnTimeout = d
	return c
}

func (c *Config) WithPooling(maxIdle, perHost, maxPerHost int) *Config {
	if maxIdle != 0 {
		c.MaxIdleConns = maxIdle
	}
	if perHost != 0 {
		c.MaxIdleConnsPerHost = perHost
	}
	if maxPerHost != 0 {
		c.MaxConnsPerHost = maxPerHost
	}
	return c
}

func (c *Config) WithTLSConfig(t *tls.Config) *Config {
	c.TLSConfig = t
	return c
}

func (c *Config) WithMaxRedirects(n int) *Config {
	c.MaxRedirects = n
	return c
}
