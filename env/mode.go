// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package env

import "fmt"

type Mode string

const (
	Dev     Mode = "dev"
	Prod    Mode = "prod"
	Staging Mode = "staging"
)

func (m Mode) IsDev() bool {
	return m == Dev
}

func (m Mode) IsProd() bool {
	return m == Prod
}

func (m Mode) IsStaging() bool {
	return m == Staging
}

func ParseMode(v string) (Mode, error) {
	switch v {
	case string(Dev):
		return Dev, nil
	case string(Prod):
		return Prod, nil
	case string(Staging):
		return Staging, nil
	default:
		return "", fmt.Errorf("invalid mode: %s", v)
	}
}
