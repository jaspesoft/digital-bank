package domain

const (
	TOPIC_NOTIFY_CLIENT_WEBHOOK     Topic = "NOTIFY_CLIENT_WEBHOOK"
	TOPIC_UPDATE_BALANCE_WALLET     Topic = "UPDATE_BALANCE_WALLET"
	TOPIC_NOTIFY_ERROR              Topic = "NOTIFY_ERROR"
	TOPIC_EXEC_INTERNAL_TRANSACTION Topic = "EXEC_INTERNAL_TRANSACTION"
	TOPIC_EXEC_WITHDRAW             Topic = "EXEC_WITHDRAW"
)

type (
	Topic string

	Message struct {
		Data  interface{}
		Topic Topic
	}

	EventBus interface {
		Emit(data interface{}, topic Topic) error
		Subscribe(topic Topic, callback func(Message))
	}
)
