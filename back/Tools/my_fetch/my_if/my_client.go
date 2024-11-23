// 封装，直接得到一个包含了ipv6地址的http.Client
package my_if

import (
	"context"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	*http.Client
	ip net.IP
}

func NewClient(prefix string, jar *cookiejar.Jar) (*Client, error) {
	ip := NewAddr(prefix)
	if err := AddAddr(ip.String()); err != nil {
		return nil, err
	}

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

	return &Client{
		Client: &http.Client{
			Transport: tr,
			Jar:       jar,
		},
		ip: ip,
	}, nil

}
