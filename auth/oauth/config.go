// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package oauth

// Config holds OAuth 2.0 client credentials and redirect URL.
type Config struct {
	ClientID     string // OAuth client ID
	ClientSecret string // OAuth client secret
	RedirectURL  string // OAuth redirect URL after authentication
}

func NewConfig(clientID, clientSecret, redirectURL string) *Config {
	return &Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
	}
}

// WithClientID sets the OAuth client ID and returns the updated config.
func (c *Config) WithClientID(clientID string) *Config {
	c.ClientID = clientID
	return c
}

// WithClientSecret sets the OAuth client secret and returns the updated config.
func (c *Config) WithClientSecret(clientSecret string) *Config {
	c.ClientSecret = clientSecret
	return c
}

// WithRedirectURL sets the OAuth redirect URL and returns the updated config.
func (c *Config) WithRedirectURL(redirectURL string) *Config {
	c.RedirectURL = redirectURL
	return c
}
