package streamdeck

import (
	"github.com/magicmonkey/go-streamdeck"
	"log"
)

var device *streamdeck.Device

var config *Config

var CurrentPage *Page

func InitStreamdeck(loadedConfig *Config){
	var err error
	device, err = streamdeck.Open()
	if err != nil {
		panic(err)
	}

	log.Printf("Found device %s\n", device.GetName())

	config = loadedConfig

	// switch to initial page
	log.Println("Loading initial page")
	switchToPage(config.InitialPage)

	// init brightness
	device.SetBrightness(config.Brightness)

	device.ButtonPress(func(btnIndex int, device *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}

		var button = CurrentPage.GetButton(btnIndex)
		if button != nil {
			button.ExecCommand()
		}
	})
}

func Shutdown(){
	device.ResetComms()
}

func switchToPage(pageName string) {
	page := config.GetPage(pageName)

	if page == nil {
		log.Print("Page "+pageName+" not found")
		return
	}

	CurrentPage = page

	// render the page
	device.ClearButtons()
	for _, button := range CurrentPage.Buttons {
		renderButton(&button)
	}
}

func renderButton(button *Button) {
	if button.Icon != "" {
		device.WriteImageToButton(button.Key, button.Icon)
	}
}
