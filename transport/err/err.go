package errx

import (
	"errors"
)

// Simple transport error struct, holding the most generic data for an error response.
// Can be used for any type of transport - htpp, grpc, unix ...
type Err struct {
	Code        string `json:"status"`      // semantic error code (e.g. "user.not_found" or whatever)
	MsgKey      string `json:"message_key"` // i18n key
	InternalMsg string `json:"-"`           // logs only
}

// New constructs a new Err with both client-facing key and internal message.
func NewWithMsgKey(code, msgKey string, internalMsg ...string) *Err {
	msg := ""
	if len(internalMsg) > 1 {
		msg = internalMsg[0]
	}

	return &Err{
		Code:        code,
		MsgKey:      msgKey,
		InternalMsg: msg,
	}
}

// New constructs a new Err with internal message key and without a i18n message for the client.
// Used in Frontend APIs, since the i18n message key is already displayed in the html template.
func New(code string, internalMsg ...string) *Err {
	msg := ""
	if len(internalMsg) > 1 {
		msg = internalMsg[0]
	}

	return &Err{
		Code:        code,
		MsgKey:      "",
		InternalMsg: msg,
	}
}

// Error implements the error interface and returns the MsgKey.
func (e *Err) Error() string {
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
