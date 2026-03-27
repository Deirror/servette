// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package oauth

import (
	"context"

	"github.com/Deirror/servette/encoding/json"
	"golang.org/x/oauth2"
)

type Client struct {
	cfg *oauth2.Config
	org Org
}

func NewClient(cfg *Config, o Org) (*Client, error) {
	return &Client{
		cfg: NewOAuth2Config(cfg, props[o]),
		org: o,
	}, nil
}

func (c *Client) AuthCodeURL(state string) (string, error) {
	return c.cfg.AuthCodeURL(state), nil
}

func (c *Client) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return c.cfg.Exchange(ctx, code)
}

func (c *Client) FetchProfile(ctx context.Context, t *oauth2.Token) (*Profile, error) {
	url := props[c.org].UserInfoURL

	cl := c.cfg.Client(ctx, t)
	resp, err := cl.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var p Profile
	if err = json.DecodeInto(resp.Body, &p); err != nil {
		return nil, err
	}
	p.Org = c.org
	return &p, nil
}
