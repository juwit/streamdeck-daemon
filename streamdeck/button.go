package streamdeck

import (
	"github.com/juwit/streamdeck-daemon/libxdo"
	"log"
	"os/exec"
)

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

/**
 * executes the key press on the button!
 */
func (button *Button) ExecCommand() {
	if button.SwitchPage != "" {
		SwitchToPage(button.SwitchPage)
	}

	if button.Write != "" {
		libxdo.TypeKeys(button.Write)
	}

	if button.Command != "" {
		err := exec.Command("/bin/sh", "-c", button.Command).Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}