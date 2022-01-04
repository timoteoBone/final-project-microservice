package errors

import (
	"fmt"
)

type ErrNotFound struct {
	Err error
}

type ErrBadRequest struct {
	Err error
}

type ErrInternal struct {
	Err error
}

func (err *ErrNotFound) Error() string {
	return fmt.Sprint(err.Err)
}

func (err *ErrBadRequest) Error() string {
	return fmt.Sprint(err.Err)
}

func (err *ErrInternal) Error() string {
	return fmt.Sprint(err.Err)
}
