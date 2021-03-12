package main

import (
	"fmt"
	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"image/color"
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
	})

	time.Sleep(5 * time.Second)

	server.StartHttpServer()
}
