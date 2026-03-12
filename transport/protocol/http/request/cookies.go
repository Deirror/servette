// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package request

import (
	"errors"
	"net/http"
)

func Cookies(r *http.Request, keys ...string) ([]*http.Cookie, error) {
	if len(keys) < 1 {
		return nil, errors.New("empty keys slice")
	}

	var cookies []*http.Cookie
	for _, key := range keys {
		val, err := r.Cookie(key)
		if err != nil {
			return nil, err
		}
		cookies = append(cookies, val)
	}

	return cookies, nil
}
