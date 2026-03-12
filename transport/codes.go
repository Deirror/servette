// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package transport

type Code = string

const (
	JSONFail         Code = "json.fail"
	HeadersWriteFail Code = "headers.write.fail"
	TransportFail    Code = "transport.fail"
	TemplFail        Code = "templ.fail"
	URLNotFound      Code = "url.not_found"
)
