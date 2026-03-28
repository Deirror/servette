// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package urlx

type Provider interface {
	GetURL() string
	WithQuery(arg, val string) Provider
	WithPath(path string) Provider
}

