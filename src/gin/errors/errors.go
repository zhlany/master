package errors

import (
	"errors"
	"fmt"
)

func NewError(arg ...any) error {
	return errors.New(fmt.Sprint(arg))
}

func NewErrorf(format string, args ...any) error {
	return errors.New(fmt.Sprintf(format, args))
}
