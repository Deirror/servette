// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package initx

import "fmt"

func (c *Config) EnsureOnlyFile(key string) error {
	rm := c.Cfgs[key].ReadMode
	if !rm.IsFile() {
		return fmt.Errorf("unsupported read mode: %s", string(rm))
	}

	rs := c.Cfgs[key].Resources
	for _, r := range rs {
		if !r.Kind().IsFilePath() {
			return fmt.Errorf("unsupported resource kind: %s", string(r.Kind()))
		}
	}

	return nil
}

func (c *Config) EnsureOnlyOS(key string) error {
	rm := c.Cfgs[key].ReadMode
	if !rm.IsOS() {
		return fmt.Errorf("unsupported read mode: %s", string(rm))
	}

	rs := c.Cfgs[key].Resources
	for _, r := range rs {
		if !r.Kind().IsUnknown() {
			return fmt.Errorf("unsupported resource kind: %s", string(r.Kind()))
		}
	}

	return nil
}

func (c *Config) EnsureOnlyExt(key string) error {
	rm := c.Cfgs[key].ReadMode
	if !rm.IsExt() {
		return fmt.Errorf("unsupported read mode: %s", string(rm))
	}

	rs := c.Cfgs[key].Resources
	for _, r := range rs {
		if !r.Kind().IsURI() {
			return fmt.Errorf("unsupported resource kind: %s", string(r.Kind()))
		}
	}

	return nil
}
