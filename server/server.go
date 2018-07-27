package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	httpServer *http.Server
)

// Start 启动
func Start() {
	rest := gin.Default()
	rest.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome cogo server")
	})
	rest.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	httpServer = &http.Server{
		Addr:    viper.GetString("system.addr"),
		Handler: rest,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

// Close 关闭
func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("HttpServer Shutdown:", err)
	}
}
