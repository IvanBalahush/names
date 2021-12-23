package api

import (
	"fmt"
	"log"
	"net/http"

	"task2/pkg/fileutil"
)

func firstNamesHandler(w http.ResponseWriter, r *http.Request) {
	err := fileutil.PageInput(w, "./pkg/data_store/in/first_names.txt")
	if err!=nil {
		log.Fatal(err)
	}
}

func lastNamesHandler(w http.ResponseWriter, r *http.Request){
	err := fileutil.PageInput(w, "./pkg/data_store/in/last_names.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func generatedNamesHandler(w http.ResponseWriter, r *http.Request) {
	firstNames, err := fileutil.ScanFile("./pkg/data_store/in/first_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	lastNames, err := fileutil.ScanFile("./pkg/data_store/in/last_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	generations := fileutil.GenerateNames(firstNames, lastNames, 3)
	for _, generation := range generations {
		_, err := fmt.Fprintln(w, generation)
		if err != nil {
			log.Fatal(err)
		}
	}
}
