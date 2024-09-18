package services

import (
	systemdomain "digital-bank/domain/system/domain"
	"github.com/slack-go/slack"
	"os"
)

type (
	Slack struct {
		notification systemdomain.Notification
	}
)

func (s *Slack) SetMessage(n systemdomain.Notification) {
	s.notification = n
}

func (s *Slack) Send() error {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	_, _, err := api.PostMessage(s.notification.Channel, slack.MsgOptionText(s.notification.Message, false))

	if err != nil {
		return err
	}

	return nil

}
