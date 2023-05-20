package proxies

import (
  "sync"
  "strings"
  "net/url"
)

func NewList(proxyString string) (p *ProxyList) {
  p = &ProxyList{
    nextProxy: 0,
    mutex: &sync.Mutex{},
  }

  proxyList := strings.Split(strings.Replace(proxyString, "\r", "", -1), "\n")

  for i := 0; i < len(proxyList); i++ {
    proxyParts := strings.Split(proxyList[i], ":")
    if len(proxyParts) == 2 {
      proxyURL, err := url.Parse("http://" + proxyParts[0] + ":" + proxyParts[1])
      if err == nil {
        p.Proxies = append(p.Proxies, proxyURL)
        p.TwoCaptchaProxies = append(p.TwoCaptchaProxies, proxyParts[0] + ":" + proxyParts[1])
      }
    } else if len(proxyParts) == 4 {
      proxyURL, err := url.Parse("http://" + proxyParts[2] + ":" + proxyParts[3] + "@" + proxyParts[0] + ":" + proxyParts[1])
      if err == nil {
        p.Proxies = append(p.Proxies, proxyURL)
        p.TwoCaptchaProxies = append(p.TwoCaptchaProxies, proxyParts[2] + ":" + proxyParts[3] + "@" + proxyParts[0] + ":" + proxyParts[1])
      }
    }
  }

  return p
}

func GetProxy(p *ProxyList) (*url.URL, string) {
  if len(p.Proxies) == 0 {
    return nil, ""
  }

  p.mutex.Lock()
  proxy := p.Proxies[p.nextProxy]
  twoCaptchaProxy := p.TwoCaptchaProxies[p.nextProxy]
  p.nextProxy++
  if len(p.Proxies) == p.nextProxy {
    p.nextProxy = 0
  }
  p.mutex.Unlock()
  return proxy, twoCaptchaProxy
}
