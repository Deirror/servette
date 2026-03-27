// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package languages

import (
	"context"
	"net/http"
	"slices"
)

type Resolver struct {
	DefaultLang    string
	SupportedLangs []string
}

func NewResolver(defaultLang string, supported ...string) *Resolver {
	// Ensure defaultLang is always included in SupportedLangs
	langs := append([]string{defaultLang}, supported...)

	return &Resolver{
		DefaultLang:    defaultLang,
		SupportedLangs: langs,
	}
}

func (rlv *Resolver) FromRequestURL(r *http.Request) string {
	lang := r.URL.Query().Get(Lang)

	if !rlv.IsSupported(lang) {
		return rlv.DefaultLang
	}

	return lang
}

func (rlv *Resolver) FromRequestCookie(r *http.Request) string {
	lang := rlv.DefaultLang

	if cookie, err := r.Cookie(Lang); err == nil {
		if rlv.IsSupported(cookie.Value) {
			lang = cookie.Value
		}
	}

	return lang
}

func (r *Resolver) FromContext(ctx context.Context) string {
	val := ctx.Value(Lang)
	if lang, ok := val.(string); ok && r.IsSupported(lang) {
		return lang
	}

	return r.DefaultLang
}

func (r *Resolver) IsSupported(lang string) bool {
	return slices.Contains(r.SupportedLangs, lang)
}
