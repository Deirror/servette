package db

import (
	"time"

	"github.com/Deirror/servette/env"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/config"
	"github.com/Deirror/servette/domain/db"
)

type MultiConfig = config.MultiConfig[db.Config]

var suffixes = []string{
	"DB_DRIVER",
	"DB_DSN",
	"DB_POOL_SIZE",
	"DB_MAX_IDLE",
	"DB_MAX_LIFETIME",
}

// LoadConfig loads the database configuration from environment variables.
// Required variables: DB_DSN, DB_POOL_SIZE, DB_MAX_IDLE, DB_MAX_LIFETIME.
func LoadConfig(prefix ...string) (*db.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	driver, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	dsn, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	size, err := env.ParseInt(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	maxIdle, err := env.ParseInt(pfx + suffixes[3])
	if err != nil {
		return nil, err
	}

	maxLT, err := env.ParseTimeDuration(pfx + suffixes[4])
	if err != nil {
		return nil, err
	}

	return db.NewConfig(driver, dsn, uint8(size), uint8(maxIdle), maxLT*time.Second), nil
}

// LoadMultiConfig loads multiple DB Config instances by scanning environment variables
// with the suffixes keys and optional prefixes.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
