package server

import (
	"github.com/dkeng/cogo/controller/application"
	"github.com/dkeng/cogo/controller/config"
	"github.com/gin-gonic/gin"
)

// setRouter 设置路由
func setRouter(handler *gin.Engine) {
	handler.GET("/", func(c *gin.Context) {
		c.String(200, "welcome cogo server")
	})

	handler.GET("/applications", application.Get)

	handler.GET("/configs", config.Get)
}
