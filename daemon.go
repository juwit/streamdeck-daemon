package main

import (
	"fmt"
	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"image/color"
	"os"
	"os/signal"
	"syscall"
	"github.com/go-vgo/robotgo"

	"github.com/juwit/streamdeck-daemon/server"
)

func main() {
	device, err := streamdeck.Open()
	if err != nil {
		panic(err)
	}

	device.ClearButtons()

	fmt.Printf("Found device %s\n", device.GetName())

	// A simple yellow button in position 26
	device.WriteColorToButton(0, color.RGBA{255, 255, 0, 255})

	device.ButtonPress(func(btnIndex int, device *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Printf("Button %d pressed\n", btnIndex)

		robotgo.TypeStr("Hello World from streamdeck" + string(btnIndex))
	})

	setupShutdownHandler(device)

	server.StartHttpServer()
}

func setupShutdownHandler(device *streamdeck.Device) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		// shutting down streamdeck
		device.ResetComms()

		os.Exit(0)
	}()
}
