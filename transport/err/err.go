// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package errx

import (
	"errors"
)

// Simple transport error struct, holding the most generic data for an error response.
// Can be used for any type of transport - htpp, grpc, unix ...
type Err struct {
	Code        string `json:"code"`              // semantic error code
	MsgKey      string `json:"message_key"`       // i18n key
	Msg         string `json:"message,omitempty"` // resolved message (optional)
	InternalMsg string `json:"-"`                 // logs only
}

// New constructs a new Err with both client-facing key and internal message.
func New(code, msgKey string, internalMsg ...string) *Err {
	msg := ""
	if len(internalMsg) > 0 {
		msg = internalMsg[0]
	}

	return &Err{
		Code:        code,
		MsgKey:      msgKey,
		InternalMsg: msg,
	}
}

// New constructs a new Err with both client-facing key and resolved message, and internal message.
func NewWithMsg(code, msgKey, msg string, internalMsg ...string) *Err {
	iMsg := ""
	if len(internalMsg) > 0 {
		iMsg = internalMsg[0]
	}

	return &Err{
		Code:        code,
		MsgKey:      msgKey,
		Msg:         msg,
		InternalMsg: iMsg,
	}
}

// Error implements the error interface and returns the MsgKey or resolved Msg.
func (e *Err) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.MsgKey
}

// FromErr attempts to convert a generic error into an *Err.
// Returns nil if the error is not of the correct type.
func FromErr(err error) *Err {
	if err == nil {
		return nil
	}
	var e *Err
	if ok := AsErr(err, &e); !ok {
		return nil
	}
	return e
}

// AsErr checks whether the error is of type *Err.
func AsErr(err error, target **Err) bool {
	return errors.As(err, target)
}
