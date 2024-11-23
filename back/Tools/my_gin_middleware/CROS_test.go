package middleware

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

// func main() {
func TestCROS(t *testing.T) {
	router := gin.Default()

	// 使用 CORS 中间件
	router.Use(CORSMiddleware())

	// 定义路由
	router.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// 启动服务器
	router.Run(":8080")
}
