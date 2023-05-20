package tuneportals

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
	"tunebot/util/plog"

	"net/http"
)

// Loading shipping info
func LoadShipping(task *Task) {
	plog.TaskStatus(task.Count, "yellow", "INFO", "Getting shipping data...")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/SecureCartShippingOptions.json?total=%s&shipping_country=%s", task.Site, task.total, task.Country), nil)

	// Probably don't need this but whatever
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while getting shipping data, please report this if it keeps happening")
		return
	}

	// I don't actually know if they care about headers that much, just wanna keep everything as accurate as possible in case they do
	req.Header = http.Header{
		"cache-control":             {`max-age=0`},
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
		plog.TaskStatus(task.Count, "red", "ERROR", "Request failure while getting shipping data (Check proxies/internet connection)")
		fmt.Println(err)
		<-time.After(3 * time.Second)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			plog.TaskStatus(task.Count, "red", "ERROR", "Failure parsing response while getting shipping data")
			<-time.After(3 * time.Second)
			return
		}
		var response shippingResponse
		json.Unmarshal(body, &response)
		// fmt.Println(string(body))
		if len(response.ShippingOptions) > 0 {
			for i := 0; i < len(response.ShippingOptions); i++ {
				if strings.TrimSpace(task.ShippingOption) == strings.TrimSpace(response.ShippingOptions[i].DeliveryType) {
					task.ShippingOption = fmt.Sprint(response.ShippingOptions[i].ID)
					fmt.Println(response.ShippingOptions[i], task.ShippingOption)
					plog.TaskStatus(task.Count, "blue", "INFO", fmt.Sprintf("Successfully loaded shipping data (Using %s)", task.ShippingOption))
					task.Stage = TOKENIZE_PAYMENT
					return
				}
			}
			plog.TaskStatus(task.Count, "blue", "INFO", fmt.Sprintf("Successfully loaded shipping data (Using %s)", response.ShippingOptions[0].DeliveryType))
			task.ShippingOption = fmt.Sprint(response.ShippingOptions[0].ID)
			task.Stage = TOKENIZE_PAYMENT
		} else {
			plog.TaskStatus(task.Count, "red", "ERROR", "No shipping options available")
			<-time.After(3 * time.Second)
		}
	} else if res.StatusCode >= 500 {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Server error while getting shipping data (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	} else {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Request error while getting shipping data (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	}
}
