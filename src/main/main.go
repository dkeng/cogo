package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dkeng/cogo/src/main/config"
	cogoStore "github.com/dkeng/cogo/src/store"
	"github.com/spf13/viper"
)

var (
	flagConfig string
	flagH      bool
)

func init() {
	flag.StringVar(&flagConfig, "config", "", "配置文件(.json,.yaml,.toml)")
	flag.BoolVar(&flagH, "help", false, "帮助")
	flag.Parse()
}

// checkFlag 检查Flag
func checkFlag() bool {
	if flagH {
		flag.PrintDefaults()
		return false
	}
	if flagConfig == "" {
		log.Fatalf("请使用-config指定配置文件")
		return false
	}
	return true
}

var (
	// 存储
	store *cogoStore.Store
)

func main() {
	// 检查输入参数
	if !checkFlag() {
		return
	}
	if !open() {
		return
	}
	defer close()
	fmt.Println("Hello,Cogo!!!")
}

// open 打开
func open() bool {
	// 初始化配置文件
	viper.SetConfigFile(flagConfig)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误：%s", err.Error())
		return false
	}
	config.Init()
	// 储存
	store = new(cogoStore.Store)
	err := store.Open()
	if err != nil {
		log.Fatalf("打开存储错误：%s", err.Error())
		return false
	}
	return true
}

// 关闭
func close() {
	store.Close()
}
