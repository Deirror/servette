// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package crypto provides simple utilities for secure password hashing
// and verification using the bcrypt algorithm.
//
// It is designed for straightforward password handling in applications,
// abstracting the underlying bcrypt implementation while maintaining
// strong security practices such as salting and adaptive hashing.
//
// # Features
//
//   - Secure password hashing using bcrypt
//   - Password verification against stored hashes
//   - Automatic salting and cost handling via bcrypt
//
// # Usage
//
// Hash a password:
//
//	hash, err := crypto.HashPassword("my-secret-password")
//
// Verify a password:
//
//	err := crypto.CheckPasswordHash(hash, "my-secret-password")
//	if err != nil {
//	    // password is invalid
//	}
//
// # Notes
//
//   - bcrypt automatically handles salting; no manual salt management is required
//   - The default cost (bcrypt.DefaultCost) is used, balancing security and performance
//   - CheckPasswordHash returns an error if the password does not match
//   - Store only the hashed password, never the plaintext password
//
// # Security Considerations
//
//   - Always use HTTPS when transmitting passwords
//   - Consider increasing the bcrypt cost for higher security if performance allows
//   - Do not log or expose password hashes unnecessarily
package crypto
