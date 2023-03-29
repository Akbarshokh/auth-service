package errs

import (
	"errors"
	"fmt"
)

type Error struct {
	err error
	msg string
}

func (e *Error) Error() string {
	if e.msg != "" {
		return fmt.Sprintf("%s; %s", e.err.Error(), e.msg)
	}

	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func Errf(err error, tmpl string, args ...interface{}) error {
	return &Error{err: err, msg: fmt.Sprintf(tmpl, args...)}
}

func New(text string) *Error {
	return &Error{err: errors.New(text), msg: ""} //nolint:goerr113
}

func Fallback(err error, fallback *Error) error {
	customErr, ok := err.(*Error) //nolint:errorlint
	if !ok {
		return Errf(fallback, err.Error())
	}

	return customErr
}