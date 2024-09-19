package systemdomain

const (
	TOPIC_UPDATE_BALANCE_WALLET     Topic = "UPDATE_BALANCE_WALLET"
	TOPIC_EXEC_INTERNAL_TRANSACTION Topic = "EXEC_INTERNAL_TRANSACTION"
	TOPIC_EXEC_WITHDRAW             Topic = "EXEC_WITHDRAW"
	TOPIC_ONBOARDING_IN_PROVIDER    Topic = "ONBOARDING_IN_PROVIDER"
	TOPIC_ONBOARDING_UPDATE_DATA    Topic = "ONBOARDING_UPDATE_DATA"
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
