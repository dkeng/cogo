package module

import (
	"log"
	"os"

	"github.com/dkeng/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	initConfigFile()
	initLog()
	initGin()
}

func initConfigFile() {
	viper.SetConfigFile("./cogo.toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误：%s", err.Error())
		os.Exit(-1)
	}
}
func initLog() {
	// 日志初始化
	logger.Init()
	// logger.RegisterSentry()
}

func initGin() {
	if viper.GetString("system.mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
