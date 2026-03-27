// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package str

import (
	"database/sql"
	"errors"
	"fmt"
)

func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func ContainsRanges(subSlices [][]string, slices [][]string) error {
	if len(slices) != len(subSlices) || len(slices) == 0 || len(subSlices) == 0 {
		return errors.New("invalid lens passed of slices")
	}

	for idx, slice := range slices {
		for _, key := range subSlices[idx] {
			if !Contains(slice, key) {
				return fmt.Errorf("key not present in slice, with idx %d: %s", idx, key)
			}
		}
	}

	return nil
}

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}
