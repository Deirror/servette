// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package htmx

import (
	"net/http"

	"github.com/Deirror/servette/transport"
	"github.com/Deirror/servette/transport/err"
	httperr "github.com/Deirror/servette/transport/protocol/http/err"
)

func IsHXRequest(r *http.Request) bool {
	return r.Header.Get(HXRequest) == "true"
}

func RequireHXRequest(r *http.Request) *errx.Err {
	if !IsHXRequest(r) {
		return httperr.New(http.StatusBadRequest, transport.RequestBad, "cannot process non-hx requests", nil)
	}
	return nil
}
