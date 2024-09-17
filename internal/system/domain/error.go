package systemdomain

type (
	ErrorMessage struct {
		HttpCode int
		Message  string
	}

	Error interface {
		Error() string
		GetHTTPCode() int
	}
)

func (e *ErrorMessage) Error() string {
	return e.Message
}

func (e *ErrorMessage) GetHTTPCode() int {
	return e.HttpCode
}
