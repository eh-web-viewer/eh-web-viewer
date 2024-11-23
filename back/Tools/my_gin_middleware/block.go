package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func BlockMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取请求中的 Cookie
		cookie, err := c.Cookie("pass")

		// 如果 Cookie 包含 pass=pass，直接继续处理请求
		if err == nil && cookie == "pass" {
			c.Next() // 继续处理请求
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

		c.Next() // 继续处理请求
	}
}

func RedirectMiddleware(host string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if host == "" {
			c.Next()
			return
		}

		// 获取请求中的 Cookie
		cookie, err := c.Cookie("pass")

		// 如果 Cookie 包含 pass=pass，直接继续处理请求
		if err == nil && cookie == "pass" {
			c.Next() // 继续处理请求
			return
		}

		// 如果不在 "CN", "" 中的任意一个。
		if !slices.Contains([]string{"CN", ""}, c.Request.Header.Get("Cf-Ipcountry")) {
			path := c.Request.URL.String()
			c.Redirect(http.StatusFound, "https://"+host+path)
			return
		}

		c.Next() // 继续处理请求
	}
}
