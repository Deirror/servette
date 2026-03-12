// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package languages

import (
	"context"
	"net/http"

	"github.com/Deirror/servette/auth/jwt"
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/translation/languages"

	"github.com/go-chi/chi/v5"
)

type Middleware struct {
	mode env.Mode

	rlv *languages.Resolver

	jwt jwt.Provider
}

func NewMiddleware(r *languages.Resolver) *Middleware {
	return &Middleware{
		rlv: r,
	}
}

func NewWebMiddleware(m env.Mode, r *languages.Resolver, jwt jwt.Provider) *Middleware {
	return &Middleware{
		mode: m,
		rlv:  r,
		jwt:  jwt,
	}
}

func (m *Middleware) LanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := m.rlv.FromRequestCookie(r)
		ctx := context.WithValue(r.Context(), languages.Lang, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) LanguageWebMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieLang := m.rlv.FromRequestCookie(r)
		urlLang := chi.URLParam(r, languages.Lang)
		lang := cookieLang
		if urlLang != cookieLang && m.rlv.IsSupported(urlLang) {
			lang = urlLang
			m.jwt.SetCookie(w, lang, m.mode.IsProd())
		}
		ctx := context.WithValue(r.Context(), languages.Lang, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
