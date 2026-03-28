// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package initx

import (
	"fmt"

	"github.com/Deirror/servette/path"
)

func (c *Config) Subresources(keys ...string) ([]pathx.Resource, error) {
	var subrs []pathx.Resource
	for _, k := range keys {
		if cfg, ok := c.Cfgs[k]; ok {
			subrs = append(subrs, cfg.Resources...)
		} else {
			return nil, fmt.Errorf("config key not found: %s", k)
		}
	}
	return subrs, nil
}

func (c *Config) Substrings(keys ...string) ([]string, error) {
	subrs, err := c.Subresources(keys...)
	if err != nil {
		return nil, err
	}
	return pathx.ResourcesToStrings(subrs), nil
}
