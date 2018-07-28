package main

import (
	"log"
	"os"
	"os/signal"
	// 加载配置文件
	"github.com/dkeng/cogo/model"
	_ "github.com/dkeng/cogo/module/config"
	// 初始化存储
	"github.com/dkeng/cogo/module/store"
	"github.com/dkeng/cogo/server"
)

func init() {
	store.Start()
	model.AutoMigrate()
	server.Start()
}
func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer close()
	log.Println("Server exiting")
}

func close() {
	server.Close()
	store.Close()
}
