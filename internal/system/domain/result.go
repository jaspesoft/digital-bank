package systemdomain

type (
	Result[T any] struct {
		value T
		err   *Error
	}

	Notification struct {
		Channel string `json:"channel"`
		Message string `json:"message"`
	}

	NotifyError interface {
		SetMessage(n Notification)
		Send() error
	}
)

func NewResult[T any](value T, err *Error) Result[T] {
	return Result[T]{
		value: value,
		err:   err,
	}
}

func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

func (r *Result[T]) GetValue() T {
	return r.value
}

func (r *Result[T]) SetNotifyError(n NotifyError) {
	go func() {
		n.SetMessage(Notification{
			Channel: "error",
			Message: r.err.Error(),
		})
		_ = n.Send()
	}()
}

func (r *Result[T]) GetError() *Error {
	return r.err
}
