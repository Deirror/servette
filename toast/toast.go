// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package toast

import (
	"net/http"

	"github.com/Deirror/servette/encoding/json"
	"github.com/Deirror/servette/transport/protocol/http/header"
)

type payload struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func trigger(w http.ResponseWriter, msg string, t string) {
	data := map[string]payload{
		"showToast": {
			Message: msg,
			Type:    t,
		},
	}

	b, _ := json.Marshal(data)

	w.Header().Set(header.HXTrigger, string(b))
}

func Success(w http.ResponseWriter, msg string) {
	trigger(w, msg, "success")
}

func Error(w http.ResponseWriter, msg string) {
	trigger(w, msg, "error")
}

func Info(w http.ResponseWriter, msg string) {
	trigger(w, msg, "info")
}

func Warning(w http.ResponseWriter, msg string) {
	trigger(w, msg, "warning")
}
