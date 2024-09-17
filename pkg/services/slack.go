package services

import (
	"digital-bank/internal/system/domain"
	"github.com/slack-go/slack"
	"os"
)

type (
	Slack struct {
		notification domain.Notification
	}
)

func (s *Slack) SetMessage(n domain.Notification) {
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
