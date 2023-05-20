package tuneportals

import (
	"fmt"
	"time"
	"tunebot/util/plog"

	"net/http"
)

// Getting session cookies
func CreateSession(task *Task) {
	plog.TaskStatus(task.Count, "yellow", "INFO", "Initializing session...")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/SecureUser.json?_=%d", task.Site, time.Now().UTC().UnixNano()/1e6), nil)

	// Probably don't need this but whatever
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while initializing session, please report this if it keeps happening")
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
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"sec-fetch-site":            {`none`},
		"sec-fetch-mode":            {`navigate`},
		"sec-fetch-user":            {`?1`},
		"sec-fetch-dest":            {`document`},
		"accept-encoding":           {"gzip, deflate, br"},
		"accept-language":           {"en-US,en;q=0.9"},
	}

	res, err := task.client.Do(req)
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Request failure while initializing session (Check proxies/internet connection)")
		fmt.Println(err)
		<-time.After(3 * time.Second)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		plog.TaskStatus(task.Count, "blue", "INFO", "Successfully loaded session")
		task.sessionTimestamp = time.Now().Unix()
		if task.itemId == 0 {
			task.Stage = LOAD_PRODUCT
		} else {
			task.Stage = ADD_TO_CART
		}
	} else if res.StatusCode >= 500 {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Server error while initializing session (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	} else {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Request error while initializing session (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	}
}
