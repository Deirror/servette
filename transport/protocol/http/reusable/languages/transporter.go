package languages

import (
	"github.com/Deirror/servette/app"
	"github.com/Deirror/servette/auth/jwt"
	"github.com/Deirror/servette/translation/languages"
)

type Transporter struct {
	*Handler
	*Middleware
}

func NewTransporter(h *Handler, m *Middleware) *Transporter {
	return &Transporter{
		Handler:    h,
		Middleware: m,
	}
}

func EmplaceTransporter(cfg *appx.Config, r *languages.Resolver, jwt *jwt.JWT) *Transporter {
	h := NewHandler(cfg, r, jwt)
	m := NewMiddleware(r)

	return NewTransporter(h, m)
}
