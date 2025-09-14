package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
* CORS 处理跨域请求的中间件。
* 允许所有来源的请求，并处理预检请求。
 */
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有域名的请求（即任何网站都能访问你的接口）。
		c.Header("Access-Control-Allow-Origin", "*")
		// 允许请求头中带有 Content-Type 字段。
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		// 允许这几种 HTTP 方法的请求。
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
