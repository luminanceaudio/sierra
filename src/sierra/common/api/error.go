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

	UserFacingMessage string `json:"message"`
}

func (e Error) Error() string {
	return e.Err.Error()
}

func (e *Error) WithUserFacingMessage(message string) *Error {
	e.UserFacingMessage = message

	if e.Err == nil {
		e.Err = fmt.Errorf("user facing message '%s'", message)
	} else {
		e.Err = errors.Wrapf(e.Err, "user facing message '%s'", message)
	}

	return e
}

func newError(err error, userFacingMessage string, httpStatusCode int) *Error {
	if aErr, ok := err.(*Error); ok {
		return aErr
	}

	err = errors.Wrapf(err, userFacingMessage)

	message := http.StatusText(httpStatusCode)
	if err == nil {
		err = fmt.Errorf("%s", message)
	} else {
		err = errors.Wrapf(err, message)
	}

	err = fmt.Errorf("api error: %w", err)
	err = errors.WithStack(err)

	return &Error{
		Err:               err,
		httpStatusCode:    httpStatusCode,
		UserFacingMessage: userFacingMessage,
	}
}

func NewInternalError(err error, internalMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error: %s", internalMessage)
	} else {
		err = errors.Wrapf(err, "api error: %s", internalMessage)
	}

	return newError(err, "Internal error", http.StatusInternalServerError)
}

// NewUnauthorizedError will log the user out upon receiving it
func NewUnauthorizedError(err error, userFacingMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}
	return newError(err, userFacingMessage, http.StatusUnauthorized)
}

func NewForbiddenError(err error, userFacingMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}
	return newError(err, userFacingMessage, http.StatusForbidden)
}

func NewBadRequestError(err error, userFacingMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}
	return newError(err, userFacingMessage, http.StatusBadRequest)
}

func NewNotFoundError(err error, userFacingMessage string) *Error {
	if err == nil {
		err = fmt.Errorf("api error")
	}
	return newError(err, userFacingMessage, http.StatusNotFound)
}

var shutdownError = fmt.Errorf("received shutdown")

// NewShutdownError instructs the server to gracefully shutdown
func NewShutdownError() *Error {
	return &Error{
		Err:            shutdownError,
		httpStatusCode: http.StatusOK,
	}
}
