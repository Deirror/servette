// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package pathx

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Deirror/servette/env"
)

// GetProjectRoot returns the project root directory.
// It first checks the prefix + APP_MODE env variable. If not set,
// It falls back to searching for go.mod in parent directories.
// In prod, the env vars are already loaded, meanwhile in dev mode they must be loaded from a file.
func GetProjectRootFromAppMode(appPrefix string) (string, error) {
	modeEnv, _ := env.Get(appPrefix + "_APP_MODE")
	mode, err := env.ParseMode(modeEnv)
	if err == nil {
		if mode.IsProd() {
			return "", err
		}
	}

	// Fallback to searching for go.mod
	root, err := FindProjectRoot("go.mod")
	if err != nil {
		return "", err
	}
	return root, nil
}

// FindProjectRoot returns the project root directory.
// Uses the executable directory and goes up until it finds a marker.
func FindProjectRoot(markers ...string) (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(exe)

	for {
		for _, m := range markers {
			if _, err := os.Stat(filepath.Join(dir, m)); err == nil {
				return dir, nil
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("project root not found")
		}
		dir = parent
	}
}
