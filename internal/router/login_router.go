package router

import (
	"github.com/1255177148/golangTask4/internal/app/controller"
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/1255177148/golangTask4/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func RegisterLoginRouter(rg *gin.RouterGroup, db *gorm.DB, sqlxDB *sqlx.DB) {
	userRepo := repository.NewUserRepository(db, sqlxDB)
	userService := service.NewUserService(db, sqlxDB, userRepo)
	loginCtrl := controller.NewLoginController(userService)
	login := rg.Group("/login")
	{
		login.GET("/getCaptcha", loginCtrl.GetCaptcha)
		login.POST("/register", loginCtrl.RegisterUser)
		login.POST("/", loginCtrl.Login)
		login.POST("/logout", middleware.JWTMiddleware(), loginCtrl.Logout) // 中间件必须在前
	}
}
