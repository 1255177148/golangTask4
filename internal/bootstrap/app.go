package bootstrap

import (
	"github.com/1255177148/golangTask4/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// InitApp 初始化gin
func InitApp(db *gorm.DB, sqlxDB *sqlx.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery()) // 不加 Logger
	router.RegisterRoutes(r, db, sqlxDB)
	return r
}
