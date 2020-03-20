package notifications

import (
	"fmt"
	"github.com/slack-go/slack"
	"glastorumornotifier/models"
)

const (
	webhookurl = `https://hooks.slack.com/services/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`
)

type SlackNotificationHandler struct{
	Config SlackConfig
}

type SlackConfig struct {
	Channel string
}

func (s SlackNotificationHandler) SendNewActs(acts models.Acts) error {

		slackMsg := &slack.WebhookMessage{
			Channel: "dptestchannel",
			Text: "test",
		}

		//if err := slack.PostWebhook(webhookurl, slackMsg); err != nil {
		//	return err
		//}

		fmt.Println(slackMsg§)

		return nil
}