package server

import (
	"encoding/json"
	"github.com/go-chi/chi/middleware"
	"github.com/juwit/streamdeck-daemon/streamdeck"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func StartHttpServer(){
	log.Println("Starting http server")

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Post("/brightness/{value}", updateBrightness)

	router.Route("/pages/{name}", func(r chi.Router){
		r.Post("/", switchToPage)
	})

	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {

		var button streamdeck.Button

		err := json.NewDecoder(request.Body).Decode(&button)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		streamdeck.CurrentPage.AddButton(&button)
	})

	http.ListenAndServe(":8081", router)
}

func updateBrightness(writer http.ResponseWriter, request *http.Request){
	value := chi.URLParam(request, "value")
	brightness, _ := strconv.Atoi(value)
	streamdeck.ChangeBrightness(brightness)
}

func switchToPage(writer http.ResponseWriter, request *http.Request){
	page := chi.URLParam(request, "name")
	streamdeck.SwitchToPage(page)
}