package events

import accountevent "digital-bank/internal/account/infrastructure/event"

func SubscribeToEvents() {
	accountevent.AccountEventHandle()
}
