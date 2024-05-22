package errorhandle

import (
	"errors"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type APIError interface {
	APIError() (int, any)
}

type errorAPIError struct {
	status int
	err    error
}

func (e errorAPIError) Error() string {
	return e.err.Error()
}

func (e errorAPIError) APIError() (int, any) {
	return e.status, Response{
		Message: e.Error(),
	}
}

func Handle(err error) (int, any) {
	var apiErr APIError
	if errors.As(err, &apiErr) {
		return apiErr.APIError()
	}

	return http.StatusInternalServerError, Response{
		Message: err.Error(),
	}
}

var (
	ErrNotFound            = &errorAPIError{status: http.StatusNotFound, err: errors.New("can not found zipcode")}
	ErrUnprocessableEntity = &errorAPIError{status: http.StatusUnprocessableEntity, err: errors.New("invalid zipcode")}
)
