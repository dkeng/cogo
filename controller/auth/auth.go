package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Authenticator 身份验证
func Authenticator(username string, password string, c *gin.Context) (interface{}, bool) {

	if username == viper.GetString("system.username") && password == viper.GetString("system.password") {
		return gin.H{
			"user_name": username,
		}, true
	}

	return nil, false
}

// Authorizator 授权
func Authorizator(user interface{}, c *gin.Context) bool {
	if v, ok := user.(string); ok && v == viper.GetString("system.username") {
		return true
	}

	return false
}

// Unauthorized 未被授权的
func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
