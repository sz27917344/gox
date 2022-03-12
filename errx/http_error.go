package errx

import "github.com/sz27917344/gox/resx"

type HttpError struct {
	Code    string
	Message string
	Err     error
}

func NewCode(httpCode resx.HttpCode) error {
	return &HttpError{Code: httpCode.Code, Message: httpCode.Message}
}

func NewCodeMessage(httpCode resx.HttpCode, message string) error {
	return &HttpError{Code: httpCode.Code, Message: message}
}

func NewError(httpCode resx.HttpCode, err error) error {
	return &HttpError{Code: httpCode.Code, Err: err}
}

func NewErrorMessage(httpCode resx.HttpCode, err error, message string) error {
	return &HttpError{Code: httpCode.Code, Err: err, Message: message}
}

func (e *HttpError) Error() string {
	return e.Message
}
