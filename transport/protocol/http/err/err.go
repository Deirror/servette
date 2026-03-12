// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package errx

import (
	"fmt"
	"strconv"

	errx "github.com/Deirror/servette/transport/err"
)

func New(status int, key, errMsg string, srvErr error) *errx.Err {
	srvErrMsg := errMsg
	if srvErr != nil {
		srvErrMsg = fmt.Sprintf("%s: %v", errMsg, srvErr)
	}

	return errx.New(strconv.Itoa(status), key, srvErrMsg)
}
