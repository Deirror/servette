// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package htmx

import (
	"context"
	"net/http"

	"github.com/Deirror/servette/transport/err"
	"github.com/Deirror/servette/transport/protocol/http/handler"
	"github.com/Deirror/servette/transport/protocol/http/htmx"
)

func RequestMiddleware(next handler.Func) handler.Func {
	return handler.Func(func(ctx context.Context, w http.ResponseWriter, r *http.Request) *errx.Err {
		if err := htmx.RequireHXRequest(r); err != nil {
			return err
		}
		return next(ctx, w, r)
	})
}
