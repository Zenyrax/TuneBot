package proxies

import (
  "sync"
  "net/url"
)

type ProxyList struct {
  Proxies []*url.URL
  TwoCaptchaProxies []string

  nextProxy int
  mutex *sync.Mutex
}
