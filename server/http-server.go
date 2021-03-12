package server

import (
	"fmt"
	"net/http"
)

func StartHttpServer(){
	fmt.Println("Starting http server")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintln(writer, "OK")
	})

	http.ListenAndServe(":8081", nil)
}
