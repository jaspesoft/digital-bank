package domain

type (
	Result[T any] struct {
		Value T
		Err   *ErrorMessage
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

func NewResult[T any](value T, err *ErrorMessage) *Result[T] {
	return &Result[T]{
		Value: value,
		Err:   err,
	}
}

func (r *Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r *Result[T]) GetValue() T {
	return r.Value
}

func (r *Result[T]) GetError() *ErrorMessage {
	if r.Err == nil {
		return nil
	}

	return r.Err
}

func (r *Result[T]) SetNotifyError(n NotifyError) {
	go func() {
		n.SetMessage(Notification{
			Channel: "error",
			Message: r.Err.Error(),
		})
		_ = n.Send()
	}()
}
