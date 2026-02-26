package client

import (
	"time"
)

// Config contains configuration settings for a HTTP client.
type Config struct {
	BaseURL      string        // Main url and schema for all requests
	ReadTimeout  time.Duration // Maximum duration for reading the entire request
	WriteTimeout time.Duration // Maximum duration before timing out writes
	IdleTimeout  time.Duration // Maximum time to wait for the next request
}

func NewConfig(url string, read, write, idle time.Duration) *Config {
	return &Config{
		BaseURL: url,
		ReadTimeout:  read,
		WriteTimeout: write,
		IdleTimeout:  idle,
	}
}

// WithBaseURL sets the BaseURL field and returns the updated Config.
func (c *Config) WithBaseURL(url string) *Config {
	c.BaseURL = url
	return c
}

// WithReadTimeout sets the ReadTimeout field and returns the updated Config.
func (c *Config) WithReadTimeout(timeout time.Duration) *Config {
	c.ReadTimeout = timeout
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
