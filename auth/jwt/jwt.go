package jwt

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT handles JWT token creation, validation, and cookie management.
type JWT struct {
	CookieName string        // Name of the cookie to store the JWT
	Secret     string        // Secret key used to sign the JWT
	TokenTTL   time.Duration // Token time-to-live duration
}

func NewJWT(jwtCfg *Config) *JWT {
	return &JWT{
		CookieName: jwtCfg.CookieName,
		Secret:     jwtCfg.Secret,
		TokenTTL:   jwtCfg.TokenTTL,
	}
}

// GenerateTokenWithClaim creates a signed JWT token containing a single string claim and expiration.
func (j *JWT) GenerateTokenWithClaim(claimKey, claimValue string) (string, error) {
	claims := jwt.MapClaims{
		claimKey: claimValue,
		Exp:      time.Now().Add(j.TokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// SetCookie sets an HTTP-only, secure cookie with the JWT token on the response writer.
// If a domain is provided, the cookie will be accessible across subdomains.
func (j *JWT) SetCookie(w http.ResponseWriter, token string, secure bool, domain ...string) {
	cookie := &http.Cookie{
		Name:     j.CookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		MaxAge:   int(j.TokenTTL.Seconds()),
		Expires:  time.Now().Add(j.TokenTTL), // set Expires for compatibility
	}

	if secure {
		cookie.SameSite = http.SameSiteNoneMode
	} else {
		cookie.SameSite = http.SameSiteLaxMode
	}

	if len(domain) > 0 && domain[0] != "" {
		cookie.Domain = domain[0]
	}

	http.SetCookie(w, cookie)
}

// RemoveCookie removes the JWT cookie by setting an expired Set-Cookie header.
func (j *JWT) RemoveCookie(w http.ResponseWriter, secure bool, domain ...string) {
	cookie := &http.Cookie{
		Name:     j.CookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}

	if secure {
		cookie.SameSite = http.SameSiteNoneMode
	} else {
		cookie.SameSite = http.SameSiteLaxMode
	}

	if len(domain) > 0 && domain[0] != "" {
		cookie.Domain = domain[0]
	}

	http.SetCookie(w, cookie)
}

// GetCookie extracts HTTP Cookie from request based on cookie name.
func (j *JWT) GetCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(j.CookieName)
}

// helper to extract "exp" claim as int64 (unix seconds). Accepts common types.
func expFromClaims(claims jwt.MapClaims) (int64, error) {
	raw, ok := claims[Exp]
	if !ok {
		return 0, fmt.Errorf("missing exp claim")
	}

	switch v := raw.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case string:
		// sometimes exp may be serialized as string
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid exp claim string: %w", err)
		}
		return i, nil
	default:
		return 0, fmt.Errorf("unsupported exp claim type: %T", raw)
	}
}

// generic validator that extracts a string claim (claimKey) from the token.
// returns the claim string (e.g. email or lang) or an error.
func (j *JWT) ValidateClaim(tokenString, claimKey string) (string, error) {
	if tokenString == "" {
		return "", fmt.Errorf("empty token")
	}

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		// prefer explicit alg check
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header[Alg])
		}
		return []byte(j.Secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return "", fmt.Errorf("parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token claims")
	}

	exp, err := expFromClaims(claims)
	if err != nil {
		return "", err
	}

	if time.Now().Unix() > exp+5 {
		return "", fmt.Errorf("token has expired")
	}

	val, ok := claims[claimKey].(string)
	if !ok || val == "" {
		return "", fmt.Errorf("%s claim missing or invalid", claimKey)
	}

	return val, nil
}
