// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package urlx

import (
	"net/url"
)

// Client contains all pices for keeping a base url,
// an additional path and raw query.
type Client struct {
	base  *url.URL
	path  string
	query url.Values
}

func NewClient(cfg *Config) (*Client, error) {
	u, err := url.Parse(cfg.URL)
	if err != nil {
		return nil, err
	}

	return &Client{
		base:  u,
		query: url.Values{},
	}, nil
}

// clone makes a copy of the current client.
func (c *Client) clone() *Client {
	newQuery := url.Values{}
	for k, v := range c.query {
		newQuery[k] = append([]string{}, v...)
	}

	var baseCopy *url.URL
	if c.base != nil {
		u := *c.base
		baseCopy = &u
	}

	return &Client{
		base:  baseCopy,
		path:  c.path,
		query: newQuery,
	}
}

// GetURL combines all variables to create the whole URL.
func (c *Client) GetURL() string {
	if c.base == nil {
		return ""
	}

	u := *c.base

	if c.path != "" {
		u.Path = c.path
	}

	if len(c.query) > 0 {
		u.RawQuery = c.query.Encode()
	}

	return u.String()
}

// WithPath creates a clone client and sets raw query.
func (c *Client) WithQuery(arg, val string) Provider {
	clone := c.clone()
	clone.query.Set(arg, val)
	return clone
}

// WithPath creates a clone client and sets path.
func (c *Client) WithPath(path string) Provider {
	clone := c.clone()
	clone.path = path
	return clone
}
