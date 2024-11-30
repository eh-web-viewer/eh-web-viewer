package main

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"slices"
	"strconv"
	"strings"

	tools "github.com/eh-web-viewer/eh-web-viewer/Tools"
	myfetch "github.com/eh-web-viewer/eh-web-viewer/Tools/my_fetch"
	middleware "github.com/eh-web-viewer/eh-web-viewer/Tools/my_gin_middleware"
	"github.com/gin-gonic/gin"
)

func SukebeiProxy() {

	var client *http.Client = func() *http.Client {
		jar, _ := cookiejar.New(nil)
		u, _ := url.Parse("sukebei.nyaa.si")
		jar.SetCookies(u, []*http.Cookie{})
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
	}()

	cp := myfetch.NewClientPool([]*http.Client{client})
	mf := myfetch.NewFetcher(nil, cp)

	r := gin.Default()

	// 设置block
	// r.Use(middleware.BlockMiddleware())

	// 设置 CORS 头
	r.Use(middleware.CORSMiddleware())

	// 定义一个简单的 GET 路由
	r.Any("/*any", func(c *gin.Context) {
		c.Header("X-Debug-Request-Host", c.Request.Host)     // 要设置 Host $http_host
		c.Header("X-Debug-Header-Host", c.GetHeader("Host")) // never

		if c.Request.Body != nil {
			defer c.Request.Body.Close()
		}
		// 读取 URL 参数
		path := c.Request.URL.String()

		host := "sukebei.nyaa.si"

		header := tools.NewHeader(c.Request.Header)

		// 如果不在 "CN", "" 中的任意一个。
		if !slices.Contains([]string{"CN", ""}, c.Request.Header.Get("Cf-Ipcountry")) {
			c.Redirect(http.StatusFound, "https://"+host+path)
			return
		}

		resp, err := mf.Fetch(
			c.Request.Method, "https://"+host+path,
			(header.Header), c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		defer resp.Body.Close()

		if strings.HasPrefix(path, "/download") {
			c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, map[string]string{
				"X-Host":    host,
				"X-Origin":  header.Get("Origin"),
				"X-Referer": header.Get("Referer"),
				"X-Cookie":  header.Get("Cookie"),
			})
			return
		}

		body, err := myfetch.ResponseToReader(resp)
		if err != nil {
			c.AbortWithError(http.StatusBadGateway, err)
			return
		}

		data, err := io.ReadAll(body)
		if err != nil {
			c.AbortWithError(http.StatusBadGateway, err)
			return
		}

		data = bytes.ReplaceAll(data, []byte("https://"+host), []byte{})
		// remove ad
		data = bytes.ReplaceAll(data, []byte(`magsrv.com`), []byte(`localhost`))
		// data = bytes.ReplaceAll(data, []byte(`<div id="dd4ce992-766a-4df0-a01d-86f13e43fd61"></div>`), []byte{})
		// data = bytes.ReplaceAll(data, []byte(`<div id="e7a3ddb6-efae-4f74-a719-607fdf4fa1a1"></div>`), []byte{})

		// 为什么自带的方法这么贵物
		c.Writer.Header().Set("Content-Encoding", "identity")
		// 和最后的方法到底用哪个。
		c.Writer.Header().Set("Content-Length", strconv.Itoa(len(data)))
		for k, vs := range resp.Header {
			if c.Writer.Header().Get(k) != "" {
				continue
			}
			for _, v := range vs {
				c.Writer.Header().Add(k, v)
			}
		}

		// log.Println(string(data[:1024]))
		if len(data) == 0 {
			c.AbortWithStatus(resp.StatusCode)
			return
		}

		// 为什么会报多写。
		c.DataFromReader(resp.StatusCode, int64(len(data)), resp.Header.Get("Content-Type"), bytes.NewReader(data), map[string]string{
			"X-Host":    host,
			"X-Origin":  header.Get("Origin"),
			"X-Referer": header.Get("Referer"),
			"X-Cookie":  header.Get("Cookie"),
		})

	})

	// 启动服务器
	r.Run("127.25.23.5:8080") // 在 8080 端口启动服务

}
