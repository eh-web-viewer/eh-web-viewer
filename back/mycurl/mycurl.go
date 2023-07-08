package mycurl

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/eh-web-viewer/eh-web-viewer/my_if"
)

// generate clients.
const (
	VANILLA  string = "vanilla"
	V6POOL   string = "v6pool"
	PROXY    string = "proxy"
	INSECURE string = "insecure"
)

var (
	jar, _ = cookiejar.New(nil)
)

func newV6Client(ip net.IP) *http.Client {
	tr := &http.Transport{
		DialContext: (&net.Dialer{ // dialer
			// LocalAddr 用于指定本地 IP 地址
			LocalAddr: &net.TCPAddr{
				IP: ip, // 将 "your_specific_ip" 替换为你要使用的特定 IP 地址
			},
			Timeout:   5 * time.Second,  // 连接超时时间
			KeepAlive: 30 * time.Second, // Keep-Alive 超时时间
			Resolver: &net.Resolver{
				PreferGo: true,
				Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
					return net.Dial("udp", "1.1.1.1:53")
				},
			},
		}).DialContext,
	}
	return &http.Client{
		Transport: tr,
		Jar:       jar,
	}
}

func newClients(clientType string /*, ips ...net.IP*/) []*http.Client {
	switch clientType {
	case VANILLA:
		return []*http.Client{&http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Jar: jar,
		}}
	case V6POOL:
		ips := my_if.GetIPBatchAndShift()
		clients := make([]*http.Client, len(ips))
		for i, ip := range ips {
			clients[i] = newV6Client(ip)
		}
		return clients
	case INSECURE:
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		return []*http.Client{&http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: tr,
			Jar:       jar,
		}}
	case PROXY:
		proxyURL, _ := url.Parse("http://localhost:10809")
		tr := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		return []*http.Client{&http.Client{
			Transport: tr,
			Jar:       jar,
		}}
	default:
		log.Println(clientType, "not defined")
		return nil
	}
}

// as a client.

const MAX_CNT = MAX_SINGLE_CNT * my_if.BATCH_SIZE

var clientType = "vanilla"
var clients = newClients(clientType)
var cnt = 0

func shift() {
	go func() {
		cnt++
		if cnt > MAX_CNT {
			cnt = 0
			clients = newClients(clientType)
			// log.Println(clients)
		}
	}()
}

func client() *http.Client {
	defer shift()
	// log.Println(cnt)
	return clients[cnt%len(clients)]
}

func fetchWithRequest(req *http.Request) (*http.Response, error) {
	return client().Do(req)
}

func generateRequest(method, url string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(
		method,
		url,
		body,
	)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func fetch(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := generateRequest(method, url, headers, body)
	if err != nil {
		return nil, err
	}
	return fetchWithRequest(req)
}

func SetClient(nextClientType string) {
	clientType = nextClientType
	clients = newClients(clientType)
}

func Fetch(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	return fetch(method, url, headers, body)
}

func Client() *http.Client {
	return client()
}
