package router

import (
	"github.com/1255177148/golangTask4/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 统一注册路由
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1", middleware.CORSMiddleware(), middleware.LoggerMiddleware(), middleware.JWTMiddleware())
	loginApi := r.Group("/api", middleware.CORSMiddleware(), middleware.LoggerMiddleware())
	// 注册登录路由
	RegisterLoginRouter(loginApi)
	// 注册用户路由
	RegisterUserRoutes(api)
	// 注册文章路由
	RegisterPostRoutes(api)
	// 注册评论路由
	RegisterCommentRouters(api)
}
