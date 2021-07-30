package streamdeck

import "log"

type Page struct {
	Name    string   `json:"name"`
	Buttons []Button `json:"buttons"`
}

func (page *Page) GetButton(index int) *Button {
	// find button on page definition
	for idx, button := range page.Buttons {
		if button.Key == index {
			return &page.Buttons[idx]
		}
	}
	return nil
}

/**
 * Adds the given button to the page,
 * replacing if the button already exist for the same key
 */
func (page *Page) AddButton(button *Button){
	if page.GetButton(button.Key) == nil {
		// button doesn't exist
		page.Buttons = append(page.Buttons, *button)
	} else {
		// find in array, and replace
		for i := range page.Buttons {
			if page.Buttons[i].Key == button.Key {
				page.Buttons[i] = *button
			}
		}
	}

	if page == CurrentPage {
		renderButton(button)
	}
}

func (page *Page) DeleteButton(key int){
	var buttonIndex int = -1
	// find the index of the button
	for idx, button := range page.Buttons {
		if button.Key == key {
			buttonIndex = idx
		}
	}
	if buttonIndex == -1 {
		log.Printf("Button %d not found on page %s", key, page.Name)
		return
	}
	// remove the button from the array
	// putting the button at the very end
	page.Buttons[ len(page.Buttons) - 1 ], page.Buttons[buttonIndex] = page.Buttons[buttonIndex], page.Buttons[ len(page.Buttons) - 1 ]
	// slicing the array to remove the last value
	page.Buttons = page.Buttons[ : len(page.Buttons) - 1]

	// clear the button on the streamdeck if the page is the current page
	if page == CurrentPage {
		clearButton(key)
	}
}
