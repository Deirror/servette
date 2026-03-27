// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package transport

import "fmt"

type Type string

const (
	TCP Type = "TCP"
	UDS Type = "UDS"
)

func (t Type) IsTCP() bool {
	return t == TCP
}

func (t Type) IsUDS() bool {
	return t == UDS
}

func ParseType(v string) (Type, error) {
	switch v {
	case string(TCP):
		return TCP, nil
	case string(UDS):
		return UDS, nil
	default:
		return "", fmt.Errorf("invalid type: %s", v)
	}
}
