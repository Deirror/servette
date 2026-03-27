// Copyright 2025 Deirror. All rights reserved.
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

type ResolveType int

const (
	Cookie            ResolveType = iota
	URLParam          ResolveType = 1
	CookieAndURLParam ResolveType = 2
)

type Middleware struct {
	mode env.Mode

	rt  ResolveType
	rlv *languages.Resolver

	jwt jwt.Provider
}

func NewMiddleware(m env.Mode, t ResolveType, r *languages.Resolver, jwt jwt.Provider) *Middleware {
	return &Middleware{
		mode: m,
		rt:   t,
		rlv:  r,
		jwt:  jwt,
	}
}

func (m *Middleware) LanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := ""
		switch m.rt {
		case Cookie:
			lang = m.rlv.FromRequestCookie(r)
		case URLParam:
			lang = chi.URLParam(r, languages.Lang)
		case CookieAndURLParam:
			cookieLang := m.rlv.FromRequestCookie(r)
			urlLang := chi.URLParam(r, languages.Lang)
			lang = cookieLang
			if urlLang != cookieLang && m.rlv.IsSupported(urlLang) {
				lang = urlLang
				m.jwt.SetCookie(w, lang, m.mode.IsProd())
			}
		default:
			break
		}
		ctx := context.WithValue(r.Context(), languages.Lang, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
