package auth

import "github.com/gin-gonic/gin"

// Authenticator 身份验证
func Authenticator(userID string, password string, c *gin.Context) (interface{}, bool) {
	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {

		// &User{
		// 	UserName:  userId,
		// 	LastName:  "Bo-Yi",
		// 	FirstName: "Wu",
		// }
		return gin.H{
			"user_name":  userID,
			"last_name":  "Deyi",
			"first_name": "Xu",
		}, true
	}

	return nil, false
}

// Authorizator 授权
func Authorizator(user interface{}, c *gin.Context) bool {
	if v, ok := user.(string); ok && v == "admin" {
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
