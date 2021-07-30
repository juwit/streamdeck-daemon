package streamdeck

import (
	"github.com/magicmonkey/go-streamdeck"
	"image/color"
	"log"
)

var device *streamdeck.Device

var config *Config

var CurrentPage *Page

func StartStreamdeck() {
	var err error
	device, err = streamdeck.Open()
	if err != nil {
		panic(err)
	}

	log.Printf("Found device %s\n", device.GetName())

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

func InitStreamdeck(loadedConfig *Config){
	config = loadedConfig

	// switch to initial page
	log.Println("Loading initial page")
	SwitchToPage(config.InitialPage)

	// init brightness
	device.SetBrightness(config.Brightness)
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
	for idx, _ := range CurrentPage.Buttons {
		renderButton(&CurrentPage.Buttons[idx])
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
