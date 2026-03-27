// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package env provides helpers for loading any configuration
// from environment variables.
//
// The subdirectories integrate with the servette configuration system and supports both
// single and multi-instance configuration loading using environment-based
// prefixes.
//
// # Notes
//
//   - Prefixes are normalized using envcfg.ModPrefix
//   - Missing required variables will result in an error
//   - Documentations for the subdirectories is omitted
package env
