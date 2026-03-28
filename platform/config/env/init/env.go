// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package initx

import (
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/path"
)

func (c *Config) LoadEnvFiles(keys ...string) error {
	filenames := []string{}
	for _, k := range keys {
		filenames = append(filenames, pathx.ResourcesToStrings(c.Cfgs[k].Resources)...)
	}

	if err := env.Load(filenames...); err != nil {
		return err
	}

	return nil
}
