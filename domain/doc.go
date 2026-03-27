// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package domain provides preconfigured, ready-to-use service clients
// for external systems such as databases, caches, and third-party services.
//
// Each subpackage under domain represents a specific integration
// (e.g. MongoDB, Upstash, etc.) and exposes a client that can be
// easily instantiated and used without additional setup complexity.
//
// # Philosophy
//
// The goal of this package is to:
//
//   - Centralize external service integrations
//   - Provide consistent and minimal client initialization
//   - Reduce boilerplate when working with infrastructure services
//
// # Structure
//
// Each subdirectory:
//
//   - Targets a specific service or provider
//   - Exposes a client with sensible defaults
//   - Is designed to be plug-and-play within applications
//
// Example:
//
//	domain/mongo   → MongoDB client
//	domain/upstash → Upstash Redis client
//
// # Usage
//
// Import the desired subpackage and create the client:
//
//	client := mongo.New(...)
//
// # Notes
//
//   - Clients are intended to be lightweight and composable
//   - Configuration is typically injected from higher-level packages
//   - Each client is isolated, making it easy to swap implementations
//   - Documentations for the subdirectories is omitted
package domain
