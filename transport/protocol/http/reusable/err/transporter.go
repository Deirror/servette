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

func EmplaceTransporter(log *slog.Logger) *Transporter {
	m := NewMiddleware(log)

	return NewTransporter(m)
}
