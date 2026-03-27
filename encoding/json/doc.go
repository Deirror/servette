// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package json provides convenience helpers for encoding, decoding,
// and working with JSON in Go applications.
//
// It wraps the standard library encoding/json package with a simpler,
// more ergonomic API, especially for HTTP use cases and generic types.
//
// # Features
//
//   - Encode Go values to JSON
//   - Decode JSON into existing or new typed values
//   - HTTP response helpers for writing JSON
//   - Generic helpers using Go type parameters
//
// # Encoding
//
// Encode writes a Go value as JSON to an io.Writer:
//
//	err := json.Encode(w, data)
//
// Marshal returns JSON as a byte slice:
//
//	b, err := json.Marshal(data)
//
// # Decoding
//
// DecodeInto decodes JSON into an existing variable:
//
//	var v MyStruct
//	err := json.DecodeInto(r, &v)
//
// Decode creates and returns a new typed value:
//
//	v, err := json.Decode[MyStruct](r)
//
// Unmarshal parses JSON from bytes into a typed value:
//
//	v, err := json.Unmarshal[MyStruct](data)
//
// # HTTP Helpers
//
// Write simplifies sending JSON responses in HTTP handlers:
//
//	err := json.Write(w, http.StatusOK, data)
//
// It automatically sets the Content-Type header to application/json.
//
// BodyToJSON reads and decodes an HTTP response body:
//
//	v, err := json.BodyToJSON[MyStruct](resp)
//
// # Notes
//
//   - All decoding functions return errors for malformed JSON
//   - BodyToJSON automatically closes the response body
//   - Write sets headers before encoding, so it should be called once per response
//   - The package is designed to reduce boilerplate when working with JSON in APIs
//
// This package is intended to standardize JSON handling across services
// while keeping full compatibility with encoding/json.
package json
