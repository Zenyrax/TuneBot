package tuneportals

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
	"tunebot/util/plog"

	"net/http"
)

// Loading cart info
func LoadCart(task *Task) {
	plog.TaskStatus(task.Count, "yellow", "INFO", "Getting cart data...")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/SecureCartTotals.json?ig=1&state=%s&city=%s", task.Site, task.State, task.City), nil)

	// Probably don't need this but whatever
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while getting cart data, please report this if it keeps happening")
		return
	}

	// I don't actually know if they care about headers that much, just wanna keep everything as accurate as possible in case they do
	req.Header = http.Header{
		"cache-control":             {`max-age=0`},
		"sec-ch-ua":                 {task.secUa},
		"sec-ch-ua-mobile":          {`?0`},
		"sec-ch-ua-platform":        {`"Windows"`},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {task.userAgent},
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
		plog.TaskStatus(task.Count, "red", "ERROR", "Request failure while getting cart data (Check proxies/internet connection)")
		fmt.Println(err)
		<-time.After(3 * time.Second)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			plog.TaskStatus(task.Count, "red", "ERROR", "Failure parsing response while getting cart data")
			<-time.After(3 * time.Second)
			return
		}
		var response cartResponse
		json.Unmarshal(body, &response)
		// fmt.Println(string(body))
		task.total = response.Prices.Total
		// fmt.Println(task.total)
		task.Stage = LOAD_SHIPPING
		plog.TaskStatus(task.Count, "blue", "INFO", "Successfully loaded cart data")
	} else if res.StatusCode >= 500 {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Server error while getting cart data (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	} else {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Request error while getting cart data (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	}
}
