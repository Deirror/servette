// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package env

import (
	"fmt"
	"strings"

	"github.com/Deirror/servette/env"
)

// GroupMap maps a prefix (e.g. "WEB", "EXE") to a set of environment variable suffixes and their values.
// For example, for "WEB_JWT_SECRET", the prefix is "WEB" and the suffix is "_JWT_SECRET".
type GroupMap map[string]map[string]string

// Extracts all env vars ending with the suffixes.
func LoadGroups(suffixes []string, filenames ...string) (GroupMap, error) {
	grouped := make(GroupMap)

	envVars, err := env.GetAll(filenames...)
	if err != nil {
		return nil, err
	}

	for key, val := range envVars {
		for _, suffix := range suffixes {
			if !strings.HasSuffix(key, "_"+suffix) {
				continue
			}

			prefix := strings.TrimSuffix(key, "_"+suffix)

			if len(prefix) == 0 {
				return nil, fmt.Errorf("prefix is empty: _ for suffix: %s", suffix)
			}

			if _, ok := grouped[prefix]; !ok {
				grouped[prefix] = make(map[string]string)
			}
			grouped[prefix][suffix] = val
		}
	}

	return grouped, nil
}

// Retrieves env vars based on prefix.
func (m GroupMap) GetGroup(prefix string) map[string]string {
	return m[prefix]
}
