// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package errx

import (
	"log/slog"
)

type Transporter struct {
	*Middleware
}

func NewTransporter(m *Middleware) *Transporter {
	return &Transporter{
		Middleware: m,
	}
}

func EmplaceTransporter(log *slog.Logger, t WriteType) *Transporter {
	m := NewMiddleware(log, t)

	return NewTransporter(m)
}
