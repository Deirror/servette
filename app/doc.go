// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package app provides a minimal application runner abstraction for managing
// the lifecycle of multiple concurrent components.
// It allows you to register multiple Runner implementations, start them
// concurrently, and gracefully shut them down when the context is canceled
// or when any runner returns an error.
package app
