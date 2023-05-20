package discord

import (
  "time"
  "strconv"
  "strings"
  "encoding/json"

  "net/http"
)

func SendWebhook(webhook *Webhook, url string) {
  out, err := json.Marshal(webhook)
  if err != nil {
      return
  }

  if !strings.Contains(url, "?wait=true") {
    url = url + "?wait=true"
  }

  for {
    payload := strings.NewReader(string(out))

    tr := &http.Transport{}

    req, err := http.NewRequest("POST", url, payload)
    if err != nil {
      return
    }
    req.Header.Add("content-type", "application/json")

    res, err := tr.RoundTrip(req)

    if err == nil {
      if res.StatusCode == 429 && res.Header["Retry-After"] != nil && len(res.Header["Retry-After"]) == 1 {
        res.Body.Close()
        i, err := strconv.Atoi(res.Header["Retry-After"][0])
        if err == nil {
          <-time.After(time.Duration(i) * time.Millisecond)
        }
      } else if res.StatusCode == 200 || res.StatusCode == 204 || res.StatusCode == 400 {
        res.Body.Close()
        break
      }
    }
  }
}
