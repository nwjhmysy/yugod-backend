package main

import (
	"log"
	"os"
	"os/signal"
	"yugod-backend/app/boot"
)

func main() {

	boot.GinServer()

	WaitExit()
}

func WaitExit() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server.")
}
