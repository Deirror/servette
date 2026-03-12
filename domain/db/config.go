// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package db

import (
	"time"
)

// Config holds the configuration parameters for a database connection.
type Config struct {
	Driver      string        // Driver is used to set db engine (postgres, mysql and so on)
	DSN         string        // DSN is used to connect to the database
	PoolSize    uint8         // Maximum number of open connections in the pool
	MaxIdle     uint8         // Maximum number of idle connections in the pool
	MaxLifetime time.Duration // Maximum lifetime of a connection before it's recycled
}

func NewConfig(driver, dsn string, size, maxIdle uint8, maxLT time.Duration) *Config {
	return &Config{
		Driver:      driver,
		DSN:         dsn,
		PoolSize:    size,
		MaxIdle:     maxIdle,
		MaxLifetime: maxLT,
	}
}

// WithPoolSize sets the PoolSize and returns the updated Config.
func (c *Config) WithPoolSize(size uint8) *Config {
	c.PoolSize = size
	return c
}

// WithIdle sets the MaxIdle value and returns the updated Config.
func (c *Config) WithIdle(idle uint8) *Config {
	c.MaxIdle = idle
	return c
}

// WithMaxLifetime sets the MaxLifetime value and returns the updated Config.
func (c *Config) WithMaxLifetime(maxTL time.Duration) *Config {
	c.MaxLifetime = maxTL
	return c
}

// WithDSN sets the DSN string and returns the updated Config.
func (c *Config) WithDSN(dsn string) *Config {
	c.DSN = dsn
	return c
}

// WithDriver sets the Driver string and returns the updated Config.
func (c *Config) WithDriver(driver string) *Config {
	c.Driver = driver
	return c
}
