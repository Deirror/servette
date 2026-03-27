// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package kv

type Storer interface {
	Set(key string, value string, ttlSeconds int) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) (int64, error)
}
