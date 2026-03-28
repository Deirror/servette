// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package config

import (
    "fmt"
    "strings"
)

type ReadMode string

const (
    File     ReadMode = "file"     // Read from a local file (.env, YAML, etc.)
    OS       ReadMode = "os"       // Read from inherited OS environment variables
    External ReadMode = "external" // Read from external source (Docker secrets, K8s secrets, etc.)
)

// ParseReadMode converts a string to a ReadMode. Case-insensitive.
func ParseReadMode(s string) (ReadMode, error) {
    normalized := strings.ToLower(strings.TrimSpace(s))
    switch normalized {
    case string(File):
        return File, nil
    case string(OS):
        return OS, nil
    case string(External):
        return External, nil
    default:
        return "", fmt.Errorf("invalid ReadMode: %q", s)
    }
}
