package server

import (
	"encoding/json"
	"fmt"
	"github.com/juwit/streamdeck-daemon/streamdeck"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func StartHttpServer(){
	log.Println("Starting http server")

	router := chi.NewRouter()

	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {

		var button streamdeck.Button

		err := json.NewDecoder(request.Body).Decode(&button)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		streamdeck.CurrentPage.AddButton(&button)

		writer.WriteHeader(http.StatusOK)
		fmt.Fprintln(writer, "OK")
	})

	http.ListenAndServe(":8081", router)
}
