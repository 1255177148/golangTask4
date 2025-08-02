package controller

import (
	"github.com/1255177148/golangTask4/internal/app/model/dto"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/1255177148/golangTask4/internal/binder"
	"github.com/1255177148/golangTask4/internal/constant"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/1255177148/golangTask4/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// AuthUser 认证用户
// @Summary 认证用户
// @Description 用户通过token和邮箱认证
// @Tags 用户
// @Produce json
// @Param request formData dto.UserAuth true "用户认证参数"
// @Success 200 {object} response.ResultResponse
// @Router /v1/users/auth [get]
func (uc *UserController) AuthUser(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var userAuth dto.UserAuth
	if err := binder.BindAndValidate(c, &userAuth); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	userAuth.ID = userId.(uint)
	if err := uc.userService.AuthUser(&userAuth); err != nil {
		log.Error(constant.AuthFail, zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.AuthFail)
		return
	}
	response.Success(c, nil)
}
