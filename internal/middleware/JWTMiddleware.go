package middleware

import (
	"github.com/1255177148/golangTask4/internal/utils"
	"github.com/1255177148/golangTask4/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTMiddleware token校验中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.ResultResponse{
				Status:  http.StatusUnauthorized,
				Message: "未登录",
				Data:    nil,
				Success: false,
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userId, err := utils.ParseToken(tokenString, false)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ResultResponse{
				Status:  http.StatusUnauthorized,
				Message: "无效token",
				Success: false,
			})
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
