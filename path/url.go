// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package pathx

import (
	"net/url"
	"strings"
)

// Join constructs a URL path by concatenating multiple path segments.
//
// The function assumes that provided segments already contain the
// correct leading slashes when necessary. No normalization or cleanup
// of slashes is performed.
//
// Example:
//
//	base := "/platform"
//	full := Join(base, "/news", "/latest")
//	// Result: "/platform/news/latest"
func Join(parts ...string) string {
	var b strings.Builder

	for _, p := range parts {
		b.WriteString(p)
	}

	return b.String()
}

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

// Prefix constructs a path by prepending a prefix segment and joining
// additional path segments.
//
// The prefix will automatically be prefixed with a leading slash.
//
// Example:
//
//	path := Prefix("api", "/users", "/123")
//	// Result: "/api/users/123"
func Prefix(prefix string, parts ...string) string {
	var b strings.Builder

	b.WriteString("/")
	b.WriteString(prefix)

	for _, p := range parts {
		b.WriteString(p)
	}

	return b.String()
}
