// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package translation

type I18n struct {
	data map[string]string
}

func (i *I18n) T(key string) string {
	if v, ok := i.data[key]; ok {
		return v
	}
	return key
}
