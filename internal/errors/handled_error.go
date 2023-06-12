package errors

const defaultMessage = "There is something wrong with the server."

type HandledError struct {
	HttpStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
}

func NewHandledError(status int, code, message, detail string) HandledError {
	return HandledError{status, code, message, detail}
}

func NewDefaultHandledError(status int, detail string) HandledError {
	return HandledError{status, "COM0000", defaultMessage, detail}
}

func (e HandledError) Error() string {
	return e.Detail
}
