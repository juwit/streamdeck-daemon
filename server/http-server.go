package server

import (
	"encoding/json"
	"fmt"
	"github.com/juwit/streamdeck-daemon/config"
	"github.com/juwit/streamdeck-daemon/streamdeck"
	"net/http"
)

func StartHttpServer(){
	fmt.Println("Starting http server")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		var button config.Button

		err := json.NewDecoder(request.Body).Decode(&button)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		streamdeck.RenderButton(button)

		writer.WriteHeader(http.StatusOK)
		fmt.Fprintln(writer, "OK")
	})

	http.ListenAndServe(":8081", nil)
}
