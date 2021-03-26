package streamdeck

import (
	"fmt"
	configService "github.com/juwit/streamdeck-daemon/config"
	"github.com/magicmonkey/go-streamdeck"
	"log"
	"os/exec"
)

var device *streamdeck.Device

var config configService.Config

var currentPage *configService.Page

func InitStreamdeck(loadedConfig configService.Config){
	var err error
	device, err = streamdeck.Open()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found device %s\n", device.GetName())

	config = loadedConfig

	// switch to initial page
	fmt.Println("Loading initial page")
	SwitchToPage(config.InitialPage)

	device.ButtonPress(func(btnIndex int, device *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}

		// find button on page definition
		for _, button := range currentPage.Buttons {
			if button.Key == btnIndex {
				if button.Write != "" {
					go exec.Command("xdotool", "type", "--delay", "0", button.Write).Start()
				}

				if button.SwitchPage != "" {
					SwitchToPage(button.SwitchPage)
				}

				if button.Command != "" {
					err = exec.Command("/bin/sh", "-c", button.Command).Start()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	})
}

func Shutdown(){
	device.ResetComms()
}

func SwitchToPage(pageName string) {
	// first, clearing buttons
	device.ClearButtons()

	currentPage = config.GetPage(pageName)
	for _, button := range currentPage.Buttons {
		RenderButton(button)
	}
}

func RenderButton(button configService.Button) {
	if button.Icon != "" {
		device.WriteImageToButton(button.Key, button.Icon)
	}
}
