package languages

import (
	"net/http"
	"context"

	"github.com/Deirror/servette/translation/languages"
)

type Middleware struct {
	rlv *languages.Resolver
}

func NewMiddleware(r *languages.Resolver) *Middleware {
	return &Middleware{
		rlv: r,
	}
}

func (m *Middleware) LanguageMiddleware(next http.Handler, reqType languages.RequestType) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := m.rlv.FromRequest(r, reqType)
		ctx := context.WithValue(r.Context(), languages.LangKey, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
