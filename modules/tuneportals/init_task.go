package tuneportals

import (
	"tunebot/util/plog"
	"tunebot/util/proxies"

	"net/http"
	"net/http/cookiejar"
)

// Create req client and get proxy if needed
func Init(task *Task) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		plog.TaskStatus(task.Count, "red", "ERROR", "Internal error while generating task data, please report this if it keeps happening")
		return
	}

	proxy, _ := proxies.GetProxy(task.Proxies)
	httpProxy := http.ProxyURL(proxy)

	if !task.UseProxy {
		httpProxy = nil
	}

	// fmt.Println(proxy, httpProxy)

	// Makes it easier to change the UA in the future if I just define it here
	task.userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
	task.secUa = `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`

	task.client = &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			Proxy:             httpProxy,
			ForceAttemptHTTP2: true,
		},
	}

	task.Stage = PRELOAD_CHECKOUT
}
