// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package oauth

import (
	"context"

	"github.com/Deirror/servette/encoding/json"
	"golang.org/x/oauth2"
)

// Client wraps an OAuth2 configuration and a specific provider (Org).
// It provides helper methods for the standard OAuth2 authorization flow.
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

// AuthCodeURL generates the OAuth2 authorization URL.
//
// This URL should be used to redirect the user to the provider’s
// consent/login page. The state parameter should be a securely
// generated random string used to prevent CSRF attacks.
//
// Example:
//
//	url, _ := client.AuthCodeURL("random-state")
//	http.Redirect(w, r, url, http.StatusFound)
func (c *Client) AuthCodeURL(state string) (string, error) {
	return c.cfg.AuthCodeURL(state), nil
}

// Exchange exchanges an authorization code for an OAuth2 token.
//
// The code is obtained from the OAuth provider after the user
// completes the authentication flow and is redirected back to
// your application.
//
// The returned token contains access and optionally refresh tokens.
func (c *Client) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return c.cfg.Exchange(ctx, code)
}

// FetchProfile retrieves the authenticated user's profile from
// the provider's user info endpoint.
//
// It uses the provided OAuth2 token to authorize the request.
// The response is decoded into a Profile struct using the
// project's JSON decoder.
//
// The returned Profile will have its Org field populated with
// the provider used to fetch it.
//
// Returns an error if the request fails or the response cannot be decoded.
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
