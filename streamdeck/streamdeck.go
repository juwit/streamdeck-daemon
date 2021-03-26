package streamdeck

import (
	"fmt"
	"github.com/magicmonkey/go-streamdeck"
)

var device *streamdeck.Device

var config Config

var currentPage *Page

func InitStreamdeck(loadedConfig Config){
	var err error
	device, err = streamdeck.Open()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found device %s\n", device.GetName())

	config = loadedConfig

	// switch to initial page
	fmt.Println("Loading initial page")
	switchToPage(config.InitialPage)

	device.ButtonPress(func(btnIndex int, device *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}

		var button = currentPage.GetButton(btnIndex)
		button.ExecCommand()
	})
}

func Shutdown(){
	device.ResetComms()
}

func switchToPage(pageName string) {
	// first, clearing buttons
	device.ClearButtons()

	currentPage = config.GetPage(pageName)

	for _, button := range currentPage.Buttons {
		renderButton(button)
	}
}

func renderButton(button Button) {
	if button.Icon != "" {
		device.WriteImageToButton(button.Key, button.Icon)
	}
}
