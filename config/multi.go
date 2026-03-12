// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package config

// Represents maps of  prefixes (e.g. "NEON", "UPSTASH", "WEB", "db", "whatever_you-want") to a set of Configs.
type MultiConfig[T any] map[string]*T
