package languages

import (
	"net/http"

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

func (m *Middleware) LanguageMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := m.rlv.ContextWithLang(r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
