package router

import (
	"github.com/1255177148/golangTask4/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// RegisterRoutes 统一注册路由
func RegisterRoutes(r *gin.Engine, db *gorm.DB, sqlxDB *sqlx.DB) {
	api := r.Group("/api/v1", middleware.CORSMiddleware(), middleware.LoggerMiddleware(), middleware.JWTMiddleware())
	loginApi := r.Group("/api", middleware.CORSMiddleware(), middleware.LoggerMiddleware())
	// 注册登录路由
	RegisterLoginRouter(loginApi, db, sqlxDB)
	// 注册用户路由
	RegisterUserRoutes(api, db, sqlxDB)
	// 注册文章路由
	RegisterPostRoutes(api, db, sqlxDB)
}
