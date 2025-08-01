package utils

import (
	"github.com/1255177148/golangTask4/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

var captchaStore = base64Captcha.DefaultMemStore

// GetCaptcha 生成图形验证码
func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, captchaStore)

	id, b64s, _, err := cp.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response.ResultResponse{
		Status: 200,
		Data: gin.H{
			"captchaId":  id,
			"captchaImg": b64s, // 前端直接 <img src="data:image/png;base64,...">
		},
		Success: true,
	})
}

func VerifyCaptcha(id, value string) bool {
	return captchaStore.Verify(id, value, true)
}
