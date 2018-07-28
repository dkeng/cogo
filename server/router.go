package server

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/dkeng/cogo/controller/application"
	"github.com/dkeng/cogo/controller/auth"
	"github.com/dkeng/cogo/controller/config"
	wgin "github.com/dkeng/pkg/context/gin"
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
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
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

		auth.GET("/applications", wgin.WrapControllerFunction(application.Get))
		auth.GET("/applications/:id", wgin.WrapControllerFunction(application.GetOne))
		auth.GET("/applications/:id/configs", wgin.WrapControllerFunction(application.GetConfigs))
		auth.POST("/applications", wgin.WrapControllerFunction(application.Post))
		auth.DELETE("/applications/:id", wgin.WrapControllerFunction(application.Delete))

		auth.GET("/configs", wgin.WrapControllerFunction(config.Get))
		auth.POST("/configs", wgin.WrapControllerFunction(config.Post))
		auth.PUT("/configs/:id", wgin.WrapControllerFunction(config.Put))
		auth.DELETE("/configs/:id", wgin.WrapControllerFunction(config.Delete))
	}
}
