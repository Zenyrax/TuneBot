package tuneportals

import (
	"fmt"
	"io"
	"strings"
	"time"
	"tunebot/util/plog"

	"net/http"
)

// Loading checkout to get stripe key
func PreloadCheckout(task *Task) {
	plog.TaskStatus(task.Count, "yellow", "INFO", "Loading site data...")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/SecureCartCheckout?anon_ok=1", task.Site), nil)

	// Probably don't need this but whatever
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while getting site data, please report this if it keeps happening")
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
		"accept-encoding":           {"gzip, deflate, br"},
		"accept-language":           {"en-US,en;q=0.9"},
	}

	res, err := task.client.Do(req)
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Request failure while getting site data (Check proxies/internet connection)")
		fmt.Println(err)
		<-time.After(3 * time.Second)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			plog.TaskStatus(task.Count, "red", "ERROR", "Failure parsing response while getting site data")
			<-time.After(3 * time.Second)
			return
		}

		//fmt.Println(string(body))
		task.stripeKey = GetStringInBetween(string(body), `stripe = Stripe("`, `");`)
		if task.stripeKey == "" {
			plog.TaskStatus(task.Count, "red", "ERROR", "Could not find payment key")
			<-time.After(3 * time.Second)
		} else {
			plog.TaskStatus(task.Count, "blue", "INFO", "Successfully loaded site data")
			task.Stage = CREATE_SESSION
		}
	} else if res.StatusCode >= 500 {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Server error while getting site data (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	} else {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Request error while getting site data (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	}
}

// S/O Stackoverflow for the broken code I had to fix >:(
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : e+s]
}
