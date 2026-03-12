// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package env

import "github.com/Deirror/servette/config"

// Template func for loading envs with their prefixes.
func LoadMultiConfig[T any](
	suffixes []string,
	loader func(prefix ...string) (*T, error),
) (config.MultiConfig[T], error) {
	grouped, err := LoadGroups(suffixes)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*T)
	for prefix := range grouped {
		conf, err := loader(prefix)
		if err != nil {
			return nil, err
		}
		result[prefix] = conf
	}

	return result, nil
}
