// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package urlx

import (
	"net/url"
	"strings"
)

// JoinQuery constructs a raw query string from key=value parts, including the '?' prefix.
func JoinQuery(parts ...string) string {
	values := url.Values{}

	for _, part := range parts {
		if kv := strings.SplitN(part, "=", 2); len(kv) == 2 {
			values.Add(kv[0], kv[1])
		}
	}

	enc := values.Encode()
	if enc == "" {
		return ""
	}
	return "?" + enc
}
