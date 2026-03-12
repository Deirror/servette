// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package htmx

import (
	"net/http"
)

func IsHXRequest(r *http.Request) bool {
	return r.Header.Get(HXRequest) == "true"
}
