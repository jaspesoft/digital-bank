package systemdomain

type (
	Error struct {
		httpCode int
		message  string
	}
)

func NewError(httpCode int, message string) *Error {
	return &Error{
		httpCode: httpCode,
		message:  message,
	}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) GetHTTPCode() int {
	return e.httpCode
}
