package router

import (
	"github.com/1255177148/golangTask4/internal/app/controller"
	"github.com/1255177148/golangTask4/internal/container"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 用户模块路由注册
func RegisterUserRoutes(rg *gin.RouterGroup) {
	userService := container.Instance.UserService
	userCtrl := controller.NewUserController(userService)

	users := rg.Group("/users")
	{
		users.GET("/auth", userCtrl.AuthUser)
	}
}
