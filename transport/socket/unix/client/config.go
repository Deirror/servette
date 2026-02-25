package client

import (
	"time"
)

// Config contains configuration settings for a HTTP client.
type Config struct {
	SocketPath   string        // "Port" for dialing.
	WriteTimeout time.Duration // Maximum duration before timing out writes
	IdleTimeout  time.Duration // Maximum time to wait for the next request
}

func NewConfig(path string, write, idle time.Duration) *Config {
	return &Config{
		SocketPath:   path,
		WriteTimeout: write,
		IdleTimeout:  idle,
	}
}

// WithSocketPath sets the path field and returns the updated Config.
func (c *Config) WithSocketPath(path string) *Config {
	c.SocketPath = path
	return c
}

// WithWriteTimeout sets the WriteTimeout field and returns the updated Config.
func (c *Config) WithWriteTimeout(timeout time.Duration) *Config {
	c.WriteTimeout = timeout
	return c
}

// WithIdleTimeout sets the IdleTimeout field and returns the updated Config.
func (c *Config) WithIdleTimeout(timeout time.Duration) *Config {
	c.IdleTimeout = timeout
	return c
}
