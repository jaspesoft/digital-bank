package event

import (
	systemdomain "digital-bank/domain/system/domain"
	eventbus "digital-bank/pkg/event_bus"
)

func onboardingInProvider() {
	eventbus.NewAWSEventBus().Subscribe(systemdomain.TOPIC_ONBOARDING_IN_PROVIDER, func(msg eventbus.Message) {
		// code to handle the event

	})
}

func onboardingUpdateData() {
	eventbus.NewAWSEventBus().Subscribe(systemdomain.TOPIC_ONBOARDING_UPDATE_DATA, func(msg eventbus.Message) {
		// code to handle the event

	})
}

func AccountEventHandle() {

	onboardingInProvider()

	onboardingUpdateData()
}
