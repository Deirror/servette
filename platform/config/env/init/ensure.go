// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package initx

import (
	"errors"
	"fmt"
)

func (c *Config) EnsureFile(key string) error {
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

func (c *Config) EnsureMultiFile(keys ...string) error {
	for _, k := range keys {
		if err := c.EnsureFile(k); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) EnsureOS(key string) error {
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

func (c *Config) EnsureMultiOS(keys ...string) error {
	for _, k := range keys {
		if err := c.EnsureOS(k); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) EnsureExt(key string) error {
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

func (c *Config) EnsureMultiExt(keys ...string) error {
	for _, k := range keys {
		if err := c.EnsureExt(k); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) EnsureLocal(key string) error {
	if err := c.EnsureExt(key); err == nil {
		return errors.New("ext is unsupported, only file or os")
	}

	return nil
}

func (c *Config) EnsureMultiLocal(keys ...string) error {
	for _, k := range keys {
		if err := c.EnsureExt(k); err == nil {
			return errors.New("ext is unsupported, only file or os")
		}
	}

	return nil
}
