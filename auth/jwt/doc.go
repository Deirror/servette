// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package jwt provides a minimal JWT utility layer for generating,
// validating, and transporting tokens via HTTP cookies.
//
// It is designed around a simple Provider interface that abstracts
// token creation, validation, and cookie management, allowing
// interchangeable implementations if needed.
//
// The default implementation (JWT) uses HMAC-SHA256 signing via
// github.com/golang-jwt/jwt/v5 and supports:
//
//   - Generating tokens with a single string claim
//   - Validating and extracting a specific claim from a token
//   - Managing JWT tokens via HTTP cookies (set, get, remove)
//
// Tokens include an expiration ("exp") claim and are validated
// manually with a small tolerance window.
//
// # Core Concepts
//
// Provider
//
//	The Provider interface defines the contract for JWT operations:
//	  - GenerateTokenWithClaim: create a token with a single claim
//	  - ValidateClaim: validate a token and extract a claim value
//	  - Cookie helpers: SetCookie, GetCookie, RemoveCookie
//
// JWT
//
//	The JWT struct is the default Provider implementation. It is
//	configured with:
//	  - CookieName: name of the HTTP cookie used to store the token
//	  - Secret: HMAC secret used for signing and validation
//	  - TokenTTL: lifetime of generated tokens
//
// # Cookie Behavior
//
// Tokens are typically transported via HTTP-only cookies:
//
//   - HttpOnly is always enabled for security
//   - Secure flag is configurable (recommended true in production)
//   - SameSite:
//   - None when Secure = true (cross-site usage)
//   - Lax when Secure = false
//   - Domain can be optionally set for cross-subdomain usage
//
// # Validation Rules
//
//   - Only HS256 signing method is accepted
//   - Token must be structurally valid and signed with the correct secret
//   - "exp" claim is required and checked manually
//   - A small grace period (+5 seconds) is allowed for expiration
//   - The requested claim must exist and be a non-empty string
//
// # Usage Example
//
//	jwtProvider := jwt.NewJWT(&jwt.Config{
//	    CookieName: "auth_token",
//	    Secret:     "super-secret",
//	    TokenTTL:   time.Hour,
//	})
//
//	token, _ := jwtProvider.GenerateTokenWithClaim("user_id", "123")
//
//	jwtProvider.SetCookie(w, token, true)
//
//	cookie, _ := jwtProvider.GetCookie(r)
//	userID, err := jwtProvider.ValidateClaim(cookie.Value, "user_id")
//
// # Notes
//
//   - This package focuses on simplicity and single-claim usage.
//   - It can be extended to support multiple claims or custom claim types.
//   - Ensure the Secret is kept secure and sufficiently random in production.
package jwt
