package streamdeck

type Page struct {
	Name    string   `json:"name"`
	Buttons []Button `json:"buttons"`
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

	renderButton(*button)
}
