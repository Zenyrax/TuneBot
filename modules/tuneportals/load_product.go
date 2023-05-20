package tuneportals

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"
	"tunebot/util/plog"

	"net/http"
)

// Loading product info
func LoadProduct(task *Task) {
	plog.TaskStatus(task.Count, "yellow", "INFO", "Getting product data...")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/KioskAPIItem.json?upc=%s", task.Site, task.UPC), nil)

	// Probably don't need this but whatever
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while getting product data, please report this if it keeps happening")
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
		"referer":                   {fmt.Sprintf("https://%s/UPC/%s", task.Site, task.UPC)},
		"accept-encoding":           {"gzip, deflate, br"},
		"accept-language":           {"en-US,en;q=0.9"},
	}

	res, err := task.client.Do(req)
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Request failure while getting product data (Check proxies/internet connection)")
		fmt.Println(err)
		<-time.After(3 * time.Second)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			plog.TaskStatus(task.Count, "red", "ERROR", "Failure parsing response while getting product data")
			<-time.After(3 * time.Second)
			return
		}
		var response productResponse
		json.Unmarshal(body, &response)
		// fmt.Println(string(body))
		// fmt.Println(response)
		price, _ := strconv.ParseFloat(response.Item.Price, 64)
		if task.PriceLimit == 0 || (price < task.PriceLimit) {
			task.itemId = response.Item.ItemID
			if response.Item.PreorderFlag == 1 {
				task.itemType = "preorder"
			} else {
				task.itemType = "new"
			}
			task.name = response.Item.Title
			task.image = response.Item.ThumbnailURL
			task.price = "$" + response.Item.Price
			task.mediaType = response.Item.MediaType
			plog.TaskStatus(task.Count, "blue", "INFO", fmt.Sprintf("Successfully loaded product data (%s)", response.Item.Title))
			task.Stage = ADD_TO_CART
		} else {
			plog.TaskStatus(task.Count, "purple", "INFO", fmt.Sprintf("Price of item ($%v) is higher than limit ($%v)", price, task.PriceLimit))
			<-time.After(3 * time.Second)
		}
	} else if res.StatusCode >= 500 {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Server error while getting product data (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	} else {
		plog.TaskStatus(task.Count, "red", "ERROR", fmt.Sprintf("Request error while getting product data (%d)", res.StatusCode))
		<-time.After(3 * time.Second)
	}
}
