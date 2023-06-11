package menu

import (
	"fmt"
	"strings"
	"time"

	"tunebot/util/appdata"

	"github.com/AlecAivazis/survey/v2"
)

func Settings() {
	Reset()

	option := ""
	prompt := &survey.Select{
		Message: "What do you want to do?",
		Options: []string{fmt.Sprintf("Set webhook - %s", appdata.CurrentConfig.Webhook), "Test webhook", "Return"},
	}
	survey.AskOne(prompt, &option)

	switch {
	case strings.Contains(option, "Set webhook"):
		webhook := ""
		prompt := &survey.Input{
			Message: "What's your webhook URL? (If you can't paste, right-click the window)",
		}
		survey.AskOne(prompt, &webhook)

		if webhook != "" {
			appdata.CurrentConfig.Webhook = webhook
			appdata.WriteConfig()
		}

		Settings()
	case option == "Test webhook":
		Reset()
		TestWebhook()
		time.Sleep(3 * time.Second)
		Settings()
	default:
		Start()
	}
}
