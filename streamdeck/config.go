package streamdeck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Brightness int     `json:"brightness"`
	InitialPage string `json:"initial_page"`
	Pages []Page       `json:"pages"`
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

