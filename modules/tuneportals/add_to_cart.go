package tuneportals

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
	"tunebot/util/plog"

	"net/http"
	"net/http/cookiejar"
)

// Add product to cart and keep trying if OOS
func AddToCart(task *Task) {
	plog.TaskStatus(task.Count, "yellow", "INFO", "Adding to cart...")

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/SecureCartAdd.json", task.Site), strings.NewReader(fmt.Sprintf("type=%s&id=%d&qty=%s", task.itemType, task.itemId, task.Quantity)))

	// Probably don't need this but whatever
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while adding to cart, please report this if it keeps happening")
		return
	}

	// I don't actually know if they care about headers that much, just wanna keep everything as accurate as possible in case they do
	req.Header = http.Header{
		"cache-control":             {`max-age=0`},
		"content-type":              {`application/x-www-form-urlencoded; charset=UTF-8`},
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
		"referer":                   {fmt.Sprintf("https://%s/UPC/%s", task.Site, task.UPC)},
		"accept-encoding":           {"gzip, deflate, br"},
		"accept-language":           {"en-US,en;q=0.9"},
	}

	res, err := task.client.Do(req)
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Request failure while adding to cart (Check proxies/internet connection)")
		fmt.Println(err)
		<-time.After(3 * time.Second)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			plog.TaskStatus(task.Count, "red", "ERROR", "Failure parsing response while adding to cart")
			<-time.After(3 * time.Second)
			return
		}
		var response atcResponse
		json.Unmarshal(body, &response)
		// fmt.Println(string(body))
		// fmt.Println(response)
		if len(response.AnonCart) > 0 {
			plog.TaskStatus(task.Count, "blue", "INFO", "Successfully added to cart")
			_, err := strconv.ParseInt(task.ShippingOption, 10, 64)
			if err != nil {
				task.Stage = LOAD_CART
			} else {
				task.Stage = TOKENIZE_PAYMENT
			}
		} else {
			plog.TaskStatus(task.Count, "purple", "INFO", "Failed to add to cart (Likely out of stock)")
			<-time.After(3 * time.Second)
			// Idk how long it takes for sessions to expire, but I know it's fast
			if time.Now().Unix()-task.sessionTimestamp > 600 {
				jar, _ := cookiejar.New(nil)
				task.client.Jar = jar
				task.Stage = CREATE_SESSION
			}
		}
	} else if res.StatusCode >= 500 {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Server error while adding to cart (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	} else {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Request error while adding to cart (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	}
}
