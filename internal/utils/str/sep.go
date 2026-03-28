// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package str

import "strings"

// SplitAndTrim splits a string by the given separator(s), trims spaces, and
// ignores empty strings. Returns a slice of strings.
func SplitAndTrim(s string, seps ...string) []string {
    if len(seps) == 0 {
        seps = []string{","} // default separator
    }

    // Use the first separator for simplicity
    parts := strings.Split(s, seps[0])
    result := make([]string, 0, len(parts))
    for _, p := range parts {
        trimmed := strings.TrimSpace(p)
        if trimmed != "" {
            result = append(result, trimmed)
        }
    }
    return result
}
