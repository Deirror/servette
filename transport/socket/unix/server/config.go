package server

import (
	"time"
)

// Config mirrors the http server config.
type Config struct {
	Port         string        // for UNIX socket: the socket path
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewConfig(port string, read, write, idle time.Duration) *Config {
	return &Config{
		Port:         port,
		ReadTimeout:  read,
		WriteTimeout: write,
		IdleTimeout:  idle,
	}
}

// WithPort sets the Port field and returns the updated Config.
func (c *Config) WithPort(port string) *Config {
	c.Port = port
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
