// 封装统一响应体

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResultResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
		Success: true,
	})
}

func Fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResultResponse{
		Status:  code,
		Message: message,
		Data:    nil,
		Success: false,
	})
}
