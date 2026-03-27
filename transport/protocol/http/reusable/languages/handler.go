// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package languages

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/Deirror/servette/auth/jwt"
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/translation/languages"
	errx "github.com/Deirror/servette/transport/err"
)

type Handler struct {
	mode env.Mode

	rlv *languages.Resolver

	jwt jwt.Provider
}

func NewHandler(m env.Mode, r *languages.Resolver, jwt jwt.Provider) *Handler {
	return &Handler{
		mode: m,
		rlv:  r,
		jwt:  jwt,
	}
}

func (h *Handler) HandleSetLanguage(ctx context.Context, w http.ResponseWriter, r *http.Request) *errx.Err {
	lang := h.rlv.FromRequestURL(r)
	h.jwt.SetCookie(w, lang, h.mode.IsProd())

	ref := r.Referer()
	if ref == "" {
		ref = "/" // fallback to root
	}

	u, err := url.Parse(ref)
	if err != nil {
		u = &url.URL{Path: "/"}
	}

	cleanPath := stripLangPrefix(u.Path)
	if cleanPath == "/" {
		cleanPath = ""
	}

	finalURL := "/" + lang + cleanPath

	if u.RawQuery != "" {
		finalURL += "?" + u.RawQuery
	}

	http.Redirect(w, r, finalURL, http.StatusTemporaryRedirect)
	return nil
}

// stripLangPrefix removes leading /{lang} from a path (expects only path, no scheme/domain)
func stripLangPrefix(path string) string {
	if path == "" || path == "/" {
		return "/"
	}

	parts := strings.SplitN(path, "/", 3) // ["", "en", "news"] or ["", "en"]
	if len(parts) < 2 {
		return path
	}

	langCandidate := parts[1]
	if len(langCandidate) == 2 { // simple heuristic
		if len(parts) == 2 {
			return "/" // only /en -> root
		}
		return "/" + parts[2] // remove lang
	}

	return path
}
