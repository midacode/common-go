package errutil

import (
	"errors"

	"github.com/midacode/common-go/httputil"
	"github.com/midacode/common-go/validator"
)

func IsExpectedError(err error) bool {
	return errors.Is(err, httputil.ErrBadRequest) || errors.Is(err, httputil.ErrUnauthorized) || errors.Is(err, httputil.ErrForbidden) || errors.Is(err, httputil.ErrNotFound) || errors.Is(err, httputil.ErrInternal) || errors.Is(err, httputil.ErrConflict) || errors.As(err, &validator.ValidationError{})
}
