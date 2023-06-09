package http

import (
	"errors"
	"fmt"
)

type ErrorHandler interface {
	HandleError(w ResponseWriter, r *Request, err error)
}

type ErrorHandlerFunc func(w ResponseWriter, r *Request, err error)

func (fn ErrorHandlerFunc) HandleError(w ResponseWriter, r *Request, err error) {
	fn(w, r, err)
}

type Error struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *Error) Error() string {
	return fmt.Sprintf("http error: %d %s: %v", e.StatusCode, e.Message, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Is(err error) bool {
	var httpErr *Error
	if !errors.As(err, &httpErr) {
		return false
	}
	return httpErr.StatusCode == e.StatusCode
}
