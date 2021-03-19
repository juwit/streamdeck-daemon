package main

import (
	"fmt"
	"github.com/juwit/streamdeck-daemon/streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"os"
	"os/signal"
	"syscall"

	configService "github.com/juwit/streamdeck-daemon/config"
	"github.com/juwit/streamdeck-daemon/server"
)

func main() {

	// load configuration
	var config = configService.LoadConfiguration()

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
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		// shutting down streamdeck
		streamdeck.Shutdown()

		os.Exit(0)
	}()
}

