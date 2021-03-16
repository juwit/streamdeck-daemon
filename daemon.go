package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"os"
	"os/signal"
	"syscall"

	configService "github.com/juwit/streamdeck-daemon/config"
	"github.com/juwit/streamdeck-daemon/server"
)

var config configService.Config

var device *streamdeck.Device

var currentPage configService.Page

func main() {
	var err error
	device, err = streamdeck.Open()
	if err != nil {
		panic(err)
	}

	device.ClearButtons()
	device.ResetComms()

	fmt.Printf("Found device %s\n", device.GetName())

	// load configuration
	config = configService.LoadConfiguration()

	// set up brightness
	fmt.Printf("Setting up device brightness to %d\n", config.Brightness)
	// device.SetBrightness(config.Brightness)

	// switch to initial page
	fmt.Println("Loading initial page")
	SwitchToPage(config.InitialPage)

	// A simple yellow button in position 26
	// device.WriteColorToButton(0, color.RGBA{255, 255, 0, 255})

	device.ButtonPress(func(btnIndex int, device *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Button %d pressed\n", btnIndex)

		// find button on page definition
		for _, button := range currentPage.Buttons {
			if button.Key == btnIndex {
				if button.Write != "" {
					robotgo.PasteStr(button.Write)
				}
			}
		}
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

func SwitchToPage(pageName string) {
	// first, clearing buttons
	device.ClearButtons()
	for _, page := range config.Pages {
		if page.Name == pageName {
			// page found !
			fmt.Println("Page found")
			currentPage = page
			for _, button := range page.Buttons {
				RenderButton(button)
			}
		}
	}
}

func RenderButton(button configService.Button) {
	if button.Icon != "" {
		device.WriteImageToButton(button.Key, button.Icon)
	}
}