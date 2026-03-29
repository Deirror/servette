// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package pathx

type Path string

// Root returns the absolute path with leading slash, e.g., "/static"
func (p Path) Root() string {
	return "/" + string(p)
}

// Relative returns the relative path with trailing slash, e.g., "static/"
func (p Path) Relative() string {
	return string(p) + "/"
}

// Mid returns a prefixed path with "/" + path + "/", e.g., "/static/"
func (p Path) Mid() string {
	return "/" + string(p) + "/"
}
