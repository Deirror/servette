// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/Deirror/servette/transport"
	errx "github.com/Deirror/servette/transport/err"
)

// A handler func which accepts context field and returns an error.
type Func func(context.Context, http.ResponseWriter, *http.Request) *errx.Err

// Custom type of an error handler func mostly used for custom transport logic.
type ErrFunc func(context.Context, http.ResponseWriter, *http.Request, *errx.Err)

// A wrapper func for handling errors from the called handler funcs.
// Uses custom func which handles error.
// Defaults to DefaultErrHandler if nothing is passed.
func Wrap(h Func, onErr ErrFunc) http.HandlerFunc {
	if onErr == nil {
		onErr = DefaultErrHandler
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, transport.ReqID, uuid.NewString())
		if err := h(ctx, w, r); err != nil {
			onErr(ctx, w, r, err)
		}
	}
}
