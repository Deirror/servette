package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/Deirror/servette/transport/err"
	"github.com/Deirror/servette/transport"
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
		ctx = context.WithValue(ctx, transport.ReqIDKey, uuid.NewString())
		if err := h(ctx, w, r); err != nil {
			onErr(ctx, w, r, err)
		}
	}
}
