package languages

import (
	"context"
	"net/http"
	"strings"

	"github.com/Deirror/servette/auth/jwt"
	"github.com/Deirror/servette/translation/languages"
	"github.com/Deirror/servette/transport/err"

	"github.com/Deirror/servette/app"
)

type Handler struct {
	rlv *languages.Resolver

	jwt jwt.Provider

	cfg *appx.Config
}

func NewHandler(cfg *appx.Config, r *languages.Resolver, jwt jwt.Provider) *Handler {
	return &Handler{
		cfg: cfg,
		rlv: r,
		jwt: jwt,
	}
}

func (h *Handler) HandleSetLanguage(ctx context.Context, w http.ResponseWriter, r *http.Request) *errx.Err {
	lang := h.rlv.FromRequestURL(r)

	h.jwt.SetCookie(w, lang, appx.IsProdMode(h.cfg.Mode))

	path := r.Referer()
	if path == "" {
		path = "/"
	}

	url := "/" + lang + stripLangPrefix(path)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}

// StripLangPrefix removes the leading /{lang} from a path if present.
// e.g. "/en/news" -> "/news", "/fr/about" -> "/about"
// if no lang prefix, returns path unchanged
func stripLangPrefix(path string) string {
	if path == "" || path == "/" {
		return path
	}

	parts := strings.SplitN(path, "/", 3) // ["", "en", "news"] or ["", "en"]
	if len(parts) < 2 {
		return path // no prefix
	}

	langCandidate := parts[1]
	if len(langCandidate) == 2 { // simple heuristic: 2-letter lang code
		if len(parts) == 2 {
			return "/" // only /en -> return root
		}
		return "/" + parts[2] // remove lang
	}

	return path // not a lang prefix
}
