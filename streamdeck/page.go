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
