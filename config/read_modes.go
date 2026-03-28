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
    Ext ReadMode = "ext" // Read from external source (Docker secrets, K8s secrets, etc.)
)

// ParseReadMode converts a string to a ReadMode. Case-insensitive.
func ParseReadMode(s string) (ReadMode, error) {
    normalized := strings.ToLower(strings.TrimSpace(s))
    switch normalized {
    case string(File):
        return File, nil
    case string(OS):
        return OS, nil
    case string(Ext):
        return Ext, nil
    default:
        return "", fmt.Errorf("invalid ReadMode: %q", s)
    }
}

func (r ReadMode) IsFile() bool {
	return r == File
}

func (r ReadMode) IsOS() bool {
	return r == OS
}

func (r ReadMode) IsExt() bool {
	return r == Ext
}
