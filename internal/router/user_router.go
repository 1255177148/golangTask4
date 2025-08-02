package router

import (
	"github.com/1255177148/golangTask4/internal/app/controller"
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// RegisterUserRoutes 用户模块路由注册
func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB, sqlxDB *sqlx.DB) {
	userRepo := repository.NewUserRepository(db, sqlxDB)
	userService := service.NewUserService(db, sqlxDB, userRepo)
	userCtrl := controller.NewUserController(userService)

	users := rg.Group("/users")
	{
		users.GET("/auth", userCtrl.AuthUser)
	}
}
