package languages

import (
	"context"
	"net/http"

	"github.com/Deirror/servette/app"
	"github.com/Deirror/servette/auth/jwt"

	"github.com/Deirror/servette/translation/languages"
	"github.com/go-chi/chi/v5"
)

type WebMiddleware struct {
	rlv *languages.Resolver

	jwt jwt.Provider

	cfg *appx.Config
}

func NewWebMiddleware(cfg *appx.Config, r *languages.Resolver, jwt jwt.Provider) *WebMiddleware {
	return &WebMiddleware{
		cfg: cfg,
		rlv: r,
		jwt: jwt,
	}
}

func (m *WebMiddleware) LanguageMiddleware(next http.Handler) http.Handler {
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

type Middleware struct {
	rlv *languages.Resolver
}

func NewMiddleware(r *languages.Resolver) *WebMiddleware {
	return &WebMiddleware{
		rlv: r,
	}
}

func (m *Middleware) LanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := m.rlv.FromRequestCookie(r)
		ctx := context.WithValue(r.Context(), languages.LangKey, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
