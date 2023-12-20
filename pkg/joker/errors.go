package joker

import (
	"fmt"
)

type (
	// Error interface represents an wrap error.
	Error interface {
		Error() string
		SetOrigin(orig error) Error
		Unwrap() error
		Code() ErrorCode
		Message() string
		SetMessage(message string) Error
	}

	// ErrorObject is the default wrap error
	// that implements the Error interface.
	ErrorObject struct {
		orig error
		msg  string
		code ErrorCode
	}
)

// WrapErrorf returns a wrapped error.
func WrapErrorf(orig error, code ErrorCode, format string, a ...interface{}) Error {
	return &ErrorObject{
		orig: orig,
		code: code,
		msg:  fmt.Sprintf(format, a...),
	}
}

// NewErrorf instantiates a new error.
func NewErrorf(code ErrorCode, format string, a ...interface{}) Error {
	return WrapErrorf(nil, code, format, a...)
}

// Error returns the message, when wrapping errors the wrapped error is returned.
func (e *ErrorObject) Error() string {
	if e.orig != nil {
		return fmt.Sprintf("%s : %v", e.msg, e.orig)
	}

	return e.msg
}

// SetError set the error's origin.
func (e *ErrorObject) SetOrigin(orig error) Error {
	e.orig = orig
	return e
}

// Unwrap returns the wrapped error, if any.
func (e *ErrorObject) Unwrap() error {
	return e.orig
}

// Code returns the code representing this error.
func (e *ErrorObject) Code() ErrorCode {
	return e.code
}

// Message return the error's message.
func (e *ErrorObject) Message() string {
	return e.msg
}

// SetMessage set the error's message.
func (e *ErrorObject) SetMessage(message string) Error {
	e.msg = message
	return e
}
