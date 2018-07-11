package config

import (
	"github.com/dkeng/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Init 配置文件初始化
func Init() {
	// 日志初始化
	logger.Init()
	// logger.RegisterSentry()

	if viper.GetString("system.mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
