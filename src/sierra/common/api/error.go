package api

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// This struct is returned from the API if the status code is not 400
type Error struct {
	// Internal, not returned from API
	Err            error `json:"-"`
	httpStatusCode int   `json:"-"`

	Code int
}

func (e Error) Error() string {
	return e.Err.Error()
}

type ErrorCode = int

func (e *Error) WithCode(code ErrorCode) *Error {
	e.Code = code

	if e.Err == nil {
		e.Err = fmt.Errorf("code %d", code)
	} else {
		e.Err = errors.Wrapf(e.Err, "code %d", code)
	}

	return e
}

func newError(err error, internalMessage string, httpStatusCode int) *Error {
	if aErr, ok := err.(*Error); ok {
		return aErr
	}

	err = errors.Wrapf(err, internalMessage)

	message := http.StatusText(httpStatusCode)
	if err == nil {
		err = fmt.Errorf("%s", message)
	} else {
		err = errors.Wrapf(err, message)
	}

	err = fmt.Errorf("api error: %w", err)
	err = errors.WithStack(err)

	return &Error{
		Err:            err,
		httpStatusCode: httpStatusCode,
	}
}

func NewInternalError(err error, internalMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}

	return newError(err, internalMessage, http.StatusInternalServerError)
}

// NewUnauthorizedError will log the user out upon receiving it
func NewUnauthorizedError(err error, internalMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}
	return newError(err, internalMessage, http.StatusUnauthorized)
}

func NewForbiddenError(err error, internalMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}
	return newError(err, internalMessage, http.StatusForbidden)
}

func NewBadRequestError(err error, internalMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}
	return newError(err, internalMessage, http.StatusBadRequest)
}

func NewNotFoundError(err error, internalMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}
	return newError(err, internalMessage, http.StatusNotFound)
}
