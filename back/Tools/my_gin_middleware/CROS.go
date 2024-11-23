package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

// 使用 strings.Join 将切片连接成一个字符串
var headers = ([]string{
	"Accept",
	"Authorization",

	"X-Host",
	"X-Origin",
	"X-Referer",

	"X-Requested-With",

	"Cache-Control",
})

// 输出结果

// CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 CORS 头
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}

		c.Header("Access-Control-Allow-Origin", origin) // 或者指定特定的源
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
		// c.Header("Access-Control-Allow-Headers", strings.Join(headers, ", "))
		c.Header("Access-Control-Allow-Headers", c.GetHeader("access-control-request-headers"))
		c.Header("Access-Control-Allow-Credentials", "true") // 允许cookie传递

		// 处理预检请求
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent) // 返回 204 No Content
			return
		}

		c.Next() // 继续处理请求

		// 为什么自带的方法这么贵物
		exposeHeaders := make([]string, 0, len(c.Writer.Header()))
		for k, _ := range c.Writer.Header() {
			exposeHeaders = append(exposeHeaders, k)
		}
		slices.Sort(exposeHeaders)
		c.Writer.Header().Add("Access-Control-Expose-Headers", strings.Join(exposeHeaders, ", "))
	}
}
