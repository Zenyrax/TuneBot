package tuneportals

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
	"tunebot/util/discord"
	"tunebot/util/plog"

	"net/http"
	"net/http/cookiejar"
)

// Submit order
func SubmitOrder(task *Task) {
	plog.TaskStatus(task.Count, "yellow", "INFO", "Submitting order...")

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/SecureCartCheckoutSubmit.json", task.Site), strings.NewReader(fmt.Sprintf("email=%s&password=&billing_is_shipping=on&shipping_address_1=%s&shipping_address_2=%s&shipping_first_name=%s&shipping_last_name=%s&shipping_city=%s&shipping_state=%s&shipping_country=%s&shipping_zip=%s&shipping_phone=%s&billing_address_1=%s&billing_address_2=%s&billing_first_name=%s&billing_last_name=%s&billing_city=%s&billing_state=%s&billing_country=%s&billing_zip=%s&billing_phone=%s&comments=&shipping_option_id=%s&payment_option=stripe&stripe_token=%s", task.Email, task.AddressLine1, task.AddressLine2, task.FirstName, task.LastName, task.City, task.State, task.Country, task.Zip, task.Phone, task.AddressLine1, task.AddressLine2, task.FirstName, task.LastName, task.City, task.State, task.Country, task.Zip, task.Phone, task.ShippingOption, task.paymentToken)))

	// Probably don't need this but whatever
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while submitting order, please report this if it keeps happening")
		return
	}

	// I don't actually know if they care about headers that much, just wanna keep everything as accurate as possible in case they do
	req.Header = http.Header{
		"cache-control":             {`max-age=0`},
		"content-type":              {`application/x-www-form-urlencoded; charset=UTF-8`},
		"sec-ch-ua":                 {`"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`},
		"sec-ch-ua-mobile":          {`?0`},
		"sec-ch-ua-platform":        {`"Windows"`},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"},
		"accept":                    {"application/json, text/javascript, */*; q=0.01"},
		"sec-fetch-site":            {`none`},
		"sec-fetch-mode":            {`cors`},
		"sec-fetch-user":            {`?1`},
		"sec-fetch-dest":            {`document`},
		"referer":                   {fmt.Sprintf("https://%s/SecureCartCheckout?anon_ok=1", task.Site)},
		"accept-encoding":           {"gzip, deflate, br"},
		"accept-language":           {"en-US,en;q=0.9"},
	}

	res, err := task.client.Do(req)
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Request failure while submitting order (Check proxies/internet connection)")
		fmt.Println(err)
		<-time.After(3 * time.Second)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			plog.TaskStatus(task.Count, "red", "ERROR", "Failure parsing response while submitting order")
			<-time.After(3 * time.Second)
			return
		}
		var response orderResponse
		json.Unmarshal(body, &response)
		//fmt.Println(string(body))
		//fmt.Println(response)
		if len(response.Errors) == 0 {
			plog.TaskStatus(task.Count, "green", "SUCCESS", "Successfully placed order!")
			webhook := &discord.Webhook{}

			embed := discord.Embed{
				Title: "Successfully placed order!",
				Color: 3468093,
				Footer: discord.Footer{
					Text: fmt.Sprintf("TuneBot • %s", time.Now().Format(time.RFC3339)),
				},
				Thumbnail: discord.Thumbnail{
					URL: "https://img.broadtime.com" + task.image,
				},
			}

			field := discord.Field{
				Name:  "Site",
				Value: task.Site,
			}
			embed.Fields = append(embed.Fields, field)

			field = discord.Field{
				Name:  "Item",
				Value: task.name,
			}
			embed.Fields = append(embed.Fields, field)

			field = discord.Field{
				Name:   "Item Price",
				Inline: true,
				Value:  task.price,
			}
			embed.Fields = append(embed.Fields, field)

			field = discord.Field{
				Name:   "Media Type",
				Inline: true,
				Value:  task.mediaType,
			}
			embed.Fields = append(embed.Fields, field)

			field = discord.Field{
				Name:   "Quantity",
				Inline: true,
				Value:  task.Quantity,
			}
			embed.Fields = append(embed.Fields, field)

			webhook.Embeds = append(webhook.Embeds, embed)

			discord.SendWebhook(webhook, task.Webhook)
			task.Stage = KILL
		} else {
			plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Error while submitting order %v, reinitializing session...", response.Errors))
			<-time.After(3 * time.Second)
			// The session seems to kinda freak out if the order fails so we make a new one
			jar, _ := cookiejar.New(nil)
			task.client.Jar = jar
			task.Stage = CREATE_SESSION
		}
	} else if res.StatusCode >= 500 {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Server error while submitting order (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	} else {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Request error while submitting order (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	}
}
