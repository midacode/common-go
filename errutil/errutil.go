package errutil

import (
	"errors"
	"fmt"

	"github.com/midacode/common-go/httputil"
	"github.com/midacode/common-go/validator"
)

func IsExpectedError(err error) bool {
	return errors.Is(err, httputil.ErrBadRequest) || errors.Is(err, httputil.ErrUnauthorized) || errors.Is(err, httputil.ErrForbidden) || errors.Is(err, httputil.ErrNotFound) || errors.Is(err, httputil.ErrInternal) || errors.Is(err, httputil.ErrConflict) || errors.As(err, &validator.ValidationError{})
}

func WrapFnError(err *error, format string, args ...interface{}) {
	if *err != nil {
		s := fmt.Sprintf(format, args...)
		*err = fmt.Errorf("%s: %w", s, *err)
	}
}
