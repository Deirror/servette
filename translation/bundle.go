// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package translation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Deirror/servette/encoding/json"
)

// Bundle consists of key-value pairs of a language and translations.
type Bundle struct {
	langs map[string]map[string]string
}

func NewBundle() *Bundle {
	return &Bundle{
		langs: make(map[string]map[string]string),
	}
}

// LoadBundle loads all .json files.
func LoadBundle(dir string) (*Bundle, error) {
	b := NewBundle()

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}

		lang := strings.TrimSuffix(e.Name(), ".json")

		f, err := os.Open(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}

		raw, err := json.Decode[map[string]any](f)
		if err != nil {
			f.Close()
			return nil, err
		}
		f.Close()

		flat := make(map[string]string)
		flatten("", raw, flat)

		b.langs[lang] = flat
	}

	return b, nil
}

// Recursive flattening - leads to <>.<> ...
func flatten(prefix string, in map[string]any, out map[string]string) {
	for k, v := range in {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}

		switch val := v.(type) {
		case string:
			out[key] = val
		case map[string]any:
			flatten(key, val, out)
		default:
			out[key] = fmt.Sprintf("%v", val)
		}
	}
}

func (b *Bundle) ForLang(lang string) *I18n {
	return &I18n{
		data: b.langs[lang],
	}
}
