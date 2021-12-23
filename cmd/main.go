package main

import (
	"log"

	"task2/api"
	"task2/pkg/fileutil"
)

func main() {
	names, err := fileutil.ScanFile("./pkg/data_store/out/names.txt")
	if err != nil {
		log.Println(err)
	}
	firstNames, lastNames := fileutil.SeparateNames(names)
	if err != nil {
		log.Fatal(err)
	}
	firstNames = fileutil.DeleteDubs(firstNames)
	lastNames = fileutil.DeleteDubs(lastNames)

	err = fileutil.WriteInFile(firstNames, "./pkg/data_store/in/first_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = fileutil.WriteInFile(lastNames, "./pkg/data_store/in/last_names.txt")
	if err != nil {
		log.Fatal(err)
	}

	api.HandleRequest()
}
