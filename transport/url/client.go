// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package urlx

import (
	"net/url"
)

type ExternalURLClient struct {
	eURL string
}

func NewExternalURLClient(cfg *Config) *ExternalURLClient {
	return &ExternalURLClient{
		eURL: cfg.URL,
	}
}

func (c *ExternalURLClient) GetURL() string {
	return c.eURL
}

func (c *ExternalURLClient) WithQuery(arg, val string) string {
	u, err := url.Parse(c.eURL)
	if err != nil {
		return c.eURL
	}

	q := u.Query()
	q.Set(arg, val)
	u.RawQuery = q.Encode()

	return u.String()
}
