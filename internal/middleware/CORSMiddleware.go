package middleware

import (
	"github.com/1255177148/golangTask4/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	_config := cors.Config{
		AllowOrigins:     config.Cfg.CORS.AllowedOrigins,
		AllowMethods:     config.Cfg.CORS.AllowedMethods,
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	c := cors.New(_config)
	return func(ctx *gin.Context) {
		c(ctx)     // 调用 cors 的 applyCors 逻辑
		ctx.Next() // 关键：调用 Next，执行后续中间件和路由
	}
}
