package pathx

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Deirror/servette/env"
)

// GetProjectRoot returns the project root directory.
// It first checks the prefix + APP_MODE env variable. If not set,
// it falls back to searching for go.mod in parent directories.
// In prod, the env vars are already loaded, meanwhile in dev mode they must be loaded from a file.
func GetProjectRoot(envVar string) (string, error) {
	if mode, _ := env.Get(envVar); mode == env.Prod {
		return "", nil
	}

	// fallback to searching for go.mod
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
