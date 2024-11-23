// gin-pack @ 2024-04-06
// azure-go @ 2023-12-21
// eh-web-viewer

package myfetch

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var (
	jar, _ = cookiejar.New(nil)
)

func SetCookieJar(newJar *cookiejar.Jar) {
	jar = newJar
}

func NewProxyClient(proxyUrl string) *http.Client {
	proxyURL, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	return &http.Client{
		Transport: tr,
		Jar:       jar,
	}
}

// client poll
type clientPool struct {
	cnt     uint8
	clients []*http.Client
}

func (cp *clientPool) Client() *http.Client {
	cp.cnt++
	if len(cp.clients) == 0 {
		return http.DefaultClient
	} else {
		return cp.clients[int(cp.cnt)%len(cp.clients)]
	}
}

func (cp *clientPool) SetClients(clients []*http.Client) {
	cp.clients = clients
}

func NewClientPool(clients []*http.Client) *clientPool {
	return &clientPool{
		cnt:     0,
		clients: clients,
	}
}

// public methods

var DefaultClientPool *clientPool

func Client() *http.Client {
	return DefaultClientPool.Client()
}

func SetClients(clients []*http.Client) {
	DefaultClientPool.SetClients(clients)
}
