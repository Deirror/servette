package domain

import (
	"fmt"
	"strings"

	"github.com/Deirror/servette/config"
)

// Contains multiple providers of given domain.
type MultiDomain[P any] = map[string]P

func NewMultiDomain[T any, P any](keys []string, constructor func(cfg *T) (P, error), cfgs config.MultiConfig[T]) (MultiDomain[P], error) {
	m := make(MultiDomain[P])
	for _, key := range keys {
		val, ok := cfgs[key]
		if !ok {
			return nil, fmt.Errorf("unsupported key passed: %s", key)
		}

		provider, err := constructor(val)
		if err != nil {
			return nil, err
		}

		key = strings.ToLower(key)
		m[key] = provider
	}
	return m, nil
}
