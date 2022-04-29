package httputil

import (
	"encoding/json"
	"errors"
	"net/http"

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
	return envelope{"error": envelope{"message": err.Error()}}
}

func ErrorResponse(w http.ResponseWriter, err error) error {
	switch {
	case errors.Is(err, ErrBadRequest):
		return WriteJSON(w, http.StatusBadRequest, errPayload(err), nil)

	case errors.As(err, &validator.ValidationError{}):
		return WriteJSON(w, http.StatusBadRequest, envelope{
			"error": envelope{
				"message": "validation error",
				"errors":  err.(validator.ValidationError).Errors,
			},
		}, nil)

	case errors.Is(err, ErrUnauthorized):
		return WriteJSON(w, http.StatusUnauthorized, errPayload(err), nil)

	case errors.Is(err, ErrForbidden):
		return WriteJSON(w, http.StatusForbidden, errPayload(err), nil)

	case errors.Is(err, ErrNotFound):
		return WriteJSON(w, http.StatusNotFound, errPayload(err), nil)

	case errors.Is(err, ErrConflict):
		return WriteJSON(w, http.StatusConflict, errPayload(err), nil)

	case errors.Is(err, ErrInternal):
		return WriteJSON(w, http.StatusForbidden, errPayload(err), nil)

	default:
		return WriteJSON(w, http.StatusInternalServerError, errPayload(ErrInternal), nil)
	}
}

func DataResponse(w http.ResponseWriter, dst interface{}) error {
	return WriteJSON(w, http.StatusOK, envelope{"data": dst}, nil)
}
