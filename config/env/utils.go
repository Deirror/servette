// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package env

import "strings"

func ModPrefix(prefix ...string) string {
	pfx := ""
	if len(prefix) > 0 && prefix[0] != "" {
		pfx = prefix[0]
		if !strings.HasSuffix(pfx, "_") {
			pfx += "_"
		}
	}
	return pfx
}
