// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package pathx provides utilities for working with file system paths
// and URL-style paths in Go applications.
//
// It includes helpers to:
//   - Detect the project root directory based on environment variables or marker files
//   - Join URL path segments consistently
//   - Prepend prefixes to paths for routing or API construction
//
// # Project Root Detection
//
// GetProjectRootFromAppMode attempts to determine the project root directory
// by first checking environment variables (APP_MODE) and falling back to
// searching for marker files such as go.mod in parent directories.
//
// FindProjectRoot scans parent directories starting from the executable
// location until it finds a specified marker file.
//
// # URL Path Helpers
//
// Join concatenates multiple path segments into a single string.
//   - Assumes segments are correctly formatted with slashes
//   - Does not normalize or clean slashes
//
// Prefix prepends a prefix segment (automatically adding a leading slash)
// and joins additional path segments.
//
// # Usage Examples
//
// Detect project root:
//
//	root, err := pathx.GetProjectRootFromAppMode("MYAPP")
//
// Join paths:
//
//	fullPath := pathx.Join("/platform", "/news", "/latest")
//
// Prefix paths:
//
//	apiPath := pathx.Prefix("api", "/users", "/123")
//
// # Notes
//
//   - GetProjectRootFromAppMode depends on environment variables loaded
//     via the env package; in development, call env.Load first
//   - FindProjectRoot can search for any marker files provided
//   - Join and Prefix are intended for constructing URL paths, not file system paths
package pathx
