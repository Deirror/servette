// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package env provides utilities for loading, accessing, and parsing
// environment variables.
//
// It wraps standard library functions and third-party helpers to offer
// a more ergonomic and error-aware API for working with environment data,
// especially in development and configuration-heavy applications.
//
// # Features
//
//   - Load environment variables from .env files (development-friendly)
//   - Safe retrieval with explicit error handling
//   - Fallback to OS environment variables when .env is not present
//   - Helper parsers for common types (bool, int, duration)
//   - Default value support
//
// # Loading Environment Variables
//
// The Load function uses github.com/joho/godotenv to load variables
// from .env files. This is typically used in development:
//
//	err := env.Load()
//	err := env.Load(".env.local", ".env")
//
// Note: Load should be called early in application startup.
//
// # Accessing Variables
//
// Use Get for strict retrieval:
//
//	val, err := env.Get("APP_PORT")
//
// Or GetValOrDefault for fallback behavior:
//
//	port := env.GetValOrDefault("APP_PORT", "8080")
//
// # Reading All Variables
//
// GetAll attempts to read from .env files first. If none are found,
// it falls back to the OS environment:
//
//	vars, err := env.GetAll()
//
// # Parsing Helpers
//
// The package includes helpers for common types:
//
//	b, err := env.ParseBool("FEATURE_ENABLED")
//	i, err := env.ParseInt("PORT")
//	d, err := env.ParseTimeDuration("TIMEOUT")
//
// Supported boolean values include:
//
//	true, 1, yes, y
//	false, 0, no, n
//
// # Notes
//
//   - Get returns an error if a variable is not set
//   - Parsing functions depend on Get and will fail if the variable is missing
//   - GetAll prioritizes .env files but gracefully falls back to OS variables
//   - fileExists is an internal helper used to validate file presence
//
// This package is intended to standardize environment handling across
// applications and reduce repetitive boilerplate.
package env
