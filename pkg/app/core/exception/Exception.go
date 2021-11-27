package exception

import (
	"net/http"
)

type Exception struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func BadRequestException(message string, err error) *Exception {

	exception := &Exception{
		Message: message,
		Code:    http.StatusBadRequest,
	}

	if err != nil {
		exception.Error = err.Error()
	}

	return exception
}

func InternalServerErrorException(message string, err error) *Exception {

	exception := &Exception{
		Message: message,
		Code:    http.StatusInternalServerError,
	}

	if err != nil {
		exception.Error = err.Error()
	}

	return exception
}

func NotFoundException(message string, err error) *Exception {

	exception := &Exception{
		Message: message,
		Code:    http.StatusNotFound,
	}

	if err != nil {
		exception.Error = err.Error()
	}

	return exception
}
