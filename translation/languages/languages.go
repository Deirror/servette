// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package languages

import (
	"context"
	"net/http"
	"slices"
)

// Resolver determines the appropriate language for a request.
type Resolver struct {
	DefaultLang    string   // Fallback language if none is provided or supported
	SupportedLangs []string // List of supported language codes
}

func NewResolver(defaultLang string, supported ...string) *Resolver {
	// Ensure defaultLang is always included in SupportedLangs
	langs := append([]string{defaultLang}, supported...)

	return &Resolver{
		DefaultLang:    defaultLang,
		SupportedLangs: langs,
	}
}

// FromRequestURL returns the language code from the request query parameter.
// Falls back to DefaultLang if the value is missing or not supported.
func (rlv *Resolver) FromRequestURL(r *http.Request) string {
	lang := r.URL.Query().Get(Lang)

	if !rlv.IsSupported(lang) {
		return rlv.DefaultLang
	}

	return lang
}

// FromRequestCookie returns the language code stored in the request cookie.
// Falls back to DefaultLang if the cookie is missing or contains an unsupported value.
func (rlv *Resolver) FromRequestCookie(r *http.Request) string {
	lang := rlv.DefaultLang

	if cookie, err := r.Cookie(Lang); err == nil {
		if rlv.IsSupported(cookie.Value) {
			lang = cookie.Value
		}
	}

	return lang
}

// FromContext returns the language code stored in the context.
// Falls back to DefaultLang if the value is missing or not supported.
func (r *Resolver) FromContext(ctx context.Context) string {
	val := ctx.Value(Lang)
	if lang, ok := val.(string); ok && r.IsSupported(lang) {
		return lang
	}

	return r.DefaultLang
}

// IsSupported returns true if the given language code is in SupportedLangs.
func (r *Resolver) IsSupported(lang string) bool {
	return slices.Contains(r.SupportedLangs, lang)
}
