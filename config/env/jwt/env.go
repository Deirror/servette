package jwt

import (
	"github.com/Deirror/servette/auth/jwt"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/config"
	"github.com/Deirror/servette/env"
)

type MultiConfig = config.MultiConfig[jwt.Config]

var suffixes = []string{"JWT_SECRET", "JWT_COOKIE_NAME", "JWT_TOKEN_TTL"}

// LoadConfig loads the JWT configuration from environment variables with optional prefix:
// JWT_COOKIE_NAME, JWT_SECRET, and JWT_TOKEN_TTL.
func LoadConfig(prefix ...string) (*jwt.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	secret, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	name, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	ttl, err := env.ParseTimeDuration(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	return jwt.NewConfig(name, secret, ttl), nil
}

// LoadMultiConfig scans env vars and builds JWT configs based on their prefix.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
