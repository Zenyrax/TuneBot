package menu

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"tunebot/util/appdata"
	"tunebot/util/discord"

	"github.com/gookit/color"
)

func TestWebhook() {
	if appdata.CurrentConfig.Webhook != "" {
		color.Yellowln("Testing webhook...")

		webhook := &discord.Webhook{}

		embed := discord.Embed{
			Title: "Your Tunebot webhook is working!",
			Color: 3468093,
			Footer: discord.Footer{
				Text: fmt.Sprintf("Tunebot • %s", time.Now().Format(time.RFC3339)),
			},
		}

		// Turns webhook into a string
		webhook.Embeds = append(webhook.Embeds, embed)
		webhookString, _ := json.Marshal(webhook)
		payload := strings.NewReader(string(webhookString))

		// Adds ?wait=true to the end of the webhook if it's not already there
		webhookUrl := appdata.CurrentConfig.Webhook
		if !strings.Contains(appdata.CurrentConfig.Webhook, "?wait=true") {
			webhookUrl = appdata.CurrentConfig.Webhook + "?wait=true"
		}

		req, err := http.NewRequest("POST", webhookUrl, payload)

		if err != nil {
			color.Redln("Failed to send webhook:", err)
			return
		}

		req.Header.Add("content-type", "application/json")

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			color.Redln("Failed to send webhook:", err)
			return
		}

		if res.StatusCode != 200 && res.StatusCode != 204 {
			color.Redln("Failed to send webhook: Server returned error code", res.Status)
			return
		}

		color.Greenln("Successfully sent test webhook!")
	}
}
