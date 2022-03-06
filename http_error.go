package gox

type HttpError struct {
	Code    string
	Message string
	Err     error
}

func NewCode(httpCode HttpCode) error {
	return &HttpError{Code: httpCode.Code, Message: httpCode.Message}
}

func NewCodeMessage(httpCode HttpCode, message string) error {
	return &HttpError{Code: httpCode.Code, Message: message}
}

func NewError(httpCode HttpCode, err error) error {
	return &HttpError{Code: httpCode.Code, Err: err}
}

func NewErrorMessage(httpCode HttpCode, err error, message string) error {
	return &HttpError{Code: httpCode.Code, Err: err, Message: message}
}

func (e *HttpError) Error() string {
	return e.Message
}
