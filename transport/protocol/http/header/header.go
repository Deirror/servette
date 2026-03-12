// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package header

import (
	"net/http"
)

func Set(w http.ResponseWriter, key string, val string) {
	w.Header().Set(key, val)
}
