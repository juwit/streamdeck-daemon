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
	streamdeck.StartStreamdeck()
	streamdeck.InitStreamdeck(config)

	setupShutdownHandler()
	setupSIGHUPHandler()

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

func setupSIGHUPHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP)
	go func() {
		for range c {
			log.Println("SIGHUP received - reloading configuration")
			var config = streamdeck.LoadConfiguration()
			streamdeck.InitStreamdeck(config)
		}
	}()
}
