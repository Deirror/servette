// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package translation

import (
	"github.com/Deirror/servette/translation/languages"
)

// Wrapper struct, containing all needed data for translation.
type Translator struct {
	Bundle   *Bundle
	Resolver *languages.Resolver
}

func New(b *Bundle, r *languages.Resolver) *Translator {
	return &Translator{
		Bundle:   b,
		Resolver: r,
	}
}

func Emplace(bundlePaths []string, defaultLang string, supportedLangs ...string) (*Translator, error) {
	rlv := languages.NewResolver(defaultLang, supportedLangs...)
	bundle := NewBundle()
	for _, p := range bundlePaths {
		b, err := LoadBundle(p)
		if err != nil {
			return nil, err
		}
		bundle.Merge(b)
	}
	return New(bundle, rlv), nil
}
