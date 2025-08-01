package controller

import (
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUser(c *gin.Context) {

}

func (uc *UserController) CreateUser(c *gin.Context) {}

func (uc *UserController) UpdateUser(c *gin.Context) {}

func (uc *UserController) DeleteUser(c *gin.Context) {}
