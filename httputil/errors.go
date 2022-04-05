package httputil

import (
	"errors"
)

var (
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrNotFound     = errors.New("not found")
	ErrInternal     = errors.New("internal server error")
)

func IsUserError(err error) bool {
	return errors.Is(err, ErrBadRequest) && errors.Is(err, ErrUnauthorized) && errors.Is(err, ErrForbidden) && errors.Is(err, ErrNotFound) && errors.Is(err, ErrInternal)
}
