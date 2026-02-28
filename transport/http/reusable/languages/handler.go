package languages

import (
	"context"
	"net/http"

	"github.com/Deirror/servette/auth/jwt"
	"github.com/Deirror/servette/translation/languages"
	"github.com/Deirror/servette/transport/err"
	"github.com/Deirror/servette/transport/http/handler"

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

	handler.SafeRedirect(w, r)
	return nil
}
