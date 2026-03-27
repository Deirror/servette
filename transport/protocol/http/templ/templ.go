// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package templ

import (
	"bytes"
	"context"
	"net/http"

	"github.com/a-h/templ"

	"github.com/Deirror/servette/transport"
	errx "github.com/Deirror/servette/transport/err"
	httperr "github.com/Deirror/servette/transport/protocol/http/err"
	"github.com/Deirror/servette/transport/protocol/http/header"
)

func Render(ctx context.Context, c templ.Component, w http.ResponseWriter) *errx.Err {
	// Render page into a buffer
	var buf bytes.Buffer
	if err := c.Render(ctx, &buf); err != nil {
		return httperr.New(http.StatusInternalServerError, transport.TemplFail, "cannot render template", err)
	}

	// Success - write headers + body
	w.Header().Set(header.ContentType, header.TextHTML)
	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)

	return nil
}
