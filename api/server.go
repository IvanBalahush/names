package api

import (
	"github.com/gorilla/mux"
	"log"

	"net/http"
)

// HandleRequest starts a server
func HandleRequest(){
	router := mux.NewRouter()
	router.HandleFunc("/firstnames/", firstNamesHandler)
	router.HandleFunc("/lastnames/",lastNamesHandler)
	router.HandleFunc("/generations/", generatedNamesHandler)
	http.Handle("/",router)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}