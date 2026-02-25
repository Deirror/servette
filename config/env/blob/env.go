package blob

import (
	envcfg "github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/config"
	"github.com/Deirror/servette/domain/blob"
	"github.com/Deirror/servette/env"
)

type MultiConfig = config.MultiConfig[blob.Config]

var suffixes = []string{"BLOB_PROJECT_URL", "BLOB_API_KEY", "BLOB_BUCKET"}

// LoadConfig loads Blob Config from env vars with optional prefix.
func LoadConfig(prefix ...string) (*blob.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	url, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	apiKey, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	bucket, err := env.Get(pfx + suffixes[2])
	if err != nil {
		return nil, err
	}

	return blob.NewConfig(url, apiKey, bucket), nil
}

// LoadMultiConfig loads multiple Blob Configs by scanning env vars with blob suffixes.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
