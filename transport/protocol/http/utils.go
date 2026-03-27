// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package httpx

func IsSuccessStatus(status int) bool {
	return status/100 == 2
}
