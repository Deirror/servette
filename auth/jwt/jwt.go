// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package jwt

import (
	"net/http"
)

type Provider interface {
	GenerateTokenWithClaim(claimKey, claimValue string) (string, error)
	ValidateClaim(tokenString, claimKey string) (string, error)

	SetCookie(w http.ResponseWriter, token string, secure bool, domain ...string)
	RemoveCookie(w http.ResponseWriter, secure bool, domain ...string)
	GetCookie(r *http.Request) (*http.Cookie, error)
}
