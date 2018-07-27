package server

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/dkeng/cogo/controller/application"
	"github.com/dkeng/cogo/controller/auth"
	"github.com/dkeng/cogo/controller/config"
	"github.com/gin-gonic/gin"
)

var (
	// the jwt middleware
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: auth.Authenticator,
		Authorizator:  auth.Authorizator,
		Unauthorized:  auth.Unauthorized,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
)

// setRouter 设置路由
func setRouter(handler *gin.Engine) {
	handler.GET("/", func(c *gin.Context) {
		c.String(200, "welcome cogo server")
	})

	handler.POST("/login", authMiddleware.LoginHandler)
	auth := handler.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		auth.GET("/applications", application.Get)
		auth.POST("/applications", application.Post)
		auth.DELETE("/applications/:id", application.Delete)

		auth.GET("/configs", config.Get)
		auth.POST("/configs", config.Post)
		auth.PUT("/configs/:id", config.Put)
		auth.DELETE("/configs/:id", config.Delete)
	}
}
