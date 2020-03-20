package notifications

import (
	"glastorumornotifier/models"
	"testing"
)

func TestSlackNotification(t *testing.T) {

	slackConfig := SlackConfig{
		Channel: "dptestchannel",
	}
	notifier := SlackNotificationHandler{
		slackConfig,
	}

	acts := models.Acts{"a","b","c"}

	if err := notifier.SendNewActs(acts); err != nil {
		t.Fatal(err)
	}
}
