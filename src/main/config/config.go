package config

import "github.com/dkeng/pkg/logger"

// Init 配置文件初始化
func Init() {
	// 日志初始化
	logger.Init()
	// logger.RegisterSentry()
}
