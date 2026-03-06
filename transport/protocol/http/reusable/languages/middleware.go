package languages

import (
	"context"
	"net/http"

	"github.com/Deirror/servette/app"
	"github.com/Deirror/servette/auth/jwt"

	"github.com/Deirror/servette/translation/languages"
	"github.com/go-chi/chi/v5"
)

type Middleware struct {
	cfg *appx.Config

	rlv *languages.Resolver

	jwt jwt.Provider
}

func NewMiddleware(r *languages.Resolver) *Middleware {
	return &Middleware{
		rlv: r,
	}
}

func NewWebMiddleware(cfg *appx.Config, r *languages.Resolver, jwt jwt.Provider) *Middleware {
	return &Middleware{
		cfg: cfg,
		rlv: r,
		jwt: jwt,
	}
}

func (m *Middleware) LanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := m.rlv.FromRequestCookie(r)
		ctx := context.WithValue(r.Context(), languages.LangKey, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) LanguageWebMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieLang := m.rlv.FromRequestCookie(r)
		urlLang := chi.URLParam(r, languages.LangKey)
		lang := cookieLang
		if urlLang != cookieLang && m.rlv.IsSupported(urlLang) {
			lang = urlLang
			m.jwt.SetCookie(w, lang, appx.IsProdMode(m.cfg.Mode))
		}
		ctx := context.WithValue(r.Context(), languages.LangKey, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
