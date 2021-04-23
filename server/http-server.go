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

	router.Route("/pages/{pageName}", func(r chi.Router){
		r.Post("/", switchToPage)
		r.Post("/buttons/{key}", setupButton)
		r.Delete("/buttons/{key}", removeButton)
	})

	http.ListenAndServe(":8081", router)
}

func updateBrightness(writer http.ResponseWriter, request *http.Request){
	value := chi.URLParam(request, "value")
	brightness, _ := strconv.Atoi(value)
	streamdeck.ChangeBrightness(brightness)
}

func switchToPage(writer http.ResponseWriter, request *http.Request){
	page := chi.URLParam(request, "pageName")
	streamdeck.SwitchToPage(page)
}

func setupButton(writer http.ResponseWriter, request *http.Request){
	page := chi.URLParam(request, "pageName")

	keyStr := chi.URLParam(request, "key")
	key, _ := strconv.Atoi(keyStr)

	var button streamdeck.Button
	err := json.NewDecoder(request.Body).Decode(&button)
	button.Key = key

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	streamdeck.GetPage(page).AddButton(&button)
}

func removeButton(writer http.ResponseWriter, request *http.Request){
	page := chi.URLParam(request, "pageName")

	keyStr := chi.URLParam(request, "key")
	key, _ := strconv.Atoi(keyStr)


	streamdeck.GetPage(page).DeleteButton(key)
}