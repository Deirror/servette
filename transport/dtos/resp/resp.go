// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package respx

type Resp struct {
	Code    string `json:"code"`
	MsgKey  string `json:"message_key"`
	Payload any    `json:"payload,omitempty"`
}

func New(code, key string, payload any) *Resp {
	return &Resp{
		Code:    code,
		MsgKey:  key,
		Payload: payload,
	}
}
