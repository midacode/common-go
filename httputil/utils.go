package httputil

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/midacode/common-go/errutil"
	"github.com/midacode/common-go/validator"
)

func ReadJSON(r *http.Request, dst interface{}) error {
	// Decode the body into the target destination.
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, dst interface{}, headers http.Header) error {
	js, err := json.Marshal(dst)
	if err != nil {
		return err
	}
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

type envelope map[string]interface{}

func errPayload(err error) envelope {

	uErr := errors.Unwrap(err)
	if uErr != nil {
		err = uErr
	}

	return envelope{"error": envelope{"message": err.Error()}}
}

func ErrorResponse(w http.ResponseWriter, err error) error {
	switch {
	case errors.Is(err, errutil.ErrBadRequest):
		return WriteJSON(w, http.StatusBadRequest, errPayload(err), nil)

	case errors.As(err, &validator.ValidationError{}):
		return WriteJSON(w, http.StatusBadRequest, envelope{
			"error": envelope{
				"message": "validation error",
				"errors":  err.(validator.ValidationError).Errors,
			},
		}, nil)

	case errors.Is(err, errutil.ErrUnauthorized):
		return WriteJSON(w, http.StatusUnauthorized, errPayload(err), nil)

	case errors.Is(err, errutil.ErrForbidden):
		return WriteJSON(w, http.StatusForbidden, errPayload(err), nil)

	case errors.Is(err, errutil.ErrNotFound):
		return WriteJSON(w, http.StatusNotFound, errPayload(err), nil)

	case errors.Is(err, errutil.ErrConflict):
		return WriteJSON(w, http.StatusConflict, errPayload(err), nil)

	case errors.Is(err, errutil.ErrInternal):
		return WriteJSON(w, http.StatusForbidden, errPayload(err), nil)

	default:
		return WriteJSON(w, http.StatusInternalServerError, errPayload(errutil.ErrInternal), nil)
	}
}

func DataResponse(w http.ResponseWriter, dst interface{}) error {
	return WriteJSON(w, http.StatusOK, envelope{"data": dst}, nil)
}
