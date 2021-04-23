package streamdeck

import (
	"github.com/magicmonkey/go-streamdeck"
	"image/color"
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
	SwitchToPage(config.InitialPage)

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

func GetPage(pageName string) *Page {
	return config.GetPage(pageName)
}

func SwitchToPage(pageName string) {
	page := GetPage(pageName)

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

func clearButton(key int) {
	device.WriteColorToButton(key, color.Black)
}

func ChangeBrightness(value int){
	device.SetBrightness(value)
}
