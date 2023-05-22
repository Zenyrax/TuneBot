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

// Create stripe payment token
func TokenizePayment(task *Task) {
	plog.TaskStatus(task.Count, "yellow", "INFO", "Creating payment token...")

	req, err := http.NewRequest("POST", "https://api.stripe.com/v1/tokens", strings.NewReader(fmt.Sprintf("card[number]=%s&card[cvc]=%s&card[exp_month]=%s&card[exp_year]=%s&card[address_zip]=%s&key=%s", task.CardNumber, task.CVC, task.CardMonth, strings.ReplaceAll(task.CardYear, "20", ""), task.Zip, task.stripeKey)))

	// Probably don't need this but whatever
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while creating payment token, please report this if it keeps happening")
		return
	}

	// I don't actually know if they care about headers that much, just wanna keep everything as accurate as possible in case they do
	req.Header = http.Header{
		"cache-control":             {`max-age=0`},
		"content-type":              {`application/x-www-form-urlencoded`},
		"sec-ch-ua":                 {`"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`},
		"sec-ch-ua-mobile":          {`?0`},
		"sec-ch-ua-platform":        {`"Windows"`},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"},
		"accept":                    {"application/json"},
		"sec-fetch-site":            {`none`},
		"sec-fetch-mode":            {`cors`},
		"sec-fetch-user":            {`?1`},
		"sec-fetch-dest":            {`document`},
		"referer":                   {"https://js.stripe.com/"},
		"accept-encoding":           {"gzip, deflate, br"},
		"accept-language":           {"en-US,en;q=0.9"},
	}

	res, err := task.client.Do(req)
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Request failure while creating payment token (Check proxies/internet connection)")
		fmt.Println(err)
		<-time.After(3 * time.Second)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			plog.TaskStatus(task.Count, "red", "ERROR", "Failure parsing response while creating payment token")
			<-time.After(3 * time.Second)
			return
		}
		var response stripeResponse
		json.Unmarshal(body, &response)
		// fmt.Println(string(body))
		// fmt.Println(response)
		task.paymentToken = response.ID
		task.Stage = SUBMIT_ORDER
		plog.TaskStatus(task.Count, "blue", "INFO", "Successfully created payment token")
	} else if res.StatusCode >= 500 {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Server error while creating payment token (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	} else {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			plog.TaskStatus(task.Count, "red", "ERROR", "Failure parsing error while creating payment token")
			<-time.After(3 * time.Second)
			return
		}
		var response stripeError
		json.Unmarshal(body, &response)
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Request error while creating payment token [%s] (%d)", response.Error.Message, res.StatusCode))
		<-time.After(3 * time.Second)
	}
}
