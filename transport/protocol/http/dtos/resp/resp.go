// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package respx

import (
	"strconv"

	respx "github.com/Deirror/servette/transport/dtos/resp"
)

func New(status int, key string, payload any) *respx.Resp {
	return respx.New(strconv.Itoa(status), key, payload)
}
