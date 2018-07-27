package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	start()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer close()
	log.Println("Server exiting")
}
