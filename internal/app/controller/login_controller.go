package controller

import (
	"github.com/1255177148/golangTask4/internal/app/model/dto"
	"github.com/1255177148/golangTask4/internal/app/model/po"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/1255177148/golangTask4/internal/binder"
	"github.com/1255177148/golangTask4/internal/constant"
	"github.com/1255177148/golangTask4/internal/utils"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/1255177148/golangTask4/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type LoginController struct {
	userService *service.UserService
}

func NewLoginController(userService *service.UserService) *LoginController {
	return &LoginController{userService: userService}
}

func (ctrl *LoginController) GetCaptcha(c *gin.Context) {
	utils.GetCaptcha(c)
}

// RegisterUser 注册用户
func (ctrl *LoginController) RegisterUser(c *gin.Context) {
	var user po.User
	if err := binder.BindAndValidate(c, &user); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	if err := ctrl.userService.RegisterUser(&user); err != nil {
		log.Error(constant.RegisterFail, zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// Login 用户登录
func (ctrl *LoginController) Login(c *gin.Context) {
	var user dto.UserDTO
	if err := binder.BindAndValidate(c, &user); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	token, err := ctrl.userService.CheckLogin(&user)
	if err != nil {
		log.Error(constant.LoginFail, zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.LoginFail)
		return
	}
	response.Success(c, token)
}

// AuthUser 认证用户
func (ctrl *LoginController) AuthUser(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var userAuth dto.UserAuth
	if err := binder.BindAndValidate(c, &userAuth); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	userAuth.ID = userId.(uint)
	if err := ctrl.userService.AuthUser(&userAuth); err != nil {
		log.Error(constant.AuthFail, zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.AuthFail)
		return
	}
	response.Success(c, nil)

}
