package jwt

import (
	"net/http"
)

type Provider interface {
	GenerateTokenWithClaim(claimKey, claimValue string) (string, error)
	SetCookie(w http.ResponseWriter, token string, secure bool, domain ...string)
	RemoveCookie(w http.ResponseWriter, secure bool, domain ...string)
	GetCookie(r *http.Request) (*http.Cookie, error)
	ValidateClaim(tokenString, claimKey string) (string, error)
}
