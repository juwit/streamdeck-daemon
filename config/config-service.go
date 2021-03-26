package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

type Config struct {
	Brightness int `json:"brightness"`
	InitialPage string `json:"initial_page"`
	Pages []Page `json:"pages"`
}

type Page struct {
	Name string `json:"name"`
	Buttons []Button `json:"buttons"`
}

type Button struct {
	// the index of the key to put the button on !
	Key int `json:"key"`
	// the icon to show on the button
	Icon string `json:"icon"`
	// switch on another page when the button pressed
	SwitchPage string `json:"switch_page"`
	// text to write when the button is pressed
	Write string `json:"write"`
	// command to execute when the button is pressed
	Command string `json:"command"`
}

func LoadConfiguration() Config {
	byteValue, err := ioutil.ReadFile("config.json")

	if err != nil {
		panic(err)
	}

	var config Config

	json.Unmarshal(byteValue, &config)

	fmt.Println("Loaded configuration file")
	fmt.Println(config)

	return config
}


func (config *Config) GetPage(pageName string) *Page {
	for _, page := range config.Pages {
		if page.Name == pageName {
			return &page
		}
	}
	return nil
}

func (page *Page) GetButton(index int) *Button {
	// find button on page definition
	for _, button := range page.Buttons {
		if button.Key == index {
			return &button
		}
	}
	return nil
}

/**
 * executes the key press on the button!
 */
func (button *Button) ExecCommand() {
	if button.Write != "" {
		go exec.Command("xdotool", "type", "--delay", "0", button.Write).Start()
	}

	if button.Command != "" {
		err := exec.Command("/bin/sh", "-c", button.Command).Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}