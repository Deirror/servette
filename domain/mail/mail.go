// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package mail

type Sender interface {
	Send(to, subject, body string) error
}
