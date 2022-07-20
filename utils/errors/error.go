package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	name string
	desc string
	err  error
}

func New(name string) *Error {
	return &Error{name: name}
}

func (e *Error) Error() string {
	desc := "[" + e.name + "]: " + e.desc
	if e.err != nil {
		desc += "(" + e.err.Error() + ")"
	}
	return desc
}

func (e *Error) Is(err error) bool {
	return e == err
}

func (e *Error) Desc(desc string) *Error {
	return &Error{name: e.name, desc: desc}
}

func (e *Error) Err(err error, desc string) *Error {
	return &Error{name: e.name, desc: desc, err: err}
}

func (e *Error) Errf(err error, format string, a ...interface{}) *Error {
	return e.Err(err, fmt.Sprintf(format, a...))
}

func (e *Error) Code(desc string, err error, code codes.Code) *status.Status {
	return status.New(code, e.Err(err, desc).Error())
}

func (e *Error) Unwrap() error {
	return e.err
}
