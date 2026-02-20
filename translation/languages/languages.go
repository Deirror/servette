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

func (rlv *Resolver) LangFromRequest(r *http.Request) string {
	lang := r.URL.Query().Get(LangKey)

	if !rlv.IsSupported(lang) {
		return rlv.DefaultLang
	}

	return lang
}

func (rlv *Resolver) ContextWithLang(r *http.Request) context.Context {
	lang := rlv.DefaultLang

	if cookie, err := r.Cookie(LangKey); err == nil {
		if rlv.IsSupported(cookie.Value) {
			lang = cookie.Value
		}
	}

	return context.WithValue(r.Context(), LangKey, lang)
}

func (r *Resolver) LangFromContext(ctx context.Context) string {
	val := ctx.Value(LangKey)
	if lang, ok := val.(string); ok && r.IsSupported(lang) {
		return lang
	}

	return r.DefaultLang
}

func (r *Resolver) IsSupported(lang string) bool {
	return slices.Contains(r.SupportedLangs, lang)
}
