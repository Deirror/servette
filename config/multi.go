package config

// Represents maps of  prefixes (e.g. "NEON", "UPSTASH", "WEB", "db", "whatever_you-want") to a set of Configs.
type MultiConfig[T any] map[string]*T
