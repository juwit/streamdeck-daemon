package streamdeck

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Brightness int     `json:"brightness"`
	InitialPage string `json:"initial_page"`
	Pages []Page       `json:"pages"`
}

func LoadConfiguration() *Config {
	byteValue, err := ioutil.ReadFile("config.json")

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

