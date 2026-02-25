package doc

import (
	"github.com/Deirror/servette/config"
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/domain/doc"
	"github.com/Deirror/servette/env"
)

type MultiConfig = config.MultiConfig[doc.Config]

var suffixes = []string{
	"DOC_STORE_URL",
	"DOC_DATABASE",
}

// LoadConfig loads the document store configuration from environment variables,
// supporting an optional prefix.
func LoadConfig(prefix ...string) (*doc.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	url, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	db, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	return doc.NewConfig(url, db), nil
}

// LoadMultiConfig loads multiple Config instances by scanning env vars with suffixes.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
