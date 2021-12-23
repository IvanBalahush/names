package api

import (
	"fmt"
	"log"
	"net/http"

	"task2/pkg/fileutil"
)

const fileDirectory = "./pkg/data_store/in/"

func firstNamesHandler(w http.ResponseWriter, r *http.Request) {
	err := fileutil.PageInput(w, fileDirectory + "first_names.txt")
	if err != nil {
		log.Print("HTTP 500 - Internal server error ")
		http.Error(w, "Internal server error", 500)
	}
}

func lastNamesHandler(w http.ResponseWriter, r *http.Request) {
	err := fileutil.PageInput(w, fileDirectory + "last_names.txt")
	if err != nil {
		log.Print("HTTP 500 - Internal server error ")
		http.Error(w, "Internal server error", 500)
	}
}

func generatedNamesHandler(w http.ResponseWriter, r *http.Request) {
	firstNames, err := fileutil.ScanFile(fileDirectory + "first_names.txt")
	if err != nil {
		log.Fatal(err)
	}

	lastNames, err := fileutil.ScanFile(fileDirectory + "last_names.txt")
	if err != nil {
		log.Fatal(err)
	}

	generations := fileutil.Generator(firstNames, lastNames, 3)
	_, err = fmt.Fprintln(w, generations)
	if err != nil {
		log.Print("HTTP 500 - Internal server error")
		http.Error(w, "Internal server error", 500)
	}
}
