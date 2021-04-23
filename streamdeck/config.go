package streamdeck

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Brightness int     `json:"brightness"`
	InitialPage string `json:"initial_page"`
	Pages []Page       `json:"pages"`
}

func LoadConfiguration() *Config {
	homeDirectory, _ := os.UserHomeDir()
	configDirectory := "/.config/streamdeck"
	configFile := "/config.json"
	byteValue, err := ioutil.ReadFile(homeDirectory + configDirectory + configFile)

	if err != nil {
		panic(err)
	}

	var config Config

	json.Unmarshal(byteValue, &config)

	log.Print("Configuration file loaded")

	return &config
}


func (config *Config) GetPage(pageName string) *Page {
	for idx, _ := range config.Pages {
		var page = &config.Pages[idx]
		if page.Name == pageName {
			return page
		}
	}
	return nil
}

