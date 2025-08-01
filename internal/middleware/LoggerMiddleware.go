package middleware

import (
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/gin-gonic/gin"
	"time"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 继续处理请求
		c.Next()

		duration := time.Since(start)
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		log.DebugF("[GIN] %s %s | %d | %v",
			c.Request.Method,
			path,
			c.Writer.Status(),
			duration,
		)
	}
}
