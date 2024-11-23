package myclient

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

var CloudFlareDNSDial = func(ctx context.Context, network, address string) (net.Conn, error) {
	return net.Dial("udp", "1.1.1.1:53") // manually set to cloudlfare dns for resolve ipv6 address. 240725
}
var CloudFlareResolver = &net.Resolver{
	PreferGo: false, // it should be flase. 240725
	// or it will give unable to lookup error.
	Dial: CloudFlareDNSDial, // 就为什么要自己处理只给控制conn这个参数啊，就很无助。
}
var DefaultCookieJar, _ = cookiejar.New(nil)

func NewClientWithLocalAddress(localAddress net.IP, jar *cookiejar.Jar) *http.Client {
	tr := &http.Transport{
		DialContext: (&net.Dialer{ // dialer
			// LocalAddr 用于指定本地 IP 地址
			LocalAddr: &net.TCPAddr{
				IP: localAddress, // 替换为你要使用的特定 IP 地址
			},
			Timeout:   5 * time.Second,  // 连接超时时间
			KeepAlive: 30 * time.Second, // Keep-Alive 超时时间
			Resolver:  CloudFlareResolver,
		}).DialContext,
	}
	return &http.Client{
		Transport: tr,
		Jar:       jar,
	}
}

func NewInsecureClient(jar *cookiejar.Jar) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
		Jar:       jar,
	}
}

func NewProxyClient(proxyURL *url.URL, jar *cookiejar.Jar) *http.Client {
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	return &http.Client{
		Transport: tr,
		Jar:       jar,
	}
}
