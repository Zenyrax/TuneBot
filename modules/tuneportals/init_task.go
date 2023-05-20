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

	task.client = &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			Proxy:             httpProxy,
			ForceAttemptHTTP2: true,
		},
	}

	task.Stage = PRELOAD_CHECKOUT
}
