// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package initx

import (
	"fmt"

	"github.com/Deirror/servette/config"
)

// Helper to get config by key safely
func (c *Config) getCfg(key string) (*config.Config, error) {
	cfg, ok := c.Cfgs[key]
	if !ok {
		return nil, fmt.Errorf("config key not found: %s", key)
	}
	return cfg, nil
}

// EnsureFile checks that ReadMode is File and all resources are FilePath
func (c *Config) EnsureFile(key string) error {
	cfg, err := c.getCfg(key)
	if err != nil {
		return err
	}

	if !cfg.ReadMode.IsFile() {
		return fmt.Errorf("unsupported read mode for %s: %s", key, cfg.ReadMode)
	}

	for _, r := range cfg.Resources {
		if !r.Kind().IsFilePath() {
			return fmt.Errorf("unsupported resource kind for %s: %s", key, r.Kind())
		}
	}

	return nil
}

// EnsureMultiFile applies EnsureFile to multiple keys
func (c *Config) EnsureMultiFile(keys ...string) error {
	for _, k := range keys {
		if err := c.EnsureFile(k); err != nil {
			return err
		}
	}
	return nil
}

// EnsureOS checks that ReadMode is OS and resources are unknown
func (c *Config) EnsureOS(key string) error {
	cfg, err := c.getCfg(key)
	if err != nil {
		return err
	}

	if !cfg.ReadMode.IsOS() {
		return fmt.Errorf("unsupported read mode for %s: %s", key, cfg.ReadMode)
	}

	for _, r := range cfg.Resources {
		if !r.Kind().IsUnknown() {
			return fmt.Errorf("unsupported resource kind for %s: %s", key, r.Kind())
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

// EnsureExt checks that ReadMode is External and resources are URI
func (c *Config) EnsureExt(key string) error {
	cfg, err := c.getCfg(key)
	if err != nil {
		return err
	}

	if !cfg.ReadMode.IsExt() {
		return fmt.Errorf("unsupported read mode for %s: %s", key, cfg.ReadMode)
	}

	for _, r := range cfg.Resources {
		if !r.Kind().IsURI() {
			return fmt.Errorf("unsupported resource kind for %s: %s", key, r.Kind())
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

// EnsureLocal ensures the resource is only File or OS (not External)
func (c *Config) EnsureLocal(key string) error {
	cfg, err := c.getCfg(key)
	if err != nil {
		return err
	}

	if cfg.ReadMode.IsExt() {
		return fmt.Errorf("key %s: ext is unsupported, only file or os", key)
	}

	return nil
}

func (c *Config) EnsureMultiLocal(keys ...string) error {
	for _, k := range keys {
		if err := c.EnsureLocal(k); err != nil {
			return err
		}
	}
	return nil
}
