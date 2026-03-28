// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package urlx

import (
	"net/url"
)

type URLClient struct {
	URL string
}

func NewURLClient(cfg *Config) *URLClient {
	return &URLClient{
		URL: cfg.URL,
	}
}

func (c *URLClient) GetURL() string {
	return c.URL
}

func (c *URLClient) WithQuery(arg, val string) string {
	u, err := url.Parse(c.URL)
	if err != nil {
		return c.URL
	}

	q := u.Query()
	q.Set(arg, val)
	u.RawQuery = q.Encode()

	return u.String()
}
