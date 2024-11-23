
package main

import (
	"bytes"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	tools "github.com/eh-web-viewer/eh-web-viewer/Tools"
	"github.com/eh-web-viewer/eh-web-viewer/Tools/debug"
	myfetch "github.com/eh-web-viewer/eh-web-viewer/Tools/my_fetch"
	"github.com/eh-web-viewer/eh-web-viewer/Tools/my_fetch/my_if"
	middleware "github.com/eh-web-viewer/eh-web-viewer/Tools/my_gin_middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/net/html"

	"github.com/antchfx/htmlquery"
)

const InnerText string = "INNER_TEXT"

func findOneAndSelectAttr(top *html.Node, expr string, name string) (v string, err error) {
	elem := htmlquery.FindOne(top, expr)
	if elem == nil {
		err = errors.New(expr + ":" + name + "is null")
		return
	}
	if name == InnerText {
		v = htmlquery.InnerText(elem)
	} else {
		v = htmlquery.SelectAttr(elem, name)
	}
	return
}

func findAll(top *html.Node, expr, name string) (v []string) {
	elemArray := htmlquery.Find(top, expr)
	v = make([]string, len(elemArray))
	for i, e := range elemArray {
		if name == InnerText {
			v[i] = htmlquery.InnerText(e)
		} else {
			v[i] = htmlquery.SelectAttr(e, name)
		}
	}
	return
}

var jar *cookiejar.Jar = nil

func ExhProxy() {
	godotenv.Load(".env")

	debug.LogLevel = debug.Fatal

	prefix := tools.NewSlice(
		os.Getenv("EXHENTAI_PROXY_PREFIX"),
		"2001:470:c:6c:",
	).FirstNonDefaultValue("")

	ips := []net.IP{my_if.NewAddr(prefix), my_if.NewAddr(prefix)}
	ipidx := 0

	my_if.AddAddr(ips[ipidx].String())
	cp := myfetch.NewClientPool([]*http.Client{
		myfetch.NewV6Client(ips[ipidx], jar),
	})

	// cp = nil // debug
	mf := myfetch.NewFetcher(nil, cp)

	// gin
	r := gin.Default()

	// 设置block
	// r.Use(middleware.BlockMiddleware()) // 改到下面

	// 设置 CORS 头
	r.Use(middleware.CORSMiddleware())

	// 定义一个简单的 GET 路由
	r.Any("/*any", func(c *gin.Context) {

		if strings.HasPrefix(c.Request.URL.String(), "/api") {
			c.AbortWithStatus(http.StatusGone)
			return
		}

		if strings.HasPrefix(c.Request.URL.String(), "/fullimg") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if strings.HasPrefix(c.Request.URL.String(), "/archiver.php") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// 遗留问题, image这个path重定向到用param的请求
		if strings.HasPrefix(c.Request.URL.String(), "/image") {
			path := strings.TrimPrefix(c.Request.URL.String(), "/image")
			parsedURL, err := url.Parse(path)
			if err != nil {
				c.Header("X-Error", parsedURL.String())
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			query := parsedURL.Query()
			query.Set("redirect_to", "image")
			parsedURL.RawQuery = query.Encode()

			c.Redirect(http.StatusMovedPermanently, parsedURL.String())
			return
		}

	}, func(c *gin.Context) {

		// 获取请求中的 Cookie
		cookie, err := c.Cookie("pass")

		// 如果 Cookie 包含 pass=pass，直接继续处理请求
		if err == nil && cookie == "pass" {
			return
		}

		if c.Query("redirect_to") != "" {
			return
		}

		// 如果不在 "CN", "" 中的任意一个。
		if !slices.Contains([]string{"CN", ""}, c.Request.Header.Get("Cf-Ipcountry")) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message":      "禁止访问",
				"Cf-Ipcountry": c.Request.Header.Get("Cf-Ipcountry"),
			})
			return
		}

	}, func(c *gin.Context) {
		c.Header("X-Debug-Request-Host", c.Request.Host)     // 要设置 Host $http_host
		c.Header("X-Debug-Header-Host", c.GetHeader("Host")) // never

		if c.Request.Body != nil {
			defer c.Request.Body.Close()
		}
		// 读取 URL 参数
		path := c.Request.URL.String()

		host := "exhentai.org"

		header := tools.NewHeader(c.Request.Header)
		header.Set(
			"Cookie",
			tools.NewSlice(
				c.GetHeader("X-Cookie"),
				os.Getenv("EXHENTAI_PROXY_COOKIE"),
			).FirstNonDefaultValue(""),
		)

		resp, err := mf.Fetch(
			c.Request.Method, "https://"+host+path,
			(header.Header), c.Request.Body)
		if err != nil {
			debug.E("why", err.Error())
			c.Header("X-Error", err.Error())
			c.AbortWithError(http.StatusBadGateway, err)
			return
		}
		defer resp.Body.Close()

		if mf.Count() > 250 {
			ipidx = (ipidx + 1) % len(ips)
			defer func(ip string) {
				time.Sleep(60 * time.Second)
				my_if.DelAddr(ip)
			}(ips[ipidx].String())
			ips[ipidx] = my_if.NewAddr(prefix)
			my_if.AddAddr(ips[ipidx].String())
			newCp := myfetch.NewClientPool([]*http.Client{myfetch.NewV6Client(ips[ipidx], jar)})
			mf.SetClientPool(newCp)
		}

		// resp 获得 text
		// 需要override https://exhentai.org/s/, https://exhentai.org/g/
		body, err := myfetch.ResponseToReader(resp)
		if err != nil {
			debug.E("why", err.Error())
			c.Header("X-Error", err.Error())
			c.AbortWithError(http.StatusBadGateway, err)
			return
		}

		data, err := io.ReadAll(body)
		if err != nil {
			debug.E("why", err.Error())
			c.Header("X-Error", err.Error())
			c.AbortWithError(http.StatusBadGateway, err)
			return
		}

		// data = bytes.ReplaceAll(data, []byte("https://exhentai.org/g/"), []byte("https://"+"ex.nmbyd1.top"+"/g/"))
		// data = bytes.ReplaceAll(data, []byte("https://exhentai.org/s/"), []byte("https://"+"ex.nmbyd1.top"+"/s/"))
		// data = bytes.ReplaceAll(data, []byte("https://exhentai.org/z/"), []byte("https://"+"ex.nmbyd1.top"+"/z/"))
		// data = bytes.ReplaceAll(data, []byte("https://exhentai.org/img/"), []byte("https://"+"ex.nmbyd1.top"+"/img/"))
		data = bytes.ReplaceAll(data, []byte("https://exhentai.org"), []byte{})
		c.Header("X-Debug", c.GetHeader("Host"))
		data = bytes.ReplaceAll(data, []byte("https://s.exhentai.org"), []byte("https://s-ex.moonchan.xyz"))
		if strings.HasPrefix(path, `/s/`) {
			data = []byte(addWaterFallViewButton(string(data)))
		}
		data = addFloatingIframeAtRightBottom(data)

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

		// 获得param，redirect_to=image
		// 没经过类型检查
		if strings.HasPrefix(path, "/s/") && c.Query("redirect_to") == "image" {
			doc, err := htmlquery.Parse(bytes.NewReader(data))
			if err != nil {
				c.Header("X-Error", err.Error())
				c.AbortWithStatus(http.StatusBadGateway)
				return
			}

			image, err := findOneAndSelectAttr(doc, "//img[@id='img']", "src")
			c.Redirect(http.StatusFound, image)
			return
		}

		if strings.HasPrefix(path, "/g/") && c.Query("redirect_to") == "cover" {
			doc, err := htmlquery.Parse(bytes.NewReader(data))
			if err != nil {
				c.Header("X-Error", err.Error())
				c.AbortWithStatus(http.StatusBadGateway)
				return
			}
			hrefArray := findAll(doc, "//a", "href")
			for _, href := range hrefArray {
				if strings.HasPrefix(href, "/s/") {
					parsedURL, err := url.Parse(href)
					if err != nil {
						c.Header("X-Error", parsedURL.String())
						c.AbortWithStatus(http.StatusBadRequest)
						return
					}
					query := parsedURL.Query()
					query.Set("redirect_to", "image")
					parsedURL.RawQuery = query.Encode()

					c.Redirect(http.StatusFound, parsedURL.String())
					return
				}
			}
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if len(data) == 0 {
			debug.E("why", resp.Status) // 304
			c.Header("X-Error", resp.Status)
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
	r.Run("127.25.23.2:8080") // 在 8080 端口启动服务

}

func SProxy() {
	r := gin.Default()

	// 设置 CORS 头
	r.Use(middleware.CORSMiddleware())

	// 定义一个简单的 GET 路由
	r.Any("/*any", func(c *gin.Context) {
		if c.Request.Body != nil {
			defer c.Request.Body.Close()
		}
		// 读取 URL 参数
		path := c.Request.URL.String()

		if strings.HasPrefix(path, "/api") {
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}

		// if strings.HasPrefix(path, "/fullimg/") {
		// 	c.AbortWithStatus(http.StatusServiceUnavailable)
		// 	return
		// }

		host := "s.exhentai.org"

		header := tools.NewHeader(c.Request.Header)

		resp, err := myfetch.Fetch(
			c.Request.Method, "https://"+host+path,
			(header.Header), c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		defer resp.Body.Close()

		for k, vs := range resp.Header {
			if c.Writer.Header().Get(k) != "" {
				continue
			}
			for _, v := range vs {
				c.Writer.Header().Add(k, v)
			}
		}

		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, map[string]string{
			"X-Host":    host,
			"X-Origin":  header.Get("Origin"),
			"X-Referer": header.Get("Referer"),
			"X-Cookie":  header.Get("Cookie"),
		})

	})

	// 启动服务器
	r.Run("127.25.23.3:8080") // 在 8080 端口启动服务

}

func addWaterFallViewButton(html string) string {
	return strings.Replace(html, "<body>", `<body>
	<div style="
	  height: 60px;
	  width: 100px;
	  text-align: center;
	  /* background-color: violet; */
	  position: fixed;
	  right: 20px; 
	  top: 20px;
	  z-index: 99;
	  display: table-cell;
	  vertical-align: middle;
	  /* float: right; */
	">
	  <button id="waterfall" style="
			width: 100%;    
			height: 100%;
			font-size: x-large;
	  ">
			下拉式
	  </button>
	</div>
  <script type="text/javascript">
	async function execWaterfall(){
		console.log('!');
		document.getElementById("waterfall").remove();
		let pn = document.createElement('div');
		let lp = location.href;
		let ln = location.href;
		const element = document.getElementById('i1');
		element.appendChild(pn);
		let hn = document.getElementById('next').href;
		while (hn != ln) {
		  let doc;
		  while(!doc) {
			doc = await fetch(hn).then(resp => resp.text())			
			  .then(data => {
			    console.log(data);
			    let parser = new DOMParser();
			    let doc = parser.parseFromString(data, "text/html");
			    return doc;
			  });
			}
		  console.log(doc);
		  let img = document.createElement('img');
		  let element = doc.getElementById('img');
		  if (element) {
			img.src = element.src;
			pn.appendChild(img);
			ln = hn;
			hn = doc.getElementById('next').href;
		  }
		}
		let p = document.createElement('p');
		p.innerHTML = hn;
	  }
	document.getElementById("waterfall").addEventListener("click", execWaterfall, false); 
	</script>`, 1)
}

func addFloatingIframeAtRightBottom(html []byte) []byte {
	html = bytes.Replace(html,
		[]byte("</head>"),
		[]byte(`
	<style>
		#moonchan-floating-iframe {
			position: fixed;
			bottom: 20px; /* 距离底部的距离 */
			right: 20px; /* 距离右侧的距离 */
			width: 300px; /* 根据需要调整宽度 */
			height: 200px; /* 根据需要调整高度 */
			border: 2px solid #ccc; /* 边框样式 */
			border-radius: 8px; /* 圆角边框 */
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.2); /* 阴影效果 */
			z-index: 100000; /* 确保在最上层 */
			overflow: hidden; /* 确保内容不超出边框 */
		}       
		#moonchan-close-button {
            position: absolute;
            top: 5px;
            right: 5px;
            background-color: red; /* 按钮颜色 */
            color: white; /* 字体颜色 */
            border: none;
            border-radius: 50%;
            width: 25px;
            height: 25px;
            cursor: pointer;
            font-size: 18px;
            line-height: 25px; /* 垂直居中 */
            text-align: center;
        }
	</style>
</head>`), 1)
	html = bytes.Replace(html,
		[]byte("<body>"),
		[]byte(`<body>
    <div id="moonchan-floating-iframe" style="display: none;">
        <button id="moonchan-close-button" onclick="moonchanCloseIframe()">×</button>
        <iframe src="https://moonchan.xyz/iframe.html" style="border: none; width: 100%; height: calc(100% - 30px);"></iframe>
    </div>

    <script>
        // 检查 localStorage 中的值
        if (localStorage.getItem('iframeClosed') !== 'true') {
            document.getElementById('moonchan-floating-iframe').style.display = 'block';
        }

        function moonchanCloseIframe() {
            const iframeContainer = document.getElementById('moonchan-floating-iframe');
            iframeContainer.style.display = 'none'; // 隐藏 iframe
            localStorage.setItem('iframeClosed', 'true'); // 设置 localStorage 标记
        }
    </script>

`), 1)
	return html
}
