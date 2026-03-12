// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package client

import (
	"net/http"
)

type RequestOpt func(req *http.Request)

func WithCookies(cookies ...*http.Cookie) RequestOpt {
	return func(req *http.Request) {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}
}

func WithHeader(key, value string) RequestOpt {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}
