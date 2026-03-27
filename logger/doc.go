// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package logger provides a small wrapper around Go's slog package
// with opinionated defaults and helper utilities for structured logging.
//
// The package is designed to:
//   - Initialize a logger based on application environment (development or production)
//   - Standardize logging of function calls
//   - Attach contextual metadata such as request IDs
//   - Provide optional timing measurements for function execution
//
// # Logger Initialization
//
// Use New to create a logger configured for the current environment:
//
//   - Development mode uses a human-readable text format with debug-level logs
//
//   - Production mode uses structured JSON output with info-level logs
//
//     logger := logger.New(mode)
//
// # Function Logging
//
// The package provides helpers to log function execution consistently:
//
//   - LogFunc logs a function call without timing
//   - LogFuncWithTiming logs a function call including execution duration
//
// Both functions:
//   - Extract the request ID from the context (if present)
//   - Include error information automatically
//   - Allow attaching additional structured attributes
//   - Adjust log level based on the presence of an error
//
// Example:
//
//	func handler(ctx context.Context, logger *slog.Logger) error {
//	    start := time.Now()
//	    err := doWork()
//	    logger.LogFuncWithTiming(ctx, logger, "handler", start, err)
//	    return err
//	}
//
// # Context Integration
//
// The logger expects a request ID to be stored in the context using
// the transport.ReqID key. If present, it will be included in all logs.
//
// Notes
//
//   - Errors are logged under a consistent "error" attribute
//   - Duration is recorded using slog.Duration when timing is enabled
//   - Additional attributes can be passed for richer structured logs
package logger
