// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package languages

import (
	"github.com/Deirror/servette/auth/jwt"
	"github.com/Deirror/servette/env"
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

func EmplaceTransporter(mode env.Mode, t ResolveType, r *languages.Resolver, jwt *jwt.JWT) *Transporter {
	h := NewHandler(mode, r, jwt)
	m := NewMiddleware(mode, t, r, jwt)

	return NewTransporter(h, m)
}
