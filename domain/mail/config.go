// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package mail

// Config holds SMTP configuration for sending emails.
type Config struct {
	Host     string // SMTP server host
	Port     string // SMTP server port
	Username string // SMTP username for authentication
	Password string // SMTP password or token for authentication
	From     string // Default "From" email address
}

func NewConfig(host, port, username, password, from string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

// WithHost sets the SMTP host and returns the updated Config.
func (c *Config) WithHost(host string) *Config {
	c.Host = host
	return c
}

// WithPort sets the SMTP port and returns the updated Config.
func (c *Config) WithPort(port string) *Config {
	c.Port = port
	return c
}

// WithUsername sets the SMTP username and returns the updated Config.
func (c *Config) WithUsername(username string) *Config {
	c.Username = username
	return c
}

// WithPassword sets the SMTP password and returns the updated Config.
func (c *Config) WithPassword(password string) *Config {
	c.Password = password
	return c
}

// WithFrom sets the default "From" email address and returns the updated Config.
func (c *Config) WithFrom(from string) *Config {
	c.From = from
	return c
}
