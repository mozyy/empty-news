package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	name string
	desc string
	err  error
}

func New(name string) Error {
	return Error{name: name}
}

func (e Error) Error() string {
	desc := "[" + e.name + "]: " + e.desc
	if e.err != nil {
		desc += "(" + e.err.Error() + ")"
	}
	return desc
}

func (e Error) Desc(desc string) Error {
	return Error{name: e.name, desc: desc}
}

func (e Error) Err(desc string, err error) Error {
	return Error{name: e.name, desc: desc, err: err}
}
func (e Error) Code(desc string, err error, code codes.Code) *status.Status {
	return status.New(code, e.Err(desc, err).Error())
}

func (e Error) Unwrap() error {
	return e.err
}
