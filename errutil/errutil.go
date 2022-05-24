package errutil

import (
	"errors"
	"fmt"

	"github.com/midacode/common-go/validator"
)

func IsExpectedError(err error) bool {
	return errors.Is(err, ErrBadRequest) || errors.Is(err, ErrUnauthorized) || errors.Is(err, ErrForbidden) || errors.Is(err, ErrNotFound) || errors.Is(err, ErrInternal) || errors.Is(err, ErrConflict) || errors.As(err, &validator.ValidationError{})
}

func WrapFnError(err *error, format string, args ...interface{}) {
	if *err != nil && !IsExpectedError(*err) {
		s := fmt.Sprintf(format, args...)
		*err = fmt.Errorf("%s: %w", s, *err)
	}
}
