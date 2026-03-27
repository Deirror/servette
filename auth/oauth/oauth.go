// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

type Profile struct {
	Email      string `json:"email"`
	ExternalID string `json:"id"`
	Org
}

type Provider interface {
	AuthCodeURL(state string) (string, error)
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	FetchProfile(ctx context.Context, token *oauth2.Token) (*Profile, error)
}
