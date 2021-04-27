package main

import (
	"github.com/juwit/streamdeck-daemon/streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/juwit/streamdeck-daemon/server"
)

func main() {

	// load configuration
	var config = streamdeck.LoadConfiguration()

	// init streamdeck
	streamdeck.InitStreamdeck(config)

	setupShutdownHandler()

	// init http server
	server.StartHttpServer()
}

func setupShutdownHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("SIGTERM received - shutting down")

		// shutting down streamdeck
		streamdeck.Shutdown()

		os.Exit(0)
	}()
}

