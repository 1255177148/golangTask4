package router

import (
	"github.com/1255177148/golangTask4/internal/app/controller"
	"github.com/1255177148/golangTask4/internal/container"
	"github.com/1255177148/golangTask4/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterLoginRouter(rg *gin.RouterGroup) {
	userService := container.Instance.UserService
	loginCtrl := controller.NewLoginController(userService)
	login := rg.Group("/login")
	{
		login.GET("/getCaptcha", loginCtrl.GetCaptcha)
		login.POST("/register", loginCtrl.RegisterUser)
		login.POST("/", loginCtrl.Login)
		login.POST("/logout", middleware.JWTMiddleware(), loginCtrl.Logout) // 中间件必须在前
	}
}
