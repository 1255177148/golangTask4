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
// @Summary 注册用户
// @Description 注册一个用户
// @Tags 登录
// @Produce json
// @Param User body dto.UserDTO true "注册表单参数"
// @Success 200 {object} response.ResultResponse
// @Router /login/register [post]
func (ctrl *LoginController) RegisterUser(c *gin.Context) {
	var user po.User
	if err := binder.BindAndValidate(c, &user); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	if err := ctrl.userService.RegisterUser(&user); err != nil {
		log.Error(constant.RegisterFail, zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.ServiceFail)
		return
	}
	response.Success(c, nil)
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户通过用户名和密码登录
// @Tags 登录
// @Produce json
// @Param UserDTO body dto.UserDTO true "用户登录参数"
// @Success 200 {object} response.ResultResponse
// @Router /login [post]
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

// Logout 用户退出登录
// @Summary 用户退出登录
// @Description 用户退出登录
// @Tags 登录
// @Produce json
// @Success 200 {object} response.ResultResponse
// @Router /login/logout [post]
func (ctrl *LoginController) Logout(c *gin.Context) {
	userId, _ := c.Get("user_id")
	err := ctrl.userService.Logout(userId.(uint))
	if err != nil {
		log.Error(constant.LogoutFail, zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.LogoutFail)
		return
	}
	response.Success(c, nil)
}
