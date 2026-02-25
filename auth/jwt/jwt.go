package jwt

import (
	"net/http"
)

type Provider interface {
	GenerateToken(email string) (string, error)
	SetCookie(w http.ResponseWriter, token string, secure bool, domains ...string)
	RemoveCookie(w http.ResponseWriter, secure bool, domains ...string)
	ValidateJWT(token string) (string, error)
	GetCookie(r *http.Request) (*http.Cookie, error)
}
