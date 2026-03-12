// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package oauth

import (
	"golang.org/x/oauth2"
)

func NewOAuth2Config(cfg *Config, p Props) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       p.Scopes,
		Endpoint:     p.Endpoint,
	}
}
