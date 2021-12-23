package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRequest starts a server
func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/firstname", firstNamesHandler)
	router.HandleFunc("/lastname", lastNamesHandler)
	router.HandleFunc("/generation", generatedNamesHandler)

	http.Handle("/", router)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
